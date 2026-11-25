# \IntegrationsAPI

All URIs are relative to *https://api.torbox.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateOauth**](IntegrationsAPI.md#AuthenticateOauth) | **Get** /v1/api/integration/oauth/{provider} | Authenticate OAuth
[**CancelSpecificJob**](IntegrationsAPI.md#CancelSpecificJob) | **Delete** /v1/api/integration/job/{job_id} | Cancel Specific Job
[**GetAllJobs**](IntegrationsAPI.md#GetAllJobs) | **Get** /v1/api/integration/jobs | Get All Jobs
[**GetAllJobsByHash**](IntegrationsAPI.md#GetAllJobsByHash) | **Get** /v1/api/integration/jobs/{hash} | Get All Jobs By Hash
[**GetOauthIntegrations**](IntegrationsAPI.md#GetOauthIntegrations) | **Get** /v1/api/integration/oauth/me | Get OAuth Integrations
[**GetSpecificJob**](IntegrationsAPI.md#GetSpecificJob) | **Get** /v1/api/integration/job/{job_id} | Get Specific Job
[**LinkDiscordRoles**](IntegrationsAPI.md#LinkDiscordRoles) | **Post** /v1/api/integration/oauth/discord/linked_roles | Link Discord Roles
[**OauthCallbackGet**](IntegrationsAPI.md#OauthCallbackGet) | **Get** /v1/api/integration/oauth/{provider}/callback | OAuth Callback Get
[**OauthCallbackPost**](IntegrationsAPI.md#OauthCallbackPost) | **Post** /v1/api/integration/oauth/{provider}/callback | OAuth Callback Post
[**OauthSuccess**](IntegrationsAPI.md#OauthSuccess) | **Get** /v1/api/integration/oauth/{provider}/success | OAuth Success
[**RegisterOauthLink**](IntegrationsAPI.md#RegisterOauthLink) | **Post** /v1/api/integration/oauth/{provider}/register | Register OAuth Link
[**UnregisterOauthLink**](IntegrationsAPI.md#UnregisterOauthLink) | **Delete** /v1/api/integration/oauth/{provider}/unregister | Unregister OAuth Link
[**UploadTo1fichier**](IntegrationsAPI.md#UploadTo1fichier) | **Post** /v1/api/integration/1fichier | Upload To 1Fichier
[**UploadToDropbox**](IntegrationsAPI.md#UploadToDropbox) | **Post** /v1/api/integration/dropbox | Upload To Dropbox
[**UploadToGofile**](IntegrationsAPI.md#UploadToGofile) | **Post** /v1/api/integration/gofile | Upload To Gofile
[**UploadToGoogleDrive**](IntegrationsAPI.md#UploadToGoogleDrive) | **Post** /v1/api/integration/googledrive | Upload To Google Drive
[**UploadToOnedrive**](IntegrationsAPI.md#UploadToOnedrive) | **Post** /v1/api/integration/onedrive | Upload To OneDrive
[**UploadToPixeldrain**](IntegrationsAPI.md#UploadToPixeldrain) | **Post** /v1/api/integration/pixeldrain | Upload To Pixeldrain



## AuthenticateOauth

> AuthenticateOauth(ctx, provider).Execute()

Authenticate OAuth



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
	provider := "provider_example" // string | The service to log in with. One of `google`, `dropbox`, `discord`, `onedrive`, `patreon`, `apple`, `anilist`, `mal`, `trakt` or `simkl`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.AuthenticateOauth(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.AuthenticateOauth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The service to log in with. One of &#x60;google&#x60;, &#x60;dropbox&#x60;, &#x60;discord&#x60;, &#x60;onedrive&#x60;, &#x60;patreon&#x60;, &#x60;apple&#x60;, &#x60;anilist&#x60;, &#x60;mal&#x60;, &#x60;trakt&#x60; or &#x60;simkl&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateOauthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSpecificJob

> CancelSpecificJobResponse CancelSpecificJob(ctx, jobId).Execute()

Cancel Specific Job



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
	jobId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.CancelSpecificJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.CancelSpecificJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSpecificJob`: CancelSpecificJobResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.CancelSpecificJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSpecificJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelSpecificJobResponse**](CancelSpecificJobResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllJobs

> GetAllJobsResponse GetAllJobs(ctx).Execute()

Get All Jobs



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
	resp, r, err := apiClient.IntegrationsAPI.GetAllJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetAllJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllJobs`: GetAllJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetAllJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobsRequest struct via the builder pattern


### Return type

[**GetAllJobsResponse**](GetAllJobsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllJobsByHash

> GetAllJobsByHashResponse GetAllJobsByHash(ctx, hash).Execute()

Get All Jobs By Hash



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
	hash := "hash_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetAllJobsByHash(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetAllJobsByHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllJobsByHash`: GetAllJobsByHashResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetAllJobsByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobsByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllJobsByHashResponse**](GetAllJobsByHashResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthIntegrations

> GetOauthIntegrationsResponse GetOauthIntegrations(ctx).Execute()

Get OAuth Integrations



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
	resp, r, err := apiClient.IntegrationsAPI.GetOauthIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetOauthIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauthIntegrations`: GetOauthIntegrationsResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetOauthIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthIntegrationsRequest struct via the builder pattern


### Return type

[**GetOauthIntegrationsResponse**](GetOauthIntegrationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificJob

> GetSpecificJobResponse GetSpecificJob(ctx, jobId).Execute()

Get Specific Job



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
	jobId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.GetSpecificJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.GetSpecificJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecificJob`: GetSpecificJobResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.GetSpecificJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSpecificJobResponse**](GetSpecificJobResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkDiscordRoles

> LinkDiscordRolesResponse LinkDiscordRoles(ctx).LinkedRolesData(linkedRolesData).Execute()

Link Discord Roles



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
	linkedRolesData := *openapiclient.NewLinkedRolesData("DiscordToken_example") // LinkedRolesData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.LinkDiscordRoles(context.Background()).LinkedRolesData(linkedRolesData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.LinkDiscordRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkDiscordRoles`: LinkDiscordRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.LinkDiscordRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkDiscordRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkedRolesData** | [**LinkedRolesData**](LinkedRolesData.md) |  | 

### Return type

[**LinkDiscordRolesResponse**](LinkDiscordRolesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthCallbackGet

> OauthCallbackGet(ctx, provider).Execute()

OAuth Callback Get



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
	provider := "provider_example" // string | The service the login was started with. One of `google`, `dropbox`, `discord`, `onedrive`, `patreon`, `apple`, `anilist`, `mal`, `trakt` or `simkl`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.OauthCallbackGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OauthCallbackGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The service the login was started with. One of &#x60;google&#x60;, &#x60;dropbox&#x60;, &#x60;discord&#x60;, &#x60;onedrive&#x60;, &#x60;patreon&#x60;, &#x60;apple&#x60;, &#x60;anilist&#x60;, &#x60;mal&#x60;, &#x60;trakt&#x60; or &#x60;simkl&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthCallbackGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthCallbackPost

> OauthCallbackPost(ctx, provider).Execute()

OAuth Callback Post



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
	provider := "provider_example" // string | The service the login was started with. One of `google`, `dropbox`, `discord`, `onedrive`, `patreon`, `apple`, `anilist`, `mal`, `trakt` or `simkl`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.OauthCallbackPost(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OauthCallbackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The service the login was started with. One of &#x60;google&#x60;, &#x60;dropbox&#x60;, &#x60;discord&#x60;, &#x60;onedrive&#x60;, &#x60;patreon&#x60;, &#x60;apple&#x60;, &#x60;anilist&#x60;, &#x60;mal&#x60;, &#x60;trakt&#x60; or &#x60;simkl&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthCallbackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthSuccess

> OauthSuccessResponse OauthSuccess(ctx, provider).Execute()

OAuth Success



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
	provider := "provider_example" // string | The service. One of `google`, `dropbox`, `discord`, `onedrive` or `patreon`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.OauthSuccess(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.OauthSuccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthSuccess`: OauthSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.OauthSuccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The service. One of &#x60;google&#x60;, &#x60;dropbox&#x60;, &#x60;discord&#x60;, &#x60;onedrive&#x60; or &#x60;patreon&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOauthSuccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OauthSuccessResponse**](OauthSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterOauthLink

> RegisterOauthLinkResponse RegisterOauthLink(ctx, provider).OAuthRegisterRequest(oAuthRegisterRequest).Execute()

Register OAuth Link



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
	provider := "provider_example" // string | The service to link. One of `google`, `dropbox`, `discord`, `onedrive`, `patreon`, `apple`, `anilist`, `mal`, `trakt` or `simkl`.
	oAuthRegisterRequest := *openapiclient.NewOAuthRegisterRequest("Token_example", "RefreshToken_example") // OAuthRegisterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.RegisterOauthLink(context.Background(), provider).OAuthRegisterRequest(oAuthRegisterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.RegisterOauthLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterOauthLink`: RegisterOauthLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.RegisterOauthLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The service to link. One of &#x60;google&#x60;, &#x60;dropbox&#x60;, &#x60;discord&#x60;, &#x60;onedrive&#x60;, &#x60;patreon&#x60;, &#x60;apple&#x60;, &#x60;anilist&#x60;, &#x60;mal&#x60;, &#x60;trakt&#x60; or &#x60;simkl&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterOauthLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuthRegisterRequest** | [**OAuthRegisterRequest**](OAuthRegisterRequest.md) |  | 

### Return type

[**RegisterOauthLinkResponse**](RegisterOauthLinkResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterOauthLink

> UnregisterOauthLinkResponse UnregisterOauthLink(ctx, provider).Execute()

Unregister OAuth Link



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
	provider := "provider_example" // string | The service to unlink, by the name it was registered with in Register OAuth Link (`google`, `dropbox`, `discord`, `onedrive`, `patreon`, `apple`, `anilist`, `mal`, `trakt` or `simkl`). The name is not checked by the route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UnregisterOauthLink(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UnregisterOauthLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnregisterOauthLink`: UnregisterOauthLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UnregisterOauthLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The service to unlink, by the name it was registered with in Register OAuth Link (&#x60;google&#x60;, &#x60;dropbox&#x60;, &#x60;discord&#x60;, &#x60;onedrive&#x60;, &#x60;patreon&#x60;, &#x60;apple&#x60;, &#x60;anilist&#x60;, &#x60;mal&#x60;, &#x60;trakt&#x60; or &#x60;simkl&#x60;). The name is not checked by the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterOauthLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnregisterOauthLinkResponse**](UnregisterOauthLinkResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTo1fichier

> UploadTo1fichierResponse UploadTo1fichier(ctx).AddTo1Fichier(addTo1Fichier).Execute()

Upload To 1Fichier



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
	addTo1Fichier := *openapiclient.NewAddTo1Fichier(int64(123)) // AddTo1Fichier | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UploadTo1fichier(context.Background()).AddTo1Fichier(addTo1Fichier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UploadTo1fichier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadTo1fichier`: UploadTo1fichierResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UploadTo1fichier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTo1fichierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addTo1Fichier** | [**AddTo1Fichier**](AddTo1Fichier.md) |  | 

### Return type

[**UploadTo1fichierResponse**](UploadTo1fichierResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToDropbox

> UploadToDropbox(ctx).AddToDropbox(addToDropbox).Execute()

Upload To Dropbox



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
	addToDropbox := *openapiclient.NewAddToDropbox(int64(123), "DropboxToken_example") // AddToDropbox | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.UploadToDropbox(context.Background()).AddToDropbox(addToDropbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UploadToDropbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadToDropboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToDropbox** | [**AddToDropbox**](AddToDropbox.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToGofile

> UploadToGofileResponse UploadToGofile(ctx).AddToGofile(addToGofile).Execute()

Upload To Gofile



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
	addToGofile := *openapiclient.NewAddToGofile(int64(123)) // AddToGofile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UploadToGofile(context.Background()).AddToGofile(addToGofile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UploadToGofile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadToGofile`: UploadToGofileResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UploadToGofile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadToGofileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToGofile** | [**AddToGofile**](AddToGofile.md) |  | 

### Return type

[**UploadToGofileResponse**](UploadToGofileResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToGoogleDrive

> UploadToGoogleDriveResponse UploadToGoogleDrive(ctx).AddToGoogleDrive(addToGoogleDrive).Execute()

Upload To Google Drive



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
	addToGoogleDrive := *openapiclient.NewAddToGoogleDrive(int64(123), "GoogleToken_example") // AddToGoogleDrive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UploadToGoogleDrive(context.Background()).AddToGoogleDrive(addToGoogleDrive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UploadToGoogleDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadToGoogleDrive`: UploadToGoogleDriveResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UploadToGoogleDrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadToGoogleDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToGoogleDrive** | [**AddToGoogleDrive**](AddToGoogleDrive.md) |  | 

### Return type

[**UploadToGoogleDriveResponse**](UploadToGoogleDriveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToOnedrive

> UploadToOnedriveResponse UploadToOnedrive(ctx).AddToOnedrive(addToOnedrive).Execute()

Upload To OneDrive



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
	addToOnedrive := *openapiclient.NewAddToOnedrive(int64(123), "OnedriveToken_example") // AddToOnedrive | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UploadToOnedrive(context.Background()).AddToOnedrive(addToOnedrive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UploadToOnedrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadToOnedrive`: UploadToOnedriveResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UploadToOnedrive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadToOnedriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToOnedrive** | [**AddToOnedrive**](AddToOnedrive.md) |  | 

### Return type

[**UploadToOnedriveResponse**](UploadToOnedriveResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadToPixeldrain

> UploadToPixeldrainResponse UploadToPixeldrain(ctx).AddToPixeldrain(addToPixeldrain).Execute()

Upload To Pixeldrain



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
	addToPixeldrain := *openapiclient.NewAddToPixeldrain(int64(123)) // AddToPixeldrain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.UploadToPixeldrain(context.Background()).AddToPixeldrain(addToPixeldrain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.UploadToPixeldrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadToPixeldrain`: UploadToPixeldrainResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.UploadToPixeldrain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadToPixeldrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addToPixeldrain** | [**AddToPixeldrain**](AddToPixeldrain.md) |  | 

### Return type

[**UploadToPixeldrainResponse**](UploadToPixeldrainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

