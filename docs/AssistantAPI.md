# \AssistantAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatWithAssistant**](AssistantAPI.md#ChatWithAssistant) | **Post** /v1/api/assistant/chat | Chat With Assistant
[**DeleteAssistantConversation**](AssistantAPI.md#DeleteAssistantConversation) | **Delete** /v1/api/assistant/conversations/{session_id} | Delete Assistant Conversation
[**DeleteAssistantConversations**](AssistantAPI.md#DeleteAssistantConversations) | **Delete** /v1/api/assistant/conversations | Delete All Assistant Conversations
[**GetAssistantUsage**](AssistantAPI.md#GetAssistantUsage) | **Get** /v1/api/assistant/usage | Get Assistant Usage
[**ListAssistantConversations**](AssistantAPI.md#ListAssistantConversations) | **Get** /v1/api/assistant/conversations | Get Assistant Conversations
[**ReadAssistantConversation**](AssistantAPI.md#ReadAssistantConversation) | **Get** /v1/api/assistant/conversations/{session_id} | Get Assistant Conversation



## ChatWithAssistant

> string ChatWithAssistant(ctx).ChatRequest(chatRequest).Execute()

Chat With Assistant



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
	chatRequest := *openapiclient.NewChatRequest("Message_example") // ChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantAPI.ChatWithAssistant(context.Background()).ChatRequest(chatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantAPI.ChatWithAssistant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatWithAssistant`: string
	fmt.Fprintf(os.Stdout, "Response from `AssistantAPI.ChatWithAssistant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatWithAssistantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatRequest** | [**ChatRequest**](ChatRequest.md) |  | 

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistantConversation

> DeleteAssistantConversationResponse DeleteAssistantConversation(ctx, sessionId).Execute()

Delete Assistant Conversation



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
	sessionId := "sessionId_example" // string | The session you want to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantAPI.DeleteAssistantConversation(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantAPI.DeleteAssistantConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssistantConversation`: DeleteAssistantConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantAPI.DeleteAssistantConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssistantConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAssistantConversationResponse**](DeleteAssistantConversationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistantConversations

> DeleteAssistantConversationsResponse DeleteAssistantConversations(ctx).Execute()

Delete All Assistant Conversations



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
	resp, r, err := apiClient.AssistantAPI.DeleteAssistantConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantAPI.DeleteAssistantConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssistantConversations`: DeleteAssistantConversationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantAPI.DeleteAssistantConversations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssistantConversationsRequest struct via the builder pattern


### Return type

[**DeleteAssistantConversationsResponse**](DeleteAssistantConversationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssistantUsage

> GetAssistantUsageResponse GetAssistantUsage(ctx).Execute()

Get Assistant Usage



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
	resp, r, err := apiClient.AssistantAPI.GetAssistantUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantAPI.GetAssistantUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssistantUsage`: GetAssistantUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantAPI.GetAssistantUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssistantUsageRequest struct via the builder pattern


### Return type

[**GetAssistantUsageResponse**](GetAssistantUsageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssistantConversations

> ListAssistantConversationsResponse ListAssistantConversations(ctx).Execute()

Get Assistant Conversations



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
	resp, r, err := apiClient.AssistantAPI.ListAssistantConversations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantAPI.ListAssistantConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssistantConversations`: ListAssistantConversationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantAPI.ListAssistantConversations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAssistantConversationsRequest struct via the builder pattern


### Return type

[**ListAssistantConversationsResponse**](ListAssistantConversationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAssistantConversation

> ReadAssistantConversationResponse ReadAssistantConversation(ctx, sessionId).Since(since).Execute()

Get Assistant Conversation



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
	sessionId := "sessionId_example" // string | The session you want to retrieve.
	since := int64(789) // int64 | Unix timestamp of the conversation you want to get since this time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssistantAPI.ReadAssistantConversation(context.Background(), sessionId).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssistantAPI.ReadAssistantConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAssistantConversation`: ReadAssistantConversationResponse
	fmt.Fprintf(os.Stdout, "Response from `AssistantAPI.ReadAssistantConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAssistantConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **int64** | Unix timestamp of the conversation you want to get since this time. | 

### Return type

[**ReadAssistantConversationResponse**](ReadAssistantConversationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

