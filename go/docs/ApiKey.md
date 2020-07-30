# APIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**LastUsedAt** | Pointer to **int64** |  | [optional] 
**LastUsedIp** | Pointer to **string** |  | [optional] 
**LastUsedUa** | Pointer to **string** |  | [optional] 
**Counter** | Pointer to **int64** |  | [optional] 
**UsageToday** | Pointer to **int64** |  | [optional] 
**UsageLimit** | Pointer to **int64** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 

## Methods

### NewAPIKey

`func NewAPIKey(id string, ) *APIKey`

NewAPIKey instantiates a new APIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyWithDefaults

`func NewAPIKeyWithDefaults() *APIKey`

NewAPIKeyWithDefaults instantiates a new APIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIKey) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *APIKey) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *APIKey) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *APIKey) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *APIKey) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *APIKey) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *APIKey) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *APIKey) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *APIKey) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *APIKey) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIKey) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIKey) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *APIKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *APIKey) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APIKey) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APIKey) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *APIKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetComment

`func (o *APIKey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *APIKey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *APIKey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *APIKey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *APIKey) GetLastUsedAt() int64`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *APIKey) GetLastUsedAtOk() (*int64, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *APIKey) SetLastUsedAt(v int64)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *APIKey) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLastUsedIp

`func (o *APIKey) GetLastUsedIp() string`

GetLastUsedIp returns the LastUsedIp field if non-nil, zero value otherwise.

### GetLastUsedIpOk

`func (o *APIKey) GetLastUsedIpOk() (*string, bool)`

GetLastUsedIpOk returns a tuple with the LastUsedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedIp

`func (o *APIKey) SetLastUsedIp(v string)`

SetLastUsedIp sets LastUsedIp field to given value.

### HasLastUsedIp

`func (o *APIKey) HasLastUsedIp() bool`

HasLastUsedIp returns a boolean if a field has been set.

### GetLastUsedUa

`func (o *APIKey) GetLastUsedUa() string`

GetLastUsedUa returns the LastUsedUa field if non-nil, zero value otherwise.

### GetLastUsedUaOk

`func (o *APIKey) GetLastUsedUaOk() (*string, bool)`

GetLastUsedUaOk returns a tuple with the LastUsedUa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedUa

`func (o *APIKey) SetLastUsedUa(v string)`

SetLastUsedUa sets LastUsedUa field to given value.

### HasLastUsedUa

`func (o *APIKey) HasLastUsedUa() bool`

HasLastUsedUa returns a boolean if a field has been set.

### GetCounter

`func (o *APIKey) GetCounter() int64`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *APIKey) GetCounterOk() (*int64, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *APIKey) SetCounter(v int64)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *APIKey) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetUsageToday

`func (o *APIKey) GetUsageToday() int64`

GetUsageToday returns the UsageToday field if non-nil, zero value otherwise.

### GetUsageTodayOk

`func (o *APIKey) GetUsageTodayOk() (*int64, bool)`

GetUsageTodayOk returns a tuple with the UsageToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageToday

`func (o *APIKey) SetUsageToday(v int64)`

SetUsageToday sets UsageToday field to given value.

### HasUsageToday

`func (o *APIKey) HasUsageToday() bool`

HasUsageToday returns a boolean if a field has been set.

### GetUsageLimit

`func (o *APIKey) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *APIKey) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *APIKey) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *APIKey) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetToken

`func (o *APIKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *APIKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *APIKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *APIKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetInactive

`func (o *APIKey) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *APIKey) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *APIKey) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *APIKey) HasInactive() bool`

HasInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


