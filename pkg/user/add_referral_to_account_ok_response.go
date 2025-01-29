package user

import (
	"encoding/json"
)

type AddReferralToAccountOkResponse struct {
	Data    any     `json:"data,omitempty"`
	Detail  *string `json:"detail,omitempty"`
	Error   any     `json:"error,omitempty"`
	Success *bool   `json:"success,omitempty"`
	touched map[string]bool
}

func (a *AddReferralToAccountOkResponse) GetData() any {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *AddReferralToAccountOkResponse) SetData(data any) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = data
}

func (a *AddReferralToAccountOkResponse) SetDataNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = nil
}

func (a *AddReferralToAccountOkResponse) GetDetail() *string {
	if a == nil {
		return nil
	}
	return a.Detail
}

func (a *AddReferralToAccountOkResponse) SetDetail(detail string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Detail"] = true
	a.Detail = &detail
}

func (a *AddReferralToAccountOkResponse) SetDetailNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Detail"] = true
	a.Detail = nil
}

func (a *AddReferralToAccountOkResponse) GetError() any {
	if a == nil {
		return nil
	}
	return a.Error
}

func (a *AddReferralToAccountOkResponse) SetError(error any) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Error"] = true
	a.Error = error
}

func (a *AddReferralToAccountOkResponse) SetErrorNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Error"] = true
	a.Error = nil
}

func (a *AddReferralToAccountOkResponse) GetSuccess() *bool {
	if a == nil {
		return nil
	}
	return a.Success
}

func (a *AddReferralToAccountOkResponse) SetSuccess(success bool) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Success"] = true
	a.Success = &success
}

func (a *AddReferralToAccountOkResponse) SetSuccessNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Success"] = true
	a.Success = nil
}

func (a AddReferralToAccountOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Data"] && a.Data == nil {
		data["data"] = nil
	} else if a.Data != nil {
		data["data"] = a.Data
	}

	if a.touched["Detail"] && a.Detail == nil {
		data["detail"] = nil
	} else if a.Detail != nil {
		data["detail"] = a.Detail
	}

	if a.touched["Error"] && a.Error == nil {
		data["error"] = nil
	} else if a.Error != nil {
		data["error"] = a.Error
	}

	if a.touched["Success"] && a.Success == nil {
		data["success"] = nil
	} else if a.Success != nil {
		data["success"] = a.Success
	}

	return json.Marshal(data)
}

func (a AddReferralToAccountOkResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddReferralToAccountOkResponse to string"
	}
	return string(jsonData)
}
