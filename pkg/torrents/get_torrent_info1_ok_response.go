package torrents

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetTorrentInfo1OkResponse struct {
	Data    *GetTorrentInfo1OkResponseData `json:"data,omitempty"`
	Detail  *string                        `json:"detail,omitempty"`
	Error   *util.Nullable[any]            `json:"error,omitempty"`
	Success *bool                          `json:"success,omitempty"`
}

func (g *GetTorrentInfo1OkResponse) GetData() *GetTorrentInfo1OkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentInfo1OkResponse) SetData(data GetTorrentInfo1OkResponseData) {
	g.Data = &data
}

func (g *GetTorrentInfo1OkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentInfo1OkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentInfo1OkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentInfo1OkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetTorrentInfo1OkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetTorrentInfo1OkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetTorrentInfo1OkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetTorrentInfo1OkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfo1OkResponse to string"
	}
	return string(jsonData)
}

type GetTorrentInfo1OkResponseData struct {
	Files    []DataFiles3 `json:"files,omitempty"`
	Hash     *string      `json:"hash,omitempty"`
	Name     *string      `json:"name,omitempty"`
	Peers    *float64     `json:"peers,omitempty"`
	Seeds    *float64     `json:"seeds,omitempty"`
	Size     *float64     `json:"size,omitempty"`
	Trackers []any        `json:"trackers,omitempty"`
}

func (g *GetTorrentInfo1OkResponseData) GetFiles() []DataFiles3 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetTorrentInfo1OkResponseData) SetFiles(files []DataFiles3) {
	g.Files = files
}

func (g *GetTorrentInfo1OkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentInfo1OkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetTorrentInfo1OkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentInfo1OkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetTorrentInfo1OkResponseData) GetPeers() *float64 {
	if g == nil {
		return nil
	}
	return g.Peers
}

func (g *GetTorrentInfo1OkResponseData) SetPeers(peers float64) {
	g.Peers = &peers
}

func (g *GetTorrentInfo1OkResponseData) GetSeeds() *float64 {
	if g == nil {
		return nil
	}
	return g.Seeds
}

func (g *GetTorrentInfo1OkResponseData) SetSeeds(seeds float64) {
	g.Seeds = &seeds
}

func (g *GetTorrentInfo1OkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentInfo1OkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetTorrentInfo1OkResponseData) GetTrackers() []any {
	if g == nil {
		return nil
	}
	return g.Trackers
}

func (g *GetTorrentInfo1OkResponseData) SetTrackers(trackers []any) {
	g.Trackers = trackers
}

func (g GetTorrentInfo1OkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfo1OkResponseData to string"
	}
	return string(jsonData)
}

type DataFiles3 struct {
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
}

func (d *DataFiles3) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles3) SetName(name string) {
	d.Name = &name
}

func (d *DataFiles3) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

func (d *DataFiles3) SetSize(size float64) {
	d.Size = &size
}

func (d DataFiles3) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataFiles3 to string"
	}
	return string(jsonData)
}
