package sharedmodels

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/clients/rest/httptransport"
	"net/http"
)

// MainAPISDKResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type MainAPISDKResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata MainAPISDKResponseMetadata
}

// MainAPISDKResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type MainAPISDKResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewMainAPISDKResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewMainAPISDKResponse[T any](resp *httptransport.Response[T]) *MainAPISDKResponse[T] {
	return &MainAPISDKResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: MainAPISDKResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *MainAPISDKResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r MainAPISDKResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: MainAPISDKResponse to string"
	}
	return string(jsonData)
}
