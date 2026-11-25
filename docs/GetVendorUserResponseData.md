# GetVendorUserResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UserReferral** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **int64** |  | [optional] 
**TotalDownloaded** | Pointer to **int64** |  | [optional] 
**Customer** | Pointer to **string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**PremiumExpiresAt** | Pointer to **string** |  | [optional] 
**CooldownUntil** | Pointer to **string** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 
**BaseEmail** | Pointer to **string** |  | [optional] 
**TotalBytesDownloaded** | Pointer to **int64** |  | [optional] 
**TotalBytesUploaded** | Pointer to **int64** |  | [optional] 
**TorrentsDownloaded** | Pointer to **float32** |  | [optional] 
**WebDownloadsDownloaded** | Pointer to **float32** |  | [optional] 
**UsenetDownloadsDownloaded** | Pointer to **float32** |  | [optional] 
**AdditionalConcurrentSlots** | Pointer to **int64** |  | [optional] 
**LongTermSeeding** | Pointer to **bool** |  | [optional] 
**VendorId** | Pointer to **int64** |  | [optional] 
**IsVendor** | Pointer to **bool** |  | [optional] 
**PurchasesReferred** | Pointer to **float32** |  | [optional] 
**LongTermStorage** | Pointer to **bool** |  | [optional] 
**DailyBandwidthUsage** | Pointer to **int64** |  | [optional] 
**ApiToken** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVendorUserResponseData

`func NewGetVendorUserResponseData() *GetVendorUserResponseData`

NewGetVendorUserResponseData instantiates a new GetVendorUserResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVendorUserResponseDataWithDefaults

`func NewGetVendorUserResponseDataWithDefaults() *GetVendorUserResponseData`

NewGetVendorUserResponseDataWithDefaults instantiates a new GetVendorUserResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVendorUserResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVendorUserResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVendorUserResponseData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVendorUserResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserReferral

`func (o *GetVendorUserResponseData) GetUserReferral() string`

GetUserReferral returns the UserReferral field if non-nil, zero value otherwise.

### GetUserReferralOk

`func (o *GetVendorUserResponseData) GetUserReferralOk() (*string, bool)`

GetUserReferralOk returns a tuple with the UserReferral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReferral

`func (o *GetVendorUserResponseData) SetUserReferral(v string)`

SetUserReferral sets UserReferral field to given value.

### HasUserReferral

`func (o *GetVendorUserResponseData) HasUserReferral() bool`

HasUserReferral returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetVendorUserResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetVendorUserResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetVendorUserResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetVendorUserResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetVendorUserResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetVendorUserResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetVendorUserResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetVendorUserResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *GetVendorUserResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetVendorUserResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetVendorUserResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetVendorUserResponseData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPlan

`func (o *GetVendorUserResponseData) GetPlan() int64`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetVendorUserResponseData) GetPlanOk() (*int64, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetVendorUserResponseData) SetPlan(v int64)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetVendorUserResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetTotalDownloaded

`func (o *GetVendorUserResponseData) GetTotalDownloaded() int64`

GetTotalDownloaded returns the TotalDownloaded field if non-nil, zero value otherwise.

### GetTotalDownloadedOk

`func (o *GetVendorUserResponseData) GetTotalDownloadedOk() (*int64, bool)`

GetTotalDownloadedOk returns a tuple with the TotalDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownloaded

`func (o *GetVendorUserResponseData) SetTotalDownloaded(v int64)`

SetTotalDownloaded sets TotalDownloaded field to given value.

### HasTotalDownloaded

`func (o *GetVendorUserResponseData) HasTotalDownloaded() bool`

HasTotalDownloaded returns a boolean if a field has been set.

### GetCustomer

