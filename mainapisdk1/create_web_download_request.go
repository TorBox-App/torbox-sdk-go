package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type CreateWebDownloadRequest struct {
	Link            *param.Nullable[string] `json:"link,omitempty" xml:"link,omitempty"`
	Password        *param.Nullable[string] `json:"password,omitempty" xml:"password,omitempty"`
	Name            *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	AsQueued        *param.Nullable[string] `json:"as_queued,omitempty" xml:"as_queued,omitempty"`
	AddOnlyIfCached *param.Nullable[string] `json:"add_only_if_cached,omitempty" xml:"add_only_if_cached,omitempty"`
}

func (c CreateWebDownloadRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebDownloadRequest to string"
	}
	return string(jsonData)
}

func (c *CreateWebDownloadRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
