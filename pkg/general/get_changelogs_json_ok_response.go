package general

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetChangelogsJsonOkResponse struct {
	Data    []GetChangelogsJsonOkResponseData `json:"data,omitempty"`
	Detail  *string                           `json:"detail,omitempty"`
	Error   *util.Nullable[any]               `json:"error,omitempty"`
	Success *bool                             `json:"success,omitempty"`
}

func (g *GetChangelogsJsonOkResponse) GetData() []GetChangelogsJsonOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetChangelogsJsonOkResponse) SetData(data []GetChangelogsJsonOkResponseData) {
	g.Data = data
}

func (g *GetChangelogsJsonOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetChangelogsJsonOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetChangelogsJsonOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetChangelogsJsonOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetChangelogsJsonOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetChangelogsJsonOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetChangelogsJsonOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetChangelogsJsonOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetChangelogsJsonOkResponse to string"
	}
	return string(jsonData)
}

type GetChangelogsJsonOkResponseData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Html      *string `json:"html,omitempty"`
	Id        *string `json:"id,omitempty"`
	Link      *string `json:"link,omitempty"`
	Markdown  *string `json:"markdown,omitempty"`
	Name      *string `json:"name,omitempty"`
}

func (g *GetChangelogsJsonOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetChangelogsJsonOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetChangelogsJsonOkResponseData) GetHtml() *string {
	if g == nil {
		return nil
	}
	return g.Html
}

func (g *GetChangelogsJsonOkResponseData) SetHtml(html string) {
	g.Html = &html
}

func (g *GetChangelogsJsonOkResponseData) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetChangelogsJsonOkResponseData) SetId(id string) {
	g.Id = &id
}

func (g *GetChangelogsJsonOkResponseData) GetLink() *string {
	if g == nil {
		return nil
	}
	return g.Link
}

func (g *GetChangelogsJsonOkResponseData) SetLink(link string) {
	g.Link = &link
}

func (g *GetChangelogsJsonOkResponseData) GetMarkdown() *string {
	if g == nil {
		return nil
	}
	return g.Markdown
}

func (g *GetChangelogsJsonOkResponseData) SetMarkdown(markdown string) {
	g.Markdown = &markdown
}

func (g *GetChangelogsJsonOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetChangelogsJsonOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g GetChangelogsJsonOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetChangelogsJsonOkResponseData to string"
	}
	return string(jsonData)
}
