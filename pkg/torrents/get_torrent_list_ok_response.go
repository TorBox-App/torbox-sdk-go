package torrents

import (
	"encoding/json"
)

type GetTorrentListOkResponse struct {
	Data    []GetTorrentListOkResponseData `json:"data,omitempty"`
	Detail  *string                        `json:"detail,omitempty"`
	Error   any                            `json:"error,omitempty"`
	Success *bool                          `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetTorrentListOkResponse) GetData() []GetTorrentListOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetTorrentListOkResponse) SetData(data []GetTorrentListOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = data
}

func (g *GetTorrentListOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetTorrentListOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetTorrentListOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetTorrentListOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetTorrentListOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetTorrentListOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetTorrentListOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetTorrentListOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetTorrentListOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetTorrentListOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetTorrentListOkResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Data"] && g.Data == nil {
		data["data"] = nil
	} else if g.Data != nil {
		data["data"] = g.Data
	}

	if g.touched["Detail"] && g.Detail == nil {
		data["detail"] = nil
	} else if g.Detail != nil {
		data["detail"] = g.Detail
	}

	if g.touched["Error"] && g.Error == nil {
		data["error"] = nil
	} else if g.Error != nil {
		data["error"] = g.Error
	}

	if g.touched["Success"] && g.Success == nil {
		data["success"] = nil
	} else if g.Success != nil {
		data["success"] = g.Success
	}

	return json.Marshal(data)
}

func (g GetTorrentListOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentListOkResponse to string"
	}
	return string(jsonData)
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
	touched          map[string]bool
}

func (g *GetTorrentListOkResponseData) GetActive() *bool {
	if g == nil {
		return nil
	}
	return g.Active
}

func (g *GetTorrentListOkResponseData) SetActive(active bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Active"] = true
	g.Active = &active
}

func (g *GetTorrentListOkResponseData) SetActiveNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Active"] = true
	g.Active = nil
}

func (g *GetTorrentListOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetTorrentListOkResponseData) SetAuthId(authId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = &authId
}

func (g *GetTorrentListOkResponseData) SetAuthIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = nil
}

func (g *GetTorrentListOkResponseData) GetAvailability() *float64 {
	if g == nil {
		return nil
	}
	return g.Availability
}

func (g *GetTorrentListOkResponseData) SetAvailability(availability float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Availability"] = true
	g.Availability = &availability
}

func (g *GetTorrentListOkResponseData) SetAvailabilityNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Availability"] = true
	g.Availability = nil
}

func (g *GetTorrentListOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetTorrentListOkResponseData) SetCreatedAt(createdAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = &createdAt
}

func (g *GetTorrentListOkResponseData) SetCreatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = nil
}

func (g *GetTorrentListOkResponseData) GetDownloadFinished() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadFinished
}

func (g *GetTorrentListOkResponseData) SetDownloadFinished(downloadFinished bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadFinished"] = true
	g.DownloadFinished = &downloadFinished
}

func (g *GetTorrentListOkResponseData) SetDownloadFinishedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadFinished"] = true
	g.DownloadFinished = nil
}

func (g *GetTorrentListOkResponseData) GetDownloadPresent() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadPresent
}

func (g *GetTorrentListOkResponseData) SetDownloadPresent(downloadPresent bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadPresent"] = true
	g.DownloadPresent = &downloadPresent
}

func (g *GetTorrentListOkResponseData) SetDownloadPresentNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadPresent"] = true
	g.DownloadPresent = nil
}

func (g *GetTorrentListOkResponseData) GetDownloadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.DownloadSpeed
}

func (g *GetTorrentListOkResponseData) SetDownloadSpeed(downloadSpeed float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadSpeed"] = true
	g.DownloadSpeed = &downloadSpeed
}

func (g *GetTorrentListOkResponseData) SetDownloadSpeedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadSpeed"] = true
	g.DownloadSpeed = nil
}

func (g *GetTorrentListOkResponseData) GetDownloadState() *string {
	if g == nil {
		return nil
	}
	return g.DownloadState
}

func (g *GetTorrentListOkResponseData) SetDownloadState(downloadState string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadState"] = true
	g.DownloadState = &downloadState
}

func (g *GetTorrentListOkResponseData) SetDownloadStateNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadState"] = true
	g.DownloadState = nil
}

func (g *GetTorrentListOkResponseData) GetEta() *float64 {
	if g == nil {
		return nil
	}
	return g.Eta
}

func (g *GetTorrentListOkResponseData) SetEta(eta float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Eta"] = true
	g.Eta = &eta
}

func (g *GetTorrentListOkResponseData) SetEtaNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Eta"] = true
	g.Eta = nil
}

func (g *GetTorrentListOkResponseData) GetExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.ExpiresAt
}

func (g *GetTorrentListOkResponseData) SetExpiresAt(expiresAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ExpiresAt"] = true
	g.ExpiresAt = &expiresAt
}

func (g *GetTorrentListOkResponseData) SetExpiresAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ExpiresAt"] = true
	g.ExpiresAt = nil
}

func (g *GetTorrentListOkResponseData) GetFiles() []DataFiles1 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetTorrentListOkResponseData) SetFiles(files []DataFiles1) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Files"] = true
	g.Files = files
}

func (g *GetTorrentListOkResponseData) SetFilesNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Files"] = true
	g.Files = nil
}

func (g *GetTorrentListOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetTorrentListOkResponseData) SetHash(hash string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = &hash
}

func (g *GetTorrentListOkResponseData) SetHashNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = nil
}

func (g *GetTorrentListOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetTorrentListOkResponseData) SetId(id float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GetTorrentListOkResponseData) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GetTorrentListOkResponseData) GetInactiveCheck() *float64 {
	if g == nil {
		return nil
	}
	return g.InactiveCheck
}

func (g *GetTorrentListOkResponseData) SetInactiveCheck(inactiveCheck float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["InactiveCheck"] = true
	g.InactiveCheck = &inactiveCheck
}

func (g *GetTorrentListOkResponseData) SetInactiveCheckNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["InactiveCheck"] = true
	g.InactiveCheck = nil
}

func (g *GetTorrentListOkResponseData) GetMagnet() *string {
	if g == nil {
		return nil
	}
	return g.Magnet
}

func (g *GetTorrentListOkResponseData) SetMagnet(magnet string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Magnet"] = true
	g.Magnet = &magnet
}

func (g *GetTorrentListOkResponseData) SetMagnetNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Magnet"] = true
	g.Magnet = nil
}

func (g *GetTorrentListOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetTorrentListOkResponseData) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GetTorrentListOkResponseData) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}

func (g *GetTorrentListOkResponseData) GetPeers() *float64 {
	if g == nil {
		return nil
	}
	return g.Peers
}

func (g *GetTorrentListOkResponseData) SetPeers(peers float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Peers"] = true
	g.Peers = &peers
}

func (g *GetTorrentListOkResponseData) SetPeersNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Peers"] = true
	g.Peers = nil
}

func (g *GetTorrentListOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetTorrentListOkResponseData) SetProgress(progress float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = &progress
}

func (g *GetTorrentListOkResponseData) SetProgressNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = nil
}

func (g *GetTorrentListOkResponseData) GetRatio() *float64 {
	if g == nil {
		return nil
	}
	return g.Ratio
}

func (g *GetTorrentListOkResponseData) SetRatio(ratio float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Ratio"] = true
	g.Ratio = &ratio
}

func (g *GetTorrentListOkResponseData) SetRatioNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Ratio"] = true
	g.Ratio = nil
}

func (g *GetTorrentListOkResponseData) GetSeeds() *float64 {
	if g == nil {
		return nil
	}
	return g.Seeds
}

func (g *GetTorrentListOkResponseData) SetSeeds(seeds float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Seeds"] = true
	g.Seeds = &seeds
}

func (g *GetTorrentListOkResponseData) SetSeedsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Seeds"] = true
	g.Seeds = nil
}

func (g *GetTorrentListOkResponseData) GetServer() *float64 {
	if g == nil {
		return nil
	}
	return g.Server
}

func (g *GetTorrentListOkResponseData) SetServer(server float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Server"] = true
	g.Server = &server
}

func (g *GetTorrentListOkResponseData) SetServerNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Server"] = true
	g.Server = nil
}

func (g *GetTorrentListOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetTorrentListOkResponseData) SetSize(size float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = &size
}

func (g *GetTorrentListOkResponseData) SetSizeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = nil
}

func (g *GetTorrentListOkResponseData) GetTorrentFile() *bool {
	if g == nil {
		return nil
	}
	return g.TorrentFile
}

func (g *GetTorrentListOkResponseData) SetTorrentFile(torrentFile bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TorrentFile"] = true
	g.TorrentFile = &torrentFile
}

func (g *GetTorrentListOkResponseData) SetTorrentFileNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TorrentFile"] = true
	g.TorrentFile = nil
}

func (g *GetTorrentListOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetTorrentListOkResponseData) SetUpdatedAt(updatedAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = &updatedAt
}

func (g *GetTorrentListOkResponseData) SetUpdatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = nil
}

func (g *GetTorrentListOkResponseData) GetUploadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.UploadSpeed
}

func (g *GetTorrentListOkResponseData) SetUploadSpeed(uploadSpeed float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UploadSpeed"] = true
	g.UploadSpeed = &uploadSpeed
}

func (g *GetTorrentListOkResponseData) SetUploadSpeedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UploadSpeed"] = true
	g.UploadSpeed = nil
}

func (g GetTorrentListOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Active"] && g.Active == nil {
		data["active"] = nil
	} else if g.Active != nil {
		data["active"] = g.Active
	}

	if g.touched["AuthId"] && g.AuthId == nil {
		data["auth_id"] = nil
	} else if g.AuthId != nil {
		data["auth_id"] = g.AuthId
	}

	if g.touched["Availability"] && g.Availability == nil {
		data["availability"] = nil
	} else if g.Availability != nil {
		data["availability"] = g.Availability
	}

	if g.touched["CreatedAt"] && g.CreatedAt == nil {
		data["created_at"] = nil
	} else if g.CreatedAt != nil {
		data["created_at"] = g.CreatedAt
	}

	if g.touched["DownloadFinished"] && g.DownloadFinished == nil {
		data["download_finished"] = nil
	} else if g.DownloadFinished != nil {
		data["download_finished"] = g.DownloadFinished
	}

	if g.touched["DownloadPresent"] && g.DownloadPresent == nil {
		data["download_present"] = nil
	} else if g.DownloadPresent != nil {
		data["download_present"] = g.DownloadPresent
	}

	if g.touched["DownloadSpeed"] && g.DownloadSpeed == nil {
		data["download_speed"] = nil
	} else if g.DownloadSpeed != nil {
		data["download_speed"] = g.DownloadSpeed
	}

	if g.touched["DownloadState"] && g.DownloadState == nil {
		data["download_state"] = nil
	} else if g.DownloadState != nil {
		data["download_state"] = g.DownloadState
	}

	if g.touched["Eta"] && g.Eta == nil {
		data["eta"] = nil
	} else if g.Eta != nil {
		data["eta"] = g.Eta
	}

	if g.touched["ExpiresAt"] && g.ExpiresAt == nil {
		data["expires_at"] = nil
	} else if g.ExpiresAt != nil {
		data["expires_at"] = g.ExpiresAt
	}

	if g.touched["Files"] && g.Files == nil {
		data["files"] = nil
	} else if g.Files != nil {
		data["files"] = g.Files
	}

	if g.touched["Hash"] && g.Hash == nil {
		data["hash"] = nil
	} else if g.Hash != nil {
		data["hash"] = g.Hash
	}

	if g.touched["Id"] && g.Id == nil {
		data["id"] = nil
	} else if g.Id != nil {
		data["id"] = g.Id
	}

	if g.touched["InactiveCheck"] && g.InactiveCheck == nil {
		data["inactive_check"] = nil
	} else if g.InactiveCheck != nil {
		data["inactive_check"] = g.InactiveCheck
	}

	if g.touched["Magnet"] && g.Magnet == nil {
		data["magnet"] = nil
	} else if g.Magnet != nil {
		data["magnet"] = g.Magnet
	}

	if g.touched["Name"] && g.Name == nil {
		data["name"] = nil
	} else if g.Name != nil {
		data["name"] = g.Name
	}

	if g.touched["Peers"] && g.Peers == nil {
		data["peers"] = nil
	} else if g.Peers != nil {
		data["peers"] = g.Peers
	}

	if g.touched["Progress"] && g.Progress == nil {
		data["progress"] = nil
	} else if g.Progress != nil {
		data["progress"] = g.Progress
	}

	if g.touched["Ratio"] && g.Ratio == nil {
		data["ratio"] = nil
	} else if g.Ratio != nil {
		data["ratio"] = g.Ratio
	}

	if g.touched["Seeds"] && g.Seeds == nil {
		data["seeds"] = nil
	} else if g.Seeds != nil {
		data["seeds"] = g.Seeds
	}

	if g.touched["Server"] && g.Server == nil {
		data["server"] = nil
	} else if g.Server != nil {
		data["server"] = g.Server
	}

	if g.touched["Size"] && g.Size == nil {
		data["size"] = nil
	} else if g.Size != nil {
		data["size"] = g.Size
	}

	if g.touched["TorrentFile"] && g.TorrentFile == nil {
		data["torrent_file"] = nil
	} else if g.TorrentFile != nil {
		data["torrent_file"] = g.TorrentFile
	}

	if g.touched["UpdatedAt"] && g.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if g.UpdatedAt != nil {
		data["updated_at"] = g.UpdatedAt
	}

	if g.touched["UploadSpeed"] && g.UploadSpeed == nil {
		data["upload_speed"] = nil
	} else if g.UploadSpeed != nil {
		data["upload_speed"] = g.UploadSpeed
	}

	return json.Marshal(data)
}

func (g GetTorrentListOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTorrentListOkResponseData to string"
	}
	return string(jsonData)
}

type DataFiles1 struct {
	Id        *float64 `json:"id,omitempty"`
	Md5       *string  `json:"md5,omitempty"`
	Mimetype  *string  `json:"mimetype,omitempty"`
	Name      *string  `json:"name,omitempty"`
	S3Path    *string  `json:"s3_path,omitempty"`
	ShortName *string  `json:"short_name,omitempty"`
	Size      *float64 `json:"size,omitempty"`
	touched   map[string]bool
}

func (d *DataFiles1) GetId() *float64 {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DataFiles1) SetId(id float64) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = &id
}

func (d *DataFiles1) SetIdNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = nil
}

func (d *DataFiles1) GetMd5() *string {
	if d == nil {
		return nil
	}
	return d.Md5
}

func (d *DataFiles1) SetMd5(md5 string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Md5"] = true
	d.Md5 = &md5
}

func (d *DataFiles1) SetMd5Nil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Md5"] = true
	d.Md5 = nil
}

func (d *DataFiles1) GetMimetype() *string {
	if d == nil {
		return nil
	}
	return d.Mimetype
}

func (d *DataFiles1) SetMimetype(mimetype string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Mimetype"] = true
	d.Mimetype = &mimetype
}

func (d *DataFiles1) SetMimetypeNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Mimetype"] = true
	d.Mimetype = nil
}

func (d *DataFiles1) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles1) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DataFiles1) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DataFiles1) GetS3Path() *string {
	if d == nil {
		return nil
	}
	return d.S3Path
}

func (d *DataFiles1) SetS3Path(s3Path string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["S3Path"] = true
	d.S3Path = &s3Path
}

func (d *DataFiles1) SetS3PathNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["S3Path"] = true
	d.S3Path = nil
}

func (d *DataFiles1) GetShortName() *string {
	if d == nil {
		return nil
	}
	return d.ShortName
}

func (d *DataFiles1) SetShortName(shortName string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["ShortName"] = true
	d.ShortName = &shortName
}

func (d *DataFiles1) SetShortNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["ShortName"] = true
	d.ShortName = nil
}

func (d *DataFiles1) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

func (d *DataFiles1) SetSize(size float64) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Size"] = true
	d.Size = &size
}

func (d *DataFiles1) SetSizeNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Size"] = true
	d.Size = nil
}

func (d DataFiles1) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Id"] && d.Id == nil {
		data["id"] = nil
	} else if d.Id != nil {
		data["id"] = d.Id
	}

	if d.touched["Md5"] && d.Md5 == nil {
		data["md5"] = nil
	} else if d.Md5 != nil {
		data["md5"] = d.Md5
	}

	if d.touched["Mimetype"] && d.Mimetype == nil {
		data["mimetype"] = nil
	} else if d.Mimetype != nil {
		data["mimetype"] = d.Mimetype
	}

	if d.touched["Name"] && d.Name == nil {
		data["name"] = nil
	} else if d.Name != nil {
		data["name"] = d.Name
	}

	if d.touched["S3Path"] && d.S3Path == nil {
		data["s3_path"] = nil
	} else if d.S3Path != nil {
		data["s3_path"] = d.S3Path
	}

	if d.touched["ShortName"] && d.ShortName == nil {
		data["short_name"] = nil
	} else if d.ShortName != nil {
		data["short_name"] = d.ShortName
	}

	if d.touched["Size"] && d.Size == nil {
		data["size"] = nil
	} else if d.Size != nil {
		data["size"] = d.Size
	}

	return json.Marshal(data)
}

func (d DataFiles1) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataFiles1 to string"
	}
	return string(jsonData)
}
