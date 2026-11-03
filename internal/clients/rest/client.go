package rest

import (
	"example.com/main-api-sdk/internal/clients/rest/handlers"
	"example.com/main-api-sdk/internal/clients/rest/hooks"
	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"example.com/main-api-sdk/mainapisdkconfig"
)

// RestClient is a generic HTTP client that handles API requests through a chain of handlers.
// It supports both regular and streaming requests with type-safe response handling.
// T is the response type, E is the error type.
type RestClient[T any, E any] struct {
	handlers *handlers.HandlerChain[T, E]
}

// NewRestClient creates a new REST client with the configured handler chain.
// Initializes all handlers in the correct order for request processing.
func NewRestClient[T any, E any](config mainapisdkconfig.Config, hook hooks.Hook) *RestClient[T, E] {
	retryHandler := handlers.NewRetryHandler[T, E]()
	bearerTokenHandler := handlers.NewAccessTokenHandler[T, E]()
	responseValidationHandler := handlers.NewResponseValidationHandler[T, E]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T, E]()
	requestValidationHandler := handlers.NewRequestValidationHandler[T, E]()
	hookHandler := handlers.NewHookHandler[T, E](hook)
	terminatingHandler := handlers.NewTerminatingHandler[T, E]()

	handlers := handlers.BuildHandlerChain[T, E]().
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(responseValidationHandler).
		AddHandler(unmarshalHandler).
		AddHandler(requestValidationHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T, E]{
		handlers: handlers,
	}
}

// Call executes a regular HTTP request through the handler chain.
// Returns the response with deserialized data or an error response.
func (client *RestClient[T, E]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	return client.handlers.CallApi(request)
}

// Stream executes a streaming HTTP request through the handler chain.
// Returns a stream for consuming response chunks or an error response.
func (client *RestClient[T, E]) Stream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	return client.handlers.StreamApi(request)
}
