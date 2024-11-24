package httptransport

import "fmt"

type ErrorResponse[T any] struct {
	Err         error
	IsHttpError bool
	StatusCode  int
	Headers     map[string]string
	Body        []byte
	Data        T
}

func NewErrorResponse[T any](err error, resp *Response[T]) *ErrorResponse[T] {
	if resp == nil {
		return &ErrorResponse[T]{
			Err:         err,
			IsHttpError: false,
		}
	}

	return &ErrorResponse[T]{
		Err:         err,
		IsHttpError: true,
		StatusCode:  resp.StatusCode,
		Headers:     resp.Headers,
		Body:        resp.Body,
		Data:        resp.Data,
	}
}

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

func (r *ErrorResponse[T]) Error() string {
	return fmt.Sprintf("%s", r.Err)
}

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
