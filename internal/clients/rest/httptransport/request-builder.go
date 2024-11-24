package httptransport

import (
	"context"

	"torbox-sdk-go/pkg/torboxapiconfig"
)

type RequestBuilder struct {
	request *Request
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: &Request{
			Headers:     make(map[string]string),
			QueryParams: make(map[string]string),
			PathParams:  make(map[string]string),
		},
	}
}

func (rb *RequestBuilder) WithContext(ctx context.Context) *RequestBuilder {
	rb.request.Context = ctx
	return rb
}

func (rb *RequestBuilder) WithMethod(method string) *RequestBuilder {
	rb.request.Method = method
	return rb
}

func (rb *RequestBuilder) WithPath(path string) *RequestBuilder {
	rb.request.Path = path
	return rb
}

func (rb *RequestBuilder) AddHeader(key string, value string) *RequestBuilder {
	rb.request.SetHeader(key, value)
	return rb
}

func (rb *RequestBuilder) AddPathParam(key string, value any) *RequestBuilder {
	rb.request.SetPathParam(key, value)
	return rb
}

func (rb *RequestBuilder) WithOptions(options any) *RequestBuilder {
	rb.request.Options = options
	return rb
}

func (rb *RequestBuilder) WithBody(body any) *RequestBuilder {
	rb.request.Body = body
	return rb
}

func (rb *RequestBuilder) WithConfig(config torboxapiconfig.Config) *RequestBuilder {
	rb.request.Config = config
	return rb
}

func (rb *RequestBuilder) WithContentType(contentType ContentType) *RequestBuilder {
	rb.request.ContentType = contentType
	return rb
}

func (rb *RequestBuilder) WithResponseContentType(contentType ContentType) *RequestBuilder {
	rb.request.ResponseContentType = contentType
	return rb
}

func (rb *RequestBuilder) Build() *Request {
	return rb.request
}
