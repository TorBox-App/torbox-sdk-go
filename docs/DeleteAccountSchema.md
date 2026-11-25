# DeleteAccountSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionToken** | **string** | The user&#39;s session token from the website. It must belong to the same account as the API key. | 
**ConfirmationCode** | **int64** | The confirmation code that was sent to the user&#39;s email by the Get Confirmation Code endpoint. | 

## Methods

### NewDeleteAccountSchema

`func NewDeleteAccountSchema(sessionToken string, confirmationCode int64, ) *DeleteAccountSchema`

NewDeleteAccountSchema instantiates a new DeleteAccountSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAccountSchemaWithDefaults

`func NewDeleteAccountSchemaWithDefaults() *DeleteAccountSchema`

NewDeleteAccountSchemaWithDefaults instantiates a new DeleteAccountSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionToken

`func (o *DeleteAccountSchema) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *DeleteAccountSchema) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *DeleteAccountSchema) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.


### GetConfirmationCode

`func (o *DeleteAccountSchema) GetConfirmationCode() int64`

GetConfirmationCode returns the ConfirmationCode field if non-nil, zero value otherwise.

### GetConfirmationCodeOk

`func (o *DeleteAccountSchema) GetConfirmationCodeOk() (*int64, bool)`

GetConfirmationCodeOk returns a tuple with the ConfirmationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationCode

`func (o *DeleteAccountSchema) SetConfirmationCode(v int64)`

SetConfirmationCode sets ConfirmationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


