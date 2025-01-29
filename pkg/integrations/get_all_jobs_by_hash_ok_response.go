package integrations

import (
	"encoding/json"
)

type GetAllJobsByHashOkResponse struct {
	Data    []GetAllJobsByHashOkResponseData `json:"data,omitempty"`
	Detail  *string                          `json:"detail,omitempty"`
	Error   any                              `json:"error,omitempty"`
	Success *bool                            `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetAllJobsByHashOkResponse) GetData() []GetAllJobsByHashOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetAllJobsByHashOkResponse) SetData(data []GetAllJobsByHashOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = data
}

func (g *GetAllJobsByHashOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetAllJobsByHashOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsByHashOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetAllJobsByHashOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetAllJobsByHashOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetAllJobsByHashOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetAllJobsByHashOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetAllJobsByHashOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetAllJobsByHashOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetAllJobsByHashOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetAllJobsByHashOkResponse) MarshalJSON() ([]byte, error) {
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

func (g GetAllJobsByHashOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAllJobsByHashOkResponse to string"
	}
	return string(jsonData)
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
	touched     map[string]bool
}

func (g *GetAllJobsByHashOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetAllJobsByHashOkResponseData) SetAuthId(authId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = &authId
}

func (g *GetAllJobsByHashOkResponseData) SetAuthIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = nil
}

func (g *GetAllJobsByHashOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetAllJobsByHashOkResponseData) SetCreatedAt(createdAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = &createdAt
}

func (g *GetAllJobsByHashOkResponseData) SetCreatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = nil
}

func (g *GetAllJobsByHashOkResponseData) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetAllJobsByHashOkResponseData) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetAllJobsByHashOkResponseData) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetAllJobsByHashOkResponseData) GetDownloadUrl() *string {
	if g == nil {
		return nil
	}
	return g.DownloadUrl
}

func (g *GetAllJobsByHashOkResponseData) SetDownloadUrl(downloadUrl string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadUrl"] = true
	g.DownloadUrl = &downloadUrl
}

func (g *GetAllJobsByHashOkResponseData) SetDownloadUrlNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DownloadUrl"] = true
	g.DownloadUrl = nil
}

func (g *GetAllJobsByHashOkResponseData) GetFileId() *float64 {
	if g == nil {
		return nil
	}
	return g.FileId
}

func (g *GetAllJobsByHashOkResponseData) SetFileId(fileId float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["FileId"] = true
	g.FileId = &fileId
}

func (g *GetAllJobsByHashOkResponseData) SetFileIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["FileId"] = true
	g.FileId = nil
}

func (g *GetAllJobsByHashOkResponseData) GetHash() *string {
	if g == nil {
		return nil
	}
	return g.Hash
}

func (g *GetAllJobsByHashOkResponseData) SetHash(hash string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = &hash
}

func (g *GetAllJobsByHashOkResponseData) SetHashNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Hash"] = true
	g.Hash = nil
}

func (g *GetAllJobsByHashOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetAllJobsByHashOkResponseData) SetId(id float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GetAllJobsByHashOkResponseData) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GetAllJobsByHashOkResponseData) GetIntegration() *string {
	if g == nil {
		return nil
	}
	return g.Integration
}

func (g *GetAllJobsByHashOkResponseData) SetIntegration(integration string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Integration"] = true
	g.Integration = &integration
}

func (g *GetAllJobsByHashOkResponseData) SetIntegrationNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Integration"] = true
	g.Integration = nil
}

func (g *GetAllJobsByHashOkResponseData) GetProgress() *float64 {
	if g == nil {
		return nil
	}
	return g.Progress
}

func (g *GetAllJobsByHashOkResponseData) SetProgress(progress float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = &progress
}

func (g *GetAllJobsByHashOkResponseData) SetProgressNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Progress"] = true
	g.Progress = nil
}

func (g *GetAllJobsByHashOkResponseData) GetStatus() *string {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetAllJobsByHashOkResponseData) SetStatus(status string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = &status
}

func (g *GetAllJobsByHashOkResponseData) SetStatusNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = nil
}

func (g *GetAllJobsByHashOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetAllJobsByHashOkResponseData) SetType_(type_ string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Type_"] = true
	g.Type_ = &type_
}

func (g *GetAllJobsByHashOkResponseData) SetType_Nil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Type_"] = true
	g.Type_ = nil
}

func (g *GetAllJobsByHashOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetAllJobsByHashOkResponseData) SetUpdatedAt(updatedAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = &updatedAt
}

func (g *GetAllJobsByHashOkResponseData) SetUpdatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = nil
}

func (g *GetAllJobsByHashOkResponseData) GetZip() *bool {
	if g == nil {
		return nil
	}
	return g.Zip
}

func (g *GetAllJobsByHashOkResponseData) SetZip(zip bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Zip"] = true
	g.Zip = &zip
}

func (g *GetAllJobsByHashOkResponseData) SetZipNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Zip"] = true
	g.Zip = nil
}

func (g GetAllJobsByHashOkResponseData) MarshalJSON() ([]byte, error) {
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

func (g GetAllJobsByHashOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAllJobsByHashOkResponseData to string"
	}
	return string(jsonData)
}
