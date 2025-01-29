package notifications

type GetRssNotificationFeedRequestParams struct {
	Token *string `explode:"true" serializationStyle:"form" queryParam:"token"`
}

func (params *GetRssNotificationFeedRequestParams) SetToken(token string) {
	params.Token = &token
}
