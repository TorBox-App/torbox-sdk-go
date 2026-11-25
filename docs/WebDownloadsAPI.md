# \WebDownloadsAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncCreateWebDownload**](WebDownloadsAPI.md#AsyncCreateWebDownload) | **Post** /v1/api/webdl/asynccreatewebdownload | Async Create Web Download
[**ControlWebDownload**](WebDownloadsAPI.md#ControlWebDownload) | **Post** /v1/api/webdl/controlwebdownload | Control Web Download
[**CreateWebDownload**](WebDownloadsAPI.md#CreateWebDownload) | **Post** /v1/api/webdl/createwebdownload | Create Web Download
[**EditWebDownloadItem**](WebDownloadsAPI.md#EditWebDownloadItem) | **Put** /v1/api/webdl/editwebdownload | Edit Web Download Item
[**GetHosterList**](WebDownloadsAPI.md#GetHosterList) | **Get** /v1/api/webdl/hosters | Get Hoster List
[**GetWebDownloadCachedAvailability**](WebDownloadsAPI.md#GetWebDownloadCachedAvailability) | **Get** /v1/api/webdl/checkcached | Get Web Download Cached Availability
[**GetWebDownloadCachedAvailabilityBatch**](WebDownloadsAPI.md#GetWebDownloadCachedAvailabilityBatch) | **Post** /v1/api/webdl/checkcached | Get Web Download Cached Availability Batch
[**GetWebDownloadList**](WebDownloadsAPI.md#GetWebDownloadList) | **Get** /v1/api/webdl/mylist | Get Web Download List
[**RequestWebDownloadLink**](WebDownloadsAPI.md#RequestWebDownloadLink) | **Get** /v1/api/webdl/requestdl | Request Download Link
[**RequestWebDownloadLinksGet**](WebDownloadsAPI.md#RequestWebDownloadLinksGet) | **Get** /v1/api/webdl/requestdl/batch | Request Download Links - GET
[**RequestWebDownloadLinksPost**](WebDownloadsAPI.md#RequestWebDownloadLinksPost) | **Post** /v1/api/webdl/requestdl/batch | Request Download Links - POST



## AsyncCreateWebDownload

> AsyncCreateWebDownloadResponse AsyncCreateWebDownload(ctx).Link(link).Password(password).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()

Async Create Web Download



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
	link := "link_example" // string | An accessible link to any file on the internet. Cannot be a redirection.
	password := "password_example" // string | Password to the file, if the files requires it. Optional. (optional)
	name := "name_example" // string | The name you want the web download to be. Optional. (optional)
	asQueued := true // bool | Tells TorBox you want this web download instantly queued. Optional. (optional) (default to false)
	addOnlyIfCached := true // bool | Only adds the download if it is cached on TorBox. If not cached, it won't be added. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.AsyncCreateWebDownload(context.Background()).Link(link).Password(password).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.AsyncCreateWebDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncCreateWebDownload`: AsyncCreateWebDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.AsyncCreateWebDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsyncCreateWebDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** | An accessible link to any file on the internet. Cannot be a redirection. | 
 **password** | **string** | Password to the file, if the files requires it. Optional. | 
 **name** | **string** | The name you want the web download to be. Optional. | 
 **asQueued** | **bool** | Tells TorBox you want this web download instantly queued. Optional. | [default to false]
 **addOnlyIfCached** | **bool** | Only adds the download if it is cached on TorBox. If not cached, it won&#39;t be added. | [default to false]

### Return type

[**AsyncCreateWebDownloadResponse**](AsyncCreateWebDownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlWebDownload

> ControlWebDownloadResponse ControlWebDownload(ctx).ControlWebDownload(controlWebDownload).Execute()

Control Web Download



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
	controlWebDownload := *openapiclient.NewControlWebDownload("Operation_example") // ControlWebDownload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.ControlWebDownload(context.Background()).ControlWebDownload(controlWebDownload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.ControlWebDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlWebDownload`: ControlWebDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.ControlWebDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlWebDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlWebDownload** | [**ControlWebDownload**](ControlWebDownload.md) |  | 

### Return type

[**ControlWebDownloadResponse**](ControlWebDownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebDownload

> CreateWebDownloadResponse CreateWebDownload(ctx).Link(link).Password(password).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()

Create Web Download



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
	link := "link_example" // string | An accessible link to any file on the internet. Cannot be a redirection.
	password := "password_example" // string | Password to the file, if the files requires it. Optional. (optional)
	name := "name_example" // string | The name you want the torrent to be. Optional. (optional)
	asQueued := true // bool | Tells TorBox you want this web download.instantly queued. Optional. (optional) (default to false)
	addOnlyIfCached := true // bool | Only adds the download if it is cached on TorBox. If not cached, it won't be added. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.CreateWebDownload(context.Background()).Link(link).Password(password).Name(name).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.CreateWebDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebDownload`: CreateWebDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.CreateWebDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** | An accessible link to any file on the internet. Cannot be a redirection. | 
 **password** | **string** | Password to the file, if the files requires it. Optional. | 
 **name** | **string** | The name you want the torrent to be. Optional. | 
 **asQueued** | **bool** | Tells TorBox you want this web download.instantly queued. Optional. | [default to false]
 **addOnlyIfCached** | **bool** | Only adds the download if it is cached on TorBox. If not cached, it won&#39;t be added. | [default to false]

### Return type

[**CreateWebDownloadResponse**](CreateWebDownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWebDownloadItem

> EditWebDownloadItemResponse EditWebDownloadItem(ctx).EditWebDownload(editWebDownload).Execute()

Edit Web Download Item



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
	editWebDownload := *openapiclient.NewEditWebDownload(int64(123)) // EditWebDownload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.EditWebDownloadItem(context.Background()).EditWebDownload(editWebDownload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.EditWebDownloadItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditWebDownloadItem`: EditWebDownloadItemResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.EditWebDownloadItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditWebDownloadItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editWebDownload** | [**EditWebDownload**](EditWebDownload.md) |  | 

### Return type

[**EditWebDownloadItemResponse**](EditWebDownloadItemResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHosterList

> GetHosterListResponse GetHosterList(ctx).Execute()

Get Hoster List



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
	resp, r, err := apiClient.WebDownloadsAPI.GetHosterList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.GetHosterList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHosterList`: GetHosterListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.GetHosterList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHosterListRequest struct via the builder pattern


### Return type

[**GetHosterListResponse**](GetHosterListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebDownloadCachedAvailability

> GetWebDownloadCachedAvailabilityResponse GetWebDownloadCachedAvailability(ctx).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()

Get Web Download Cached Availability



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
	hash := []string{"Inner_example"} // []string | The list of web hashes you want to check. Comma seperated. To find the hash, md5 the link. (optional)
	format := "format_example" // string | Format you want the data in. Acceptable is either \"object\" or \"list\". List is the most performant option as it doesn't require modification of the list. (optional) (default to "object")
	listFiles := true // bool | Allows you to list the files found inside the cached data. (optional) (default to false)
	checkCached := *openapiclient.NewCheckCached([]string{"Hashes_example"}) // CheckCached |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.GetWebDownloadCachedAvailability(context.Background()).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.GetWebDownloadCachedAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebDownloadCachedAvailability`: GetWebDownloadCachedAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.GetWebDownloadCachedAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebDownloadCachedAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **[]string** | The list of web hashes you want to check. Comma seperated. To find the hash, md5 the link. | 
 **format** | **string** | Format you want the data in. Acceptable is either \&quot;object\&quot; or \&quot;list\&quot;. List is the most performant option as it doesn&#39;t require modification of the list. | [default to &quot;object&quot;]
 **listFiles** | **bool** | Allows you to list the files found inside the cached data. | [default to false]
 **checkCached** | [**CheckCached**](CheckCached.md) |  | 

### Return type

[**GetWebDownloadCachedAvailabilityResponse**](GetWebDownloadCachedAvailabilityResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebDownloadCachedAvailabilityBatch

> GetWebDownloadCachedAvailabilityBatchResponse GetWebDownloadCachedAvailabilityBatch(ctx).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()

Get Web Download Cached Availability Batch



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
	format := "format_example" // string | Format you want the data in. Acceptable is either \"object\" or \"list\". List is the most performant option as it doesn't require modification of the list. (optional) (default to "object")
	listFiles := true // bool | Allows you to list the files found inside the cached data. (optional) (default to false)
	checkCached := *openapiclient.NewCheckCached([]string{"Hashes_example"}) // CheckCached |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.GetWebDownloadCachedAvailabilityBatch(context.Background()).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.GetWebDownloadCachedAvailabilityBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebDownloadCachedAvailabilityBatch`: GetWebDownloadCachedAvailabilityBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.GetWebDownloadCachedAvailabilityBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebDownloadCachedAvailabilityBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **[]string** |  | 
 **format** | **string** | Format you want the data in. Acceptable is either \&quot;object\&quot; or \&quot;list\&quot;. List is the most performant option as it doesn&#39;t require modification of the list. | [default to &quot;object&quot;]
 **listFiles** | **bool** | Allows you to list the files found inside the cached data. | [default to false]
 **checkCached** | [**CheckCached**](CheckCached.md) |  | 

### Return type

[**GetWebDownloadCachedAvailabilityBatchResponse**](GetWebDownloadCachedAvailabilityBatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebDownloadList

> GetWebDownloadListResponse GetWebDownloadList(ctx).Id(id).BypassCache(bypassCache).Offset(offset).Limit(limit).Files(files).Active(active).Execute()

Get Web Download List



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
	resp, r, err := apiClient.WebDownloadsAPI.GetWebDownloadList(context.Background()).Id(id).BypassCache(bypassCache).Offset(offset).Limit(limit).Files(files).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.GetWebDownloadList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebDownloadList`: GetWebDownloadListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.GetWebDownloadList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebDownloadListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | Determines the torrent requested, will return an object rather than list. Optional. | 
 **bypassCache** | **bool** | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. | [default to false]
 **offset** | **int64** | Determines the offset of items to get from the database. Default is 0. Optional. | [default to 0]
 **limit** | **int64** | Determines the number of items to recieve per request. Default is 1000. Optional. | [default to 1000]
 **files** | **bool** | Exclude files from response. Optional | [default to true]
 **active** | **bool** | Shows only active downloads. Optional | 

### Return type

[**GetWebDownloadListResponse**](GetWebDownloadListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestWebDownloadLink

> RequestWebDownloadLinkResponse RequestWebDownloadLink(ctx).Token(token).WebId(webId).FileId(fileId).ZipLink(zipLink).UserIp(userIp).Redirect(redirect).AppendName(appendName).Execute()

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
	webId := int64(789) // int64 | The web download's ID that you want to download
	fileId := int64(789) // int64 | The files's ID that you want to download (optional) (default to 0)
	zipLink := true // bool | If you want a zip link. Required if no file_id. Takes precedence over file_id if both are given. (optional) (default to false)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	redirect := true // bool | If you want to redirect the user to the CDN link. This is useful for creating permalinks so that you can just make this request URL the link. (optional) (default to false)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.RequestWebDownloadLink(context.Background()).Token(token).WebId(webId).FileId(fileId).ZipLink(zipLink).UserIp(userIp).Redirect(redirect).AppendName(appendName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.RequestWebDownloadLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestWebDownloadLink`: RequestWebDownloadLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.RequestWebDownloadLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestWebDownloadLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **webId** | **int64** | The web download&#39;s ID that you want to download | 
 **fileId** | **int64** | The files&#39;s ID that you want to download | [default to 0]
 **zipLink** | **bool** | If you want a zip link. Required if no file_id. Takes precedence over file_id if both are given. | [default to false]
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **redirect** | **bool** | If you want to redirect the user to the CDN link. This is useful for creating permalinks so that you can just make this request URL the link. | [default to false]
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]

### Return type

[**RequestWebDownloadLinkResponse**](RequestWebDownloadLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestWebDownloadLinksGet

> RequestWebDownloadLinksGetResponse RequestWebDownloadLinksGet(ctx).Token(token).WebId(webId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()

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
	webId := int64(789) // int64 | The web download's ID that you want to download
	fileIds := []int64{int64(123)} // []int64 | The files's ID that you want to download (optional)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)
	requestDownloadBatch := *openapiclient.NewRequestDownloadBatch() // RequestDownloadBatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.RequestWebDownloadLinksGet(context.Background()).Token(token).WebId(webId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.RequestWebDownloadLinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestWebDownloadLinksGet`: RequestWebDownloadLinksGetResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.RequestWebDownloadLinksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestWebDownloadLinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **webId** | **int64** | The web download&#39;s ID that you want to download | 
 **fileIds** | **[]int64** | The files&#39;s ID that you want to download | 
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]
 **requestDownloadBatch** | [**RequestDownloadBatch**](RequestDownloadBatch.md) |  | 

### Return type

[**RequestWebDownloadLinksGetResponse**](RequestWebDownloadLinksGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestWebDownloadLinksPost

> RequestWebDownloadLinksPostResponse RequestWebDownloadLinksPost(ctx).Token(token).WebId(webId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()

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
	webId := int64(789) // int64 | The web downloads' ID that you want to download
	fileIds := []int64{int64(123)} // []int64 |  (optional)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)
	requestDownloadBatch := *openapiclient.NewRequestDownloadBatch() // RequestDownloadBatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebDownloadsAPI.RequestWebDownloadLinksPost(context.Background()).Token(token).WebId(webId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebDownloadsAPI.RequestWebDownloadLinksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestWebDownloadLinksPost`: RequestWebDownloadLinksPostResponse
	fmt.Fprintf(os.Stdout, "Response from `WebDownloadsAPI.RequestWebDownloadLinksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestWebDownloadLinksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **webId** | **int64** | The web downloads&#39; ID that you want to download | 
 **fileIds** | **[]int64** |  | 
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]
 **requestDownloadBatch** | [**RequestDownloadBatch**](RequestDownloadBatch.md) |  | 

### Return type

[**RequestWebDownloadLinksPostResponse**](RequestWebDownloadLinksPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

