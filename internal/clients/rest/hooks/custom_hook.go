package hooks

import (
	"fmt"
)

type CustomHook struct{}

func NewCustomHook() Hook {
	return &CustomHook{}
}

func (h *CustomHook) BeforeRequest(req Request, params map[string]string) Request {
	fmt.Printf("BeforeRequest: %#v\n", req)
	return req
}

func (h *CustomHook) AfterResponse(req Request, resp Response, params map[string]string) Response {
	fmt.Printf("AfterResponse: %#v\n", resp)
	return resp
}

func (h *CustomHook) OnError(req Request, resp ErrorResponse, params map[string]string) ErrorResponse {
	fmt.Printf("On Error: %#v\n", resp)
	return resp
}
