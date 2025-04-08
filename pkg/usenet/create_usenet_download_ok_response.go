package usenet

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type CreateUsenetDownloadOkResponse struct {
	Data    *CreateUsenetDownloadOkResponseData `json:"data,omitempty"`
	Detail  *string                             `json:"detail,omitempty"`
	Error   *util.Nullable[any]                 `json:"error,omitempty"`
	Success *bool                               `json:"success,omitempty"`
}

func (c *CreateUsenetDownloadOkResponse) GetData() *CreateUsenetDownloadOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateUsenetDownloadOkResponse) SetData(data CreateUsenetDownloadOkResponseData) {
	c.Data = &data
}

func (c *CreateUsenetDownloadOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateUsenetDownloadOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *CreateUsenetDownloadOkResponse) GetError() *util.Nullable[any] {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateUsenetDownloadOkResponse) SetError(error util.Nullable[any]) {
	c.Error = &error
}

func (c *CreateUsenetDownloadOkResponse) SetErrorNull() {
	c.Error = &util.Nullable[any]{IsNull: true}
}

func (c *CreateUsenetDownloadOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateUsenetDownloadOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c CreateUsenetDownloadOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUsenetDownloadOkResponse to string"
	}
	return string(jsonData)
}

type CreateUsenetDownloadOkResponseData struct {
	AuthId           *string `json:"auth_id,omitempty"`
	Hash             *string `json:"hash,omitempty"`
	UsenetdownloadId *string `json:"usenetdownload_id,omitempty"`
}

func (c *CreateUsenetDownloadOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateUsenetDownloadOkResponseData) SetAuthId(authId string) {
	c.AuthId = &authId
}

func (c *CreateUsenetDownloadOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateUsenetDownloadOkResponseData) SetHash(hash string) {
	c.Hash = &hash
}

func (c *CreateUsenetDownloadOkResponseData) GetUsenetdownloadId() *string {
	if c == nil {
		return nil
	}
	return c.UsenetdownloadId
}

func (c *CreateUsenetDownloadOkResponseData) SetUsenetdownloadId(usenetdownloadId string) {
	c.UsenetdownloadId = &usenetdownloadId
}

func (c CreateUsenetDownloadOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUsenetDownloadOkResponseData to string"
	}
	return string(jsonData)
}
