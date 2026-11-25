# EditTorrent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TorrentId** | **int64** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**AlternativeHashes** | Pointer to **[]string** |  | [optional] 
**Airlocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewEditTorrent

`func NewEditTorrent(torrentId int64, ) *EditTorrent`

NewEditTorrent instantiates a new EditTorrent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditTorrentWithDefaults

`func NewEditTorrentWithDefaults() *EditTorrent`

NewEditTorrentWithDefaults instantiates a new EditTorrent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTorrentId

`func (o *EditTorrent) GetTorrentId() int64`

GetTorrentId returns the TorrentId field if non-nil, zero value otherwise.

### GetTorrentIdOk

`func (o *EditTorrent) GetTorrentIdOk() (*int64, bool)`

GetTorrentIdOk returns a tuple with the TorrentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentId

`func (o *EditTorrent) SetTorrentId(v int64)`

SetTorrentId sets TorrentId field to given value.


### GetName

`func (o *EditTorrent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditTorrent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditTorrent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditTorrent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EditTorrent) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EditTorrent) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EditTorrent) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EditTorrent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlternativeHashes

`func (o *EditTorrent) GetAlternativeHashes() []string`

GetAlternativeHashes returns the AlternativeHashes field if non-nil, zero value otherwise.

### GetAlternativeHashesOk

`func (o *EditTorrent) GetAlternativeHashesOk() (*[]string, bool)`

GetAlternativeHashesOk returns a tuple with the AlternativeHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeHashes

`func (o *EditTorrent) SetAlternativeHashes(v []string)`

SetAlternativeHashes sets AlternativeHashes field to given value.

### HasAlternativeHashes

`func (o *EditTorrent) HasAlternativeHashes() bool`

HasAlternativeHashes returns a boolean if a field has been set.

### GetAirlocked

`func (o *EditTorrent) GetAirlocked() bool`

GetAirlocked returns the Airlocked field if non-nil, zero value otherwise.

### GetAirlockedOk

`func (o *EditTorrent) GetAirlockedOk() (*bool, bool)`

GetAirlockedOk returns a tuple with the Airlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlocked

`func (o *EditTorrent) SetAirlocked(v bool)`

SetAirlocked sets Airlocked field to given value.

### HasAirlocked

`func (o *EditTorrent) HasAirlocked() bool`

HasAirlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


