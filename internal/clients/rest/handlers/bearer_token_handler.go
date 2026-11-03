package handlers

import (
	"errors"
	"fmt"

	"example.com/main-api-sdk/internal/clients/rest/httptransport"
)

// AccessTokenHandler adds a Bearer token to the Authorization header for authentication.
// It injects the configured access token into requests that require token-based auth.
// T is the response type, E is the error type.
type AccessTokenHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewAccessTokenHandler creates a new Bearer token authentication handler.
// Returns a handler that will inject the access token as a Bearer token in the Authorization header.
func NewAccessTokenHandler[T any, E any]() *AccessTokenHandler[T, E] {
	return &AccessTokenHandler[T, E]{
		nextHandler: nil,
	}
}

// prepareRequest clones the request and adds the Bearer token header if configured.
// Formats the token as "Bearer <token>" and sets the Authorization header.
func (h *AccessTokenHandler[T, E]) prepareRequest(request httptransport.Request) (httptransport.Request, error) {
	if h.nextHandler == nil {
		return httptransport.Request{}, errors.New("Handler chain terminated without terminating handler")
	}

	nextRequest := request.Clone()

	// The bearer handler stamps `Authorization: Bearer <AccessToken>` for two cases:
	//
	//  1. The operation accepts an HTTP bearer scheme (the obvious case).
	//  2. The operation accepts OAuth2 but no in-runtime OAuth handler covers it.
	//     Only `client_credentials` is driven by the embedded token manager; for
	//     `authorization_code`/`device_code`/`implicit`/`openIdConnect` the
	//     consumer obtains the token out-of-band and configures it on the SDK.
	//     Without this branch, OAuth2-only specs whose flow the SDK can't
	//     drive end up sending no Authorization header at all.
	//
	// Legacy hand-built consumers leave `SecuritySchemes == nil`, which both
	// `ShouldApplyAuthScheme` calls treat as "apply" — same behaviour as before
	// per-operation gating landed.
	if !httptransport.ShouldApplyAuthScheme(request.SecuritySchemes, httptransport.AuthSchemeBearer) &&
		!httptransport.ShouldApplyAuthScheme(request.SecuritySchemes, httptransport.AuthSchemeOAuth2) {
		return nextRequest, nil
	}

	if request.Config.AccessToken != nil {
		nextRequest.SetHeader("Authorization", fmt.Sprintf("Bearer %s", *request.Config.AccessToken))
	}

	return nextRequest, nil
}

// Handle processes a regular request by adding the Bearer token if configured.
// Returns the response from the next handler after authentication is applied.
func (h *AccessTokenHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	nextRequest, err := h.prepareRequest(request)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}
	return h.nextHandler.Handle(nextRequest)
}

// HandleStream processes a streaming request by adding the Bearer token if configured.
// Returns the stream from the next handler after authentication is applied.
func (h *AccessTokenHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	nextRequest, err := h.prepareRequest(request)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}
	return h.nextHandler.HandleStream(nextRequest)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *AccessTokenHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
