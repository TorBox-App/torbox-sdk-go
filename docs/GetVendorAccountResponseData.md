# GetVendorAccountResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **int64** |  | [optional] 
**VendorToken** | Pointer to **string** |  | [optional] 
**VendorName** | Pointer to **string** |  | [optional] 
**VendorEmail** | Pointer to **string** |  | [optional] 
**VendorUrl** | Pointer to **string** |  | [optional] 
**VendorStatus** | Pointer to **string** |  | [optional] 
**UsersAllowed** | Pointer to **float32** |  | [optional] 
**CurrentUsers** | Pointer to **float32** |  | [optional] 
**CanRegisterNew** | Pointer to **bool** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **int64** |  | [optional] 
**PaidUntil** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**CanGenerateVouchers** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetVendorAccountResponseData

`func NewGetVendorAccountResponseData() *GetVendorAccountResponseData`

NewGetVendorAccountResponseData instantiates a new GetVendorAccountResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVendorAccountResponseDataWithDefaults

`func NewGetVendorAccountResponseDataWithDefaults() *GetVendorAccountResponseData`

NewGetVendorAccountResponseDataWithDefaults instantiates a new GetVendorAccountResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *GetVendorAccountResponseData) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GetVendorAccountResponseData) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GetVendorAccountResponseData) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GetVendorAccountResponseData) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorToken

`func (o *GetVendorAccountResponseData) GetVendorToken() string`

GetVendorToken returns the VendorToken field if non-nil, zero value otherwise.

### GetVendorTokenOk

`func (o *GetVendorAccountResponseData) GetVendorTokenOk() (*string, bool)`

GetVendorTokenOk returns a tuple with the VendorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorToken

`func (o *GetVendorAccountResponseData) SetVendorToken(v string)`

SetVendorToken sets VendorToken field to given value.

### HasVendorToken

`func (o *GetVendorAccountResponseData) HasVendorToken() bool`

HasVendorToken returns a boolean if a field has been set.

### GetVendorName

