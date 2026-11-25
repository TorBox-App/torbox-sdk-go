# AddToDropbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the torrent, usenet download or web download. Still required by the route, but no longer used. | 
**FileId** | Pointer to **int64** | Id of the file inside the download. No longer used. | [optional] 
**Zip** | Pointer to **bool** | Whether to upload the whole download as one zip. No longer used. | [optional] [default to false]
**Type** | Pointer to **string** | Which kind of download &#x60;id&#x60; refers to. One of &#x60;torrent&#x60;, &#x60;usenet&#x60; or &#x60;webdownload&#x60;. Optional, default &#x60;torrent&#x60;. Still checked before the route answers. | [optional] [default to "torrent"]
**DropboxToken** | **string** | Dropbox OAuth access token. Still required by the route, but no longer used. | 

## Methods

### NewAddToDropbox

`func NewAddToDropbox(id int64, dropboxToken string, ) *AddToDropbox`

NewAddToDropbox instantiates a new AddToDropbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddToDropboxWithDefaults

`func NewAddToDropboxWithDefaults() *AddToDropbox`

NewAddToDropboxWithDefaults instantiates a new AddToDropbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddToDropbox) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddToDropbox) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddToDropbox) SetId(v int64)`

SetId sets Id field to given value.


### GetFileId

`func (o *AddToDropbox) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AddToDropbox) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AddToDropbox) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AddToDropbox) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetZip

`func (o *AddToDropbox) GetZip() bool`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddToDropbox) GetZipOk() (*bool, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddToDropbox) SetZip(v bool)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddToDropbox) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetType

`func (o *AddToDropbox) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddToDropbox) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddToDropbox) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddToDropbox) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDropboxToken

`func (o *AddToDropbox) GetDropboxToken() string`

GetDropboxToken returns the DropboxToken field if non-nil, zero value otherwise.

### GetDropboxTokenOk

`func (o *AddToDropbox) GetDropboxTokenOk() (*string, bool)`

GetDropboxTokenOk returns a tuple with the DropboxToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropboxToken

`func (o *AddToDropbox) SetDropboxToken(v string)`

SetDropboxToken sets DropboxToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


