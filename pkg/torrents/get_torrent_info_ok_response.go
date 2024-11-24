package torrents

type GetTorrentInfoOkResponse struct {
	Data    *GetTorrentInfoOkResponseData `json:"data,omitempty"`
	Detail  *string                       `json:"detail,omitempty"`
	Error   any                           `json:"error,omitempty"`
	Success *bool                         `json:"success,omitempty"`
}

func (g *GetTorrentInfoOkResponse) SetData(data GetTorrentInfoOkResponseData) {
	g.Data = &data
}

func (g *GetTorrentInfoOkResponse) GetData() *GetTorrentInfoOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentInfoOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentInfoOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentInfoOkResponse) SetError(error any) {
	g.Error = error
}

func (g *GetTorrentInfoOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentInfoOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetTorrentInfoOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

type GetTorrentInfoOkResponseData struct {
	Files []DataFiles2 `json:"files,omitempty"`
	Hash  *string      `json:"hash,omitempty"`
	Name  *string      `json:"name,omitempty"`
	Size  *float64     `json:"size,omitempty"`
}

func (g *GetTorrentInfoOkResponseData) SetFiles(files []DataFiles2) {
	g.Files = files
}

func (g *GetTorrentInfoOkResponseData) GetFiles() []DataFiles2 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetTorrentInfoOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetTorrentInfoOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentInfoOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetTorrentInfoOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentInfoOkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetTorrentInfoOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

type DataFiles2 struct {
	Name *string  `json:"name,omitempty"`
	Size *float64 `json:"size,omitempty"`
}

func (d *DataFiles2) SetName(name string) {
	d.Name = &name
}

func (d *DataFiles2) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles2) SetSize(size float64) {
	d.Size = &size
}

func (d *DataFiles2) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}
