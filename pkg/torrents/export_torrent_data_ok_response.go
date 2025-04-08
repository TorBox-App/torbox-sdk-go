package torrents

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type ExportTorrentDataOkResponse struct {
	Data    *string             `json:"data,omitempty"`
	Detail  *string             `json:"detail,omitempty"`
	Error   *util.Nullable[any] `json:"error,omitempty"`
	Success *bool               `json:"success,omitempty"`
}

func (e *ExportTorrentDataOkResponse) GetData() *string {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ExportTorrentDataOkResponse) SetData(data string) {
	e.Data = &data
}

func (e *ExportTorrentDataOkResponse) GetDetail() *string {
	if e == nil {
		return nil
	}
	return e.Detail
}

func (e *ExportTorrentDataOkResponse) SetDetail(detail string) {
	e.Detail = &detail
}

func (e *ExportTorrentDataOkResponse) GetError() *util.Nullable[any] {
	if e == nil {
		return nil
	}
	return e.Error
}

func (e *ExportTorrentDataOkResponse) SetError(error util.Nullable[any]) {
	e.Error = &error
}

func (e *ExportTorrentDataOkResponse) SetErrorNull() {
	e.Error = &util.Nullable[any]{IsNull: true}
}

func (e *ExportTorrentDataOkResponse) GetSuccess() *bool {
	if e == nil {
		return nil
	}
	return e.Success
}

func (e *ExportTorrentDataOkResponse) SetSuccess(success bool) {
	e.Success = &success
}

func (e ExportTorrentDataOkResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExportTorrentDataOkResponse to string"
	}
	return string(jsonData)
}
