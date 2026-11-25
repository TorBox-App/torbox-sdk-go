# ControlUsenetDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | the operation you want to perform on the usenet download. Can be \&quot;delete\&quot;, \&quot;pause\&quot; or \&quot;resume\&quot;. | 
**UsenetId** | Pointer to **int64** | the usenet download&#39;s id, optional if using the \&quot;all\&quot; parameter. | [optional] 
**All** | Pointer to **bool** | if you want to do this operation to all usenet downloads in your account, optional if using the \&quot;usenet_id\&quot; parameter | [optional] [default to false]

## Methods

### NewControlUsenetDownload

`func NewControlUsenetDownload(operation string, ) *ControlUsenetDownload`

NewControlUsenetDownload instantiates a new ControlUsenetDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlUsenetDownloadWithDefaults

`func NewControlUsenetDownloadWithDefaults() *ControlUsenetDownload`

NewControlUsenetDownloadWithDefaults instantiates a new ControlUsenetDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ControlUsenetDownload) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ControlUsenetDownload) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ControlUsenetDownload) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetUsenetId

`func (o *ControlUsenetDownload) GetUsenetId() int64`

GetUsenetId returns the UsenetId field if non-nil, zero value otherwise.

### GetUsenetIdOk

`func (o *ControlUsenetDownload) GetUsenetIdOk() (*int64, bool)`

GetUsenetIdOk returns a tuple with the UsenetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsenetId

`func (o *ControlUsenetDownload) SetUsenetId(v int64)`

SetUsenetId sets UsenetId field to given value.

### HasUsenetId

`func (o *ControlUsenetDownload) HasUsenetId() bool`

HasUsenetId returns a boolean if a field has been set.

### GetAll

`func (o *ControlUsenetDownload) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ControlUsenetDownload) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ControlUsenetDownload) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ControlUsenetDownload) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


