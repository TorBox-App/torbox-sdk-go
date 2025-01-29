package torrents

import (
	"encoding/json"
)

type CreateTorrentOkResponse struct {
	Data    *CreateTorrentOkResponseData `json:"data,omitempty"`
	Detail  *string                      `json:"detail,omitempty"`
	Error   any                          `json:"error,omitempty"`
	Success *bool                        `json:"success,omitempty"`
	touched map[string]bool
}

func (c *CreateTorrentOkResponse) GetData() *CreateTorrentOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateTorrentOkResponse) SetData(data CreateTorrentOkResponseData) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = &data
}

func (c *CreateTorrentOkResponse) SetDataNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = nil
}

func (c *CreateTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateTorrentOkResponse) SetDetail(detail string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = &detail
}

func (c *CreateTorrentOkResponse) SetDetailNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = nil
}

func (c *CreateTorrentOkResponse) GetError() any {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateTorrentOkResponse) SetError(error any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = error
}

func (c *CreateTorrentOkResponse) SetErrorNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = nil
}

func (c *CreateTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateTorrentOkResponse) SetSuccess(success bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = &success
}

func (c *CreateTorrentOkResponse) SetSuccessNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = nil
}

func (c CreateTorrentOkResponse) MarshalJSON() ([]byte, error) {
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

func (c CreateTorrentOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTorrentOkResponse to string"
	}
	return string(jsonData)
}

type CreateTorrentOkResponseData struct {
	ActiveLimit            *float64 `json:"active_limit,omitempty"`
	AuthId                 *string  `json:"auth_id,omitempty"`
	CurrentActiveDownloads *float64 `json:"current_active_downloads,omitempty"`
	Hash                   *string  `json:"hash,omitempty"`
	QueuedId               *float64 `json:"queued_id,omitempty"`
	TorrentId              *float64 `json:"torrent_id,omitempty"`
	touched                map[string]bool
}

func (c *CreateTorrentOkResponseData) GetActiveLimit() *float64 {
	if c == nil {
		return nil
	}
	return c.ActiveLimit
}

func (c *CreateTorrentOkResponseData) SetActiveLimit(activeLimit float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ActiveLimit"] = true
	c.ActiveLimit = &activeLimit
}

func (c *CreateTorrentOkResponseData) SetActiveLimitNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ActiveLimit"] = true
	c.ActiveLimit = nil
}

func (c *CreateTorrentOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateTorrentOkResponseData) SetAuthId(authId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AuthId"] = true
	c.AuthId = &authId
}

func (c *CreateTorrentOkResponseData) SetAuthIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AuthId"] = true
	c.AuthId = nil
}

func (c *CreateTorrentOkResponseData) GetCurrentActiveDownloads() *float64 {
	if c == nil {
		return nil
	}
	return c.CurrentActiveDownloads
}

func (c *CreateTorrentOkResponseData) SetCurrentActiveDownloads(currentActiveDownloads float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CurrentActiveDownloads"] = true
	c.CurrentActiveDownloads = &currentActiveDownloads
}

func (c *CreateTorrentOkResponseData) SetCurrentActiveDownloadsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["CurrentActiveDownloads"] = true
	c.CurrentActiveDownloads = nil
}

func (c *CreateTorrentOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateTorrentOkResponseData) SetHash(hash string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = &hash
}

func (c *CreateTorrentOkResponseData) SetHashNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = nil
}

func (c *CreateTorrentOkResponseData) GetQueuedId() *float64 {
	if c == nil {
		return nil
	}
	return c.QueuedId
}

func (c *CreateTorrentOkResponseData) SetQueuedId(queuedId float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueuedId"] = true
	c.QueuedId = &queuedId
}

func (c *CreateTorrentOkResponseData) SetQueuedIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["QueuedId"] = true
	c.QueuedId = nil
}

func (c *CreateTorrentOkResponseData) GetTorrentId() *float64 {
	if c == nil {
		return nil
	}
	return c.TorrentId
}

func (c *CreateTorrentOkResponseData) SetTorrentId(torrentId float64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["TorrentId"] = true
	c.TorrentId = &torrentId
}

func (c *CreateTorrentOkResponseData) SetTorrentIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["TorrentId"] = true
	c.TorrentId = nil
}

func (c CreateTorrentOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["ActiveLimit"] && c.ActiveLimit == nil {
		data["active_limit"] = nil
	} else if c.ActiveLimit != nil {
		data["active_limit"] = c.ActiveLimit
	}

	if c.touched["AuthId"] && c.AuthId == nil {
		data["auth_id"] = nil
	} else if c.AuthId != nil {
		data["auth_id"] = c.AuthId
	}

	if c.touched["CurrentActiveDownloads"] && c.CurrentActiveDownloads == nil {
		data["current_active_downloads"] = nil
	} else if c.CurrentActiveDownloads != nil {
		data["current_active_downloads"] = c.CurrentActiveDownloads
	}

	if c.touched["Hash"] && c.Hash == nil {
		data["hash"] = nil
	} else if c.Hash != nil {
		data["hash"] = c.Hash
	}

	if c.touched["QueuedId"] && c.QueuedId == nil {
		data["queued_id"] = nil
	} else if c.QueuedId != nil {
		data["queued_id"] = c.QueuedId
	}

	if c.touched["TorrentId"] && c.TorrentId == nil {
		data["torrent_id"] = nil
	} else if c.TorrentId != nil {
		data["torrent_id"] = c.TorrentId
	}

	return json.Marshal(data)
}

func (c CreateTorrentOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTorrentOkResponseData to string"
	}
	return string(jsonData)
}
