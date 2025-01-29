package torrents

import (
	"encoding/json"
)

type GetTorrentInfoOkResponse struct {
	Data    *GetTorrentInfoOkResponseData `json:"data,omitempty"`
	Detail  *string                       `json:"detail,omitempty"`
	Error   any                           `json:"error,omitempty"`
	Success *bool                         `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetTorrentInfoOkResponse) GetData() *GetTorrentInfoOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentInfoOkResponse) SetData(data GetTorrentInfoOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = &data
}

func (g *GetTorrentInfoOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetTorrentInfoOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentInfoOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetTorrentInfoOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetTorrentInfoOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentInfoOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetTorrentInfoOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetTorrentInfoOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetTorrentInfoOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetTorrentInfoOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetTorrentInfoOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Data"] && g.Data == nil {
		data["data"] = nil
	} else if g.Data != nil {
		data["data"] = g.Data
	}

	if g.touched["Detail"] && g.Detail == nil {
		data["detail"] = nil
	} else if g.Detail != nil {
		data["detail"] = g.Detail
	}

	if g.touched["Error"] && g.Error == nil {
		data["error"] = nil
	} else if g.Error != nil {
		data["error"] = g.Error
	}

	if g.touched["Success"] && g.Success == nil {
		data["success"] = nil
	} else if g.Success != nil {
		data["success"] = g.Success
	}

	return json.Marshal(data)
}

func (g GetTorrentInfoOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfoOkResponse to string"
	}
	return string(jsonData)
}

type GetTorrentInfoOkResponseData struct {
	Files   []DataFiles2 `json:"files,omitempty"`
	Hash    *string      `json:"hash,omitempty"`
	Name    *string      `json:"name,omitempty"`
	Size    *float64     `json:"size,omitempty"`
	touched map[string]bool
}

func (g *GetTorrentInfoOkResponseData) GetFiles() []DataFiles2 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetTorrentInfoOkResponseData) SetFiles(files []DataFiles2) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Files"] = true
	g.Files = files
}

func (g *GetTorrentInfoOkResponseData) SetFilesNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Files"] = true
	g.Files = nil
}

func (g *GetTorrentInfoOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentInfoOkResponseData) SetHash(hash string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = &hash
}

func (g *GetTorrentInfoOkResponseData) SetHashNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = nil
}

func (g *GetTorrentInfoOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentInfoOkResponseData) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GetTorrentInfoOkResponseData) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}

func (g *GetTorrentInfoOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentInfoOkResponseData) SetSize(size float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = &size
}

func (g *GetTorrentInfoOkResponseData) SetSizeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = nil
}

func (g GetTorrentInfoOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Files"] && g.Files == nil {
		data["files"] = nil
	} else if g.Files != nil {
		data["files"] = g.Files
	}

	if g.touched["Hash"] && g.Hash == nil {
		data["hash"] = nil
	} else if g.Hash != nil {
		data["hash"] = g.Hash
	}

	if g.touched["Name"] && g.Name == nil {
		data["name"] = nil
	} else if g.Name != nil {
		data["name"] = g.Name
	}

	if g.touched["Size"] && g.Size == nil {
		data["size"] = nil
	} else if g.Size != nil {
		data["size"] = g.Size
	}

	return json.Marshal(data)
}

func (g GetTorrentInfoOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentInfoOkResponseData to string"
	}
	return string(jsonData)
}

type DataFiles2 struct {
	Name    *string  `json:"name,omitempty"`
	Size    *float64 `json:"size,omitempty"`
	touched map[string]bool
}

func (d *DataFiles2) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles2) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DataFiles2) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DataFiles2) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

func (d *DataFiles2) SetSize(size float64) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Size"] = true
	d.Size = &size
}

func (d *DataFiles2) SetSizeNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Size"] = true
	d.Size = nil
}

func (d DataFiles2) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Name"] && d.Name == nil {
		data["name"] = nil
	} else if d.Name != nil {
		data["name"] = d.Name
	}

	if d.touched["Size"] && d.Size == nil {
		data["size"] = nil
	} else if d.Size != nil {
		data["size"] = d.Size
	}

	return json.Marshal(data)
}

func (d DataFiles2) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataFiles2 to string"
	}
	return string(jsonData)
}
