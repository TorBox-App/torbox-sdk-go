# \NotificationsAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearAllNotifications**](NotificationsAPI.md#ClearAllNotifications) | **Post** /v1/api/notifications/clear | Clear All Notifications
[**ClearSingleNotification**](NotificationsAPI.md#ClearSingleNotification) | **Post** /v1/api/notifications/clear/{id} | Clear Single Notification
[**GetNotificationFeed**](NotificationsAPI.md#GetNotificationFeed) | **Get** /v1/api/notifications/mynotifications | Get Notification Feed
[**GetRssNotificationFeed**](NotificationsAPI.md#GetRssNotificationFeed) | **Get** /v1/api/notifications/rss | Get RSS Notification Feed
[**SendTestNotification**](NotificationsAPI.md#SendTestNotification) | **Post** /v1/api/notifications/test | Send Test Notification



## ClearAllNotifications

> ClearAllNotificationsResponse ClearAllNotifications(ctx).Execute()

Clear All Notifications



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
	resp, r, err := apiClient.NotificationsAPI.ClearAllNotifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ClearAllNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearAllNotifications`: ClearAllNotificationsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ClearAllNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearAllNotificationsRequest struct via the builder pattern


### Return type

[**ClearAllNotificationsResponse**](ClearAllNotificationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearSingleNotification

> ClearSingleNotificationResponse ClearSingleNotification(ctx, id).Execute()

Clear Single Notification



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ClearSingleNotification(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ClearSingleNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearSingleNotification`: ClearSingleNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ClearSingleNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSingleNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClearSingleNotificationResponse**](ClearSingleNotificationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationFeed

> GetNotificationFeedResponse GetNotificationFeed(ctx).BypassCache(bypassCache).Execute()

Get Notification Feed



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
	bypassCache := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationFeed(context.Background()).BypassCache(bypassCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationFeed`: GetNotificationFeedResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bypassCache** | **bool** |  | [default to false]

### Return type

[**GetNotificationFeedResponse**](GetNotificationFeedResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRssNotificationFeed

> string GetRssNotificationFeed(ctx).Token(token).Execute()

Get RSS Notification Feed



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
	token := "token_example" // string | The user's API key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetRssNotificationFeed(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetRssNotificationFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRssNotificationFeed`: string
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetRssNotificationFeed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRssNotificationFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The user&#39;s API key. | 

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


## SendTestNotification

> SendTestNotificationResponse SendTestNotification(ctx).Execute()

Send Test Notification



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
	resp, r, err := apiClient.NotificationsAPI.SendTestNotification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.SendTestNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestNotification`: SendTestNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.SendTestNotification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestNotificationRequest struct via the builder pattern


### Return type

[**SendTestNotificationResponse**](SendTestNotificationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

