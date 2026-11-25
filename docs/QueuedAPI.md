# \QueuedAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlQueuedDownloads**](QueuedAPI.md#ControlQueuedDownloads) | **Post** /v1/api/queued/controlqueued | Control Queued Downloads
[**GetQueuedDownloads**](QueuedAPI.md#GetQueuedDownloads) | **Get** /v1/api/queued/getqueued | Get Queued Downloads



## ControlQueuedDownloads

> ControlQueuedDownloadsResponse ControlQueuedDownloads(ctx).ControlQueuedDownload(controlQueuedDownload).Execute()

Control Queued Downloads



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
	controlQueuedDownload := *openapiclient.NewControlQueuedDownload("Operation_example") // ControlQueuedDownload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuedAPI.ControlQueuedDownloads(context.Background()).ControlQueuedDownload(controlQueuedDownload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuedAPI.ControlQueuedDownloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlQueuedDownloads`: ControlQueuedDownloadsResponse
	fmt.Fprintf(os.Stdout, "Response from `QueuedAPI.ControlQueuedDownloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlQueuedDownloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlQueuedDownload** | [**ControlQueuedDownload**](ControlQueuedDownload.md) |  | 

### Return type

[**ControlQueuedDownloadsResponse**](ControlQueuedDownloadsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueuedDownloads

> GetQueuedDownloadsResponse GetQueuedDownloads(ctx).Id(id).Type_(type_).BypassCache(bypassCache).Offset(offset).Limit(limit).Execute()

Get Queued Downloads



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
	id := int64(789) // int64 | Determines the queued download requested, will return an object rather than list. Optional. (optional)
	type_ := "type__example" // string | The type of the queued download you want to retrieve. Can be \"torrent\", \"usenet\" or \"webdl\". Optional. Default is \"torrent\". (optional) (default to "torrent")
	bypassCache := true // bool | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. (optional) (default to false)
	offset := int64(789) // int64 | Determines the offset of items to get from the database. Default is 0. Optional. (optional) (default to 0)
	limit := int64(789) // int64 | Determines the number of items to recieve per request. Default is 1000. Optional. (optional) (default to 1000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuedAPI.GetQueuedDownloads(context.Background()).Id(id).Type_(type_).BypassCache(bypassCache).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuedAPI.GetQueuedDownloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueuedDownloads`: GetQueuedDownloadsResponse
	fmt.Fprintf(os.Stdout, "Response from `QueuedAPI.GetQueuedDownloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueuedDownloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | Determines the queued download requested, will return an object rather than list. Optional. | 
 **type_** | **string** | The type of the queued download you want to retrieve. Can be \&quot;torrent\&quot;, \&quot;usenet\&quot; or \&quot;webdl\&quot;. Optional. Default is \&quot;torrent\&quot;. | [default to &quot;torrent&quot;]
 **bypassCache** | **bool** | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. | [default to false]
 **offset** | **int64** | Determines the offset of items to get from the database. Default is 0. Optional. | [default to 0]
 **limit** | **int64** | Determines the number of items to recieve per request. Default is 1000. Optional. | [default to 1000]

### Return type

[**GetQueuedDownloadsResponse**](GetQueuedDownloadsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

