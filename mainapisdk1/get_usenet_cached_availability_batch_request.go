package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type GetUsenetCachedAvailabilityBatchRequest struct {
	Hashes *param.Nullable[[]string] `json:"hashes,omitempty" xml:"hashes,omitempty"`
}

func (g GetUsenetCachedAvailabilityBatchRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetUsenetCachedAvailabilityBatchRequest to string"
	}
	return string(jsonData)
}

func (g *GetUsenetCachedAvailabilityBatchRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, g)
}
