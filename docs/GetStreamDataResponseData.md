# GetStreamDataResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HlsUrl** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**PresignedToken** | Pointer to **string** |  | [optional] 
**SubtitleIndex** | Pointer to **int64** |  | [optional] 
**AudioIndex** | Pointer to **int64** |  | [optional] 
**ResolutionIndex** | Pointer to **int64** |  | [optional] 
**FileToken** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**IsTranscoding** | Pointer to **bool** |  | [optional] 
**NeedsTranscoding** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**CreateStreamResponseDataMetadata**](CreateStreamResponseDataMetadata.md) |  | [optional] 
**SearchMetadata** | Pointer to [**CreateStreamResponseDataSearchMetadata**](CreateStreamResponseDataSearchMetadata.md) |  | [optional] 

## Methods

### NewGetStreamDataResponseData

`func NewGetStreamDataResponseData() *GetStreamDataResponseData`

NewGetStreamDataResponseData instantiates a new GetStreamDataResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStreamDataResponseDataWithDefaults

`func NewGetStreamDataResponseDataWithDefaults() *GetStreamDataResponseData`

NewGetStreamDataResponseDataWithDefaults instantiates a new GetStreamDataResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHlsUrl

`func (o *GetStreamDataResponseData) GetHlsUrl() string`

GetHlsUrl returns the HlsUrl field if non-nil, zero value otherwise.

### GetHlsUrlOk

`func (o *GetStreamDataResponseData) GetHlsUrlOk() (*string, bool)`

GetHlsUrlOk returns a tuple with the HlsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlsUrl

`func (o *GetStreamDataResponseData) SetHlsUrl(v string)`

SetHlsUrl sets HlsUrl field to given value.

### HasHlsUrl

`func (o *GetStreamDataResponseData) HasHlsUrl() bool`

HasHlsUrl returns a boolean if a field has been set.

### GetDomain

`func (o *GetStreamDataResponseData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetStreamDataResponseData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetStreamDataResponseData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetStreamDataResponseData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPresignedToken

`func (o *GetStreamDataResponseData) GetPresignedToken() string`

GetPresignedToken returns the PresignedToken field if non-nil, zero value otherwise.

### GetPresignedTokenOk

`func (o *GetStreamDataResponseData) GetPresignedTokenOk() (*string, bool)`

GetPresignedTokenOk returns a tuple with the PresignedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedToken

`func (o *GetStreamDataResponseData) SetPresignedToken(v string)`

SetPresignedToken sets PresignedToken field to given value.

### HasPresignedToken

`func (o *GetStreamDataResponseData) HasPresignedToken() bool`

HasPresignedToken returns a boolean if a field has been set.

### GetSubtitleIndex

`func (o *GetStreamDataResponseData) GetSubtitleIndex() int64`

GetSubtitleIndex returns the SubtitleIndex field if non-nil, zero value otherwise.

### GetSubtitleIndexOk

`func (o *GetStreamDataResponseData) GetSubtitleIndexOk() (*int64, bool)`

GetSubtitleIndexOk returns a tuple with the SubtitleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleIndex

`func (o *GetStreamDataResponseData) SetSubtitleIndex(v int64)`

SetSubtitleIndex sets SubtitleIndex field to given value.

### HasSubtitleIndex

`func (o *GetStreamDataResponseData) HasSubtitleIndex() bool`

HasSubtitleIndex returns a boolean if a field has been set.

### GetAudioIndex

`func (o *GetStreamDataResponseData) GetAudioIndex() int64`

GetAudioIndex returns the AudioIndex field if non-nil, zero value otherwise.

### GetAudioIndexOk

`func (o *GetStreamDataResponseData) GetAudioIndexOk() (*int64, bool)`

GetAudioIndexOk returns a tuple with the AudioIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioIndex

`func (o *GetStreamDataResponseData) SetAudioIndex(v int64)`

SetAudioIndex sets AudioIndex field to given value.

### HasAudioIndex

`func (o *GetStreamDataResponseData) HasAudioIndex() bool`

HasAudioIndex returns a boolean if a field has been set.

### GetResolutionIndex

`func (o *GetStreamDataResponseData) GetResolutionIndex() int64`

GetResolutionIndex returns the ResolutionIndex field if non-nil, zero value otherwise.

### GetResolutionIndexOk

`func (o *GetStreamDataResponseData) GetResolutionIndexOk() (*int64, bool)`

GetResolutionIndexOk returns a tuple with the ResolutionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionIndex

`func (o *GetStreamDataResponseData) SetResolutionIndex(v int64)`

SetResolutionIndex sets ResolutionIndex field to given value.

### HasResolutionIndex

`func (o *GetStreamDataResponseData) HasResolutionIndex() bool`

HasResolutionIndex returns a boolean if a field has been set.

### GetFileToken

`func (o *GetStreamDataResponseData) GetFileToken() string`

GetFileToken returns the FileToken field if non-nil, zero value otherwise.

### GetFileTokenOk

`func (o *GetStreamDataResponseData) GetFileTokenOk() (*string, bool)`

GetFileTokenOk returns a tuple with the FileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileToken

`func (o *GetStreamDataResponseData) SetFileToken(v string)`

SetFileToken sets FileToken field to given value.

### HasFileToken

`func (o *GetStreamDataResponseData) HasFileToken() bool`

HasFileToken returns a boolean if a field has been set.

### GetToken

`func (o *GetStreamDataResponseData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetStreamDataResponseData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetStreamDataResponseData) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetStreamDataResponseData) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetIsTranscoding

