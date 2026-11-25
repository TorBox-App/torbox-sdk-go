# CreateStreamResponseDataSearchMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Titles** | Pointer to **[]string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Actors** | Pointer to **[]string** |  | [optional] 
**ReleaseYears** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**Characters** | Pointer to **[]string** |  | [optional] 
**GlobalID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **float32** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**ContentRating** | Pointer to **string** |  | [optional] 
**Trailer** | Pointer to **map[string]interface{}** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**IsAdult** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ReleasedDate** | Pointer to **string** |  | [optional] 
**EpisodesNumber** | Pointer to **int64** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**TitlesFull** | Pointer to [**[]CreateStreamResponseDataSearchMetadataTitlesFullInner**](CreateStreamResponseDataSearchMetadataTitlesFullInner.md) |  | [optional] 
**Backdrop** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateStreamResponseDataSearchMetadata

`func NewCreateStreamResponseDataSearchMetadata() *CreateStreamResponseDataSearchMetadata`

NewCreateStreamResponseDataSearchMetadata instantiates a new CreateStreamResponseDataSearchMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStreamResponseDataSearchMetadataWithDefaults

`func NewCreateStreamResponseDataSearchMetadataWithDefaults() *CreateStreamResponseDataSearchMetadata`

NewCreateStreamResponseDataSearchMetadataWithDefaults instantiates a new CreateStreamResponseDataSearchMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateStreamResponseDataSearchMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateStreamResponseDataSearchMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateStreamResponseDataSearchMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateStreamResponseDataSearchMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitles

`func (o *CreateStreamResponseDataSearchMetadata) GetTitles() []string`

GetTitles returns the Titles field if non-nil, zero value otherwise.

### GetTitlesOk

`func (o *CreateStreamResponseDataSearchMetadata) GetTitlesOk() (*[]string, bool)`

GetTitlesOk returns a tuple with the Titles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitles

`func (o *CreateStreamResponseDataSearchMetadata) SetTitles(v []string)`

SetTitles sets Titles field to given value.

### HasTitles

`func (o *CreateStreamResponseDataSearchMetadata) HasTitles() bool`

HasTitles returns a boolean if a field has been set.

### GetKeywords

`func (o *CreateStreamResponseDataSearchMetadata) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CreateStreamResponseDataSearchMetadata) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CreateStreamResponseDataSearchMetadata) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CreateStreamResponseDataSearchMetadata) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetGenres

`func (o *CreateStreamResponseDataSearchMetadata) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *CreateStreamResponseDataSearchMetadata) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *CreateStreamResponseDataSearchMetadata) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *CreateStreamResponseDataSearchMetadata) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetActors

`func (o *CreateStreamResponseDataSearchMetadata) GetActors() []string`

GetActors returns the Actors field if non-nil, zero value otherwise.

### GetActorsOk

`func (o *CreateStreamResponseDataSearchMetadata) GetActorsOk() (*[]string, bool)`

GetActorsOk returns a tuple with the Actors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActors

`func (o *CreateStreamResponseDataSearchMetadata) SetActors(v []string)`

SetActors sets Actors field to given value.

### HasActors

`func (o *CreateStreamResponseDataSearchMetadata) HasActors() bool`

HasActors returns a boolean if a field has been set.

### GetReleaseYears

`func (o *CreateStreamResponseDataSearchMetadata) GetReleaseYears() string`

GetReleaseYears returns the ReleaseYears field if non-nil, zero value otherwise.

### GetReleaseYearsOk

`func (o *CreateStreamResponseDataSearchMetadata) GetReleaseYearsOk() (*string, bool)`

GetReleaseYearsOk returns a tuple with the ReleaseYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseYears

`func (o *CreateStreamResponseDataSearchMetadata) SetReleaseYears(v string)`

SetReleaseYears sets ReleaseYears field to given value.

### HasReleaseYears

`func (o *CreateStreamResponseDataSearchMetadata) HasReleaseYears() bool`

HasReleaseYears returns a boolean if a field has been set.

### GetMediaType

