package user

type GetUserDataOkResponse struct {
	Data    *GetUserDataOkResponseData `json:"data,omitempty"`
	Detail  *string                    `json:"detail,omitempty"`
	Error   any                        `json:"error,omitempty"`
	Success *bool                      `json:"success,omitempty"`
}

func (g *GetUserDataOkResponse) SetData(data GetUserDataOkResponseData) {
	g.Data = &data
}

func (g *GetUserDataOkResponse) GetData() *GetUserDataOkResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetUserDataOkResponse) SetDetail(detail string) {
	g.Detail = &detail
}

func (g *GetUserDataOkResponse) GetDetail() *string {
	if g == nil {
		return nil
	}
	return g.Detail
}

func (g *GetUserDataOkResponse) SetError(error any) {
	g.Error = error
}

func (g *GetUserDataOkResponse) GetError() any {
	if g == nil {
		return nil
	}
	return g.Error
}

func (g *GetUserDataOkResponse) SetSuccess(success bool) {
	g.Success = &success
}

func (g *GetUserDataOkResponse) GetSuccess() *bool {
	if g == nil {
		return nil
	}
	return g.Success
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
}

func (g *GetUserDataOkResponseData) SetAuthId(authId string) {
	g.AuthId = &authId
}

func (g *GetUserDataOkResponseData) GetAuthId() *string {
	if g == nil {
		return nil
	}
	return g.AuthId
}

func (g *GetUserDataOkResponseData) SetBaseEmail(baseEmail string) {
	g.BaseEmail = &baseEmail
}

func (g *GetUserDataOkResponseData) GetBaseEmail() *string {
	if g == nil {
		return nil
	}
	return g.BaseEmail
}

func (g *GetUserDataOkResponseData) SetCooldownUntil(cooldownUntil string) {
	g.CooldownUntil = &cooldownUntil
}

func (g *GetUserDataOkResponseData) GetCooldownUntil() *string {
	if g == nil {
		return nil
	}
	return g.CooldownUntil
}

func (g *GetUserDataOkResponseData) SetCreatedAt(createdAt string) {
	g.CreatedAt = &createdAt
}

func (g *GetUserDataOkResponseData) GetCreatedAt() *string {
	if g == nil {
		return nil
	}
	return g.CreatedAt
}

func (g *GetUserDataOkResponseData) SetCustomer(customer string) {
	g.Customer = &customer
}

func (g *GetUserDataOkResponseData) GetCustomer() *string {
	if g == nil {
		return nil
	}
	return g.Customer
}

func (g *GetUserDataOkResponseData) SetEmail(email string) {
	g.Email = &email
}

func (g *GetUserDataOkResponseData) GetEmail() *string {
	if g == nil {
		return nil
	}
	return g.Email
}

func (g *GetUserDataOkResponseData) SetId(id float64) {
	g.Id = &id
}

func (g *GetUserDataOkResponseData) GetId() *float64 {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GetUserDataOkResponseData) SetIsSubscribed(isSubscribed bool) {
	g.IsSubscribed = &isSubscribed
}

func (g *GetUserDataOkResponseData) GetIsSubscribed() *bool {
	if g == nil {
		return nil
	}
	return g.IsSubscribed
}

func (g *GetUserDataOkResponseData) SetPlan(plan float64) {
	g.Plan = &plan
}

func (g *GetUserDataOkResponseData) GetPlan() *float64 {
	if g == nil {
		return nil
	}
	return g.Plan
}

func (g *GetUserDataOkResponseData) SetPremiumExpiresAt(premiumExpiresAt string) {
	g.PremiumExpiresAt = &premiumExpiresAt
}

func (g *GetUserDataOkResponseData) GetPremiumExpiresAt() *string {
	if g == nil {
		return nil
	}
	return g.PremiumExpiresAt
}

func (g *GetUserDataOkResponseData) SetServer(server float64) {
	g.Server = &server
}

func (g *GetUserDataOkResponseData) GetServer() *float64 {
	if g == nil {
		return nil
	}
	return g.Server
}

func (g *GetUserDataOkResponseData) SetSettings(settings Settings) {
	g.Settings = &settings
}

func (g *GetUserDataOkResponseData) GetSettings() *Settings {
	if g == nil {
		return nil
	}
	return g.Settings
}

func (g *GetUserDataOkResponseData) SetTotalDownloaded(totalDownloaded float64) {
	g.TotalDownloaded = &totalDownloaded
}

func (g *GetUserDataOkResponseData) GetTotalDownloaded() *float64 {
	if g == nil {
		return nil
	}
	return g.TotalDownloaded
}

func (g *GetUserDataOkResponseData) SetUpdatedAt(updatedAt string) {
	g.UpdatedAt = &updatedAt
}

func (g *GetUserDataOkResponseData) GetUpdatedAt() *string {
	if g == nil {
		return nil
	}
	return g.UpdatedAt
}

func (g *GetUserDataOkResponseData) SetUserReferral(userReferral string) {
	g.UserReferral = &userReferral
}

func (g *GetUserDataOkResponseData) GetUserReferral() *string {
	if g == nil {
		return nil
	}
	return g.UserReferral
}

type Settings struct {
	Anothersetting *string `json:"anothersetting,omitempty"`
	Setting        *string `json:"setting,omitempty"`
}

func (s *Settings) SetAnothersetting(anothersetting string) {
	s.Anothersetting = &anothersetting
}

func (s *Settings) GetAnothersetting() *string {
	if s == nil {
		return nil
	}
	return s.Anothersetting
}

func (s *Settings) SetSetting(setting string) {
	s.Setting = &setting
}

func (s *Settings) GetSetting() *string {
	if s == nil {
		return nil
	}
	return s.Setting
}
