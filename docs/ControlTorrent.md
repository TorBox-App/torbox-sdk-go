# ControlTorrent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | the operation you want to perform on the torrent | 
**TorrentId** | Pointer to **int64** | the torrent&#39;s id, optional if using \&quot;all\&quot; parameter | [optional] 
**All** | Pointer to **bool** | if you want to do this operation to all torrents in your account, optional if using the \&quot;torrent_id\&quot; parameter | [optional] [default to false]

## Methods

### NewControlTorrent

`func NewControlTorrent(operation string, ) *ControlTorrent`

NewControlTorrent instantiates a new ControlTorrent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlTorrentWithDefaults

`func NewControlTorrentWithDefaults() *ControlTorrent`

NewControlTorrentWithDefaults instantiates a new ControlTorrent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ControlTorrent) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ControlTorrent) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ControlTorrent) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetTorrentId

`func (o *ControlTorrent) GetTorrentId() int64`

GetTorrentId returns the TorrentId field if non-nil, zero value otherwise.

### GetTorrentIdOk

`func (o *ControlTorrent) GetTorrentIdOk() (*int64, bool)`

GetTorrentIdOk returns a tuple with the TorrentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentId

`func (o *ControlTorrent) SetTorrentId(v int64)`

SetTorrentId sets TorrentId field to given value.

### HasTorrentId

`func (o *ControlTorrent) HasTorrentId() bool`

HasTorrentId returns a boolean if a field has been set.

### GetAll

`func (o *ControlTorrent) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ControlTorrent) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ControlTorrent) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ControlTorrent) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