`func (o *GetVendorUserResponseData) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GetVendorUserResponseData) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GetVendorUserResponseData) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GetVendorUserResponseData) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *GetVendorUserResponseData) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *GetVendorUserResponseData) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *GetVendorUserResponseData) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *GetVendorUserResponseData) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetPremiumExpiresAt

`func (o *GetVendorUserResponseData) GetPremiumExpiresAt() string`

GetPremiumExpiresAt returns the PremiumExpiresAt field if non-nil, zero value otherwise.

### GetPremiumExpiresAtOk

`func (o *GetVendorUserResponseData) GetPremiumExpiresAtOk() (*string, bool)`

GetPremiumExpiresAtOk returns a tuple with the PremiumExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumExpiresAt

`func (o *GetVendorUserResponseData) SetPremiumExpiresAt(v string)`

SetPremiumExpiresAt sets PremiumExpiresAt field to given value.

### HasPremiumExpiresAt

`func (o *GetVendorUserResponseData) HasPremiumExpiresAt() bool`

HasPremiumExpiresAt returns a boolean if a field has been set.

### GetCooldownUntil

`func (o *GetVendorUserResponseData) GetCooldownUntil() string`

GetCooldownUntil returns the CooldownUntil field if non-nil, zero value otherwise.

### GetCooldownUntilOk

`func (o *GetVendorUserResponseData) GetCooldownUntilOk() (*string, bool)`

GetCooldownUntilOk returns a tuple with the CooldownUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownUntil

`func (o *GetVendorUserResponseData) SetCooldownUntil(v string)`

SetCooldownUntil sets CooldownUntil field to given value.

### HasCooldownUntil

`func (o *GetVendorUserResponseData) HasCooldownUntil() bool`

HasCooldownUntil returns a boolean if a field has been set.

### GetAuthId

`func (o *GetVendorUserResponseData) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *GetVendorUserResponseData) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *GetVendorUserResponseData) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *GetVendorUserResponseData) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetBaseEmail

`func (o *GetVendorUserResponseData) GetBaseEmail() string`

GetBaseEmail returns the BaseEmail field if non-nil, zero value otherwise.

### GetBaseEmailOk

`func (o *GetVendorUserResponseData) GetBaseEmailOk() (*string, bool)`

GetBaseEmailOk returns a tuple with the BaseEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEmail

`func (o *GetVendorUserResponseData) SetBaseEmail(v string)`

SetBaseEmail sets BaseEmail field to given value.

### HasBaseEmail

`func (o *GetVendorUserResponseData) HasBaseEmail() bool`

HasBaseEmail returns a boolean if a field has been set.

### GetTotalBytesDownloaded

`func (o *GetVendorUserResponseData) GetTotalBytesDownloaded() int64`

GetTotalBytesDownloaded returns the TotalBytesDownloaded field if non-nil, zero value otherwise.

### GetTotalBytesDownloadedOk

`func (o *GetVendorUserResponseData) GetTotalBytesDownloadedOk() (*int64, bool)`

GetTotalBytesDownloadedOk returns a tuple with the TotalBytesDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesDownloaded

`func (o *GetVendorUserResponseData) SetTotalBytesDownloaded(v int64)`

SetTotalBytesDownloaded sets TotalBytesDownloaded field to given value.

### HasTotalBytesDownloaded

`func (o *GetVendorUserResponseData) HasTotalBytesDownloaded() bool`

HasTotalBytesDownloaded returns a boolean if a field has been set.

### GetTotalBytesUploaded

`func (o *GetVendorUserResponseData) GetTotalBytesUploaded() int64`

GetTotalBytesUploaded returns the TotalBytesUploaded field if non-nil, zero value otherwise.

### GetTotalBytesUploadedOk

`func (o *GetVendorUserResponseData) GetTotalBytesUploadedOk() (*int64, bool)`

GetTotalBytesUploadedOk returns a tuple with the TotalBytesUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesUploaded

`func (o *GetVendorUserResponseData) SetTotalBytesUploaded(v int64)`

SetTotalBytesUploaded sets TotalBytesUploaded field to given value.

