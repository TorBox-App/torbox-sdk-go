package handlers

import (
	"fmt"
	"net/http"

	"torbox-sdk-go/internal/clients/rest/httptransport"
)

type TerminatingHandler[T any] struct{}

func NewTerminatingHandler[T any]() *TerminatingHandler[T] {
	return &TerminatingHandler[T]{}
}

func (h *TerminatingHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	requestClone := request.Clone()

	client := http.Client{Timeout: *requestClone.Config.Timeout}

	req, err := requestClone.CreateHttpRequest()
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	transportResponse, responseErr := httptransport.NewResponse[T](resp)
	if responseErr != nil {
		return nil, httptransport.NewErrorResponse[T](responseErr, transportResponse)
	}

	if transportResponse.StatusCode >= 400 {
		err := fmt.Errorf("HTTP request failed with status code %d", transportResponse.StatusCode)
		return nil, httptransport.NewErrorResponse[T](err, transportResponse)
	}

	return transportResponse, nil
}

func (h *TerminatingHandler[T]) SetNext(handler Handler[T]) {
	fmt.Println("WARNING: SetNext should not be called on the terminating handler.")
}
