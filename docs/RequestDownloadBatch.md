# RequestDownloadBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileIds** | Pointer to **[]int64** | The files&#39;s ID that you want to download | [optional] [default to {}]

## Methods

### NewRequestDownloadBatch

`func NewRequestDownloadBatch() *RequestDownloadBatch`

NewRequestDownloadBatch instantiates a new RequestDownloadBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestDownloadBatchWithDefaults

`func NewRequestDownloadBatchWithDefaults() *RequestDownloadBatch`

NewRequestDownloadBatchWithDefaults instantiates a new RequestDownloadBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileIds

`func (o *RequestDownloadBatch) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *RequestDownloadBatch) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *RequestDownloadBatch) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *RequestDownloadBatch) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


