package usenet

import (
	"encoding/json"
)

type CreateUsenetDownloadOkResponse struct {
	Data    *CreateUsenetDownloadOkResponseData `json:"data,omitempty"`
	Detail  *string                             `json:"detail,omitempty"`
	Error   any                                 `json:"error,omitempty"`
	Success *bool                               `json:"success,omitempty"`
	touched map[string]bool
}

func (c *CreateUsenetDownloadOkResponse) GetData() *CreateUsenetDownloadOkResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateUsenetDownloadOkResponse) SetData(data CreateUsenetDownloadOkResponseData) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = &data
}

func (c *CreateUsenetDownloadOkResponse) SetDataNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Data"] = true
	c.Data = nil
}

func (c *CreateUsenetDownloadOkResponse) GetDetail() *string {
	if c == nil {
		return nil
	}
	return c.Detail
}

func (c *CreateUsenetDownloadOkResponse) SetDetail(detail string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = &detail
}

func (c *CreateUsenetDownloadOkResponse) SetDetailNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Detail"] = true
	c.Detail = nil
}

func (c *CreateUsenetDownloadOkResponse) GetError() any {
	if c == nil {
		return nil
	}
	return c.Error
}

func (c *CreateUsenetDownloadOkResponse) SetError(error any) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = error
}

func (c *CreateUsenetDownloadOkResponse) SetErrorNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Error"] = true
	c.Error = nil
}

func (c *CreateUsenetDownloadOkResponse) GetSuccess() *bool {
	if c == nil {
		return nil
	}
	return c.Success
}

func (c *CreateUsenetDownloadOkResponse) SetSuccess(success bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = &success
}

func (c *CreateUsenetDownloadOkResponse) SetSuccessNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Success"] = true
	c.Success = nil
}

func (c CreateUsenetDownloadOkResponse) MarshalJSON() ([]byte, error) {
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
	touched          map[string]bool
}

func (c *CreateUsenetDownloadOkResponseData) GetAuthId() *string {
	if c == nil {
		return nil
	}
	return c.AuthId
}

func (c *CreateUsenetDownloadOkResponseData) SetAuthId(authId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AuthId"] = true
	c.AuthId = &authId
}

func (c *CreateUsenetDownloadOkResponseData) SetAuthIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["AuthId"] = true
	c.AuthId = nil
}

func (c *CreateUsenetDownloadOkResponseData) GetHash() *string {
	if c == nil {
		return nil
	}
	return c.Hash
}

func (c *CreateUsenetDownloadOkResponseData) SetHash(hash string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = &hash
}

func (c *CreateUsenetDownloadOkResponseData) SetHashNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Hash"] = true
	c.Hash = nil
}

func (c *CreateUsenetDownloadOkResponseData) GetUsenetdownloadId() *string {
	if c == nil {
		return nil
	}
	return c.UsenetdownloadId
}

func (c *CreateUsenetDownloadOkResponseData) SetUsenetdownloadId(usenetdownloadId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["UsenetdownloadId"] = true
	c.UsenetdownloadId = &usenetdownloadId
}

func (c *CreateUsenetDownloadOkResponseData) SetUsenetdownloadIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["UsenetdownloadId"] = true
	c.UsenetdownloadId = nil
}

func (c CreateUsenetDownloadOkResponseData) MarshalJSON() ([]byte, error) {
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

	if c.touched["UsenetdownloadId"] && c.UsenetdownloadId == nil {
		data["usenetdownload_id"] = nil
	} else if c.UsenetdownloadId != nil {
		data["usenetdownload_id"] = c.UsenetdownloadId
	}

	return json.Marshal(data)
}

func (c CreateUsenetDownloadOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUsenetDownloadOkResponseData to string"
	}
	return string(jsonData)
}
