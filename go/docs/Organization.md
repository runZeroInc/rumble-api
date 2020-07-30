# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**DownloadToken** | Pointer to **string** |  | [optional] 
**DownloadTokenCreatedAt** | Pointer to **int64** |  | [optional] 
**Permanent** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**DeactivatedAt** | Pointer to **int64** |  | [optional] 
**ServiceCount** | Pointer to **int64** |  | [optional] 
**ServiceCountTcp** | Pointer to **int64** |  | [optional] 
**ServiceCountUdp** | Pointer to **int64** |  | [optional] 
**ServiceCountArp** | Pointer to **int64** |  | [optional] 
**ServiceCountIcmp** | Pointer to **int64** |  | [optional] 
**AssetCount** | Pointer to **int64** |  | [optional] 
**ExportToken** | Pointer to **string** |  | [optional] 
**ExportTokenCreatedAt** | Pointer to **int64** |  | [optional] 
**ExportTokenLastUsedAt** | Pointer to **int64** |  | [optional] 
**ExportTokenLastUsedBy** | Pointer to **string** |  | [optional] 
**ExportTokenCounter** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization(id string, name string, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Organization) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Organization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Organization) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Organization) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Organization) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Organization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClientId

`func (o *Organization) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Organization) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Organization) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Organization) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDownloadToken

`func (o *Organization) GetDownloadToken() string`

GetDownloadToken returns the DownloadToken field if non-nil, zero value otherwise.

### GetDownloadTokenOk

`func (o *Organization) GetDownloadTokenOk() (*string, bool)`

GetDownloadTokenOk returns a tuple with the DownloadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadToken

`func (o *Organization) SetDownloadToken(v string)`

SetDownloadToken sets DownloadToken field to given value.

### HasDownloadToken

`func (o *Organization) HasDownloadToken() bool`

HasDownloadToken returns a boolean if a field has been set.

### GetDownloadTokenCreatedAt

`func (o *Organization) GetDownloadTokenCreatedAt() int64`

GetDownloadTokenCreatedAt returns the DownloadTokenCreatedAt field if non-nil, zero value otherwise.

### GetDownloadTokenCreatedAtOk

`func (o *Organization) GetDownloadTokenCreatedAtOk() (*int64, bool)`

GetDownloadTokenCreatedAtOk returns a tuple with the DownloadTokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadTokenCreatedAt

`func (o *Organization) SetDownloadTokenCreatedAt(v int64)`

SetDownloadTokenCreatedAt sets DownloadTokenCreatedAt field to given value.

### HasDownloadTokenCreatedAt

`func (o *Organization) HasDownloadTokenCreatedAt() bool`

HasDownloadTokenCreatedAt returns a boolean if a field has been set.

### GetPermanent

`func (o *Organization) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *Organization) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *Organization) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *Organization) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInactive

