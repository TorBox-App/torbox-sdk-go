package torrents

type RequestDownloadLinkRequestParams struct {
	Token       *string `queryParam:"token"`
	TorrentId   *string `queryParam:"torrent_id"`
	FileId      *string `queryParam:"file_id"`
	ZipLink     *string `queryParam:"zip_link"`
	TorrentFile *string `queryParam:"torrent_file"`
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
func (params *RequestDownloadLinkRequestParams) SetTorrentFile(torrentFile string) {
	params.TorrentFile = &torrentFile
}

type GetTorrentListRequestParams struct {
	BypassCache *string `queryParam:"bypass_cache"`
	Id          *string `queryParam:"id"`
	Offset      *string `queryParam:"offset"`
	Limit       *string `queryParam:"limit"`
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
	Hash      *string `queryParam:"hash"`
	Format    *string `queryParam:"format"`
	ListFiles *string `queryParam:"list_files"`
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

type SearchAllTorrentsFromScraperRequestParams struct {
	Query *string `queryParam:"query"`
}

func (params *SearchAllTorrentsFromScraperRequestParams) SetQuery(query string) {
	params.Query = &query
}

type ExportTorrentDataRequestParams struct {
	TorrentId *string `queryParam:"torrent_id"`
	Type_     *string `queryParam:"type"`
}

func (params *ExportTorrentDataRequestParams) SetTorrentId(torrentId string) {
	params.TorrentId = &torrentId
}
func (params *ExportTorrentDataRequestParams) SetType_(type_ string) {
	params.Type_ = &type_
}

type GetTorrentInfoRequestParams struct {
	Hash    *string `queryParam:"hash"`
	Timeout *string `queryParam:"timeout"`
}

func (params *GetTorrentInfoRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetTorrentInfoRequestParams) SetTimeout(timeout string) {
	params.Timeout = &timeout
}
