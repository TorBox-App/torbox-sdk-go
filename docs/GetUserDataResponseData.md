# GetUserDataResponseData

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
**Settings** | Pointer to [**GetUserDataResponseDataSettings**](GetUserDataResponseDataSettings.md) |  | [optional] 

## Methods

### NewGetUserDataResponseData

`func NewGetUserDataResponseData() *GetUserDataResponseData`

NewGetUserDataResponseData instantiates a new GetUserDataResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDataResponseDataWithDefaults

`func NewGetUserDataResponseDataWithDefaults() *GetUserDataResponseData`

NewGetUserDataResponseDataWithDefaults instantiates a new GetUserDataResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUserDataResponseData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserDataResponseData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserDataResponseData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserDataResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserReferral

`func (o *GetUserDataResponseData) GetUserReferral() string`

GetUserReferral returns the UserReferral field if non-nil, zero value otherwise.

### GetUserReferralOk

`func (o *GetUserDataResponseData) GetUserReferralOk() (*string, bool)`

GetUserReferralOk returns a tuple with the UserReferral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserReferral

`func (o *GetUserDataResponseData) SetUserReferral(v string)`

SetUserReferral sets UserReferral field to given value.

### HasUserReferral

`func (o *GetUserDataResponseData) HasUserReferral() bool`

HasUserReferral returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetUserDataResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUserDataResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUserDataResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetUserDataResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetUserDataResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetUserDataResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetUserDataResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetUserDataResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *GetUserDataResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUserDataResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUserDataResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUserDataResponseData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPlan

`func (o *GetUserDataResponseData) GetPlan() int64`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetUserDataResponseData) GetPlanOk() (*int64, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetUserDataResponseData) SetPlan(v int64)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetUserDataResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetTotalDownloaded

`func (o *GetUserDataResponseData) GetTotalDownloaded() int64`

GetTotalDownloaded returns the TotalDownloaded field if non-nil, zero value otherwise.

### GetTotalDownloadedOk

`func (o *GetUserDataResponseData) GetTotalDownloadedOk() (*int64, bool)`

GetTotalDownloadedOk returns a tuple with the TotalDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownloaded

`func (o *GetUserDataResponseData) SetTotalDownloaded(v int64)`

SetTotalDownloaded sets TotalDownloaded field to given value.

### HasTotalDownloaded

`func (o *GetUserDataResponseData) HasTotalDownloaded() bool`

HasTotalDownloaded returns a boolean if a field has been set.

### GetCustomer

`func (o *GetUserDataResponseData) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GetUserDataResponseData) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GetUserDataResponseData) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GetUserDataResponseData) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *GetUserDataResponseData) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *GetUserDataResponseData) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *GetUserDataResponseData) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *GetUserDataResponseData) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetPremiumExpiresAt

`func (o *GetUserDataResponseData) GetPremiumExpiresAt() string`

GetPremiumExpiresAt returns the PremiumExpiresAt field if non-nil, zero value otherwise.

### GetPremiumExpiresAtOk

`func (o *GetUserDataResponseData) GetPremiumExpiresAtOk() (*string, bool)`

GetPremiumExpiresAtOk returns a tuple with the PremiumExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumExpiresAt

`func (o *GetUserDataResponseData) SetPremiumExpiresAt(v string)`

SetPremiumExpiresAt sets PremiumExpiresAt field to given value.

### HasPremiumExpiresAt

`func (o *GetUserDataResponseData) HasPremiumExpiresAt() bool`

HasPremiumExpiresAt returns a boolean if a field has been set.

### GetCooldownUntil

`func (o *GetUserDataResponseData) GetCooldownUntil() string`

GetCooldownUntil returns the CooldownUntil field if non-nil, zero value otherwise.

### GetCooldownUntilOk

`func (o *GetUserDataResponseData) GetCooldownUntilOk() (*string, bool)`

GetCooldownUntilOk returns a tuple with the CooldownUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownUntil

`func (o *GetUserDataResponseData) SetCooldownUntil(v string)`

SetCooldownUntil sets CooldownUntil field to given value.

### HasCooldownUntil

`func (o *GetUserDataResponseData) HasCooldownUntil() bool`

HasCooldownUntil returns a boolean if a field has been set.

### GetAuthId

`func (o *GetUserDataResponseData) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *GetUserDataResponseData) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *GetUserDataResponseData) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *GetUserDataResponseData) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.

### GetBaseEmail

`func (o *GetUserDataResponseData) GetBaseEmail() string`

