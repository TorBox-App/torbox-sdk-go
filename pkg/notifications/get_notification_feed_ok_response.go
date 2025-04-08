package notifications

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetNotificationFeedOkResponse struct {
	Data    []GetNotificationFeedOkResponseData `json:"data,omitempty"`
	Detail  *string                             `json:"detail,omitempty"`
	Error   *util.Nullable[any]                 `json:"error,omitempty"`
	Success *bool                               `json:"success,omitempty"`
}

func (g *GetNotificationFeedOkResponse) GetData() []GetNotificationFeedOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetNotificationFeedOkResponse) SetData(data []GetNotificationFeedOkResponseData) {
	g.Data = data
}

func (g *GetNotificationFeedOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetNotificationFeedOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetNotificationFeedOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetNotificationFeedOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetNotificationFeedOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetNotificationFeedOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetNotificationFeedOkResponse) SetSuccess(success bool) {
	g.Success = &success
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
}

func (g *GetNotificationFeedOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetNotificationFeedOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetNotificationFeedOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetNotificationFeedOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetNotificationFeedOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetNotificationFeedOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetNotificationFeedOkResponseData) GetMessage() *string {
	if g == nil {
		return nil
	}
	return g.Message
}

func (g *GetNotificationFeedOkResponseData) SetMessage(message string) {
	g.Message = &message
}

func (g *GetNotificationFeedOkResponseData) GetTitle() *string {
	if g == nil {
		return nil
	}
	return g.Title
}

func (g *GetNotificationFeedOkResponseData) SetTitle(title string) {
	g.Title = &title
}

func (g GetNotificationFeedOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetNotificationFeedOkResponseData to string"
	}
	return string(jsonData)
}
