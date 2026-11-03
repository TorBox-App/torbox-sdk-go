package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type GetWebDownloadCachedAvailabilityBatchRequest struct {
	Hashes *param.Nullable[[]string] `json:"hashes,omitempty" xml:"hashes,omitempty"`
}

func (g GetWebDownloadCachedAvailabilityBatchRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetWebDownloadCachedAvailabilityBatchRequest to string"
	}
	return string(jsonData)
}

func (g *GetWebDownloadCachedAvailabilityBatchRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, g)
}
