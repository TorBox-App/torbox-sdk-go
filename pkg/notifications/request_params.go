package notifications

type GetRssNotificationFeedRequestParams struct {
	Token *string `queryParam:"token"`
}

func (params *GetRssNotificationFeedRequestParams) SetToken(token string) {
	params.Token = &token
}
