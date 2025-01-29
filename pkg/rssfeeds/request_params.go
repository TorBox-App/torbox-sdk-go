package rssfeeds

type GetUserRssFeedsRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *GetUserRssFeedsRequestParams) SetId(id string) {
	params.Id = &id
}

type GetRssFeedItemsRequestParams struct {
	RssFeedId *string `explode:"true" serializationStyle:"form" queryParam:"rss_feed_id"`
}

func (params *GetRssFeedItemsRequestParams) SetRssFeedId(rssFeedId string) {
	params.RssFeedId = &rssFeedId
}
