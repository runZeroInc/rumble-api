# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | 
**ServiceAssetId** | Pointer to **string** |  | [optional] 
**ServiceCreatedAt** | Pointer to **int64** |  | [optional] 
**ServiceUpdatedAt** | Pointer to **int64** |  | [optional] 
**ServiceAddress** | Pointer to **string** |  | [optional] 
**ServiceTransport** | Pointer to **string** |  | [optional] 
**ServiceVhost** | Pointer to **string** |  | [optional] 
**ServicePort** | Pointer to **string** |  | [optional] 
**ServiceData** | Pointer to **map[string]string** |  | [optional] 
**ServiceProtocol** | Pointer to **string** |  | [optional] 
**ServiceSummary** | Pointer to **string** |  | [optional] 
**ServiceScreenshotLink** | Pointer to **string** |  | [optional] 
**ServiceLink** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Alive** | Pointer to **bool** |  | [optional] 
**FirstSeen** | Pointer to **int64** |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
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
**ServiceCount** | Pointer to **int64** |  | [optional] 
**ServiceCountTcp** | Pointer to **int64** |  | [optional] 
**ServiceCountUdp** | Pointer to **int64** |  | [optional] 
**ServiceCountArp** | Pointer to **int64** |  | [optional] 
**ServiceCountIcmp** | Pointer to **int64** |  | [optional] 
**LowestTtl** | Pointer to **int64** |  | [optional] 
**LowestRtt** | Pointer to **int64** |  | [optional] 
**LastAgentId** | Pointer to **string** |  | [optional] 
**LastTaskId** | Pointer to **string** |  | [optional] 
**NewestMac** | Pointer to **string** |  | [optional] 
**NewestMacVendor** | Pointer to **string** |  | [optional] 
**NewestMacAge** | Pointer to **int64** |  | [optional] 
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
**Rtts** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Credentials** | Pointer to [**map[string]map[string]bool**](map.md) |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewService

