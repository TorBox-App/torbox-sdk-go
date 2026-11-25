# \StreamAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStream**](StreamAPI.md#CreateStream) | **Get** /v1/api/stream/createstream | Create Stream
[**GetStreamData**](StreamAPI.md#GetStreamData) | **Get** /v1/api/stream/getstreamdata | Get Stream Data



## CreateStream

> CreateStreamResponse CreateStream(ctx).Id(id).FileId(fileId).Type_(type_).ChosenSubtitleIndex(chosenSubtitleIndex).ChosenAudioIndex(chosenAudioIndex).ChosenResolutionIndex(chosenResolutionIndex).ScrobblingEnabled(scrobblingEnabled).Execute()

Create Stream



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
	id := int64(789) // int64 | The ID of the type of download.
	fileId := int64(789) // int64 | The file ID inside of the download you would like to stream. (optional) (default to 0)
	type_ := "type__example" // string | The type of download that this is. If choosing the incorrect one, the ID will be incorrect. (optional) (default to "torrent")
	chosenSubtitleIndex := "chosenSubtitleIndex_example" // string | The relative index of the subtitle you would like to stream. This is default null for none. 0 is the first subtitle found. (optional)
	chosenAudioIndex := int64(789) // int64 | The relative index of the subtitle you would like to stream. Defaults to 0 for the first audio stream found inside the file. Cannot be null. (optional) (default to 0)
	chosenResolutionIndex := "chosenResolutionIndex_example" // string | The index of the resolution you would like to stream. This is default null for original. (optional)
	scrobblingEnabled := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamAPI.CreateStream(context.Background()).Id(id).FileId(fileId).Type_(type_).ChosenSubtitleIndex(chosenSubtitleIndex).ChosenAudioIndex(chosenAudioIndex).ChosenResolutionIndex(chosenResolutionIndex).ScrobblingEnabled(scrobblingEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.CreateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStream`: CreateStreamResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamAPI.CreateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The ID of the type of download. | 
 **fileId** | **int64** | The file ID inside of the download you would like to stream. | [default to 0]
 **type_** | **string** | The type of download that this is. If choosing the incorrect one, the ID will be incorrect. | [default to &quot;torrent&quot;]
 **chosenSubtitleIndex** | **string** | The relative index of the subtitle you would like to stream. This is default null for none. 0 is the first subtitle found. | 
 **chosenAudioIndex** | **int64** | The relative index of the subtitle you would like to stream. Defaults to 0 for the first audio stream found inside the file. Cannot be null. | [default to 0]
 **chosenResolutionIndex** | **string** | The index of the resolution you would like to stream. This is default null for original. | 
 **scrobblingEnabled** | **bool** |  | [default to true]

### Return type

[**CreateStreamResponse**](CreateStreamResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamData

> GetStreamDataResponse GetStreamData(ctx).Token(token).PresignedToken(presignedToken).ChosenSubtitleIndex(chosenSubtitleIndex).ChosenAudioIndex(chosenAudioIndex).ChosenResolutionIndex(chosenResolutionIndex).Execute()

Get Stream Data



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
	token := "token_example" // string | The user's API token, can also be used from the original, /createstream request.
	presignedToken := "presignedToken_example" // string | The presigned token given from the original, /createstream request.
	chosenSubtitleIndex := "chosenSubtitleIndex_example" // string | The chosen subtitle index that you are requesting information about. (optional)
	chosenAudioIndex := int64(789) // int64 | The chosen audio index that you are requesting information about. (optional) (default to 0)
	chosenResolutionIndex := "chosenResolutionIndex_example" // string | The index of the resolution you would like to stream. This is default null for original. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamAPI.GetStreamData(context.Background()).Token(token).PresignedToken(presignedToken).ChosenSubtitleIndex(chosenSubtitleIndex).ChosenAudioIndex(chosenAudioIndex).ChosenResolutionIndex(chosenResolutionIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.GetStreamData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamData`: GetStreamDataResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamAPI.GetStreamData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The user&#39;s API token, can also be used from the original, /createstream request. | 
 **presignedToken** | **string** | The presigned token given from the original, /createstream request. | 
 **chosenSubtitleIndex** | **string** | The chosen subtitle index that you are requesting information about. | 
 **chosenAudioIndex** | **int64** | The chosen audio index that you are requesting information about. | [default to 0]
 **chosenResolutionIndex** | **string** | The index of the resolution you would like to stream. This is default null for original. | 

### Return type

[**GetStreamDataResponse**](GetStreamDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