`func (o *GetVendorAccountResponseData) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *GetVendorAccountResponseData) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *GetVendorAccountResponseData) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *GetVendorAccountResponseData) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetVendorEmail

`func (o *GetVendorAccountResponseData) GetVendorEmail() string`

GetVendorEmail returns the VendorEmail field if non-nil, zero value otherwise.

### GetVendorEmailOk

`func (o *GetVendorAccountResponseData) GetVendorEmailOk() (*string, bool)`

GetVendorEmailOk returns a tuple with the VendorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorEmail

`func (o *GetVendorAccountResponseData) SetVendorEmail(v string)`

SetVendorEmail sets VendorEmail field to given value.

### HasVendorEmail

`func (o *GetVendorAccountResponseData) HasVendorEmail() bool`

HasVendorEmail returns a boolean if a field has been set.

### GetVendorUrl

`func (o *GetVendorAccountResponseData) GetVendorUrl() string`

GetVendorUrl returns the VendorUrl field if non-nil, zero value otherwise.

### GetVendorUrlOk

`func (o *GetVendorAccountResponseData) GetVendorUrlOk() (*string, bool)`

GetVendorUrlOk returns a tuple with the VendorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorUrl

`func (o *GetVendorAccountResponseData) SetVendorUrl(v string)`

SetVendorUrl sets VendorUrl field to given value.

### HasVendorUrl

`func (o *GetVendorAccountResponseData) HasVendorUrl() bool`

HasVendorUrl returns a boolean if a field has been set.

### GetVendorStatus

`func (o *GetVendorAccountResponseData) GetVendorStatus() string`

GetVendorStatus returns the VendorStatus field if non-nil, zero value otherwise.

### GetVendorStatusOk

`func (o *GetVendorAccountResponseData) GetVendorStatusOk() (*string, bool)`

GetVendorStatusOk returns a tuple with the VendorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorStatus

`func (o *GetVendorAccountResponseData) SetVendorStatus(v string)`

SetVendorStatus sets VendorStatus field to given value.

### HasVendorStatus

`func (o *GetVendorAccountResponseData) HasVendorStatus() bool`

HasVendorStatus returns a boolean if a field has been set.

### GetUsersAllowed

`func (o *GetVendorAccountResponseData) GetUsersAllowed() float32`

GetUsersAllowed returns the UsersAllowed field if non-nil, zero value otherwise.

### GetUsersAllowedOk

`func (o *GetVendorAccountResponseData) GetUsersAllowedOk() (*float32, bool)`

GetUsersAllowedOk returns a tuple with the UsersAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersAllowed

`func (o *GetVendorAccountResponseData) SetUsersAllowed(v float32)`

SetUsersAllowed sets UsersAllowed field to given value.

### HasUsersAllowed

`func (o *GetVendorAccountResponseData) HasUsersAllowed() bool`

HasUsersAllowed returns a boolean if a field has been set.

### GetCurrentUsers

`func (o *GetVendorAccountResponseData) GetCurrentUsers() float32`

GetCurrentUsers returns the CurrentUsers field if non-nil, zero value otherwise.

### GetCurrentUsersOk

`func (o *GetVendorAccountResponseData) GetCurrentUsersOk() (*float32, bool)`

GetCurrentUsersOk returns a tuple with the CurrentUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsers

`func (o *GetVendorAccountResponseData) SetCurrentUsers(v float32)`

SetCurrentUsers sets CurrentUsers field to given value.

### HasCurrentUsers

`func (o *GetVendorAccountResponseData) HasCurrentUsers() bool`

HasCurrentUsers returns a boolean if a field has been set.

### GetCanRegisterNew

`func (o *GetVendorAccountResponseData) GetCanRegisterNew() bool`

GetCanRegisterNew returns the CanRegisterNew field if non-nil, zero value otherwise.

### GetCanRegisterNewOk

`func (o *GetVendorAccountResponseData) GetCanRegisterNewOk() (*bool, bool)`

GetCanRegisterNewOk returns a tuple with the CanRegisterNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRegisterNew

`func (o *GetVendorAccountResponseData) SetCanRegisterNew(v bool)`

SetCanRegisterNew sets CanRegisterNew field to given value.

### HasCanRegisterNew

`func (o *GetVendorAccountResponseData) HasCanRegisterNew() bool`

HasCanRegisterNew returns a boolean if a field has been set.

### GetAccountType

`func (o *GetVendorAccountResponseData) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetVendorAccountResponseData) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetVendorAccountResponseData) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetVendorAccountResponseData) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetPlan

`func (o *GetVendorAccountResponseData) GetPlan() int64`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetVendorAccountResponseData) GetPlanOk() (*int64, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetVendorAccountResponseData) SetPlan(v int64)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetVendorAccountResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPaidUntil

`func (o *GetVendorAccountResponseData) GetPaidUntil() string`

GetPaidUntil returns the PaidUntil field if non-nil, zero value otherwise.

### GetPaidUntilOk

`func (o *GetVendorAccountResponseData) GetPaidUntilOk() (*string, bool)`

GetPaidUntilOk returns a tuple with the PaidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidUntil

`func (o *GetVendorAccountResponseData) SetPaidUntil(v string)`

SetPaidUntil sets PaidUntil field to given value.

### HasPaidUntil

`func (o *GetVendorAccountResponseData) HasPaidUntil() bool`

HasPaidUntil returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetVendorAccountResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetVendorAccountResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetVendorAccountResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetVendorAccountResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetVendorAccountResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetVendorAccountResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetVendorAccountResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetVendorAccountResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCanGenerateVouchers

`func (o *GetVendorAccountResponseData) GetCanGenerateVouchers() bool`

GetCanGenerateVouchers returns the CanGenerateVouchers field if non-nil, zero value otherwise.

### GetCanGenerateVouchersOk

`func (o *GetVendorAccountResponseData) GetCanGenerateVouchersOk() (*bool, bool)`

GetCanGenerateVouchersOk returns a tuple with the CanGenerateVouchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanGenerateVouchers

`func (o *GetVendorAccountResponseData) SetCanGenerateVouchers(v bool)`

SetCanGenerateVouchers sets CanGenerateVouchers field to given value.

### HasCanGenerateVouchers

`func (o *GetVendorAccountResponseData) HasCanGenerateVouchers() bool`

HasCanGenerateVouchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


