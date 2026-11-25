# GetTorrentInfoResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Trackers** | Pointer to **[]string** |  | [optional] 
**Seeds** | Pointer to **int64** |  | [optional] 
**Peers** | Pointer to **int64** |  | [optional] 
**Files** | Pointer to [**[]GetTorrentInfoResponseDataFilesInner**](GetTorrentInfoResponseDataFilesInner.md) |  | [optional] 

## Methods

### NewGetTorrentInfoResponseData

`func NewGetTorrentInfoResponseData() *GetTorrentInfoResponseData`

NewGetTorrentInfoResponseData instantiates a new GetTorrentInfoResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTorrentInfoResponseDataWithDefaults

`func NewGetTorrentInfoResponseDataWithDefaults() *GetTorrentInfoResponseData`

NewGetTorrentInfoResponseDataWithDefaults instantiates a new GetTorrentInfoResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTorrentInfoResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTorrentInfoResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTorrentInfoResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTorrentInfoResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHash

`func (o *GetTorrentInfoResponseData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetTorrentInfoResponseData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetTorrentInfoResponseData) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetTorrentInfoResponseData) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetSize

`func (o *GetTorrentInfoResponseData) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTorrentInfoResponseData) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTorrentInfoResponseData) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetTorrentInfoResponseData) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTrackers

`func (o *GetTorrentInfoResponseData) GetTrackers() []string`

GetTrackers returns the Trackers field if non-nil, zero value otherwise.

### GetTrackersOk

`func (o *GetTorrentInfoResponseData) GetTrackersOk() (*[]string, bool)`

GetTrackersOk returns a tuple with the Trackers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackers

`func (o *GetTorrentInfoResponseData) SetTrackers(v []string)`

SetTrackers sets Trackers field to given value.

### HasTrackers

`func (o *GetTorrentInfoResponseData) HasTrackers() bool`

HasTrackers returns a boolean if a field has been set.

### GetSeeds

`func (o *GetTorrentInfoResponseData) GetSeeds() int64`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *GetTorrentInfoResponseData) GetSeedsOk() (*int64, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *GetTorrentInfoResponseData) SetSeeds(v int64)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *GetTorrentInfoResponseData) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### GetPeers

`func (o *GetTorrentInfoResponseData) GetPeers() int64`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *GetTorrentInfoResponseData) GetPeersOk() (*int64, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *GetTorrentInfoResponseData) SetPeers(v int64)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *GetTorrentInfoResponseData) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetFiles

`func (o *GetTorrentInfoResponseData) GetFiles() []GetTorrentInfoResponseDataFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetTorrentInfoResponseData) GetFilesOk() (*[]GetTorrentInfoResponseDataFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetTorrentInfoResponseData) SetFiles(v []GetTorrentInfoResponseDataFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetTorrentInfoResponseData) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


