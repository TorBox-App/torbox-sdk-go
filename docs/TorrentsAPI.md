# \TorrentsAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncCreateTorrent**](TorrentsAPI.md#AsyncCreateTorrent) | **Post** /v1/api/torrents/asynccreatetorrent | Async Create Torrent
[**ControlQueuedTorrents**](TorrentsAPI.md#ControlQueuedTorrents) | **Post** /v1/api/torrents/controlqueued | Control Queued Torrents
[**ControlTorrent**](TorrentsAPI.md#ControlTorrent) | **Post** /v1/api/torrents/controltorrent | Control Torrent
[**CreateTorrent**](TorrentsAPI.md#CreateTorrent) | **Post** /v1/api/torrents/createtorrent | Create Torrent
[**EditTorrentItem**](TorrentsAPI.md#EditTorrentItem) | **Put** /v1/api/torrents/edittorrent | Edit Torrent Item
[**ExportTorrentData**](TorrentsAPI.md#ExportTorrentData) | **Get** /v1/api/torrents/exportdata | Export Torrent Data
[**GetQueuedTorrents**](TorrentsAPI.md#GetQueuedTorrents) | **Get** /v1/api/torrents/getqueued | Get Queued Torrents
[**GetTorrentCachedAvailability**](TorrentsAPI.md#GetTorrentCachedAvailability) | **Get** /v1/api/torrents/checkcached | Get Torrent Cached Availability
[**GetTorrentCachedAvailabilityBatch**](TorrentsAPI.md#GetTorrentCachedAvailabilityBatch) | **Post** /v1/api/torrents/checkcached | Get Torrent Cached Availability Batch
[**GetTorrentInfo**](TorrentsAPI.md#GetTorrentInfo) | **Get** /v1/api/torrents/torrentinfo | Get Torrent Info
[**GetTorrentInfoFromFile**](TorrentsAPI.md#GetTorrentInfoFromFile) | **Post** /v1/api/torrents/torrentinfo | Get Torrent Info
[**GetTorrentList**](TorrentsAPI.md#GetTorrentList) | **Get** /v1/api/torrents/mylist | Get Torrent List
[**MagnetToFile**](TorrentsAPI.md#MagnetToFile) | **Post** /v1/api/torrents/magnettofile | Magnet To File
[**RequestTorrentDownloadLink**](TorrentsAPI.md#RequestTorrentDownloadLink) | **Get** /v1/api/torrents/requestdl | Request Download Link
[**RequestTorrentDownloadLinksGet**](TorrentsAPI.md#RequestTorrentDownloadLinksGet) | **Get** /v1/api/torrents/requestdl/batch | Request Download Links - GET
[**RequestTorrentDownloadLinksPost**](TorrentsAPI.md#RequestTorrentDownloadLinksPost) | **Post** /v1/api/torrents/requestdl/batch | Request Download Links - POST



## AsyncCreateTorrent

> AsyncCreateTorrentResponse AsyncCreateTorrent(ctx).File(file).Magnet(magnet).Seed(seed).AllowZip(allowZip).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()

Async Create Torrent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | The torrent's torrent file. Optional. (optional)
	magnet := "magnet_example" // string | The torrent's magnet link. Optional. (optional)
	seed := int64(789) // int64 | Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. (optional)
	allowZip := true // bool | Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger. (optional) (default to true)
	name := "name_example" // string | The name you want the torrent to be. Optional. (optional)
	asQueued := true // bool | Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional. (optional) (default to false)
	addOnlyIfCached := true // bool | Only adds the download if it is cached on TorBox. If not cached, it won't be added. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.AsyncCreateTorrent(context.Background()).File(file).Magnet(magnet).Seed(seed).AllowZip(allowZip).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.AsyncCreateTorrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncCreateTorrent`: AsyncCreateTorrentResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.AsyncCreateTorrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsyncCreateTorrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The torrent&#39;s torrent file. Optional. | 
 **magnet** | **string** | The torrent&#39;s magnet link. Optional. | 
 **seed** | **int64** | Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don&#39;t seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. | 
 **allowZip** | **bool** | Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger. | [default to true]
 **name** | **string** | The name you want the torrent to be. Optional. | 
 **asQueued** | **bool** | Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional. | [default to false]
 **addOnlyIfCached** | **bool** | Only adds the download if it is cached on TorBox. If not cached, it won&#39;t be added. | [default to false]

### Return type

[**AsyncCreateTorrentResponse**](AsyncCreateTorrentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlQueuedTorrents

> ControlQueuedTorrents(ctx).Execute()

Control Queued Torrents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TorrentsAPI.ControlQueuedTorrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.ControlQueuedTorrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiControlQueuedTorrentsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlTorrent

> ControlTorrentResponse ControlTorrent(ctx).ControlTorrent(controlTorrent).Execute()

Control Torrent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	controlTorrent := *openapiclient.NewControlTorrent("Operation_example") // ControlTorrent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.ControlTorrent(context.Background()).ControlTorrent(controlTorrent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.ControlTorrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlTorrent`: ControlTorrentResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.ControlTorrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlTorrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlTorrent** | [**ControlTorrent**](ControlTorrent.md) |  | 

### Return type

[**ControlTorrentResponse**](ControlTorrentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTorrent

> CreateTorrentResponse CreateTorrent(ctx).File(file).Magnet(magnet).Seed(seed).AllowZip(allowZip).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()

Create Torrent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | The torrent's torrent file. Optional. (optional)
	magnet := "magnet_example" // string | The torrent's magnet link. Optional. (optional)
	seed := int64(789) // int64 | Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don't seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. (optional)
	allowZip := true // bool | Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger. (optional) (default to true)
	name := "name_example" // string | The name you want the torrent to be. Optional. (optional)
	asQueued := true // bool | Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional. (optional) (default to false)
	addOnlyIfCached := true // bool | Only adds the download if it is cached on TorBox. If not cached, it won't be added. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.CreateTorrent(context.Background()).File(file).Magnet(magnet).Seed(seed).AllowZip(allowZip).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.CreateTorrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTorrent`: CreateTorrentResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.CreateTorrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTorrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The torrent&#39;s torrent file. Optional. | 
 **magnet** | **string** | The torrent&#39;s magnet link. Optional. | 
 **seed** | **int64** | Tells TorBox your preference for seeding this torrent. 1 is auto. 2 is seed. 3 is don&#39;t seed. Optional. Default is 1, or whatever the user has in their settings. Overwrites option in settings. | 
 **allowZip** | **bool** | Tells TorBox if you want to allow this torrent to be zipped or not. TorBox only zips if the torrent is 100 files or larger. | [default to true]
 **name** | **string** | The name you want the torrent to be. Optional. | 
 **asQueued** | **bool** | Tells TorBox you want this torrent instantly queued. This is bypassed if user is on free plan, and will process the request as normal in this case. Optional. | [default to false]
 **addOnlyIfCached** | **bool** | Only adds the download if it is cached on TorBox. If not cached, it won&#39;t be added. | [default to false]

### Return type

[**CreateTorrentResponse**](CreateTorrentResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTorrentItem

> EditTorrentItemResponse EditTorrentItem(ctx).EditTorrent(editTorrent).Execute()

Edit Torrent Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	editTorrent := *openapiclient.NewEditTorrent(int64(123)) // EditTorrent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.EditTorrentItem(context.Background()).EditTorrent(editTorrent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.EditTorrentItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTorrentItem`: EditTorrentItemResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.EditTorrentItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditTorrentItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editTorrent** | [**EditTorrent**](EditTorrent.md) |  | 

### Return type

[**EditTorrentItemResponse**](EditTorrentItemResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTorrentData

> ExportTorrentDataResponse ExportTorrentData(ctx).TorrentId(torrentId).Type_(type_).Execute()

Export Torrent Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	torrentId := int64(789) // int64 | The torrent's ID.
	type_ := "type__example" // string | Either \"magnet\" or \"file\". Tells how the API what to get, and what to respond as. If magnet, it returns a JSON response with the magnet as a string in the data key. If file, it responds with a torrent file download.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.ExportTorrentData(context.Background()).TorrentId(torrentId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.ExportTorrentData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportTorrentData`: ExportTorrentDataResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.ExportTorrentData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTorrentDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **torrentId** | **int64** | The torrent&#39;s ID. | 
 **type_** | **string** | Either \&quot;magnet\&quot; or \&quot;file\&quot;. Tells how the API what to get, and what to respond as. If magnet, it returns a JSON response with the magnet as a string in the data key. If file, it responds with a torrent file download. | 

### Return type

[**ExportTorrentDataResponse**](ExportTorrentDataResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-bittorrent

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueuedTorrents

> GetQueuedTorrents(ctx).Execute()

Get Queued Torrents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TorrentsAPI.GetQueuedTorrents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.GetQueuedTorrents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueuedTorrentsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTorrentCachedAvailability

> GetTorrentCachedAvailabilityResponse GetTorrentCachedAvailability(ctx).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()

Get Torrent Cached Availability



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	hash := []string{"Inner_example"} // []string | The list of torrent hashes you want to check. Comma seperated. (optional)
	format := "format_example" // string | Format you want the data in. Acceptable is either \"object\" or \"list\". List is the most performant option as it doesn't require modification of the list. (optional) (default to "object")
	listFiles := true // bool | Allows you to list the files found inside the cached data. (optional) (default to false)
	checkCached := *openapiclient.NewCheckCached([]string{"Hashes_example"}) // CheckCached |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.GetTorrentCachedAvailability(context.Background()).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.GetTorrentCachedAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTorrentCachedAvailability`: GetTorrentCachedAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.GetTorrentCachedAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTorrentCachedAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **[]string** | The list of torrent hashes you want to check. Comma seperated. | 
 **format** | **string** | Format you want the data in. Acceptable is either \&quot;object\&quot; or \&quot;list\&quot;. List is the most performant option as it doesn&#39;t require modification of the list. | [default to &quot;object&quot;]
 **listFiles** | **bool** | Allows you to list the files found inside the cached data. | [default to false]
 **checkCached** | [**CheckCached**](CheckCached.md) |  | 

### Return type

[**GetTorrentCachedAvailabilityResponse**](GetTorrentCachedAvailabilityResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTorrentCachedAvailabilityBatch

> GetTorrentCachedAvailabilityBatchResponse GetTorrentCachedAvailabilityBatch(ctx).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()

Get Torrent Cached Availability Batch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	hash := []string{"Inner_example"} // []string |  (optional)
	format := "format_example" // string |  (optional) (default to "object")
	listFiles := true // bool |  (optional) (default to false)
	checkCached := *openapiclient.NewCheckCached([]string{"Hashes_example"}) // CheckCached |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.GetTorrentCachedAvailabilityBatch(context.Background()).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.GetTorrentCachedAvailabilityBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTorrentCachedAvailabilityBatch`: GetTorrentCachedAvailabilityBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.GetTorrentCachedAvailabilityBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTorrentCachedAvailabilityBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **[]string** |  | 
 **format** | **string** |  | [default to &quot;object&quot;]
 **listFiles** | **bool** |  | [default to false]
 **checkCached** | [**CheckCached**](CheckCached.md) |  | 

### Return type

[**GetTorrentCachedAvailabilityBatchResponse**](GetTorrentCachedAvailabilityBatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTorrentInfo

> GetTorrentInfoResponse GetTorrentInfo(ctx).Hash(hash).Timeout(timeout).UseCacheLookup(useCacheLookup).Execute()

Get Torrent Info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	hash := "hash_example" // string | Hash of the torrent you want to get info for. This is required.
	timeout := int64(789) // int64 | The amount of time you want TorBox to search for the torrent on the Bittorrent network. If the number of seeders is low or even zero, this value may be helpful to move up. Default is 10. Optional. (optional) (default to 30)
	useCacheLookup := true // bool | If you want to use a pre-cached request if sent before. This can save a lot of time. Default false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.GetTorrentInfo(context.Background()).Hash(hash).Timeout(timeout).UseCacheLookup(useCacheLookup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.GetTorrentInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTorrentInfo`: GetTorrentInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.GetTorrentInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTorrentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **string** | Hash of the torrent you want to get info for. This is required. | 
 **timeout** | **int64** | The amount of time you want TorBox to search for the torrent on the Bittorrent network. If the number of seeders is low or even zero, this value may be helpful to move up. Default is 10. Optional. | [default to 30]
 **useCacheLookup** | **bool** | If you want to use a pre-cached request if sent before. This can save a lot of time. Default false. | [default to false]

### Return type

[**GetTorrentInfoResponse**](GetTorrentInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTorrentInfoFromFile

> GetTorrentInfoFromFileResponse GetTorrentInfoFromFile(ctx).Timeout(timeout).Hash(hash).Magnet(magnet).File(file).UseCacheLookup(useCacheLookup).PeersOnly(peersOnly).Execute()

Get Torrent Info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	timeout := int64(789) // int64 | The amount of time you want TorBox to search for the torrent on the Bittorrent network. If the number of seeders is low or even zero, this value may be helpful to move up. Default is 10. Optional. (optional) (default to 10)
	hash := "hash_example" // string | A torrent hash for requesting information about the torrent. Optional. (optional)
	magnet := "magnet_example" // string | A magnet link for requesting information abouit the torrent. Optional. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | A torrent file for requesting information about the torrent. Optional. (optional)
	useCacheLookup := true // bool | If you want to use a pre-cached request if sent before. This can save a lot of time. Default false. (optional) (default to false)
	peersOnly := true // bool | Uses a quick lookup method using the trackers contained within the magnet or file. Requires the submitted magnet or torrent file to contain trackers of UDP or HTTP(S) to work. This will not get metadata such as name, size, or files. Best used for simply and quickly gathering the current seeds and peers according to the trackers. Uses the highest number of all the trackers as the final value. Default off. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.GetTorrentInfoFromFile(context.Background()).Timeout(timeout).Hash(hash).Magnet(magnet).File(file).UseCacheLookup(useCacheLookup).PeersOnly(peersOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.GetTorrentInfoFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTorrentInfoFromFile`: GetTorrentInfoFromFileResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.GetTorrentInfoFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTorrentInfoFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeout** | **int64** | The amount of time you want TorBox to search for the torrent on the Bittorrent network. If the number of seeders is low or even zero, this value may be helpful to move up. Default is 10. Optional. | [default to 10]
 **hash** | **string** | A torrent hash for requesting information about the torrent. Optional. | 
 **magnet** | **string** | A magnet link for requesting information abouit the torrent. Optional. | 
 **file** | ***os.File** | A torrent file for requesting information about the torrent. Optional. | 
 **useCacheLookup** | **bool** | If you want to use a pre-cached request if sent before. This can save a lot of time. Default false. | [default to false]
 **peersOnly** | **bool** | Uses a quick lookup method using the trackers contained within the magnet or file. Requires the submitted magnet or torrent file to contain trackers of UDP or HTTP(S) to work. This will not get metadata such as name, size, or files. Best used for simply and quickly gathering the current seeds and peers according to the trackers. Uses the highest number of all the trackers as the final value. Default off. | [default to false]

### Return type

[**GetTorrentInfoFromFileResponse**](GetTorrentInfoFromFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTorrentList

> GetTorrentListResponse GetTorrentList(ctx).Id(id).BypassCache(bypassCache).Offset(offset).Limit(limit).Files(files).Active(active).Execute()

Get Torrent List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	id := int64(789) // int64 | Determines the torrent requested, will return an object rather than list. Optional. (optional)
	bypassCache := true // bool | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. (optional) (default to false)
	offset := int64(789) // int64 | Determines the offset of items to get from the database. Default is 0. Optional. (optional) (default to 0)
	limit := int64(789) // int64 | Determines the number of items to recieve per request. Default is 1000. Optional. (optional) (default to 1000)
	files := true // bool | Exclude files from response. Optional (optional) (default to true)
	active := true // bool | Shows only active downloads. Optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.GetTorrentList(context.Background()).Id(id).BypassCache(bypassCache).Offset(offset).Limit(limit).Files(files).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.GetTorrentList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTorrentList`: GetTorrentListResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.GetTorrentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTorrentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | Determines the torrent requested, will return an object rather than list. Optional. | 
 **bypassCache** | **bool** | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. | [default to false]
 **offset** | **int64** | Determines the offset of items to get from the database. Default is 0. Optional. | [default to 0]
 **limit** | **int64** | Determines the number of items to recieve per request. Default is 1000. Optional. | [default to 1000]
 **files** | **bool** | Exclude files from response. Optional | [default to true]
 **active** | **bool** | Shows only active downloads. Optional | 

### Return type

[**GetTorrentListResponse**](GetTorrentListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MagnetToFile

> string MagnetToFile(ctx).MagnetToTorrent(magnetToTorrent).Execute()

Magnet To File



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	magnetToTorrent := *openapiclient.NewMagnetToTorrent("Magnet_example") // MagnetToTorrent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.MagnetToFile(context.Background()).MagnetToTorrent(magnetToTorrent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.MagnetToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MagnetToFile`: string
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.MagnetToFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMagnetToFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **magnetToTorrent** | [**MagnetToTorrent**](MagnetToTorrent.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-bittorrent, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestTorrentDownloadLink

> RequestTorrentDownloadLinkResponse RequestTorrentDownloadLink(ctx).Token(token).TorrentId(torrentId).FileId(fileId).ZipLink(zipLink).UserIp(userIp).Redirect(redirect).AppendName(appendName).Execute()

Request Download Link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	token := "token_example" // string | Your API Key
	torrentId := int64(789) // int64 | The torrent's ID that you want to download
	fileId := int64(789) // int64 | The files's ID that you want to download (optional) (default to 0)
	zipLink := true // bool | If you want a zip link. Required if no file_id. Takes precedence over file_id if both are given. (optional) (default to false)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	redirect := true // bool | If you want to redirect the user to the CDN link. This is useful for creating permalinks so that you can just make this request URL the link. (optional) (default to false)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.RequestTorrentDownloadLink(context.Background()).Token(token).TorrentId(torrentId).FileId(fileId).ZipLink(zipLink).UserIp(userIp).Redirect(redirect).AppendName(appendName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.RequestTorrentDownloadLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestTorrentDownloadLink`: RequestTorrentDownloadLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.RequestTorrentDownloadLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestTorrentDownloadLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **torrentId** | **int64** | The torrent&#39;s ID that you want to download | 
 **fileId** | **int64** | The files&#39;s ID that you want to download | [default to 0]
 **zipLink** | **bool** | If you want a zip link. Required if no file_id. Takes precedence over file_id if both are given. | [default to false]
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **redirect** | **bool** | If you want to redirect the user to the CDN link. This is useful for creating permalinks so that you can just make this request URL the link. | [default to false]
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]

### Return type

[**RequestTorrentDownloadLinkResponse**](RequestTorrentDownloadLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestTorrentDownloadLinksGet

> RequestTorrentDownloadLinksGetResponse RequestTorrentDownloadLinksGet(ctx).Token(token).TorrentId(torrentId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()

Request Download Links - GET



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	token := "token_example" // string | Your API Key
	torrentId := int64(789) // int64 | The torrent's ID that you want to download
	fileIds := []int64{int64(123)} // []int64 | The files's ID that you want to download (optional)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)
	requestDownloadBatch := *openapiclient.NewRequestDownloadBatch() // RequestDownloadBatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.RequestTorrentDownloadLinksGet(context.Background()).Token(token).TorrentId(torrentId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.RequestTorrentDownloadLinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestTorrentDownloadLinksGet`: RequestTorrentDownloadLinksGetResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.RequestTorrentDownloadLinksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestTorrentDownloadLinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **torrentId** | **int64** | The torrent&#39;s ID that you want to download | 
 **fileIds** | **[]int64** | The files&#39;s ID that you want to download | 
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]
 **requestDownloadBatch** | [**RequestDownloadBatch**](RequestDownloadBatch.md) |  | 

### Return type

[**RequestTorrentDownloadLinksGetResponse**](RequestTorrentDownloadLinksGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestTorrentDownloadLinksPost

> RequestTorrentDownloadLinksPostResponse RequestTorrentDownloadLinksPost(ctx).Token(token).TorrentId(torrentId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()

Request Download Links - POST



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	token := "token_example" // string | Your API Key
	torrentId := int64(789) // int64 | The torrent's ID that you want to download
	fileIds := []int64{int64(123)} // []int64 |  (optional)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)
	requestDownloadBatch := *openapiclient.NewRequestDownloadBatch() // RequestDownloadBatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TorrentsAPI.RequestTorrentDownloadLinksPost(context.Background()).Token(token).TorrentId(torrentId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TorrentsAPI.RequestTorrentDownloadLinksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestTorrentDownloadLinksPost`: RequestTorrentDownloadLinksPostResponse
	fmt.Fprintf(os.Stdout, "Response from `TorrentsAPI.RequestTorrentDownloadLinksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestTorrentDownloadLinksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **torrentId** | **int64** | The torrent&#39;s ID that you want to download | 
 **fileIds** | **[]int64** |  | 
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]
 **requestDownloadBatch** | [**RequestDownloadBatch**](RequestDownloadBatch.md) |  | 

### Return type

[**RequestTorrentDownloadLinksPostResponse**](RequestTorrentDownloadLinksPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

