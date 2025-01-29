![Logo](https://raw.githubusercontent.com/TorBox-App/torbox-sdk-go/main/assets/banner.png)

## Table of Contents

- [Setup \& Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
  - [Authentication](#authentication)
    - [Access Token Authentication](#access-token-authentication)
      - [Setting the Access Token](#setting-the-access-token)
  - [Services](#services)
    - [Response Wrappers](#response-wrappers)
      - [`TorboxApiResponse[T]`](#torboxapiresponset)
      - [`TorboxApiError`](#torboxapierror)
      - [`TorboxApiResponseMetadata`](#torboxapiresponsemetadata)
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
| [QueuedService](documentation/services/queued_service.md)                           |

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

| Name                                                                                                          | Description |
| :------------------------------------------------------------------------------------------------------------ | :---------- |
| [CreateTorrentRequest](documentation/models/create_torrent_request.md)                                        |             |
| [CreateTorrentOkResponse](documentation/models/create_torrent_ok_response.md)                                 |             |
| [ControlTorrentOkResponse](documentation/models/control_torrent_ok_response.md)                               |             |
| [RequestDownloadLinkOkResponse](documentation/models/request_download_link_ok_response.md)                    |             |
| [GetTorrentListOkResponse](documentation/models/get_torrent_list_ok_response.md)                              |             |
| [GetTorrentCachedAvailabilityOkResponse](documentation/models/get_torrent_cached_availability_ok_response.md) |             |
| [ExportTorrentDataOkResponse](documentation/models/export_torrent_data_ok_response.md)                        |             |
| [GetTorrentInfoOkResponse](documentation/models/get_torrent_info_ok_response.md)                              |             |
| [CreateUsenetDownloadRequest](documentation/models/create_usenet_download_request.md)                         |             |
| [CreateUsenetDownloadOkResponse](documentation/models/create_usenet_download_ok_response.md)                  |             |
| [GetUsenetListOkResponse](documentation/models/get_usenet_list_ok_response.md)                                |             |
| [CreateWebDownloadRequest](documentation/models/create_web_download_request.md)                               |             |
| [CreateWebDownloadOkResponse](documentation/models/create_web_download_ok_response.md)                        |             |
| [GetWebDownloadListOkResponse](documentation/models/get_web_download_list_ok_response.md)                     |             |
| [GetHosterListOkResponse](documentation/models/get_hoster_list_ok_response.md)                                |             |
| [GetUpStatusOkResponse](documentation/models/get_up_status_ok_response.md)                                    |             |
| [GetStatsOkResponse](documentation/models/get_stats_ok_response.md)                                           |             |
| [GetNotificationFeedOkResponse](documentation/models/get_notification_feed_ok_response.md)                    |             |
| [GetUserDataOkResponse](documentation/models/get_user_data_ok_response.md)                                    |             |
| [AddReferralToAccountOkResponse](documentation/models/add_referral_to_account_ok_response.md)                 |             |
| [GetAllJobsOkResponse](documentation/models/get_all_jobs_ok_response.md)                                      |             |
| [GetAllJobsByHashOkResponse](documentation/models/get_all_jobs_by_hash_ok_response.md)                        |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.
