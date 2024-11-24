package torrents

type RequestDownloadLinkOkResponse struct {
	Data    *string `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

func (r *RequestDownloadLinkOkResponse) SetData(data string) {
	r.Data = &data
}

func (r *RequestDownloadLinkOkResponse) GetData() *string {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *RequestDownloadLinkOkResponse) SetDetail(detail string) {
	r.Detail = &detail
}

func (r *RequestDownloadLinkOkResponse) GetDetail() *string {
	if r == nil {
		return nil
	}
	return r.Detail
}

func (r *RequestDownloadLinkOkResponse) SetSuccess(success bool) {
	r.Success = &success
}

func (r *RequestDownloadLinkOkResponse) GetSuccess() *bool {
	if r == nil {
		return nil
	}
	return r.Success
}