`func (o *CreateStreamResponseDataSearchMetadata) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CreateStreamResponseDataSearchMetadata) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CreateStreamResponseDataSearchMetadata) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *CreateStreamResponseDataSearchMetadata) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetCharacters

`func (o *CreateStreamResponseDataSearchMetadata) GetCharacters() []string`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *CreateStreamResponseDataSearchMetadata) GetCharactersOk() (*[]string, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *CreateStreamResponseDataSearchMetadata) SetCharacters(v []string)`

SetCharacters sets Characters field to given value.

### HasCharacters

`func (o *CreateStreamResponseDataSearchMetadata) HasCharacters() bool`

HasCharacters returns a boolean if a field has been set.

### GetGlobalID

`func (o *CreateStreamResponseDataSearchMetadata) GetGlobalID() string`

GetGlobalID returns the GlobalID field if non-nil, zero value otherwise.

### GetGlobalIDOk

`func (o *CreateStreamResponseDataSearchMetadata) GetGlobalIDOk() (*string, bool)`

GetGlobalIDOk returns a tuple with the GlobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalID

`func (o *CreateStreamResponseDataSearchMetadata) SetGlobalID(v string)`

SetGlobalID sets GlobalID field to given value.

### HasGlobalID

`func (o *CreateStreamResponseDataSearchMetadata) HasGlobalID() bool`

HasGlobalID returns a boolean if a field has been set.

### GetId

`func (o *CreateStreamResponseDataSearchMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateStreamResponseDataSearchMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateStreamResponseDataSearchMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateStreamResponseDataSearchMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLink

`func (o *CreateStreamResponseDataSearchMetadata) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CreateStreamResponseDataSearchMetadata) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CreateStreamResponseDataSearchMetadata) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CreateStreamResponseDataSearchMetadata) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDescription

`func (o *CreateStreamResponseDataSearchMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStreamResponseDataSearchMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStreamResponseDataSearchMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStreamResponseDataSearchMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRating

`func (o *CreateStreamResponseDataSearchMetadata) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *CreateStreamResponseDataSearchMetadata) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *CreateStreamResponseDataSearchMetadata) SetRating(v float32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *CreateStreamResponseDataSearchMetadata) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLanguages

`func (o *CreateStreamResponseDataSearchMetadata) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CreateStreamResponseDataSearchMetadata) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CreateStreamResponseDataSearchMetadata) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CreateStreamResponseDataSearchMetadata) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetContentRating

`func (o *CreateStreamResponseDataSearchMetadata) GetContentRating() string`

GetContentRating returns the ContentRating field if non-nil, zero value otherwise.

### GetContentRatingOk

`func (o *CreateStreamResponseDataSearchMetadata) GetContentRatingOk() (*string, bool)`

GetContentRatingOk returns a tuple with the ContentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRating

`func (o *CreateStreamResponseDataSearchMetadata) SetContentRating(v string)`

SetContentRating sets ContentRating field to given value.

### HasContentRating

`func (o *CreateStreamResponseDataSearchMetadata) HasContentRating() bool`

HasContentRating returns a boolean if a field has been set.

### GetTrailer

`func (o *CreateStreamResponseDataSearchMetadata) GetTrailer() map[string]interface{}`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *CreateStreamResponseDataSearchMetadata) GetTrailerOk() (*map[string]interface{}, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *CreateStreamResponseDataSearchMetadata) SetTrailer(v map[string]interface{})`

SetTrailer sets Trailer field to given value.

### HasTrailer

`func (o *CreateStreamResponseDataSearchMetadata) HasTrailer() bool`

HasTrailer returns a boolean if a field has been set.

### GetImage

