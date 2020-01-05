# ScanOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to **string** |  | 
**Excludes** | Pointer to **string** |  | [optional] 
**ScanName** | Pointer to **string** |  | [optional] 
**ScanDescription** | Pointer to **string** |  | [optional] 
**ScanFrequency** | Pointer to **string** |  | [optional] 
**ScanStart** | Pointer to **int32** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **int32** |  | [optional] 
**MaxHostRate** | Pointer to **int32** |  | [optional] 
**Passes** | Pointer to **int32** |  | [optional] 
**MaxSockets** | Pointer to **int32** |  | [optional] 
**MaxGroupSize** | Pointer to **int32** |  | [optional] 
**TcpPorts** | Pointer to **string** |  | [optional] 
**Screenshots** | Pointer to **bool** |  | [optional] 
**Nameservers** | Pointer to **string** |  | [optional] 
**Probes** | Pointer to **string** | Optional probe list, otherwise all probes are used | [optional] 

## Methods

### GetTargets

`func (o *ScanOptions) GetTargets() string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ScanOptions) GetTargetsOk() (string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargets

`func (o *ScanOptions) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargets

`func (o *ScanOptions) SetTargets(v string)`

SetTargets gets a reference to the given string and assigns it to the Targets field.

### GetExcludes

`func (o *ScanOptions) GetExcludes() string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *ScanOptions) GetExcludesOk() (string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExcludes

`func (o *ScanOptions) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### SetExcludes

`func (o *ScanOptions) SetExcludes(v string)`

SetExcludes gets a reference to the given string and assigns it to the Excludes field.

### GetScanName

`func (o *ScanOptions) GetScanName() string`

GetScanName returns the ScanName field if non-nil, zero value otherwise.

### GetScanNameOk

`func (o *ScanOptions) GetScanNameOk() (string, bool)`

GetScanNameOk returns a tuple with the ScanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScanName

`func (o *ScanOptions) HasScanName() bool`

HasScanName returns a boolean if a field has been set.

### SetScanName

`func (o *ScanOptions) SetScanName(v string)`

SetScanName gets a reference to the given string and assigns it to the ScanName field.

### GetScanDescription

`func (o *ScanOptions) GetScanDescription() string`

GetScanDescription returns the ScanDescription field if non-nil, zero value otherwise.

### GetScanDescriptionOk

`func (o *ScanOptions) GetScanDescriptionOk() (string, bool)`

GetScanDescriptionOk returns a tuple with the ScanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScanDescription

`func (o *ScanOptions) HasScanDescription() bool`

HasScanDescription returns a boolean if a field has been set.

### SetScanDescription

`func (o *ScanOptions) SetScanDescription(v string)`

SetScanDescription gets a reference to the given string and assigns it to the ScanDescription field.

### GetScanFrequency

`func (o *ScanOptions) GetScanFrequency() string`

GetScanFrequency returns the ScanFrequency field if non-nil, zero value otherwise.

### GetScanFrequencyOk

`func (o *ScanOptions) GetScanFrequencyOk() (string, bool)`

GetScanFrequencyOk returns a tuple with the ScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScanFrequency

`func (o *ScanOptions) HasScanFrequency() bool`

HasScanFrequency returns a boolean if a field has been set.

### SetScanFrequency

`func (o *ScanOptions) SetScanFrequency(v string)`

SetScanFrequency gets a reference to the given string and assigns it to the ScanFrequency field.

### GetScanStart

`func (o *ScanOptions) GetScanStart() int32`

GetScanStart returns the ScanStart field if non-nil, zero value otherwise.

### GetScanStartOk

`func (o *ScanOptions) GetScanStartOk() (int32, bool)`

GetScanStartOk returns a tuple with the ScanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScanStart

`func (o *ScanOptions) HasScanStart() bool`

HasScanStart returns a boolean if a field has been set.

### SetScanStart

`func (o *ScanOptions) SetScanStart(v int32)`

SetScanStart gets a reference to the given int32 and assigns it to the ScanStart field.

### GetAgent

`func (o *ScanOptions) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ScanOptions) GetAgentOk() (string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgent

`func (o *ScanOptions) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### SetAgent

`func (o *ScanOptions) SetAgent(v string)`

SetAgent gets a reference to the given string and assigns it to the Agent field.

### GetRate

