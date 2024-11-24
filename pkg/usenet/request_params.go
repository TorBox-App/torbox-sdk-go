package usenet

type RequestDownloadLink1RequestParams struct {
	Token       *string `queryParam:"token"`
	UsenetId    *string `queryParam:"usenet_id"`
	FileId      *string `queryParam:"file_id"`
	ZipLink     *string `queryParam:"zip_link"`
	TorrentFile *string `queryParam:"torrent_file"`
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
func (params *RequestDownloadLink1RequestParams) SetTorrentFile(torrentFile string) {
	params.TorrentFile = &torrentFile
}

type GetUsenetListRequestParams struct {
	BypassCache *string `queryParam:"bypass_cache"`
	Id          *string `queryParam:"id"`
	Offset      *string `queryParam:"offset"`
	Limit       *string `queryParam:"limit"`
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
	Hash   *string `queryParam:"hash"`
	Format *string `queryParam:"format"`
}

func (params *GetUsenetCachedAvailabilityRequestParams) SetHash(hash string) {
	params.Hash = &hash
}
func (params *GetUsenetCachedAvailabilityRequestParams) SetFormat(format string) {
	params.Format = &format
}
