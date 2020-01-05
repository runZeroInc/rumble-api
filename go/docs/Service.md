# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | 
**ServiceAssetId** | Pointer to **string** |  | [optional] 
**ServiceCreatedAt** | Pointer to **float32** |  | [optional] 
**ServiceUpdatedAt** | Pointer to **float32** |  | [optional] 
**ServiceAddress** | Pointer to **string** |  | [optional] 
**ServiceTransport** | Pointer to **string** |  | [optional] 
**ServiceVhost** | Pointer to **string** |  | [optional] 
**ServicePort** | Pointer to **float32** |  | [optional] 
**ServiceData** | Pointer to **map[string]string** |  | [optional] 
**ServiceProtocol** | Pointer to **string** |  | [optional] 
**ServiceSummary** | Pointer to **string** |  | [optional] 
**ServiceScreenshotLink** | Pointer to **string** |  | [optional] 
**ServiceLink** | Pointer to **string** |  | [optional] 
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

### GetServiceId

`func (o *Service) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Service) GetServiceIdOk() (string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceId

`func (o *Service) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceId

`func (o *Service) SetServiceId(v string)`

SetServiceId gets a reference to the given string and assigns it to the ServiceId field.

### GetServiceAssetId

`func (o *Service) GetServiceAssetId() string`

GetServiceAssetId returns the ServiceAssetId field if non-nil, zero value otherwise.

### GetServiceAssetIdOk

`func (o *Service) GetServiceAssetIdOk() (string, bool)`

GetServiceAssetIdOk returns a tuple with the ServiceAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceAssetId

`func (o *Service) HasServiceAssetId() bool`

HasServiceAssetId returns a boolean if a field has been set.

### SetServiceAssetId

`func (o *Service) SetServiceAssetId(v string)`

SetServiceAssetId gets a reference to the given string and assigns it to the ServiceAssetId field.

### GetServiceCreatedAt

`func (o *Service) GetServiceCreatedAt() float32`

GetServiceCreatedAt returns the ServiceCreatedAt field if non-nil, zero value otherwise.

### GetServiceCreatedAtOk

`func (o *Service) GetServiceCreatedAtOk() (float32, bool)`

GetServiceCreatedAtOk returns a tuple with the ServiceCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCreatedAt

`func (o *Service) HasServiceCreatedAt() bool`

HasServiceCreatedAt returns a boolean if a field has been set.

### SetServiceCreatedAt

`func (o *Service) SetServiceCreatedAt(v float32)`

SetServiceCreatedAt gets a reference to the given float32 and assigns it to the ServiceCreatedAt field.

### GetServiceUpdatedAt

`func (o *Service) GetServiceUpdatedAt() float32`

GetServiceUpdatedAt returns the ServiceUpdatedAt field if non-nil, zero value otherwise.

### GetServiceUpdatedAtOk

`func (o *Service) GetServiceUpdatedAtOk() (float32, bool)`

GetServiceUpdatedAtOk returns a tuple with the ServiceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceUpdatedAt

`func (o *Service) HasServiceUpdatedAt() bool`

HasServiceUpdatedAt returns a boolean if a field has been set.

### SetServiceUpdatedAt

`func (o *Service) SetServiceUpdatedAt(v float32)`

SetServiceUpdatedAt gets a reference to the given float32 and assigns it to the ServiceUpdatedAt field.

### GetServiceAddress

`func (o *Service) GetServiceAddress() string`

GetServiceAddress returns the ServiceAddress field if non-nil, zero value otherwise.

### GetServiceAddressOk

`func (o *Service) GetServiceAddressOk() (string, bool)`

GetServiceAddressOk returns a tuple with the ServiceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceAddress

`func (o *Service) HasServiceAddress() bool`

HasServiceAddress returns a boolean if a field has been set.

### SetServiceAddress

`func (o *Service) SetServiceAddress(v string)`

SetServiceAddress gets a reference to the given string and assigns it to the ServiceAddress field.

### GetServiceTransport

`func (o *Service) GetServiceTransport() string`

GetServiceTransport returns the ServiceTransport field if non-nil, zero value otherwise.

### GetServiceTransportOk

`func (o *Service) GetServiceTransportOk() (string, bool)`

GetServiceTransportOk returns a tuple with the ServiceTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceTransport

`func (o *Service) HasServiceTransport() bool`

HasServiceTransport returns a boolean if a field has been set.

