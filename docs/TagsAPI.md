# \TagsAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlTag**](TagsAPI.md#ControlTag) | **Post** /v1/api/tags/controltag | Control Tag
[**EditTag**](TagsAPI.md#EditTag) | **Put** /v1/api/tags/edittag | Edit Tag
[**GetTags**](TagsAPI.md#GetTags) | **Get** /v1/api/tags/mytags | Get User Tags



## ControlTag

> ControlTagResponse ControlTag(ctx).ControlTag(controlTag).Execute()

Control Tag



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
	controlTag := *openapiclient.NewControlTag(int64(123), "Operation_example") // ControlTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.ControlTag(context.Background()).ControlTag(controlTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ControlTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlTag`: ControlTagResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ControlTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlTag** | [**ControlTag**](ControlTag.md) |  | 

### Return type

[**ControlTagResponse**](ControlTagResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTag

> EditTagResponse EditTag(ctx).EditTag(editTag).Execute()

Edit Tag



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
	editTag := *openapiclient.NewEditTag(int64(123), "NewName_example") // EditTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.EditTag(context.Background()).EditTag(editTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.EditTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTag`: EditTagResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.EditTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editTag** | [**EditTag**](EditTag.md) |  | 

### Return type

[**EditTagResponse**](EditTagResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> GetTagsResponse GetTags(ctx).Execute()

Get User Tags



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
	resp, r, err := apiClient.TagsAPI.GetTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: GetTagsResponse
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


### Return type

[**GetTagsResponse**](GetTagsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

