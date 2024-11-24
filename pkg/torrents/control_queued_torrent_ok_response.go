package torrents

type ControlQueuedTorrentOkResponse struct {
	Data    any     `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

func (c *ControlQueuedTorrentOkResponse) SetData(data any) {
	c.Data = data
}

func (c *ControlQueuedTorrentOkResponse) GetData() any {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *ControlQueuedTorrentOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *ControlQueuedTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *ControlQueuedTorrentOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c *ControlQueuedTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}