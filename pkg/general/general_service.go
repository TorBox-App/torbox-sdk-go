package general

import (
	"context"
	"time"
	restClient "torbox-sdk-go/internal/clients/rest"
	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/shared"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type GeneralService struct {
	manager *configmanager.ConfigManager
}

func NewGeneralService() *GeneralService {
	return &GeneralService{
		manager: configmanager.NewConfigManager(torboxapiconfig.Config{}),
	}
}

func (api *GeneralService) WithConfigManager(manager *configmanager.ConfigManager) *GeneralService {
	api.manager = manager
	return api
}

func (api *GeneralService) getConfig() *torboxapiconfig.Config {
	return api.manager.GetGeneral()
}

func (api *GeneralService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *GeneralService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *GeneralService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// ### Overview
//
// Simply gets if the TorBox API is available for use or not. Also might include information about downtimes.
//
// ### Authorization
//
// None needed.
func (api *GeneralService) GetUpStatus(ctx context.Context) (*shared.TorboxApiResponse[GetUpStatusOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetUpStatusOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetUpStatusOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetUpStatusOkResponse](resp), nil
}

// ### Overview
//
// Simply gets general stats about the TorBox service.
//
// ### Authorization
//
// None needed.
func (api *GeneralService) GetStats(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[GetStatsOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/stats").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetStatsOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetStatsOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetStatsOkResponse](resp), nil
}

// ### Overview
//
// Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for people who want updates on the TorBox changelog. Includes all the necessary items to make this compatible with RSS feed viewers for notifications, and proper HTML viewing.
//
// ### Authorization
//
// None needed.
func (api *GeneralService) GetChangelogsRssFeed(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[[]byte], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/changelogs/rss").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeText).
		Build()

	client := restClient.NewRestClient[[]byte](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[[]byte](err)
	}

	return shared.NewTorboxApiResponse[[]byte](resp), nil
}

// ### Overview
//
// Gets most recent 100 changelogs from [https://feedback.torbox.app/changelog.](https://feedback.torbox.app/changelog.) This is useful for developers who want to integrate the changelog into their apps for their users to see. Includes content in HTML and markdown for developers to easily render the text in a fancy yet simple way.
//
// ### Authorization
//
// None needed.
func (api *GeneralService) GetChangelogsJson(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[GetChangelogsJsonOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/changelogs/json").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetChangelogsJsonOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetChangelogsJsonOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetChangelogsJsonOkResponse](resp), nil
}

// ### Overview
//
// Gets CDN speedtest files. This can be used for speedtesting TorBox for users or other usages, such as checking download speeds for verification. Provides all necessary data such as region, server name, and even coordinates. Uses the requesting IP to determine if the server is the closest to the user.
//
// You also have the ability to choose between long tests or short tests via the "test_length" parameter. You may also force the region selection by passing the "region" as a specific region.
//
// ### Authorization
//
// None needed.
func (api *GeneralService) GetSpeedtestFiles(ctx context.Context, apiVersion string, params GetSpeedtestFilesRequestParams) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/speedtest").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
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
