# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **float32** |  | [optional] 
**UpdatedAt** | Pointer to **float32** |  | [optional] 
**HostId** | Pointer to **string** |  | [optional] 
**HubId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**LastCheckin** | Pointer to **float32** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**SystemInfo** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**DeactivatedAt** | Pointer to **float32** |  | [optional] 

## Methods

### GetId

`func (o *Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Agent) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetClientId

`func (o *Agent) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Agent) GetClientIdOk() (string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientId

`func (o *Agent) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientId

`func (o *Agent) SetClientId(v string)`

SetClientId gets a reference to the given string and assigns it to the ClientId field.

### GetOrganizationId

`func (o *Agent) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Agent) GetOrganizationIdOk() (string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationId

`func (o *Agent) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationId

`func (o *Agent) SetOrganizationId(v string)`

SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.

### GetCreatedAt

`func (o *Agent) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Agent) GetCreatedAtOk() (float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *Agent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *Agent) SetCreatedAt(v float32)`

SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.

### GetUpdatedAt

`func (o *Agent) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Agent) GetUpdatedAtOk() (float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedAt

`func (o *Agent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAt

`func (o *Agent) SetUpdatedAt(v float32)`

SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.

### GetHostId

`func (o *Agent) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *Agent) GetHostIdOk() (string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostId

`func (o *Agent) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### SetHostId

`func (o *Agent) SetHostId(v string)`

SetHostId gets a reference to the given string and assigns it to the HostId field.

### GetHubId

`func (o *Agent) GetHubId() string`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *Agent) GetHubIdOk() (string, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHubId

`func (o *Agent) HasHubId() bool`

HasHubId returns a boolean if a field has been set.

### SetHubId

`func (o *Agent) SetHubId(v string)`

SetHubId gets a reference to the given string and assigns it to the HubId field.

### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Agent) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSiteId

`func (o *Agent) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Agent) GetSiteIdOk() (string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteId

`func (o *Agent) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteId

`func (o *Agent) SetSiteId(v string)`

SetSiteId gets a reference to the given string and assigns it to the SiteId field.

### GetLastCheckin

`func (o *Agent) GetLastCheckin() float32`

GetLastCheckin returns the LastCheckin field if non-nil, zero value otherwise.

### GetLastCheckinOk

`func (o *Agent) GetLastCheckinOk() (float32, bool)`

GetLastCheckinOk returns a tuple with the LastCheckin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastCheckin

`func (o *Agent) HasLastCheckin() bool`

HasLastCheckin returns a boolean if a field has been set.

### SetLastCheckin

`func (o *Agent) SetLastCheckin(v float32)`

SetLastCheckin gets a reference to the given float32 and assigns it to the LastCheckin field.

### GetOs

`func (o *Agent) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Agent) GetOsOk() (string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOs

`func (o *Agent) HasOs() bool`

HasOs returns a boolean if a field has been set.

### SetOs

`func (o *Agent) SetOs(v string)`

SetOs gets a reference to the given string and assigns it to the Os field.

### GetArch

`func (o *Agent) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *Agent) GetArchOk() (string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArch

`func (o *Agent) HasArch() bool`

HasArch returns a boolean if a field has been set.

### SetArch

`func (o *Agent) SetArch(v string)`

SetArch gets a reference to the given string and assigns it to the Arch field.

### GetVersion

`func (o *Agent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Agent) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *Agent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *Agent) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### GetExternalIp

`func (o *Agent) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *Agent) GetExternalIpOk() (string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalIp

`func (o *Agent) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIp

`func (o *Agent) SetExternalIp(v string)`

SetExternalIp gets a reference to the given string and assigns it to the ExternalIp field.

### GetInternalIp

`func (o *Agent) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *Agent) GetInternalIpOk() (string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternalIp

`func (o *Agent) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIp

`func (o *Agent) SetInternalIp(v string)`

SetInternalIp gets a reference to the given string and assigns it to the InternalIp field.

### GetSystemInfo

`func (o *Agent) GetSystemInfo() map[string]map[string]interface{}`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *Agent) GetSystemInfoOk() (map[string]map[string]interface{}, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSystemInfo

`func (o *Agent) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.

### SetSystemInfo

`func (o *Agent) SetSystemInfo(v map[string]map[string]interface{})`

SetSystemInfo gets a reference to the given map[string]map[string]interface{} and assigns it to the SystemInfo field.

### GetConnected

`func (o *Agent) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *Agent) GetConnectedOk() (bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnected

`func (o *Agent) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### SetConnected

`func (o *Agent) SetConnected(v bool)`

SetConnected gets a reference to the given bool and assigns it to the Connected field.

### GetInactive

`func (o *Agent) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *Agent) GetInactiveOk() (bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInactive

`func (o *Agent) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### SetInactive

`func (o *Agent) SetInactive(v bool)`

SetInactive gets a reference to the given bool and assigns it to the Inactive field.

### GetDeactivatedAt

`func (o *Agent) GetDeactivatedAt() float32`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *Agent) GetDeactivatedAtOk() (float32, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeactivatedAt

`func (o *Agent) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.

### SetDeactivatedAt

`func (o *Agent) SetDeactivatedAt(v float32)`

SetDeactivatedAt gets a reference to the given float32 and assigns it to the DeactivatedAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


