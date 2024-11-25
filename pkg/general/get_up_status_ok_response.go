package general

type GetUpStatusOkResponse struct {
	Data    any     `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

func (g *GetUpStatusOkResponse) SetData(data any) {
	g.Data = data
}

func (g *GetUpStatusOkResponse) GetData() any {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetUpStatusOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetUpStatusOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetUpStatusOkResponse) SetError(error any) {
	g.Error = error
}

func (g *GetUpStatusOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetUpStatusOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetUpStatusOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}
