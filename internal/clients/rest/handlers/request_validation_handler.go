package handlers

import (
	"errors"

	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"example.com/main-api-sdk/internal/validation"
)

// RequestValidationHandler validates request body and options before sending.
// It ensures that all required fields are present and values meet the specified constraints.
// T is the response type, E is the error type.
type RequestValidationHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewRequestValidationHandler creates a new request validation handler.
// Returns a handler that will validate request body and options against defined schemas.
func NewRequestValidationHandler[T any, E any]() *RequestValidationHandler[T, E] {
	return &RequestValidationHandler[T, E]{
		nextHandler: nil,
	}
}

// Handle validates the request body and options before passing to the next handler.
// Returns a validation error if any constraints are violated, otherwise continues the chain.
func (h *RequestValidationHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	err := validation.ValidateData(request.Body)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	err = validation.ValidateData(request.Options)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return h.nextHandler.Handle(request)
}

// HandleStream validates the request body and options before initiating a stream.
// Returns a validation error if any constraints are violated, otherwise continues the chain.
func (h *RequestValidationHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	err := validation.ValidateData(request.Body)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	err = validation.ValidateData(request.Options)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	return h.nextHandler.HandleStream(request)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *RequestValidationHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