`func (o *GetStreamDataResponseData) GetIsTranscoding() bool`

GetIsTranscoding returns the IsTranscoding field if non-nil, zero value otherwise.

### GetIsTranscodingOk

`func (o *GetStreamDataResponseData) GetIsTranscodingOk() (*bool, bool)`

GetIsTranscodingOk returns a tuple with the IsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranscoding

`func (o *GetStreamDataResponseData) SetIsTranscoding(v bool)`

SetIsTranscoding sets IsTranscoding field to given value.

### HasIsTranscoding

`func (o *GetStreamDataResponseData) HasIsTranscoding() bool`

HasIsTranscoding returns a boolean if a field has been set.

### GetNeedsTranscoding

`func (o *GetStreamDataResponseData) GetNeedsTranscoding() bool`

GetNeedsTranscoding returns the NeedsTranscoding field if non-nil, zero value otherwise.

### GetNeedsTranscodingOk

`func (o *GetStreamDataResponseData) GetNeedsTranscodingOk() (*bool, bool)`

GetNeedsTranscodingOk returns a tuple with the NeedsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsTranscoding

`func (o *GetStreamDataResponseData) SetNeedsTranscoding(v bool)`

SetNeedsTranscoding sets NeedsTranscoding field to given value.

### HasNeedsTranscoding

`func (o *GetStreamDataResponseData) HasNeedsTranscoding() bool`

HasNeedsTranscoding returns a boolean if a field has been set.

### GetMetadata

`func (o *GetStreamDataResponseData) GetMetadata() CreateStreamResponseDataMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetStreamDataResponseData) GetMetadataOk() (*CreateStreamResponseDataMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetStreamDataResponseData) SetMetadata(v CreateStreamResponseDataMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetStreamDataResponseData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSearchMetadata

`func (o *GetStreamDataResponseData) GetSearchMetadata() CreateStreamResponseDataSearchMetadata`

GetSearchMetadata returns the SearchMetadata field if non-nil, zero value otherwise.

### GetSearchMetadataOk

`func (o *GetStreamDataResponseData) GetSearchMetadataOk() (*CreateStreamResponseDataSearchMetadata, bool)`

GetSearchMetadataOk returns a tuple with the SearchMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMetadata

`func (o *GetStreamDataResponseData) SetSearchMetadata(v CreateStreamResponseDataSearchMetadata)`

SetSearchMetadata sets SearchMetadata field to given value.

### HasSearchMetadata

`func (o *GetStreamDataResponseData) HasSearchMetadata() bool`

HasSearchMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


