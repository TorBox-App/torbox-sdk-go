package webdownloadsdebrid

import (
	"encoding/json"
)

type GetHosterListOkResponse struct {
	Data    []GetHosterListOkResponseData `json:"data,omitempty"`
	Detail  *string                       `json:"detail,omitempty"`
	Error   any                           `json:"error,omitempty"`
	Success *bool                         `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetHosterListOkResponse) GetData() []GetHosterListOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetHosterListOkResponse) SetData(data []GetHosterListOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = data
}

func (g *GetHosterListOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetHosterListOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetHosterListOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetHosterListOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetHosterListOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetHosterListOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetHosterListOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetHosterListOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetHosterListOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetHosterListOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetHosterListOkResponse) MarshalJSON() ([]byte, error) {
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

func (g GetHosterListOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetHosterListOkResponse to string"
	}
	return string(jsonData)
}

type GetHosterListOkResponseData struct {
	DailyBandwidthLimit *float64 `json:"daily_bandwidth_limit,omitempty"`
	DailyBandwidthUsed  *float64 `json:"daily_bandwidth_used,omitempty"`
	DailyLinkLimit      *float64 `json:"daily_link_limit,omitempty"`
	DailyLinkUsed       *float64 `json:"daily_link_used,omitempty"`
	Domains             []string `json:"domains,omitempty"`
	Domais              []string `json:"domais,omitempty"`
	Domaisn             []string `json:"domaisn,omitempty"`
	Icon                *string  `json:"icon,omitempty"`
	Limit               *float64 `json:"limit,omitempty"`
	Name                *string  `json:"name,omitempty"`
	Note                *string  `json:"note,omitempty"`
	Status              *bool    `json:"status,omitempty"`
	Type_               *string  `json:"type,omitempty"`
	Url                 *string  `json:"url,omitempty"`
	touched             map[string]bool
}

func (g *GetHosterListOkResponseData) GetDailyBandwidthLimit() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyBandwidthLimit
}

func (g *GetHosterListOkResponseData) SetDailyBandwidthLimit(dailyBandwidthLimit float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyBandwidthLimit"] = true
	g.DailyBandwidthLimit = &dailyBandwidthLimit
}

func (g *GetHosterListOkResponseData) SetDailyBandwidthLimitNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyBandwidthLimit"] = true
	g.DailyBandwidthLimit = nil
}

func (g *GetHosterListOkResponseData) GetDailyBandwidthUsed() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyBandwidthUsed
}

func (g *GetHosterListOkResponseData) SetDailyBandwidthUsed(dailyBandwidthUsed float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyBandwidthUsed"] = true
	g.DailyBandwidthUsed = &dailyBandwidthUsed
}

func (g *GetHosterListOkResponseData) SetDailyBandwidthUsedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyBandwidthUsed"] = true
	g.DailyBandwidthUsed = nil
}

func (g *GetHosterListOkResponseData) GetDailyLinkLimit() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyLinkLimit
}

func (g *GetHosterListOkResponseData) SetDailyLinkLimit(dailyLinkLimit float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyLinkLimit"] = true
	g.DailyLinkLimit = &dailyLinkLimit
}

func (g *GetHosterListOkResponseData) SetDailyLinkLimitNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyLinkLimit"] = true
	g.DailyLinkLimit = nil
}

func (g *GetHosterListOkResponseData) GetDailyLinkUsed() *float64 {
	if g == nil {
		return nil
	}
	return g.DailyLinkUsed
}

func (g *GetHosterListOkResponseData) SetDailyLinkUsed(dailyLinkUsed float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyLinkUsed"] = true
	g.DailyLinkUsed = &dailyLinkUsed
}

func (g *GetHosterListOkResponseData) SetDailyLinkUsedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["DailyLinkUsed"] = true
	g.DailyLinkUsed = nil
}

func (g *GetHosterListOkResponseData) GetDomains() []string {
	if g == nil {
		return nil
	}
	return g.Domains
}

func (g *GetHosterListOkResponseData) SetDomains(domains []string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Domains"] = true
	g.Domains = domains
}

func (g *GetHosterListOkResponseData) SetDomainsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Domains"] = true
	g.Domains = nil
}

func (g *GetHosterListOkResponseData) GetDomais() []string {
	if g == nil {
		return nil
	}
	return g.Domais
}

func (g *GetHosterListOkResponseData) SetDomais(domais []string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Domais"] = true
	g.Domais = domais
}

func (g *GetHosterListOkResponseData) SetDomaisNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Domais"] = true
	g.Domais = nil
}

func (g *GetHosterListOkResponseData) GetDomaisn() []string {
	if g == nil {
		return nil
	}
	return g.Domaisn
}

func (g *GetHosterListOkResponseData) SetDomaisn(domaisn []string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Domaisn"] = true
	g.Domaisn = domaisn
}

func (g *GetHosterListOkResponseData) SetDomaisnNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Domaisn"] = true
	g.Domaisn = nil
}

func (g *GetHosterListOkResponseData) GetIcon() *string {
	if g == nil {
		return nil
	}
	return g.Icon
}

func (g *GetHosterListOkResponseData) SetIcon(icon string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Icon"] = true
	g.Icon = &icon
}

func (g *GetHosterListOkResponseData) SetIconNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Icon"] = true
	g.Icon = nil
}

func (g *GetHosterListOkResponseData) GetLimit() *float64 {
	if g == nil {
		return nil
	}
	return g.Limit
}

func (g *GetHosterListOkResponseData) SetLimit(limit float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Limit"] = true
	g.Limit = &limit
}

func (g *GetHosterListOkResponseData) SetLimitNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Limit"] = true
	g.Limit = nil
}

func (g *GetHosterListOkResponseData) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GetHosterListOkResponseData) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GetHosterListOkResponseData) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}

func (g *GetHosterListOkResponseData) GetNote() *string {
	if g == nil {
		return nil
	}
	return g.Note
}

func (g *GetHosterListOkResponseData) SetNote(note string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Note"] = true
	g.Note = &note
}

func (g *GetHosterListOkResponseData) SetNoteNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Note"] = true
	g.Note = nil
}

func (g *GetHosterListOkResponseData) GetStatus() *bool {
	if g == nil {
		return nil
	}
	return g.Status
}

func (g *GetHosterListOkResponseData) SetStatus(status bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = &status
}

func (g *GetHosterListOkResponseData) SetStatusNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Status"] = true
	g.Status = nil
}

func (g *GetHosterListOkResponseData) GetType_() *string {
	if g == nil {
		return nil
	}
	return g.Type_
}

func (g *GetHosterListOkResponseData) SetType_(type_ string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Type_"] = true
	g.Type_ = &type_
}

func (g *GetHosterListOkResponseData) SetType_Nil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Type_"] = true
	g.Type_ = nil
}

func (g *GetHosterListOkResponseData) GetUrl() *string {
	if g == nil {
		return nil
	}
	return g.Url
}

func (g *GetHosterListOkResponseData) SetUrl(url string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Url"] = true
	g.Url = &url
}

func (g *GetHosterListOkResponseData) SetUrlNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Url"] = true
	g.Url = nil
}

func (g GetHosterListOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["DailyBandwidthLimit"] && g.DailyBandwidthLimit == nil {
		data["daily_bandwidth_limit"] = nil
	} else if g.DailyBandwidthLimit != nil {
		data["daily_bandwidth_limit"] = g.DailyBandwidthLimit
	}

	if g.touched["DailyBandwidthUsed"] && g.DailyBandwidthUsed == nil {
		data["daily_bandwidth_used"] = nil
	} else if g.DailyBandwidthUsed != nil {
		data["daily_bandwidth_used"] = g.DailyBandwidthUsed
	}

	if g.touched["DailyLinkLimit"] && g.DailyLinkLimit == nil {
		data["daily_link_limit"] = nil
	} else if g.DailyLinkLimit != nil {
		data["daily_link_limit"] = g.DailyLinkLimit
	}

	if g.touched["DailyLinkUsed"] && g.DailyLinkUsed == nil {
		data["daily_link_used"] = nil
	} else if g.DailyLinkUsed != nil {
		data["daily_link_used"] = g.DailyLinkUsed
	}

	if g.touched["Domains"] && g.Domains == nil {
		data["domains"] = nil
	} else if g.Domains != nil {
		data["domains"] = g.Domains
	}

	if g.touched["Domais"] && g.Domais == nil {
		data["domais"] = nil
	} else if g.Domais != nil {
		data["domais"] = g.Domais
	}

	if g.touched["Domaisn"] && g.Domaisn == nil {
		data["domaisn"] = nil
	} else if g.Domaisn != nil {
		data["domaisn"] = g.Domaisn
	}

	if g.touched["Icon"] && g.Icon == nil {
		data["icon"] = nil
	} else if g.Icon != nil {
		data["icon"] = g.Icon
	}

	if g.touched["Limit"] && g.Limit == nil {
		data["limit"] = nil
	} else if g.Limit != nil {
		data["limit"] = g.Limit
	}

	if g.touched["Name"] && g.Name == nil {
		data["name"] = nil
	} else if g.Name != nil {
		data["name"] = g.Name
	}

	if g.touched["Note"] && g.Note == nil {
		data["note"] = nil
	} else if g.Note != nil {
		data["note"] = g.Note
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

	if g.touched["Url"] && g.Url == nil {
		data["url"] = nil
	} else if g.Url != nil {
		data["url"] = g.Url
	}

	return json.Marshal(data)
}

func (g GetHosterListOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetHosterListOkResponseData to string"
	}
	return string(jsonData)
}
