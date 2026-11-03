package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type AsyncCreateTorrentRequest struct {
	File            *param.Nullable[[]byte] `json:"file,omitempty" xml:"file,omitempty"`
	Magnet          *param.Nullable[string] `json:"magnet,omitempty" xml:"magnet,omitempty"`
	Seed            *param.Nullable[string] `json:"seed,omitempty" xml:"seed,omitempty"`
	AllowZip        *param.Nullable[string] `json:"allow_zip,omitempty" xml:"allow_zip,omitempty"`
	Name            *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	AsQueued        *param.Nullable[string] `json:"as_queued,omitempty" xml:"as_queued,omitempty"`
	AddOnlyIfCached *param.Nullable[string] `json:"add_only_if_cached,omitempty" xml:"add_only_if_cached,omitempty"`
}

func (a AsyncCreateTorrentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AsyncCreateTorrentRequest to string"
	}
	return string(jsonData)
}

func (a *AsyncCreateTorrentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
