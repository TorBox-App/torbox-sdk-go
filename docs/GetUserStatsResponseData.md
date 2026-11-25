# GetUserStatsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**GetUserStatsResponseDataGeneral**](GetUserStatsResponseDataGeneral.md) |  | [optional] 
**Bandwidth** | Pointer to [**[]GetUserStatsResponseDataBandwidthInner**](GetUserStatsResponseDataBandwidthInner.md) |  | [optional] 

## Methods

### NewGetUserStatsResponseData

`func NewGetUserStatsResponseData() *GetUserStatsResponseData`

NewGetUserStatsResponseData instantiates a new GetUserStatsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserStatsResponseDataWithDefaults

`func NewGetUserStatsResponseDataWithDefaults() *GetUserStatsResponseData`

NewGetUserStatsResponseDataWithDefaults instantiates a new GetUserStatsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *GetUserStatsResponseData) GetGeneral() GetUserStatsResponseDataGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *GetUserStatsResponseData) GetGeneralOk() (*GetUserStatsResponseDataGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *GetUserStatsResponseData) SetGeneral(v GetUserStatsResponseDataGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *GetUserStatsResponseData) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetBandwidth

`func (o *GetUserStatsResponseData) GetBandwidth() []GetUserStatsResponseDataBandwidthInner`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *GetUserStatsResponseData) GetBandwidthOk() (*[]GetUserStatsResponseDataBandwidthInner, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *GetUserStatsResponseData) SetBandwidth(v []GetUserStatsResponseDataBandwidthInner)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *GetUserStatsResponseData) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


