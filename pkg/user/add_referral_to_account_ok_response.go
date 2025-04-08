package user

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type AddReferralToAccountOkResponse struct {
	Data    *util.Nullable[any] `json:"data,omitempty"`
	Detail  *string             `json:"detail,omitempty"`
	Error   *util.Nullable[any] `json:"error,omitempty"`
	Success *bool               `json:"success,omitempty"`
}

func (a *AddReferralToAccountOkResponse) GetData() *util.Nullable[any] {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *AddReferralToAccountOkResponse) SetData(data util.Nullable[any]) {
	a.Data = &data
}

func (a *AddReferralToAccountOkResponse) SetDataNull() {
	a.Data = &util.Nullable[any]{IsNull: true}
}

func (a *AddReferralToAccountOkResponse) GetDetail() *string {
	if a == nil {
		return nil
	}
	return a.Detail
}

func (a *AddReferralToAccountOkResponse) SetDetail(detail string) {
	a.Detail = &detail
}

func (a *AddReferralToAccountOkResponse) GetError() *util.Nullable[any] {
	if a == nil {
		return nil
	}
	return a.Error
}

func (a *AddReferralToAccountOkResponse) SetError(error util.Nullable[any]) {
	a.Error = &error
}

func (a *AddReferralToAccountOkResponse) SetErrorNull() {
	a.Error = &util.Nullable[any]{IsNull: true}
}

func (a *AddReferralToAccountOkResponse) GetSuccess() *bool {
	if a == nil {
		return nil
	}
	return a.Success
}

func (a *AddReferralToAccountOkResponse) SetSuccess(success bool) {
	a.Success = &success
}

func (a AddReferralToAccountOkResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddReferralToAccountOkResponse to string"
	}
	return string(jsonData)
}
