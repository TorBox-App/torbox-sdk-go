package webdownloadsdebrid

type ControlWebDownloadRequestParams struct {
	BypassCache *string `queryParam:"bypass_cache"`
	Id          *string `queryParam:"id"`
}

func (params *ControlWebDownloadRequestParams) SetBypassCache(bypassCache string) {
	params.BypassCache = &bypassCache
}
func (params *ControlWebDownloadRequestParams) SetId(id string) {
	params.Id = &id
}

type RequestDownloadLink2RequestParams struct {
	Token       *string `queryParam:"token"`
	WebId       *string `queryParam:"web_id"`
	FileId      *string `queryParam:"file_id"`
	ZipLink     *string `queryParam:"zip_link"`
	TorrentFile *string `queryParam:"torrent_file"`
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
func (params *RequestDownloadLink2RequestParams) SetTorrentFile(torrentFile string) {
	params.TorrentFile = &torrentFile
}

type GetWebDownloadListRequestParams struct {
	BypassCache *string `queryParam:"bypass_cache"`
	Id          *string `queryParam:"id"`
	Offset      *string `queryParam:"offset"`
	Limit       *string `queryParam:"limit"`
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
	Hash   *string `queryParam:"hash"`
	Format *string `queryParam:"format"`
}

func (params *GetWebDownloadCachedAvailabilityRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetWebDownloadCachedAvailabilityRequestParams) SetFormat(format string) {
	params.Format = &format
}
