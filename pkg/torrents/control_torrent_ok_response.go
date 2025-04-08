package torrents

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type ControlTorrentOkResponse struct {
	Data    *util.Nullable[any] `json:"data,omitempty"`
	Detail  *string             `json:"detail,omitempty"`
	Error   *util.Nullable[any] `json:"error,omitempty"`
	Success *bool               `json:"success,omitempty"`
}

func (c *ControlTorrentOkResponse) GetData() *util.Nullable[any] {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *ControlTorrentOkResponse) SetData(data util.Nullable[any]) {
	c.Data = &data
}

func (c *ControlTorrentOkResponse) SetDataNull() {
	c.Data = &util.Nullable[any]{IsNull: true}
}

func (c *ControlTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *ControlTorrentOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *ControlTorrentOkResponse) GetError() *util.Nullable[any] {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *ControlTorrentOkResponse) SetError(error util.Nullable[any]) {
	c.Error = &error
}

func (c *ControlTorrentOkResponse) SetErrorNull() {
	c.Error = &util.Nullable[any]{IsNull: true}
}

func (c *ControlTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *ControlTorrentOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c ControlTorrentOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ControlTorrentOkResponse to string"
	}
	return string(jsonData)
}
