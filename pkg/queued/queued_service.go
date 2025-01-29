package queued

import (
	"context"
	"time"
	restClient "torbox-sdk-go/internal/clients/rest"
	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/shared"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type QueuedService struct {
	manager *configmanager.ConfigManager
}

func NewQueuedService() *QueuedService {
	return &QueuedService{
		manager: configmanager.NewConfigManager(torboxapiconfig.Config{}),
	}
}

func (api *QueuedService) WithConfigManager(manager *configmanager.ConfigManager) *QueuedService {
	api.manager = manager
	return api
}

func (api *QueuedService) getConfig() *torboxapiconfig.Config {
	return api.manager.GetQueued()
}

func (api *QueuedService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *QueuedService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *QueuedService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// ### Overview
//
// Retrieves all of a user's queued downloads by type. If you want to get all 3 types, "torrent", "usenet" and "webdl" then you will need to run this request 3 times, each with the different type.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *QueuedService) GetQueuedDownloads(ctx context.Context, apiVersion string, params GetQueuedDownloadsRequestParams) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/queued/getqueued").
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

// ### Overview
//
// Controls a queued torrent. By sending the queued torrent's ID and the type of operation you want to perform, it will perform that action on the queued torrent.
//
// Operations are either:
//
// - **Delete** `deletes the queued download from your account`
//
// - **Start** `starts a queued download, cannot be used with the "all" parameter`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *QueuedService) ControlQueuedDownloads(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/queued/controlqueued").
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
