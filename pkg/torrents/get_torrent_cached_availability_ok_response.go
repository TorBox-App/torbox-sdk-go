package torrents

import (
	"torbox-sdk-go/internal/utils"
)

type GetTorrentCachedAvailabilityOkResponse struct {
	Data    map[string]*GetTorrentCachedAvailabilityOkResponseData `json:"data,omitempty"`
	Detail  *string                                                `json:"detail,omitempty"`
	Error   *string                                                `json:"error,omitempty"`
	Success *bool                                                  `json:"success,omitempty"`
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetData(data map[string]*GetTorrentCachedAvailabilityOkResponseData) {
	g.Data = utils.CloneMap(data)
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetData() map[string]*GetTorrentCachedAvailabilityOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetError(error string) {
	g.Error = &error
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetError() *string {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

type GetTorrentCachedAvailabilityOkResponseData struct {
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
	Hash *string  `json:"hash,omitempty"`
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}
