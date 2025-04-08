package general

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetUpStatusOkResponse struct {
	Data    *util.Nullable[any] `json:"data,omitempty"`
	Detail  *string             `json:"detail,omitempty"`
	Error   *util.Nullable[any] `json:"error,omitempty"`
	Success *bool               `json:"success,omitempty"`
}

func (g *GetUpStatusOkResponse) GetData() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetUpStatusOkResponse) SetData(data util.Nullable[any]) {
	g.Data = &data
}

func (g *GetUpStatusOkResponse) SetDataNull() {
	g.Data = &util.Nullable[any]{IsNull: true}
}

func (g *GetUpStatusOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetUpStatusOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetUpStatusOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetUpStatusOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetUpStatusOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetUpStatusOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetUpStatusOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetUpStatusOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetUpStatusOkResponse to string"
	}
	return string(jsonData)
}
