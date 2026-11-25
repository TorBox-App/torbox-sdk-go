# GetTorrentListResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **float32** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Magnet** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**DownloadState** | Pointer to **string** |  | [optional] 
**Seeds** | Pointer to **int64** |  | [optional] 
**Peers** | Pointer to **int64** |  | [optional] 
**Ratio** | Pointer to **float32** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**DownloadSpeed** | Pointer to **float32** |  | [optional] 
**UploadSpeed** | Pointer to **float32** |  | [optional] 
**Eta** | Pointer to **float32** |  | [optional] 
**TorrentFile** | Pointer to **bool** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**DownloadPresent** | Pointer to **bool** |  | [optional] 
**Files** | Pointer to [**[]GetTorrentListResponseDataInnerFilesInner**](GetTorrentListResponseDataInnerFilesInner.md) |  | [optional] 
**DownloadPath** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**DownloadFinished** | Pointer to **bool** |  | [optional] 
**Tracker** | Pointer to **string** |  | [optional] 
**TotalUploaded** | Pointer to **int64** |  | [optional] 
**TotalDownloaded** | Pointer to **int64** |  | [optional] 
**Cached** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**SeedTorrent** | Pointer to **bool** |  | [optional] 
**AllowZipped** | Pointer to **bool** |  | [optional] 
**LongTermSeeding** | Pointer to **bool** |  | [optional] 
**TrackerMessage** | Pointer to **string** |  | [optional] 
**CachedAt** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**AlternativeHashes** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Airlocked** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetTorrentListResponseDataInner

`func NewGetTorrentListResponseDataInner() *GetTorrentListResponseDataInner`

NewGetTorrentListResponseDataInner instantiates a new GetTorrentListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTorrentListResponseDataInnerWithDefaults

`func NewGetTorrentListResponseDataInnerWithDefaults() *GetTorrentListResponseDataInner`

NewGetTorrentListResponseDataInnerWithDefaults instantiates a new GetTorrentListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTorrentListResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTorrentListResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTorrentListResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTorrentListResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthId

`func (o *GetTorrentListResponseDataInner) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *GetTorrentListResponseDataInner) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *GetTorrentListResponseDataInner) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *GetTorrentListResponseDataInner) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetServer

`func (o *GetTorrentListResponseDataInner) GetServer() float32`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetTorrentListResponseDataInner) GetServerOk() (*float32, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetTorrentListResponseDataInner) SetServer(v float32)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetTorrentListResponseDataInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetHash

`func (o *GetTorrentListResponseDataInner) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetTorrentListResponseDataInner) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetTorrentListResponseDataInner) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetTorrentListResponseDataInner) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetName

`func (o *GetTorrentListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTorrentListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTorrentListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTorrentListResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMagnet

`func (o *GetTorrentListResponseDataInner) GetMagnet() string`

GetMagnet returns the Magnet field if non-nil, zero value otherwise.

### GetMagnetOk

`func (o *GetTorrentListResponseDataInner) GetMagnetOk() (*string, bool)`

GetMagnetOk returns a tuple with the Magnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnet

`func (o *GetTorrentListResponseDataInner) SetMagnet(v string)`

SetMagnet sets Magnet field to given value.

### HasMagnet

`func (o *GetTorrentListResponseDataInner) HasMagnet() bool`

HasMagnet returns a boolean if a field has been set.

### GetSize

`func (o *GetTorrentListResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTorrentListResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTorrentListResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetTorrentListResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetActive

`func (o *GetTorrentListResponseDataInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetTorrentListResponseDataInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetTorrentListResponseDataInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetTorrentListResponseDataInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetTorrentListResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTorrentListResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTorrentListResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTorrentListResponseDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetTorrentListResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTorrentListResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTorrentListResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTorrentListResponseDataInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDownloadState

`func (o *GetTorrentListResponseDataInner) GetDownloadState() string`

GetDownloadState returns the DownloadState field if non-nil, zero value otherwise.

### GetDownloadStateOk

`func (o *GetTorrentListResponseDataInner) GetDownloadStateOk() (*string, bool)`

GetDownloadStateOk returns a tuple with the DownloadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadState

`func (o *GetTorrentListResponseDataInner) SetDownloadState(v string)`

SetDownloadState sets DownloadState field to given value.

### HasDownloadState

`func (o *GetTorrentListResponseDataInner) HasDownloadState() bool`

HasDownloadState returns a boolean if a field has been set.

### GetSeeds

`func (o *GetTorrentListResponseDataInner) GetSeeds() int64`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *GetTorrentListResponseDataInner) GetSeedsOk() (*int64, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *GetTorrentListResponseDataInner) SetSeeds(v int64)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *GetTorrentListResponseDataInner) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### GetPeers

`func (o *GetTorrentListResponseDataInner) GetPeers() int64`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *GetTorrentListResponseDataInner) GetPeersOk() (*int64, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *GetTorrentListResponseDataInner) SetPeers(v int64)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *GetTorrentListResponseDataInner) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetRatio

`func (o *GetTorrentListResponseDataInner) GetRatio() float32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *GetTorrentListResponseDataInner) GetRatioOk() (*float32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *GetTorrentListResponseDataInner) SetRatio(v float32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *GetTorrentListResponseDataInner) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetProgress

