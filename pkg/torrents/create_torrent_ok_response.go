package torrents

type CreateTorrentOkResponse struct {
	Data    *CreateTorrentOkResponseData `json:"data,omitempty"`
	Detail  *string                      `json:"detail,omitempty"`
	Error   any                          `json:"error,omitempty"`
	Success *bool                        `json:"success,omitempty"`
}

func (c *CreateTorrentOkResponse) SetData(data CreateTorrentOkResponseData) {
	c.Data = &data
}

func (c *CreateTorrentOkResponse) GetData() *CreateTorrentOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateTorrentOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *CreateTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateTorrentOkResponse) SetError(error any) {
	c.Error = error
}

func (c *CreateTorrentOkResponse) GetError() any {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateTorrentOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c *CreateTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

type CreateTorrentOkResponseData struct {
	ActiveLimit            *float64 `json:"active_limit,omitempty"`
	CurrentActiveDownloads *float64 `json:"current_active_downloads,omitempty"`
	Hash                   *string  `json:"hash,omitempty"`
	Name                   *string  `json:"name,omitempty"`
	QueuedId               *float64 `json:"queued_id,omitempty"`
	TorrentId              *float64 `json:"torrent_id,omitempty"`
}

func (c *CreateTorrentOkResponseData) SetActiveLimit(activeLimit float64) {
	c.ActiveLimit = &activeLimit
}

func (c *CreateTorrentOkResponseData) GetActiveLimit() *float64 {
	if c == nil {
		return nil
	}
	return c.ActiveLimit
}

func (c *CreateTorrentOkResponseData) SetCurrentActiveDownloads(currentActiveDownloads float64) {
	c.CurrentActiveDownloads = &currentActiveDownloads
}

func (c *CreateTorrentOkResponseData) GetCurrentActiveDownloads() *float64 {
	if c == nil {
		return nil
	}
	return c.CurrentActiveDownloads
}

func (c *CreateTorrentOkResponseData) SetHash(hash string) {
	c.Hash = &hash
}

func (c *CreateTorrentOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateTorrentOkResponseData) SetName(name string) {
	c.Name = &name
}

func (c *CreateTorrentOkResponseData) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTorrentOkResponseData) SetQueuedId(queuedId float64) {
	c.QueuedId = &queuedId
}

func (c *CreateTorrentOkResponseData) GetQueuedId() *float64 {
	if c == nil {
		return nil
	}
	return c.QueuedId
}

func (c *CreateTorrentOkResponseData) SetTorrentId(torrentId float64) {
	c.TorrentId = &torrentId
}

func (c *CreateTorrentOkResponseData) GetTorrentId() *float64 {
	if c == nil {
		return nil
	}
	return c.TorrentId
}
