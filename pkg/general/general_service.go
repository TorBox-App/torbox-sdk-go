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

func NewGeneralService(manager *configmanager.ConfigManager) *GeneralService {
	return &GeneralService{
		manager: manager,
	}
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
