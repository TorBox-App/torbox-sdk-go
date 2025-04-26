package webdownloadsdebrid

type ControlWebDownloadRequestParams struct {
	BypassCache *string `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	Id          *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *ControlWebDownloadRequestParams) SetBypassCache(bypassCache string) {
	params.BypassCache = &bypassCache
}
func (params *ControlWebDownloadRequestParams) SetId(id string) {
	params.Id = &id
}

type RequestDownloadLink2RequestParams struct {
	Token    *string `explode:"true" serializationStyle:"form" queryParam:"token"`
	WebId    *string `explode:"true" serializationStyle:"form" queryParam:"web_id"`
	FileId   *string `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	ZipLink  *string `explode:"true" serializationStyle:"form" queryParam:"zip_link"`
	UserIp   *string `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	Redirect *string `explode:"true" serializationStyle:"form" queryParam:"redirect"`
}

func (params *RequestDownloadLink2RequestParams) SetToken(token string) {
	params.Token = &token
}
func (params *RequestDownloadLink2RequestParams) SetWebId(webId string) {
	params.WebId = &webId
}
func (params *RequestDownloadLink2RequestParams) SetFileId(fileId string) {
	params.FileId = &fileId
}
func (params *RequestDownloadLink2RequestParams) SetZipLink(zipLink string) {
	params.ZipLink = &zipLink
}
func (params *RequestDownloadLink2RequestParams) SetUserIp(userIp string) {
	params.UserIp = &userIp
}
func (params *RequestDownloadLink2RequestParams) SetRedirect(redirect string) {
	params.Redirect = &redirect
}

type GetWebDownloadListRequestParams struct {
	BypassCache *string `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	Id          *string `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *string `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *string `explode:"true" serializationStyle:"form" queryParam:"limit"`
}

func (params *GetWebDownloadListRequestParams) SetBypassCache(bypassCache string) {
	params.BypassCache = &bypassCache
}
func (params *GetWebDownloadListRequestParams) SetId(id string) {
	params.Id = &id
}
func (params *GetWebDownloadListRequestParams) SetOffset(offset string) {
	params.Offset = &offset
}
func (params *GetWebDownloadListRequestParams) SetLimit(limit string) {
	params.Limit = &limit
}

type GetWebDownloadCachedAvailabilityRequestParams struct {
	Hash   *string `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Format *string `explode:"true" serializationStyle:"form" queryParam:"format"`
}

func (params *GetWebDownloadCachedAvailabilityRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetWebDownloadCachedAvailabilityRequestParams) SetFormat(format string) {
	params.Format = &format
}
