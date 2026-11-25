# CreateStreamResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**FileId** | Pointer to **int64** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**S3Path** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Mimetype** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**OpenSubtitlesHash** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**UserToken** | Pointer to **string** |  | [optional] 
**WebdavUrl** | Pointer to **string** |  | [optional] 
**IntroInformation** | Pointer to [**CreateStreamResponseDataIntroInformation**](CreateStreamResponseDataIntroInformation.md) |  | [optional] 
**ScrobblingEnabled** | Pointer to **bool** |  | [optional] 
**HlsUrl** | Pointer to **string** |  | [optional] 
**PresignedToken** | Pointer to **string** |  | [optional] 
**SubtitleIndex** | Pointer to **int64** |  | [optional] 
**AudioIndex** | Pointer to **int64** |  | [optional] 
**ResolutionIndex** | Pointer to **int64** |  | [optional] 
**IsTranscoding** | Pointer to **bool** |  | [optional] 
**NeedsTranscoding** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**CreateStreamResponseDataMetadata**](CreateStreamResponseDataMetadata.md) |  | [optional] 
**SearchMetadata** | Pointer to [**CreateStreamResponseDataSearchMetadata**](CreateStreamResponseDataSearchMetadata.md) |  | [optional] 

## Methods

### NewCreateStreamResponseData

`func NewCreateStreamResponseData() *CreateStreamResponseData`

NewCreateStreamResponseData instantiates a new CreateStreamResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStreamResponseDataWithDefaults

`func NewCreateStreamResponseDataWithDefaults() *CreateStreamResponseData`

NewCreateStreamResponseDataWithDefaults instantiates a new CreateStreamResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateStreamResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStreamResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStreamResponseData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateStreamResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileId

`func (o *CreateStreamResponseData) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *CreateStreamResponseData) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *CreateStreamResponseData) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *CreateStreamResponseData) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetAuthId

`func (o *CreateStreamResponseData) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *CreateStreamResponseData) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *CreateStreamResponseData) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *CreateStreamResponseData) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateStreamResponseData) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateStreamResponseData) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateStreamResponseData) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateStreamResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *CreateStreamResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStreamResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStreamResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStreamResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *CreateStreamResponseData) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CreateStreamResponseData) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CreateStreamResponseData) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *CreateStreamResponseData) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetS3Path

`func (o *CreateStreamResponseData) GetS3Path() string`

GetS3Path returns the S3Path field if non-nil, zero value otherwise.

### GetS3PathOk

`func (o *CreateStreamResponseData) GetS3PathOk() (*string, bool)`

GetS3PathOk returns a tuple with the S3Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Path

`func (o *CreateStreamResponseData) SetS3Path(v string)`

SetS3Path sets S3Path field to given value.

### HasS3Path

`func (o *CreateStreamResponseData) HasS3Path() bool`

HasS3Path returns a boolean if a field has been set.

### GetSize

`func (o *CreateStreamResponseData) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateStreamResponseData) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateStreamResponseData) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateStreamResponseData) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMimetype

`func (o *CreateStreamResponseData) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *CreateStreamResponseData) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *CreateStreamResponseData) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *CreateStreamResponseData) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetHash

