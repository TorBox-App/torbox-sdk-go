package usenet

type CreateUsenetDownloadOkResponse struct {
	Data    *CreateUsenetDownloadOkResponseData `json:"data,omitempty"`
	Detail  *string                             `json:"detail,omitempty"`
	Success *bool                               `json:"success,omitempty"`
}

func (c *CreateUsenetDownloadOkResponse) SetData(data CreateUsenetDownloadOkResponseData) {
	c.Data = &data
}

func (c *CreateUsenetDownloadOkResponse) GetData() *CreateUsenetDownloadOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateUsenetDownloadOkResponse) SetDetail(detail string) {
	c.Detail = &detail
}

func (c *CreateUsenetDownloadOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateUsenetDownloadOkResponse) SetSuccess(success bool) {
	c.Success = &success
}

func (c *CreateUsenetDownloadOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

type CreateUsenetDownloadOkResponseData struct {
	AuthId           *string `json:"auth_id,omitempty"`
	Hash             *string `json:"hash,omitempty"`
	UsenetdownloadId *string `json:"usenetdownload_id,omitempty"`
}

func (c *CreateUsenetDownloadOkResponseData) SetAuthId(authId string) {
	c.AuthId = &authId
}

func (c *CreateUsenetDownloadOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateUsenetDownloadOkResponseData) SetHash(hash string) {
	c.Hash = &hash
}

func (c *CreateUsenetDownloadOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateUsenetDownloadOkResponseData) SetUsenetdownloadId(usenetdownloadId string) {
	c.UsenetdownloadId = &usenetdownloadId
}

func (c *CreateUsenetDownloadOkResponseData) GetUsenetdownloadId() *string {
	if c == nil {
		return nil
	}
	return c.UsenetdownloadId
}