`func NewService(serviceId string, id string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *Service) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Service) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Service) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceAssetId

`func (o *Service) GetServiceAssetId() string`

GetServiceAssetId returns the ServiceAssetId field if non-nil, zero value otherwise.

### GetServiceAssetIdOk

`func (o *Service) GetServiceAssetIdOk() (*string, bool)`

GetServiceAssetIdOk returns a tuple with the ServiceAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAssetId

`func (o *Service) SetServiceAssetId(v string)`

SetServiceAssetId sets ServiceAssetId field to given value.

### HasServiceAssetId

`func (o *Service) HasServiceAssetId() bool`

HasServiceAssetId returns a boolean if a field has been set.

### GetServiceCreatedAt

`func (o *Service) GetServiceCreatedAt() int64`

GetServiceCreatedAt returns the ServiceCreatedAt field if non-nil, zero value otherwise.

### GetServiceCreatedAtOk

`func (o *Service) GetServiceCreatedAtOk() (*int64, bool)`

GetServiceCreatedAtOk returns a tuple with the ServiceCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCreatedAt

`func (o *Service) SetServiceCreatedAt(v int64)`

SetServiceCreatedAt sets ServiceCreatedAt field to given value.

### HasServiceCreatedAt

`func (o *Service) HasServiceCreatedAt() bool`

HasServiceCreatedAt returns a boolean if a field has been set.

### GetServiceUpdatedAt

`func (o *Service) GetServiceUpdatedAt() int64`

GetServiceUpdatedAt returns the ServiceUpdatedAt field if non-nil, zero value otherwise.

### GetServiceUpdatedAtOk

`func (o *Service) GetServiceUpdatedAtOk() (*int64, bool)`

GetServiceUpdatedAtOk returns a tuple with the ServiceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUpdatedAt

`func (o *Service) SetServiceUpdatedAt(v int64)`

SetServiceUpdatedAt sets ServiceUpdatedAt field to given value.

### HasServiceUpdatedAt

`func (o *Service) HasServiceUpdatedAt() bool`

HasServiceUpdatedAt returns a boolean if a field has been set.

### GetServiceAddress

`func (o *Service) GetServiceAddress() string`

GetServiceAddress returns the ServiceAddress field if non-nil, zero value otherwise.

### GetServiceAddressOk

`func (o *Service) GetServiceAddressOk() (*string, bool)`

GetServiceAddressOk returns a tuple with the ServiceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAddress

`func (o *Service) SetServiceAddress(v string)`

SetServiceAddress sets ServiceAddress field to given value.

### HasServiceAddress

`func (o *Service) HasServiceAddress() bool`

HasServiceAddress returns a boolean if a field has been set.

### GetServiceTransport

`func (o *Service) GetServiceTransport() string`

GetServiceTransport returns the ServiceTransport field if non-nil, zero value otherwise.

### GetServiceTransportOk

`func (o *Service) GetServiceTransportOk() (*string, bool)`

GetServiceTransportOk returns a tuple with the ServiceTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTransport

`func (o *Service) SetServiceTransport(v string)`

SetServiceTransport sets ServiceTransport field to given value.

### HasServiceTransport

`func (o *Service) HasServiceTransport() bool`

HasServiceTransport returns a boolean if a field has been set.

### GetServiceVhost

`func (o *Service) GetServiceVhost() string`

GetServiceVhost returns the ServiceVhost field if non-nil, zero value otherwise.

### GetServiceVhostOk

`func (o *Service) GetServiceVhostOk() (*string, bool)`

GetServiceVhostOk returns a tuple with the ServiceVhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVhost

`func (o *Service) SetServiceVhost(v string)`

SetServiceVhost sets ServiceVhost field to given value.

### HasServiceVhost

`func (o *Service) HasServiceVhost() bool`

HasServiceVhost returns a boolean if a field has been set.

### GetServicePort

`func (o *Service) GetServicePort() string`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *Service) GetServicePortOk() (*string, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *Service) SetServicePort(v string)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *Service) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceData

`func (o *Service) GetServiceData() map[string]string`

GetServiceData returns the ServiceData field if non-nil, zero value otherwise.

### GetServiceDataOk

`func (o *Service) GetServiceDataOk() (*map[string]string, bool)`

GetServiceDataOk returns a tuple with the ServiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceData

`func (o *Service) SetServiceData(v map[string]string)`

SetServiceData sets ServiceData field to given value.

### HasServiceData

`func (o *Service) HasServiceData() bool`

HasServiceData returns a boolean if a field has been set.

### GetServiceProtocol

`func (o *Service) GetServiceProtocol() string`

GetServiceProtocol returns the ServiceProtocol field if non-nil, zero value otherwise.

### GetServiceProtocolOk

`func (o *Service) GetServiceProtocolOk() (*string, bool)`

GetServiceProtocolOk returns a tuple with the ServiceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProtocol

`func (o *Service) SetServiceProtocol(v string)`

SetServiceProtocol sets ServiceProtocol field to given value.

### HasServiceProtocol

`func (o *Service) HasServiceProtocol() bool`

HasServiceProtocol returns a boolean if a field has been set.

### GetServiceSummary

`func (o *Service) GetServiceSummary() string`

GetServiceSummary returns the ServiceSummary field if non-nil, zero value otherwise.

### GetServiceSummaryOk

`func (o *Service) GetServiceSummaryOk() (*string, bool)`

GetServiceSummaryOk returns a tuple with the ServiceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSummary

`func (o *Service) SetServiceSummary(v string)`

SetServiceSummary sets ServiceSummary field to given value.

### HasServiceSummary

`func (o *Service) HasServiceSummary() bool`

HasServiceSummary returns a boolean if a field has been set.

### GetServiceScreenshotLink

`func (o *Service) GetServiceScreenshotLink() string`

GetServiceScreenshotLink returns the ServiceScreenshotLink field if non-nil, zero value otherwise.

### GetServiceScreenshotLinkOk

`func (o *Service) GetServiceScreenshotLinkOk() (*string, bool)`

GetServiceScreenshotLinkOk returns a tuple with the ServiceScreenshotLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceScreenshotLink

`func (o *Service) SetServiceScreenshotLink(v string)`

SetServiceScreenshotLink sets ServiceScreenshotLink field to given value.

### HasServiceScreenshotLink

`func (o *Service) HasServiceScreenshotLink() bool`

HasServiceScreenshotLink returns a boolean if a field has been set.

### GetServiceLink

`func (o *Service) GetServiceLink() string`

GetServiceLink returns the ServiceLink field if non-nil, zero value otherwise.

### GetServiceLinkOk

`func (o *Service) GetServiceLinkOk() (*string, bool)`

GetServiceLinkOk returns a tuple with the ServiceLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLink

`func (o *Service) SetServiceLink(v string)`

SetServiceLink sets ServiceLink field to given value.

### HasServiceLink

`func (o *Service) HasServiceLink() bool`

HasServiceLink returns a boolean if a field has been set.

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Service) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Service) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Service) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Service) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Service) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Service) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Service) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Service) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Service) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Service) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSiteId

`func (o *Service) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Service) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Service) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Service) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetAlive

