package torrents

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type RequestDownloadLinkOkResponse struct {
	Data    *string             `json:"data,omitempty"`
	Detail  *string             `json:"detail,omitempty"`
	Error   *util.Nullable[any] `json:"error,omitempty"`
	Success *bool               `json:"success,omitempty"`
}

func (r *RequestDownloadLinkOkResponse) GetData() *string {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *RequestDownloadLinkOkResponse) SetData(data string) {
	r.Data = &data
}

func (r *RequestDownloadLinkOkResponse) GetDetail() *string {
	if r == nil {
		return nil
	}
	return r.Detail
}

func (r *RequestDownloadLinkOkResponse) SetDetail(detail string) {
	r.Detail = &detail
}

func (r *RequestDownloadLinkOkResponse) GetError() *util.Nullable[any] {
	if r == nil {
		return nil
	}
	return r.Error
}

func (r *RequestDownloadLinkOkResponse) SetError(error util.Nullable[any]) {
	r.Error = &error
}

func (r *RequestDownloadLinkOkResponse) SetErrorNull() {
	r.Error = &util.Nullable[any]{IsNull: true}
}

func (r *RequestDownloadLinkOkResponse) GetSuccess() *bool {
	if r == nil {
		return nil
	}
	return r.Success
}

func (r *RequestDownloadLinkOkResponse) SetSuccess(success bool) {
	r.Success = &success
}

func (r RequestDownloadLinkOkResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RequestDownloadLinkOkResponse to string"
	}
	return string(jsonData)
}
