package user

import (
	"encoding/json"
)

type GetUserDataOkResponse struct {
	Data    *GetUserDataOkResponseData `json:"data,omitempty"`
	Detail  *string                    `json:"detail,omitempty"`
	Error   any                        `json:"error,omitempty"`
	Success *bool                      `json:"success,omitempty"`
	touched map[string]bool
}

func (g *GetUserDataOkResponse) GetData() *GetUserDataOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetUserDataOkResponse) SetData(data GetUserDataOkResponseData) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = &data
}

func (g *GetUserDataOkResponse) SetDataNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Data"] = true
	g.Data = nil
}

func (g *GetUserDataOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetUserDataOkResponse) SetDetail(detail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = &detail
}

func (g *GetUserDataOkResponse) SetDetailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Detail"] = true
	g.Detail = nil
}

func (g *GetUserDataOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetUserDataOkResponse) SetError(error any) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = error
}

func (g *GetUserDataOkResponse) SetErrorNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Error"] = true
	g.Error = nil
}

func (g *GetUserDataOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
}

func (g *GetUserDataOkResponse) SetSuccess(success bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = &success
}

func (g *GetUserDataOkResponse) SetSuccessNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Success"] = true
	g.Success = nil
}

func (g GetUserDataOkResponse) MarshalJSON() ([]byte, error) {
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

func (g GetUserDataOkResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetUserDataOkResponse to string"
	}
	return string(jsonData)
}

type GetUserDataOkResponseData struct {
	AuthId           *string   `json:"auth_id,omitempty"`
	BaseEmail        *string   `json:"base_email,omitempty"`
	CooldownUntil    *string   `json:"cooldown_until,omitempty"`
	CreatedAt        *string   `json:"created_at,omitempty"`
	Customer         *string   `json:"customer,omitempty"`
	Email            *string   `json:"email,omitempty"`
	Id               *float64  `json:"id,omitempty"`
	IsSubscribed     *bool     `json:"is_subscribed,omitempty"`
	Plan             *float64  `json:"plan,omitempty"`
	PremiumExpiresAt *string   `json:"premium_expires_at,omitempty"`
	Server           *float64  `json:"server,omitempty"`
	Settings         *Settings `json:"settings,omitempty"`
	TotalDownloaded  *float64  `json:"total_downloaded,omitempty"`
	UpdatedAt        *string   `json:"updated_at,omitempty"`
	UserReferral     *string   `json:"user_referral,omitempty"`
	touched          map[string]bool
}

func (g *GetUserDataOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetUserDataOkResponseData) SetAuthId(authId string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = &authId
}

func (g *GetUserDataOkResponseData) SetAuthIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["AuthId"] = true
	g.AuthId = nil
}

func (g *GetUserDataOkResponseData) GetBaseEmail() *string {
	if g == nil {
		return nil
	}
	return g.BaseEmail
}

func (g *GetUserDataOkResponseData) SetBaseEmail(baseEmail string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["BaseEmail"] = true
	g.BaseEmail = &baseEmail
}

func (g *GetUserDataOkResponseData) SetBaseEmailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["BaseEmail"] = true
	g.BaseEmail = nil
}

func (g *GetUserDataOkResponseData) GetCooldownUntil() *string {
	if g == nil {
		return nil
	}
	return g.CooldownUntil
}

func (g *GetUserDataOkResponseData) SetCooldownUntil(cooldownUntil string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CooldownUntil"] = true
	g.CooldownUntil = &cooldownUntil
}

func (g *GetUserDataOkResponseData) SetCooldownUntilNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CooldownUntil"] = true
	g.CooldownUntil = nil
}

func (g *GetUserDataOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetUserDataOkResponseData) SetCreatedAt(createdAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = &createdAt
}

func (g *GetUserDataOkResponseData) SetCreatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["CreatedAt"] = true
	g.CreatedAt = nil
}

func (g *GetUserDataOkResponseData) GetCustomer() *string {
	if g == nil {
		return nil
	}
	return g.Customer
}