`func (o *GetTorrentListResponseDataInner) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GetTorrentListResponseDataInner) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GetTorrentListResponseDataInner) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *GetTorrentListResponseDataInner) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetDownloadSpeed

`func (o *GetTorrentListResponseDataInner) GetDownloadSpeed() float32`

GetDownloadSpeed returns the DownloadSpeed field if non-nil, zero value otherwise.

### GetDownloadSpeedOk

`func (o *GetTorrentListResponseDataInner) GetDownloadSpeedOk() (*float32, bool)`

GetDownloadSpeedOk returns a tuple with the DownloadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeed

`func (o *GetTorrentListResponseDataInner) SetDownloadSpeed(v float32)`

SetDownloadSpeed sets DownloadSpeed field to given value.

### HasDownloadSpeed

`func (o *GetTorrentListResponseDataInner) HasDownloadSpeed() bool`

HasDownloadSpeed returns a boolean if a field has been set.

### GetUploadSpeed

`func (o *GetTorrentListResponseDataInner) GetUploadSpeed() float32`

GetUploadSpeed returns the UploadSpeed field if non-nil, zero value otherwise.

### GetUploadSpeedOk

`func (o *GetTorrentListResponseDataInner) GetUploadSpeedOk() (*float32, bool)`

GetUploadSpeedOk returns a tuple with the UploadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpeed

`func (o *GetTorrentListResponseDataInner) SetUploadSpeed(v float32)`

SetUploadSpeed sets UploadSpeed field to given value.

### HasUploadSpeed

`func (o *GetTorrentListResponseDataInner) HasUploadSpeed() bool`

HasUploadSpeed returns a boolean if a field has been set.

### GetEta

`func (o *GetTorrentListResponseDataInner) GetEta() float32`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *GetTorrentListResponseDataInner) GetEtaOk() (*float32, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *GetTorrentListResponseDataInner) SetEta(v float32)`

SetEta sets Eta field to given value.

### HasEta

`func (o *GetTorrentListResponseDataInner) HasEta() bool`

HasEta returns a boolean if a field has been set.

### GetTorrentFile

`func (o *GetTorrentListResponseDataInner) GetTorrentFile() bool`

GetTorrentFile returns the TorrentFile field if non-nil, zero value otherwise.

### GetTorrentFileOk

`func (o *GetTorrentListResponseDataInner) GetTorrentFileOk() (*bool, bool)`

GetTorrentFileOk returns a tuple with the TorrentFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentFile

`func (o *GetTorrentListResponseDataInner) SetTorrentFile(v bool)`

SetTorrentFile sets TorrentFile field to given value.

### HasTorrentFile

`func (o *GetTorrentListResponseDataInner) HasTorrentFile() bool`

HasTorrentFile returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetTorrentListResponseDataInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetTorrentListResponseDataInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetTorrentListResponseDataInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetTorrentListResponseDataInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDownloadPresent

`func (o *GetTorrentListResponseDataInner) GetDownloadPresent() bool`

GetDownloadPresent returns the DownloadPresent field if non-nil, zero value otherwise.

### GetDownloadPresentOk

`func (o *GetTorrentListResponseDataInner) GetDownloadPresentOk() (*bool, bool)`

GetDownloadPresentOk returns a tuple with the DownloadPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPresent

`func (o *GetTorrentListResponseDataInner) SetDownloadPresent(v bool)`

SetDownloadPresent sets DownloadPresent field to given value.

### HasDownloadPresent

`func (o *GetTorrentListResponseDataInner) HasDownloadPresent() bool`

HasDownloadPresent returns a boolean if a field has been set.

### GetFiles

