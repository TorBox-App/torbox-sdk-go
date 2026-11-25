# \UserAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReferralToAccount**](UserAPI.md#AddReferralToAccount) | **Post** /v1/api/user/addreferral | Add Referral To Account
[**DeleteAccount**](UserAPI.md#DeleteAccount) | **Delete** /v1/api/user/deleteme | Delete Account
[**EditUserSettings**](UserAPI.md#EditUserSettings) | **Put** /v1/api/user/settings/editsettings | Edit User Settings
[**GetConfirmationCode**](UserAPI.md#GetConfirmationCode) | **Get** /v1/api/user/getconfirmation | Get Confirmation Code
[**GetReferralData**](UserAPI.md#GetReferralData) | **Get** /v1/api/user/referraldata | Get Referral Data
[**GetSubscriptions**](UserAPI.md#GetSubscriptions) | **Get** /v1/api/user/subscriptions | Get Subscriptions
[**GetTokenFromDeviceCodeProcess**](UserAPI.md#GetTokenFromDeviceCodeProcess) | **Post** /v1/api/user/auth/device/token | Get Token From Device Code Process
[**GetTransactionPdf**](UserAPI.md#GetTransactionPdf) | **Get** /v1/api/user/transaction/pdf | Get Transaction PDF
[**GetTransactions**](UserAPI.md#GetTransactions) | **Get** /v1/api/user/transactions | Get Transactions
[**GetUserData**](UserAPI.md#GetUserData) | **Get** /v1/api/user/me | Get User Data
[**GetUserStats**](UserAPI.md#GetUserStats) | **Get** /v1/api/user/stats | Get User Stats
[**RefreshApiToken**](UserAPI.md#RefreshApiToken) | **Post** /v1/api/user/refreshtoken | Refresh API Token
[**StartDeviceCodeAuthorizationProcess**](UserAPI.md#StartDeviceCodeAuthorizationProcess) | **Get** /v1/api/user/auth/device/start | Start Device Code Authorization Process



## AddReferralToAccount

> AddReferralToAccountResponse AddReferralToAccount(ctx).Referral(referral).Execute()

Add Referral To Account



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
	referral := "referral_example" // string | A referral code. Must be UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AddReferralToAccount(context.Background()).Referral(referral).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AddReferralToAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddReferralToAccount`: AddReferralToAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AddReferralToAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddReferralToAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referral** | **string** | A referral code. Must be UUID. | 

### Return type

[**AddReferralToAccountResponse**](AddReferralToAccountResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccountResponse DeleteAccount(ctx).DeleteAccountSchema(deleteAccountSchema).Execute()

Delete Account



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
	deleteAccountSchema := *openapiclient.NewDeleteAccountSchema("SessionToken_example", int64(123)) // DeleteAccountSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.DeleteAccount(context.Background()).DeleteAccountSchema(deleteAccountSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccount`: DeleteAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAccountSchema** | [**DeleteAccountSchema**](DeleteAccountSchema.md) |  | 

### Return type

[**DeleteAccountResponse**](DeleteAccountResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditUserSettings

> EditUserSettingsResponse EditUserSettings(ctx).BaseSettingsModel(baseSettingsModel).Execute()

Edit User Settings



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
	baseSettingsModel := *openapiclient.NewBaseSettingsModel() // BaseSettingsModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.EditUserSettings(context.Background()).BaseSettingsModel(baseSettingsModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.EditUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditUserSettings`: EditUserSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.EditUserSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseSettingsModel** | [**BaseSettingsModel**](BaseSettingsModel.md) |  | 

### Return type

[**EditUserSettingsResponse**](EditUserSettingsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfirmationCode

> GetConfirmationCodeResponse GetConfirmationCode(ctx).Execute()

Get Confirmation Code



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
	resp, r, err := apiClient.UserAPI.GetConfirmationCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetConfirmationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfirmationCode`: GetConfirmationCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetConfirmationCode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfirmationCodeRequest struct via the builder pattern


### Return type

[**GetConfirmationCodeResponse**](GetConfirmationCodeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferralData

> GetReferralDataResponse GetReferralData(ctx).Execute()

Get Referral Data



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
	resp, r, err := apiClient.UserAPI.GetReferralData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetReferralData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferralData`: GetReferralDataResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetReferralData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralDataRequest struct via the builder pattern


### Return type

[**GetReferralDataResponse**](GetReferralDataResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> GetSubscriptionsResponse GetSubscriptions(ctx).Execute()

Get Subscriptions



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
	resp, r, err := apiClient.UserAPI.GetSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: GetSubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


### Return type

[**GetSubscriptionsResponse**](GetSubscriptionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenFromDeviceCodeProcess

> GetTokenFromDeviceCodeProcessResponse GetTokenFromDeviceCodeProcess(ctx).DeviceTokenSchema(deviceTokenSchema).Execute()

Get Token From Device Code Process



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
	deviceTokenSchema := *openapiclient.NewDeviceTokenSchema("DeviceCode_example") // DeviceTokenSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTokenFromDeviceCodeProcess(context.Background()).DeviceTokenSchema(deviceTokenSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTokenFromDeviceCodeProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenFromDeviceCodeProcess`: GetTokenFromDeviceCodeProcessResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTokenFromDeviceCodeProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenFromDeviceCodeProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceTokenSchema** | [**DeviceTokenSchema**](DeviceTokenSchema.md) |  | 

### Return type

[**GetTokenFromDeviceCodeProcessResponse**](GetTokenFromDeviceCodeProcessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionPdf

> string GetTransactionPdf(ctx).TransactionId(transactionId).Execute()

Get Transaction PDF



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
	transactionId := "transactionId_example" // string | The transaction ID that you want to generate the PDF for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTransactionPdf(context.Background()).TransactionId(transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTransactionPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionPdf`: string
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTransactionPdf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **string** | The transaction ID that you want to generate the PDF for. | 

### Return type

**string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> GetTransactionsResponse GetTransactions(ctx).Execute()

Get Transactions



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
	resp, r, err := apiClient.UserAPI.GetTransactions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactions`: GetTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTransactions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


### Return type

[**GetTransactionsResponse**](GetTransactionsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserData

> GetUserDataResponse GetUserData(ctx).Settings(settings).DownloadStats(downloadStats).Referral(referral).Execute()

Get User Data



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
	settings := true // bool | Allows you to retrieve user settings. (optional) (default to false)
	downloadStats := true // bool |  (optional) (default to false)
	referral := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserData(context.Background()).Settings(settings).DownloadStats(downloadStats).Referral(referral).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserData`: GetUserDataResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | **bool** | Allows you to retrieve user settings. | [default to false]
 **downloadStats** | **bool** |  | [default to false]
 **referral** | **bool** |  | [default to false]

### Return type

[**GetUserDataResponse**](GetUserDataResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserStats

> GetUserStatsResponse GetUserStats(ctx).General(general).Bandwidth(bandwidth).BandwidthGrouping(bandwidthGrouping).Execute()

Get User Stats



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
	general := true // bool | Retrieve general stats. (optional) (default to true)
	bandwidth := true // bool | Retrieve bandwidth stats. (optional) (default to false)
	bandwidthGrouping := "bandwidthGrouping_example" // string | If the \"bandwidth\" query parameter is sent, this value is used. Allowed is \"hour\", \"day\", \"week\", and \"month\", this narrates how bandwidth is grouped. More granular shows more information. (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserStats(context.Background()).General(general).Bandwidth(bandwidth).BandwidthGrouping(bandwidthGrouping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserStats`: GetUserStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **general** | **bool** | Retrieve general stats. | [default to true]
 **bandwidth** | **bool** | Retrieve bandwidth stats. | [default to false]
 **bandwidthGrouping** | **string** | If the \&quot;bandwidth\&quot; query parameter is sent, this value is used. Allowed is \&quot;hour\&quot;, \&quot;day\&quot;, \&quot;week\&quot;, and \&quot;month\&quot;, this narrates how bandwidth is grouped. More granular shows more information. | [default to &quot;day&quot;]

### Return type

[**GetUserStatsResponse**](GetUserStatsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshApiToken

> RefreshApiTokenResponse RefreshApiToken(ctx).RefreshTokenSchema(refreshTokenSchema).Execute()

Refresh API Token



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
	refreshTokenSchema := *openapiclient.NewRefreshTokenSchema("SessionToken_example") // RefreshTokenSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RefreshApiToken(context.Background()).RefreshTokenSchema(refreshTokenSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RefreshApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshApiToken`: RefreshApiTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RefreshApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenSchema** | [**RefreshTokenSchema**](RefreshTokenSchema.md) |  | 

### Return type

[**RefreshApiTokenResponse**](RefreshApiTokenResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartDeviceCodeAuthorizationProcess

> StartDeviceCodeAuthorizationProcessResponse StartDeviceCodeAuthorizationProcess(ctx).App(app).Execute()

Start Device Code Authorization Process



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
	app := "app_example" // string | The app name that the user will see when verifying the code. This should be the name of your app. (optional) (default to "Third Party App")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.StartDeviceCodeAuthorizationProcess(context.Background()).App(app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.StartDeviceCodeAuthorizationProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartDeviceCodeAuthorizationProcess`: StartDeviceCodeAuthorizationProcessResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.StartDeviceCodeAuthorizationProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartDeviceCodeAuthorizationProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **string** | The app name that the user will see when verifying the code. This should be the name of your app. | [default to &quot;Third Party App&quot;]

### Return type

[**StartDeviceCodeAuthorizationProcessResponse**](StartDeviceCodeAuthorizationProcessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