`func (o *Service) GetAlive() bool`

GetAlive returns the Alive field if non-nil, zero value otherwise.

### GetAliveOk

`func (o *Service) GetAliveOk() (*bool, bool)`

GetAliveOk returns a tuple with the Alive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlive

`func (o *Service) SetAlive(v bool)`

SetAlive sets Alive field to given value.

### HasAlive

`func (o *Service) HasAlive() bool`

HasAlive returns a boolean if a field has been set.

### GetFirstSeen

`func (o *Service) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *Service) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *Service) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *Service) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *Service) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Service) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Service) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Service) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetDetectedBy

`func (o *Service) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *Service) GetDetectedByOk() (*string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBy

`func (o *Service) SetDetectedBy(v string)`

SetDetectedBy sets DetectedBy field to given value.

### HasDetectedBy

`func (o *Service) HasDetectedBy() bool`

HasDetectedBy returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOs

`func (o *Service) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Service) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Service) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Service) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *Service) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Service) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Service) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *Service) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetHw

`func (o *Service) GetHw() string`

GetHw returns the Hw field if non-nil, zero value otherwise.

### GetHwOk

`func (o *Service) GetHwOk() (*string, bool)`

GetHwOk returns a tuple with the Hw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHw

`func (o *Service) SetHw(v string)`

SetHw sets Hw field to given value.

### HasHw

`func (o *Service) HasHw() bool`

HasHw returns a boolean if a field has been set.

### GetAddresses

`func (o *Service) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Service) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Service) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Service) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAddressesExtra

`func (o *Service) GetAddressesExtra() []string`

GetAddressesExtra returns the AddressesExtra field if non-nil, zero value otherwise.

### GetAddressesExtraOk

`func (o *Service) GetAddressesExtraOk() (*[]string, bool)`

GetAddressesExtraOk returns a tuple with the AddressesExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesExtra

`func (o *Service) SetAddressesExtra(v []string)`

SetAddressesExtra sets AddressesExtra field to given value.

### HasAddressesExtra

`func (o *Service) HasAddressesExtra() bool`

HasAddressesExtra returns a boolean if a field has been set.

### GetMacs

`func (o *Service) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *Service) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *Service) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *Service) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetMacVendors

`func (o *Service) GetMacVendors() []string`

GetMacVendors returns the MacVendors field if non-nil, zero value otherwise.

### GetMacVendorsOk

`func (o *Service) GetMacVendorsOk() (*[]string, bool)`

GetMacVendorsOk returns a tuple with the MacVendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacVendors

`func (o *Service) SetMacVendors(v []string)`

SetMacVendors sets MacVendors field to given value.

### HasMacVendors

`func (o *Service) HasMacVendors() bool`

HasMacVendors returns a boolean if a field has been set.

### GetNames

`func (o *Service) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Service) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Service) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *Service) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetDomains

