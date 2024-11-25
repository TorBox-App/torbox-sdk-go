package torrents

type GetTorrentListOkResponse struct {
	Data    []GetTorrentListOkResponseData `json:"data,omitempty"`
	Detail  *string                        `json:"detail,omitempty"`
	Error   any                            `json:"error,omitempty"`
	Success *bool                          `json:"success,omitempty"`
}

func (g *GetTorrentListOkResponse) SetData(data []GetTorrentListOkResponseData) {
	g.Data = data
}

func (g *GetTorrentListOkResponse) GetData() []GetTorrentListOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentListOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetTorrentListOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentListOkResponse) SetError(error any) {
	g.Error = error
}

func (g *GetTorrentListOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentListOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetTorrentListOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

type GetTorrentListOkResponseData struct {
	Active           *bool        `json:"active,omitempty"`
	AuthId           *string      `json:"auth_id,omitempty"`
	Availability     *float64     `json:"availability,omitempty"`
	CreatedAt        *string      `json:"created_at,omitempty"`
	DownloadFinished *bool        `json:"download_finished,omitempty"`
	DownloadPresent  *bool        `json:"download_present,omitempty"`
	DownloadSpeed    *float64     `json:"download_speed,omitempty"`
	DownloadState    *string      `json:"download_state,omitempty"`
	Eta              *float64     `json:"eta,omitempty"`
	ExpiresAt        *string      `json:"expires_at,omitempty"`
	Files            []DataFiles1 `json:"files,omitempty"`
	Hash             *string      `json:"hash,omitempty"`
	Id               *float64     `json:"id,omitempty"`
	InactiveCheck    *float64     `json:"inactive_check,omitempty"`
	Magnet           *string      `json:"magnet,omitempty"`
	Name             *string      `json:"name,omitempty"`
	Peers            *float64     `json:"peers,omitempty"`
	Progress         *float64     `json:"progress,omitempty"`
	Ratio            *float64     `json:"ratio,omitempty"`
	Seeds            *float64     `json:"seeds,omitempty"`
	Server           *float64     `json:"server,omitempty"`
	Size             *float64     `json:"size,omitempty"`
	TorrentFile      *bool        `json:"torrent_file,omitempty"`
	UpdatedAt        *string      `json:"updated_at,omitempty"`
	UploadSpeed      *float64     `json:"upload_speed,omitempty"`
}

func (g *GetTorrentListOkResponseData) SetActive(active bool) {
	g.Active = &active
}

func (g *GetTorrentListOkResponseData) GetActive() *bool {
	if g == nil {
		return nil
	}
	return g.Active
}

func (g *GetTorrentListOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetTorrentListOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetTorrentListOkResponseData) SetAvailability(availability float64) {
	g.Availability = &availability
}

func (g *GetTorrentListOkResponseData) GetAvailability() *float64 {
	if g == nil {
		return nil
	}
	return g.Availability
}

func (g *GetTorrentListOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetTorrentListOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetTorrentListOkResponseData) SetDownloadFinished(downloadFinished bool) {
	g.DownloadFinished = &downloadFinished
}

func (g *GetTorrentListOkResponseData) GetDownloadFinished() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadFinished
}

func (g *GetTorrentListOkResponseData) SetDownloadPresent(downloadPresent bool) {
	g.DownloadPresent = &downloadPresent
}

func (g *GetTorrentListOkResponseData) GetDownloadPresent() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadPresent
}

func (g *GetTorrentListOkResponseData) SetDownloadSpeed(downloadSpeed float64) {
	g.DownloadSpeed = &downloadSpeed
}

func (g *GetTorrentListOkResponseData) GetDownloadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.DownloadSpeed
}

func (g *GetTorrentListOkResponseData) SetDownloadState(downloadState string) {
	g.DownloadState = &downloadState
}

func (g *GetTorrentListOkResponseData) GetDownloadState() *string {
	if g == nil {
		return nil
	}
	return g.DownloadState
}

func (g *GetTorrentListOkResponseData) SetEta(eta float64) {
	g.Eta = &eta
}

func (g *GetTorrentListOkResponseData) GetEta() *float64 {
	if g == nil {
		return nil
	}
	return g.Eta
}

func (g *GetTorrentListOkResponseData) SetExpiresAt(expiresAt string) {
	g.ExpiresAt = &expiresAt
}

