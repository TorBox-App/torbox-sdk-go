package shared

import (
	"torbox-sdk-go/internal/clients/rest/httptransport"
)

type TorboxApiError struct {
	Err      error
	Body     []byte
	Metadata TorboxApiErrorMetadata
}

type TorboxApiErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewTorboxApiError[T any](transportError *httptransport.ErrorResponse[T]) *TorboxApiError {
	return &TorboxApiError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: TorboxApiErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *TorboxApiError) Error() string {
	return e.Err.Error()
}
