# ClearSingleNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewClearSingleNotificationResponse

`func NewClearSingleNotificationResponse(success bool, detail string, ) *ClearSingleNotificationResponse`

NewClearSingleNotificationResponse instantiates a new ClearSingleNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearSingleNotificationResponseWithDefaults

`func NewClearSingleNotificationResponseWithDefaults() *ClearSingleNotificationResponse`

NewClearSingleNotificationResponseWithDefaults instantiates a new ClearSingleNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ClearSingleNotificationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ClearSingleNotificationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ClearSingleNotificationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *ClearSingleNotificationResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClearSingleNotificationResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClearSingleNotificationResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ClearSingleNotificationResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *ClearSingleNotificationResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ClearSingleNotificationResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ClearSingleNotificationResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *ClearSingleNotificationResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ClearSingleNotificationResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ClearSingleNotificationResponse) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ClearSingleNotificationResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ClearSingleNotificationResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ClearSingleNotificationResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