### SetServiceTransport

`func (o *Service) SetServiceTransport(v string)`

SetServiceTransport gets a reference to the given string and assigns it to the ServiceTransport field.

### GetServiceVhost

`func (o *Service) GetServiceVhost() string`

GetServiceVhost returns the ServiceVhost field if non-nil, zero value otherwise.

### GetServiceVhostOk

`func (o *Service) GetServiceVhostOk() (string, bool)`

GetServiceVhostOk returns a tuple with the ServiceVhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceVhost

`func (o *Service) HasServiceVhost() bool`

HasServiceVhost returns a boolean if a field has been set.

### SetServiceVhost

`func (o *Service) SetServiceVhost(v string)`

SetServiceVhost gets a reference to the given string and assigns it to the ServiceVhost field.

### GetServicePort

`func (o *Service) GetServicePort() float32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *Service) GetServicePortOk() (float32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePort

`func (o *Service) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePort

`func (o *Service) SetServicePort(v float32)`

SetServicePort gets a reference to the given float32 and assigns it to the ServicePort field.

### GetServiceData

`func (o *Service) GetServiceData() map[string]string`

GetServiceData returns the ServiceData field if non-nil, zero value otherwise.

### GetServiceDataOk

`func (o *Service) GetServiceDataOk() (map[string]string, bool)`

GetServiceDataOk returns a tuple with the ServiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceData

`func (o *Service) HasServiceData() bool`

HasServiceData returns a boolean if a field has been set.

### SetServiceData

`func (o *Service) SetServiceData(v map[string]string)`

SetServiceData gets a reference to the given map[string]string and assigns it to the ServiceData field.

### GetServiceProtocol

`func (o *Service) GetServiceProtocol() string`

GetServiceProtocol returns the ServiceProtocol field if non-nil, zero value otherwise.

### GetServiceProtocolOk

`func (o *Service) GetServiceProtocolOk() (string, bool)`

GetServiceProtocolOk returns a tuple with the ServiceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceProtocol

`func (o *Service) HasServiceProtocol() bool`

HasServiceProtocol returns a boolean if a field has been set.

### SetServiceProtocol

`func (o *Service) SetServiceProtocol(v string)`

SetServiceProtocol gets a reference to the given string and assigns it to the ServiceProtocol field.

### GetServiceSummary

`func (o *Service) GetServiceSummary() string`

GetServiceSummary returns the ServiceSummary field if non-nil, zero value otherwise.

### GetServiceSummaryOk

`func (o *Service) GetServiceSummaryOk() (string, bool)`

GetServiceSummaryOk returns a tuple with the ServiceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceSummary

`func (o *Service) HasServiceSummary() bool`

HasServiceSummary returns a boolean if a field has been set.

### SetServiceSummary

`func (o *Service) SetServiceSummary(v string)`

SetServiceSummary gets a reference to the given string and assigns it to the ServiceSummary field.

### GetServiceScreenshotLink

`func (o *Service) GetServiceScreenshotLink() string`

GetServiceScreenshotLink returns the ServiceScreenshotLink field if non-nil, zero value otherwise.

### GetServiceScreenshotLinkOk

`func (o *Service) GetServiceScreenshotLinkOk() (string, bool)`

GetServiceScreenshotLinkOk returns a tuple with the ServiceScreenshotLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceScreenshotLink

`func (o *Service) HasServiceScreenshotLink() bool`

HasServiceScreenshotLink returns a boolean if a field has been set.

### SetServiceScreenshotLink

`func (o *Service) SetServiceScreenshotLink(v string)`

SetServiceScreenshotLink gets a reference to the given string and assigns it to the ServiceScreenshotLink field.

### GetServiceLink

`func (o *Service) GetServiceLink() string`

GetServiceLink returns the ServiceLink field if non-nil, zero value otherwise.

### GetServiceLinkOk

`func (o *Service) GetServiceLinkOk() (string, bool)`

GetServiceLinkOk returns a tuple with the ServiceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceLink

`func (o *Service) HasServiceLink() bool`

HasServiceLink returns a boolean if a field has been set.

### SetServiceLink

`func (o *Service) SetServiceLink(v string)`

SetServiceLink gets a reference to the given string and assigns it to the ServiceLink field.

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedAt

`func (o *Service) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *Service) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v float32)`

SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.

### GetUpdatedAt

