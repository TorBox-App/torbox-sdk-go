package torrents

import (
	"encoding/json"
)

type ExportTorrentDataOkResponse struct {
	Data    *string `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
	touched map[string]bool
}

func (e *ExportTorrentDataOkResponse) GetData() *string {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ExportTorrentDataOkResponse) SetData(data string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Data"] = true
	e.Data = &data
}

func (e *ExportTorrentDataOkResponse) SetDataNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Data"] = true
	e.Data = nil
}

func (e *ExportTorrentDataOkResponse) GetDetail() *string {
	if e == nil {
		return nil
	}
	return e.Detail
}

func (e *ExportTorrentDataOkResponse) SetDetail(detail string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Detail"] = true
	e.Detail = &detail
}

func (e *ExportTorrentDataOkResponse) SetDetailNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Detail"] = true
	e.Detail = nil
}

func (e *ExportTorrentDataOkResponse) GetError() any {
	if e == nil {
		return nil
	}
	return e.Error
}

func (e *ExportTorrentDataOkResponse) SetError(error any) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Error"] = true
	e.Error = error
}

func (e *ExportTorrentDataOkResponse) SetErrorNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Error"] = true
	e.Error = nil
}

func (e *ExportTorrentDataOkResponse) GetSuccess() *bool {
	if e == nil {
		return nil
	}
	return e.Success
}

func (e *ExportTorrentDataOkResponse) SetSuccess(success bool) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Success"] = true
	e.Success = &success
}

func (e *ExportTorrentDataOkResponse) SetSuccessNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Success"] = true
	e.Success = nil
}

func (e ExportTorrentDataOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["Data"] && e.Data == nil {
		data["data"] = nil
	} else if e.Data != nil {
		data["data"] = e.Data
	}

	if e.touched["Detail"] && e.Detail == nil {
		data["detail"] = nil
	} else if e.Detail != nil {
		data["detail"] = e.Detail
	}

	if e.touched["Error"] && e.Error == nil {
		data["error"] = nil
	} else if e.Error != nil {
		data["error"] = e.Error
	}

	if e.touched["Success"] && e.Success == nil {
		data["success"] = nil
	} else if e.Success != nil {
		data["success"] = e.Success
	}

	return json.Marshal(data)
}

func (e ExportTorrentDataOkResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExportTorrentDataOkResponse to string"
	}
	return string(jsonData)
}
