package handlers

import (
	"errors"

	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"example.com/main-api-sdk/internal/validation"
)

// ResponseValidationHandler validates response data after deserialization.
// It ensures that the response data meets the expected schema and constraints.
// T is the response type, E is the error type.
type ResponseValidationHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewResponseValidationHandler creates a new response validation handler.
// Returns a handler that will validate response data against defined schemas.
func NewResponseValidationHandler[T any, E any]() *ResponseValidationHandler[T, E] {
	return &ResponseValidationHandler[T, E]{
		nextHandler: nil,
	}
}

// Handle validates the response data after it's been processed by the handler chain.
// Returns a validation error if the response data doesn't meet the schema, otherwise returns the response.
func (h *ResponseValidationHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	resp, handlerError := h.nextHandler.Handle(request)
	if handlerError != nil {
		return nil, handlerError
	}

	err := validation.ValidateData(resp.Data)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return resp, nil
}

// HandleStream passes through streaming requests without validation.
// Streaming responses are validated as they are consumed by the caller.
func (h *ResponseValidationHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	// Streaming responses are validated as they are consumed, so just pass through
	return h.nextHandler.HandleStream(request)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *ResponseValidationHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
