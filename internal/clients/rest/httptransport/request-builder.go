package httptransport

import (
	"context"

	"example.com/main-api-sdk/mainapisdkconfig"
)

// RequestBuilder provides a fluent interface for constructing HTTP requests.
// Uses the builder pattern to incrementally configure request parameters, headers, body, and options.
type RequestBuilder struct {
	request *Request
}

// NewRequestBuilder creates a new request builder with default settings.
// Initializes empty maps for headers, query params, and path params.
func NewRequestBuilder() *RequestBuilder {
	rb := &RequestBuilder{
		request: &Request{
			Headers:     make(map[string]string),
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		},
	}
	rb.request.SetHeader("User-Agent", "postman-codegen/3.20.0 mainapisdk/1.0.7 (go)")
	return rb
}

// WithContext sets the request context for cancellation and timeout control.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithContext(ctx context.Context) *RequestBuilder {
	rb.request.Context = ctx
	return rb
}

// WithMethod sets the HTTP method (GET, POST, PUT, DELETE, etc.).
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithMethod(method string) *RequestBuilder {
	rb.request.Method = method
	return rb
}

// WithPath sets the URL path for the request.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithPath(path string) *RequestBuilder {
	rb.request.Path = path
	return rb
}

// AddHeader adds a header key-value pair to the request.
// Returns the builder for method chaining.
func (rb *RequestBuilder) AddHeader(key string, value string) *RequestBuilder {
	rb.request.SetHeader(key, value)
	return rb
}

// AddPathParam adds a path parameter to be substituted into the URL path.
// Returns the builder for method chaining.
func (rb *RequestBuilder) AddPathParam(key string, value any) *RequestBuilder {
	rb.request.SetPathParam(key, value)
	return rb
}

// WithOptions sets the request options containing query/header parameters.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithOptions(options any) *RequestBuilder {
	rb.request.Options = options
	return rb
}

// WithBody sets the request body to be serialized.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithBody(body any) *RequestBuilder {
	rb.request.Body = body
	return rb
}

// WithConfig sets the SDK configuration including base URL, timeout, and auth credentials.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithConfig(config mainapisdkconfig.Config) *RequestBuilder {
	rb.request.Config = config
	return rb
}

// WithContentType sets the content type for serializing the request body.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithContentType(contentType ContentType) *RequestBuilder {
	rb.request.ContentType = contentType
	return rb
}

// WithResponseContentType sets the expected content type for deserializing the response.
// Returns the builder for method chaining.
func (rb *RequestBuilder) WithResponseContentType(contentType ContentType) *RequestBuilder {
	rb.request.ResponseContentType = contentType
	return rb
}

// WithSecuritySchemes records the auth schemes the operation accepts.
// Each request-chain auth handler consults this list via
// ShouldApplyAuthScheme to decide whether to apply its credentials —
// preventing the SDK from sending an apiKey header on an OAuth-only
// operation (and vice-versa) when multiple credentials are configured.
// Passing nil opts out of per-operation gating: every configured
// handler runs (backward-compat for hand-constructed requests).
func (rb *RequestBuilder) WithSecuritySchemes(schemes []AuthScheme) *RequestBuilder {
	rb.request.SecuritySchemes = schemes
	return rb
}

// Build returns the fully configured Request.
// Should be called after all configuration methods to obtain the final request.
func (rb *RequestBuilder) Build() *Request {
	return rb.request
}
