package integrations

import (
	"encoding/json"
)

type GetAllJobsOkResponse struct {
	Data    []GetAllJobsOkResponseData `json:"data,omitempty"`
	Detail  *string                    `json:"detail,omitempty"`
	Error   any                        `json:"error,omitempty"`
	Success *bool                      `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetAllJobsOkResponse) GetData() []GetAllJobsOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetAllJobsOkResponse) SetData(data []GetAllJobsOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = data
}

func (g *GetAllJobsOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetAllJobsOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetAllJobsOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetAllJobsOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetAllJobsOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetAllJobsOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetAllJobsOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetAllJobsOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetAllJobsOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetAllJobsOkResponse) MarshalJSON() ([]byte, error) {
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

func (g GetAllJobsOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAllJobsOkResponse to string"
	}
	return string(jsonData)
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
	touched     map[string]bool
}

func (g *GetAllJobsOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetAllJobsOkResponseData) SetAuthId(authId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = &authId
}

func (g *GetAllJobsOkResponseData) SetAuthIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = nil
}

func (g *GetAllJobsOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetAllJobsOkResponseData) SetCreatedAt(createdAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = &createdAt
}

func (g *GetAllJobsOkResponseData) SetCreatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = nil
}

func (g *GetAllJobsOkResponseData) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsOkResponseData) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetAllJobsOkResponseData) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetAllJobsOkResponseData) GetDownloadUrl() *string {
	if g == nil {
		return nil
	}
	return g.DownloadUrl
}

func (g *GetAllJobsOkResponseData) SetDownloadUrl(downloadUrl string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadUrl"] = true
	g.DownloadUrl = &downloadUrl
}

func (g *GetAllJobsOkResponseData) SetDownloadUrlNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadUrl"] = true
	g.DownloadUrl = nil
}

func (g *GetAllJobsOkResponseData) GetFileId() *float64 {
	if g == nil {
		return nil
	}
	return g.FileId
}

func (g *GetAllJobsOkResponseData) SetFileId(fileId float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["FileId"] = true
	g.FileId = &fileId
}

func (g *GetAllJobsOkResponseData) SetFileIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["FileId"] = true
	g.FileId = nil
}

func (g *GetAllJobsOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetAllJobsOkResponseData) SetHash(hash string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = &hash
}

func (g *GetAllJobsOkResponseData) SetHashNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = nil
}

func (g *GetAllJobsOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetAllJobsOkResponseData) SetId(id float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GetAllJobsOkResponseData) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GetAllJobsOkResponseData) GetIntegration() *string {
	if g == nil {
		return nil
	}
	return g.Integration
}

func (g *GetAllJobsOkResponseData) SetIntegration(integration string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Integration"] = true
	g.Integration = &integration
}

func (g *GetAllJobsOkResponseData) SetIntegrationNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Integration"] = true
	g.Integration = nil
}

func (g *GetAllJobsOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetAllJobsOkResponseData) SetProgress(progress float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = &progress
}

func (g *GetAllJobsOkResponseData) SetProgressNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = nil
}

func (g *GetAllJobsOkResponseData) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetAllJobsOkResponseData) SetStatus(status string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = &status
}

func (g *GetAllJobsOkResponseData) SetStatusNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = nil
}

func (g *GetAllJobsOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetAllJobsOkResponseData) SetType_(type_ string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Type_"] = true
	g.Type_ = &type_
}

func (g *GetAllJobsOkResponseData) SetType_Nil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Type_"] = true
	g.Type_ = nil
}

func (g *GetAllJobsOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetAllJobsOkResponseData) SetUpdatedAt(updatedAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = &updatedAt
}

func (g *GetAllJobsOkResponseData) SetUpdatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = nil
}

func (g *GetAllJobsOkResponseData) GetZip() *bool {
	if g == nil {
		return nil
	}
	return g.Zip
}

func (g *GetAllJobsOkResponseData) SetZip(zip bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Zip"] = true
	g.Zip = &zip
}

func (g *GetAllJobsOkResponseData) SetZipNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Zip"] = true
	g.Zip = nil
}

func (g GetAllJobsOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["AuthId"] && g.AuthId == nil {
		data["auth_id"] = nil
	} else if g.AuthId != nil {
		data["auth_id"] = g.AuthId
	}

	if g.touched["CreatedAt"] && g.CreatedAt == nil {
		data["created_at"] = nil
	} else if g.CreatedAt != nil {
		data["created_at"] = g.CreatedAt
	}

	if g.touched["Detail"] && g.Detail == nil {
		data["detail"] = nil
	} else if g.Detail != nil {
		data["detail"] = g.Detail
	}

	if g.touched["DownloadUrl"] && g.DownloadUrl == nil {
		data["download_url"] = nil
	} else if g.DownloadUrl != nil {
		data["download_url"] = g.DownloadUrl
	}

	if g.touched["FileId"] && g.FileId == nil {
		data["file_id"] = nil
	} else if g.FileId != nil {
		data["file_id"] = g.FileId
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

	if g.touched["Integration"] && g.Integration == nil {
		data["integration"] = nil
	} else if g.Integration != nil {
		data["integration"] = g.Integration
	}

	if g.touched["Progress"] && g.Progress == nil {
		data["progress"] = nil
	} else if g.Progress != nil {
		data["progress"] = g.Progress
	}

	if g.touched["Status"] && g.Status == nil {
		data["status"] = nil
	} else if g.Status != nil {
		data["status"] = g.Status
	}

	if g.touched["Type_"] && g.Type_ == nil {
		data["type"] = nil
	} else if g.Type_ != nil {
		data["type"] = g.Type_
	}

	if g.touched["UpdatedAt"] && g.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if g.UpdatedAt != nil {
		data["updated_at"] = g.UpdatedAt
	}

	if g.touched["Zip"] && g.Zip == nil {
		data["zip"] = nil
	} else if g.Zip != nil {
		data["zip"] = g.Zip
	}

	return json.Marshal(data)
}

func (g GetAllJobsOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAllJobsOkResponseData to string"
	}
	return string(jsonData)
}
