package torrents

import (
	"encoding/json"
	"torbox-sdk-go/internal/utils"
	"torbox-sdk-go/pkg/util"
)

type GetTorrentCachedAvailabilityOkResponse struct {
	Data    map[string]*GetTorrentCachedAvailabilityOkResponseData `json:"data,omitempty"`
	Detail  *string                                                `json:"detail,omitempty"`
	Error   *util.Nullable[string]                                 `json:"error,omitempty"`
	Success *bool                                                  `json:"success,omitempty"`
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetData() map[string]*GetTorrentCachedAvailabilityOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetData(data map[string]*GetTorrentCachedAvailabilityOkResponseData) {
	g.Data = utils.CloneMap(data)
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetError() *util.Nullable[string] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetError(error util.Nullable[string]) {
	g.Error = &error
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[string]{IsNull: true}
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetTorrentCachedAvailabilityOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentCachedAvailabilityOkResponse to string"
	}
	return string(jsonData)
}

type GetTorrentCachedAvailabilityOkResponseData struct {
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Hash *string  `json:"hash,omitempty"`
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g GetTorrentCachedAvailabilityOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentCachedAvailabilityOkResponseData to string"
	}
	return string(jsonData)
}
