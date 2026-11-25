# RequestWebDownloadLinksGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]RequestTorrentDownloadLinksGetResponseDataInner**](RequestTorrentDownloadLinksGetResponseDataInner.md) |  | [optional] 

## Methods

### NewRequestWebDownloadLinksGetResponse

`func NewRequestWebDownloadLinksGetResponse(success bool, detail string, ) *RequestWebDownloadLinksGetResponse`

NewRequestWebDownloadLinksGetResponse instantiates a new RequestWebDownloadLinksGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWebDownloadLinksGetResponseWithDefaults

`func NewRequestWebDownloadLinksGetResponseWithDefaults() *RequestWebDownloadLinksGetResponse`

NewRequestWebDownloadLinksGetResponseWithDefaults instantiates a new RequestWebDownloadLinksGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RequestWebDownloadLinksGetResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RequestWebDownloadLinksGetResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RequestWebDownloadLinksGetResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *RequestWebDownloadLinksGetResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RequestWebDownloadLinksGetResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RequestWebDownloadLinksGetResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RequestWebDownloadLinksGetResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *RequestWebDownloadLinksGetResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RequestWebDownloadLinksGetResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RequestWebDownloadLinksGetResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *RequestWebDownloadLinksGetResponse) GetData() []RequestTorrentDownloadLinksGetResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestWebDownloadLinksGetResponse) GetDataOk() (*[]RequestTorrentDownloadLinksGetResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestWebDownloadLinksGetResponse) SetData(v []RequestTorrentDownloadLinksGetResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *RequestWebDownloadLinksGetResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


