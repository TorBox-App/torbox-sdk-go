package torrents

import (
	"encoding/json"
)

type RequestDownloadLinkOkResponse struct {
	Data    *string `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
	touched map[string]bool
}

func (r *RequestDownloadLinkOkResponse) GetData() *string {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *RequestDownloadLinkOkResponse) SetData(data string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Data"] = true
	r.Data = &data
}

func (r *RequestDownloadLinkOkResponse) SetDataNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Data"] = true
	r.Data = nil
}

func (r *RequestDownloadLinkOkResponse) GetDetail() *string {
	if r == nil {
		return nil
	}
	return r.Detail
}

func (r *RequestDownloadLinkOkResponse) SetDetail(detail string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Detail"] = true
	r.Detail = &detail
}

func (r *RequestDownloadLinkOkResponse) SetDetailNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Detail"] = true
	r.Detail = nil
}

func (r *RequestDownloadLinkOkResponse) GetError() any {
	if r == nil {
		return nil
	}
	return r.Error
}

func (r *RequestDownloadLinkOkResponse) SetError(error any) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Error"] = true
	r.Error = error
}

func (r *RequestDownloadLinkOkResponse) SetErrorNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Error"] = true
	r.Error = nil
}

func (r *RequestDownloadLinkOkResponse) GetSuccess() *bool {
	if r == nil {
		return nil
	}
	return r.Success
}

func (r *RequestDownloadLinkOkResponse) SetSuccess(success bool) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Success"] = true
	r.Success = &success
}

func (r *RequestDownloadLinkOkResponse) SetSuccessNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Success"] = true
	r.Success = nil
}

func (r RequestDownloadLinkOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Data"] && r.Data == nil {
		data["data"] = nil
	} else if r.Data != nil {
		data["data"] = r.Data
	}

	if r.touched["Detail"] && r.Detail == nil {
		data["detail"] = nil
	} else if r.Detail != nil {
		data["detail"] = r.Detail
	}

	if r.touched["Error"] && r.Error == nil {
		data["error"] = nil
	} else if r.Error != nil {
		data["error"] = r.Error
	}

	if r.touched["Success"] && r.Success == nil {
		data["success"] = nil
	} else if r.Success != nil {
		data["success"] = r.Success
	}

	return json.Marshal(data)
}

func (r RequestDownloadLinkOkResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RequestDownloadLinkOkResponse to string"
	}
	return string(jsonData)
}
