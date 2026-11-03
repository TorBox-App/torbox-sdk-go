package mainapisdk1

import (
	"encoding/json"
	"example.com/main-api-sdk/internal/unmarshal"
	"example.com/main-api-sdk/param"
)

type GetTorrentInfoRequest struct {
	Magnet         *param.Nullable[string] `json:"magnet,omitempty" xml:"magnet,omitempty"`
	File           *param.Nullable[[]byte] `json:"file,omitempty" xml:"file,omitempty"`
	Hash           *param.Nullable[string] `json:"hash,omitempty" xml:"hash,omitempty"`
	Timeout        *param.Nullable[string] `json:"timeout,omitempty" xml:"timeout,omitempty"`
	UseCacheLookup *param.Nullable[string] `json:"use_cache_lookup,omitempty" xml:"use_cache_lookup,omitempty"`
	PeersOnly      *param.Nullable[string] `json:"peers_only,omitempty" xml:"peers_only,omitempty"`
}

func (g GetTorrentInfoRequest) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfoRequest to string"
	}
	return string(jsonData)
}

func (g *GetTorrentInfoRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, g)
}
