package notifications

type GetNotificationFeedOkResponse struct {
	Data   []GetNotificationFeedOkResponseData `json:"data,omitempty"`
	Detail *string                             `json:"detail,omitempty"`
}

func (g *GetNotificationFeedOkResponse) SetData(data []GetNotificationFeedOkResponseData) {
	g.Data = data
}

func (g *GetNotificationFeedOkResponse) GetData() []GetNotificationFeedOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetNotificationFeedOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetNotificationFeedOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

type GetNotificationFeedOkResponseData struct {
	AuthId    *string  `json:"auth_id,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	Id        *float64 `json:"id,omitempty"`
	Message   *string  `json:"message,omitempty"`
	Title     *string  `json:"title,omitempty"`
}

func (g *GetNotificationFeedOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetNotificationFeedOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetNotificationFeedOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetNotificationFeedOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetNotificationFeedOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetNotificationFeedOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetNotificationFeedOkResponseData) SetMessage(message string) {
	g.Message = &message
}

func (g *GetNotificationFeedOkResponseData) GetMessage() *string {
	if g == nil {
		return nil
	}
	return g.Message
}

func (g *GetNotificationFeedOkResponseData) SetTitle(title string) {
	g.Title = &title
}

func (g *GetNotificationFeedOkResponseData) GetTitle() *string {
	if g == nil {
		return nil
	}
	return g.Title
}
