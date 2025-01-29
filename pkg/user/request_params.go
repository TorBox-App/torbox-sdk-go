package user

type GetUserDataRequestParams struct {
	Settings *string `explode:"true" serializationStyle:"form" queryParam:"settings"`
}

func (params *GetUserDataRequestParams) SetSettings(settings string) {
	params.Settings = &settings
}

type AddReferralToAccountRequestParams struct {
	Referral *string `explode:"true" serializationStyle:"form" queryParam:"referral"`
}

func (params *AddReferralToAccountRequestParams) SetReferral(referral string) {
	params.Referral = &referral
}
