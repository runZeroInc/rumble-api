# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**CreatedAt** | Pointer to **float32** |  | [optional] 
**UpdatedAt** | Pointer to **float32** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Alive** | Pointer to **bool** |  | [optional] 
**FirstSeen** | Pointer to **float32** |  | [optional] 
**LastSeen** | Pointer to **float32** |  | [optional] 
**DetectedBy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Hw** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to **[]string** |  | [optional] 
**AddressesExtra** | Pointer to **[]string** |  | [optional] 
**Macs** | Pointer to **[]string** |  | [optional] 
**MacVendors** | Pointer to **[]string** |  | [optional] 
**Names** | Pointer to **[]string** |  | [optional] 
**Domains** | Pointer to **[]string** |  | [optional] 
**ServiceCount** | Pointer to **float32** |  | [optional] 
**ServiceCountTcp** | Pointer to **float32** |  | [optional] 
**ServiceCountUdp** | Pointer to **float32** |  | [optional] 
**ServiceCountArp** | Pointer to **float32** |  | [optional] 
**ServiceCountIcmp** | Pointer to **float32** |  | [optional] 
**LowestTtl** | Pointer to **float32** |  | [optional] 
**LowestRtt** | Pointer to **float32** |  | [optional] 
**LastAgentId** | Pointer to **string** |  | [optional] 
**LastTaskId** | Pointer to **string** |  | [optional] 
**NewestMac** | Pointer to **string** |  | [optional] 
**NewestMacVendor** | Pointer to **string** |  | [optional] 
**NewestMacAge** | Pointer to **float32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ServicePortsTcp** | Pointer to **[]string** |  | [optional] 
**ServicePortsUdp** | Pointer to **[]string** |  | [optional] 
**ServicePortsProtocols** | Pointer to **[]string** |  | [optional] 
**ServicePortsProducts** | Pointer to **[]string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**AgentName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Services** | Pointer to [**map[string]map[string]string**](map.md) |  | [optional] 
**Rtts** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Credentials** | Pointer to [**map[string]map[string]bool**](map.md) |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedAt

`func (o *Asset) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Asset) GetCreatedAtOk() (float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *Asset) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *Asset) SetCreatedAt(v float32)`

SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.

### GetUpdatedAt

`func (o *Asset) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Asset) GetUpdatedAtOk() (float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedAt

`func (o *Asset) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAt

`func (o *Asset) SetUpdatedAt(v float32)`

SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.

### GetOrganizationId

`func (o *Asset) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Asset) GetOrganizationIdOk() (string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationId

`func (o *Asset) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationId

`func (o *Asset) SetOrganizationId(v string)`

SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.

### GetSiteId

`func (o *Asset) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Asset) GetSiteIdOk() (string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteId

`func (o *Asset) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteId

`func (o *Asset) SetSiteId(v string)`

SetSiteId gets a reference to the given string and assigns it to the SiteId field.

### GetAlive

`func (o *Asset) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *Asset) GetAliveOk() (bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlive

`func (o *Asset) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### SetAlive

`func (o *Asset) SetAlive(v bool)`

SetAlive gets a reference to the given bool and assigns it to the Alive field.

### GetFirstSeen

`func (o *Asset) GetFirstSeen() float32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *Asset) GetFirstSeenOk() (float32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstSeen

`func (o *Asset) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### SetFirstSeen

`func (o *Asset) SetFirstSeen(v float32)`

SetFirstSeen gets a reference to the given float32 and assigns it to the FirstSeen field.

### GetLastSeen

`func (o *Asset) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Asset) GetLastSeenOk() (float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSeen

`func (o *Asset) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### SetLastSeen

`func (o *Asset) SetLastSeen(v float32)`

SetLastSeen gets a reference to the given float32 and assigns it to the LastSeen field.

### GetDetectedBy

`func (o *Asset) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *Asset) GetDetectedByOk() (string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetectedBy

`func (o *Asset) HasDetectedBy() bool`

HasDetectedBy returns a boolean if a field has been set.

### SetDetectedBy

`func (o *Asset) SetDetectedBy(v string)`

SetDetectedBy gets a reference to the given string and assigns it to the DetectedBy field.

### GetType

`func (o *Asset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Asset) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Asset) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Asset) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetOs

