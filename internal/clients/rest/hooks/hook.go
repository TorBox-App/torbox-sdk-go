package hooks

// Hook defines the interface for intercepting and modifying requests, responses, and errors.
// Implementations can inspect, modify, or log HTTP interactions at key points in the request lifecycle.
type Hook interface {
	BeforeRequest(req Request, params map[string]string) Request
	AfterResponse(req Request, resp Response, params map[string]string) Response
	OnError(req Request, resp ErrorResponse, params map[string]string) ErrorResponse
}

// Request defines the interface for accessing and modifying HTTP request properties within hooks.
type Request interface {
	GetMethod() string
	SetMethod(method string)
	GetBaseURL() string
	SetBaseURL(baseURL string)
	GetPath() string
	SetPath(path string)
	GetPathParam(param string) string
	SetPathParam(param string, value any)
	GetHeader(header string) string
	SetHeader(header string, value string)
	GetQueryParam(header string) string
	SetQueryParam(header string, value string)
	GetOptions() any
	SetOptions(options any)
	GetBody() any
	SetBody(body any)
}

// Response defines the interface for accessing and modifying HTTP response properties within hooks.
type Response interface {
	GetStatusCode() int
	SetStatusCode(statusCode int)
	GetHeader(header string) string
	SetHeader(header string, value string)
	GetBody() []byte
	SetBody(body []byte)
}

// ErrorResponse defines the interface for accessing and modifying HTTP error responses within hooks.
type ErrorResponse interface {
	Error() string
	GetError() error
	GetStatusCode() int
	SetStatusCode(statusCode int)
	GetHeader(header string) string
	SetHeader(header string, value string)
	GetBody() []byte
	SetBody(body []byte)
}