`func (o *GetTorrentListResponseDataInner) GetFiles() []GetTorrentListResponseDataInnerFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GetTorrentListResponseDataInner) GetFilesOk() (*[]GetTorrentListResponseDataInnerFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GetTorrentListResponseDataInner) SetFiles(v []GetTorrentListResponseDataInnerFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GetTorrentListResponseDataInner) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDownloadPath

`func (o *GetTorrentListResponseDataInner) GetDownloadPath() string`

GetDownloadPath returns the DownloadPath field if non-nil, zero value otherwise.

### GetDownloadPathOk

`func (o *GetTorrentListResponseDataInner) GetDownloadPathOk() (*string, bool)`

GetDownloadPathOk returns a tuple with the DownloadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPath

`func (o *GetTorrentListResponseDataInner) SetDownloadPath(v string)`

SetDownloadPath sets DownloadPath field to given value.

### HasDownloadPath

`func (o *GetTorrentListResponseDataInner) HasDownloadPath() bool`

HasDownloadPath returns a boolean if a field has been set.

### GetAvailability

`func (o *GetTorrentListResponseDataInner) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *GetTorrentListResponseDataInner) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *GetTorrentListResponseDataInner) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *GetTorrentListResponseDataInner) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetDownloadFinished

`func (o *GetTorrentListResponseDataInner) GetDownloadFinished() bool`

GetDownloadFinished returns the DownloadFinished field if non-nil, zero value otherwise.

### GetDownloadFinishedOk

`func (o *GetTorrentListResponseDataInner) GetDownloadFinishedOk() (*bool, bool)`

GetDownloadFinishedOk returns a tuple with the DownloadFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFinished

`func (o *GetTorrentListResponseDataInner) SetDownloadFinished(v bool)`

SetDownloadFinished sets DownloadFinished field to given value.

### HasDownloadFinished

`func (o *GetTorrentListResponseDataInner) HasDownloadFinished() bool`

HasDownloadFinished returns a boolean if a field has been set.

### GetTracker

`func (o *GetTorrentListResponseDataInner) GetTracker() string`

GetTracker returns the Tracker field if non-nil, zero value otherwise.

### GetTrackerOk

`func (o *GetTorrentListResponseDataInner) GetTrackerOk() (*string, bool)`

GetTrackerOk returns a tuple with the Tracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracker

`func (o *GetTorrentListResponseDataInner) SetTracker(v string)`

SetTracker sets Tracker field to given value.

### HasTracker

`func (o *GetTorrentListResponseDataInner) HasTracker() bool`

HasTracker returns a boolean if a field has been set.

### GetTotalUploaded

`func (o *GetTorrentListResponseDataInner) GetTotalUploaded() int64`

GetTotalUploaded returns the TotalUploaded field if non-nil, zero value otherwise.

### GetTotalUploadedOk

`func (o *GetTorrentListResponseDataInner) GetTotalUploadedOk() (*int64, bool)`

GetTotalUploadedOk returns a tuple with the TotalUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUploaded

`func (o *GetTorrentListResponseDataInner) SetTotalUploaded(v int64)`

SetTotalUploaded sets TotalUploaded field to given value.

### HasTotalUploaded

`func (o *GetTorrentListResponseDataInner) HasTotalUploaded() bool`

HasTotalUploaded returns a boolean if a field has been set.

### GetTotalDownloaded

`func (o *GetTorrentListResponseDataInner) GetTotalDownloaded() int64`

GetTotalDownloaded returns the TotalDownloaded field if non-nil, zero value otherwise.

### GetTotalDownloadedOk

`func (o *GetTorrentListResponseDataInner) GetTotalDownloadedOk() (*int64, bool)`

GetTotalDownloadedOk returns a tuple with the TotalDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownloaded

`func (o *GetTorrentListResponseDataInner) SetTotalDownloaded(v int64)`

SetTotalDownloaded sets TotalDownloaded field to given value.

### HasTotalDownloaded

`func (o *GetTorrentListResponseDataInner) HasTotalDownloaded() bool`

HasTotalDownloaded returns a boolean if a field has been set.

### GetCached

