package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"example.com/main-api-sdk/internal/clients/rest/httptransport"
)

// TerminatingHandler is the final handler in the chain that executes the actual HTTP request.
// It creates an HTTP client, sends the request, and converts the http.Response to a transport response.
// T is the response type, E is the error type.
type TerminatingHandler[T any, E any] struct{}

// NewTerminatingHandler creates a new terminating handler.
// Returns the handler that will execute HTTP requests and should be the last in the chain.
func NewTerminatingHandler[T any, E any]() *TerminatingHandler[T, E] {
	return &TerminatingHandler[T, E]{}
}

// Handle executes the HTTP request using the standard http.Client and returns the response.
// This is the final handler in the chain that performs the actual network call.
func (h *TerminatingHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	requestClone := request.Clone()

	client := http.Client{}
	if requestClone.Config.Timeout != 0 {
		client.Timeout = requestClone.Config.Timeout
	}

	req, err := requestClone.CreateHTTPRequest()
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	transportResponse, responseErr := httptransport.NewResponse[T](resp)
	if responseErr != nil {
		return nil, httptransport.NewErrorResponse[E](responseErr, nil)
	}

	if transportResponse.StatusCode >= 400 {
		err := fmt.Errorf("HTTP request failed with status code %d", transportResponse.StatusCode)
		errorResponse := &httptransport.Response[E]{
			StatusCode: transportResponse.StatusCode,
			Headers:    transportResponse.Headers,
			Body:       transportResponse.Body,
			Raw:        transportResponse.Raw,
		}
		return nil, httptransport.NewErrorResponse[E](err, errorResponse)
	}

	return transportResponse, nil
}

// HandleStream executes a streaming HTTP request and returns a stream for consuming the response.
// Creates an HTTP client, sends the request, and wraps the response body in a stream for chunk-by-chunk processing.
func (h *TerminatingHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	requestClone := request.Clone()

	client := http.Client{}
	if requestClone.Config.Timeout != 0 {
		client.Timeout = requestClone.Config.Timeout
	}

	req, err := requestClone.CreateHTTPRequest()
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		err := fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
		headers := make(map[string]string)
		for key, values := range resp.Header {
			if len(values) > 0 {
				headers[key] = values[0]
			}
		}
		errorResponse := &httptransport.Response[E]{
			StatusCode: resp.StatusCode,
			Headers:    headers,
			Body:       body,
			Raw:        resp,
		}
		return nil, httptransport.NewErrorResponse[E](err, errorResponse)
	}

	// Get context from request, or use background if not available
	ctx := request.Context
	if ctx == nil {
		ctx = context.Background()
	}

	stream := httptransport.NewStream[T](resp, ctx)
	return stream, nil
}

// SetNext logs a warning as the terminating handler should always be the last in the chain.
// Calling this method indicates a misconfiguration in the handler chain setup.
func (h *TerminatingHandler[T, E]) SetNext(handler Handler[T, E]) {
	fmt.Println("WARNING: SetNext should not be called on the terminating handler.")
}
