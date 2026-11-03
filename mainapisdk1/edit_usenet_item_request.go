package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type EditUsenetItemRequest struct {
	UsenetDownloadID  *param.Nullable[float64]  `json:"usenet_download_id,omitempty" xml:"usenet_download_id,omitempty"`
	Name              *param.Nullable[string]   `json:"name,omitempty" xml:"name,omitempty"`
	Tags              *param.Nullable[[]string] `json:"tags,omitempty" xml:"tags,omitempty"`
	AlternativeHashes *param.Nullable[[]string] `json:"alternative_hashes,omitempty" xml:"alternative_hashes,omitempty"`
	Airlocked         *param.Nullable[bool]     `json:"airlocked,omitempty" xml:"airlocked,omitempty"`
}

func (e EditUsenetItemRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditUsenetItemRequest to string"
	}
	return string(jsonData)
}

func (e *EditUsenetItemRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}
