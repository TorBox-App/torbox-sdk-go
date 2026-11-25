# EditUsenetDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsenetDownloadId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**AlternativeHashes** | Pointer to **[]string** |  | [optional] 
**Airlocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewEditUsenetDownload

`func NewEditUsenetDownload(usenetDownloadId int64, ) *EditUsenetDownload`

NewEditUsenetDownload instantiates a new EditUsenetDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditUsenetDownloadWithDefaults

`func NewEditUsenetDownloadWithDefaults() *EditUsenetDownload`

NewEditUsenetDownloadWithDefaults instantiates a new EditUsenetDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsenetDownloadId

`func (o *EditUsenetDownload) GetUsenetDownloadId() int64`

GetUsenetDownloadId returns the UsenetDownloadId field if non-nil, zero value otherwise.

### GetUsenetDownloadIdOk

`func (o *EditUsenetDownload) GetUsenetDownloadIdOk() (*int64, bool)`

GetUsenetDownloadIdOk returns a tuple with the UsenetDownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsenetDownloadId

`func (o *EditUsenetDownload) SetUsenetDownloadId(v int64)`

SetUsenetDownloadId sets UsenetDownloadId field to given value.


### GetName

`func (o *EditUsenetDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditUsenetDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditUsenetDownload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditUsenetDownload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EditUsenetDownload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EditUsenetDownload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EditUsenetDownload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EditUsenetDownload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlternativeHashes

`func (o *EditUsenetDownload) GetAlternativeHashes() []string`

GetAlternativeHashes returns the AlternativeHashes field if non-nil, zero value otherwise.

### GetAlternativeHashesOk

`func (o *EditUsenetDownload) GetAlternativeHashesOk() (*[]string, bool)`

GetAlternativeHashesOk returns a tuple with the AlternativeHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeHashes

`func (o *EditUsenetDownload) SetAlternativeHashes(v []string)`

SetAlternativeHashes sets AlternativeHashes field to given value.

### HasAlternativeHashes

`func (o *EditUsenetDownload) HasAlternativeHashes() bool`

HasAlternativeHashes returns a boolean if a field has been set.

### GetAirlocked

`func (o *EditUsenetDownload) GetAirlocked() bool`

GetAirlocked returns the Airlocked field if non-nil, zero value otherwise.

### GetAirlockedOk

`func (o *EditUsenetDownload) GetAirlockedOk() (*bool, bool)`

GetAirlockedOk returns a tuple with the Airlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlocked

`func (o *EditUsenetDownload) SetAirlocked(v bool)`

SetAirlocked sets Airlocked field to given value.

### HasAirlocked

`func (o *EditUsenetDownload) HasAirlocked() bool`

HasAirlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


