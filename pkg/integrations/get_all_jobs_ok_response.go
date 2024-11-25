package integrations

type GetAllJobsOkResponse struct {
	Data    []GetAllJobsOkResponseData `json:"data,omitempty"`
	Detail  *string                    `json:"detail,omitempty"`
	Error   any                        `json:"error,omitempty"`
	Success *bool                      `json:"success,omitempty"`
}

func (g *GetAllJobsOkResponse) SetData(data []GetAllJobsOkResponseData) {
	g.Data = data
}

func (g *GetAllJobsOkResponse) GetData() []GetAllJobsOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetAllJobsOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetAllJobsOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsOkResponse) SetError(error any) {
	g.Error = error
}

func (g *GetAllJobsOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetAllJobsOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetAllJobsOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

type GetAllJobsOkResponseData struct {
	AuthId      *string  `json:"auth_id,omitempty"`
	CreatedAt   *string  `json:"created_at,omitempty"`
	Detail      *string  `json:"detail,omitempty"`
	DownloadUrl *string  `json:"download_url,omitempty"`
	FileId      *float64 `json:"file_id,omitempty"`
	Hash        *string  `json:"hash,omitempty"`
	Id          *float64 `json:"id,omitempty"`
	Integration *string  `json:"integration,omitempty"`
	Progress    *float64 `json:"progress,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Type_       *string  `json:"type,omitempty"`
	UpdatedAt   *string  `json:"updated_at,omitempty"`
	Zip         *bool    `json:"zip,omitempty"`
}

func (g *GetAllJobsOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetAllJobsOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetAllJobsOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetAllJobsOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetAllJobsOkResponseData) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetAllJobsOkResponseData) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsOkResponseData) SetDownloadUrl(downloadUrl string) {
	g.DownloadUrl = &downloadUrl
}

func (g *GetAllJobsOkResponseData) GetDownloadUrl() *string {
	if g == nil {
		return nil
	}
	return g.DownloadUrl
}

func (g *GetAllJobsOkResponseData) SetFileId(fileId float64) {
	g.FileId = &fileId
}

func (g *GetAllJobsOkResponseData) GetFileId() *float64 {
	if g == nil {
		return nil
	}
	return g.FileId
}

func (g *GetAllJobsOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetAllJobsOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetAllJobsOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetAllJobsOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetAllJobsOkResponseData) SetIntegration(integration string) {
	g.Integration = &integration
}

func (g *GetAllJobsOkResponseData) GetIntegration() *string {
	if g == nil {
		return nil
	}
	return g.Integration
}

func (g *GetAllJobsOkResponseData) SetProgress(progress float64) {
	g.Progress = &progress
}

func (g *GetAllJobsOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetAllJobsOkResponseData) SetStatus(status string) {
	g.Status = &status
}

func (g *GetAllJobsOkResponseData) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetAllJobsOkResponseData) SetType_(type_ string) {
	g.Type_ = &type_
}

func (g *GetAllJobsOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetAllJobsOkResponseData) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GetAllJobsOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetAllJobsOkResponseData) SetZip(zip bool) {
	g.Zip = &zip
}

func (g *GetAllJobsOkResponseData) GetZip() *bool {
	if g == nil {
		return nil
	}
	return g.Zip
}
