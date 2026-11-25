# GetVendorUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]GetVendorUsersResponseDataInner**](GetVendorUsersResponseDataInner.md) |  | [optional] 

## Methods

### NewGetVendorUsersResponse

`func NewGetVendorUsersResponse(success bool, detail string, ) *GetVendorUsersResponse`

NewGetVendorUsersResponse instantiates a new GetVendorUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVendorUsersResponseWithDefaults

`func NewGetVendorUsersResponseWithDefaults() *GetVendorUsersResponse`

NewGetVendorUsersResponseWithDefaults instantiates a new GetVendorUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetVendorUsersResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetVendorUsersResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetVendorUsersResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetVendorUsersResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetVendorUsersResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetVendorUsersResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetVendorUsersResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetVendorUsersResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetVendorUsersResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetVendorUsersResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetVendorUsersResponse) GetData() []GetVendorUsersResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVendorUsersResponse) GetDataOk() (*[]GetVendorUsersResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVendorUsersResponse) SetData(v []GetVendorUsersResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetVendorUsersResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


