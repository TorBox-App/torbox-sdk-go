# EditWebDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebdlId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**AlternativeHashes** | Pointer to **[]string** |  | [optional] 
**Airlocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewEditWebDownload

`func NewEditWebDownload(webdlId int64, ) *EditWebDownload`

NewEditWebDownload instantiates a new EditWebDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditWebDownloadWithDefaults

`func NewEditWebDownloadWithDefaults() *EditWebDownload`

NewEditWebDownloadWithDefaults instantiates a new EditWebDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebdlId

`func (o *EditWebDownload) GetWebdlId() int64`

GetWebdlId returns the WebdlId field if non-nil, zero value otherwise.

### GetWebdlIdOk

`func (o *EditWebDownload) GetWebdlIdOk() (*int64, bool)`

GetWebdlIdOk returns a tuple with the WebdlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdlId

`func (o *EditWebDownload) SetWebdlId(v int64)`

SetWebdlId sets WebdlId field to given value.


### GetName

`func (o *EditWebDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditWebDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditWebDownload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditWebDownload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EditWebDownload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EditWebDownload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EditWebDownload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EditWebDownload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlternativeHashes

`func (o *EditWebDownload) GetAlternativeHashes() []string`

GetAlternativeHashes returns the AlternativeHashes field if non-nil, zero value otherwise.

### GetAlternativeHashesOk

`func (o *EditWebDownload) GetAlternativeHashesOk() (*[]string, bool)`

GetAlternativeHashesOk returns a tuple with the AlternativeHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeHashes

`func (o *EditWebDownload) SetAlternativeHashes(v []string)`

SetAlternativeHashes sets AlternativeHashes field to given value.

### HasAlternativeHashes

`func (o *EditWebDownload) HasAlternativeHashes() bool`

HasAlternativeHashes returns a boolean if a field has been set.

### GetAirlocked

`func (o *EditWebDownload) GetAirlocked() bool`

GetAirlocked returns the Airlocked field if non-nil, zero value otherwise.

### GetAirlockedOk

`func (o *EditWebDownload) GetAirlockedOk() (*bool, bool)`

GetAirlockedOk returns a tuple with the Airlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlocked

`func (o *EditWebDownload) SetAirlocked(v bool)`

SetAirlocked sets Airlocked field to given value.

### HasAirlocked

`func (o *EditWebDownload) HasAirlocked() bool`

HasAirlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


