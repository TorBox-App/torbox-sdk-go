package torrents

type ExportTorrentDataOkResponse struct {
	Data   *string `json:"data,omitempty"`
	Detail *string `json:"detail,omitempty"`
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
