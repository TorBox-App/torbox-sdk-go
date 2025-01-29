package webdownloadsdebrid

import (
	"encoding/json"
)

type GetWebDownloadListOkResponse struct {
	Data    []GetWebDownloadListOkResponseData `json:"data,omitempty"`
	Detail  *string                            `json:"detail,omitempty"`
	Error   any                                `json:"error,omitempty"`
	Success *bool                              `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetWebDownloadListOkResponse) GetData() []GetWebDownloadListOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetWebDownloadListOkResponse) SetData(data []GetWebDownloadListOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = data
}

func (g *GetWebDownloadListOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetWebDownloadListOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetWebDownloadListOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetWebDownloadListOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetWebDownloadListOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetWebDownloadListOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetWebDownloadListOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetWebDownloadListOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetWebDownloadListOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetWebDownloadListOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetWebDownloadListOkResponse) MarshalJSON() ([]byte, error) {
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

func (g GetWebDownloadListOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetWebDownloadListOkResponse to string"
	}
	return string(jsonData)
}

type GetWebDownloadListOkResponseData struct {
	Active           *bool        `json:"active,omitempty"`
	AuthId           *string      `json:"auth_id,omitempty"`
	Availability     *float64     `json:"availability,omitempty"`
	CreatedAt        *string      `json:"created_at,omitempty"`
	DownloadFinished *bool        `json:"download_finished,omitempty"`
	DownloadPresent  *bool        `json:"download_present,omitempty"`
	DownloadSpeed    *float64     `json:"download_speed,omitempty"`
	DownloadState    *string      `json:"download_state,omitempty"`
	Error            *string      `json:"error,omitempty"`
	Eta              *float64     `json:"eta,omitempty"`
	ExpiresAt        *string      `json:"expires_at,omitempty"`
	Files            []DataFiles4 `json:"files,omitempty"`
	Hash             *string      `json:"hash,omitempty"`
	Id               *float64     `json:"id,omitempty"`
	InactiveCheck    *float64     `json:"inactive_check,omitempty"`
	Name             *string      `json:"name,omitempty"`
	Progress         *float64     `json:"progress,omitempty"`
	Server           *float64     `json:"server,omitempty"`
	Size             *float64     `json:"size,omitempty"`
	TorrentFile      *bool        `json:"torrent_file,omitempty"`
	UpdatedAt        *string      `json:"updated_at,omitempty"`
	UploadSpeed      *float64     `json:"upload_speed,omitempty"`
	touched          map[string]bool
}

func (g *GetWebDownloadListOkResponseData) GetActive() *bool {
	if g == nil {
		return nil
	}
	return g.Active
}

func (g *GetWebDownloadListOkResponseData) SetActive(active bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Active"] = true
	g.Active = &active
}

func (g *GetWebDownloadListOkResponseData) SetActiveNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Active"] = true
	g.Active = nil
}

func (g *GetWebDownloadListOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetWebDownloadListOkResponseData) SetAuthId(authId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = &authId
}

func (g *GetWebDownloadListOkResponseData) SetAuthIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = nil
}

func (g *GetWebDownloadListOkResponseData) GetAvailability() *float64 {
	if g == nil {
		return nil
	}
	return g.Availability
}

func (g *GetWebDownloadListOkResponseData) SetAvailability(availability float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Availability"] = true
	g.Availability = &availability
}

func (g *GetWebDownloadListOkResponseData) SetAvailabilityNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Availability"] = true
	g.Availability = nil
}

func (g *GetWebDownloadListOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetWebDownloadListOkResponseData) SetCreatedAt(createdAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = &createdAt
}

func (g *GetWebDownloadListOkResponseData) SetCreatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = nil
}

func (g *GetWebDownloadListOkResponseData) GetDownloadFinished() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadFinished
}

func (g *GetWebDownloadListOkResponseData) SetDownloadFinished(downloadFinished bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadFinished"] = true
	g.DownloadFinished = &downloadFinished
}

func (g *GetWebDownloadListOkResponseData) SetDownloadFinishedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadFinished"] = true
	g.DownloadFinished = nil
}

func (g *GetWebDownloadListOkResponseData) GetDownloadPresent() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadPresent
}

func (g *GetWebDownloadListOkResponseData) SetDownloadPresent(downloadPresent bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadPresent"] = true
	g.DownloadPresent = &downloadPresent
}

func (g *GetWebDownloadListOkResponseData) SetDownloadPresentNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadPresent"] = true
	g.DownloadPresent = nil
}

func (g *GetWebDownloadListOkResponseData) GetDownloadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.DownloadSpeed
}

func (g *GetWebDownloadListOkResponseData) SetDownloadSpeed(downloadSpeed float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadSpeed"] = true
	g.DownloadSpeed = &downloadSpeed
}

func (g *GetWebDownloadListOkResponseData) SetDownloadSpeedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadSpeed"] = true
	g.DownloadSpeed = nil
}

func (g *GetWebDownloadListOkResponseData) GetDownloadState() *string {
	if g == nil {
		return nil
	}
	return g.DownloadState
}

func (g *GetWebDownloadListOkResponseData) SetDownloadState(downloadState string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadState"] = true
	g.DownloadState = &downloadState
}

func (g *GetWebDownloadListOkResponseData) SetDownloadStateNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadState"] = true
	g.DownloadState = nil
}

func (g *GetWebDownloadListOkResponseData) GetError() *string {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetWebDownloadListOkResponseData) SetError(error string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = &error
}

func (g *GetWebDownloadListOkResponseData) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetWebDownloadListOkResponseData) GetEta() *float64 {
	if g == nil {
		return nil
	}
	return g.Eta
}

func (g *GetWebDownloadListOkResponseData) SetEta(eta float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Eta"] = true
	g.Eta = &eta
}

func (g *GetWebDownloadListOkResponseData) SetEtaNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Eta"] = true
	g.Eta = nil
}

func (g *GetWebDownloadListOkResponseData) GetExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.ExpiresAt
}

func (g *GetWebDownloadListOkResponseData) SetExpiresAt(expiresAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ExpiresAt"] = true
	g.ExpiresAt = &expiresAt
}

func (g *GetWebDownloadListOkResponseData) SetExpiresAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["ExpiresAt"] = true
	g.ExpiresAt = nil
}

func (g *GetWebDownloadListOkResponseData) GetFiles() []DataFiles4 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetWebDownloadListOkResponseData) SetFiles(files []DataFiles4) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Files"] = true
	g.Files = files
}

func (g *GetWebDownloadListOkResponseData) SetFilesNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Files"] = true
	g.Files = nil
}

func (g *GetWebDownloadListOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetWebDownloadListOkResponseData) SetHash(hash string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = &hash
}

func (g *GetWebDownloadListOkResponseData) SetHashNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = nil
}

func (g *GetWebDownloadListOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetWebDownloadListOkResponseData) SetId(id float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GetWebDownloadListOkResponseData) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GetWebDownloadListOkResponseData) GetInactiveCheck() *float64 {
	if g == nil {
		return nil
	}
	return g.InactiveCheck
}

func (g *GetWebDownloadListOkResponseData) SetInactiveCheck(inactiveCheck float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["InactiveCheck"] = true
	g.InactiveCheck = &inactiveCheck
}

func (g *GetWebDownloadListOkResponseData) SetInactiveCheckNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["InactiveCheck"] = true
	g.InactiveCheck = nil
}

func (g *GetWebDownloadListOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetWebDownloadListOkResponseData) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GetWebDownloadListOkResponseData) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}

func (g *GetWebDownloadListOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetWebDownloadListOkResponseData) SetProgress(progress float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = &progress
}

func (g *GetWebDownloadListOkResponseData) SetProgressNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = nil
}

func (g *GetWebDownloadListOkResponseData) GetServer() *float64 {
	if g == nil {
		return nil
	}
	return g.Server
}

func (g *GetWebDownloadListOkResponseData) SetServer(server float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Server"] = true
	g.Server = &server
}

func (g *GetWebDownloadListOkResponseData) SetServerNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Server"] = true
	g.Server = nil
}

func (g *GetWebDownloadListOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetWebDownloadListOkResponseData) SetSize(size float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = &size
}

func (g *GetWebDownloadListOkResponseData) SetSizeNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Size"] = true
	g.Size = nil
}

func (g *GetWebDownloadListOkResponseData) GetTorrentFile() *bool {
	if g == nil {
		return nil
	}
	return g.TorrentFile
}

func (g *GetWebDownloadListOkResponseData) SetTorrentFile(torrentFile bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TorrentFile"] = true
	g.TorrentFile = &torrentFile
}

func (g *GetWebDownloadListOkResponseData) SetTorrentFileNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TorrentFile"] = true
	g.TorrentFile = nil
}

func (g *GetWebDownloadListOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetWebDownloadListOkResponseData) SetUpdatedAt(updatedAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = &updatedAt
}

func (g *GetWebDownloadListOkResponseData) SetUpdatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = nil
}

func (g *GetWebDownloadListOkResponseData) GetUploadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.UploadSpeed
}

func (g *GetWebDownloadListOkResponseData) SetUploadSpeed(uploadSpeed float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UploadSpeed"] = true
	g.UploadSpeed = &uploadSpeed
}

func (g *GetWebDownloadListOkResponseData) SetUploadSpeedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UploadSpeed"] = true
	g.UploadSpeed = nil
}

func (g GetWebDownloadListOkResponseData) MarshalJSON() ([]byte, error) {
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

	if g.touched["Error"] && g.Error == nil {
		data["error"] = nil
	} else if g.Error != nil {
		data["error"] = g.Error
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

	if g.touched["Name"] && g.Name == nil {
		data["name"] = nil
	} else if g.Name != nil {
		data["name"] = g.Name
	}

	if g.touched["Progress"] && g.Progress == nil {
		data["progress"] = nil
	} else if g.Progress != nil {
		data["progress"] = g.Progress
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

func (g GetWebDownloadListOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetWebDownloadListOkResponseData to string"
	}
	return string(jsonData)
}

type DataFiles4 struct {
	Id        *float64 `json:"id,omitempty"`
	Md5       *string  `json:"md5,omitempty"`
	Mimetype  *string  `json:"mimetype,omitempty"`
	Name      *string  `json:"name,omitempty"`
	S3Path    *string  `json:"s3_path,omitempty"`
	ShortName *string  `json:"short_name,omitempty"`
	Size      *float64 `json:"size,omitempty"`
	touched   map[string]bool
}

func (d *DataFiles4) GetId() *float64 {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DataFiles4) SetId(id float64) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = &id
}

func (d *DataFiles4) SetIdNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = nil
}

func (d *DataFiles4) GetMd5() *string {
	if d == nil {
		return nil
	}
	return d.Md5
}

func (d *DataFiles4) SetMd5(md5 string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Md5"] = true
	d.Md5 = &md5
}

func (d *DataFiles4) SetMd5Nil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Md5"] = true
	d.Md5 = nil
}

func (d *DataFiles4) GetMimetype() *string {
	if d == nil {
		return nil
	}
	return d.Mimetype
}

func (d *DataFiles4) SetMimetype(mimetype string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Mimetype"] = true
	d.Mimetype = &mimetype
}

func (d *DataFiles4) SetMimetypeNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Mimetype"] = true
	d.Mimetype = nil
}

func (d *DataFiles4) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles4) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DataFiles4) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DataFiles4) GetS3Path() *string {
	if d == nil {
		return nil
	}
	return d.S3Path
}

func (d *DataFiles4) SetS3Path(s3Path string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["S3Path"] = true
	d.S3Path = &s3Path
}

func (d *DataFiles4) SetS3PathNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["S3Path"] = true
	d.S3Path = nil
}

func (d *DataFiles4) GetShortName() *string {
	if d == nil {
		return nil
	}
	return d.ShortName
}

func (d *DataFiles4) SetShortName(shortName string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["ShortName"] = true
	d.ShortName = &shortName
}

func (d *DataFiles4) SetShortNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["ShortName"] = true
	d.ShortName = nil
}

func (d *DataFiles4) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

func (d *DataFiles4) SetSize(size float64) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Size"] = true
	d.Size = &size
}

func (d *DataFiles4) SetSizeNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Size"] = true
	d.Size = nil
}

func (d DataFiles4) MarshalJSON() ([]byte, error) {
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

func (d DataFiles4) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataFiles4 to string"
	}
	return string(jsonData)
}
