package user

type AddReferralToAccountOkResponse struct {
	Data    any     `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

func (a *AddReferralToAccountOkResponse) SetData(data any) {
	a.Data = data
}

func (a *AddReferralToAccountOkResponse) GetData() any {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *AddReferralToAccountOkResponse) SetDetail(detail string) {
	a.Detail = &detail
}

func (a *AddReferralToAccountOkResponse) GetDetail() *string {
	if a == nil {
		return nil
	}
	return a.Detail
}

func (a *AddReferralToAccountOkResponse) SetError(error any) {
	a.Error = error
}

func (a *AddReferralToAccountOkResponse) GetError() any {
	if a == nil {
		return nil
	}
	return a.Error
}

func (a *AddReferralToAccountOkResponse) SetSuccess(success bool) {
	a.Success = &success
}

func (a *AddReferralToAccountOkResponse) GetSuccess() *bool {
	if a == nil {
		return nil
	}
	return a.Success
}
