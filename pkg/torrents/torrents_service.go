package torrents

import (
	"context"
	"time"
	restClient "torbox-sdk-go/internal/clients/rest"
	"torbox-sdk-go/internal/clients/rest/httptransport"
	"torbox-sdk-go/internal/configmanager"
	"torbox-sdk-go/pkg/shared"
	"torbox-sdk-go/pkg/torboxapiconfig"
)

type TorrentsService struct {
	manager *configmanager.ConfigManager
}

func NewTorrentsService(manager *configmanager.ConfigManager) *TorrentsService {
	return &TorrentsService{
		manager: manager,
	}
}

func (api *TorrentsService) getConfig() *torboxapiconfig.Config {
	return api.manager.GetTorrents()
}

func (api *TorrentsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *TorrentsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *TorrentsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// ### Overview
//
// Creates a torrent under your account. Simply send **either** a magnet link, or a torrent file. Once they have been checked, they will begin downloading assuming your account has available active download slots, and they aren't too large.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) CreateTorrent(ctx context.Context, apiVersion string, createTorrentRequest CreateTorrentRequest) (*shared.TorboxApiResponse[CreateTorrentOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/createtorrent").
		WithConfig(config).
		WithBody(createTorrentRequest).
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CreateTorrentOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[CreateTorrentOkResponse](err)
	}

	return shared.NewTorboxApiResponse[CreateTorrentOkResponse](resp), nil
}

// ### Overview
//
// Controls a torrent. By sending the torrent's ID and the type of operation you want to perform, it will send that request to the torrent client.
//
// Operations are either:
//
// - **Reannounce** `reannounces the torrent to get new peers`
//
// - **Delete** `deletes the torrent from the client and your account permanently`
//
// - **Resume** `resumes a paused torrent`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) ControlTorrent(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[ControlTorrentOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/controltorrent").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ControlTorrentOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[ControlTorrentOkResponse](err)
	}

	return shared.NewTorboxApiResponse[ControlTorrentOkResponse](resp), nil
}

// ### Overview
//
// Controls a queued torrent. By sending the queued torrent's ID and the type of operation you want to perform, it will perform that action on the queued torrent.
//
// Operations are either:
//
// - **Delete** `deletes the queued torrent from your account`
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) ControlQueuedTorrent(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[ControlQueuedTorrentOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/{api_version}/api/torrents/controlqueued").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("api_version", apiVersion).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ControlQueuedTorrentOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[ControlQueuedTorrentOkResponse](err)
	}

	return shared.NewTorboxApiResponse[ControlQueuedTorrentOkResponse](resp), nil
}

// ### Overview
//
// Requests the download link from the server. Because downloads are metered, TorBox cannot afford to allow free access to the links directly. This endpoint opens the link for 1 hour for downloads. Once a download is started, the user has nearly unlilimited time to download the file. The 1 hour time limit is simply for starting downloads. This prevents long term link sharing.
//
// ### Authorization
//
// Requires an API key as a parameter for the `token` parameter.
func (api *TorrentsService) RequestDownloadLink(ctx context.Context, apiVersion string, params RequestDownloadLinkRequestParams) (*shared.TorboxApiResponse[RequestDownloadLinkOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/requestdl").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[RequestDownloadLinkOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[RequestDownloadLinkOkResponse](err)
	}

	return shared.NewTorboxApiResponse[RequestDownloadLinkOkResponse](resp), nil
}

