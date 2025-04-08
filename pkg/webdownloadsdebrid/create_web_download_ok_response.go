package webdownloadsdebrid

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type CreateWebDownloadOkResponse struct {
	Data    *CreateWebDownloadOkResponseData `json:"data,omitempty"`
	Detail  *string                          `json:"detail,omitempty"`
	Error   *util.Nullable[any]              `json:"error,omitempty"`
	Success *bool                            `json:"success,omitempty"`
}

func (c *CreateWebDownloadOkResponse) GetData() *CreateWebDownloadOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateWebDownloadOkResponse) SetData(data CreateWebDownloadOkResponseData) {
	c.Data = &data
}

func (c *CreateWebDownloadOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateWebDownloadOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *CreateWebDownloadOkResponse) GetError() *util.Nullable[any] {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateWebDownloadOkResponse) SetError(error util.Nullable[any]) {
	c.Error = &error
}

func (c *CreateWebDownloadOkResponse) SetErrorNull() {
	c.Error = &util.Nullable[any]{IsNull: true}
}

func (c *CreateWebDownloadOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateWebDownloadOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c CreateWebDownloadOkResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebDownloadOkResponse to string"
	}
	return string(jsonData)
}

type CreateWebDownloadOkResponseData struct {
	AuthId        *string `json:"auth_id,omitempty"`
	Hash          *string `json:"hash,omitempty"`
	WebdownloadId *string `json:"webdownload_id,omitempty"`
}

func (c *CreateWebDownloadOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateWebDownloadOkResponseData) SetAuthId(authId string) {
	c.AuthId = &authId
}

func (c *CreateWebDownloadOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateWebDownloadOkResponseData) SetHash(hash string) {
	c.Hash = &hash
}

func (c *CreateWebDownloadOkResponseData) GetWebdownloadId() *string {
	if c == nil {
		return nil
	}
	return c.WebdownloadId
}

func (c *CreateWebDownloadOkResponseData) SetWebdownloadId(webdownloadId string) {
	c.WebdownloadId = &webdownloadId
}

func (c CreateWebDownloadOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebDownloadOkResponseData to string"
	}
	return string(jsonData)
}
