package webdownloadsdebrid

type CreateWebDownloadOkResponse struct {
	Data    *CreateWebDownloadOkResponseData `json:"data,omitempty"`
	Detail  *string                          `json:"detail,omitempty"`
	Error   any                              `json:"error,omitempty"`
	Success *bool                            `json:"success,omitempty"`
}

func (c *CreateWebDownloadOkResponse) SetData(data CreateWebDownloadOkResponseData) {
	c.Data = &data
}

func (c *CreateWebDownloadOkResponse) GetData() *CreateWebDownloadOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateWebDownloadOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *CreateWebDownloadOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateWebDownloadOkResponse) SetError(error any) {
	c.Error = error
}

func (c *CreateWebDownloadOkResponse) GetError() any {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateWebDownloadOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c *CreateWebDownloadOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

type CreateWebDownloadOkResponseData struct {
	AuthId        *string `json:"auth_id,omitempty"`
	Hash          *string `json:"hash,omitempty"`
	WebdownloadId *string `json:"webdownload_id,omitempty"`
}

func (c *CreateWebDownloadOkResponseData) SetAuthId(authId string) {
	c.AuthId = &authId
}

func (c *CreateWebDownloadOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateWebDownloadOkResponseData) SetHash(hash string) {
	c.Hash = &hash
}

func (c *CreateWebDownloadOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateWebDownloadOkResponseData) SetWebdownloadId(webdownloadId string) {
	c.WebdownloadId = &webdownloadId
}

func (c *CreateWebDownloadOkResponseData) GetWebdownloadId() *string {
	if c == nil {
		return nil
	}
	return c.WebdownloadId
}