`func (o *CreateStreamResponseData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CreateStreamResponseData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CreateStreamResponseData) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CreateStreamResponseData) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetOpenSubtitlesHash

`func (o *CreateStreamResponseData) GetOpenSubtitlesHash() string`

GetOpenSubtitlesHash returns the OpenSubtitlesHash field if non-nil, zero value otherwise.

### GetOpenSubtitlesHashOk

`func (o *CreateStreamResponseData) GetOpenSubtitlesHashOk() (*string, bool)`

GetOpenSubtitlesHashOk returns a tuple with the OpenSubtitlesHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSubtitlesHash

`func (o *CreateStreamResponseData) SetOpenSubtitlesHash(v string)`

SetOpenSubtitlesHash sets OpenSubtitlesHash field to given value.

### HasOpenSubtitlesHash

`func (o *CreateStreamResponseData) HasOpenSubtitlesHash() bool`

HasOpenSubtitlesHash returns a boolean if a field has been set.

### GetType

`func (o *CreateStreamResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStreamResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStreamResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateStreamResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegion

`func (o *CreateStreamResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateStreamResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateStreamResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateStreamResponseData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDomain

`func (o *CreateStreamResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateStreamResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateStreamResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateStreamResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUserToken

`func (o *CreateStreamResponseData) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CreateStreamResponseData) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CreateStreamResponseData) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *CreateStreamResponseData) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetWebdavUrl

`func (o *CreateStreamResponseData) GetWebdavUrl() string`

GetWebdavUrl returns the WebdavUrl field if non-nil, zero value otherwise.

### GetWebdavUrlOk

`func (o *CreateStreamResponseData) GetWebdavUrlOk() (*string, bool)`

GetWebdavUrlOk returns a tuple with the WebdavUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdavUrl

`func (o *CreateStreamResponseData) SetWebdavUrl(v string)`

SetWebdavUrl sets WebdavUrl field to given value.

### HasWebdavUrl

`func (o *CreateStreamResponseData) HasWebdavUrl() bool`

HasWebdavUrl returns a boolean if a field has been set.

### GetIntroInformation

`func (o *CreateStreamResponseData) GetIntroInformation() CreateStreamResponseDataIntroInformation`

GetIntroInformation returns the IntroInformation field if non-nil, zero value otherwise.

### GetIntroInformationOk

`func (o *CreateStreamResponseData) GetIntroInformationOk() (*CreateStreamResponseDataIntroInformation, bool)`

GetIntroInformationOk returns a tuple with the IntroInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroInformation

`func (o *CreateStreamResponseData) SetIntroInformation(v CreateStreamResponseDataIntroInformation)`

SetIntroInformation sets IntroInformation field to given value.

### HasIntroInformation

`func (o *CreateStreamResponseData) HasIntroInformation() bool`

HasIntroInformation returns a boolean if a field has been set.

### GetScrobblingEnabled

`func (o *CreateStreamResponseData) GetScrobblingEnabled() bool`

GetScrobblingEnabled returns the ScrobblingEnabled field if non-nil, zero value otherwise.

### GetScrobblingEnabledOk

`func (o *CreateStreamResponseData) GetScrobblingEnabledOk() (*bool, bool)`

GetScrobblingEnabledOk returns a tuple with the ScrobblingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrobblingEnabled

`func (o *CreateStreamResponseData) SetScrobblingEnabled(v bool)`

SetScrobblingEnabled sets ScrobblingEnabled field to given value.

### HasScrobblingEnabled

`func (o *CreateStreamResponseData) HasScrobblingEnabled() bool`

HasScrobblingEnabled returns a boolean if a field has been set.

### GetHlsUrl

`func (o *CreateStreamResponseData) GetHlsUrl() string`

GetHlsUrl returns the HlsUrl field if non-nil, zero value otherwise.

### GetHlsUrlOk

`func (o *CreateStreamResponseData) GetHlsUrlOk() (*string, bool)`

GetHlsUrlOk returns a tuple with the HlsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlsUrl

`func (o *CreateStreamResponseData) SetHlsUrl(v string)`

SetHlsUrl sets HlsUrl field to given value.

### HasHlsUrl

`func (o *CreateStreamResponseData) HasHlsUrl() bool`

HasHlsUrl returns a boolean if a field has been set.

### GetPresignedToken

`func (o *CreateStreamResponseData) GetPresignedToken() string`

GetPresignedToken returns the PresignedToken field if non-nil, zero value otherwise.

### GetPresignedTokenOk

`func (o *CreateStreamResponseData) GetPresignedTokenOk() (*string, bool)`

GetPresignedTokenOk returns a tuple with the PresignedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedToken

`func (o *CreateStreamResponseData) SetPresignedToken(v string)`

SetPresignedToken sets PresignedToken field to given value.

### HasPresignedToken

`func (o *CreateStreamResponseData) HasPresignedToken() bool`

HasPresignedToken returns a boolean if a field has been set.

### GetSubtitleIndex

`func (o *CreateStreamResponseData) GetSubtitleIndex() int64`

GetSubtitleIndex returns the SubtitleIndex field if non-nil, zero value otherwise.

### GetSubtitleIndexOk

`func (o *CreateStreamResponseData) GetSubtitleIndexOk() (*int64, bool)`

GetSubtitleIndexOk returns a tuple with the SubtitleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleIndex

`func (o *CreateStreamResponseData) SetSubtitleIndex(v int64)`

SetSubtitleIndex sets SubtitleIndex field to given value.

### HasSubtitleIndex

`func (o *CreateStreamResponseData) HasSubtitleIndex() bool`

HasSubtitleIndex returns a boolean if a field has been set.

### GetAudioIndex

`func (o *CreateStreamResponseData) GetAudioIndex() int64`

GetAudioIndex returns the AudioIndex field if non-nil, zero value otherwise.

### GetAudioIndexOk

`func (o *CreateStreamResponseData) GetAudioIndexOk() (*int64, bool)`

GetAudioIndexOk returns a tuple with the AudioIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioIndex

`func (o *CreateStreamResponseData) SetAudioIndex(v int64)`

SetAudioIndex sets AudioIndex field to given value.

### HasAudioIndex

`func (o *CreateStreamResponseData) HasAudioIndex() bool`

HasAudioIndex returns a boolean if a field has been set.

### GetResolutionIndex

`func (o *CreateStreamResponseData) GetResolutionIndex() int64`

GetResolutionIndex returns the ResolutionIndex field if non-nil, zero value otherwise.

### GetResolutionIndexOk

`func (o *CreateStreamResponseData) GetResolutionIndexOk() (*int64, bool)`

GetResolutionIndexOk returns a tuple with the ResolutionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionIndex

`func (o *CreateStreamResponseData) SetResolutionIndex(v int64)`

SetResolutionIndex sets ResolutionIndex field to given value.

### HasResolutionIndex

`func (o *CreateStreamResponseData) HasResolutionIndex() bool`

HasResolutionIndex returns a boolean if a field has been set.

### GetIsTranscoding

`func (o *CreateStreamResponseData) GetIsTranscoding() bool`

GetIsTranscoding returns the IsTranscoding field if non-nil, zero value otherwise.

### GetIsTranscodingOk

`func (o *CreateStreamResponseData) GetIsTranscodingOk() (*bool, bool)`

GetIsTranscodingOk returns a tuple with the IsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranscoding

`func (o *CreateStreamResponseData) SetIsTranscoding(v bool)`

SetIsTranscoding sets IsTranscoding field to given value.

### HasIsTranscoding

`func (o *CreateStreamResponseData) HasIsTranscoding() bool`

HasIsTranscoding returns a boolean if a field has been set.

### GetNeedsTranscoding

`func (o *CreateStreamResponseData) GetNeedsTranscoding() bool`

GetNeedsTranscoding returns the NeedsTranscoding field if non-nil, zero value otherwise.

### GetNeedsTranscodingOk

`func (o *CreateStreamResponseData) GetNeedsTranscodingOk() (*bool, bool)`

GetNeedsTranscodingOk returns a tuple with the NeedsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsTranscoding

`func (o *CreateStreamResponseData) SetNeedsTranscoding(v bool)`

SetNeedsTranscoding sets NeedsTranscoding field to given value.

### HasNeedsTranscoding

`func (o *CreateStreamResponseData) HasNeedsTranscoding() bool`

HasNeedsTranscoding returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateStreamResponseData) GetMetadata() CreateStreamResponseDataMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateStreamResponseData) GetMetadataOk() (*CreateStreamResponseDataMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateStreamResponseData) SetMetadata(v CreateStreamResponseDataMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateStreamResponseData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSearchMetadata

`func (o *CreateStreamResponseData) GetSearchMetadata() CreateStreamResponseDataSearchMetadata`

GetSearchMetadata returns the SearchMetadata field if non-nil, zero value otherwise.

### GetSearchMetadataOk

`func (o *CreateStreamResponseData) GetSearchMetadataOk() (*CreateStreamResponseDataSearchMetadata, bool)`

GetSearchMetadataOk returns a tuple with the SearchMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMetadata

`func (o *CreateStreamResponseData) SetSearchMetadata(v CreateStreamResponseDataSearchMetadata)`

SetSearchMetadata sets SearchMetadata field to given value.

### HasSearchMetadata

`func (o *CreateStreamResponseData) HasSearchMetadata() bool`

HasSearchMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


