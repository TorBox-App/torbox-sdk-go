package integrations

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetAllJobsOkResponse struct {
	Data    []GetAllJobsOkResponseData `json:"data,omitempty"`
	Detail  *string                    `json:"detail,omitempty"`
	Error   *util.Nullable[any]        `json:"error,omitempty"`
	Success *bool                      `json:"success,omitempty"`
}

func (g *GetAllJobsOkResponse) GetData() []GetAllJobsOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetAllJobsOkResponse) SetData(data []GetAllJobsOkResponseData) {
	g.Data = data
}

func (g *GetAllJobsOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetAllJobsOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetAllJobsOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetAllJobsOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetAllJobsOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetAllJobsOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetAllJobsOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAllJobsOkResponse to string"
	}
	return string(jsonData)
}

type GetAllJobsOkResponseData struct {
	AuthId      *string                `json:"auth_id,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
	Detail      *string                `json:"detail,omitempty"`
	DownloadUrl *util.Nullable[string] `json:"download_url,omitempty"`
	FileId      *float64               `json:"file_id,omitempty"`
	Hash        *string                `json:"hash,omitempty"`
	Id          *float64               `json:"id,omitempty"`
	Integration *string                `json:"integration,omitempty"`
	Progress    *float64               `json:"progress,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Type_       *string                `json:"type,omitempty"`
	UpdatedAt   *string                `json:"updated_at,omitempty"`
	Zip         *bool                  `json:"zip,omitempty"`
}

func (g *GetAllJobsOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetAllJobsOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetAllJobsOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetAllJobsOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetAllJobsOkResponseData) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsOkResponseData) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetAllJobsOkResponseData) GetDownloadUrl() *util.Nullable[string] {
	if g == nil {
		return nil
	}
	return g.DownloadUrl
}

func (g *GetAllJobsOkResponseData) SetDownloadUrl(downloadUrl util.Nullable[string]) {
	g.DownloadUrl = &downloadUrl
}

func (g *GetAllJobsOkResponseData) SetDownloadUrlNull() {
	g.DownloadUrl = &util.Nullable[string]{IsNull: true}
}

func (g *GetAllJobsOkResponseData) GetFileId() *float64 {
	if g == nil {
		return nil
	}
	return g.FileId
}

func (g *GetAllJobsOkResponseData) SetFileId(fileId float64) {
	g.FileId = &fileId
}

func (g *GetAllJobsOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetAllJobsOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetAllJobsOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetAllJobsOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetAllJobsOkResponseData) GetIntegration() *string {
	if g == nil {
		return nil
	}
	return g.Integration
}

func (g *GetAllJobsOkResponseData) SetIntegration(integration string) {
	g.Integration = &integration
}

func (g *GetAllJobsOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetAllJobsOkResponseData) SetProgress(progress float64) {
	g.Progress = &progress
}

func (g *GetAllJobsOkResponseData) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetAllJobsOkResponseData) SetStatus(status string) {
	g.Status = &status
}

func (g *GetAllJobsOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetAllJobsOkResponseData) SetType_(type_ string) {
	g.Type_ = &type_
}

func (g *GetAllJobsOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetAllJobsOkResponseData) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GetAllJobsOkResponseData) GetZip() *bool {
	if g == nil {
		return nil
	}
	return g.Zip
}

func (g *GetAllJobsOkResponseData) SetZip(zip bool) {
	g.Zip = &zip
}

func (g GetAllJobsOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAllJobsOkResponseData to string"
	}
	return string(jsonData)
}
