# ControlRss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | The operation you want to perform on the RSS feed. Options are &#39;update&#39;, &#39;delete&#39;, &#39;resume&#39; and &#39;pause&#39;. | 
**RssFeedId** | **int64** | The ID of the RSS feed you want to perform the operation on. | 

## Methods

### NewControlRss

`func NewControlRss(operation string, rssFeedId int64, ) *ControlRss`

NewControlRss instantiates a new ControlRss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlRssWithDefaults

`func NewControlRssWithDefaults() *ControlRss`

NewControlRssWithDefaults instantiates a new ControlRss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ControlRss) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ControlRss) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ControlRss) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetRssFeedId

`func (o *ControlRss) GetRssFeedId() int64`

GetRssFeedId returns the RssFeedId field if non-nil, zero value otherwise.

### GetRssFeedIdOk

`func (o *ControlRss) GetRssFeedIdOk() (*int64, bool)`

GetRssFeedIdOk returns a tuple with the RssFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssFeedId

`func (o *ControlRss) SetRssFeedId(v int64)`

SetRssFeedId sets RssFeedId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


