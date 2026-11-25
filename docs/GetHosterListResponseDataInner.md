# GetHosterListResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **interface{}** |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 
**DailyLinkLimit** | Pointer to **int64** |  | [optional] 
**DailyLinkUsed** | Pointer to **float32** |  | [optional] 
**DailyBandwidthLimit** | Pointer to **int64** |  | [optional] 
**DailyBandwidthUsed** | Pointer to **int64** |  | [optional] 
**PerLinkSizeLimit** | Pointer to **int64** |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] 

## Methods

### NewGetHosterListResponseDataInner

`func NewGetHosterListResponseDataInner() *GetHosterListResponseDataInner`

NewGetHosterListResponseDataInner instantiates a new GetHosterListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHosterListResponseDataInnerWithDefaults

`func NewGetHosterListResponseDataInnerWithDefaults() *GetHosterListResponseDataInner`

NewGetHosterListResponseDataInnerWithDefaults instantiates a new GetHosterListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHosterListResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHosterListResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHosterListResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetHosterListResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetHosterListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHosterListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHosterListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHosterListResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomains

`func (o *GetHosterListResponseDataInner) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *GetHosterListResponseDataInner) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *GetHosterListResponseDataInner) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *GetHosterListResponseDataInner) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetUrl

`func (o *GetHosterListResponseDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetHosterListResponseDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetHosterListResponseDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetHosterListResponseDataInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIcon

`func (o *GetHosterListResponseDataInner) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GetHosterListResponseDataInner) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GetHosterListResponseDataInner) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GetHosterListResponseDataInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetStatus

`func (o *GetHosterListResponseDataInner) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHosterListResponseDataInner) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHosterListResponseDataInner) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetHosterListResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetHosterListResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetHosterListResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetHosterListResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetHosterListResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNote

`func (o *GetHosterListResponseDataInner) GetNote() interface{}`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetHosterListResponseDataInner) GetNoteOk() (*interface{}, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetHosterListResponseDataInner) SetNote(v interface{})`

SetNote sets Note field to given value.

### HasNote

`func (o *GetHosterListResponseDataInner) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *GetHosterListResponseDataInner) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *GetHosterListResponseDataInner) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetNsfw

`func (o *GetHosterListResponseDataInner) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *GetHosterListResponseDataInner) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *GetHosterListResponseDataInner) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *GetHosterListResponseDataInner) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetDailyLinkLimit

`func (o *GetHosterListResponseDataInner) GetDailyLinkLimit() int64`

GetDailyLinkLimit returns the DailyLinkLimit field if non-nil, zero value otherwise.

### GetDailyLinkLimitOk

`func (o *GetHosterListResponseDataInner) GetDailyLinkLimitOk() (*int64, bool)`

GetDailyLinkLimitOk returns a tuple with the DailyLinkLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLinkLimit

`func (o *GetHosterListResponseDataInner) SetDailyLinkLimit(v int64)`

SetDailyLinkLimit sets DailyLinkLimit field to given value.

### HasDailyLinkLimit

`func (o *GetHosterListResponseDataInner) HasDailyLinkLimit() bool`

HasDailyLinkLimit returns a boolean if a field has been set.

### GetDailyLinkUsed

`func (o *GetHosterListResponseDataInner) GetDailyLinkUsed() float32`

GetDailyLinkUsed returns the DailyLinkUsed field if non-nil, zero value otherwise.

### GetDailyLinkUsedOk

`func (o *GetHosterListResponseDataInner) GetDailyLinkUsedOk() (*float32, bool)`

GetDailyLinkUsedOk returns a tuple with the DailyLinkUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLinkUsed

`func (o *GetHosterListResponseDataInner) SetDailyLinkUsed(v float32)`

SetDailyLinkUsed sets DailyLinkUsed field to given value.

### HasDailyLinkUsed

`func (o *GetHosterListResponseDataInner) HasDailyLinkUsed() bool`

HasDailyLinkUsed returns a boolean if a field has been set.

### GetDailyBandwidthLimit

`func (o *GetHosterListResponseDataInner) GetDailyBandwidthLimit() int64`

GetDailyBandwidthLimit returns the DailyBandwidthLimit field if non-nil, zero value otherwise.

### GetDailyBandwidthLimitOk

`func (o *GetHosterListResponseDataInner) GetDailyBandwidthLimitOk() (*int64, bool)`

GetDailyBandwidthLimitOk returns a tuple with the DailyBandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBandwidthLimit

`func (o *GetHosterListResponseDataInner) SetDailyBandwidthLimit(v int64)`

SetDailyBandwidthLimit sets DailyBandwidthLimit field to given value.

### HasDailyBandwidthLimit

`func (o *GetHosterListResponseDataInner) HasDailyBandwidthLimit() bool`

HasDailyBandwidthLimit returns a boolean if a field has been set.

### GetDailyBandwidthUsed

`func (o *GetHosterListResponseDataInner) GetDailyBandwidthUsed() int64`

GetDailyBandwidthUsed returns the DailyBandwidthUsed field if non-nil, zero value otherwise.

### GetDailyBandwidthUsedOk

`func (o *GetHosterListResponseDataInner) GetDailyBandwidthUsedOk() (*int64, bool)`

GetDailyBandwidthUsedOk returns a tuple with the DailyBandwidthUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBandwidthUsed

`func (o *GetHosterListResponseDataInner) SetDailyBandwidthUsed(v int64)`

SetDailyBandwidthUsed sets DailyBandwidthUsed field to given value.

### HasDailyBandwidthUsed

`func (o *GetHosterListResponseDataInner) HasDailyBandwidthUsed() bool`

HasDailyBandwidthUsed returns a boolean if a field has been set.

### GetPerLinkSizeLimit

`func (o *GetHosterListResponseDataInner) GetPerLinkSizeLimit() int64`

GetPerLinkSizeLimit returns the PerLinkSizeLimit field if non-nil, zero value otherwise.

### GetPerLinkSizeLimitOk

`func (o *GetHosterListResponseDataInner) GetPerLinkSizeLimitOk() (*int64, bool)`

GetPerLinkSizeLimitOk returns a tuple with the PerLinkSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerLinkSizeLimit

`func (o *GetHosterListResponseDataInner) SetPerLinkSizeLimit(v int64)`

SetPerLinkSizeLimit sets PerLinkSizeLimit field to given value.

### HasPerLinkSizeLimit

`func (o *GetHosterListResponseDataInner) HasPerLinkSizeLimit() bool`

HasPerLinkSizeLimit returns a boolean if a field has been set.

### GetRegex

`func (o *GetHosterListResponseDataInner) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *GetHosterListResponseDataInner) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *GetHosterListResponseDataInner) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *GetHosterListResponseDataInner) HasRegex() bool`

HasRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


