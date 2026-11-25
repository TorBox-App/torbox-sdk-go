# \RssAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRssFeed**](RssAPI.md#AddRssFeed) | **Post** /v1/api/rss/addrss | Add RSS Feed
[**ControlRssFeed**](RssAPI.md#ControlRssFeed) | **Post** /v1/api/rss/controlrss | Control RSS Feed
[**EditRssFeed**](RssAPI.md#EditRssFeed) | **Post** /v1/api/rss/modifyrss | Edit RSS Feed
[**GetRssFeedItems**](RssAPI.md#GetRssFeedItems) | **Get** /v1/api/rss/getfeeditems | Get RSS Feed Items
[**GetUserRssFeeds**](RssAPI.md#GetUserRssFeeds) | **Get** /v1/api/rss/getfeeds | Get User RSS Feeds



## AddRssFeed

> AddRssFeedResponse AddRssFeed(ctx).AddRss(addRss).Execute()

Add RSS Feed



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
	addRss := *openapiclient.NewAddRss("Url_example", "Name_example") // AddRss | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RssAPI.AddRssFeed(context.Background()).AddRss(addRss).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RssAPI.AddRssFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRssFeed`: AddRssFeedResponse
	fmt.Fprintf(os.Stdout, "Response from `RssAPI.AddRssFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRssFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addRss** | [**AddRss**](AddRss.md) |  | 

### Return type

[**AddRssFeedResponse**](AddRssFeedResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlRssFeed

> ControlRssFeedResponse ControlRssFeed(ctx).ControlRss(controlRss).Execute()

Control RSS Feed



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
	controlRss := *openapiclient.NewControlRss("Operation_example", int64(123)) // ControlRss | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RssAPI.ControlRssFeed(context.Background()).ControlRss(controlRss).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RssAPI.ControlRssFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlRssFeed`: ControlRssFeedResponse
	fmt.Fprintf(os.Stdout, "Response from `RssAPI.ControlRssFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlRssFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlRss** | [**ControlRss**](ControlRss.md) |  | 

### Return type

[**ControlRssFeedResponse**](ControlRssFeedResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRssFeed

> EditRssFeedResponse EditRssFeed(ctx).ModifyRss(modifyRss).Execute()

Edit RSS Feed



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
	modifyRss := *openapiclient.NewModifyRss(int64(123)) // ModifyRss | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RssAPI.EditRssFeed(context.Background()).ModifyRss(modifyRss).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RssAPI.EditRssFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRssFeed`: EditRssFeedResponse
	fmt.Fprintf(os.Stdout, "Response from `RssAPI.EditRssFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditRssFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyRss** | [**ModifyRss**](ModifyRss.md) |  | 

### Return type

[**EditRssFeedResponse**](EditRssFeedResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRssFeedItems

> GetRssFeedItemsResponse GetRssFeedItems(ctx).RssFeedId(rssFeedId).Execute()

Get RSS Feed Items



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
	rssFeedId := int64(789) // int64 | The RSS Feed ID of the RSS Feed you want to retrieve the scraped items of.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RssAPI.GetRssFeedItems(context.Background()).RssFeedId(rssFeedId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RssAPI.GetRssFeedItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRssFeedItems`: GetRssFeedItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `RssAPI.GetRssFeedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRssFeedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rssFeedId** | **int64** | The RSS Feed ID of the RSS Feed you want to retrieve the scraped items of. | 

### Return type

[**GetRssFeedItemsResponse**](GetRssFeedItemsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRssFeeds

> GetUserRssFeedsResponse GetUserRssFeeds(ctx).Id(id).Execute()

Get User RSS Feeds



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
	id := int64(789) // int64 | A specific RSS feed ID that you want to retrieve. Optional. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RssAPI.GetUserRssFeeds(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RssAPI.GetUserRssFeeds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRssFeeds`: GetUserRssFeedsResponse
	fmt.Fprintf(os.Stdout, "Response from `RssAPI.GetUserRssFeeds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRssFeedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | A specific RSS feed ID that you want to retrieve. Optional. | 

### Return type

[**GetUserRssFeedsResponse**](GetUserRssFeedsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

