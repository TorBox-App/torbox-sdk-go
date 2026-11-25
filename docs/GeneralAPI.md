# \GeneralAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get30DayStats**](GeneralAPI.md#Get30DayStats) | **Get** /v1/api/stats/30days | Get 30 Day Stats
[**GetChangelogsJson**](GeneralAPI.md#GetChangelogsJson) | **Get** /v1/api/changelogs/json | Get Changelogs JSON
[**GetChangelogsRssFeed**](GeneralAPI.md#GetChangelogsRssFeed) | **Get** /v1/api/changelogs/rss | Get Changelogs RSS Feed
[**GetSpeedtestFiles**](GeneralAPI.md#GetSpeedtestFiles) | **Get** /v1/api/speedtest | Get Speedtest Files
[**GetStats**](GeneralAPI.md#GetStats) | **Get** /v1/api/stats | Get Stats
[**GetUpStatus**](GeneralAPI.md#GetUpStatus) | **Get** / | Get Up Status



## Get30DayStats

> Get30DayStatsResponse Get30DayStats(ctx).Execute()

Get 30 Day Stats



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
	resp, r, err := apiClient.GeneralAPI.Get30DayStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.Get30DayStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get30DayStats`: Get30DayStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.Get30DayStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGet30DayStatsRequest struct via the builder pattern


### Return type

[**Get30DayStatsResponse**](Get30DayStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangelogsJson

> GetChangelogsJsonResponse GetChangelogsJson(ctx).Offset(offset).Limit(limit).Execute()

Get Changelogs JSON



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
	offset := int64(789) // int64 |  (optional) (default to 0)
	limit := int64(789) // int64 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeneralAPI.GetChangelogsJson(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetChangelogsJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangelogsJson`: GetChangelogsJsonResponse
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetChangelogsJson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangelogsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | [default to 0]
 **limit** | **int64** |  | [default to 10]

### Return type

[**GetChangelogsJsonResponse**](GetChangelogsJsonResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangelogsRssFeed

> string GetChangelogsRssFeed(ctx).Offset(offset).Limit(limit).Execute()

Get Changelogs RSS Feed



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
	offset := int64(789) // int64 |  (optional) (default to 0)
	limit := int64(789) // int64 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeneralAPI.GetChangelogsRssFeed(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetChangelogsRssFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangelogsRssFeed`: string
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetChangelogsRssFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangelogsRssFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** |  | [default to 0]
 **limit** | **int64** |  | [default to 10]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpeedtestFiles

> GetSpeedtestFilesResponse GetSpeedtestFiles(ctx).UserIp(userIp).Region(region).TestLength(testLength).Execute()

Get Speedtest Files



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
	userIp := "userIp_example" // string | Pass a user's IP to get the closest server to them, otherwise will use the calling IP for location. (optional)
	region := "region_example" // string | Determines what cdns are returned. May be any region that TorBox is located in. To get this value, send a request without this value to determine all of the available regions that are available. (optional) (default to "all")
	testLength := "testLength_example" // string | Determines the size of the file used for the speedtest. May be \"long\" or \"short\". Optional. (optional) (default to "short")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeneralAPI.GetSpeedtestFiles(context.Background()).UserIp(userIp).Region(region).TestLength(testLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetSpeedtestFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpeedtestFiles`: GetSpeedtestFilesResponse
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetSpeedtestFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpeedtestFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIp** | **string** | Pass a user&#39;s IP to get the closest server to them, otherwise will use the calling IP for location. | 
 **region** | **string** | Determines what cdns are returned. May be any region that TorBox is located in. To get this value, send a request without this value to determine all of the available regions that are available. | [default to &quot;all&quot;]
 **testLength** | **string** | Determines the size of the file used for the speedtest. May be \&quot;long\&quot; or \&quot;short\&quot;. Optional. | [default to &quot;short&quot;]

### Return type

[**GetSpeedtestFilesResponse**](GetSpeedtestFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStats

> GetStatsResponse GetStats(ctx).Execute()

Get Stats



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
	resp, r, err := apiClient.GeneralAPI.GetStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStats`: GetStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


### Return type

[**GetStatsResponse**](GetStatsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpStatus

> GetUpStatusResponse GetUpStatus(ctx).Execute()

Get Up Status



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
	resp, r, err := apiClient.GeneralAPI.GetUpStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.GetUpStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpStatus`: GetUpStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.GetUpStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpStatusRequest struct via the builder pattern


### Return type

[**GetUpStatusResponse**](GetUpStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

