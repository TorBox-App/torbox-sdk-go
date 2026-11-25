# GetStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**GetStatsResponseData**](GetStatsResponseData.md) |  | [optional] 

## Methods

### NewGetStatsResponse

`func NewGetStatsResponse(success bool, detail string, ) *GetStatsResponse`

NewGetStatsResponse instantiates a new GetStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatsResponseWithDefaults

`func NewGetStatsResponseWithDefaults() *GetStatsResponse`

NewGetStatsResponseWithDefaults instantiates a new GetStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetStatsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetStatsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetStatsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetStatsResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetStatsResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetStatsResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetStatsResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetStatsResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetStatsResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetStatsResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetStatsResponse) GetData() GetStatsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStatsResponse) GetDataOk() (*GetStatsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStatsResponse) SetData(v GetStatsResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetStatsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


