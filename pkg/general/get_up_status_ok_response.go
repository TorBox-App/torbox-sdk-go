package general

type GetUpStatusOkResponse struct {
	Detail *string `json:"detail,omitempty"`
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
