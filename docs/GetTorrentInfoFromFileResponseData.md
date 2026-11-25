# GetTorrentInfoFromFileResponseData

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

### NewGetTorrentInfoFromFileResponseData

`func NewGetTorrentInfoFromFileResponseData() *GetTorrentInfoFromFileResponseData`

NewGetTorrentInfoFromFileResponseData instantiates a new GetTorrentInfoFromFileResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTorrentInfoFromFileResponseDataWithDefaults

`func NewGetTorrentInfoFromFileResponseDataWithDefaults() *GetTorrentInfoFromFileResponseData`

NewGetTorrentInfoFromFileResponseDataWithDefaults instantiates a new GetTorrentInfoFromFileResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTorrentInfoFromFileResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTorrentInfoFromFileResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTorrentInfoFromFileResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTorrentInfoFromFileResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHash

`func (o *GetTorrentInfoFromFileResponseData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetTorrentInfoFromFileResponseData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetTorrentInfoFromFileResponseData) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetTorrentInfoFromFileResponseData) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetSize

`func (o *GetTorrentInfoFromFileResponseData) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTorrentInfoFromFileResponseData) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTorrentInfoFromFileResponseData) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetTorrentInfoFromFileResponseData) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTrackers

`func (o *GetTorrentInfoFromFileResponseData) GetTrackers() []string`

GetTrackers returns the Trackers field if non-nil, zero value otherwise.

### GetTrackersOk

`func (o *GetTorrentInfoFromFileResponseData) GetTrackersOk() (*[]string, bool)`

GetTrackersOk returns a tuple with the Trackers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackers

`func (o *GetTorrentInfoFromFileResponseData) SetTrackers(v []string)`

SetTrackers sets Trackers field to given value.

### HasTrackers

`func (o *GetTorrentInfoFromFileResponseData) HasTrackers() bool`

HasTrackers returns a boolean if a field has been set.

### GetSeeds

`func (o *GetTorrentInfoFromFileResponseData) GetSeeds() int64`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *GetTorrentInfoFromFileResponseData) GetSeedsOk() (*int64, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *GetTorrentInfoFromFileResponseData) SetSeeds(v int64)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *GetTorrentInfoFromFileResponseData) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### GetPeers

`func (o *GetTorrentInfoFromFileResponseData) GetPeers() int64`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *GetTorrentInfoFromFileResponseData) GetPeersOk() (*int64, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *GetTorrentInfoFromFileResponseData) SetPeers(v int64)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *GetTorrentInfoFromFileResponseData) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetFiles

`func (o *GetTorrentInfoFromFileResponseData) GetFiles() []GetTorrentInfoResponseDataFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetTorrentInfoFromFileResponseData) GetFilesOk() (*[]GetTorrentInfoResponseDataFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetTorrentInfoFromFileResponseData) SetFiles(v []GetTorrentInfoResponseDataFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetTorrentInfoFromFileResponseData) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


