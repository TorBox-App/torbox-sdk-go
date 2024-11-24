package shared

import (
	"torbox-sdk-go/internal/clients/rest/httptransport"
)

type TorboxApiResponse[T any] struct {
	Data     T
	Metadata TorboxApiResponseMetadata
}

type TorboxApiResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewTorboxApiResponse[T any](resp *httptransport.Response[T]) *TorboxApiResponse[T] {
	return &TorboxApiResponse[T]{
		Data: resp.Data,
		Metadata: TorboxApiResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
