# Wireless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**LastAgentId** | Pointer to **string** |  | [optional] 
**LastTaskId** | Pointer to **string** |  | [optional] 
**Essid** | Pointer to **string** |  | [optional] 
**Bssid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**Encryption** | Pointer to **string** |  | [optional] 
**Signal** | Pointer to **int32** |  | [optional] 
**Channels** | Pointer to **string** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**SiteName** | Pointer to **string** |  | [optional] 
**AgentName** | Pointer to **string** |  | [optional] 

## Methods

### NewWireless

`func NewWireless(id string, ) *Wireless`

NewWireless instantiates a new Wireless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessWithDefaults

`func NewWirelessWithDefaults() *Wireless`

NewWirelessWithDefaults instantiates a new Wireless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Wireless) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wireless) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wireless) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Wireless) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Wireless) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Wireless) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Wireless) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSeen

`func (o *Wireless) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Wireless) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Wireless) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Wireless) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Wireless) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Wireless) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Wireless) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Wireless) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSiteId

`func (o *Wireless) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Wireless) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Wireless) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Wireless) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastAgentId

`func (o *Wireless) GetLastAgentId() string`

GetLastAgentId returns the LastAgentId field if non-nil, zero value otherwise.

### GetLastAgentIdOk

`func (o *Wireless) GetLastAgentIdOk() (*string, bool)`

GetLastAgentIdOk returns a tuple with the LastAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentId

`func (o *Wireless) SetLastAgentId(v string)`

SetLastAgentId sets LastAgentId field to given value.

### HasLastAgentId

`func (o *Wireless) HasLastAgentId() bool`

HasLastAgentId returns a boolean if a field has been set.

### GetLastTaskId

`func (o *Wireless) GetLastTaskId() string`

GetLastTaskId returns the LastTaskId field if non-nil, zero value otherwise.

### GetLastTaskIdOk

`func (o *Wireless) GetLastTaskIdOk() (*string, bool)`

GetLastTaskIdOk returns a tuple with the LastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTaskId

`func (o *Wireless) SetLastTaskId(v string)`

SetLastTaskId sets LastTaskId field to given value.

### HasLastTaskId

`func (o *Wireless) HasLastTaskId() bool`

HasLastTaskId returns a boolean if a field has been set.

### GetEssid

`func (o *Wireless) GetEssid() string`

GetEssid returns the Essid field if non-nil, zero value otherwise.

### GetEssidOk

`func (o *Wireless) GetEssidOk() (*string, bool)`

GetEssidOk returns a tuple with the Essid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEssid

`func (o *Wireless) SetEssid(v string)`

SetEssid sets Essid field to given value.

### HasEssid

`func (o *Wireless) HasEssid() bool`

HasEssid returns a boolean if a field has been set.

### GetBssid

`func (o *Wireless) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *Wireless) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *Wireless) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *Wireless) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetType

`func (o *Wireless) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Wireless) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Wireless) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Wireless) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAuthentication

`func (o *Wireless) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *Wireless) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *Wireless) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *Wireless) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetEncryption

`func (o *Wireless) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Wireless) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Wireless) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Wireless) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetSignal

`func (o *Wireless) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *Wireless) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *Wireless) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *Wireless) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetChannels

`func (o *Wireless) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Wireless) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *Wireless) SetChannels(v string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *Wireless) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetInterface

`func (o *Wireless) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *Wireless) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *Wireless) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *Wireless) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetVendor

`func (o *Wireless) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Wireless) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Wireless) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Wireless) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetFamily

`func (o *Wireless) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *Wireless) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *Wireless) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *Wireless) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetData

`func (o *Wireless) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Wireless) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Wireless) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *Wireless) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOrgName

`func (o *Wireless) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Wireless) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *Wireless) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *Wireless) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetSiteName

`func (o *Wireless) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *Wireless) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *Wireless) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *Wireless) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetAgentName

`func (o *Wireless) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *Wireless) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *Wireless) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *Wireless) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


