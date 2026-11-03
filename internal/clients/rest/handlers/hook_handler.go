package handlers

import (
	"errors"

	"example.com/main-api-sdk/internal/clients/rest/hooks"
	"example.com/main-api-sdk/internal/clients/rest/httptransport"
)

// HookHandler executes custom hook callbacks before and after request processing.
// It allows users to intercept and modify requests, responses, and errors through the hook interface.
// T is the response type, E is the error type.
type HookHandler[T any, E any] struct {
	nextHandler Handler[T, E]
	hook        hooks.Hook
}

// NewHookHandler creates a new hook handler with the provided hook implementation.
// Returns a handler that will execute BeforeRequest, AfterResponse, and OnError callbacks.
func NewHookHandler[T any, E any](hook hooks.Hook) *HookHandler[T, E] {
	return &HookHandler[T, E]{
		hook:        hook,
		nextHandler: nil,
	}
}

// Handle processes a regular request through the hook lifecycle.
// Executes BeforeRequest, passes the request through the chain, then executes AfterResponse or OnError.
func (h *HookHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	clonedReq := request.Clone()
	hookReq := h.hook.BeforeRequest(&clonedReq, clonedReq.Config.HookParams)

	nextRequest, ok := hookReq.(*httptransport.Request)
	if !ok {
		err := errors.New("hook returned invalid request")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	response, err := h.nextHandler.Handle(*nextRequest)
	if err != nil && err.IsHTTPError {
		clonedError := err.Clone()
		hookError := h.hook.OnError(hookReq, &clonedError, clonedReq.Config.HookParams)
		nextError, ok := hookError.(*httptransport.ErrorResponse[E])
		if !ok {
			err := errors.New("hook returned invalid error")
			return nil, httptransport.NewErrorResponse[E](err, nil)
		}

		return nil, nextError
	} else if err != nil {
		return nil, err
	}

	clonedResp := response.Clone()
	hookResp := h.hook.AfterResponse(hookReq, &clonedResp, clonedReq.Config.HookParams)
	nextResponse, ok := hookResp.(*httptransport.Response[T])
	if !ok {
		err := errors.New("hook returned invalid response")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return nextResponse, nil
}

// HandleStream processes a streaming request by executing the BeforeRequest hook.
// Returns the stream from the next handler after the request hook is applied.
func (h *HookHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	clonedReq := request.Clone()
	hookReq := h.hook.BeforeRequest(&clonedReq, clonedReq.Config.HookParams)

	nextRequest, ok := hookReq.(*httptransport.Request)
	if !ok {
		err := errors.New("hook returned invalid request")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return h.nextHandler.HandleStream(*nextRequest)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *HookHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