`func (o *ScanOptions) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ScanOptions) GetRateOk() (int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRate

`func (o *ScanOptions) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRate

`func (o *ScanOptions) SetRate(v int32)`

SetRate gets a reference to the given int32 and assigns it to the Rate field.

### GetMaxHostRate

`func (o *ScanOptions) GetMaxHostRate() int32`

GetMaxHostRate returns the MaxHostRate field if non-nil, zero value otherwise.

### GetMaxHostRateOk

`func (o *ScanOptions) GetMaxHostRateOk() (int32, bool)`

GetMaxHostRateOk returns a tuple with the MaxHostRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxHostRate

`func (o *ScanOptions) HasMaxHostRate() bool`

HasMaxHostRate returns a boolean if a field has been set.

### SetMaxHostRate

`func (o *ScanOptions) SetMaxHostRate(v int32)`

SetMaxHostRate gets a reference to the given int32 and assigns it to the MaxHostRate field.

### GetPasses

`func (o *ScanOptions) GetPasses() int32`

GetPasses returns the Passes field if non-nil, zero value otherwise.

### GetPassesOk

`func (o *ScanOptions) GetPassesOk() (int32, bool)`

GetPassesOk returns a tuple with the Passes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasses

`func (o *ScanOptions) HasPasses() bool`

HasPasses returns a boolean if a field has been set.

### SetPasses

`func (o *ScanOptions) SetPasses(v int32)`

SetPasses gets a reference to the given int32 and assigns it to the Passes field.

### GetMaxSockets

`func (o *ScanOptions) GetMaxSockets() int32`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *ScanOptions) GetMaxSocketsOk() (int32, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxSockets

`func (o *ScanOptions) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### SetMaxSockets

`func (o *ScanOptions) SetMaxSockets(v int32)`

SetMaxSockets gets a reference to the given int32 and assigns it to the MaxSockets field.

### GetMaxGroupSize

`func (o *ScanOptions) GetMaxGroupSize() int32`

GetMaxGroupSize returns the MaxGroupSize field if non-nil, zero value otherwise.

### GetMaxGroupSizeOk

`func (o *ScanOptions) GetMaxGroupSizeOk() (int32, bool)`

GetMaxGroupSizeOk returns a tuple with the MaxGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxGroupSize

`func (o *ScanOptions) HasMaxGroupSize() bool`

HasMaxGroupSize returns a boolean if a field has been set.

### SetMaxGroupSize

`func (o *ScanOptions) SetMaxGroupSize(v int32)`

SetMaxGroupSize gets a reference to the given int32 and assigns it to the MaxGroupSize field.

### GetTcpPorts

`func (o *ScanOptions) GetTcpPorts() string`

GetTcpPorts returns the TcpPorts field if non-nil, zero value otherwise.

### GetTcpPortsOk

`func (o *ScanOptions) GetTcpPortsOk() (string, bool)`

GetTcpPortsOk returns a tuple with the TcpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTcpPorts

`func (o *ScanOptions) HasTcpPorts() bool`

HasTcpPorts returns a boolean if a field has been set.

### SetTcpPorts

`func (o *ScanOptions) SetTcpPorts(v string)`

SetTcpPorts gets a reference to the given string and assigns it to the TcpPorts field.

### GetScreenshots

`func (o *ScanOptions) GetScreenshots() bool`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *ScanOptions) GetScreenshotsOk() (bool, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenshots

`func (o *ScanOptions) HasScreenshots() bool`

HasScreenshots returns a boolean if a field has been set.

### SetScreenshots

`func (o *ScanOptions) SetScreenshots(v bool)`

SetScreenshots gets a reference to the given bool and assigns it to the Screenshots field.

### GetNameservers

`func (o *ScanOptions) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ScanOptions) GetNameserversOk() (string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNameservers

`func (o *ScanOptions) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### SetNameservers

`func (o *ScanOptions) SetNameservers(v string)`

SetNameservers gets a reference to the given string and assigns it to the Nameservers field.

### GetProbes

`func (o *ScanOptions) GetProbes() string`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *ScanOptions) GetProbesOk() (string, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProbes

`func (o *ScanOptions) HasProbes() bool`

HasProbes returns a boolean if a field has been set.

### SetProbes

`func (o *ScanOptions) SetProbes(v string)`

SetProbes gets a reference to the given string and assigns it to the Probes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


