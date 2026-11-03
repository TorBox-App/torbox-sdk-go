package handlers

import (
	"errors"
	"fmt"
	"reflect"

	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"example.com/main-api-sdk/internal/contenttypes"
)

// UnmarshalHandler deserializes response bodies based on content type.
// It supports JSON, form data, text, and binary content types, converting raw bytes to typed data.
// T is the response type, E is the error type.
type UnmarshalHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewUnmarshalHandler creates a new unmarshaling handler.
// Returns a handler that will deserialize response bodies according to the requested content type.
func NewUnmarshalHandler[T any, E any]() *UnmarshalHandler[T, E] {
	return &UnmarshalHandler[T, E]{
		nextHandler: nil,
	}
}

// Handle deserializes the response body from the next handler based on the response content type.
// Supports JSON, form-urlencoded, multipart form data, text, and binary formats.
func (h *UnmarshalHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	resp, handlerError := h.nextHandler.Handle(request)
	if handlerError != nil {
		// Try to unmarshal error response body into the error type E
		if handlerError.IsHTTPError && len(handlerError.Body) > 0 {
			unmarshaledError := h.tryUnmarshalErrorBody(handlerError)
			if unmarshaledError != nil {
				return nil, unmarshaledError
			}
		}
		return nil, handlerError
	}

	if len(resp.Body) == 0 {
		var zeroValue T
		resp.Data = zeroValue
		return resp, nil
	}

	contentTypesToTry := []httptransport.ContentType{
		httptransport.ContentTypeJSON,
		httptransport.ContentTypeXML,
		httptransport.ContentTypeMultipartFormData,
		httptransport.ContentTypeFormUrlEncoded,
		httptransport.ContentTypeText,
		httptransport.ContentTypeBinary,
	}

	var firstErr error
	for _, contentType := range contentTypesToTry {
		tempTarget := new(T)
		err := unmarshalWithContentType(resp.Body, tempTarget, contentType)
		if err == nil {
			resp.Data = *tempTarget
			return resp, nil
		}
		// Keep the first error — JSON is tried first and gives the most meaningful message.
		if firstErr == nil {
			firstErr = err
		}
	}

	if firstErr == nil {
		firstErr = fmt.Errorf("failed to unmarshal response with any content type")
	}

	return nil, httptransport.NewErrorResponse[E](firstErr, nil)
}

// tryUnmarshalErrorBody attempts to unmarshal the error response body into the error type E.
// Returns a new ErrorResponse with unmarshaled data if successful, or nil if unmarshaling fails.
// For []byte error types, FromBinary will handle copying the body directly.
func (h *UnmarshalHandler[T, E]) tryUnmarshalErrorBody(err *httptransport.ErrorResponse[E]) *httptransport.ErrorResponse[E] {
	// Normalize empty body: convert nil to empty slice for consistent handling
	body := err.Body
	if body == nil {
		body = []byte{}
	}

	contentTypesToTry := []httptransport.ContentType{
		httptransport.ContentTypeJSON,
		httptransport.ContentTypeXML,
		httptransport.ContentTypeMultipartFormData,
		httptransport.ContentTypeFormUrlEncoded,
		httptransport.ContentTypeText,
		httptransport.ContentTypeBinary,
	}

	for _, contentType := range contentTypesToTry {
		tempTarget := new(E)
		unmarshalErr := unmarshalWithContentType(body, tempTarget, contentType)
		if unmarshalErr == nil {
			// Check if the unmarshaled value is a zero value (e.g., nil slice for []byte)
			// For zero values, we want Data to be nil, not a pointer to zero value
			if isEmptyValue(*tempTarget) {
				// Return nil to indicate unmarshaling didn't produce meaningful data
				// This allows callers to distinguish between "unmarshaled successfully" and "no data"
				return nil
			}
			// Successfully unmarshaled, create a new error response with the unmarshaled data
			unmarshaledError := err.Clone()
			unmarshaledError.Data = tempTarget
			// Ensure Body is also normalized (convert nil to empty slice)
			if unmarshaledError.Body == nil {
				unmarshaledError.Body = []byte{}
			}
			return &unmarshaledError
		}
	}

	// If unmarshaling fails for all content types, Data remains nil (pointer to E)
	return nil
}

// isEmptyValue checks if a value is considered "empty" (zero value).
// For slices, this means nil or empty slice. For other types, it's the zero value.
func isEmptyValue[T any](v T) bool {
	var zero T
	if reflect.DeepEqual(v, zero) {
		return true
	}
	// Also check if it's an empty slice (length 0)
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Slice {
		return vValue.Len() == 0
	}
	return false
}

func unmarshalWithContentType[T any](body []byte, target *T, contentType httptransport.ContentType) error {
	if contentType == httptransport.ContentTypeJSON {
		return contenttypes.FromJSON(body, target)
	} else if contentType == httptransport.ContentTypeXML {
		return contenttypes.FromXML(body, target)
	} else if contentType == httptransport.ContentTypeFormUrlEncoded {
		return contenttypes.FromFormUrlEncoded(body, target)
	} else if contentType == httptransport.ContentTypeMultipartFormData {
		return contenttypes.FromFormData(body, target)
	} else if contentType == httptransport.ContentTypeText {
		return contenttypes.FromText[T](body, target)
	} else if contentType == httptransport.ContentTypeBinary {
		return contenttypes.FromBinary(body, target)
	} else {
		return contenttypes.FromBinary(body, target)
	}
}

// HandleStream passes through streaming requests without unmarshaling.
// Streaming responses handle their own unmarshaling as chunks are consumed.
func (h *UnmarshalHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return h.nextHandler.HandleStream(request)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *UnmarshalHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
