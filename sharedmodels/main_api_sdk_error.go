package sharedmodels

import (
	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"net/http"
)

// MainAPISDKError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type MainAPISDKError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata MainAPISDKErrorMetadata
}

// MainAPISDKErrorMetadata contains HTTP metadata associated with an error response.
type MainAPISDKErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewMainAPISDKError creates a new MainAPISDKError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewMainAPISDKError[T any](transportError *httptransport.ErrorResponse[T]) *MainAPISDKError[T] {
	return &MainAPISDKError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: MainAPISDKErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *MainAPISDKError[T]) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As to traverse the chain.
func (e *MainAPISDKError[T]) Unwrap() error {
	return e.Err
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *MainAPISDKError[T]) GetData() *T {
	return e.Data
}

// GetBody returns the raw response body bytes from the error response.
// Returns nil if no response body was received.
func (e *MainAPISDKError[T]) GetBody() []byte {
	return e.Body
}
