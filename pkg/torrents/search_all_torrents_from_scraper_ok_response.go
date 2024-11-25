package torrents

type SearchAllTorrentsFromScraperOkResponse struct {
	Data    []SearchAllTorrentsFromScraperOkResponseData `json:"data,omitempty"`
	Detail  *string                                      `json:"detail,omitempty"`
	Error   any                                          `json:"error,omitempty"`
	Success *bool                                        `json:"success,omitempty"`
}

func (s *SearchAllTorrentsFromScraperOkResponse) SetData(data []SearchAllTorrentsFromScraperOkResponseData) {
	s.Data = data
}

func (s *SearchAllTorrentsFromScraperOkResponse) GetData() []SearchAllTorrentsFromScraperOkResponseData {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SearchAllTorrentsFromScraperOkResponse) SetDetail(detail string) {
	s.Detail = &detail
}

func (s *SearchAllTorrentsFromScraperOkResponse) GetDetail() *string {
	if s == nil {
		return nil
	}
	return s.Detail
}

func (s *SearchAllTorrentsFromScraperOkResponse) SetError(error any) {
	s.Error = error
}

func (s *SearchAllTorrentsFromScraperOkResponse) GetError() any {
	if s == nil {
		return nil
	}
	return s.Error
}

func (s *SearchAllTorrentsFromScraperOkResponse) SetSuccess(success bool) {
	s.Success = &success
}

func (s *SearchAllTorrentsFromScraperOkResponse) GetSuccess() *bool {
	if s == nil {
		return nil
	}
	return s.Success
}

type SearchAllTorrentsFromScraperOkResponseData struct {
	Categories    []string `json:"categories,omitempty"`
	Hash          *string  `json:"hash,omitempty"`
	Id            *string  `json:"id,omitempty"`
	Magnet        *string  `json:"magnet,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Peers         *float64 `json:"peers,omitempty"`
	PreferredType *string  `json:"preferred_type,omitempty"`
	Seeders       *float64 `json:"seeders,omitempty"`
	Size          *float64 `json:"size,omitempty"`
	Source        *string  `json:"source,omitempty"`
	TorrentFile   *string  `json:"torrent_file,omitempty"`
	UpdatedAt     *string  `json:"updated_at,omitempty"`
	Website       *string  `json:"website,omitempty"`
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetCategories(categories []string) {
	s.Categories = categories
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetCategories() []string {
	if s == nil {
		return nil
	}
	return s.Categories
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetHash(hash string) {
	s.Hash = &hash
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetHash() *string {
	if s == nil {
		return nil
	}
	return s.Hash
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetId(id string) {
	s.Id = &id
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetMagnet(magnet string) {
	s.Magnet = &magnet
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetMagnet() *string {
	if s == nil {
		return nil
	}
	return s.Magnet
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetName(name string) {
	s.Name = &name
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetPeers(peers float64) {
	s.Peers = &peers
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetPeers() *float64 {
	if s == nil {
		return nil
	}
	return s.Peers
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetPreferredType(preferredType string) {
	s.PreferredType = &preferredType
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetPreferredType() *string {
	if s == nil {
		return nil
	}
	return s.PreferredType
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetSeeders(seeders float64) {
	s.Seeders = &seeders
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetSeeders() *float64 {
	if s == nil {
		return nil
	}
	return s.Seeders
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetSize(size float64) {
	s.Size = &size
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetSize() *float64 {
	if s == nil {
		return nil
	}
	return s.Size
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetSource(source string) {
	s.Source = &source
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetSource() *string {
	if s == nil {
		return nil
	}
	return s.Source
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetTorrentFile(torrentFile string) {
	s.TorrentFile = &torrentFile
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetTorrentFile() *string {
	if s == nil {
		return nil
	}
	return s.TorrentFile
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetUpdatedAt(updatedAt string) {
	s.UpdatedAt = &updatedAt
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetUpdatedAt() *string {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *SearchAllTorrentsFromScraperOkResponseData) SetWebsite(website string) {
	s.Website = &website
}

func (s *SearchAllTorrentsFromScraperOkResponseData) GetWebsite() *string {
	if s == nil {
		return nil
	}
	return s.Website
}
