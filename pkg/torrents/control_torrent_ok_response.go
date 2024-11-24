package torrents

type ControlTorrentOkResponse struct {
	Data    any     `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

func (c *ControlTorrentOkResponse) SetData(data any) {
	c.Data = data
}

func (c *ControlTorrentOkResponse) GetData() any {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *ControlTorrentOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *ControlTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *ControlTorrentOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c *ControlTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}
