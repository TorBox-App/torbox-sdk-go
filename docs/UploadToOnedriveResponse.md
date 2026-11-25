# UploadToOnedriveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**UploadToGoogleDriveResponseData**](UploadToGoogleDriveResponseData.md) |  | [optional] 

## Methods

### NewUploadToOnedriveResponse

`func NewUploadToOnedriveResponse(success bool, detail string, ) *UploadToOnedriveResponse`

NewUploadToOnedriveResponse instantiates a new UploadToOnedriveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadToOnedriveResponseWithDefaults

`func NewUploadToOnedriveResponseWithDefaults() *UploadToOnedriveResponse`

NewUploadToOnedriveResponseWithDefaults instantiates a new UploadToOnedriveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UploadToOnedriveResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadToOnedriveResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadToOnedriveResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *UploadToOnedriveResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UploadToOnedriveResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UploadToOnedriveResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UploadToOnedriveResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *UploadToOnedriveResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UploadToOnedriveResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UploadToOnedriveResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *UploadToOnedriveResponse) GetData() UploadToGoogleDriveResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadToOnedriveResponse) GetDataOk() (*UploadToGoogleDriveResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadToOnedriveResponse) SetData(v UploadToGoogleDriveResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UploadToOnedriveResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


