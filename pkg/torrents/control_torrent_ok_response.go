package torrents

import (
	"encoding/json"
)

type ControlTorrentOkResponse struct {
	Data    any     `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
	touched map[string]bool
}

func (c *ControlTorrentOkResponse) GetData() any {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *ControlTorrentOkResponse) SetData(data any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = data
}

func (c *ControlTorrentOkResponse) SetDataNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = nil
}

func (c *ControlTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *ControlTorrentOkResponse) SetDetail(detail string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = &detail
}

func (c *ControlTorrentOkResponse) SetDetailNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = nil
}

func (c *ControlTorrentOkResponse) GetError() any {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *ControlTorrentOkResponse) SetError(error any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = error
}

func (c *ControlTorrentOkResponse) SetErrorNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = nil
}

func (c *ControlTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *ControlTorrentOkResponse) SetSuccess(success bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = &success
}

func (c *ControlTorrentOkResponse) SetSuccessNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = nil
}

func (c ControlTorrentOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Data"] && c.Data == nil {
		data["data"] = nil
	} else if c.Data != nil {
		data["data"] = c.Data
	}

	if c.touched["Detail"] && c.Detail == nil {
		data["detail"] = nil
	} else if c.Detail != nil {
		data["detail"] = c.Detail
	}

	if c.touched["Error"] && c.Error == nil {
		data["error"] = nil
	} else if c.Error != nil {
		data["error"] = c.Error
	}

	if c.touched["Success"] && c.Success == nil {
		data["success"] = nil
	} else if c.Success != nil {
		data["success"] = c.Success
	}

	return json.Marshal(data)
}

func (c ControlTorrentOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ControlTorrentOkResponse to string"
	}
	return string(jsonData)
}