`func (o *Service) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Service) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Service) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Service) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetServiceCount

`func (o *Service) GetServiceCount() int64`

GetServiceCount returns the ServiceCount field if non-nil, zero value otherwise.

### GetServiceCountOk

`func (o *Service) GetServiceCountOk() (*int64, bool)`

GetServiceCountOk returns a tuple with the ServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCount

`func (o *Service) SetServiceCount(v int64)`

SetServiceCount sets ServiceCount field to given value.

### HasServiceCount

`func (o *Service) HasServiceCount() bool`

HasServiceCount returns a boolean if a field has been set.

### GetServiceCountTcp

`func (o *Service) GetServiceCountTcp() int64`

GetServiceCountTcp returns the ServiceCountTcp field if non-nil, zero value otherwise.

### GetServiceCountTcpOk

`func (o *Service) GetServiceCountTcpOk() (*int64, bool)`

GetServiceCountTcpOk returns a tuple with the ServiceCountTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountTcp

`func (o *Service) SetServiceCountTcp(v int64)`

SetServiceCountTcp sets ServiceCountTcp field to given value.

### HasServiceCountTcp

`func (o *Service) HasServiceCountTcp() bool`

HasServiceCountTcp returns a boolean if a field has been set.

### GetServiceCountUdp

`func (o *Service) GetServiceCountUdp() int64`

GetServiceCountUdp returns the ServiceCountUdp field if non-nil, zero value otherwise.

### GetServiceCountUdpOk

`func (o *Service) GetServiceCountUdpOk() (*int64, bool)`

GetServiceCountUdpOk returns a tuple with the ServiceCountUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountUdp

`func (o *Service) SetServiceCountUdp(v int64)`

SetServiceCountUdp sets ServiceCountUdp field to given value.

### HasServiceCountUdp

`func (o *Service) HasServiceCountUdp() bool`

HasServiceCountUdp returns a boolean if a field has been set.

### GetServiceCountArp

`func (o *Service) GetServiceCountArp() int64`

GetServiceCountArp returns the ServiceCountArp field if non-nil, zero value otherwise.

### GetServiceCountArpOk

`func (o *Service) GetServiceCountArpOk() (*int64, bool)`

GetServiceCountArpOk returns a tuple with the ServiceCountArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountArp

`func (o *Service) SetServiceCountArp(v int64)`

SetServiceCountArp sets ServiceCountArp field to given value.

### HasServiceCountArp

`func (o *Service) HasServiceCountArp() bool`

HasServiceCountArp returns a boolean if a field has been set.

### GetServiceCountIcmp

`func (o *Service) GetServiceCountIcmp() int64`

GetServiceCountIcmp returns the ServiceCountIcmp field if non-nil, zero value otherwise.

### GetServiceCountIcmpOk

`func (o *Service) GetServiceCountIcmpOk() (*int64, bool)`

GetServiceCountIcmpOk returns a tuple with the ServiceCountIcmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCountIcmp

`func (o *Service) SetServiceCountIcmp(v int64)`

SetServiceCountIcmp sets ServiceCountIcmp field to given value.

### HasServiceCountIcmp

`func (o *Service) HasServiceCountIcmp() bool`

HasServiceCountIcmp returns a boolean if a field has been set.

### GetLowestTtl

`func (o *Service) GetLowestTtl() int64`

GetLowestTtl returns the LowestTtl field if non-nil, zero value otherwise.

### GetLowestTtlOk

`func (o *Service) GetLowestTtlOk() (*int64, bool)`

GetLowestTtlOk returns a tuple with the LowestTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestTtl

`func (o *Service) SetLowestTtl(v int64)`

SetLowestTtl sets LowestTtl field to given value.

### HasLowestTtl

`func (o *Service) HasLowestTtl() bool`

HasLowestTtl returns a boolean if a field has been set.

### GetLowestRtt

`func (o *Service) GetLowestRtt() int64`

GetLowestRtt returns the LowestRtt field if non-nil, zero value otherwise.

### GetLowestRttOk

`func (o *Service) GetLowestRttOk() (*int64, bool)`

GetLowestRttOk returns a tuple with the LowestRtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestRtt

`func (o *Service) SetLowestRtt(v int64)`

SetLowestRtt sets LowestRtt field to given value.

### HasLowestRtt

`func (o *Service) HasLowestRtt() bool`

HasLowestRtt returns a boolean if a field has been set.

### GetLastAgentId

`func (o *Service) GetLastAgentId() string`

GetLastAgentId returns the LastAgentId field if non-nil, zero value otherwise.

### GetLastAgentIdOk

`func (o *Service) GetLastAgentIdOk() (*string, bool)`

GetLastAgentIdOk returns a tuple with the LastAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentId

`func (o *Service) SetLastAgentId(v string)`

SetLastAgentId sets LastAgentId field to given value.

### HasLastAgentId

`func (o *Service) HasLastAgentId() bool`

HasLastAgentId returns a boolean if a field has been set.

### GetLastTaskId

`func (o *Service) GetLastTaskId() string`

GetLastTaskId returns the LastTaskId field if non-nil, zero value otherwise.

### GetLastTaskIdOk

`func (o *Service) GetLastTaskIdOk() (*string, bool)`

GetLastTaskIdOk returns a tuple with the LastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskId

`func (o *Service) SetLastTaskId(v string)`

SetLastTaskId sets LastTaskId field to given value.

### HasLastTaskId

`func (o *Service) HasLastTaskId() bool`

HasLastTaskId returns a boolean if a field has been set.

### GetNewestMac

`func (o *Service) GetNewestMac() string`

GetNewestMac returns the NewestMac field if non-nil, zero value otherwise.

### GetNewestMacOk

`func (o *Service) GetNewestMacOk() (*string, bool)`

GetNewestMacOk returns a tuple with the NewestMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestMac

`func (o *Service) SetNewestMac(v string)`

SetNewestMac sets NewestMac field to given value.

### HasNewestMac

`func (o *Service) HasNewestMac() bool`

HasNewestMac returns a boolean if a field has been set.

### GetNewestMacVendor

`func (o *Service) GetNewestMacVendor() string`

GetNewestMacVendor returns the NewestMacVendor field if non-nil, zero value otherwise.

### GetNewestMacVendorOk

`func (o *Service) GetNewestMacVendorOk() (*string, bool)`

GetNewestMacVendorOk returns a tuple with the NewestMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestMacVendor

`func (o *Service) SetNewestMacVendor(v string)`

SetNewestMacVendor sets NewestMacVendor field to given value.

### HasNewestMacVendor

`func (o *Service) HasNewestMacVendor() bool`

HasNewestMacVendor returns a boolean if a field has been set.

### GetNewestMacAge

`func (o *Service) GetNewestMacAge() int64`

GetNewestMacAge returns the NewestMacAge field if non-nil, zero value otherwise.

### GetNewestMacAgeOk

`func (o *Service) GetNewestMacAgeOk() (*int64, bool)`

GetNewestMacAgeOk returns a tuple with the NewestMacAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestMacAge

`func (o *Service) SetNewestMacAge(v int64)`

SetNewestMacAge sets NewestMacAge field to given value.

### HasNewestMacAge

`func (o *Service) HasNewestMacAge() bool`

HasNewestMacAge returns a boolean if a field has been set.

### GetComments

`func (o *Service) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Service) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Service) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Service) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetServicePortsTcp

