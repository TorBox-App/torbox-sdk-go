# GetReferralDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**GetReferralDataResponseData**](GetReferralDataResponseData.md) |  | [optional] 

## Methods

### NewGetReferralDataResponse

`func NewGetReferralDataResponse(success bool, detail string, ) *GetReferralDataResponse`

NewGetReferralDataResponse instantiates a new GetReferralDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReferralDataResponseWithDefaults

`func NewGetReferralDataResponseWithDefaults() *GetReferralDataResponse`

NewGetReferralDataResponseWithDefaults instantiates a new GetReferralDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetReferralDataResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetReferralDataResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetReferralDataResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetReferralDataResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetReferralDataResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetReferralDataResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetReferralDataResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetReferralDataResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetReferralDataResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetReferralDataResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetReferralDataResponse) GetData() GetReferralDataResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetReferralDataResponse) GetDataOk() (*GetReferralDataResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetReferralDataResponse) SetData(v GetReferralDataResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetReferralDataResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


