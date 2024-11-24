package webdownloadsdebrid

import (
	"context"
	"time"
	restClient "torbox-sdk-go/internal/clients/rest"
	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/shared"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type WebDownloadsDebridService struct {
	manager *configmanager.ConfigManager
}

func NewWebDownloadsDebridService(manager *configmanager.ConfigManager) *WebDownloadsDebridService {
	return &WebDownloadsDebridService{
		manager: manager,
	}
}

func (api *WebDownloadsDebridService) getConfig() *torboxapiconfig.Config {
	return api.manager.GetWebDownloadsDebrid()
}

func (api *WebDownloadsDebridService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *WebDownloadsDebridService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *WebDownloadsDebridService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// ### Overview
//
// Creates a web download under your account. Simply send a link to any file on the internet. Once it has been checked, it will begin downloading assuming your account has available active download slots, and they aren't too large.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *WebDownloadsDebridService) CreateWebDownload(ctx context.Context, apiVersion string, createWebDownloadRequest CreateWebDownloadRequest) (*shared.TorboxApiResponse[CreateWebDownloadOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/webdl/createwebdownload").
		WithConfig(config).
		WithBody(createWebDownloadRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CreateWebDownloadOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[CreateWebDownloadOkResponse](err)
	}

	return shared.NewTorboxApiResponse[CreateWebDownloadOkResponse](resp), nil
}

// ### Overview
//
// Controls a web download. By sending the web download's ID and the type of operation you want to perform, it will send that request to the debrid client.
//
// Operations are either:
//
// - **Delete** `deletes the download from the client and your account permanently`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *WebDownloadsDebridService) ControlWebDownload(ctx context.Context, apiVersion string, params ControlWebDownloadRequestParams) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/webdl/controlwebdownload").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
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
// Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 1 hour for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *WebDownloadsDebridService) RequestDownloadLink2(ctx context.Context, apiVersion string, params RequestDownloadLink2RequestParams) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/requestdl").
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
// Gets the user's web download list. This gives you the needed information to perform other usenet actions. Unlike Torrents, this information is updated on its own every 5 seconds for live web downloads.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *WebDownloadsDebridService) GetWebDownloadList(ctx context.Context, apiVersion string, params GetWebDownloadListRequestParams) (*shared.TorboxApiResponse[GetWebDownloadListOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/mylist").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetWebDownloadListOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetWebDownloadListOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetWebDownloadListOkResponse](resp), nil
}

// ### Overview
//
// Takes in a list of comma separated usenet hashes and checks if the web download is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as:
// `?hash=XXXX&hash=XXXX&hash=XXXX`
//
// or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX`
// and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// To get the hash of a web download, pass the link through an md5 hash algo and it will return the proper hash for it.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *WebDownloadsDebridService) GetWebDownloadCachedAvailability(ctx context.Context, apiVersion string, params GetWebDownloadCachedAvailabilityRequestParams) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/webdl/checkcached").
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