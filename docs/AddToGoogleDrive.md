# AddToGoogleDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the torrent, usenet download or web download to upload. It has to be one of your own downloads. Required. | 
**FileId** | Pointer to **int64** | Id of the file to upload, from the download&#39;s files. Leave out when &#x60;zip&#x60; is true. | [optional] 
**Zip** | Pointer to **bool** | Upload the whole download as one zip instead of a single file. Takes precedence over &#x60;file_id&#x60;. Optional, default false. | [optional] [default to false]
**Type** | Pointer to **string** | Which kind of download &#x60;id&#x60; refers to. One of &#x60;torrent&#x60;, &#x60;usenet&#x60; or &#x60;webdownload&#x60;. Optional, default &#x60;torrent&#x60;. | [optional] [default to "torrent"]
**GoogleToken** | **string** | Google OAuth access token for the Drive to upload to. It is issued by Google, not by TorBox. Required. | 

## Methods

### NewAddToGoogleDrive

`func NewAddToGoogleDrive(id int64, googleToken string, ) *AddToGoogleDrive`

NewAddToGoogleDrive instantiates a new AddToGoogleDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddToGoogleDriveWithDefaults

`func NewAddToGoogleDriveWithDefaults() *AddToGoogleDrive`

NewAddToGoogleDriveWithDefaults instantiates a new AddToGoogleDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddToGoogleDrive) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddToGoogleDrive) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddToGoogleDrive) SetId(v int64)`

SetId sets Id field to given value.


### GetFileId

`func (o *AddToGoogleDrive) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AddToGoogleDrive) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AddToGoogleDrive) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AddToGoogleDrive) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetZip

`func (o *AddToGoogleDrive) GetZip() bool`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddToGoogleDrive) GetZipOk() (*bool, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddToGoogleDrive) SetZip(v bool)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddToGoogleDrive) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetType

`func (o *AddToGoogleDrive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddToGoogleDrive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddToGoogleDrive) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddToGoogleDrive) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGoogleToken

`func (o *AddToGoogleDrive) GetGoogleToken() string`

GetGoogleToken returns the GoogleToken field if non-nil, zero value otherwise.

### GetGoogleTokenOk

`func (o *AddToGoogleDrive) GetGoogleTokenOk() (*string, bool)`

GetGoogleTokenOk returns a tuple with the GoogleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleToken

`func (o *AddToGoogleDrive) SetGoogleToken(v string)`

SetGoogleToken sets GoogleToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


