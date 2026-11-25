# AddToPixeldrain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the torrent, usenet download or web download to upload. It has to be one of your own downloads. Required. | 
**FileId** | Pointer to **int64** | Id of the file to upload, from the download&#39;s files. Leave out when &#x60;zip&#x60; is true. | [optional] 
**Zip** | Pointer to **bool** | Upload the whole download as one zip instead of a single file. Takes precedence over &#x60;file_id&#x60;. Optional, default false. | [optional] [default to false]
**Type** | Pointer to **string** | Which kind of download &#x60;id&#x60; refers to. One of &#x60;torrent&#x60;, &#x60;usenet&#x60; or &#x60;webdownload&#x60;. Optional, default &#x60;torrent&#x60;. | [optional] [default to "torrent"]
**PixeldrainToken** | Pointer to **string** | Pixeldrain API key of the account to upload to. It is issued by Pixeldrain, not by TorBox. The route accepts a request without it, but the job fails without it. | [optional] 

## Methods

### NewAddToPixeldrain

`func NewAddToPixeldrain(id int64, ) *AddToPixeldrain`

NewAddToPixeldrain instantiates a new AddToPixeldrain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddToPixeldrainWithDefaults

`func NewAddToPixeldrainWithDefaults() *AddToPixeldrain`

NewAddToPixeldrainWithDefaults instantiates a new AddToPixeldrain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddToPixeldrain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddToPixeldrain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddToPixeldrain) SetId(v int64)`

SetId sets Id field to given value.


### GetFileId

`func (o *AddToPixeldrain) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AddToPixeldrain) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AddToPixeldrain) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AddToPixeldrain) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetZip

`func (o *AddToPixeldrain) GetZip() bool`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddToPixeldrain) GetZipOk() (*bool, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddToPixeldrain) SetZip(v bool)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddToPixeldrain) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetType

`func (o *AddToPixeldrain) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddToPixeldrain) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddToPixeldrain) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddToPixeldrain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPixeldrainToken

`func (o *AddToPixeldrain) GetPixeldrainToken() string`

GetPixeldrainToken returns the PixeldrainToken field if non-nil, zero value otherwise.

### GetPixeldrainTokenOk

`func (o *AddToPixeldrain) GetPixeldrainTokenOk() (*string, bool)`

GetPixeldrainTokenOk returns a tuple with the PixeldrainToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixeldrainToken

`func (o *AddToPixeldrain) SetPixeldrainToken(v string)`

SetPixeldrainToken sets PixeldrainToken field to given value.

### HasPixeldrainToken

`func (o *AddToPixeldrain) HasPixeldrainToken() bool`

HasPixeldrainToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


