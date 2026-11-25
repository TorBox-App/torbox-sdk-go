# GetUsenetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]GetUsenetListResponseDataInner**](GetUsenetListResponseDataInner.md) |  | [optional] 

## Methods

### NewGetUsenetListResponse

`func NewGetUsenetListResponse(success bool, detail string, ) *GetUsenetListResponse`

NewGetUsenetListResponse instantiates a new GetUsenetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsenetListResponseWithDefaults

`func NewGetUsenetListResponseWithDefaults() *GetUsenetListResponse`

NewGetUsenetListResponseWithDefaults instantiates a new GetUsenetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetUsenetListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetUsenetListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetUsenetListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetUsenetListResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetUsenetListResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetUsenetListResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetUsenetListResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetUsenetListResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetUsenetListResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetUsenetListResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetUsenetListResponse) GetData() []GetUsenetListResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUsenetListResponse) GetDataOk() (*[]GetUsenetListResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUsenetListResponse) SetData(v []GetUsenetListResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetUsenetListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


