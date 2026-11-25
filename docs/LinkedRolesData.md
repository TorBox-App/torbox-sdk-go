# LinkedRolesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscordToken** | **string** | Discord OAuth access token of the user. It is issued by Discord, not by TorBox. Required. | 

## Methods

### NewLinkedRolesData

`func NewLinkedRolesData(discordToken string, ) *LinkedRolesData`

NewLinkedRolesData instantiates a new LinkedRolesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedRolesDataWithDefaults

`func NewLinkedRolesDataWithDefaults() *LinkedRolesData`

NewLinkedRolesDataWithDefaults instantiates a new LinkedRolesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscordToken

`func (o *LinkedRolesData) GetDiscordToken() string`

GetDiscordToken returns the DiscordToken field if non-nil, zero value otherwise.

### GetDiscordTokenOk

`func (o *LinkedRolesData) GetDiscordTokenOk() (*string, bool)`

GetDiscordTokenOk returns a tuple with the DiscordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordToken

`func (o *LinkedRolesData) SetDiscordToken(v string)`

SetDiscordToken sets DiscordToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


