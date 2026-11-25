# GetUsenetListResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**DownloadState** | Pointer to **string** |  | [optional] 
**DownloadSpeed** | Pointer to **float32** |  | [optional] 
**OriginalUrl** | Pointer to **string** |  | [optional] 
**Eta** | Pointer to **float32** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**DownloadId** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]GetTorrentListResponseDataInnerFilesInner**](GetTorrentListResponseDataInnerFilesInner.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Cached** | Pointer to **bool** |  | [optional] 
**DownloadPresent** | Pointer to **bool** |  | [optional] 
**DownloadFinished** | Pointer to **bool** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **float32** |  | [optional] 
**CachedAt** | Pointer to **string** |  | [optional] 
**AlternativeHashes** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Airlocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetUsenetListResponseDataInner

`func NewGetUsenetListResponseDataInner() *GetUsenetListResponseDataInner`

NewGetUsenetListResponseDataInner instantiates a new GetUsenetListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsenetListResponseDataInnerWithDefaults

`func NewGetUsenetListResponseDataInnerWithDefaults() *GetUsenetListResponseDataInner`

NewGetUsenetListResponseDataInnerWithDefaults instantiates a new GetUsenetListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUsenetListResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUsenetListResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUsenetListResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUsenetListResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetUsenetListResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUsenetListResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUsenetListResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetUsenetListResponseDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetUsenetListResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetUsenetListResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetUsenetListResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetUsenetListResponseDataInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAuthId

`func (o *GetUsenetListResponseDataInner) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *GetUsenetListResponseDataInner) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *GetUsenetListResponseDataInner) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *GetUsenetListResponseDataInner) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetName

`func (o *GetUsenetListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUsenetListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUsenetListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUsenetListResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHash

`func (o *GetUsenetListResponseDataInner) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetUsenetListResponseDataInner) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetUsenetListResponseDataInner) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetUsenetListResponseDataInner) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetDownloadState

`func (o *GetUsenetListResponseDataInner) GetDownloadState() string`

GetDownloadState returns the DownloadState field if non-nil, zero value otherwise.

### GetDownloadStateOk

`func (o *GetUsenetListResponseDataInner) GetDownloadStateOk() (*string, bool)`

GetDownloadStateOk returns a tuple with the DownloadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadState

`func (o *GetUsenetListResponseDataInner) SetDownloadState(v string)`

SetDownloadState sets DownloadState field to given value.

### HasDownloadState

`func (o *GetUsenetListResponseDataInner) HasDownloadState() bool`

HasDownloadState returns a boolean if a field has been set.

### GetDownloadSpeed

`func (o *GetUsenetListResponseDataInner) GetDownloadSpeed() float32`

GetDownloadSpeed returns the DownloadSpeed field if non-nil, zero value otherwise.

### GetDownloadSpeedOk

`func (o *GetUsenetListResponseDataInner) GetDownloadSpeedOk() (*float32, bool)`

GetDownloadSpeedOk returns a tuple with the DownloadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeed

`func (o *GetUsenetListResponseDataInner) SetDownloadSpeed(v float32)`

SetDownloadSpeed sets DownloadSpeed field to given value.

### HasDownloadSpeed

`func (o *GetUsenetListResponseDataInner) HasDownloadSpeed() bool`

HasDownloadSpeed returns a boolean if a field has been set.

### GetOriginalUrl

`func (o *GetUsenetListResponseDataInner) GetOriginalUrl() string`

GetOriginalUrl returns the OriginalUrl field if non-nil, zero value otherwise.

### GetOriginalUrlOk

`func (o *GetUsenetListResponseDataInner) GetOriginalUrlOk() (*string, bool)`

GetOriginalUrlOk returns a tuple with the OriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUrl

`func (o *GetUsenetListResponseDataInner) SetOriginalUrl(v string)`

SetOriginalUrl sets OriginalUrl field to given value.

### HasOriginalUrl

`func (o *GetUsenetListResponseDataInner) HasOriginalUrl() bool`

HasOriginalUrl returns a boolean if a field has been set.

### GetEta

`func (o *GetUsenetListResponseDataInner) GetEta() float32`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *GetUsenetListResponseDataInner) GetEtaOk() (*float32, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *GetUsenetListResponseDataInner) SetEta(v float32)`

SetEta sets Eta field to given value.

### HasEta

`func (o *GetUsenetListResponseDataInner) HasEta() bool`

HasEta returns a boolean if a field has been set.

### GetProgress

`func (o *GetUsenetListResponseDataInner) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GetUsenetListResponseDataInner) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GetUsenetListResponseDataInner) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *GetUsenetListResponseDataInner) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *GetUsenetListResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetUsenetListResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetUsenetListResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetUsenetListResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDownloadId

`func (o *GetUsenetListResponseDataInner) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *GetUsenetListResponseDataInner) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *GetUsenetListResponseDataInner) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *GetUsenetListResponseDataInner) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### GetFiles

