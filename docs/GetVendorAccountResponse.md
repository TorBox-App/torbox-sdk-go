# GetVendorAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**GetVendorAccountResponseData**](GetVendorAccountResponseData.md) |  | [optional] 

## Methods

### NewGetVendorAccountResponse

`func NewGetVendorAccountResponse(success bool, detail string, ) *GetVendorAccountResponse`

NewGetVendorAccountResponse instantiates a new GetVendorAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVendorAccountResponseWithDefaults

`func NewGetVendorAccountResponseWithDefaults() *GetVendorAccountResponse`

NewGetVendorAccountResponseWithDefaults instantiates a new GetVendorAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetVendorAccountResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetVendorAccountResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetVendorAccountResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetVendorAccountResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetVendorAccountResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetVendorAccountResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetVendorAccountResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetVendorAccountResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetVendorAccountResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetVendorAccountResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetVendorAccountResponse) GetData() GetVendorAccountResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVendorAccountResponse) GetDataOk() (*GetVendorAccountResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVendorAccountResponse) SetData(v GetVendorAccountResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetVendorAccountResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


