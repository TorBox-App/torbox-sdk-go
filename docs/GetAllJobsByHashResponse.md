# GetAllJobsByHashResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]GetSpecificJobResponseData**](GetSpecificJobResponseData.md) |  | [optional] 

## Methods

### NewGetAllJobsByHashResponse

`func NewGetAllJobsByHashResponse(success bool, detail string, ) *GetAllJobsByHashResponse`

NewGetAllJobsByHashResponse instantiates a new GetAllJobsByHashResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllJobsByHashResponseWithDefaults

`func NewGetAllJobsByHashResponseWithDefaults() *GetAllJobsByHashResponse`

NewGetAllJobsByHashResponseWithDefaults instantiates a new GetAllJobsByHashResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetAllJobsByHashResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetAllJobsByHashResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetAllJobsByHashResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetAllJobsByHashResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAllJobsByHashResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAllJobsByHashResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetAllJobsByHashResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetAllJobsByHashResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetAllJobsByHashResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetAllJobsByHashResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetAllJobsByHashResponse) GetData() []GetSpecificJobResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllJobsByHashResponse) GetDataOk() (*[]GetSpecificJobResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllJobsByHashResponse) SetData(v []GetSpecificJobResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllJobsByHashResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


