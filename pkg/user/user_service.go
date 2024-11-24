package user

import (
	"context"
	"time"
	restClient "torbox-sdk-go/internal/clients/rest"
	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/shared"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type UserService struct {
	manager *configmanager.ConfigManager
}

func NewUserService(manager *configmanager.ConfigManager) *UserService {
	return &UserService{
		manager: manager,
	}
}

func (api *UserService) getConfig() *torboxapiconfig.Config {
	return api.manager.GetUser()
}

func (api *UserService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *UserService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *UserService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// ### Overview
//
// If you want a new API token, or your old one has been compromised, you may request a new one. If you happen to forget the token, go the [TorBox settings page ](https://torbox.app/settings) and copy the current one. If this still doesn't work, you may contact us at our support email for a new one.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header as well as passing the `session_token` from the website to be passed in the body. You can find the `session_token` in the localStorage of your browser on any TorBox.app page under the key `torbox_session_token`. This is a temporary token that only lasts for 1 hour, which is why it is used here to verify the identity of a user as well as their API token.
func (api *UserService) RefreshApiToken(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/user/refreshtoken").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[any](err)
	}

	return shared.NewTorboxApiResponse[any](resp), nil
}

// ### Overview
//
// Gets a users account data and information.
//
// ### Plans
//
// `0` is Free plan
//
// `1` is Essential plan (_$3 plan_)
//
// `2` is Pro plan (_$10 plan_)
//
// `3` is Standard plan (_$5 plan_)
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *UserService) GetUserData(ctx context.Context, apiVersion string, params GetUserDataRequestParams) (*shared.TorboxApiResponse[GetUserDataOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/user/me").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetUserDataOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetUserDataOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetUserDataOkResponse](resp), nil
}

// Add Referral To Account
func (api *UserService) AddReferralToAccount(ctx context.Context, apiVersion string, params AddReferralToAccountRequestParams) (*shared.TorboxApiResponse[AddReferralToAccountOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/user/addreferral").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[AddReferralToAccountOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[AddReferralToAccountOkResponse](err)
	}

	return shared.NewTorboxApiResponse[AddReferralToAccountOkResponse](resp), nil
}
