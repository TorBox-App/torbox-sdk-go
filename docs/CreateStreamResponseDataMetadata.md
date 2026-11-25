# CreateStreamResponseDataMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Video** | Pointer to [**CreateStreamResponseDataMetadataVideo**](CreateStreamResponseDataMetadataVideo.md) |  | [optional] 
**Audios** | Pointer to [**[]CreateStreamResponseDataMetadataAudiosInner**](CreateStreamResponseDataMetadataAudiosInner.md) |  | [optional] 
**Subtitles** | Pointer to **[]interface{}** |  | [optional] 
**Thumbnail** | Pointer to **string** |  | [optional] 
**Chapters** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateStreamResponseDataMetadata

`func NewCreateStreamResponseDataMetadata() *CreateStreamResponseDataMetadata`

NewCreateStreamResponseDataMetadata instantiates a new CreateStreamResponseDataMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStreamResponseDataMetadataWithDefaults

`func NewCreateStreamResponseDataMetadataWithDefaults() *CreateStreamResponseDataMetadata`

NewCreateStreamResponseDataMetadataWithDefaults instantiates a new CreateStreamResponseDataMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideo

`func (o *CreateStreamResponseDataMetadata) GetVideo() CreateStreamResponseDataMetadataVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreateStreamResponseDataMetadata) GetVideoOk() (*CreateStreamResponseDataMetadataVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreateStreamResponseDataMetadata) SetVideo(v CreateStreamResponseDataMetadataVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreateStreamResponseDataMetadata) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAudios

`func (o *CreateStreamResponseDataMetadata) GetAudios() []CreateStreamResponseDataMetadataAudiosInner`

GetAudios returns the Audios field if non-nil, zero value otherwise.

### GetAudiosOk

`func (o *CreateStreamResponseDataMetadata) GetAudiosOk() (*[]CreateStreamResponseDataMetadataAudiosInner, bool)`

GetAudiosOk returns a tuple with the Audios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudios

`func (o *CreateStreamResponseDataMetadata) SetAudios(v []CreateStreamResponseDataMetadataAudiosInner)`

SetAudios sets Audios field to given value.

### HasAudios

`func (o *CreateStreamResponseDataMetadata) HasAudios() bool`

HasAudios returns a boolean if a field has been set.

### GetSubtitles

`func (o *CreateStreamResponseDataMetadata) GetSubtitles() []interface{}`

GetSubtitles returns the Subtitles field if non-nil, zero value otherwise.

### GetSubtitlesOk

`func (o *CreateStreamResponseDataMetadata) GetSubtitlesOk() (*[]interface{}, bool)`

GetSubtitlesOk returns a tuple with the Subtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitles

`func (o *CreateStreamResponseDataMetadata) SetSubtitles(v []interface{})`

SetSubtitles sets Subtitles field to given value.

### HasSubtitles

`func (o *CreateStreamResponseDataMetadata) HasSubtitles() bool`

HasSubtitles returns a boolean if a field has been set.

### GetThumbnail

`func (o *CreateStreamResponseDataMetadata) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *CreateStreamResponseDataMetadata) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *CreateStreamResponseDataMetadata) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *CreateStreamResponseDataMetadata) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetChapters

`func (o *CreateStreamResponseDataMetadata) GetChapters() string`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *CreateStreamResponseDataMetadata) GetChaptersOk() (*string, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *CreateStreamResponseDataMetadata) SetChapters(v string)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *CreateStreamResponseDataMetadata) HasChapters() bool`

HasChapters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


