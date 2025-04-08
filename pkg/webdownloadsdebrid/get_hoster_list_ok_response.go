package webdownloadsdebrid

import (
	"encoding/json"
	"torbox-sdk-go/pkg/util"
)

type GetHosterListOkResponse struct {
	Data    []GetHosterListOkResponseData `json:"data,omitempty"`
	Detail  *string                       `json:"detail,omitempty"`
	Error   *util.Nullable[any]           `json:"error,omitempty"`
	Success *bool                         `json:"success,omitempty"`
}

func (g *GetHosterListOkResponse) GetData() []GetHosterListOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetHosterListOkResponse) SetData(data []GetHosterListOkResponseData) {
	g.Data = data
}

func (g *GetHosterListOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetHosterListOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetHosterListOkResponse) GetError() *util.Nullable[any] {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetHosterListOkResponse) SetError(error util.Nullable[any]) {
	g.Error = &error
}

func (g *GetHosterListOkResponse) SetErrorNull() {
	g.Error = &util.Nullable[any]{IsNull: true}
}

func (g *GetHosterListOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetHosterListOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g GetHosterListOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetHosterListOkResponse to string"
	}
	return string(jsonData)
}

type GetHosterListOkResponseData struct {
	DailyBandwidthLimit *float64               `json:"daily_bandwidth_limit,omitempty"`
	DailyBandwidthUsed  *float64               `json:"daily_bandwidth_used,omitempty"`
	DailyLinkLimit      *float64               `json:"daily_link_limit,omitempty"`
	DailyLinkUsed       *float64               `json:"daily_link_used,omitempty"`
	Domains             []string               `json:"domains,omitempty"`
	Domais              []string               `json:"domais,omitempty"`
	Domaisn             []string               `json:"domaisn,omitempty"`
	Icon                *string                `json:"icon,omitempty"`
	Limit               *float64               `json:"limit,omitempty"`
	Name                *string                `json:"name,omitempty"`
	Note                *util.Nullable[string] `json:"note,omitempty"`
	Status              *bool                  `json:"status,omitempty"`
	Type_               *string                `json:"type,omitempty"`
	Url                 *string                `json:"url,omitempty"`
}

func (g *GetHosterListOkResponseData) GetDailyBandwidthLimit() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyBandwidthLimit
}

func (g *GetHosterListOkResponseData) SetDailyBandwidthLimit(dailyBandwidthLimit float64) {
	g.DailyBandwidthLimit = &dailyBandwidthLimit
}

func (g *GetHosterListOkResponseData) GetDailyBandwidthUsed() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyBandwidthUsed
}

func (g *GetHosterListOkResponseData) SetDailyBandwidthUsed(dailyBandwidthUsed float64) {
	g.DailyBandwidthUsed = &dailyBandwidthUsed
}

func (g *GetHosterListOkResponseData) GetDailyLinkLimit() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyLinkLimit
}

func (g *GetHosterListOkResponseData) SetDailyLinkLimit(dailyLinkLimit float64) {
	g.DailyLinkLimit = &dailyLinkLimit
}

func (g *GetHosterListOkResponseData) GetDailyLinkUsed() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyLinkUsed
}

func (g *GetHosterListOkResponseData) SetDailyLinkUsed(dailyLinkUsed float64) {
	g.DailyLinkUsed = &dailyLinkUsed
}

func (g *GetHosterListOkResponseData) GetDomains() []string {
	if g == nil {
		return nil
	}
	return g.Domains
}

func (g *GetHosterListOkResponseData) SetDomains(domains []string) {
	g.Domains = domains
}

func (g *GetHosterListOkResponseData) GetDomais() []string {
	if g == nil {
		return nil
	}
	return g.Domais
}

func (g *GetHosterListOkResponseData) SetDomais(domais []string) {
	g.Domais = domais
}

func (g *GetHosterListOkResponseData) GetDomaisn() []string {
	if g == nil {
		return nil
	}
	return g.Domaisn
}

func (g *GetHosterListOkResponseData) SetDomaisn(domaisn []string) {
	g.Domaisn = domaisn
}

func (g *GetHosterListOkResponseData) GetIcon() *string {
	if g == nil {
		return nil
	}
	return g.Icon
}

func (g *GetHosterListOkResponseData) SetIcon(icon string) {
	g.Icon = &icon
}

func (g *GetHosterListOkResponseData) GetLimit() *float64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

func (g *GetHosterListOkResponseData) SetLimit(limit float64) {
	g.Limit = &limit
}

func (g *GetHosterListOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetHosterListOkResponseData) SetName(name string) {
	g.Name = &name
}

func (g *GetHosterListOkResponseData) GetNote() *util.Nullable[string] {
	if g == nil {
		return nil
	}
	return g.Note
}

func (g *GetHosterListOkResponseData) SetNote(note util.Nullable[string]) {
	g.Note = &note
}

func (g *GetHosterListOkResponseData) SetNoteNull() {
	g.Note = &util.Nullable[string]{IsNull: true}
}

func (g *GetHosterListOkResponseData) GetStatus() *bool {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetHosterListOkResponseData) SetStatus(status bool) {
	g.Status = &status
}

func (g *GetHosterListOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetHosterListOkResponseData) SetType_(type_ string) {
	g.Type_ = &type_
}

func (g *GetHosterListOkResponseData) GetUrl() *string {
	if g == nil {
		return nil
	}
	return g.Url
}

func (g *GetHosterListOkResponseData) SetUrl(url string) {
	g.Url = &url
}

func (g GetHosterListOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetHosterListOkResponseData to string"
	}
	return string(jsonData)
}
