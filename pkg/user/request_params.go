package user

type GetUserDataRequestParams struct {
	Settings *string `queryParam:"settings"`
}

func (params *GetUserDataRequestParams) SetSettings(settings string) {
	params.Settings = &settings
}

type AddReferralToAccountRequestParams struct {
	Referral *string `queryParam:"referral"`
}

func (params *AddReferralToAccountRequestParams) SetReferral(referral string) {
	params.Referral = &referral
}