`func (o *GetTorrentListResponseDataInner) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *GetTorrentListResponseDataInner) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *GetTorrentListResponseDataInner) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *GetTorrentListResponseDataInner) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetOwner

`func (o *GetTorrentListResponseDataInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetTorrentListResponseDataInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetTorrentListResponseDataInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetTorrentListResponseDataInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSeedTorrent

`func (o *GetTorrentListResponseDataInner) GetSeedTorrent() bool`

GetSeedTorrent returns the SeedTorrent field if non-nil, zero value otherwise.

### GetSeedTorrentOk

`func (o *GetTorrentListResponseDataInner) GetSeedTorrentOk() (*bool, bool)`

GetSeedTorrentOk returns a tuple with the SeedTorrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedTorrent

`func (o *GetTorrentListResponseDataInner) SetSeedTorrent(v bool)`

SetSeedTorrent sets SeedTorrent field to given value.

### HasSeedTorrent

`func (o *GetTorrentListResponseDataInner) HasSeedTorrent() bool`

HasSeedTorrent returns a boolean if a field has been set.

### GetAllowZipped

`func (o *GetTorrentListResponseDataInner) GetAllowZipped() bool`

GetAllowZipped returns the AllowZipped field if non-nil, zero value otherwise.

### GetAllowZippedOk

`func (o *GetTorrentListResponseDataInner) GetAllowZippedOk() (*bool, bool)`

GetAllowZippedOk returns a tuple with the AllowZipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZipped

`func (o *GetTorrentListResponseDataInner) SetAllowZipped(v bool)`

SetAllowZipped sets AllowZipped field to given value.

### HasAllowZipped

`func (o *GetTorrentListResponseDataInner) HasAllowZipped() bool`

HasAllowZipped returns a boolean if a field has been set.

### GetLongTermSeeding

`func (o *GetTorrentListResponseDataInner) GetLongTermSeeding() bool`

GetLongTermSeeding returns the LongTermSeeding field if non-nil, zero value otherwise.

### GetLongTermSeedingOk

`func (o *GetTorrentListResponseDataInner) GetLongTermSeedingOk() (*bool, bool)`

GetLongTermSeedingOk returns a tuple with the LongTermSeeding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSeeding

`func (o *GetTorrentListResponseDataInner) SetLongTermSeeding(v bool)`

SetLongTermSeeding sets LongTermSeeding field to given value.

### HasLongTermSeeding

`func (o *GetTorrentListResponseDataInner) HasLongTermSeeding() bool`

HasLongTermSeeding returns a boolean if a field has been set.

### GetTrackerMessage

`func (o *GetTorrentListResponseDataInner) GetTrackerMessage() string`

GetTrackerMessage returns the TrackerMessage field if non-nil, zero value otherwise.

### GetTrackerMessageOk

`func (o *GetTorrentListResponseDataInner) GetTrackerMessageOk() (*string, bool)`

GetTrackerMessageOk returns a tuple with the TrackerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerMessage

`func (o *GetTorrentListResponseDataInner) SetTrackerMessage(v string)`

SetTrackerMessage sets TrackerMessage field to given value.

### HasTrackerMessage

`func (o *GetTorrentListResponseDataInner) HasTrackerMessage() bool`

HasTrackerMessage returns a boolean if a field has been set.

### GetCachedAt

`func (o *GetTorrentListResponseDataInner) GetCachedAt() string`

GetCachedAt returns the CachedAt field if non-nil, zero value otherwise.

### GetCachedAtOk

`func (o *GetTorrentListResponseDataInner) GetCachedAtOk() (*string, bool)`

GetCachedAtOk returns a tuple with the CachedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedAt

`func (o *GetTorrentListResponseDataInner) SetCachedAt(v string)`

SetCachedAt sets CachedAt field to given value.

### HasCachedAt

`func (o *GetTorrentListResponseDataInner) HasCachedAt() bool`

HasCachedAt returns a boolean if a field has been set.

### GetPrivate

`func (o *GetTorrentListResponseDataInner) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GetTorrentListResponseDataInner) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GetTorrentListResponseDataInner) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GetTorrentListResponseDataInner) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetAlternativeHashes

`func (o *GetTorrentListResponseDataInner) GetAlternativeHashes() []string`

GetAlternativeHashes returns the AlternativeHashes field if non-nil, zero value otherwise.

### GetAlternativeHashesOk

`func (o *GetTorrentListResponseDataInner) GetAlternativeHashesOk() (*[]string, bool)`

GetAlternativeHashesOk returns a tuple with the AlternativeHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeHashes

`func (o *GetTorrentListResponseDataInner) SetAlternativeHashes(v []string)`

SetAlternativeHashes sets AlternativeHashes field to given value.

### HasAlternativeHashes

`func (o *GetTorrentListResponseDataInner) HasAlternativeHashes() bool`

HasAlternativeHashes returns a boolean if a field has been set.

### GetTags

`func (o *GetTorrentListResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetTorrentListResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetTorrentListResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetTorrentListResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAirlocked

`func (o *GetTorrentListResponseDataInner) GetAirlocked() bool`

GetAirlocked returns the Airlocked field if non-nil, zero value otherwise.

### GetAirlockedOk

`func (o *GetTorrentListResponseDataInner) GetAirlockedOk() (*bool, bool)`

GetAirlockedOk returns a tuple with the Airlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlocked

`func (o *GetTorrentListResponseDataInner) SetAirlocked(v bool)`

SetAirlocked sets Airlocked field to given value.

### HasAirlocked

`func (o *GetTorrentListResponseDataInner) HasAirlocked() bool`

HasAirlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