`func (o *GetUsenetListResponseDataInner) GetFiles() []GetTorrentListResponseDataInnerFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetUsenetListResponseDataInner) GetFilesOk() (*[]GetTorrentListResponseDataInnerFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetUsenetListResponseDataInner) SetFiles(v []GetTorrentListResponseDataInnerFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetUsenetListResponseDataInner) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetActive

`func (o *GetUsenetListResponseDataInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetUsenetListResponseDataInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetUsenetListResponseDataInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetUsenetListResponseDataInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCached

`func (o *GetUsenetListResponseDataInner) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *GetUsenetListResponseDataInner) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *GetUsenetListResponseDataInner) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *GetUsenetListResponseDataInner) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetDownloadPresent

`func (o *GetUsenetListResponseDataInner) GetDownloadPresent() bool`

GetDownloadPresent returns the DownloadPresent field if non-nil, zero value otherwise.

### GetDownloadPresentOk

`func (o *GetUsenetListResponseDataInner) GetDownloadPresentOk() (*bool, bool)`

GetDownloadPresentOk returns a tuple with the DownloadPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPresent

`func (o *GetUsenetListResponseDataInner) SetDownloadPresent(v bool)`

SetDownloadPresent sets DownloadPresent field to given value.

### HasDownloadPresent

`func (o *GetUsenetListResponseDataInner) HasDownloadPresent() bool`

HasDownloadPresent returns a boolean if a field has been set.

### GetDownloadFinished

`func (o *GetUsenetListResponseDataInner) GetDownloadFinished() bool`

GetDownloadFinished returns the DownloadFinished field if non-nil, zero value otherwise.

### GetDownloadFinishedOk

`func (o *GetUsenetListResponseDataInner) GetDownloadFinishedOk() (*bool, bool)`

GetDownloadFinishedOk returns a tuple with the DownloadFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFinished

`func (o *GetUsenetListResponseDataInner) SetDownloadFinished(v bool)`

SetDownloadFinished sets DownloadFinished field to given value.

### HasDownloadFinished

`func (o *GetUsenetListResponseDataInner) HasDownloadFinished() bool`

HasDownloadFinished returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetUsenetListResponseDataInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetUsenetListResponseDataInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetUsenetListResponseDataInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetUsenetListResponseDataInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetServer

`func (o *GetUsenetListResponseDataInner) GetServer() float32`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetUsenetListResponseDataInner) GetServerOk() (*float32, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetUsenetListResponseDataInner) SetServer(v float32)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetUsenetListResponseDataInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetCachedAt

`func (o *GetUsenetListResponseDataInner) GetCachedAt() string`

GetCachedAt returns the CachedAt field if non-nil, zero value otherwise.

### GetCachedAtOk

`func (o *GetUsenetListResponseDataInner) GetCachedAtOk() (*string, bool)`

GetCachedAtOk returns a tuple with the CachedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedAt

`func (o *GetUsenetListResponseDataInner) SetCachedAt(v string)`

SetCachedAt sets CachedAt field to given value.

### HasCachedAt

`func (o *GetUsenetListResponseDataInner) HasCachedAt() bool`

HasCachedAt returns a boolean if a field has been set.

### GetAlternativeHashes

`func (o *GetUsenetListResponseDataInner) GetAlternativeHashes() []string`

GetAlternativeHashes returns the AlternativeHashes field if non-nil, zero value otherwise.

### GetAlternativeHashesOk

`func (o *GetUsenetListResponseDataInner) GetAlternativeHashesOk() (*[]string, bool)`

GetAlternativeHashesOk returns a tuple with the AlternativeHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeHashes

`func (o *GetUsenetListResponseDataInner) SetAlternativeHashes(v []string)`

SetAlternativeHashes sets AlternativeHashes field to given value.

### HasAlternativeHashes

`func (o *GetUsenetListResponseDataInner) HasAlternativeHashes() bool`

HasAlternativeHashes returns a boolean if a field has been set.

### GetTags

`func (o *GetUsenetListResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetUsenetListResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetUsenetListResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetUsenetListResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAirlocked

`func (o *GetUsenetListResponseDataInner) GetAirlocked() bool`

GetAirlocked returns the Airlocked field if non-nil, zero value otherwise.

### GetAirlockedOk

`func (o *GetUsenetListResponseDataInner) GetAirlockedOk() (*bool, bool)`

GetAirlockedOk returns a tuple with the Airlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlocked

`func (o *GetUsenetListResponseDataInner) SetAirlocked(v bool)`

SetAirlocked sets Airlocked field to given value.

### HasAirlocked

`func (o *GetUsenetListResponseDataInner) HasAirlocked() bool`

HasAirlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


