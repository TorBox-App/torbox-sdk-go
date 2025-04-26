package usenet

type RequestDownloadLink1RequestParams struct {
	Token    *string `explode:"true" serializationStyle:"form" queryParam:"token"`
	UsenetId *string `explode:"true" serializationStyle:"form" queryParam:"usenet_id"`
	FileId   *string `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	ZipLink  *string `explode:"true" serializationStyle:"form" queryParam:"zip_link"`
	UserIp   *string `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	Redirect *string `explode:"true" serializationStyle:"form" queryParam:"redirect"`
}

func (params *RequestDownloadLink1RequestParams) SetToken(token string) {
	params.Token = &token
}
func (params *RequestDownloadLink1RequestParams) SetUsenetId(usenetId string) {
	params.UsenetId = &usenetId
}
func (params *RequestDownloadLink1RequestParams) SetFileId(fileId string) {
	params.FileId = &fileId
}
func (params *RequestDownloadLink1RequestParams) SetZipLink(zipLink string) {
	params.ZipLink = &zipLink
}
func (params *RequestDownloadLink1RequestParams) SetUserIp(userIp string) {
	params.UserIp = &userIp
}
func (params *RequestDownloadLink1RequestParams) SetRedirect(redirect string) {
	params.Redirect = &redirect
}

type GetUsenetListRequestParams struct {
	BypassCache *string `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	Id          *string `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *string `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *string `explode:"true" serializationStyle:"form" queryParam:"limit"`
}

func (params *GetUsenetListRequestParams) SetBypassCache(bypassCache string) {
	params.BypassCache = &bypassCache
}
func (params *GetUsenetListRequestParams) SetId(id string) {
	params.Id = &id
}
func (params *GetUsenetListRequestParams) SetOffset(offset string) {
	params.Offset = &offset
}
func (params *GetUsenetListRequestParams) SetLimit(limit string) {
	params.Limit = &limit
}

type GetUsenetCachedAvailabilityRequestParams struct {
	Hash   *string `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Format *string `explode:"true" serializationStyle:"form" queryParam:"format"`
}

func (params *GetUsenetCachedAvailabilityRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetUsenetCachedAvailabilityRequestParams) SetFormat(format string) {
	params.Format = &format
}
