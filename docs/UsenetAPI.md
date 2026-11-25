# \UsenetAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncCreateUsenetDownload**](UsenetAPI.md#AsyncCreateUsenetDownload) | **Post** /v1/api/usenet/asynccreateusenetdownload | Async Create Usenet Download
[**ControlUsenetDownload**](UsenetAPI.md#ControlUsenetDownload) | **Post** /v1/api/usenet/controlusenetdownload | Control Usenet Download
[**CreateUsenetDownload**](UsenetAPI.md#CreateUsenetDownload) | **Post** /v1/api/usenet/createusenetdownload | Create Usenet Download
[**EditUsenetItem**](UsenetAPI.md#EditUsenetItem) | **Put** /v1/api/usenet/editusenetdownload | Edit Usenet Item
[**GetNewsServerAccount**](UsenetAPI.md#GetNewsServerAccount) | **Get** /v1/api/usenet/provider/account | Get News Server Account
[**GetNewsServerConnectionInformation**](UsenetAPI.md#GetNewsServerConnectionInformation) | **Get** /v1/api/usenet/provider/connection | Get News Server Connection Information
[**GetUsenetCachedAvailability**](UsenetAPI.md#GetUsenetCachedAvailability) | **Get** /v1/api/usenet/checkcached | Get Usenet Cached Availability
[**GetUsenetCachedAvailabilityBatch**](UsenetAPI.md#GetUsenetCachedAvailabilityBatch) | **Post** /v1/api/usenet/checkcached | Get Usenet Cached Availability Batch
[**GetUsenetList**](UsenetAPI.md#GetUsenetList) | **Get** /v1/api/usenet/mylist | Get Usenet List
[**RequestUsenetDownloadLink**](UsenetAPI.md#RequestUsenetDownloadLink) | **Get** /v1/api/usenet/requestdl | Request Download Link
[**RequestUsenetDownloadLinksGet**](UsenetAPI.md#RequestUsenetDownloadLinksGet) | **Get** /v1/api/usenet/requestdl/batch | Request Download Links - GET
[**RequestUsenetDownloadLinksPost**](UsenetAPI.md#RequestUsenetDownloadLinksPost) | **Post** /v1/api/usenet/requestdl/batch | Request Download Links - POST
[**ResetNewsServerAccountPassword**](UsenetAPI.md#ResetNewsServerAccountPassword) | **Post** /v1/api/usenet/provider/account/resetpw | Reset News Server Account Password



## AsyncCreateUsenetDownload

> AsyncCreateUsenetDownloadResponse AsyncCreateUsenetDownload(ctx).Link(link).File(file).Name(name).Password(password).PostProcessing(postProcessing).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()

Async Create Usenet Download



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
	link := "link_example" // string | An accessible link to an NZB file. Cannot be a redirection. Optional. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An NZB File. Optional. (optional)
	name := "name_example" // string | The name you want the usenet download to be. Optional. (optional)
	password := "password_example" // string | The password to send to the Usenet client. This is used for extracting the RAR at the end. Optional. (optional)
	postProcessing := int64(789) // int64 | The post processing option to use. Read the overview of the Create Usenet Download endpoint for the options. Default is -1. Optional. (optional) (default to -1)
	asQueued := true // bool | Tells TorBox you want this usenet download instantly queued. Optional. (optional) (default to false)
	addOnlyIfCached := true // bool | Only adds the download if it is cached on TorBox. If not cached, it won't be added. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.AsyncCreateUsenetDownload(context.Background()).Link(link).File(file).Name(name).Password(password).PostProcessing(postProcessing).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.AsyncCreateUsenetDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsyncCreateUsenetDownload`: AsyncCreateUsenetDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.AsyncCreateUsenetDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsyncCreateUsenetDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** | An accessible link to an NZB file. Cannot be a redirection. Optional. | 
 **file** | ***os.File** | An NZB File. Optional. | 
 **name** | **string** | The name you want the usenet download to be. Optional. | 
 **password** | **string** | The password to send to the Usenet client. This is used for extracting the RAR at the end. Optional. | 
 **postProcessing** | **int64** | The post processing option to use. Read the overview of the Create Usenet Download endpoint for the options. Default is -1. Optional. | [default to -1]
 **asQueued** | **bool** | Tells TorBox you want this usenet download instantly queued. Optional. | [default to false]
 **addOnlyIfCached** | **bool** | Only adds the download if it is cached on TorBox. If not cached, it won&#39;t be added. | [default to false]

### Return type

[**AsyncCreateUsenetDownloadResponse**](AsyncCreateUsenetDownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlUsenetDownload

> ControlUsenetDownloadResponse ControlUsenetDownload(ctx).ControlUsenetDownload(controlUsenetDownload).Execute()

Control Usenet Download



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
	controlUsenetDownload := *openapiclient.NewControlUsenetDownload("Operation_example") // ControlUsenetDownload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.ControlUsenetDownload(context.Background()).ControlUsenetDownload(controlUsenetDownload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.ControlUsenetDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlUsenetDownload`: ControlUsenetDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.ControlUsenetDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiControlUsenetDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controlUsenetDownload** | [**ControlUsenetDownload**](ControlUsenetDownload.md) |  | 

### Return type

[**ControlUsenetDownloadResponse**](ControlUsenetDownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsenetDownload

> CreateUsenetDownloadResponse CreateUsenetDownload(ctx).Link(link).File(file).Name(name).Password(password).PostProcessing(postProcessing).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()

Create Usenet Download



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
	link := "link_example" // string | An accessible link to an NZB file. Cannot be a redirection. Optional. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | An NZB File. Optional. (optional)
	name := "name_example" // string | The name you want the usenet download to be. Optional. (optional)
	password := "password_example" // string | The password to send to the Usenet client. This is used for extracting the RAR at the end. Optional. (optional)
	postProcessing := int64(789) // int64 | The post processing option to use. Read the overview for the options. Default is -1. Optional. (optional) (default to -1)
	asQueued := true // bool | Tells TorBox you want this usenet download.instantly queued. Optional. (optional) (default to false)
	addOnlyIfCached := true // bool | Only adds the download if it is cached on TorBox. If not cached, it won't be added. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.CreateUsenetDownload(context.Background()).Link(link).File(file).Name(name).Password(password).PostProcessing(postProcessing).AsQueued(asQueued).AddOnlyIfCached(addOnlyIfCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.CreateUsenetDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUsenetDownload`: CreateUsenetDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.CreateUsenetDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsenetDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **link** | **string** | An accessible link to an NZB file. Cannot be a redirection. Optional. | 
 **file** | ***os.File** | An NZB File. Optional. | 
 **name** | **string** | The name you want the usenet download to be. Optional. | 
 **password** | **string** | The password to send to the Usenet client. This is used for extracting the RAR at the end. Optional. | 
 **postProcessing** | **int64** | The post processing option to use. Read the overview for the options. Default is -1. Optional. | [default to -1]
 **asQueued** | **bool** | Tells TorBox you want this usenet download.instantly queued. Optional. | [default to false]
 **addOnlyIfCached** | **bool** | Only adds the download if it is cached on TorBox. If not cached, it won&#39;t be added. | [default to false]

### Return type

[**CreateUsenetDownloadResponse**](CreateUsenetDownloadResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditUsenetItem

> EditUsenetItemResponse EditUsenetItem(ctx).EditUsenetDownload(editUsenetDownload).Execute()

Edit Usenet Item



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
	editUsenetDownload := *openapiclient.NewEditUsenetDownload(int64(123)) // EditUsenetDownload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.EditUsenetItem(context.Background()).EditUsenetDownload(editUsenetDownload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.EditUsenetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditUsenetItem`: EditUsenetItemResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.EditUsenetItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditUsenetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editUsenetDownload** | [**EditUsenetDownload**](EditUsenetDownload.md) |  | 

### Return type

[**EditUsenetItemResponse**](EditUsenetItemResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewsServerAccount

> GetNewsServerAccountResponse GetNewsServerAccount(ctx).Execute()

Get News Server Account



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
	resp, r, err := apiClient.UsenetAPI.GetNewsServerAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.GetNewsServerAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewsServerAccount`: GetNewsServerAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.GetNewsServerAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewsServerAccountRequest struct via the builder pattern


### Return type

[**GetNewsServerAccountResponse**](GetNewsServerAccountResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewsServerConnectionInformation

> GetNewsServerConnectionInformationResponse GetNewsServerConnectionInformation(ctx).Execute()

Get News Server Connection Information



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
	resp, r, err := apiClient.UsenetAPI.GetNewsServerConnectionInformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.GetNewsServerConnectionInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewsServerConnectionInformation`: GetNewsServerConnectionInformationResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.GetNewsServerConnectionInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewsServerConnectionInformationRequest struct via the builder pattern


### Return type

[**GetNewsServerConnectionInformationResponse**](GetNewsServerConnectionInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsenetCachedAvailability

> GetUsenetCachedAvailabilityResponse GetUsenetCachedAvailability(ctx).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()

Get Usenet Cached Availability



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
	hash := []string{"Inner_example"} // []string | The list of usenet hashes you want to check. Comma seperated. To find the hash, md5 the link of the usenet link or file. (optional)
	format := "format_example" // string | Format you want the data in. Acceptable is either \"object\" or \"list\". List is the most performant option as it doesn't require modification of the list. (optional) (default to "object")
	listFiles := true // bool | Allows you to list the files found inside the cached data. Optional. (optional) (default to false)
	checkCached := *openapiclient.NewCheckCached([]string{"Hashes_example"}) // CheckCached |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.GetUsenetCachedAvailability(context.Background()).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.GetUsenetCachedAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsenetCachedAvailability`: GetUsenetCachedAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.GetUsenetCachedAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsenetCachedAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **[]string** | The list of usenet hashes you want to check. Comma seperated. To find the hash, md5 the link of the usenet link or file. | 
 **format** | **string** | Format you want the data in. Acceptable is either \&quot;object\&quot; or \&quot;list\&quot;. List is the most performant option as it doesn&#39;t require modification of the list. | [default to &quot;object&quot;]
 **listFiles** | **bool** | Allows you to list the files found inside the cached data. Optional. | [default to false]
 **checkCached** | [**CheckCached**](CheckCached.md) |  | 

### Return type

[**GetUsenetCachedAvailabilityResponse**](GetUsenetCachedAvailabilityResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsenetCachedAvailabilityBatch

> GetUsenetCachedAvailabilityBatchResponse GetUsenetCachedAvailabilityBatch(ctx).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()

Get Usenet Cached Availability Batch



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
	resp, r, err := apiClient.UsenetAPI.GetUsenetCachedAvailabilityBatch(context.Background()).Hash(hash).Format(format).ListFiles(listFiles).CheckCached(checkCached).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.GetUsenetCachedAvailabilityBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsenetCachedAvailabilityBatch`: GetUsenetCachedAvailabilityBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.GetUsenetCachedAvailabilityBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsenetCachedAvailabilityBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hash** | **[]string** |  | 
 **format** | **string** |  | [default to &quot;object&quot;]
 **listFiles** | **bool** |  | [default to false]
 **checkCached** | [**CheckCached**](CheckCached.md) |  | 

### Return type

[**GetUsenetCachedAvailabilityBatchResponse**](GetUsenetCachedAvailabilityBatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsenetList

> GetUsenetListResponse GetUsenetList(ctx).Id(id).BypassCache(bypassCache).Offset(offset).Limit(limit).Files(files).Active(active).Execute()

Get Usenet List



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
	id := int64(789) // int64 | Determines the usenet download requested, will return an object rather than list. Optional. (optional)
	bypassCache := true // bool | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. (optional) (default to false)
	offset := int64(789) // int64 | Determines the offset of items to get from the database. Default is 0. Optional. (optional) (default to 0)
	limit := int64(789) // int64 | Determines the number of items to recieve per request. Default is 1000. Optional. (optional) (default to 1000)
	files := true // bool | Exclude files from response. Optional (optional) (default to true)
	active := true // bool | Shows only active downloads. Optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.GetUsenetList(context.Background()).Id(id).BypassCache(bypassCache).Offset(offset).Limit(limit).Files(files).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.GetUsenetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsenetList`: GetUsenetListResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.GetUsenetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsenetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | Determines the usenet download requested, will return an object rather than list. Optional. | 
 **bypassCache** | **bool** | Allows you to bypass the cached data, and always get fresh information. Useful if constantly querying for fresh download stats. Otherwise, we request that you save our database a few calls. | [default to false]
 **offset** | **int64** | Determines the offset of items to get from the database. Default is 0. Optional. | [default to 0]
 **limit** | **int64** | Determines the number of items to recieve per request. Default is 1000. Optional. | [default to 1000]
 **files** | **bool** | Exclude files from response. Optional | [default to true]
 **active** | **bool** | Shows only active downloads. Optional | 

### Return type

[**GetUsenetListResponse**](GetUsenetListResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestUsenetDownloadLink

> RequestUsenetDownloadLinkResponse RequestUsenetDownloadLink(ctx).Token(token).UsenetId(usenetId).FileId(fileId).ZipLink(zipLink).UserIp(userIp).Redirect(redirect).AppendName(appendName).Execute()

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
	usenetId := int64(789) // int64 | The usenet download's ID that you want to download
	fileId := int64(789) // int64 | The files's ID that you want to download. Optional if using \"zip_link\". (optional) (default to 0)
	zipLink := true // bool | If you want a zip link. Required if no file_id. Takes precedence over file_id if both are given. (optional) (default to false)
	userIp := "userIp_example" // string | The user's IPv4 to determine the closest CDN. Optional. (optional)
	redirect := true // bool | If you want to redirect the user to the CDN link. This is useful for creating permalinks so that you can just make this request URL the link. (optional) (default to false)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.RequestUsenetDownloadLink(context.Background()).Token(token).UsenetId(usenetId).FileId(fileId).ZipLink(zipLink).UserIp(userIp).Redirect(redirect).AppendName(appendName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.RequestUsenetDownloadLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestUsenetDownloadLink`: RequestUsenetDownloadLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.RequestUsenetDownloadLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestUsenetDownloadLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **usenetId** | **int64** | The usenet download&#39;s ID that you want to download | 
 **fileId** | **int64** | The files&#39;s ID that you want to download. Optional if using \&quot;zip_link\&quot;. | [default to 0]
 **zipLink** | **bool** | If you want a zip link. Required if no file_id. Takes precedence over file_id if both are given. | [default to false]
 **userIp** | **string** | The user&#39;s IPv4 to determine the closest CDN. Optional. | 
 **redirect** | **bool** | If you want to redirect the user to the CDN link. This is useful for creating permalinks so that you can just make this request URL the link. | [default to false]
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]

### Return type

[**RequestUsenetDownloadLinkResponse**](RequestUsenetDownloadLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestUsenetDownloadLinksGet

> RequestUsenetDownloadLinksGetResponse RequestUsenetDownloadLinksGet(ctx).Token(token).UsenetId(usenetId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()

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
	usenetId := int64(789) // int64 | The usenet download's ID that you want to download
	fileIds := []int64{int64(123)} // []int64 | The files's ID that you want to download (optional)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)
	requestDownloadBatch := *openapiclient.NewRequestDownloadBatch() // RequestDownloadBatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.RequestUsenetDownloadLinksGet(context.Background()).Token(token).UsenetId(usenetId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.RequestUsenetDownloadLinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestUsenetDownloadLinksGet`: RequestUsenetDownloadLinksGetResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.RequestUsenetDownloadLinksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestUsenetDownloadLinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **usenetId** | **int64** | The usenet download&#39;s ID that you want to download | 
 **fileIds** | **[]int64** | The files&#39;s ID that you want to download | 
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]
 **requestDownloadBatch** | [**RequestDownloadBatch**](RequestDownloadBatch.md) |  | 

### Return type

[**RequestUsenetDownloadLinksGetResponse**](RequestUsenetDownloadLinksGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestUsenetDownloadLinksPost

> RequestUsenetDownloadLinksPostResponse RequestUsenetDownloadLinksPost(ctx).Token(token).UsenetId(usenetId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()

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
	usenetId := int64(789) // int64 | The usenet downloads' ID that you want to download
	fileIds := []int64{int64(123)} // []int64 |  (optional)
	userIp := "userIp_example" // string | The user's IP to determine the closest CDN. Optional. (optional)
	appendName := true // bool | If you want to add the \"filename\" parameter to the returned CDN link. Optional. Default false. (optional) (default to false)
	requestDownloadBatch := *openapiclient.NewRequestDownloadBatch() // RequestDownloadBatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsenetAPI.RequestUsenetDownloadLinksPost(context.Background()).Token(token).UsenetId(usenetId).FileIds(fileIds).UserIp(userIp).AppendName(appendName).RequestDownloadBatch(requestDownloadBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.RequestUsenetDownloadLinksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestUsenetDownloadLinksPost`: RequestUsenetDownloadLinksPostResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.RequestUsenetDownloadLinksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestUsenetDownloadLinksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Your API Key | 
 **usenetId** | **int64** | The usenet downloads&#39; ID that you want to download | 
 **fileIds** | **[]int64** |  | 
 **userIp** | **string** | The user&#39;s IP to determine the closest CDN. Optional. | 
 **appendName** | **bool** | If you want to add the \&quot;filename\&quot; parameter to the returned CDN link. Optional. Default false. | [default to false]
 **requestDownloadBatch** | [**RequestDownloadBatch**](RequestDownloadBatch.md) |  | 

### Return type

[**RequestUsenetDownloadLinksPostResponse**](RequestUsenetDownloadLinksPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetNewsServerAccountPassword

> ResetNewsServerAccountPasswordResponse ResetNewsServerAccountPassword(ctx).Execute()

Reset News Server Account Password



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
	resp, r, err := apiClient.UsenetAPI.ResetNewsServerAccountPassword(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsenetAPI.ResetNewsServerAccountPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetNewsServerAccountPassword`: ResetNewsServerAccountPasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `UsenetAPI.ResetNewsServerAccountPassword`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetNewsServerAccountPasswordRequest struct via the builder pattern


### Return type

[**ResetNewsServerAccountPasswordResponse**](ResetNewsServerAccountPasswordResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

