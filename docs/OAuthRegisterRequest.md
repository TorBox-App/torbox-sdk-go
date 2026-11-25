# OAuthRegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | OAuth access token the service issued for the user. Required. | 
**RefreshToken** | **string** | OAuth refresh token the service issued for the user. Required. | 

## Methods

### NewOAuthRegisterRequest

`func NewOAuthRegisterRequest(token string, refreshToken string, ) *OAuthRegisterRequest`

NewOAuthRegisterRequest instantiates a new OAuthRegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRegisterRequestWithDefaults

`func NewOAuthRegisterRequestWithDefaults() *OAuthRegisterRequest`

NewOAuthRegisterRequestWithDefaults instantiates a new OAuthRegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *OAuthRegisterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OAuthRegisterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OAuthRegisterRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetRefreshToken

`func (o *OAuthRegisterRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuthRegisterRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuthRegisterRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


