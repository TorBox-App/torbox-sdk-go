# Get30DayStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]Get30DayStatsResponseDataInner**](Get30DayStatsResponseDataInner.md) |  | [optional] 

## Methods

### NewGet30DayStatsResponse

`func NewGet30DayStatsResponse(success bool, detail string, ) *Get30DayStatsResponse`

NewGet30DayStatsResponse instantiates a new Get30DayStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGet30DayStatsResponseWithDefaults

`func NewGet30DayStatsResponseWithDefaults() *Get30DayStatsResponse`

NewGet30DayStatsResponseWithDefaults instantiates a new Get30DayStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Get30DayStatsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Get30DayStatsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Get30DayStatsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *Get30DayStatsResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Get30DayStatsResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Get30DayStatsResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Get30DayStatsResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *Get30DayStatsResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Get30DayStatsResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Get30DayStatsResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *Get30DayStatsResponse) GetData() []Get30DayStatsResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Get30DayStatsResponse) GetDataOk() (*[]Get30DayStatsResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Get30DayStatsResponse) SetData(v []Get30DayStatsResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *Get30DayStatsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


