# GetChangelogsJsonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Detail** | **string** |  | 
**Data** | Pointer to [**[]GetChangelogsJsonResponseDataInner**](GetChangelogsJsonResponseDataInner.md) |  | [optional] 

## Methods

### NewGetChangelogsJsonResponse

`func NewGetChangelogsJsonResponse(success bool, detail string, ) *GetChangelogsJsonResponse`

NewGetChangelogsJsonResponse instantiates a new GetChangelogsJsonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChangelogsJsonResponseWithDefaults

`func NewGetChangelogsJsonResponseWithDefaults() *GetChangelogsJsonResponse`

NewGetChangelogsJsonResponseWithDefaults instantiates a new GetChangelogsJsonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetChangelogsJsonResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetChangelogsJsonResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetChangelogsJsonResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *GetChangelogsJsonResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetChangelogsJsonResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetChangelogsJsonResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetChangelogsJsonResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetail

`func (o *GetChangelogsJsonResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetChangelogsJsonResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetChangelogsJsonResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetData

`func (o *GetChangelogsJsonResponse) GetData() []GetChangelogsJsonResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChangelogsJsonResponse) GetDataOk() (*[]GetChangelogsJsonResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChangelogsJsonResponse) SetData(v []GetChangelogsJsonResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetChangelogsJsonResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


