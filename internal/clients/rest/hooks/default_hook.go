package hooks

type DefaultHook struct{}

func NewDefaultHook() Hook {
	return &DefaultHook{}
}

func (h *DefaultHook) BeforeRequest(req Request, params map[string]string) Request {
	return req
}

func (h *DefaultHook) AfterResponse(req Request, resp Response, params map[string]string) Response {
	return resp
}

func (h *DefaultHook) OnError(req Request, resp ErrorResponse, params map[string]string) ErrorResponse {
	return resp
}
