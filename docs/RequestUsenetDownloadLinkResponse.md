# RequestUsenetDownloadLinkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestUsenetDownloadLinkResponse

`func NewRequestUsenetDownloadLinkResponse(success bool, detail string, ) *RequestUsenetDownloadLinkResponse`

NewRequestUsenetDownloadLinkResponse instantiates a new RequestUsenetDownloadLinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestUsenetDownloadLinkResponseWithDefaults

`func NewRequestUsenetDownloadLinkResponseWithDefaults() *RequestUsenetDownloadLinkResponse`

NewRequestUsenetDownloadLinkResponseWithDefaults instantiates a new RequestUsenetDownloadLinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RequestUsenetDownloadLinkResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RequestUsenetDownloadLinkResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RequestUsenetDownloadLinkResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *RequestUsenetDownloadLinkResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RequestUsenetDownloadLinkResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RequestUsenetDownloadLinkResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RequestUsenetDownloadLinkResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *RequestUsenetDownloadLinkResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RequestUsenetDownloadLinkResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RequestUsenetDownloadLinkResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *RequestUsenetDownloadLinkResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestUsenetDownloadLinkResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestUsenetDownloadLinkResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *RequestUsenetDownloadLinkResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


