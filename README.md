# TorboxApi Go SDK 1.0.0

Welcome to the TorboxApi SDK documentation. This guide will help you get started with integrating and using the TorboxApi SDK in your project.

## Versions

- API version: `1.0.0`
- SDK version: `1.0.0`

## About the API

### Full documentation for TorBox Current API Base: `https://api.torbox.app` Current API Version: `v1` You can find more API docs here: [http://api.torbox.app/docs](http://api.torbox.app/docs) (they aren't as great as these, but is in a format most people would be familiar to). ### General Information - You can rely on both `success` booleans or status codes to determine if a call was a success. Status code `200` always means a success. `403` means authentication error. `500` means something went wrong on TorBox's end. `400` means the user did something wrong, or an input wasn't correct, or expected. - There will always be a user-friendly message in `detail` for a reason why a call was a failure, or a success message. You may forward these directly to users. - The `/usenet` and the `/webdl` API endpoints are nearly the same as the `/torrents` API endpoints apart from some different named inputs and outputs. - All outputs are JSON unless otherwise specified. Check out the examples to be sure. ### Rate Limits - Unless stated below, all endpoints are rate limited to **5/sec** per IP, no edge rate limiting. - `POST` /torrents/createtorrent is rate limited to **50/hour** per IP and **10/min** at edge. - `POST` /usenet/createusenetdownload is rate limited to **25/hour** per IP and **10/min** at edge. - `POST` /webdl/createwebdownload is rate limited to **25/hour** per IP and **10/min** at edge. - If you have a specific application, please contact us with your IP and reason to get unlimited requests. ### Standard Response `json {      success : boolean, // whether a response is successful or not      error : string, // an error code if there was an error, or null if success      detail : string, // a user-friendly message      data : any // usually an object/dict, but can be anything, check the examples }  ` ### Errors Table Errors codes are part of the standard response type. If the code ends in ERROR , the error is the server's fault else that error is something that the client caused. | **Error Code** | Error Meaning | | --- | --- | | DATABASE*ERROR | Could not access internal database/memory store information. | | UNKNOWN_ERROR | The reason for the error is unknown. Usually there will be error data attached in the data key. In these cases please report the request to [contact@torbox.app](https://mailto:contact@torbox.app). | | NO_AUTH | There are no provided credentials. | | BAD_TOKEN | The provided token is invalid. | | AUTH_ERROR | There was an error verifying the given authentication. | | INVALID_OPTION | The provided option is invalid. | | REDIRECT_ERROR | The server tried redirecting, but it faulted. | | OAUTH_VERIFICATION_ERROR | The server tried verifying your OAuth token, but it was not accepted by the provider. | | ENDPOINT_NOT_FOUND | If you have hit an endpoint that doesn't exist. | | ITEM_NOT_FOUND | The item you queried cannot be found. | | PLAN_RESTRICTED_FEATURE | This feature is restricted to users of higher plans. The user is recommended to upgrade their plan to use this endpoint. | | DUPLICATE_ITEM | This item already exists. | | BOZO_RSS_FEED | This RSS feed is invalid or not a well-formed XML. | | SELLIX_ERROR | There was an error with the Sellix API. Usually in the case of payments. | | TOO_MUCH_DATA | Client sent too much data to the API. Please keep requests under 100MB in size. | | DOWNLOAD_TOO_LARGE | This download is oversized for the user's plan. The user is recommended to upgrade their plan to download this file. <br> <br>Free Plan Limit: **10737418240** bytes <br>Essential Plan Limit: **214748364800** bytes <br>Standard Plan Limit: **214748364800** bytes <br>Pro Plan Limit: **536870912000** bytes | | MISSING_REQUIRED_OPTION | The API is missing required information to process the request. | | TOO_MANY_OPTIONS | Client sent too many options. Usually this has to do with the API requiring only 1 option but the client sent more than the required. | | BOZO_TORRENT | The torrent sent is not a valid torrent. | | NO_SERVERS_AVAILABLE_ERROR | There are no download servers available to handle this request. This should never happen. If you receieve this error, please contact us at [contact@torbox.app](https://mailto:contact@torbox.app). | | MONTHLY_LIMIT | User has hit the maximum monthly limit. It is recommended user upgrade their account to be able to download more. | | COOLDOWN_LIMIT | User is on download cooldown. It is recommended user upgrade their account to be able to bypass this restriction. | | ACTIVE_LIMIT | User has hit their max active download limit. It is recommended user upgrade their account or purchase addons to bypass this restriction. | | DOWNLOAD_SERVER_ERROR | There was an error interacting with the download on the download server. It is recommdned to simply wait some time before trying again. | | BOZO_NZB | The NZB sent is not a valid NZB file. | | SEARCH_ERROR | There was an error searching using the TorBox Search API. | | INVALID_DEVICE | The client is sending requests from the incorrect device. | | DIFF_ISSUE | The request parameters sent does not allow for this request to complete. | | LINK_OFFLINE | The link given is inaccessible or has no online files. | | VENDOR_DISABLED | This vendor account has been disabled. Please contact support. | | BOZO_REGEX | The regex you entered is bad. | ### Translation to Real-Debrid API Useful for if you want to add TorBox to an app where Real-Debrid is already existing. These are the API routes that are nearly the same with Real-Debrid API routes. You will have to change things such as parameters and check outputs as they both are not the same. You can find docs to Real-Debrid's API [here](https://api.real-debrid.com/). | Real-Debrid API Routes | TorBox API Routes | | --- | --- | | `GET` /torrents | `GET` /torrents/mylist | | `GET` /torrents/info/{id} | `GET` /torrents/mylist?id={id} | | `GET` /torrents/instantAvailability/{hash} | `GET` /torrents/checkcached | | `PUT` /torrents/addTorrent | `POST` /torrents/createtorrent | | `POST` /torrents/addMagnet | `POST` /torrents/createtorrent | | `POST` /torrents/selectFiles/{id} | **None**, \_not needed. Torrents will download all files. This will not be changed.* | | `DELETE` /torrents/delete/{id} | `POST` /torrents/controltorrent | | `POST` /unrestrict/link | `GET` /torrents/requestdl | ### Dates All dates returned from the API are normalized to UTC time. All dates are in the format: `%Y-%m-%dT%H:%M:%SZ` Example: `2024-10-21T20:47:03Z` = The 21st of October, 2024 at 8:47:03 PM UTC. This makes it easy for users and integrations to parse these dates and change them to the user's time zone for better UX. For more information on these formats, you can take a look at this [cheat sheet](https://strftime.org/).

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
  - [Installation](#installation)
- [Authentication](#authentication)
  - [Access Token Authentication](#access-token-authentication)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### Access Token Authentication

The torbox-api API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "torbox-sdk-go/pkg/torboxapi"
    "torbox-sdk-go/pkg/torboxapiconfig"
  )

config := torboxapiconfig.NewConfig()
config.SetBearerToken("YOUR-TOKEN")

sdk := torboxapi.NewTorboxApi(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "torbox-sdk-go/pkg/torboxapi"
    "torbox-sdk-go/pkg/torboxapiconfig"
  )

config := torboxapiconfig.NewConfig()

sdk := torboxapi.NewTorboxApi(config)
sdk.SetBearerToken("YOUR-TOKEN")
```

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                                |
| :---------------------------------------------------------------------------------- |
| [TorrentsService](documentation/services/torrents_service.md)                       |
| [UsenetService](documentation/services/usenet_service.md)                           |
| [WebDownloadsDebridService](documentation/services/web_downloads_debrid_service.md) |
| [GeneralService](documentation/services/general_service.md)                         |
| [NotificationsService](documentation/services/notifications_service.md)             |
| [UserService](documentation/services/user_service.md)                               |
| [RssFeedsService](documentation/services/rss_feeds_service.md)                      |
| [IntegrationsService](documentation/services/integrations_service.md)               |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `TorboxApiResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                        | Description                                 |
| :------- | :-------------------------- | :------------------------------------------ |
| Data     | `T`                         | The body of the API response                |
| Metadata | `TorboxApiResponseMetadata` | Status code and headers returned by the API |

#### `TorboxApiError`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                        | Description                                 |
| :------- | :-------------------------- | :------------------------------------------ |
| Err      | `error`                     | The error that occurred                     |
| Body     | `T`                         | The body of the API response                |
| Metadata | `TorboxApiResponseMetadata` | Status code and headers returned by the API |

#### `TorboxApiResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details> 
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                                           | Description |
| :------------------------------------------------------------------------------------------------------------- | :---------- |
| [CreateTorrentRequest](documentation/models/create_torrent_request.md)                                         |             |
| [CreateTorrentOkResponse](documentation/models/create_torrent_ok_response.md)                                  |             |
| [ControlTorrentOkResponse](documentation/models/control_torrent_ok_response.md)                                |             |
| [ControlQueuedTorrentOkResponse](documentation/models/control_queued_torrent_ok_response.md)                   |             |
| [RequestDownloadLinkOkResponse](documentation/models/request_download_link_ok_response.md)                     |             |
| [GetTorrentListOkResponse](documentation/models/get_torrent_list_ok_response.md)                               |             |
| [GetTorrentCachedAvailabilityOkResponse](documentation/models/get_torrent_cached_availability_ok_response.md)  |             |
| [SearchAllTorrentsFromScraperOkResponse](documentation/models/search_all_torrents_from_scraper_ok_response.md) |             |
| [ExportTorrentDataOkResponse](documentation/models/export_torrent_data_ok_response.md)                         |             |
| [GetTorrentInfoOkResponse](documentation/models/get_torrent_info_ok_response.md)                               |             |
| [CreateUsenetDownloadRequest](documentation/models/create_usenet_download_request.md)                          |             |
| [CreateUsenetDownloadOkResponse](documentation/models/create_usenet_download_ok_response.md)                   |             |
| [GetUsenetListOkResponse](documentation/models/get_usenet_list_ok_response.md)                                 |             |
| [CreateWebDownloadRequest](documentation/models/create_web_download_request.md)                                |             |
| [CreateWebDownloadOkResponse](documentation/models/create_web_download_ok_response.md)                         |             |
| [GetWebDownloadListOkResponse](documentation/models/get_web_download_list_ok_response.md)                      |             |
| [GetUpStatusOkResponse](documentation/models/get_up_status_ok_response.md)                                     |             |
| [GetStatsOkResponse](documentation/models/get_stats_ok_response.md)                                            |             |
| [GetNotificationFeedOkResponse](documentation/models/get_notification_feed_ok_response.md)                     |             |
| [GetUserDataOkResponse](documentation/models/get_user_data_ok_response.md)                                     |             |
| [AddReferralToAccountOkResponse](documentation/models/add_referral_to_account_ok_response.md)                  |             |
| [GetAllJobsOkResponse](documentation/models/get_all_jobs_ok_response.md)                                       |             |
| [GetAllJobsByHashOkResponse](documentation/models/get_all_jobs_by_hash_ok_response.md)                         |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.
