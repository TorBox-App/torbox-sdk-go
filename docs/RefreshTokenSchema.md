# RefreshTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionToken** | **string** | you can get this by logging into https://torbox.app and finding the torbox_session_token in local storage | 

## Methods

### NewRefreshTokenSchema

`func NewRefreshTokenSchema(sessionToken string, ) *RefreshTokenSchema`

NewRefreshTokenSchema instantiates a new RefreshTokenSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenSchemaWithDefaults

`func NewRefreshTokenSchemaWithDefaults() *RefreshTokenSchema`

NewRefreshTokenSchemaWithDefaults instantiates a new RefreshTokenSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionToken

`func (o *RefreshTokenSchema) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *RefreshTokenSchema) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *RefreshTokenSchema) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


