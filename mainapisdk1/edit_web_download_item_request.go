package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type EditWebDownloadItemRequest struct {
	WebdlID           *param.Nullable[float64]  `json:"webdl_id,omitempty" xml:"webdl_id,omitempty"`
	Name              *param.Nullable[string]   `json:"name,omitempty" xml:"name,omitempty"`
	Tags              *param.Nullable[[]string] `json:"tags,omitempty" xml:"tags,omitempty"`
	AlternativeHashes *param.Nullable[[]string] `json:"alternative_hashes,omitempty" xml:"alternative_hashes,omitempty"`
	Airlocked         *param.Nullable[bool]     `json:"airlocked,omitempty" xml:"airlocked,omitempty"`
}

func (e EditWebDownloadItemRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditWebDownloadItemRequest to string"
	}
	return string(jsonData)
}

func (e *EditWebDownloadItemRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}
