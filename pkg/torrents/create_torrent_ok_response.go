package torrents

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type CreateTorrentOkResponse struct {
	Data    *CreateTorrentOkResponseData `json:"data,omitempty"`
	Detail  *string                      `json:"detail,omitempty"`
	Error   *util.Nullable[any]          `json:"error,omitempty"`
	Success *bool                        `json:"success,omitempty"`
}

func (c *CreateTorrentOkResponse) GetData() *CreateTorrentOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateTorrentOkResponse) SetData(data CreateTorrentOkResponseData) {
	c.Data = &data
}

func (c *CreateTorrentOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateTorrentOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *CreateTorrentOkResponse) GetError() *util.Nullable[any] {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateTorrentOkResponse) SetError(error util.Nullable[any]) {
	c.Error = &error
}

func (c *CreateTorrentOkResponse) SetErrorNull() {
	c.Error = &util.Nullable[any]{IsNull: true}
}

func (c *CreateTorrentOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateTorrentOkResponse) SetSuccess(success bool) {
	c.Success = &success
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
}

func (c *CreateTorrentOkResponseData) GetActiveLimit() *float64 {
	if c == nil {
		return nil
	}
	return c.ActiveLimit
}

func (c *CreateTorrentOkResponseData) SetActiveLimit(activeLimit float64) {
	c.ActiveLimit = &activeLimit
}

func (c *CreateTorrentOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateTorrentOkResponseData) SetAuthId(authId string) {
	c.AuthId = &authId
}

func (c *CreateTorrentOkResponseData) GetCurrentActiveDownloads() *float64 {
	if c == nil {
		return nil
	}
	return c.CurrentActiveDownloads
}

func (c *CreateTorrentOkResponseData) SetCurrentActiveDownloads(currentActiveDownloads float64) {
	c.CurrentActiveDownloads = &currentActiveDownloads
}

func (c *CreateTorrentOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateTorrentOkResponseData) SetHash(hash string) {
	c.Hash = &hash
}

func (c *CreateTorrentOkResponseData) GetQueuedId() *float64 {
	if c == nil {
		return nil
	}
	return c.QueuedId
}

func (c *CreateTorrentOkResponseData) SetQueuedId(queuedId float64) {
	c.QueuedId = &queuedId
}

func (c *CreateTorrentOkResponseData) GetTorrentId() *float64 {
	if c == nil {
		return nil
	}
	return c.TorrentId
}

func (c *CreateTorrentOkResponseData) SetTorrentId(torrentId float64) {
	c.TorrentId = &torrentId
}

func (c CreateTorrentOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTorrentOkResponseData to string"
	}
	return string(jsonData)
}
