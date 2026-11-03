package handlers

import (
	"example.com/main-api-sdk/internal/clients/rest/httptransport"
)

// Handler defines the interface for request processing in the handler chain.
// Each handler can process both regular and streaming requests.
// T is the response type, E is the error type.
type Handler[T any, E any] interface {
	Handle(req httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E])
	HandleStream(req httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E])
	SetNext(handler Handler[T, E])
}

// HandlerChain manages a chain of handlers for processing requests.
// Implements the chain of responsibility pattern for request/response processing.
type HandlerChain[T any, E any] struct {
	head Handler[T, E]
	tail Handler[T, E]
}

// BuildHandlerChain creates a new empty handler chain.
// Handlers can be added using AddHandler to build the processing pipeline.
func BuildHandlerChain[T any, E any]() *HandlerChain[T, E] {
	return &HandlerChain[T, E]{}
}

// AddHandler appends a handler to the end of the chain.
// Returns the chain for method chaining. Handlers execute in the order they are added.
func (chain *HandlerChain[T, E]) AddHandler(handler Handler[T, E]) *HandlerChain[T, E] {
	if chain.head == nil {
		chain.head = handler
		chain.tail = handler
		return chain
	}

	chain.tail.SetNext(handler)
	chain.tail = handler

	return chain
}

// CallApi processes a regular HTTP request through the handler chain.
// Returns the processed response or an error response.
func (chain *HandlerChain[T, E]) CallApi(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	return chain.head.Handle(request)
}

// StreamApi processes a streaming HTTP request through the handler chain.
// Returns a stream for consuming response chunks or an error response.
func (chain *HandlerChain[T, E]) StreamApi(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	return chain.head.HandleStream(request)
}
