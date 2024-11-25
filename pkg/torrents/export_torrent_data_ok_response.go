package torrents

type ExportTorrentDataOkResponse struct {
	Data    *string `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

func (e *ExportTorrentDataOkResponse) SetData(data string) {
	e.Data = &data
}

func (e *ExportTorrentDataOkResponse) GetData() *string {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ExportTorrentDataOkResponse) SetDetail(detail string) {
	e.Detail = &detail
}

func (e *ExportTorrentDataOkResponse) GetDetail() *string {
	if e == nil {
		return nil
	}
	return e.Detail
}

func (e *ExportTorrentDataOkResponse) SetError(error any) {
	e.Error = error
}

func (e *ExportTorrentDataOkResponse) GetError() any {
	if e == nil {
		return nil
	}
	return e.Error
}

func (e *ExportTorrentDataOkResponse) SetSuccess(success bool) {
	e.Success = &success
}

func (e *ExportTorrentDataOkResponse) GetSuccess() *bool {
	if e == nil {
		return nil
	}
	return e.Success
}
