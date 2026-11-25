# \VendorsAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendorAccount**](VendorsAPI.md#GetVendorAccount) | **Get** /v1/api/vendors/account | Get Vendor Account
[**GetVendorUser**](VendorsAPI.md#GetVendorUser) | **Get** /v1/api/vendors/getaccount | Get Vendor User
[**GetVendorUsers**](VendorsAPI.md#GetVendorUsers) | **Get** /v1/api/vendors/getaccounts | Get Vendor Users
[**RefreshVendorUsers**](VendorsAPI.md#RefreshVendorUsers) | **Patch** /v1/api/vendors/refresh | Refresh Vendor Users
[**RegisterVendor**](VendorsAPI.md#RegisterVendor) | **Post** /v1/api/vendors/register | Register Vendor
[**RegisterVendorUser**](VendorsAPI.md#RegisterVendorUser) | **Post** /v1/api/vendors/registeruser | Register Vendor User
[**RemoveVendorUser**](VendorsAPI.md#RemoveVendorUser) | **Delete** /v1/api/vendors/removeuser | Remove Vendor User
[**UpdateVendorAccount**](VendorsAPI.md#UpdateVendorAccount) | **Put** /v1/api/vendors/updateaccount | Update Vendor Account



## GetVendorAccount

> GetVendorAccountResponse GetVendorAccount(ctx).Execute()

Get Vendor Account



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
	resp, r, err := apiClient.VendorsAPI.GetVendorAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.GetVendorAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorAccount`: GetVendorAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.GetVendorAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorAccountRequest struct via the builder pattern


### Return type

[**GetVendorAccountResponse**](GetVendorAccountResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorUser

> GetVendorUserResponse GetVendorUser(ctx).UserAuthId(userAuthId).Execute()

Get Vendor User



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
	userAuthId := "userAuthId_example" // string | Auth id of the user, from the Get Vendor Users endpoint. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorsAPI.GetVendorUser(context.Background()).UserAuthId(userAuthId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.GetVendorUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorUser`: GetVendorUserResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.GetVendorUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuthId** | **string** | Auth id of the user, from the Get Vendor Users endpoint. Required. | 

### Return type

[**GetVendorUserResponse**](GetVendorUserResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorUsers

> GetVendorUsersResponse GetVendorUsers(ctx).Execute()

Get Vendor Users



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
	resp, r, err := apiClient.VendorsAPI.GetVendorUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.GetVendorUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorUsers`: GetVendorUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.GetVendorUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorUsersRequest struct via the builder pattern


### Return type

[**GetVendorUsersResponse**](GetVendorUsersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshVendorUsers

> RefreshVendorUsersResponse RefreshVendorUsers(ctx).Execute()

Refresh Vendor Users



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
	resp, r, err := apiClient.VendorsAPI.RefreshVendorUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.RefreshVendorUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshVendorUsers`: RefreshVendorUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.RefreshVendorUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshVendorUsersRequest struct via the builder pattern


### Return type

[**RefreshVendorUsersResponse**](RefreshVendorUsersResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterVendor

> RegisterVendorResponse RegisterVendor(ctx).VendorName(vendorName).VendorUrl(vendorUrl).Execute()

Register Vendor



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
	vendorName := "vendorName_example" // string | Name of the vendor. Required.
	vendorUrl := "vendorUrl_example" // string | Website of the vendor. Must start with http:// or https:// and cannot contain a port, a query string or a fragment. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorsAPI.RegisterVendor(context.Background()).VendorName(vendorName).VendorUrl(vendorUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.RegisterVendor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterVendor`: RegisterVendorResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.RegisterVendor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterVendorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendorName** | **string** | Name of the vendor. Required. | 
 **vendorUrl** | **string** | Website of the vendor. Must start with http:// or https:// and cannot contain a port, a query string or a fragment. Required. | 

### Return type

[**RegisterVendorResponse**](RegisterVendorResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterVendorUser

> RegisterVendorUserResponse RegisterVendorUser(ctx).UserEmail(userEmail).Execute()

Register Vendor User



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
	userEmail := "userEmail_example" // string | Email address of the user to register. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorsAPI.RegisterVendorUser(context.Background()).UserEmail(userEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.RegisterVendorUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterVendorUser`: RegisterVendorUserResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.RegisterVendorUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterVendorUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userEmail** | **string** | Email address of the user to register. Required. | 

### Return type

[**RegisterVendorUserResponse**](RegisterVendorUserResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVendorUser

> RemoveVendorUserResponse RemoveVendorUser(ctx).UserAuthId(userAuthId).Execute()

Remove Vendor User



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
	userAuthId := "userAuthId_example" // string | Auth id of the user to remove, from the Get Vendor Users endpoint. Required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorsAPI.RemoveVendorUser(context.Background()).UserAuthId(userAuthId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.RemoveVendorUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveVendorUser`: RemoveVendorUserResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.RemoveVendorUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVendorUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuthId** | **string** | Auth id of the user to remove, from the Get Vendor Users endpoint. Required. | 

### Return type

[**RemoveVendorUserResponse**](RemoveVendorUserResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVendorAccount

> UpdateVendorAccountResponse UpdateVendorAccount(ctx).VendorName(vendorName).VendorUrl(vendorUrl).Execute()

Update Vendor Account



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
	vendorName := "vendorName_example" // string | New name of the vendor. Optional, but at least one of vendor_name and vendor_url must be sent. (optional)
	vendorUrl := "vendorUrl_example" // string | New website of the vendor. Must start with http:// or https:// and cannot contain a port, a query string or a fragment. Optional, but at least one of vendor_name and vendor_url must be sent. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorsAPI.UpdateVendorAccount(context.Background()).VendorName(vendorName).VendorUrl(vendorUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorsAPI.UpdateVendorAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVendorAccount`: UpdateVendorAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `VendorsAPI.UpdateVendorAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVendorAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendorName** | **string** | New name of the vendor. Optional, but at least one of vendor_name and vendor_url must be sent. | 
 **vendorUrl** | **string** | New website of the vendor. Must start with http:// or https:// and cannot contain a port, a query string or a fragment. Optional, but at least one of vendor_name and vendor_url must be sent. | 

### Return type

[**UpdateVendorAccountResponse**](UpdateVendorAccountResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