GetBaseEmail returns the BaseEmail field if non-nil, zero value otherwise.

### GetBaseEmailOk

`func (o *GetUserDataResponseData) GetBaseEmailOk() (*string, bool)`

GetBaseEmailOk returns a tuple with the BaseEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEmail

`func (o *GetUserDataResponseData) SetBaseEmail(v string)`

SetBaseEmail sets BaseEmail field to given value.

### HasBaseEmail

`func (o *GetUserDataResponseData) HasBaseEmail() bool`

HasBaseEmail returns a boolean if a field has been set.

### GetTotalBytesDownloaded

`func (o *GetUserDataResponseData) GetTotalBytesDownloaded() int64`

GetTotalBytesDownloaded returns the TotalBytesDownloaded field if non-nil, zero value otherwise.

### GetTotalBytesDownloadedOk

`func (o *GetUserDataResponseData) GetTotalBytesDownloadedOk() (*int64, bool)`

GetTotalBytesDownloadedOk returns a tuple with the TotalBytesDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesDownloaded

`func (o *GetUserDataResponseData) SetTotalBytesDownloaded(v int64)`

SetTotalBytesDownloaded sets TotalBytesDownloaded field to given value.

### HasTotalBytesDownloaded

`func (o *GetUserDataResponseData) HasTotalBytesDownloaded() bool`

HasTotalBytesDownloaded returns a boolean if a field has been set.

### GetTotalBytesUploaded

`func (o *GetUserDataResponseData) GetTotalBytesUploaded() int64`

GetTotalBytesUploaded returns the TotalBytesUploaded field if non-nil, zero value otherwise.

### GetTotalBytesUploadedOk

`func (o *GetUserDataResponseData) GetTotalBytesUploadedOk() (*int64, bool)`

GetTotalBytesUploadedOk returns a tuple with the TotalBytesUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesUploaded

`func (o *GetUserDataResponseData) SetTotalBytesUploaded(v int64)`

SetTotalBytesUploaded sets TotalBytesUploaded field to given value.

### HasTotalBytesUploaded

`func (o *GetUserDataResponseData) HasTotalBytesUploaded() bool`

HasTotalBytesUploaded returns a boolean if a field has been set.

### GetTorrentsDownloaded

`func (o *GetUserDataResponseData) GetTorrentsDownloaded() float32`

GetTorrentsDownloaded returns the TorrentsDownloaded field if non-nil, zero value otherwise.

### GetTorrentsDownloadedOk

`func (o *GetUserDataResponseData) GetTorrentsDownloadedOk() (*float32, bool)`

GetTorrentsDownloadedOk returns a tuple with the TorrentsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentsDownloaded

`func (o *GetUserDataResponseData) SetTorrentsDownloaded(v float32)`

SetTorrentsDownloaded sets TorrentsDownloaded field to given value.

### HasTorrentsDownloaded

`func (o *GetUserDataResponseData) HasTorrentsDownloaded() bool`

HasTorrentsDownloaded returns a boolean if a field has been set.

### GetWebDownloadsDownloaded

`func (o *GetUserDataResponseData) GetWebDownloadsDownloaded() float32`

GetWebDownloadsDownloaded returns the WebDownloadsDownloaded field if non-nil, zero value otherwise.

### GetWebDownloadsDownloadedOk

`func (o *GetUserDataResponseData) GetWebDownloadsDownloadedOk() (*float32, bool)`

GetWebDownloadsDownloadedOk returns a tuple with the WebDownloadsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadsDownloaded

`func (o *GetUserDataResponseData) SetWebDownloadsDownloaded(v float32)`

SetWebDownloadsDownloaded sets WebDownloadsDownloaded field to given value.

### HasWebDownloadsDownloaded

`func (o *GetUserDataResponseData) HasWebDownloadsDownloaded() bool`

HasWebDownloadsDownloaded returns a boolean if a field has been set.

### GetUsenetDownloadsDownloaded

`func (o *GetUserDataResponseData) GetUsenetDownloadsDownloaded() float32`

GetUsenetDownloadsDownloaded returns the UsenetDownloadsDownloaded field if non-nil, zero value otherwise.

### GetUsenetDownloadsDownloadedOk

`func (o *GetUserDataResponseData) GetUsenetDownloadsDownloadedOk() (*float32, bool)`

GetUsenetDownloadsDownloadedOk returns a tuple with the UsenetDownloadsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsenetDownloadsDownloaded

`func (o *GetUserDataResponseData) SetUsenetDownloadsDownloaded(v float32)`

