# ControlQueuedDownloadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewControlQueuedDownloadsResponse

`func NewControlQueuedDownloadsResponse(success bool, detail string, ) *ControlQueuedDownloadsResponse`

NewControlQueuedDownloadsResponse instantiates a new ControlQueuedDownloadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlQueuedDownloadsResponseWithDefaults

`func NewControlQueuedDownloadsResponseWithDefaults() *ControlQueuedDownloadsResponse`

NewControlQueuedDownloadsResponseWithDefaults instantiates a new ControlQueuedDownloadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ControlQueuedDownloadsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ControlQueuedDownloadsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ControlQueuedDownloadsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *ControlQueuedDownloadsResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ControlQueuedDownloadsResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ControlQueuedDownloadsResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ControlQueuedDownloadsResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *ControlQueuedDownloadsResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ControlQueuedDownloadsResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ControlQueuedDownloadsResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *ControlQueuedDownloadsResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ControlQueuedDownloadsResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ControlQueuedDownloadsResponse) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ControlQueuedDownloadsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ControlQueuedDownloadsResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ControlQueuedDownloadsResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


