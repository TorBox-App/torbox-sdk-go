# UploadToGofileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**UploadToGoogleDriveResponseData**](UploadToGoogleDriveResponseData.md) |  | [optional] 

## Methods

### NewUploadToGofileResponse

`func NewUploadToGofileResponse(success bool, detail string, ) *UploadToGofileResponse`

NewUploadToGofileResponse instantiates a new UploadToGofileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadToGofileResponseWithDefaults

`func NewUploadToGofileResponseWithDefaults() *UploadToGofileResponse`

NewUploadToGofileResponseWithDefaults instantiates a new UploadToGofileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UploadToGofileResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadToGofileResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadToGofileResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *UploadToGofileResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UploadToGofileResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UploadToGofileResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UploadToGofileResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *UploadToGofileResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UploadToGofileResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UploadToGofileResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *UploadToGofileResponse) GetData() UploadToGoogleDriveResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadToGofileResponse) GetDataOk() (*UploadToGoogleDriveResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadToGofileResponse) SetData(v UploadToGoogleDriveResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UploadToGofileResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