`func (o *Service) GetServicePortsTcp() []string`

GetServicePortsTcp returns the ServicePortsTcp field if non-nil, zero value otherwise.

### GetServicePortsTcpOk

`func (o *Service) GetServicePortsTcpOk() (*[]string, bool)`

GetServicePortsTcpOk returns a tuple with the ServicePortsTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortsTcp

`func (o *Service) SetServicePortsTcp(v []string)`

SetServicePortsTcp sets ServicePortsTcp field to given value.

### HasServicePortsTcp

`func (o *Service) HasServicePortsTcp() bool`

HasServicePortsTcp returns a boolean if a field has been set.

### GetServicePortsUdp

`func (o *Service) GetServicePortsUdp() []string`

GetServicePortsUdp returns the ServicePortsUdp field if non-nil, zero value otherwise.

### GetServicePortsUdpOk

`func (o *Service) GetServicePortsUdpOk() (*[]string, bool)`

GetServicePortsUdpOk returns a tuple with the ServicePortsUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortsUdp

`func (o *Service) SetServicePortsUdp(v []string)`

SetServicePortsUdp sets ServicePortsUdp field to given value.

### HasServicePortsUdp

`func (o *Service) HasServicePortsUdp() bool`

HasServicePortsUdp returns a boolean if a field has been set.

