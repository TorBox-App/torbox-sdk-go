# RequestTorrentDownloadLinksPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]RequestTorrentDownloadLinksGetResponseDataInner**](RequestTorrentDownloadLinksGetResponseDataInner.md) |  | [optional] 

## Methods

### NewRequestTorrentDownloadLinksPostResponse

`func NewRequestTorrentDownloadLinksPostResponse(success bool, detail string, ) *RequestTorrentDownloadLinksPostResponse`

NewRequestTorrentDownloadLinksPostResponse instantiates a new RequestTorrentDownloadLinksPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTorrentDownloadLinksPostResponseWithDefaults

`func NewRequestTorrentDownloadLinksPostResponseWithDefaults() *RequestTorrentDownloadLinksPostResponse`

NewRequestTorrentDownloadLinksPostResponseWithDefaults instantiates a new RequestTorrentDownloadLinksPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RequestTorrentDownloadLinksPostResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RequestTorrentDownloadLinksPostResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RequestTorrentDownloadLinksPostResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *RequestTorrentDownloadLinksPostResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RequestTorrentDownloadLinksPostResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RequestTorrentDownloadLinksPostResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RequestTorrentDownloadLinksPostResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *RequestTorrentDownloadLinksPostResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RequestTorrentDownloadLinksPostResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RequestTorrentDownloadLinksPostResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *RequestTorrentDownloadLinksPostResponse) GetData() []RequestTorrentDownloadLinksGetResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestTorrentDownloadLinksPostResponse) GetDataOk() (*[]RequestTorrentDownloadLinksGetResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestTorrentDownloadLinksPostResponse) SetData(v []RequestTorrentDownloadLinksGetResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *RequestTorrentDownloadLinksPostResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


