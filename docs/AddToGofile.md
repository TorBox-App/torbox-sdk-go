# AddToGofile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the torrent, usenet download or web download to upload. It has to be one of your own downloads. Required. | 
**FileId** | Pointer to **int64** | Id of the file to upload, from the download&#39;s files. Leave out when &#x60;zip&#x60; is true. | [optional] 
**Zip** | Pointer to **bool** | Upload the whole download as one zip instead of a single file. Takes precedence over &#x60;file_id&#x60;. Optional, default false. | [optional] [default to false]
**Type** | Pointer to **string** | Which kind of download &#x60;id&#x60; refers to. One of &#x60;torrent&#x60;, &#x60;usenet&#x60; or &#x60;webdownload&#x60;. Optional, default &#x60;torrent&#x60;. | [optional] [default to "torrent"]
**GofileToken** | Pointer to **string** | Gofile API token of the account to upload to. It is issued by Gofile, not by TorBox. Optional. | [optional] 

## Methods

### NewAddToGofile

`func NewAddToGofile(id int64, ) *AddToGofile`

NewAddToGofile instantiates a new AddToGofile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddToGofileWithDefaults

`func NewAddToGofileWithDefaults() *AddToGofile`

NewAddToGofileWithDefaults instantiates a new AddToGofile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddToGofile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddToGofile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddToGofile) SetId(v int64)`

SetId sets Id field to given value.


### GetFileId

`func (o *AddToGofile) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AddToGofile) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AddToGofile) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AddToGofile) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetZip

`func (o *AddToGofile) GetZip() bool`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddToGofile) GetZipOk() (*bool, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddToGofile) SetZip(v bool)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddToGofile) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetType

`func (o *AddToGofile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddToGofile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddToGofile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddToGofile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGofileToken

`func (o *AddToGofile) GetGofileToken() string`

GetGofileToken returns the GofileToken field if non-nil, zero value otherwise.

### GetGofileTokenOk

`func (o *AddToGofile) GetGofileTokenOk() (*string, bool)`

GetGofileTokenOk returns a tuple with the GofileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGofileToken

`func (o *AddToGofile) SetGofileToken(v string)`

SetGofileToken sets GofileToken field to given value.

### HasGofileToken

`func (o *AddToGofile) HasGofileToken() bool`

HasGofileToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


