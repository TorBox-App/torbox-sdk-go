# GetQueuedDownloadsResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Magnet** | Pointer to **string** |  | [optional] 
**TorrentFile** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**NameOverride** | Pointer to **string** |  | [optional] 
**SeedTorrentOverride** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetQueuedDownloadsResponseDataInner

`func NewGetQueuedDownloadsResponseDataInner() *GetQueuedDownloadsResponseDataInner`

NewGetQueuedDownloadsResponseDataInner instantiates a new GetQueuedDownloadsResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQueuedDownloadsResponseDataInnerWithDefaults

`func NewGetQueuedDownloadsResponseDataInnerWithDefaults() *GetQueuedDownloadsResponseDataInner`

NewGetQueuedDownloadsResponseDataInnerWithDefaults instantiates a new GetQueuedDownloadsResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetQueuedDownloadsResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetQueuedDownloadsResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetQueuedDownloadsResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetQueuedDownloadsResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetQueuedDownloadsResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetQueuedDownloadsResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetQueuedDownloadsResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetQueuedDownloadsResponseDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMagnet

`func (o *GetQueuedDownloadsResponseDataInner) GetMagnet() string`

GetMagnet returns the Magnet field if non-nil, zero value otherwise.

### GetMagnetOk

`func (o *GetQueuedDownloadsResponseDataInner) GetMagnetOk() (*string, bool)`

GetMagnetOk returns a tuple with the Magnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnet

`func (o *GetQueuedDownloadsResponseDataInner) SetMagnet(v string)`

SetMagnet sets Magnet field to given value.

### HasMagnet

`func (o *GetQueuedDownloadsResponseDataInner) HasMagnet() bool`

HasMagnet returns a boolean if a field has been set.

### GetTorrentFile

`func (o *GetQueuedDownloadsResponseDataInner) GetTorrentFile() string`

GetTorrentFile returns the TorrentFile field if non-nil, zero value otherwise.

### GetTorrentFileOk

`func (o *GetQueuedDownloadsResponseDataInner) GetTorrentFileOk() (*string, bool)`

GetTorrentFileOk returns a tuple with the TorrentFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentFile

`func (o *GetQueuedDownloadsResponseDataInner) SetTorrentFile(v string)`

SetTorrentFile sets TorrentFile field to given value.

### HasTorrentFile

`func (o *GetQueuedDownloadsResponseDataInner) HasTorrentFile() bool`

HasTorrentFile returns a boolean if a field has been set.

### GetHash

`func (o *GetQueuedDownloadsResponseDataInner) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetQueuedDownloadsResponseDataInner) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetQueuedDownloadsResponseDataInner) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetQueuedDownloadsResponseDataInner) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetName

`func (o *GetQueuedDownloadsResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetQueuedDownloadsResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetQueuedDownloadsResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetQueuedDownloadsResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetQueuedDownloadsResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetQueuedDownloadsResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetQueuedDownloadsResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetQueuedDownloadsResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNameOverride

`func (o *GetQueuedDownloadsResponseDataInner) GetNameOverride() string`

GetNameOverride returns the NameOverride field if non-nil, zero value otherwise.

### GetNameOverrideOk

`func (o *GetQueuedDownloadsResponseDataInner) GetNameOverrideOk() (*string, bool)`

GetNameOverrideOk returns a tuple with the NameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOverride

`func (o *GetQueuedDownloadsResponseDataInner) SetNameOverride(v string)`

SetNameOverride sets NameOverride field to given value.

### HasNameOverride

`func (o *GetQueuedDownloadsResponseDataInner) HasNameOverride() bool`

HasNameOverride returns a boolean if a field has been set.

### GetSeedTorrentOverride

`func (o *GetQueuedDownloadsResponseDataInner) GetSeedTorrentOverride() int64`

GetSeedTorrentOverride returns the SeedTorrentOverride field if non-nil, zero value otherwise.

### GetSeedTorrentOverrideOk

`func (o *GetQueuedDownloadsResponseDataInner) GetSeedTorrentOverrideOk() (*int64, bool)`

GetSeedTorrentOverrideOk returns a tuple with the SeedTorrentOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedTorrentOverride

`func (o *GetQueuedDownloadsResponseDataInner) SetSeedTorrentOverride(v int64)`

SetSeedTorrentOverride sets SeedTorrentOverride field to given value.

### HasSeedTorrentOverride

`func (o *GetQueuedDownloadsResponseDataInner) HasSeedTorrentOverride() bool`

HasSeedTorrentOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