func (g *GetTorrentListOkResponseData) GetExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.ExpiresAt
}

func (g *GetTorrentListOkResponseData) SetFiles(files []DataFiles1) {
	g.Files = files
}

func (g *GetTorrentListOkResponseData) GetFiles() []DataFiles1 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetTorrentListOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetTorrentListOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentListOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetTorrentListOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetTorrentListOkResponseData) SetInactiveCheck(inactiveCheck float64) {
	g.InactiveCheck = &inactiveCheck
}

func (g *GetTorrentListOkResponseData) GetInactiveCheck() *float64 {
	if g == nil {
		return nil
	}
	return g.InactiveCheck
}

func (g *GetTorrentListOkResponseData) SetMagnet(magnet string) {
	g.Magnet = &magnet
}

func (g *GetTorrentListOkResponseData) GetMagnet() *string {
	if g == nil {
		return nil
	}
	return g.Magnet
}

func (g *GetTorrentListOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetTorrentListOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentListOkResponseData) SetPeers(peers float64) {
	g.Peers = &peers
}

func (g *GetTorrentListOkResponseData) GetPeers() *float64 {
	if g == nil {
		return nil
	}
	return g.Peers
}

func (g *GetTorrentListOkResponseData) SetProgress(progress float64) {
	g.Progress = &progress
}

func (g *GetTorrentListOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetTorrentListOkResponseData) SetRatio(ratio float64) {
	g.Ratio = &ratio
}

func (g *GetTorrentListOkResponseData) GetRatio() *float64 {
	if g == nil {
		return nil
	}
	return g.Ratio
}

func (g *GetTorrentListOkResponseData) SetSeeds(seeds float64) {
	g.Seeds = &seeds
}

func (g *GetTorrentListOkResponseData) GetSeeds() *float64 {
	if g == nil {
		return nil
	}
	return g.Seeds
}

func (g *GetTorrentListOkResponseData) SetServer(server float64) {
	g.Server = &server
}

func (g *GetTorrentListOkResponseData) GetServer() *float64 {
	if g == nil {
		return nil
	}
	return g.Server
}

func (g *GetTorrentListOkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetTorrentListOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentListOkResponseData) SetTorrentFile(torrentFile bool) {
	g.TorrentFile = &torrentFile
}

func (g *GetTorrentListOkResponseData) GetTorrentFile() *bool {
	if g == nil {
		return nil
	}
	return g.TorrentFile
}

func (g *GetTorrentListOkResponseData) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GetTorrentListOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetTorrentListOkResponseData) SetUploadSpeed(uploadSpeed float64) {
	g.UploadSpeed = &uploadSpeed
}

func (g *GetTorrentListOkResponseData) GetUploadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.UploadSpeed
}

type DataFiles1 struct {
	Id        *float64 `json:"id,omitempty"`
	Md5       *string  `json:"md5,omitempty"`
	Mimetype  *string  `json:"mimetype,omitempty"`
	Name      *string  `json:"name,omitempty"`
	S3Path    *string  `json:"s3_path,omitempty"`
	ShortName *string  `json:"short_name,omitempty"`
	Size      *float64 `json:"size,omitempty"`
}

func (d *DataFiles1) SetId(id float64) {
	d.Id = &id
}

func (d *DataFiles1) GetId() *float64 {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DataFiles1) SetMd5(md5 string) {
	d.Md5 = &md5
}

func (d *DataFiles1) GetMd5() *string {
	if d == nil {
		return nil
	}
	return d.Md5
}

func (d *DataFiles1) SetMimetype(mimetype string) {
	d.Mimetype = &mimetype
}

func (d *DataFiles1) GetMimetype() *string {
	if d == nil {
		return nil
	}
	return d.Mimetype
}

func (d *DataFiles1) SetName(name string) {
	d.Name = &name
}

func (d *DataFiles1) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles1) SetS3Path(s3Path string) {
	d.S3Path = &s3Path
}

func (d *DataFiles1) GetS3Path() *string {
	if d == nil {
		return nil
	}
	return d.S3Path
}

func (d *DataFiles1) SetShortName(shortName string) {
	d.ShortName = &shortName
}

func (d *DataFiles1) GetShortName() *string {
	if d == nil {
		return nil
	}
	return d.ShortName
}

func (d *DataFiles1) SetSize(size float64) {
	d.Size = &size
}

func (d *DataFiles1) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}
