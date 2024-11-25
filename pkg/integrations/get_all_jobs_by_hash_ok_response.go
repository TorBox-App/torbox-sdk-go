package integrations

type GetAllJobsByHashOkResponse struct {
	Data    []GetAllJobsByHashOkResponseData `json:"data,omitempty"`
	Detail  *string                          `json:"detail,omitempty"`
	Error   any                              `json:"error,omitempty"`
	Success *bool                            `json:"success,omitempty"`
}

func (g *GetAllJobsByHashOkResponse) SetData(data []GetAllJobsByHashOkResponseData) {
	g.Data = data
}

func (g *GetAllJobsByHashOkResponse) GetData() []GetAllJobsByHashOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetAllJobsByHashOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetAllJobsByHashOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsByHashOkResponse) SetError(error any) {
	g.Error = error
}

func (g *GetAllJobsByHashOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetAllJobsByHashOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetAllJobsByHashOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

type GetAllJobsByHashOkResponseData struct {
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

func (g *GetAllJobsByHashOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetAllJobsByHashOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetAllJobsByHashOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetAllJobsByHashOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetAllJobsByHashOkResponseData) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetAllJobsByHashOkResponseData) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsByHashOkResponseData) SetDownloadUrl(downloadUrl string) {
	g.DownloadUrl = &downloadUrl
}

func (g *GetAllJobsByHashOkResponseData) GetDownloadUrl() *string {
	if g == nil {
		return nil
	}
	return g.DownloadUrl
}

func (g *GetAllJobsByHashOkResponseData) SetFileId(fileId float64) {
	g.FileId = &fileId
}

func (g *GetAllJobsByHashOkResponseData) GetFileId() *float64 {
	if g == nil {
		return nil
	}
	return g.FileId
}

func (g *GetAllJobsByHashOkResponseData) SetHash(hash string) {
	g.Hash = &hash
}

func (g *GetAllJobsByHashOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetAllJobsByHashOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetAllJobsByHashOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetAllJobsByHashOkResponseData) SetIntegration(integration string) {
	g.Integration = &integration
}

func (g *GetAllJobsByHashOkResponseData) GetIntegration() *string {
	if g == nil {
		return nil
	}
	return g.Integration
}

func (g *GetAllJobsByHashOkResponseData) SetProgress(progress float64) {
	g.Progress = &progress
}

func (g *GetAllJobsByHashOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetAllJobsByHashOkResponseData) SetStatus(status string) {
	g.Status = &status
}

func (g *GetAllJobsByHashOkResponseData) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetAllJobsByHashOkResponseData) SetType_(type_ string) {
	g.Type_ = &type_
}

func (g *GetAllJobsByHashOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetAllJobsByHashOkResponseData) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GetAllJobsByHashOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetAllJobsByHashOkResponseData) SetZip(zip bool) {
	g.Zip = &zip
}

func (g *GetAllJobsByHashOkResponseData) GetZip() *bool {
	if g == nil {
		return nil
	}
	return g.Zip
}
