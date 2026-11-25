# ControlQueuedDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | the operation you want to perform on the torrent | 
**QueuedId** | Pointer to **int64** | the queued torrent&#39;s id, optional if using the \&quot;all\&quot; parameter | [optional] 
**All** | Pointer to **bool** | if you want to do this operation to all queued torrents in your account, optional if using the \&quot;queued_id\&quot; parameter | [optional] [default to false]

## Methods

### NewControlQueuedDownload

`func NewControlQueuedDownload(operation string, ) *ControlQueuedDownload`

NewControlQueuedDownload instantiates a new ControlQueuedDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlQueuedDownloadWithDefaults

`func NewControlQueuedDownloadWithDefaults() *ControlQueuedDownload`

NewControlQueuedDownloadWithDefaults instantiates a new ControlQueuedDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ControlQueuedDownload) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ControlQueuedDownload) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ControlQueuedDownload) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetQueuedId

`func (o *ControlQueuedDownload) GetQueuedId() int64`

GetQueuedId returns the QueuedId field if non-nil, zero value otherwise.

### GetQueuedIdOk

`func (o *ControlQueuedDownload) GetQueuedIdOk() (*int64, bool)`

GetQueuedIdOk returns a tuple with the QueuedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedId

`func (o *ControlQueuedDownload) SetQueuedId(v int64)`

SetQueuedId sets QueuedId field to given value.

### HasQueuedId

`func (o *ControlQueuedDownload) HasQueuedId() bool`

HasQueuedId returns a boolean if a field has been set.

### GetAll

`func (o *ControlQueuedDownload) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ControlQueuedDownload) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ControlQueuedDownload) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ControlQueuedDownload) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


