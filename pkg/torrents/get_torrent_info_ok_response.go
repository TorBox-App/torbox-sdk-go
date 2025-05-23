package torrents

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetTorrentInfoOkResponse struct {
	Data    *GetTorrentInfoOkResponseData `json:"data,omitempty"`
	Detail  *string                       `json:"detail,omitempty"`
	Error   *util.Nullable[any]           `json:"error,omitempty"`
	Success *bool                         `json:"success,omitempty"`
}

func (g *GetTorrentInfoOkResponse) GetData() *GetTorrentInfoOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentInfoOkResponse) SetData(data GetTorrentInfoOkResponseData) {
	g.Data = &data
}

func (g *GetTorrentInfoOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentInfoOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentInfoOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentInfoOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetTorrentInfoOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetTorrentInfoOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetTorrentInfoOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetTorrentInfoOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfoOkResponse to string"
	}
	return string(jsonData)
}

type GetTorrentInfoOkResponseData struct {
	Files    []DataFiles2 `json:"files,omitempty"`
	Hash     *string      `json:"hash,omitempty"`
	Name     *string      `json:"name,omitempty"`
	Peers    *float64     `json:"peers,omitempty"`
	Seeds    *float64     `json:"seeds,omitempty"`
	Size     *float64     `json:"size,omitempty"`
	Trackers []any        `json:"trackers,omitempty"`
}

func (g *GetTorrentInfoOkResponseData) GetFiles() []DataFiles2 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetTorrentInfoOkResponseData) SetFiles(files []DataFiles2) {
	g.Files = files
}

func (g *GetTorrentInfoOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentInfoOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetTorrentInfoOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentInfoOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetTorrentInfoOkResponseData) GetPeers() *float64 {
	if g == nil {
		return nil
	}
	return g.Peers
}

func (g *GetTorrentInfoOkResponseData) SetPeers(peers float64) {
	g.Peers = &peers
}

func (g *GetTorrentInfoOkResponseData) GetSeeds() *float64 {
	if g == nil {
		return nil
	}
	return g.Seeds
}

func (g *GetTorrentInfoOkResponseData) SetSeeds(seeds float64) {
	g.Seeds = &seeds
}

func (g *GetTorrentInfoOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentInfoOkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetTorrentInfoOkResponseData) GetTrackers() []any {
	if g == nil {
		return nil
	}
	return g.Trackers
}

func (g *GetTorrentInfoOkResponseData) SetTrackers(trackers []any) {
	g.Trackers = trackers
}

func (g GetTorrentInfoOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfoOkResponseData to string"
	}
	return string(jsonData)
}

type DataFiles2 struct {
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
}

func (d *DataFiles2) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles2) SetName(name string) {
	d.Name = &name
}

func (d *DataFiles2) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

func (d *DataFiles2) SetSize(size float64) {
	d.Size = &size
}

func (d DataFiles2) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataFiles2 to string"
	}
	return string(jsonData)
}
