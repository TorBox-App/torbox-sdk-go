# GetConfirmationCodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGetConfirmationCodeResponse

`func NewGetConfirmationCodeResponse(success bool, detail string, ) *GetConfirmationCodeResponse`

NewGetConfirmationCodeResponse instantiates a new GetConfirmationCodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfirmationCodeResponseWithDefaults

`func NewGetConfirmationCodeResponseWithDefaults() *GetConfirmationCodeResponse`

NewGetConfirmationCodeResponseWithDefaults instantiates a new GetConfirmationCodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetConfirmationCodeResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetConfirmationCodeResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetConfirmationCodeResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetConfirmationCodeResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetConfirmationCodeResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetConfirmationCodeResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetConfirmationCodeResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetConfirmationCodeResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetConfirmationCodeResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetConfirmationCodeResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetConfirmationCodeResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetConfirmationCodeResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetConfirmationCodeResponse) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GetConfirmationCodeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetConfirmationCodeResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetConfirmationCodeResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