`func (o *Organization) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *Organization) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *Organization) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *Organization) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetDeactivatedAt

`func (o *Organization) GetDeactivatedAt() int64`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *Organization) GetDeactivatedAtOk() (*int64, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *Organization) SetDeactivatedAt(v int64)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *Organization) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.

### GetServiceCount

`func (o *Organization) GetServiceCount() int64`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *Organization) GetServiceCountOk() (*int64, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCount

`func (o *Organization) SetServiceCount(v int64)`

SetServiceCount sets ServiceCount field to given value.

### HasServiceCount

`func (o *Organization) HasServiceCount() bool`

HasServiceCount returns a boolean if a field has been set.

### GetServiceCountTcp

`func (o *Organization) GetServiceCountTcp() int64`

GetServiceCountTcp returns the ServiceCountTcp field if non-nil, zero value otherwise.

### GetServiceCountTcpOk

`func (o *Organization) GetServiceCountTcpOk() (*int64, bool)`

GetServiceCountTcpOk returns a tuple with the ServiceCountTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountTcp

`func (o *Organization) SetServiceCountTcp(v int64)`

SetServiceCountTcp sets ServiceCountTcp field to given value.

### HasServiceCountTcp

`func (o *Organization) HasServiceCountTcp() bool`

HasServiceCountTcp returns a boolean if a field has been set.

### GetServiceCountUdp

`func (o *Organization) GetServiceCountUdp() int64`

GetServiceCountUdp returns the ServiceCountUdp field if non-nil, zero value otherwise.

### GetServiceCountUdpOk

`func (o *Organization) GetServiceCountUdpOk() (*int64, bool)`

GetServiceCountUdpOk returns a tuple with the ServiceCountUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountUdp

`func (o *Organization) SetServiceCountUdp(v int64)`

SetServiceCountUdp sets ServiceCountUdp field to given value.

### HasServiceCountUdp

`func (o *Organization) HasServiceCountUdp() bool`

HasServiceCountUdp returns a boolean if a field has been set.

### GetServiceCountArp

`func (o *Organization) GetServiceCountArp() int64`

GetServiceCountArp returns the ServiceCountArp field if non-nil, zero value otherwise.

### GetServiceCountArpOk

`func (o *Organization) GetServiceCountArpOk() (*int64, bool)`

GetServiceCountArpOk returns a tuple with the ServiceCountArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountArp

`func (o *Organization) SetServiceCountArp(v int64)`

SetServiceCountArp sets ServiceCountArp field to given value.

### HasServiceCountArp

`func (o *Organization) HasServiceCountArp() bool`

HasServiceCountArp returns a boolean if a field has been set.

### GetServiceCountIcmp

`func (o *Organization) GetServiceCountIcmp() int64`

GetServiceCountIcmp returns the ServiceCountIcmp field if non-nil, zero value otherwise.

### GetServiceCountIcmpOk

`func (o *Organization) GetServiceCountIcmpOk() (*int64, bool)`

GetServiceCountIcmpOk returns a tuple with the ServiceCountIcmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountIcmp

`func (o *Organization) SetServiceCountIcmp(v int64)`

SetServiceCountIcmp sets ServiceCountIcmp field to given value.

### HasServiceCountIcmp

`func (o *Organization) HasServiceCountIcmp() bool`

HasServiceCountIcmp returns a boolean if a field has been set.

### GetAssetCount

`func (o *Organization) GetAssetCount() int64`

GetAssetCount returns the AssetCount field if non-nil, zero value otherwise.

### GetAssetCountOk

`func (o *Organization) GetAssetCountOk() (*int64, bool)`

GetAssetCountOk returns a tuple with the AssetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCount

`func (o *Organization) SetAssetCount(v int64)`

SetAssetCount sets AssetCount field to given value.

### HasAssetCount

`func (o *Organization) HasAssetCount() bool`

HasAssetCount returns a boolean if a field has been set.

### GetExportToken

`func (o *Organization) GetExportToken() string`

GetExportToken returns the ExportToken field if non-nil, zero value otherwise.

### GetExportTokenOk

`func (o *Organization) GetExportTokenOk() (*string, bool)`

GetExportTokenOk returns a tuple with the ExportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToken

`func (o *Organization) SetExportToken(v string)`

SetExportToken sets ExportToken field to given value.

### HasExportToken

`func (o *Organization) HasExportToken() bool`

HasExportToken returns a boolean if a field has been set.

### GetExportTokenCreatedAt

`func (o *Organization) GetExportTokenCreatedAt() int64`

GetExportTokenCreatedAt returns the ExportTokenCreatedAt field if non-nil, zero value otherwise.

### GetExportTokenCreatedAtOk

`func (o *Organization) GetExportTokenCreatedAtOk() (*int64, bool)`

GetExportTokenCreatedAtOk returns a tuple with the ExportTokenCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTokenCreatedAt

`func (o *Organization) SetExportTokenCreatedAt(v int64)`

SetExportTokenCreatedAt sets ExportTokenCreatedAt field to given value.

### HasExportTokenCreatedAt

`func (o *Organization) HasExportTokenCreatedAt() bool`

HasExportTokenCreatedAt returns a boolean if a field has been set.

### GetExportTokenLastUsedAt

`func (o *Organization) GetExportTokenLastUsedAt() int64`

GetExportTokenLastUsedAt returns the ExportTokenLastUsedAt field if non-nil, zero value otherwise.

### GetExportTokenLastUsedAtOk

`func (o *Organization) GetExportTokenLastUsedAtOk() (*int64, bool)`

GetExportTokenLastUsedAtOk returns a tuple with the ExportTokenLastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTokenLastUsedAt

`func (o *Organization) SetExportTokenLastUsedAt(v int64)`

SetExportTokenLastUsedAt sets ExportTokenLastUsedAt field to given value.

### HasExportTokenLastUsedAt

`func (o *Organization) HasExportTokenLastUsedAt() bool`

HasExportTokenLastUsedAt returns a boolean if a field has been set.

### GetExportTokenLastUsedBy

`func (o *Organization) GetExportTokenLastUsedBy() string`

GetExportTokenLastUsedBy returns the ExportTokenLastUsedBy field if non-nil, zero value otherwise.

### GetExportTokenLastUsedByOk

`func (o *Organization) GetExportTokenLastUsedByOk() (*string, bool)`

GetExportTokenLastUsedByOk returns a tuple with the ExportTokenLastUsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTokenLastUsedBy

`func (o *Organization) SetExportTokenLastUsedBy(v string)`

SetExportTokenLastUsedBy sets ExportTokenLastUsedBy field to given value.

### HasExportTokenLastUsedBy

`func (o *Organization) HasExportTokenLastUsedBy() bool`

HasExportTokenLastUsedBy returns a boolean if a field has been set.

### GetExportTokenCounter

`func (o *Organization) GetExportTokenCounter() int64`

GetExportTokenCounter returns the ExportTokenCounter field if non-nil, zero value otherwise.

### GetExportTokenCounterOk

`func (o *Organization) GetExportTokenCounterOk() (*int64, bool)`

GetExportTokenCounterOk returns a tuple with the ExportTokenCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTokenCounter

`func (o *Organization) SetExportTokenCounter(v int64)`

SetExportTokenCounter sets ExportTokenCounter field to given value.

### HasExportTokenCounter

`func (o *Organization) HasExportTokenCounter() bool`

HasExportTokenCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