SetUsenetDownloadsDownloaded sets UsenetDownloadsDownloaded field to given value.

### HasUsenetDownloadsDownloaded

`func (o *GetUserDataResponseData) HasUsenetDownloadsDownloaded() bool`

HasUsenetDownloadsDownloaded returns a boolean if a field has been set.

### GetAdditionalConcurrentSlots

`func (o *GetUserDataResponseData) GetAdditionalConcurrentSlots() int64`

GetAdditionalConcurrentSlots returns the AdditionalConcurrentSlots field if non-nil, zero value otherwise.

### GetAdditionalConcurrentSlotsOk

`func (o *GetUserDataResponseData) GetAdditionalConcurrentSlotsOk() (*int64, bool)`

GetAdditionalConcurrentSlotsOk returns a tuple with the AdditionalConcurrentSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalConcurrentSlots

`func (o *GetUserDataResponseData) SetAdditionalConcurrentSlots(v int64)`

SetAdditionalConcurrentSlots sets AdditionalConcurrentSlots field to given value.

### HasAdditionalConcurrentSlots

`func (o *GetUserDataResponseData) HasAdditionalConcurrentSlots() bool`

HasAdditionalConcurrentSlots returns a boolean if a field has been set.

### GetLongTermSeeding

`func (o *GetUserDataResponseData) GetLongTermSeeding() bool`

GetLongTermSeeding returns the LongTermSeeding field if non-nil, zero value otherwise.

### GetLongTermSeedingOk

`func (o *GetUserDataResponseData) GetLongTermSeedingOk() (*bool, bool)`

GetLongTermSeedingOk returns a tuple with the LongTermSeeding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermSeeding

`func (o *GetUserDataResponseData) SetLongTermSeeding(v bool)`

SetLongTermSeeding sets LongTermSeeding field to given value.

### HasLongTermSeeding

`func (o *GetUserDataResponseData) HasLongTermSeeding() bool`

HasLongTermSeeding returns a boolean if a field has been set.

### GetVendorId

`func (o *GetUserDataResponseData) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GetUserDataResponseData) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GetUserDataResponseData) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GetUserDataResponseData) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetIsVendor

`func (o *GetUserDataResponseData) GetIsVendor() bool`

GetIsVendor returns the IsVendor field if non-nil, zero value otherwise.

### GetIsVendorOk

`func (o *GetUserDataResponseData) GetIsVendorOk() (*bool, bool)`

GetIsVendorOk returns a tuple with the IsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendor

`func (o *GetUserDataResponseData) SetIsVendor(v bool)`

SetIsVendor sets IsVendor field to given value.

### HasIsVendor

`func (o *GetUserDataResponseData) HasIsVendor() bool`

HasIsVendor returns a boolean if a field has been set.

### GetPurchasesReferred

`func (o *GetUserDataResponseData) GetPurchasesReferred() float32`

GetPurchasesReferred returns the PurchasesReferred field if non-nil, zero value otherwise.

### GetPurchasesReferredOk

`func (o *GetUserDataResponseData) GetPurchasesReferredOk() (*float32, bool)`

GetPurchasesReferredOk returns a tuple with the PurchasesReferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesReferred

`func (o *GetUserDataResponseData) SetPurchasesReferred(v float32)`

SetPurchasesReferred sets PurchasesReferred field to given value.

### HasPurchasesReferred

`func (o *GetUserDataResponseData) HasPurchasesReferred() bool`

HasPurchasesReferred returns a boolean if a field has been set.

### GetLongTermStorage

`func (o *GetUserDataResponseData) GetLongTermStorage() bool`

GetLongTermStorage returns the LongTermStorage field if non-nil, zero value otherwise.

### GetLongTermStorageOk

`func (o *GetUserDataResponseData) GetLongTermStorageOk() (*bool, bool)`

GetLongTermStorageOk returns a tuple with the LongTermStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermStorage

`func (o *GetUserDataResponseData) SetLongTermStorage(v bool)`

SetLongTermStorage sets LongTermStorage field to given value.

### HasLongTermStorage

`func (o *GetUserDataResponseData) HasLongTermStorage() bool`

HasLongTermStorage returns a boolean if a field has been set.

### GetSettings

`func (o *GetUserDataResponseData) GetSettings() GetUserDataResponseDataSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GetUserDataResponseData) GetSettingsOk() (*GetUserDataResponseDataSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GetUserDataResponseData) SetSettings(v GetUserDataResponseDataSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GetUserDataResponseData) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