`func (o *Service) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedAt

`func (o *Service) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v float32)`

SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.

### GetOrganizationId

`func (o *Service) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Service) GetOrganizationIdOk() (string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationId

`func (o *Service) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationId

`func (o *Service) SetOrganizationId(v string)`

SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.

### GetSiteId

`func (o *Service) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Service) GetSiteIdOk() (string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteId

`func (o *Service) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteId

`func (o *Service) SetSiteId(v string)`

SetSiteId gets a reference to the given string and assigns it to the SiteId field.

### GetAlive

`func (o *Service) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *Service) GetAliveOk() (bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlive

`func (o *Service) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### SetAlive

`func (o *Service) SetAlive(v bool)`

SetAlive gets a reference to the given bool and assigns it to the Alive field.

### GetFirstSeen

`func (o *Service) GetFirstSeen() float32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *Service) GetFirstSeenOk() (float32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstSeen

`func (o *Service) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### SetFirstSeen

`func (o *Service) SetFirstSeen(v float32)`

SetFirstSeen gets a reference to the given float32 and assigns it to the FirstSeen field.

### GetLastSeen

`func (o *Service) GetLastSeen() float32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Service) GetLastSeenOk() (float32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSeen

`func (o *Service) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### SetLastSeen

`func (o *Service) SetLastSeen(v float32)`

SetLastSeen gets a reference to the given float32 and assigns it to the LastSeen field.

### GetDetectedBy

`func (o *Service) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *Service) GetDetectedByOk() (string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetectedBy

`func (o *Service) HasDetectedBy() bool`

HasDetectedBy returns a boolean if a field has been set.

### SetDetectedBy

`func (o *Service) SetDetectedBy(v string)`

SetDetectedBy gets a reference to the given string and assigns it to the DetectedBy field.

### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetOs

`func (o *Service) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Service) GetOsOk() (string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOs

`func (o *Service) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOs

`func (o *Service) SetOs(v string)`

SetOs gets a reference to the given string and assigns it to the Os field.

### GetOsVersion

`func (o *Service) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Service) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *Service) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *Service) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### GetHw

`func (o *Service) GetHw() string`

GetHw returns the Hw field if non-nil, zero value otherwise.

### GetHwOk

`func (o *Service) GetHwOk() (string, bool)`

GetHwOk returns a tuple with the Hw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHw

`func (o *Service) HasHw() bool`

HasHw returns a boolean if a field has been set.

### SetHw

`func (o *Service) SetHw(v string)`

SetHw gets a reference to the given string and assigns it to the Hw field.

### GetAddresses

`func (o *Service) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Service) GetAddressesOk() ([]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddresses

`func (o *Service) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddresses

`func (o *Service) SetAddresses(v []string)`

SetAddresses gets a reference to the given []string and assigns it to the Addresses field.

### GetAddressesExtra

`func (o *Service) GetAddressesExtra() []string`

GetAddressesExtra returns the AddressesExtra field if non-nil, zero value otherwise.

### GetAddressesExtraOk

`func (o *Service) GetAddressesExtraOk() ([]string, bool)`

GetAddressesExtraOk returns a tuple with the AddressesExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddressesExtra

`func (o *Service) HasAddressesExtra() bool`

HasAddressesExtra returns a boolean if a field has been set.

### SetAddressesExtra

`func (o *Service) SetAddressesExtra(v []string)`

SetAddressesExtra gets a reference to the given []string and assigns it to the AddressesExtra field.

### GetMacs

`func (o *Service) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *Service) GetMacsOk() ([]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacs

`func (o *Service) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### SetMacs

`func (o *Service) SetMacs(v []string)`

SetMacs gets a reference to the given []string and assigns it to the Macs field.

### GetMacVendors

`func (o *Service) GetMacVendors() []string`

GetMacVendors returns the MacVendors field if non-nil, zero value otherwise.

### GetMacVendorsOk

`func (o *Service) GetMacVendorsOk() ([]string, bool)`

GetMacVendorsOk returns a tuple with the MacVendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacVendors

`func (o *Service) HasMacVendors() bool`

HasMacVendors returns a boolean if a field has been set.

### SetMacVendors

`func (o *Service) SetMacVendors(v []string)`

SetMacVendors gets a reference to the given []string and assigns it to the MacVendors field.

### GetNames

`func (o *Service) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Service) GetNamesOk() ([]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *Service) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *Service) SetNames(v []string)`

SetNames gets a reference to the given []string and assigns it to the Names field.

### GetDomains

`func (o *Service) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Service) GetDomainsOk() ([]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomains

`func (o *Service) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomains

`func (o *Service) SetDomains(v []string)`

SetDomains gets a reference to the given []string and assigns it to the Domains field.

### GetServiceCount

`func (o *Service) GetServiceCount() float32`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *Service) GetServiceCountOk() (float32, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCount

`func (o *Service) HasServiceCount() bool`

HasServiceCount returns a boolean if a field has been set.

### SetServiceCount

`func (o *Service) SetServiceCount(v float32)`

SetServiceCount gets a reference to the given float32 and assigns it to the ServiceCount field.

### GetServiceCountTcp

`func (o *Service) GetServiceCountTcp() float32`

GetServiceCountTcp returns the ServiceCountTcp field if non-nil, zero value otherwise.

### GetServiceCountTcpOk

`func (o *Service) GetServiceCountTcpOk() (float32, bool)`

GetServiceCountTcpOk returns a tuple with the ServiceCountTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountTcp

`func (o *Service) HasServiceCountTcp() bool`

HasServiceCountTcp returns a boolean if a field has been set.

### SetServiceCountTcp

`func (o *Service) SetServiceCountTcp(v float32)`

SetServiceCountTcp gets a reference to the given float32 and assigns it to the ServiceCountTcp field.

### GetServiceCountUdp

`func (o *Service) GetServiceCountUdp() float32`

GetServiceCountUdp returns the ServiceCountUdp field if non-nil, zero value otherwise.

### GetServiceCountUdpOk

`func (o *Service) GetServiceCountUdpOk() (float32, bool)`

GetServiceCountUdpOk returns a tuple with the ServiceCountUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountUdp

`func (o *Service) HasServiceCountUdp() bool`

HasServiceCountUdp returns a boolean if a field has been set.

### SetServiceCountUdp

`func (o *Service) SetServiceCountUdp(v float32)`

SetServiceCountUdp gets a reference to the given float32 and assigns it to the ServiceCountUdp field.

### GetServiceCountArp

`func (o *Service) GetServiceCountArp() float32`

GetServiceCountArp returns the ServiceCountArp field if non-nil, zero value otherwise.

### GetServiceCountArpOk

`func (o *Service) GetServiceCountArpOk() (float32, bool)`

GetServiceCountArpOk returns a tuple with the ServiceCountArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountArp

`func (o *Service) HasServiceCountArp() bool`

HasServiceCountArp returns a boolean if a field has been set.

### SetServiceCountArp

`func (o *Service) SetServiceCountArp(v float32)`

SetServiceCountArp gets a reference to the given float32 and assigns it to the ServiceCountArp field.

### GetServiceCountIcmp

`func (o *Service) GetServiceCountIcmp() float32`

GetServiceCountIcmp returns the ServiceCountIcmp field if non-nil, zero value otherwise.

### GetServiceCountIcmpOk

`func (o *Service) GetServiceCountIcmpOk() (float32, bool)`

GetServiceCountIcmpOk returns a tuple with the ServiceCountIcmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceCountIcmp

`func (o *Service) HasServiceCountIcmp() bool`

HasServiceCountIcmp returns a boolean if a field has been set.

### SetServiceCountIcmp

`func (o *Service) SetServiceCountIcmp(v float32)`

SetServiceCountIcmp gets a reference to the given float32 and assigns it to the ServiceCountIcmp field.

### GetLowestTtl

`func (o *Service) GetLowestTtl() float32`

GetLowestTtl returns the LowestTtl field if non-nil, zero value otherwise.

### GetLowestTtlOk

`func (o *Service) GetLowestTtlOk() (float32, bool)`

GetLowestTtlOk returns a tuple with the LowestTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLowestTtl

`func (o *Service) HasLowestTtl() bool`

HasLowestTtl returns a boolean if a field has been set.

### SetLowestTtl

`func (o *Service) SetLowestTtl(v float32)`

SetLowestTtl gets a reference to the given float32 and assigns it to the LowestTtl field.

### GetLowestRtt

`func (o *Service) GetLowestRtt() float32`

GetLowestRtt returns the LowestRtt field if non-nil, zero value otherwise.

### GetLowestRttOk

`func (o *Service) GetLowestRttOk() (float32, bool)`

GetLowestRttOk returns a tuple with the LowestRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLowestRtt

`func (o *Service) HasLowestRtt() bool`

HasLowestRtt returns a boolean if a field has been set.

### SetLowestRtt

`func (o *Service) SetLowestRtt(v float32)`

SetLowestRtt gets a reference to the given float32 and assigns it to the LowestRtt field.

### GetLastAgentId

`func (o *Service) GetLastAgentId() string`

GetLastAgentId returns the LastAgentId field if non-nil, zero value otherwise.

### GetLastAgentIdOk

`func (o *Service) GetLastAgentIdOk() (string, bool)`

GetLastAgentIdOk returns a tuple with the LastAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastAgentId

`func (o *Service) HasLastAgentId() bool`

HasLastAgentId returns a boolean if a field has been set.

### SetLastAgentId

`func (o *Service) SetLastAgentId(v string)`

SetLastAgentId gets a reference to the given string and assigns it to the LastAgentId field.

### GetLastTaskId

`func (o *Service) GetLastTaskId() string`

GetLastTaskId returns the LastTaskId field if non-nil, zero value otherwise.

### GetLastTaskIdOk

`func (o *Service) GetLastTaskIdOk() (string, bool)`

GetLastTaskIdOk returns a tuple with the LastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastTaskId

`func (o *Service) HasLastTaskId() bool`

HasLastTaskId returns a boolean if a field has been set.

### SetLastTaskId

`func (o *Service) SetLastTaskId(v string)`

SetLastTaskId gets a reference to the given string and assigns it to the LastTaskId field.

### GetNewestMac

`func (o *Service) GetNewestMac() string`

GetNewestMac returns the NewestMac field if non-nil, zero value otherwise.

### GetNewestMacOk

`func (o *Service) GetNewestMacOk() (string, bool)`

GetNewestMacOk returns a tuple with the NewestMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewestMac

`func (o *Service) HasNewestMac() bool`

HasNewestMac returns a boolean if a field has been set.

### SetNewestMac

`func (o *Service) SetNewestMac(v string)`

SetNewestMac gets a reference to the given string and assigns it to the NewestMac field.

### GetNewestMacVendor

`func (o *Service) GetNewestMacVendor() string`

GetNewestMacVendor returns the NewestMacVendor field if non-nil, zero value otherwise.

### GetNewestMacVendorOk

`func (o *Service) GetNewestMacVendorOk() (string, bool)`

GetNewestMacVendorOk returns a tuple with the NewestMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewestMacVendor

`func (o *Service) HasNewestMacVendor() bool`

HasNewestMacVendor returns a boolean if a field has been set.

### SetNewestMacVendor

`func (o *Service) SetNewestMacVendor(v string)`

SetNewestMacVendor gets a reference to the given string and assigns it to the NewestMacVendor field.

### GetNewestMacAge

`func (o *Service) GetNewestMacAge() float32`

GetNewestMacAge returns the NewestMacAge field if non-nil, zero value otherwise.

### GetNewestMacAgeOk

`func (o *Service) GetNewestMacAgeOk() (float32, bool)`

GetNewestMacAgeOk returns a tuple with the NewestMacAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewestMacAge

`func (o *Service) HasNewestMacAge() bool`

HasNewestMacAge returns a boolean if a field has been set.

### SetNewestMacAge

`func (o *Service) SetNewestMacAge(v float32)`

SetNewestMacAge gets a reference to the given float32 and assigns it to the NewestMacAge field.

### GetComments

`func (o *Service) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Service) GetCommentsOk() (string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComments

`func (o *Service) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetComments

`func (o *Service) SetComments(v string)`

SetComments gets a reference to the given string and assigns it to the Comments field.

### GetServicePortsTcp

`func (o *Service) GetServicePortsTcp() []string`

GetServicePortsTcp returns the ServicePortsTcp field if non-nil, zero value otherwise.

### GetServicePortsTcpOk

`func (o *Service) GetServicePortsTcpOk() ([]string, bool)`

GetServicePortsTcpOk returns a tuple with the ServicePortsTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsTcp

`func (o *Service) HasServicePortsTcp() bool`

HasServicePortsTcp returns a boolean if a field has been set.

### SetServicePortsTcp

`func (o *Service) SetServicePortsTcp(v []string)`

SetServicePortsTcp gets a reference to the given []string and assigns it to the ServicePortsTcp field.

### GetServicePortsUdp

`func (o *Service) GetServicePortsUdp() []string`

GetServicePortsUdp returns the ServicePortsUdp field if non-nil, zero value otherwise.

### GetServicePortsUdpOk

`func (o *Service) GetServicePortsUdpOk() ([]string, bool)`

GetServicePortsUdpOk returns a tuple with the ServicePortsUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsUdp

`func (o *Service) HasServicePortsUdp() bool`

HasServicePortsUdp returns a boolean if a field has been set.

### SetServicePortsUdp

`func (o *Service) SetServicePortsUdp(v []string)`

SetServicePortsUdp gets a reference to the given []string and assigns it to the ServicePortsUdp field.

### GetServicePortsProtocols

`func (o *Service) GetServicePortsProtocols() []string`

GetServicePortsProtocols returns the ServicePortsProtocols field if non-nil, zero value otherwise.

### GetServicePortsProtocolsOk

`func (o *Service) GetServicePortsProtocolsOk() ([]string, bool)`

GetServicePortsProtocolsOk returns a tuple with the ServicePortsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsProtocols

`func (o *Service) HasServicePortsProtocols() bool`

HasServicePortsProtocols returns a boolean if a field has been set.

### SetServicePortsProtocols

`func (o *Service) SetServicePortsProtocols(v []string)`

SetServicePortsProtocols gets a reference to the given []string and assigns it to the ServicePortsProtocols field.

### GetServicePortsProducts

`func (o *Service) GetServicePortsProducts() []string`

GetServicePortsProducts returns the ServicePortsProducts field if non-nil, zero value otherwise.

### GetServicePortsProductsOk

`func (o *Service) GetServicePortsProductsOk() ([]string, bool)`

GetServicePortsProductsOk returns a tuple with the ServicePortsProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServicePortsProducts

`func (o *Service) HasServicePortsProducts() bool`

HasServicePortsProducts returns a boolean if a field has been set.

### SetServicePortsProducts

`func (o *Service) SetServicePortsProducts(v []string)`

SetServicePortsProducts gets a reference to the given []string and assigns it to the ServicePortsProducts field.

### GetOrgName

`func (o *Service) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Service) GetOrgNameOk() (string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrgName

`func (o *Service) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### SetOrgName

`func (o *Service) SetOrgName(v string)`

SetOrgName gets a reference to the given string and assigns it to the OrgName field.

### GetSiteName

`func (o *Service) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *Service) GetSiteNameOk() (string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteName

`func (o *Service) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteName

`func (o *Service) SetSiteName(v string)`

SetSiteName gets a reference to the given string and assigns it to the SiteName field.

### GetAgentName

`func (o *Service) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *Service) GetAgentNameOk() (string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgentName

`func (o *Service) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### SetAgentName

`func (o *Service) SetAgentName(v string)`

SetAgentName gets a reference to the given string and assigns it to the AgentName field.

### GetTags

`func (o *Service) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Service) SetTags(v map[string]string)`

SetTags gets a reference to the given map[string]string and assigns it to the Tags field.

### GetServices

`func (o *Service) GetServices() map[string]map[string]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Service) GetServicesOk() (map[string]map[string]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServices

`func (o *Service) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServices

`func (o *Service) SetServices(v map[string]map[string]string)`

SetServices gets a reference to the given map[string]map[string]string and assigns it to the Services field.

### GetRtts

`func (o *Service) GetRtts() map[string]map[string]interface{}`

GetRtts returns the Rtts field if non-nil, zero value otherwise.

### GetRttsOk

`func (o *Service) GetRttsOk() (map[string]map[string]interface{}, bool)`

GetRttsOk returns a tuple with the Rtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRtts

`func (o *Service) HasRtts() bool`

HasRtts returns a boolean if a field has been set.

### SetRtts

`func (o *Service) SetRtts(v map[string]map[string]interface{})`

SetRtts gets a reference to the given map[string]map[string]interface{} and assigns it to the Rtts field.

### GetCredentials

`func (o *Service) GetCredentials() map[string]map[string]bool`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Service) GetCredentialsOk() (map[string]map[string]bool, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCredentials

`func (o *Service) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentials

`func (o *Service) SetCredentials(v map[string]map[string]bool)`

SetCredentials gets a reference to the given map[string]map[string]bool and assigns it to the Credentials field.

### GetAttributes

`func (o *Service) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Service) GetAttributesOk() (map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Service) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Service) SetAttributes(v map[string]string)`

SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


