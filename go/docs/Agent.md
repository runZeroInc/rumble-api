# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**HostId** | Pointer to **string** |  | [optional] 
**HubId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**LastCheckin** | Pointer to **int64** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**SystemInfo** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**DeactivatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAgent

`func NewAgent(id string, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v string)`

SetId sets Id field to given value.


### GetClientId

`func (o *Agent) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Agent) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Agent) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Agent) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Agent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Agent) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Agent) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Agent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Agent) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Agent) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Agent) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Agent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Agent) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Agent) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Agent) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Agent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHostId

`func (o *Agent) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *Agent) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *Agent) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *Agent) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetHubId

`func (o *Agent) GetHubId() string`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *Agent) GetHubIdOk() (*string, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *Agent) SetHubId(v string)`

SetHubId sets HubId field to given value.

### HasHubId

`func (o *Agent) HasHubId() bool`

HasHubId returns a boolean if a field has been set.

### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Agent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteId

`func (o *Agent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Agent) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Agent) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Agent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetLastCheckin

`func (o *Agent) GetLastCheckin() int64`

GetLastCheckin returns the LastCheckin field if non-nil, zero value otherwise.

### GetLastCheckinOk

`func (o *Agent) GetLastCheckinOk() (*int64, bool)`

GetLastCheckinOk returns a tuple with the LastCheckin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckin

`func (o *Agent) SetLastCheckin(v int64)`

SetLastCheckin sets LastCheckin field to given value.

### HasLastCheckin

`func (o *Agent) HasLastCheckin() bool`

HasLastCheckin returns a boolean if a field has been set.

### GetOs

`func (o *Agent) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Agent) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Agent) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Agent) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetArch

`func (o *Agent) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *Agent) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *Agent) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *Agent) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetVersion

`func (o *Agent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Agent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Agent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Agent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetExternalIp

`func (o *Agent) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *Agent) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *Agent) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *Agent) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *Agent) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *Agent) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *Agent) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *Agent) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetSystemInfo

`func (o *Agent) GetSystemInfo() map[string]map[string]interface{}`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *Agent) GetSystemInfoOk() (*map[string]map[string]interface{}, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *Agent) SetSystemInfo(v map[string]map[string]interface{})`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *Agent) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.

### GetConnected

`func (o *Agent) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *Agent) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *Agent) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *Agent) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetInactive

`func (o *Agent) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *Agent) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *Agent) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *Agent) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetDeactivatedAt

`func (o *Agent) GetDeactivatedAt() int64`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *Agent) GetDeactivatedAtOk() (*int64, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *Agent) SetDeactivatedAt(v int64)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *Agent) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


