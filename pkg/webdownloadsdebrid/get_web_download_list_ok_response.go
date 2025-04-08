package webdownloadsdebrid

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetWebDownloadListOkResponse struct {
	Data    []GetWebDownloadListOkResponseData `json:"data,omitempty"`
	Detail  *string                            `json:"detail,omitempty"`
	Error   *util.Nullable[any]                `json:"error,omitempty"`
	Success *bool                              `json:"success,omitempty"`
}

func (g *GetWebDownloadListOkResponse) GetData() []GetWebDownloadListOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetWebDownloadListOkResponse) SetData(data []GetWebDownloadListOkResponseData) {
	g.Data = data
}

func (g *GetWebDownloadListOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetWebDownloadListOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetWebDownloadListOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetWebDownloadListOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetWebDownloadListOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetWebDownloadListOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetWebDownloadListOkResponse) SetSuccess(success bool) {
	g.Success = &success
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
}

func (g *GetWebDownloadListOkResponseData) GetActive() *bool {
	if g == nil {
		return nil
	}
	return g.Active
}

func (g *GetWebDownloadListOkResponseData) SetActive(active bool) {
	g.Active = &active
}

func (g *GetWebDownloadListOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetWebDownloadListOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetWebDownloadListOkResponseData) GetAvailability() *float64 {
	if g == nil {
		return nil
	}
	return g.Availability
}

func (g *GetWebDownloadListOkResponseData) SetAvailability(availability float64) {
	g.Availability = &availability
}

func (g *GetWebDownloadListOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetWebDownloadListOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetWebDownloadListOkResponseData) GetDownloadFinished() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadFinished
}

func (g *GetWebDownloadListOkResponseData) SetDownloadFinished(downloadFinished bool) {
	g.DownloadFinished = &downloadFinished
}

func (g *GetWebDownloadListOkResponseData) GetDownloadPresent() *bool {
	if g == nil {
		return nil
	}
	return g.DownloadPresent
}

func (g *GetWebDownloadListOkResponseData) SetDownloadPresent(downloadPresent bool) {
	g.DownloadPresent = &downloadPresent
}

func (g *GetWebDownloadListOkResponseData) GetDownloadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.DownloadSpeed
}

func (g *GetWebDownloadListOkResponseData) SetDownloadSpeed(downloadSpeed float64) {
	g.DownloadSpeed = &downloadSpeed
}

func (g *GetWebDownloadListOkResponseData) GetDownloadState() *string {
	if g == nil {
		return nil
	}
	return g.DownloadState
}

func (g *GetWebDownloadListOkResponseData) SetDownloadState(downloadState string) {
	g.DownloadState = &downloadState
}

func (g *GetWebDownloadListOkResponseData) GetError() *string {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetWebDownloadListOkResponseData) SetError(error string) {
	g.Error = &error
}

func (g *GetWebDownloadListOkResponseData) GetEta() *float64 {
	if g == nil {
		return nil
	}
	return g.Eta
}

func (g *GetWebDownloadListOkResponseData) SetEta(eta float64) {
	g.Eta = &eta
}

func (g *GetWebDownloadListOkResponseData) GetExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.ExpiresAt
}

func (g *GetWebDownloadListOkResponseData) SetExpiresAt(expiresAt string) {
	g.ExpiresAt = &expiresAt
}

func (g *GetWebDownloadListOkResponseData) GetFiles() []DataFiles4 {
	if g == nil {
		return nil
	}
	return g.Files
}

func (g *GetWebDownloadListOkResponseData) SetFiles(files []DataFiles4) {
	g.Files = files
}

func (g *GetWebDownloadListOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetWebDownloadListOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetWebDownloadListOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetWebDownloadListOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetWebDownloadListOkResponseData) GetInactiveCheck() *float64 {
	if g == nil {
		return nil
	}
	return g.InactiveCheck
}

func (g *GetWebDownloadListOkResponseData) SetInactiveCheck(inactiveCheck float64) {
	g.InactiveCheck = &inactiveCheck
}

func (g *GetWebDownloadListOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetWebDownloadListOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetWebDownloadListOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetWebDownloadListOkResponseData) SetProgress(progress float64) {
	g.Progress = &progress
}

func (g *GetWebDownloadListOkResponseData) GetServer() *float64 {
	if g == nil {
		return nil
	}
	return g.Server
}

func (g *GetWebDownloadListOkResponseData) SetServer(server float64) {
	g.Server = &server
}

func (g *GetWebDownloadListOkResponseData) GetSize() *float64 {
	if g == nil {
		return nil
	}
	return g.Size
}

func (g *GetWebDownloadListOkResponseData) SetSize(size float64) {
	g.Size = &size
}

func (g *GetWebDownloadListOkResponseData) GetTorrentFile() *bool {
	if g == nil {
		return nil
	}
	return g.TorrentFile
}

func (g *GetWebDownloadListOkResponseData) SetTorrentFile(torrentFile bool) {
	g.TorrentFile = &torrentFile
}

func (g *GetWebDownloadListOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetWebDownloadListOkResponseData) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GetWebDownloadListOkResponseData) GetUploadSpeed() *float64 {
	if g == nil {
		return nil
	}
	return g.UploadSpeed
}

func (g *GetWebDownloadListOkResponseData) SetUploadSpeed(uploadSpeed float64) {
	g.UploadSpeed = &uploadSpeed
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
}

func (d *DataFiles4) GetId() *float64 {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DataFiles4) SetId(id float64) {
	d.Id = &id
}

func (d *DataFiles4) GetMd5() *string {
	if d == nil {
		return nil
	}
	return d.Md5
}

func (d *DataFiles4) SetMd5(md5 string) {
	d.Md5 = &md5
}

func (d *DataFiles4) GetMimetype() *string {
	if d == nil {
		return nil
	}
	return d.Mimetype
}

func (d *DataFiles4) SetMimetype(mimetype string) {
	d.Mimetype = &mimetype
}

func (d *DataFiles4) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DataFiles4) SetName(name string) {
	d.Name = &name
}

func (d *DataFiles4) GetS3Path() *string {
	if d == nil {
		return nil
	}
	return d.S3Path
}

func (d *DataFiles4) SetS3Path(s3Path string) {
	d.S3Path = &s3Path
}

func (d *DataFiles4) GetShortName() *string {
	if d == nil {
		return nil
	}
	return d.ShortName
}

func (d *DataFiles4) SetShortName(shortName string) {
	d.ShortName = &shortName
}

func (d *DataFiles4) GetSize() *float64 {
	if d == nil {
		return nil
	}
	return d.Size
}

func (d *DataFiles4) SetSize(size float64) {
	d.Size = &size
}

func (d DataFiles4) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataFiles4 to string"
	}
	return string(jsonData)
}