### HasTotalBytesUploaded

`func (o *GetVendorUserResponseData) HasTotalBytesUploaded() bool`

HasTotalBytesUploaded returns a boolean if a field has been set.

### GetTorrentsDownloaded

`func (o *GetVendorUserResponseData) GetTorrentsDownloaded() float32`

GetTorrentsDownloaded returns the TorrentsDownloaded field if non-nil, zero value otherwise.

### GetTorrentsDownloadedOk

`func (o *GetVendorUserResponseData) GetTorrentsDownloadedOk() (*float32, bool)`

GetTorrentsDownloadedOk returns a tuple with the TorrentsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentsDownloaded

`func (o *GetVendorUserResponseData) SetTorrentsDownloaded(v float32)`

SetTorrentsDownloaded sets TorrentsDownloaded field to given value.

### HasTorrentsDownloaded

`func (o *GetVendorUserResponseData) HasTorrentsDownloaded() bool`

HasTorrentsDownloaded returns a boolean if a field has been set.

### GetWebDownloadsDownloaded

`func (o *GetVendorUserResponseData) GetWebDownloadsDownloaded() float32`

GetWebDownloadsDownloaded returns the WebDownloadsDownloaded field if non-nil, zero value otherwise.

### GetWebDownloadsDownloadedOk

`func (o *GetVendorUserResponseData) GetWebDownloadsDownloadedOk() (*float32, bool)`

GetWebDownloadsDownloadedOk returns a tuple with the WebDownloadsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadsDownloaded

`func (o *GetVendorUserResponseData) SetWebDownloadsDownloaded(v float32)`

SetWebDownloadsDownloaded sets WebDownloadsDownloaded field to given value.

### HasWebDownloadsDownloaded

`func (o *GetVendorUserResponseData) HasWebDownloadsDownloaded() bool`

HasWebDownloadsDownloaded returns a boolean if a field has been set.

### GetUsenetDownloadsDownloaded

`func (o *GetVendorUserResponseData) GetUsenetDownloadsDownloaded() float32`

GetUsenetDownloadsDownloaded returns the UsenetDownloadsDownloaded field if non-nil, zero value otherwise.

### GetUsenetDownloadsDownloadedOk

`func (o *GetVendorUserResponseData) GetUsenetDownloadsDownloadedOk() (*float32, bool)`

GetUsenetDownloadsDownloadedOk returns a tuple with the UsenetDownloadsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsenetDownloadsDownloaded

`func (o *GetVendorUserResponseData) SetUsenetDownloadsDownloaded(v float32)`

SetUsenetDownloadsDownloaded sets UsenetDownloadsDownloaded field to given value.

### HasUsenetDownloadsDownloaded

`func (o *GetVendorUserResponseData) HasUsenetDownloadsDownloaded() bool`

HasUsenetDownloadsDownloaded returns a boolean if a field has been set.

### GetAdditionalConcurrentSlots

`func (o *GetVendorUserResponseData) GetAdditionalConcurrentSlots() int64`

GetAdditionalConcurrentSlots returns the AdditionalConcurrentSlots field if non-nil, zero value otherwise.

### GetAdditionalConcurrentSlotsOk

`func (o *GetVendorUserResponseData) GetAdditionalConcurrentSlotsOk() (*int64, bool)`

GetAdditionalConcurrentSlotsOk returns a tuple with the AdditionalConcurrentSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConcurrentSlots

`func (o *GetVendorUserResponseData) SetAdditionalConcurrentSlots(v int64)`

SetAdditionalConcurrentSlots sets AdditionalConcurrentSlots field to given value.

### HasAdditionalConcurrentSlots

`func (o *GetVendorUserResponseData) HasAdditionalConcurrentSlots() bool`

HasAdditionalConcurrentSlots returns a boolean if a field has been set.

### GetLongTermSeeding

`func (o *GetVendorUserResponseData) GetLongTermSeeding() bool`