func (g *GetUserDataOkResponseData) SetCustomer(customer string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Customer"] = true
	g.Customer = &customer
}

func (g *GetUserDataOkResponseData) SetCustomerNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Customer"] = true
	g.Customer = nil
}

func (g *GetUserDataOkResponseData) GetEmail() *string {
	if g == nil {
		return nil
	}
	return g.Email
}

func (g *GetUserDataOkResponseData) SetEmail(email string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Email"] = true
	g.Email = &email
}

func (g *GetUserDataOkResponseData) SetEmailNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Email"] = true
	g.Email = nil
}

func (g *GetUserDataOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetUserDataOkResponseData) SetId(id float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GetUserDataOkResponseData) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GetUserDataOkResponseData) GetIsSubscribed() *bool {
	if g == nil {
		return nil
	}
	return g.IsSubscribed
}

func (g *GetUserDataOkResponseData) SetIsSubscribed(isSubscribed bool) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["IsSubscribed"] = true
	g.IsSubscribed = &isSubscribed
}

func (g *GetUserDataOkResponseData) SetIsSubscribedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["IsSubscribed"] = true
	g.IsSubscribed = nil
}

func (g *GetUserDataOkResponseData) GetPlan() *float64 {
	if g == nil {
		return nil
	}
	return g.Plan
}

func (g *GetUserDataOkResponseData) SetPlan(plan float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Plan"] = true
	g.Plan = &plan
}

func (g *GetUserDataOkResponseData) SetPlanNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Plan"] = true
	g.Plan = nil
}

func (g *GetUserDataOkResponseData) GetPremiumExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.PremiumExpiresAt
}

func (g *GetUserDataOkResponseData) SetPremiumExpiresAt(premiumExpiresAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["PremiumExpiresAt"] = true
	g.PremiumExpiresAt = &premiumExpiresAt
}

func (g *GetUserDataOkResponseData) SetPremiumExpiresAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["PremiumExpiresAt"] = true
	g.PremiumExpiresAt = nil
}

func (g *GetUserDataOkResponseData) GetServer() *float64 {
	if g == nil {
		return nil
	}
	return g.Server
}

func (g *GetUserDataOkResponseData) SetServer(server float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Server"] = true
	g.Server = &server
}

func (g *GetUserDataOkResponseData) SetServerNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Server"] = true
	g.Server = nil
}

func (g *GetUserDataOkResponseData) GetSettings() *Settings {
	if g == nil {
		return nil
	}
	return g.Settings
}

func (g *GetUserDataOkResponseData) SetSettings(settings Settings) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Settings"] = true
	g.Settings = &settings
}

func (g *GetUserDataOkResponseData) SetSettingsNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Settings"] = true
	g.Settings = nil
}

func (g *GetUserDataOkResponseData) GetTotalDownloaded() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalDownloaded
}

func (g *GetUserDataOkResponseData) SetTotalDownloaded(totalDownloaded float64) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalDownloaded"] = true
	g.TotalDownloaded = &totalDownloaded
}

func (g *GetUserDataOkResponseData) SetTotalDownloadedNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["TotalDownloaded"] = true
	g.TotalDownloaded = nil
}

func (g *GetUserDataOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetUserDataOkResponseData) SetUpdatedAt(updatedAt string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = &updatedAt
}

func (g *GetUserDataOkResponseData) SetUpdatedAtNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UpdatedAt"] = true
	g.UpdatedAt = nil
}

func (g *GetUserDataOkResponseData) GetUserReferral() *string {
	if g == nil {
		return nil
	}
	return g.UserReferral
}

func (g *GetUserDataOkResponseData) SetUserReferral(userReferral string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UserReferral"] = true
	g.UserReferral = &userReferral
}

func (g *GetUserDataOkResponseData) SetUserReferralNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["UserReferral"] = true
	g.UserReferral = nil
}