### GetServicePortsProtocols

`func (o *Service) GetServicePortsProtocols() []string`

GetServicePortsProtocols returns the ServicePortsProtocols field if non-nil, zero value otherwise.

### GetServicePortsProtocolsOk

`func (o *Service) GetServicePortsProtocolsOk() (*[]string, bool)`

GetServicePortsProtocolsOk returns a tuple with the ServicePortsProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortsProtocols

`func (o *Service) SetServicePortsProtocols(v []string)`

SetServicePortsProtocols sets ServicePortsProtocols field to given value.

### HasServicePortsProtocols

`func (o *Service) HasServicePortsProtocols() bool`

HasServicePortsProtocols returns a boolean if a field has been set.

### GetServicePortsProducts

`func (o *Service) GetServicePortsProducts() []string`

GetServicePortsProducts returns the ServicePortsProducts field if non-nil, zero value otherwise.

### GetServicePortsProductsOk

`func (o *Service) GetServicePortsProductsOk() (*[]string, bool)`

GetServicePortsProductsOk returns a tuple with the ServicePortsProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortsProducts

`func (o *Service) SetServicePortsProducts(v []string)`

SetServicePortsProducts sets ServicePortsProducts field to given value.

### HasServicePortsProducts

`func (o *Service) HasServicePortsProducts() bool`

HasServicePortsProducts returns a boolean if a field has been set.

### GetOrgName

`func (o *Service) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Service) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *Service) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *Service) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetSiteName

`func (o *Service) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *Service) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *Service) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *Service) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetAgentName

`func (o *Service) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *Service) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *Service) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *Service) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetTags

`func (o *Service) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Service) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetServices

`func (o *Service) GetServices() map[string]map[string]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Service) GetServicesOk() (*map[string]map[string]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Service) SetServices(v map[string]map[string]string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Service) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetRtts

`func (o *Service) GetRtts() map[string]map[string]interface{}`

GetRtts returns the Rtts field if non-nil, zero value otherwise.

### GetRttsOk

`func (o *Service) GetRttsOk() (*map[string]map[string]interface{}, bool)`

GetRttsOk returns a tuple with the Rtts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtts

`func (o *Service) SetRtts(v map[string]map[string]interface{})`

SetRtts sets Rtts field to given value.

### HasRtts

`func (o *Service) HasRtts() bool`

HasRtts returns a boolean if a field has been set.

### GetCredentials

`func (o *Service) GetCredentials() map[string]map[string]bool`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Service) GetCredentialsOk() (*map[string]map[string]bool, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Service) SetCredentials(v map[string]map[string]bool)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Service) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetAttributes

`func (o *Service) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Service) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Service) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Service) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


