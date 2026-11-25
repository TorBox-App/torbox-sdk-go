# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | False when the request failed. A list with nothing to return answers &#x60;ITEM_NOT_FOUND&#x60; with &#x60;success&#x60; true: the request worked, there is no data. | 
**Error** | **string** | The error code, such as &#x60;BAD_TOKEN&#x60;. Listed in the introduction. | 
**Detail** | **string** | What went wrong, in words. | 
**Data** | Pointer to **interface{}** | Extra data for some errors, usually null. | [optional] 

## Methods

### NewApiError

`func NewApiError(success bool, error_ string, detail string, ) *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ApiError) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiError) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiError) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *ApiError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiError) SetError(v string)`

SetError sets Error field to given value.


### GetDetail

`func (o *ApiError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ApiError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ApiError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *ApiError) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiError) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiError) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ApiError) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ApiError) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ApiError) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


