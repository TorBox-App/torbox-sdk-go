# AddTo1Fichier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the torrent, usenet download or web download to upload. It has to be one of your own downloads. Required. | 
**FileId** | Pointer to **int64** | Id of the file to upload, from the download&#39;s files. Leave out when &#x60;zip&#x60; is true. | [optional] 
**Zip** | Pointer to **bool** | Upload the whole download as one zip instead of a single file. Takes precedence over &#x60;file_id&#x60;. Optional, default false. | [optional] [default to false]
**Type** | Pointer to **string** | Which kind of download &#x60;id&#x60; refers to. One of &#x60;torrent&#x60;, &#x60;usenet&#x60; or &#x60;webdownload&#x60;. Optional, default &#x60;torrent&#x60;. | [optional] [default to "torrent"]
**OnefichierToken** | Pointer to **string** | 1Fichier API token of the account to upload to. It is issued by 1Fichier, not by TorBox. Optional. | [optional] 

## Methods

### NewAddTo1Fichier

`func NewAddTo1Fichier(id int64, ) *AddTo1Fichier`

NewAddTo1Fichier instantiates a new AddTo1Fichier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTo1FichierWithDefaults

`func NewAddTo1FichierWithDefaults() *AddTo1Fichier`

NewAddTo1FichierWithDefaults instantiates a new AddTo1Fichier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTo1Fichier) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTo1Fichier) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTo1Fichier) SetId(v int64)`

SetId sets Id field to given value.


### GetFileId

`func (o *AddTo1Fichier) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AddTo1Fichier) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AddTo1Fichier) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AddTo1Fichier) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetZip

`func (o *AddTo1Fichier) GetZip() bool`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddTo1Fichier) GetZipOk() (*bool, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddTo1Fichier) SetZip(v bool)`

SetZip sets Zip field to given value.

### HasZip

`func (o *AddTo1Fichier) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetType

`func (o *AddTo1Fichier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddTo1Fichier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddTo1Fichier) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddTo1Fichier) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOnefichierToken

`func (o *AddTo1Fichier) GetOnefichierToken() string`

GetOnefichierToken returns the OnefichierToken field if non-nil, zero value otherwise.

### GetOnefichierTokenOk

`func (o *AddTo1Fichier) GetOnefichierTokenOk() (*string, bool)`

GetOnefichierTokenOk returns a tuple with the OnefichierToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnefichierToken

`func (o *AddTo1Fichier) SetOnefichierToken(v string)`

SetOnefichierToken sets OnefichierToken field to given value.

### HasOnefichierToken

`func (o *AddTo1Fichier) HasOnefichierToken() bool`

HasOnefichierToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


