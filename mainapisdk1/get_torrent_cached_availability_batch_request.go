package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type GetTorrentCachedAvailabilityBatchRequest struct {
	Hashes *param.Nullable[[]string] `json:"hashes,omitempty" xml:"hashes,omitempty"`
}

func (g GetTorrentCachedAvailabilityBatchRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentCachedAvailabilityBatchRequest to string"
	}
	return string(jsonData)
}

func (g *GetTorrentCachedAvailabilityBatchRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, g)
}
