package handlers

import (
	"errors"

	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/internal/validation"
)

type ResponseValidationHandler[T any] struct {
	nextHandler Handler[T]
}

func NewResponseValidationHandler[T any]() *ResponseValidationHandler[T] {
	return &ResponseValidationHandler[T]{
		nextHandler: nil,
	}
}

func (h *ResponseValidationHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp, handlerError := h.nextHandler.Handle(request)
	if handlerError != nil {
		return nil, handlerError
	}

	err := validation.ValidateData(resp.Data)
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	return resp, nil
}

func (h *ResponseValidationHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
