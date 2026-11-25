# UploadToGoogleDriveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**UploadToGoogleDriveResponseData**](UploadToGoogleDriveResponseData.md) |  | [optional] 

## Methods

### NewUploadToGoogleDriveResponse

`func NewUploadToGoogleDriveResponse(success bool, detail string, ) *UploadToGoogleDriveResponse`

NewUploadToGoogleDriveResponse instantiates a new UploadToGoogleDriveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadToGoogleDriveResponseWithDefaults

`func NewUploadToGoogleDriveResponseWithDefaults() *UploadToGoogleDriveResponse`

NewUploadToGoogleDriveResponseWithDefaults instantiates a new UploadToGoogleDriveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UploadToGoogleDriveResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadToGoogleDriveResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadToGoogleDriveResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *UploadToGoogleDriveResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UploadToGoogleDriveResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UploadToGoogleDriveResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UploadToGoogleDriveResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *UploadToGoogleDriveResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UploadToGoogleDriveResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UploadToGoogleDriveResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *UploadToGoogleDriveResponse) GetData() UploadToGoogleDriveResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadToGoogleDriveResponse) GetDataOk() (*UploadToGoogleDriveResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadToGoogleDriveResponse) SetData(v UploadToGoogleDriveResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UploadToGoogleDriveResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


