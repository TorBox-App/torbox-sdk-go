# CreateTorrentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**CreateTorrentResponseData**](CreateTorrentResponseData.md) |  | [optional] 

## Methods

### NewCreateTorrentResponse

`func NewCreateTorrentResponse(success bool, detail string, ) *CreateTorrentResponse`

NewCreateTorrentResponse instantiates a new CreateTorrentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTorrentResponseWithDefaults

`func NewCreateTorrentResponseWithDefaults() *CreateTorrentResponse`

NewCreateTorrentResponseWithDefaults instantiates a new CreateTorrentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateTorrentResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateTorrentResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateTorrentResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *CreateTorrentResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateTorrentResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateTorrentResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreateTorrentResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *CreateTorrentResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CreateTorrentResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CreateTorrentResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *CreateTorrentResponse) GetData() CreateTorrentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTorrentResponse) GetDataOk() (*CreateTorrentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTorrentResponse) SetData(v CreateTorrentResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateTorrentResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


