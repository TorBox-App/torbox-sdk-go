package queued

type GetQueuedDownloadsRequestParams struct {
	BypassCache *string `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	Id          *string `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *string `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *string `explode:"true" serializationStyle:"form" queryParam:"limit"`
	Type_       *string `explode:"true" serializationStyle:"form" queryParam:"type"`
}

func (params *GetQueuedDownloadsRequestParams) SetBypassCache(bypassCache string) {
	params.BypassCache = &bypassCache
}
func (params *GetQueuedDownloadsRequestParams) SetId(id string) {
	params.Id = &id
}
func (params *GetQueuedDownloadsRequestParams) SetOffset(offset string) {
	params.Offset = &offset
}
func (params *GetQueuedDownloadsRequestParams) SetLimit(limit string) {
	params.Limit = &limit
}
func (params *GetQueuedDownloadsRequestParams) SetType_(type_ string) {
	params.Type_ = &type_
}
