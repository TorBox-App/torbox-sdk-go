# ControlWebDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | the operation you want to perform on the web download. Can be \&quot;delete\&quot;. | 
**WebdlId** | Pointer to **int64** | the web download&#39;s id, optional if using the \&quot;all\&quot; parameter. | [optional] 
**All** | Pointer to **bool** | if you want to do this operation to all web downloads in your account, optional if using the \&quot;webdl_id\&quot; parameter | [optional] [default to false]

## Methods

### NewControlWebDownload

`func NewControlWebDownload(operation string, ) *ControlWebDownload`

NewControlWebDownload instantiates a new ControlWebDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlWebDownloadWithDefaults

`func NewControlWebDownloadWithDefaults() *ControlWebDownload`

NewControlWebDownloadWithDefaults instantiates a new ControlWebDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ControlWebDownload) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ControlWebDownload) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ControlWebDownload) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetWebdlId

`func (o *ControlWebDownload) GetWebdlId() int64`

GetWebdlId returns the WebdlId field if non-nil, zero value otherwise.

### GetWebdlIdOk

`func (o *ControlWebDownload) GetWebdlIdOk() (*int64, bool)`

GetWebdlIdOk returns a tuple with the WebdlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebdlId

`func (o *ControlWebDownload) SetWebdlId(v int64)`

SetWebdlId sets WebdlId field to given value.

### HasWebdlId

`func (o *ControlWebDownload) HasWebdlId() bool`

HasWebdlId returns a boolean if a field has been set.

### GetAll

`func (o *ControlWebDownload) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ControlWebDownload) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ControlWebDownload) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ControlWebDownload) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


