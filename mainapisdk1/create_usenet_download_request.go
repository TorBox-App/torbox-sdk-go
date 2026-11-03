package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type CreateUsenetDownloadRequest struct {
	File            *param.Nullable[[]byte] `json:"file,omitempty" xml:"file,omitempty"`
	Link            *param.Nullable[string] `json:"link,omitempty" xml:"link,omitempty"`
	Name            *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	Password        *param.Nullable[string] `json:"password,omitempty" xml:"password,omitempty"`
	PostProcessing  *param.Nullable[string] `json:"post_processing,omitempty" xml:"post_processing,omitempty"`
	AsQueued        *param.Nullable[string] `json:"as_queued,omitempty" xml:"as_queued,omitempty"`
	AddOnlyIfCached *param.Nullable[string] `json:"add_only_if_cached,omitempty" xml:"add_only_if_cached,omitempty"`
}

func (c CreateUsenetDownloadRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUsenetDownloadRequest to string"
	}
	return string(jsonData)
}

func (c *CreateUsenetDownloadRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
