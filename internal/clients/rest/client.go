package rest

import (
	"torbox-sdk-go/internal/clients/rest/handlers"
	"torbox-sdk-go/internal/clients/rest/hooks"
	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type RestClient[T any] struct {
	handlers *handlers.HandlerChain[T]
}

func NewRestClient[T any](config torboxapiconfig.Config) *RestClient[T] {
	retryHandler := handlers.NewRetryHandler[T]()
	bearerTokenHandler := handlers.NewAccessTokenHandler[T]()
	responseValidationHandler := handlers.NewResponseValidationHandler[T]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T]()
	requestValidationHandler := handlers.NewRequestValidationHandler[T]()
	hookHandler := handlers.NewHookHandler[T](hooks.NewCustomHook())
	terminatingHandler := handlers.NewTerminatingHandler[T]()

	handlers := handlers.BuildHandlerChain[T]().
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(responseValidationHandler).
		AddHandler(unmarshalHandler).
		AddHandler(requestValidationHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T]{
		handlers: handlers,
	}
}

func (client *RestClient[T]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return client.handlers.CallApi(request)
}