`func (o *CreateStreamResponseDataSearchMetadata) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateStreamResponseDataSearchMetadata) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateStreamResponseDataSearchMetadata) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateStreamResponseDataSearchMetadata) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIsAdult

`func (o *CreateStreamResponseDataSearchMetadata) GetIsAdult() bool`

GetIsAdult returns the IsAdult field if non-nil, zero value otherwise.

### GetIsAdultOk

`func (o *CreateStreamResponseDataSearchMetadata) GetIsAdultOk() (*bool, bool)`

GetIsAdultOk returns a tuple with the IsAdult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdult

`func (o *CreateStreamResponseDataSearchMetadata) SetIsAdult(v bool)`

SetIsAdult sets IsAdult field to given value.

### HasIsAdult

`func (o *CreateStreamResponseDataSearchMetadata) HasIsAdult() bool`

HasIsAdult returns a boolean if a field has been set.

### GetType

`func (o *CreateStreamResponseDataSearchMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateStreamResponseDataSearchMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateStreamResponseDataSearchMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateStreamResponseDataSearchMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReleasedDate

`func (o *CreateStreamResponseDataSearchMetadata) GetReleasedDate() string`

GetReleasedDate returns the ReleasedDate field if non-nil, zero value otherwise.

### GetReleasedDateOk

`func (o *CreateStreamResponseDataSearchMetadata) GetReleasedDateOk() (*string, bool)`

GetReleasedDateOk returns a tuple with the ReleasedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedDate

`func (o *CreateStreamResponseDataSearchMetadata) SetReleasedDate(v string)`

SetReleasedDate sets ReleasedDate field to given value.

### HasReleasedDate

`func (o *CreateStreamResponseDataSearchMetadata) HasReleasedDate() bool`

HasReleasedDate returns a boolean if a field has been set.

### GetEpisodesNumber

`func (o *CreateStreamResponseDataSearchMetadata) GetEpisodesNumber() int64`

GetEpisodesNumber returns the EpisodesNumber field if non-nil, zero value otherwise.

### GetEpisodesNumberOk

`func (o *CreateStreamResponseDataSearchMetadata) GetEpisodesNumberOk() (*int64, bool)`

GetEpisodesNumberOk returns a tuple with the EpisodesNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodesNumber

`func (o *CreateStreamResponseDataSearchMetadata) SetEpisodesNumber(v int64)`

SetEpisodesNumber sets EpisodesNumber field to given value.

### HasEpisodesNumber

`func (o *CreateStreamResponseDataSearchMetadata) HasEpisodesNumber() bool`

HasEpisodesNumber returns a boolean if a field has been set.

### GetRuntime

`func (o *CreateStreamResponseDataSearchMetadata) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *CreateStreamResponseDataSearchMetadata) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *CreateStreamResponseDataSearchMetadata) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *CreateStreamResponseDataSearchMetadata) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTitlesFull

`func (o *CreateStreamResponseDataSearchMetadata) GetTitlesFull() []CreateStreamResponseDataSearchMetadataTitlesFullInner`

GetTitlesFull returns the TitlesFull field if non-nil, zero value otherwise.

### GetTitlesFullOk

`func (o *CreateStreamResponseDataSearchMetadata) GetTitlesFullOk() (*[]CreateStreamResponseDataSearchMetadataTitlesFullInner, bool)`

GetTitlesFullOk returns a tuple with the TitlesFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitlesFull

`func (o *CreateStreamResponseDataSearchMetadata) SetTitlesFull(v []CreateStreamResponseDataSearchMetadataTitlesFullInner)`

SetTitlesFull sets TitlesFull field to given value.

### HasTitlesFull

`func (o *CreateStreamResponseDataSearchMetadata) HasTitlesFull() bool`

HasTitlesFull returns a boolean if a field has been set.

### GetBackdrop

`func (o *CreateStreamResponseDataSearchMetadata) GetBackdrop() string`

GetBackdrop returns the Backdrop field if non-nil, zero value otherwise.

### GetBackdropOk

`func (o *CreateStreamResponseDataSearchMetadata) GetBackdropOk() (*string, bool)`

GetBackdropOk returns a tuple with the Backdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdrop

`func (o *CreateStreamResponseDataSearchMetadata) SetBackdrop(v string)`

SetBackdrop sets Backdrop field to given value.

### HasBackdrop

`func (o *CreateStreamResponseDataSearchMetadata) HasBackdrop() bool`

HasBackdrop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