func (g GetUserDataOkResponseData) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["AuthId"] && g.AuthId == nil {
		data["auth_id"] = nil
	} else if g.AuthId != nil {
		data["auth_id"] = g.AuthId
	}

	if g.touched["BaseEmail"] && g.BaseEmail == nil {
		data["base_email"] = nil
	} else if g.BaseEmail != nil {
		data["base_email"] = g.BaseEmail
	}

	if g.touched["CooldownUntil"] && g.CooldownUntil == nil {
		data["cooldown_until"] = nil
	} else if g.CooldownUntil != nil {
		data["cooldown_until"] = g.CooldownUntil
	}

	if g.touched["CreatedAt"] && g.CreatedAt == nil {
		data["created_at"] = nil
	} else if g.CreatedAt != nil {
		data["created_at"] = g.CreatedAt
	}

	if g.touched["Customer"] && g.Customer == nil {
		data["customer"] = nil
	} else if g.Customer != nil {
		data["customer"] = g.Customer
	}

	if g.touched["Email"] && g.Email == nil {
		data["email"] = nil
	} else if g.Email != nil {
		data["email"] = g.Email
	}

	if g.touched["Id"] && g.Id == nil {
		data["id"] = nil
	} else if g.Id != nil {
		data["id"] = g.Id
	}

	if g.touched["IsSubscribed"] && g.IsSubscribed == nil {
		data["is_subscribed"] = nil
	} else if g.IsSubscribed != nil {
		data["is_subscribed"] = g.IsSubscribed
	}

	if g.touched["Plan"] && g.Plan == nil {
		data["plan"] = nil
	} else if g.Plan != nil {
		data["plan"] = g.Plan
	}

	if g.touched["PremiumExpiresAt"] && g.PremiumExpiresAt == nil {
		data["premium_expires_at"] = nil
	} else if g.PremiumExpiresAt != nil {
		data["premium_expires_at"] = g.PremiumExpiresAt
	}

	if g.touched["Server"] && g.Server == nil {
		data["server"] = nil
	} else if g.Server != nil {
		data["server"] = g.Server
	}

	if g.touched["Settings"] && g.Settings == nil {
		data["settings"] = nil
	} else if g.Settings != nil {
		data["settings"] = g.Settings
	}

	if g.touched["TotalDownloaded"] && g.TotalDownloaded == nil {
		data["total_downloaded"] = nil
	} else if g.TotalDownloaded != nil {
		data["total_downloaded"] = g.TotalDownloaded
	}

	if g.touched["UpdatedAt"] && g.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if g.UpdatedAt != nil {
		data["updated_at"] = g.UpdatedAt
	}

	if g.touched["UserReferral"] && g.UserReferral == nil {
		data["user_referral"] = nil
	} else if g.UserReferral != nil {
		data["user_referral"] = g.UserReferral
	}

	return json.Marshal(data)
}

func (g GetUserDataOkResponseData) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetUserDataOkResponseData to string"
	}
	return string(jsonData)
}

type Settings struct {
	Anothersetting *string `json:"anothersetting,omitempty"`
	Setting        *string `json:"setting,omitempty"`
	touched        map[string]bool
}

func (s *Settings) GetAnothersetting() *string {
	if s == nil {
		return nil
	}
	return s.Anothersetting
}

func (s *Settings) SetAnothersetting(anothersetting string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Anothersetting"] = true
	s.Anothersetting = &anothersetting
}

func (s *Settings) SetAnothersettingNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Anothersetting"] = true
	s.Anothersetting = nil
}

func (s *Settings) GetSetting() *string {
	if s == nil {
		return nil
	}
	return s.Setting
}

func (s *Settings) SetSetting(setting string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Setting"] = true
	s.Setting = &setting
}

func (s *Settings) SetSettingNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Setting"] = true
	s.Setting = nil
}

func (s Settings) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Anothersetting"] && s.Anothersetting == nil {
		data["anothersetting"] = nil
	} else if s.Anothersetting != nil {
		data["anothersetting"] = s.Anothersetting
	}

	if s.touched["Setting"] && s.Setting == nil {
		data["setting"] = nil
	} else if s.Setting != nil {
		data["setting"] = s.Setting
	}

	return json.Marshal(data)
}

func (s Settings) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Settings to string"
	}
	return string(jsonData)
}
