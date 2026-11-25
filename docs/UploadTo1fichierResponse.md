# UploadTo1fichierResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**UploadToGoogleDriveResponseData**](UploadToGoogleDriveResponseData.md) |  | [optional] 

## Methods

### NewUploadTo1fichierResponse

`func NewUploadTo1fichierResponse(success bool, detail string, ) *UploadTo1fichierResponse`

NewUploadTo1fichierResponse instantiates a new UploadTo1fichierResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadTo1fichierResponseWithDefaults

`func NewUploadTo1fichierResponseWithDefaults() *UploadTo1fichierResponse`

NewUploadTo1fichierResponseWithDefaults instantiates a new UploadTo1fichierResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UploadTo1fichierResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadTo1fichierResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadTo1fichierResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *UploadTo1fichierResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UploadTo1fichierResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UploadTo1fichierResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UploadTo1fichierResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *UploadTo1fichierResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UploadTo1fichierResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UploadTo1fichierResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *UploadTo1fichierResponse) GetData() UploadToGoogleDriveResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadTo1fichierResponse) GetDataOk() (*UploadToGoogleDriveResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadTo1fichierResponse) SetData(v UploadToGoogleDriveResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UploadTo1fichierResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