GetLongTermSeeding returns the LongTermSeeding field if non-nil, zero value otherwise.

### GetLongTermSeedingOk

`func (o *GetVendorUserResponseData) GetLongTermSeedingOk() (*bool, bool)`

GetLongTermSeedingOk returns a tuple with the LongTermSeeding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSeeding

`func (o *GetVendorUserResponseData) SetLongTermSeeding(v bool)`

SetLongTermSeeding sets LongTermSeeding field to given value.

### HasLongTermSeeding

`func (o *GetVendorUserResponseData) HasLongTermSeeding() bool`

HasLongTermSeeding returns a boolean if a field has been set.

### GetVendorId

`func (o *GetVendorUserResponseData) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GetVendorUserResponseData) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GetVendorUserResponseData) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GetVendorUserResponseData) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetIsVendor

`func (o *GetVendorUserResponseData) GetIsVendor() bool`

GetIsVendor returns the IsVendor field if non-nil, zero value otherwise.

### GetIsVendorOk

`func (o *GetVendorUserResponseData) GetIsVendorOk() (*bool, bool)`

GetIsVendorOk returns a tuple with the IsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendor

`func (o *GetVendorUserResponseData) SetIsVendor(v bool)`

SetIsVendor sets IsVendor field to given value.

### HasIsVendor

`func (o *GetVendorUserResponseData) HasIsVendor() bool`

HasIsVendor returns a boolean if a field has been set.

### GetPurchasesReferred

`func (o *GetVendorUserResponseData) GetPurchasesReferred() float32`

GetPurchasesReferred returns the PurchasesReferred field if non-nil, zero value otherwise.

### GetPurchasesReferredOk

`func (o *GetVendorUserResponseData) GetPurchasesReferredOk() (*float32, bool)`

GetPurchasesReferredOk returns a tuple with the PurchasesReferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesReferred

`func (o *GetVendorUserResponseData) SetPurchasesReferred(v float32)`

SetPurchasesReferred sets PurchasesReferred field to given value.

### HasPurchasesReferred

`func (o *GetVendorUserResponseData) HasPurchasesReferred() bool`

HasPurchasesReferred returns a boolean if a field has been set.

### GetLongTermStorage

`func (o *GetVendorUserResponseData) GetLongTermStorage() bool`

GetLongTermStorage returns the LongTermStorage field if non-nil, zero value otherwise.

### GetLongTermStorageOk

`func (o *GetVendorUserResponseData) GetLongTermStorageOk() (*bool, bool)`

GetLongTermStorageOk returns a tuple with the LongTermStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermStorage

`func (o *GetVendorUserResponseData) SetLongTermStorage(v bool)`

SetLongTermStorage sets LongTermStorage field to given value.

### HasLongTermStorage

`func (o *GetVendorUserResponseData) HasLongTermStorage() bool`

HasLongTermStorage returns a boolean if a field has been set.

### GetDailyBandwidthUsage

`func (o *GetVendorUserResponseData) GetDailyBandwidthUsage() int64`

GetDailyBandwidthUsage returns the DailyBandwidthUsage field if non-nil, zero value otherwise.

### GetDailyBandwidthUsageOk

`func (o *GetVendorUserResponseData) GetDailyBandwidthUsageOk() (*int64, bool)`

GetDailyBandwidthUsageOk returns a tuple with the DailyBandwidthUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBandwidthUsage

`func (o *GetVendorUserResponseData) SetDailyBandwidthUsage(v int64)`

SetDailyBandwidthUsage sets DailyBandwidthUsage field to given value.

### HasDailyBandwidthUsage

`func (o *GetVendorUserResponseData) HasDailyBandwidthUsage() bool`

HasDailyBandwidthUsage returns a boolean if a field has been set.

### GetApiToken

`func (o *GetVendorUserResponseData) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *GetVendorUserResponseData) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *GetVendorUserResponseData) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *GetVendorUserResponseData) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


