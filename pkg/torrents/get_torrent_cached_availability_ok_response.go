package torrents

import (
	"encoding/json"
	"torbox-sdk-go/internal/utils"
)

type GetTorrentCachedAvailabilityOkResponse struct {
	Data    map[string]*GetTorrentCachedAvailabilityOkResponseData `json:"data,omitempty"`
	Detail  *string                                                `json:"detail,omitempty"`
	Error   *string                                                `json:"error,omitempty"`
	Success *bool                                                  `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetData() map[string]*GetTorrentCachedAvailabilityOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetData(data map[string]*GetTorrentCachedAvailabilityOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = utils.CloneMap(data)
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetError() *string {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetError(error string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = &error
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetTorrentCachedAvailabilityOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetTorrentCachedAvailabilityOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetTorrentCachedAvailabilityOkResponse) MarshalJSON() ([]byte, error) {
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

func (g GetTorrentCachedAvailabilityOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentCachedAvailabilityOkResponse to string"
	}
	return string(jsonData)
}

type GetTorrentCachedAvailabilityOkResponseData struct {
	Name    *string  `json:"name,omitempty"`
	Size    *float64 `json:"size,omitempty"`
	Hash    *string  `json:"hash,omitempty"`
	touched map[string]bool
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetSize(size float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = &size
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetSizeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = nil
}

func (g *GetTorrentCachedAvailabilityOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetHash(hash string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = &hash
}

func (g *GetTorrentCachedAvailabilityOkResponseData) SetHashNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = nil
}

func (g GetTorrentCachedAvailabilityOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

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

	if g.touched["Hash"] && g.Hash == nil {
		data["hash"] = nil
	} else if g.Hash != nil {
		data["hash"] = g.Hash
	}

	return json.Marshal(data)
}

func (g GetTorrentCachedAvailabilityOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentCachedAvailabilityOkResponseData to string"
	}
	return string(jsonData)
}
