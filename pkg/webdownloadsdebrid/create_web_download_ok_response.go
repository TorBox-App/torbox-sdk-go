package webdownloadsdebrid

import (
	"encoding/json"
)

type CreateWebDownloadOkResponse struct {
	Data    *CreateWebDownloadOkResponseData `json:"data,omitempty"`
	Detail  *string                          `json:"detail,omitempty"`
	Error   any                              `json:"error,omitempty"`
	Success *bool                            `json:"success,omitempty"`
	touched map[string]bool
}

func (c *CreateWebDownloadOkResponse) GetData() *CreateWebDownloadOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateWebDownloadOkResponse) SetData(data CreateWebDownloadOkResponseData) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = &data
}

func (c *CreateWebDownloadOkResponse) SetDataNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = nil
}

func (c *CreateWebDownloadOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateWebDownloadOkResponse) SetDetail(detail string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = &detail
}

func (c *CreateWebDownloadOkResponse) SetDetailNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = nil
}

func (c *CreateWebDownloadOkResponse) GetError() any {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateWebDownloadOkResponse) SetError(error any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = error
}

func (c *CreateWebDownloadOkResponse) SetErrorNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = nil
}

func (c *CreateWebDownloadOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateWebDownloadOkResponse) SetSuccess(success bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = &success
}

func (c *CreateWebDownloadOkResponse) SetSuccessNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = nil
}

func (c CreateWebDownloadOkResponse) MarshalJSON() ([]byte, error) {
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
	touched       map[string]bool
}

func (c *CreateWebDownloadOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateWebDownloadOkResponseData) SetAuthId(authId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AuthId"] = true
	c.AuthId = &authId
}

func (c *CreateWebDownloadOkResponseData) SetAuthIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AuthId"] = true
	c.AuthId = nil
}

func (c *CreateWebDownloadOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateWebDownloadOkResponseData) SetHash(hash string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = &hash
}

func (c *CreateWebDownloadOkResponseData) SetHashNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = nil
}

func (c *CreateWebDownloadOkResponseData) GetWebdownloadId() *string {
	if c == nil {
		return nil
	}
	return c.WebdownloadId
}

func (c *CreateWebDownloadOkResponseData) SetWebdownloadId(webdownloadId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["WebdownloadId"] = true
	c.WebdownloadId = &webdownloadId
}

func (c *CreateWebDownloadOkResponseData) SetWebdownloadIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["WebdownloadId"] = true
	c.WebdownloadId = nil
}

func (c CreateWebDownloadOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["AuthId"] && c.AuthId == nil {
		data["auth_id"] = nil
	} else if c.AuthId != nil {
		data["auth_id"] = c.AuthId
	}

	if c.touched["Hash"] && c.Hash == nil {
		data["hash"] = nil
	} else if c.Hash != nil {
		data["hash"] = c.Hash
	}

	if c.touched["WebdownloadId"] && c.WebdownloadId == nil {
		data["webdownload_id"] = nil
	} else if c.WebdownloadId != nil {
		data["webdownload_id"] = c.WebdownloadId
	}

	return json.Marshal(data)
}

func (c CreateWebDownloadOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateWebDownloadOkResponseData to string"
	}
	return string(jsonData)
}
