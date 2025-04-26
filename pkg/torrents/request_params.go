package torrents

type RequestDownloadLinkRequestParams struct {
	Token     *string `explode:"true" serializationStyle:"form" queryParam:"token"`
	TorrentId *string `explode:"true" serializationStyle:"form" queryParam:"torrent_id"`
	FileId    *string `explode:"true" serializationStyle:"form" queryParam:"file_id"`
	ZipLink   *string `explode:"true" serializationStyle:"form" queryParam:"zip_link"`
	UserIp    *string `explode:"true" serializationStyle:"form" queryParam:"user_ip"`
	Redirect  *string `explode:"true" serializationStyle:"form" queryParam:"redirect"`
}

func (params *RequestDownloadLinkRequestParams) SetToken(token string) {
	params.Token = &token
}
func (params *RequestDownloadLinkRequestParams) SetTorrentId(torrentId string) {
	params.TorrentId = &torrentId
}
func (params *RequestDownloadLinkRequestParams) SetFileId(fileId string) {
	params.FileId = &fileId
}
func (params *RequestDownloadLinkRequestParams) SetZipLink(zipLink string) {
	params.ZipLink = &zipLink
}
func (params *RequestDownloadLinkRequestParams) SetUserIp(userIp string) {
	params.UserIp = &userIp
}
func (params *RequestDownloadLinkRequestParams) SetRedirect(redirect string) {
	params.Redirect = &redirect
}

type GetTorrentListRequestParams struct {
	BypassCache *string `explode:"true" serializationStyle:"form" queryParam:"bypass_cache"`
	Id          *string `explode:"true" serializationStyle:"form" queryParam:"id"`
	Offset      *string `explode:"true" serializationStyle:"form" queryParam:"offset"`
	Limit       *string `explode:"true" serializationStyle:"form" queryParam:"limit"`
}

func (params *GetTorrentListRequestParams) SetBypassCache(bypassCache string) {
	params.BypassCache = &bypassCache
}
func (params *GetTorrentListRequestParams) SetId(id string) {
	params.Id = &id
}
func (params *GetTorrentListRequestParams) SetOffset(offset string) {
	params.Offset = &offset
}
func (params *GetTorrentListRequestParams) SetLimit(limit string) {
	params.Limit = &limit
}

type GetTorrentCachedAvailabilityRequestParams struct {
	Hash      *string `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Format    *string `explode:"true" serializationStyle:"form" queryParam:"format"`
	ListFiles *string `explode:"true" serializationStyle:"form" queryParam:"list_files"`
}

func (params *GetTorrentCachedAvailabilityRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetTorrentCachedAvailabilityRequestParams) SetFormat(format string) {
	params.Format = &format
}
func (params *GetTorrentCachedAvailabilityRequestParams) SetListFiles(listFiles string) {
	params.ListFiles = &listFiles
}

type ExportTorrentDataRequestParams struct {
	TorrentId *string `explode:"true" serializationStyle:"form" queryParam:"torrent_id"`
	Type_     *string `explode:"true" serializationStyle:"form" queryParam:"type"`
}

func (params *ExportTorrentDataRequestParams) SetTorrentId(torrentId string) {
	params.TorrentId = &torrentId
}
func (params *ExportTorrentDataRequestParams) SetType_(type_ string) {
	params.Type_ = &type_
}

type GetTorrentInfoRequestParams struct {
	Hash    *string `explode:"true" serializationStyle:"form" queryParam:"hash"`
	Timeout *string `explode:"true" serializationStyle:"form" queryParam:"timeout"`
}

func (params *GetTorrentInfoRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetTorrentInfoRequestParams) SetTimeout(timeout string) {
	params.Timeout = &timeout
}