// ### Overview
//
// Gets the user's torrent list. This gives you the needed information to perform other torrent actions. This information only gets updated every 600 seconds, or when the _Request Update On Torrent_ request is sent to the _relay API_.
//
// #### Download States:
//
// - "downloading" -> The torrent is currently downloading.
//
// - "uploading" -> The torrent is currently seeding.
//
// - "stalled (no seeds)" -> The torrent is trying to download, but there are no seeds connected to download from.
//
// - "paused" -> The torrent is paused.
//
// - "completed" -> The torrent is completely downloaded. Do not use this for download completion status.
//
// - "cached" -> The torrent is cached from the server.
//
// - "metaDL" -> The torrent is downloading metadata from the hoard.
//
// - "checkingResumeData" -> The torrent is checking resumable data.
//
// All other statuses are basic qBittorrent states. [Check out here for the full list](https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#torrent-management).
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) GetTorrentList(ctx context.Context, apiVersion string, params GetTorrentListRequestParams) (*shared.TorboxApiResponse[GetTorrentListOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/mylist").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetTorrentListOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetTorrentListOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetTorrentListOkResponse](resp), nil
}

// ### Overview
//
// Takes in a list of comma separated torrent hashes and checks if the torrent is cached. This endpoint only gets a max of around 100 at a time, due to http limits in queries. If you want to do more, you can simply do more hash queries. Such as:
// `?hash=XXXX&hash=XXXX&hash=XXXX`
//
// or `?hash=XXXX,XXXX&hash=XXXX&hash=XXXX,XXXX`
// and this will work too. Performance is very fast. Less than 1 second per 100. Time is approximately O(log n) time for those interested in taking it to its max. That is without caching as well. This endpoint stores a cache for an hour.
//
// You may also pass a `format` parameter with the format you want the data in. Options are either `object` or `list`. You can view examples of both below.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) GetTorrentCachedAvailability(ctx context.Context, apiVersion string, params GetTorrentCachedAvailabilityRequestParams) (*shared.TorboxApiResponse[GetTorrentCachedAvailabilityOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/checkcached").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetTorrentCachedAvailabilityOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetTorrentCachedAvailabilityOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetTorrentCachedAvailabilityOkResponse](resp), nil
}

// ### Overview
//
// Uses Meilisearch to search for scraped torrents. This is a basic torrent aggregator system and has no real relation to TorBox. It is simply a tool you can use. It does not have cache information, or anything special like that, and will not have any of that information. This is simply a torrent search, nothing more.
//
// You may use this for anything. TorBox uses it in the website to make it easy for users to find torrents without having to go and find them on sketchy websites.
//
// ### Authorization
//
// None required.
func (api *TorrentsService) SearchAllTorrentsFromScraper(ctx context.Context, apiVersion string, params SearchAllTorrentsFromScraperRequestParams) (*shared.TorboxApiResponse[SearchAllTorrentsFromScraperOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/search").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SearchAllTorrentsFromScraperOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[SearchAllTorrentsFromScraperOkResponse](err)
	}

	return shared.NewTorboxApiResponse[SearchAllTorrentsFromScraperOkResponse](resp), nil
}

// ### Overview
//
// Exports the magnet or torrent file. Requires a type to be passed. If type is **magnet**, it will return a JSON response with the magnet as a string in the _data_ key. If type is **file**, it will return a bittorrent file as a download. Not compatible with cached downloads.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) ExportTorrentData(ctx context.Context, apiVersion string, params ExportTorrentDataRequestParams) (*shared.TorboxApiResponse[ExportTorrentDataOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/exportdata").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExportTorrentDataOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[ExportTorrentDataOkResponse](err)
	}

	return shared.NewTorboxApiResponse[ExportTorrentDataOkResponse](resp), nil
}

// ### Overview
//
// A general route that allows you to give a hash, and TorBox will return data about the torrent. This data is retrieved from the Bittorrent network, so expect it to take some time. If the request goes longer than 10 seconds, TorBox will cancel it. You can adjust this if you like, but the default is 10 seconds. This route is cached as well, so subsequent requests will be instant.
//
// ### Authorization
//
// None required.
func (api *TorrentsService) GetTorrentInfo(ctx context.Context, apiVersion string, params GetTorrentInfoRequestParams) (*shared.TorboxApiResponse[GetTorrentInfoOkResponse], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/torrentinfo").
		WithConfig(config).
		AddPathParam("api_version", apiVersion).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetTorrentInfoOkResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewTorboxApiError[GetTorrentInfoOkResponse](err)
	}

	return shared.NewTorboxApiResponse[GetTorrentInfoOkResponse](resp), nil
}

// ### Overview
//
// Retrieves all of a user's queued torrents.
//
// ### Authorization
//
// Requires an API key using the Authorization Bearer Header.
func (api *TorrentsService) GetQueuedTorrents(ctx context.Context, apiVersion string) (*shared.TorboxApiResponse[any], *shared.TorboxApiError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/{api_version}/api/torrents/getqueued").
		WithConfig(config).
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