`func (o *Asset) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Asset) GetOsOk() (string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOs

`func (o *Asset) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOs

`func (o *Asset) SetOs(v string)`

SetOs gets a reference to the given string and assigns it to the Os field.

### GetOsVersion

`func (o *Asset) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Asset) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *Asset) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *Asset) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### GetHw

`func (o *Asset) GetHw() string`

GetHw returns the Hw field if non-nil, zero value otherwise.

### GetHwOk

`func (o *Asset) GetHwOk() (string, bool)`

GetHwOk returns a tuple with the Hw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHw

`func (o *Asset) HasHw() bool`

HasHw returns a boolean if a field has been set.

### SetHw

`func (o *Asset) SetHw(v string)`

SetHw gets a reference to the given string and assigns it to the Hw field.

### GetAddresses

`func (o *Asset) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Asset) GetAddressesOk() ([]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddresses

`func (o *Asset) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddresses

`func (o *Asset) SetAddresses(v []string)`

SetAddresses gets a reference to the given []string and assigns it to the Addresses field.

### GetAddressesExtra

`func (o *Asset) GetAddressesExtra() []string`

GetAddressesExtra returns the AddressesExtra field if non-nil, zero value otherwise.

### GetAddressesExtraOk

`func (o *Asset) GetAddressesExtraOk() ([]string, bool)`

GetAddressesExtraOk returns a tuple with the AddressesExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddressesExtra

`func (o *Asset) HasAddressesExtra() bool`

HasAddressesExtra returns a boolean if a field has been set.

### SetAddressesExtra

`func (o *Asset) SetAddressesExtra(v []string)`

SetAddressesExtra gets a reference to the given []string and assigns it to the AddressesExtra field.

### GetMacs

`func (o *Asset) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *Asset) GetMacsOk() ([]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacs

`func (o *Asset) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### SetMacs

`func (o *Asset) SetMacs(v []string)`

SetMacs gets a reference to the given []string and assigns it to the Macs field.

### GetMacVendors

`func (o *Asset) GetMacVendors() []string`

GetMacVendors returns the MacVendors field if non-nil, zero value otherwise.

### GetMacVendorsOk

`func (o *Asset) GetMacVendorsOk() ([]string, bool)`

GetMacVendorsOk returns a tuple with the MacVendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacVendors

`func (o *Asset) HasMacVendors() bool`

HasMacVendors returns a boolean if a field has been set.

### SetMacVendors

`func (o *Asset) SetMacVendors(v []string)`

SetMacVendors gets a reference to the given []string and assigns it to the MacVendors field.

### GetNames

`func (o *Asset) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Asset) GetNamesOk() ([]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *Asset) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *Asset) SetNames(v []string)`

SetNames gets a reference to the given []string and assigns it to the Names field.

### GetDomains

`func (o *Asset) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Asset) GetDomainsOk() ([]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomains

`func (o *Asset) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomains

`func (o *Asset) SetDomains(v []string)`

SetDomains gets a reference to the given []string and assigns it to the Domains field.

### GetServiceCount

`func (o *Asset) GetServiceCount() float32`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *Asset) GetServiceCountOk() (float32, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCount

`func (o *Asset) HasServiceCount() bool`

HasServiceCount returns a boolean if a field has been set.

### SetServiceCount

`func (o *Asset) SetServiceCount(v float32)`

SetServiceCount gets a reference to the given float32 and assigns it to the ServiceCount field.

### GetServiceCountTcp

`func (o *Asset) GetServiceCountTcp() float32`

GetServiceCountTcp returns the ServiceCountTcp field if non-nil, zero value otherwise.

### GetServiceCountTcpOk

`func (o *Asset) GetServiceCountTcpOk() (float32, bool)`

GetServiceCountTcpOk returns a tuple with the ServiceCountTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountTcp

`func (o *Asset) HasServiceCountTcp() bool`

HasServiceCountTcp returns a boolean if a field has been set.

### SetServiceCountTcp

`func (o *Asset) SetServiceCountTcp(v float32)`

SetServiceCountTcp gets a reference to the given float32 and assigns it to the ServiceCountTcp field.

### GetServiceCountUdp

`func (o *Asset) GetServiceCountUdp() float32`

GetServiceCountUdp returns the ServiceCountUdp field if non-nil, zero value otherwise.

### GetServiceCountUdpOk

`func (o *Asset) GetServiceCountUdpOk() (float32, bool)`

GetServiceCountUdpOk returns a tuple with the ServiceCountUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountUdp

`func (o *Asset) HasServiceCountUdp() bool`

HasServiceCountUdp returns a boolean if a field has been set.

### SetServiceCountUdp

`func (o *Asset) SetServiceCountUdp(v float32)`

SetServiceCountUdp gets a reference to the given float32 and assigns it to the ServiceCountUdp field.

### GetServiceCountArp

`func (o *Asset) GetServiceCountArp() float32`

GetServiceCountArp returns the ServiceCountArp field if non-nil, zero value otherwise.

### GetServiceCountArpOk

`func (o *Asset) GetServiceCountArpOk() (float32, bool)`

GetServiceCountArpOk returns a tuple with the ServiceCountArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountArp

`func (o *Asset) HasServiceCountArp() bool`

HasServiceCountArp returns a boolean if a field has been set.

### SetServiceCountArp

`func (o *Asset) SetServiceCountArp(v float32)`

SetServiceCountArp gets a reference to the given float32 and assigns it to the ServiceCountArp field.

### GetServiceCountIcmp

`func (o *Asset) GetServiceCountIcmp() float32`

GetServiceCountIcmp returns the ServiceCountIcmp field if non-nil, zero value otherwise.

### GetServiceCountIcmpOk

`func (o *Asset) GetServiceCountIcmpOk() (float32, bool)`

GetServiceCountIcmpOk returns a tuple with the ServiceCountIcmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountIcmp

`func (o *Asset) HasServiceCountIcmp() bool`

HasServiceCountIcmp returns a boolean if a field has been set.

### SetServiceCountIcmp

`func (o *Asset) SetServiceCountIcmp(v float32)`

SetServiceCountIcmp gets a reference to the given float32 and assigns it to the ServiceCountIcmp field.

### GetLowestTtl

`func (o *Asset) GetLowestTtl() float32`

GetLowestTtl returns the LowestTtl field if non-nil, zero value otherwise.

### GetLowestTtlOk

`func (o *Asset) GetLowestTtlOk() (float32, bool)`

GetLowestTtlOk returns a tuple with the LowestTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLowestTtl

`func (o *Asset) HasLowestTtl() bool`

HasLowestTtl returns a boolean if a field has been set.

### SetLowestTtl

`func (o *Asset) SetLowestTtl(v float32)`

SetLowestTtl gets a reference to the given float32 and assigns it to the LowestTtl field.

### GetLowestRtt

`func (o *Asset) GetLowestRtt() float32`

GetLowestRtt returns the LowestRtt field if non-nil, zero value otherwise.

### GetLowestRttOk

`func (o *Asset) GetLowestRttOk() (float32, bool)`

GetLowestRttOk returns a tuple with the LowestRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLowestRtt

`func (o *Asset) HasLowestRtt() bool`

HasLowestRtt returns a boolean if a field has been set.

### SetLowestRtt

`func (o *Asset) SetLowestRtt(v float32)`

SetLowestRtt gets a reference to the given float32 and assigns it to the LowestRtt field.

### GetLastAgentId

`func (o *Asset) GetLastAgentId() string`

GetLastAgentId returns the LastAgentId field if non-nil, zero value otherwise.

### GetLastAgentIdOk

`func (o *Asset) GetLastAgentIdOk() (string, bool)`

GetLastAgentIdOk returns a tuple with the LastAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastAgentId

`func (o *Asset) HasLastAgentId() bool`

HasLastAgentId returns a boolean if a field has been set.

### SetLastAgentId

`func (o *Asset) SetLastAgentId(v string)`

SetLastAgentId gets a reference to the given string and assigns it to the LastAgentId field.

### GetLastTaskId

`func (o *Asset) GetLastTaskId() string`

GetLastTaskId returns the LastTaskId field if non-nil, zero value otherwise.

### GetLastTaskIdOk

`func (o *Asset) GetLastTaskIdOk() (string, bool)`

GetLastTaskIdOk returns a tuple with the LastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastTaskId

`func (o *Asset) HasLastTaskId() bool`

HasLastTaskId returns a boolean if a field has been set.

### SetLastTaskId

`func (o *Asset) SetLastTaskId(v string)`

SetLastTaskId gets a reference to the given string and assigns it to the LastTaskId field.

### GetNewestMac

`func (o *Asset) GetNewestMac() string`

GetNewestMac returns the NewestMac field if non-nil, zero value otherwise.

### GetNewestMacOk

`func (o *Asset) GetNewestMacOk() (string, bool)`

GetNewestMacOk returns a tuple with the NewestMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewestMac

`func (o *Asset) HasNewestMac() bool`

HasNewestMac returns a boolean if a field has been set.

### SetNewestMac

`func (o *Asset) SetNewestMac(v string)`

SetNewestMac gets a reference to the given string and assigns it to the NewestMac field.

### GetNewestMacVendor

`func (o *Asset) GetNewestMacVendor() string`

GetNewestMacVendor returns the NewestMacVendor field if non-nil, zero value otherwise.

### GetNewestMacVendorOk

`func (o *Asset) GetNewestMacVendorOk() (string, bool)`

GetNewestMacVendorOk returns a tuple with the NewestMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewestMacVendor

`func (o *Asset) HasNewestMacVendor() bool`

HasNewestMacVendor returns a boolean if a field has been set.

### SetNewestMacVendor

`func (o *Asset) SetNewestMacVendor(v string)`

SetNewestMacVendor gets a reference to the given string and assigns it to the NewestMacVendor field.

### GetNewestMacAge

`func (o *Asset) GetNewestMacAge() float32`

GetNewestMacAge returns the NewestMacAge field if non-nil, zero value otherwise.

### GetNewestMacAgeOk

`func (o *Asset) GetNewestMacAgeOk() (float32, bool)`

GetNewestMacAgeOk returns a tuple with the NewestMacAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewestMacAge

`func (o *Asset) HasNewestMacAge() bool`

HasNewestMacAge returns a boolean if a field has been set.

### SetNewestMacAge

`func (o *Asset) SetNewestMacAge(v float32)`

SetNewestMacAge gets a reference to the given float32 and assigns it to the NewestMacAge field.

### GetComments

`func (o *Asset) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Asset) GetCommentsOk() (string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComments

`func (o *Asset) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetComments

`func (o *Asset) SetComments(v string)`

SetComments gets a reference to the given string and assigns it to the Comments field.

### GetServicePortsTcp

`func (o *Asset) GetServicePortsTcp() []string`

GetServicePortsTcp returns the ServicePortsTcp field if non-nil, zero value otherwise.

### GetServicePortsTcpOk

`func (o *Asset) GetServicePortsTcpOk() ([]string, bool)`

GetServicePortsTcpOk returns a tuple with the ServicePortsTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsTcp

`func (o *Asset) HasServicePortsTcp() bool`

HasServicePortsTcp returns a boolean if a field has been set.

### SetServicePortsTcp

`func (o *Asset) SetServicePortsTcp(v []string)`

SetServicePortsTcp gets a reference to the given []string and assigns it to the ServicePortsTcp field.

### GetServicePortsUdp

`func (o *Asset) GetServicePortsUdp() []string`

GetServicePortsUdp returns the ServicePortsUdp field if non-nil, zero value otherwise.

### GetServicePortsUdpOk

`func (o *Asset) GetServicePortsUdpOk() ([]string, bool)`

GetServicePortsUdpOk returns a tuple with the ServicePortsUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsUdp

`func (o *Asset) HasServicePortsUdp() bool`

HasServicePortsUdp returns a boolean if a field has been set.

### SetServicePortsUdp

`func (o *Asset) SetServicePortsUdp(v []string)`

SetServicePortsUdp gets a reference to the given []string and assigns it to the ServicePortsUdp field.

### GetServicePortsProtocols

`func (o *Asset) GetServicePortsProtocols() []string`

GetServicePortsProtocols returns the ServicePortsProtocols field if non-nil, zero value otherwise.

### GetServicePortsProtocolsOk

`func (o *Asset) GetServicePortsProtocolsOk() ([]string, bool)`

GetServicePortsProtocolsOk returns a tuple with the ServicePortsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsProtocols

`func (o *Asset) HasServicePortsProtocols() bool`

HasServicePortsProtocols returns a boolean if a field has been set.

### SetServicePortsProtocols

`func (o *Asset) SetServicePortsProtocols(v []string)`

SetServicePortsProtocols gets a reference to the given []string and assigns it to the ServicePortsProtocols field.

### GetServicePortsProducts

`func (o *Asset) GetServicePortsProducts() []string`

GetServicePortsProducts returns the ServicePortsProducts field if non-nil, zero value otherwise.

### GetServicePortsProductsOk

`func (o *Asset) GetServicePortsProductsOk() ([]string, bool)`

GetServicePortsProductsOk returns a tuple with the ServicePortsProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsProducts

`func (o *Asset) HasServicePortsProducts() bool`

HasServicePortsProducts returns a boolean if a field has been set.

### SetServicePortsProducts

`func (o *Asset) SetServicePortsProducts(v []string)`

SetServicePortsProducts gets a reference to the given []string and assigns it to the ServicePortsProducts field.

### GetOrgName

`func (o *Asset) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Asset) GetOrgNameOk() (string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrgName

`func (o *Asset) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgName

`func (o *Asset) SetOrgName(v string)`

SetOrgName gets a reference to the given string and assigns it to the OrgName field.

### GetSiteName

`func (o *Asset) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *Asset) GetSiteNameOk() (string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteName

`func (o *Asset) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteName

`func (o *Asset) SetSiteName(v string)`

SetSiteName gets a reference to the given string and assigns it to the SiteName field.

### GetAgentName

`func (o *Asset) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *Asset) GetAgentNameOk() (string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgentName

`func (o *Asset) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### SetAgentName

`func (o *Asset) SetAgentName(v string)`

SetAgentName gets a reference to the given string and assigns it to the AgentName field.

### GetTags

`func (o *Asset) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Asset) GetTagsOk() (map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Asset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Asset) SetTags(v map[string]string)`

SetTags gets a reference to the given map[string]string and assigns it to the Tags field.

### GetServices

`func (o *Asset) GetServices() map[string]map[string]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Asset) GetServicesOk() (map[string]map[string]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServices

`func (o *Asset) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServices

`func (o *Asset) SetServices(v map[string]map[string]string)`

SetServices gets a reference to the given map[string]map[string]string and assigns it to the Services field.

### GetRtts

`func (o *Asset) GetRtts() map[string]map[string]interface{}`

GetRtts returns the Rtts field if non-nil, zero value otherwise.

### GetRttsOk

`func (o *Asset) GetRttsOk() (map[string]map[string]interface{}, bool)`

GetRttsOk returns a tuple with the Rtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRtts

`func (o *Asset) HasRtts() bool`

HasRtts returns a boolean if a field has been set.

### SetRtts

`func (o *Asset) SetRtts(v map[string]map[string]interface{})`

SetRtts gets a reference to the given map[string]map[string]interface{} and assigns it to the Rtts field.

### GetCredentials

`func (o *Asset) GetCredentials() map[string]map[string]bool`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Asset) GetCredentialsOk() (map[string]map[string]bool, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCredentials

`func (o *Asset) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentials

`func (o *Asset) SetCredentials(v map[string]map[string]bool)`

SetCredentials gets a reference to the given map[string]map[string]bool and assigns it to the Credentials field.

### GetAttributes

`func (o *Asset) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Asset) GetAttributesOk() (map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Asset) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Asset) SetAttributes(v map[string]string)`

SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


