package notifications

import (
	"encoding/json"
)

type GetNotificationFeedOkResponse struct {
	Data    []GetNotificationFeedOkResponseData `json:"data,omitempty"`
	Detail  *string                             `json:"detail,omitempty"`
	Error   any                                 `json:"error,omitempty"`
	Success *bool                               `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetNotificationFeedOkResponse) GetData() []GetNotificationFeedOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetNotificationFeedOkResponse) SetData(data []GetNotificationFeedOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = data
}

func (g *GetNotificationFeedOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetNotificationFeedOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetNotificationFeedOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetNotificationFeedOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetNotificationFeedOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetNotificationFeedOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetNotificationFeedOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetNotificationFeedOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetNotificationFeedOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetNotificationFeedOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetNotificationFeedOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Data"] && g.Data == nil {
		data["data"] = nil
	} else if g.Data != nil {
		data["data"] = g.Data
	}

	if g.touched["Detail"] && g.Detail == nil {
		data["detail"] = nil
	} else if g.Detail != nil {
		data["detail"] = g.Detail
	}

	if g.touched["Error"] && g.Error == nil {
		data["error"] = nil
	} else if g.Error != nil {
		data["error"] = g.Error
	}

	if g.touched["Success"] && g.Success == nil {
		data["success"] = nil
	} else if g.Success != nil {
		data["success"] = g.Success
	}

	return json.Marshal(data)
}

func (g GetNotificationFeedOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetNotificationFeedOkResponse to string"
	}
	return string(jsonData)
}

type GetNotificationFeedOkResponseData struct {
	AuthId    *string  `json:"auth_id,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Message   *string  `json:"message,omitempty"`
	Title     *string  `json:"title,omitempty"`
	touched   map[string]bool
}

func (g *GetNotificationFeedOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetNotificationFeedOkResponseData) SetAuthId(authId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = &authId
}

func (g *GetNotificationFeedOkResponseData) SetAuthIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = nil
}

func (g *GetNotificationFeedOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetNotificationFeedOkResponseData) SetCreatedAt(createdAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = &createdAt
}

func (g *GetNotificationFeedOkResponseData) SetCreatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = nil
}

func (g *GetNotificationFeedOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetNotificationFeedOkResponseData) SetId(id float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GetNotificationFeedOkResponseData) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GetNotificationFeedOkResponseData) GetMessage() *string {
	if g == nil {
		return nil
	}
	return g.Message
}

func (g *GetNotificationFeedOkResponseData) SetMessage(message string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Message"] = true
	g.Message = &message
}

func (g *GetNotificationFeedOkResponseData) SetMessageNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Message"] = true
	g.Message = nil
}

func (g *GetNotificationFeedOkResponseData) GetTitle() *string {
	if g == nil {
		return nil
	}
	return g.Title
}

func (g *GetNotificationFeedOkResponseData) SetTitle(title string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Title"] = true
	g.Title = &title
}

func (g *GetNotificationFeedOkResponseData) SetTitleNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Title"] = true
	g.Title = nil
}

func (g GetNotificationFeedOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["AuthId"] && g.AuthId == nil {
		data["auth_id"] = nil
	} else if g.AuthId != nil {
		data["auth_id"] = g.AuthId
	}

	if g.touched["CreatedAt"] && g.CreatedAt == nil {
		data["created_at"] = nil
	} else if g.CreatedAt != nil {
		data["created_at"] = g.CreatedAt
	}

	if g.touched["Id"] && g.Id == nil {
		data["id"] = nil
	} else if g.Id != nil {
		data["id"] = g.Id
	}

	if g.touched["Message"] && g.Message == nil {
		data["message"] = nil
	} else if g.Message != nil {
		data["message"] = g.Message
	}

	if g.touched["Title"] && g.Title == nil {
		data["title"] = nil
	} else if g.Title != nil {
		data["title"] = g.Title
	}

	return json.Marshal(data)
}

func (g GetNotificationFeedOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetNotificationFeedOkResponseData to string"
	}
	return string(jsonData)
}
