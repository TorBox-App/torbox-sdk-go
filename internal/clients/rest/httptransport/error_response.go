package httptransport

import (
	"fmt"
	"net/http"
)

// ErrorResponse represents a failed HTTP request with error details and optional response data.
// It distinguishes between HTTP errors (with status codes) and other errors (network, serialization, etc.).
// Data is a pointer to allow nil when unmarshaling fails.
type ErrorResponse[T any] struct {
	Err         error
	IsHTTPError bool
	StatusCode  int
	Headers     map[string]string
	Body        []byte
	Data        *T
	Raw         *http.Response
}

// NewErrorResponse creates an ErrorResponse from an error and an optional response.
// If response is provided, sets IsHTTPError to true and includes response metadata.
func NewErrorResponse[T any](err error, resp *Response[T]) *ErrorResponse[T] {
	if resp == nil {
		return &ErrorResponse[T]{
			Err:         err,
			IsHTTPError: false,
		}
	}

	return &ErrorResponse[T]{
		Err:         err,
		IsHTTPError: true,
		StatusCode:  resp.StatusCode,
		Headers:     resp.Headers,
		Body:        resp.Body,
		Raw:         resp.Raw,
	}
}

// Clone creates a deep copy of the ErrorResponse including headers.
// Returns a new ErrorResponse with copied values.
func (r *ErrorResponse[T]) Clone() ErrorResponse[T] {
	if r == nil {
		return ErrorResponse[T]{}
	}

	clone := *r
	clone.Headers = make(map[string]string)
	for header, value := range r.Headers {
		clone.Headers[header] = value
	}
	return clone
}

// Error implements the error interface, returning the error message string.
func (r *ErrorResponse[T]) Error() string {
	return fmt.Sprintf("%s", r.Err)
}

// GetError returns the underlying error.
func (r *ErrorResponse[T]) GetError() error {
	return r.Err
}

func (r *ErrorResponse[T]) GetStatusCode() int {
	return r.StatusCode
}

func (r *ErrorResponse[T]) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

func (r *ErrorResponse[T]) GetHeaders() map[string]string {
	return r.Headers
}

func (r *ErrorResponse[T]) GetHeader(header string) string {
	return r.Headers[header]
}

func (r *ErrorResponse[T]) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *ErrorResponse[T]) GetBody() []byte {
	return r.Body
}

func (r *ErrorResponse[T]) SetBody(body []byte) {
	r.Body = body
}
