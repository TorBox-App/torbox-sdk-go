package hooks

// DefaultHook is a no-op hook implementation that passes through all requests and responses unchanged.
// Used as the default hook when no custom hook is provided.
type DefaultHook struct{}

// NewDefaultHook creates a new default hook that performs no modifications.
// Returns a hook that passes through all requests, responses, and errors unchanged.
func NewDefaultHook() Hook {
	return &DefaultHook{}
}

// BeforeRequest passes the request through unchanged.
func (h *DefaultHook) BeforeRequest(req Request, params map[string]string) Request {
	return req
}

// AfterResponse passes the response through unchanged.
func (h *DefaultHook) AfterResponse(req Request, resp Response, params map[string]string) Response {
	return resp
}

// OnError passes the error response through unchanged.
func (h *DefaultHook) OnError(req Request, resp ErrorResponse, params map[string]string) ErrorResponse {
	return resp
}
