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

### NewScanOptions

`func NewScanOptions(targets string, ) *ScanOptions`

NewScanOptions instantiates a new ScanOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanOptionsWithDefaults

`func NewScanOptionsWithDefaults() *ScanOptions`

NewScanOptionsWithDefaults instantiates a new ScanOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ScanOptions) GetTargets() string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ScanOptions) GetTargetsOk() (*string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ScanOptions) SetTargets(v string)`

SetTargets sets Targets field to given value.


### GetExcludes

`func (o *ScanOptions) GetExcludes() string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *ScanOptions) GetExcludesOk() (*string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *ScanOptions) SetExcludes(v string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *ScanOptions) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetScanName

`func (o *ScanOptions) GetScanName() string`

GetScanName returns the ScanName field if non-nil, zero value otherwise.

### GetScanNameOk

`func (o *ScanOptions) GetScanNameOk() (*string, bool)`

GetScanNameOk returns a tuple with the ScanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanName

`func (o *ScanOptions) SetScanName(v string)`

SetScanName sets ScanName field to given value.

### HasScanName

`func (o *ScanOptions) HasScanName() bool`

HasScanName returns a boolean if a field has been set.

### GetScanDescription

`func (o *ScanOptions) GetScanDescription() string`

GetScanDescription returns the ScanDescription field if non-nil, zero value otherwise.

### GetScanDescriptionOk

`func (o *ScanOptions) GetScanDescriptionOk() (*string, bool)`

GetScanDescriptionOk returns a tuple with the ScanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDescription

`func (o *ScanOptions) SetScanDescription(v string)`

SetScanDescription sets ScanDescription field to given value.

### HasScanDescription

`func (o *ScanOptions) HasScanDescription() bool`

HasScanDescription returns a boolean if a field has been set.

### GetScanFrequency

`func (o *ScanOptions) GetScanFrequency() string`

GetScanFrequency returns the ScanFrequency field if non-nil, zero value otherwise.

### GetScanFrequencyOk

`func (o *ScanOptions) GetScanFrequencyOk() (*string, bool)`

GetScanFrequencyOk returns a tuple with the ScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFrequency

`func (o *ScanOptions) SetScanFrequency(v string)`

SetScanFrequency sets ScanFrequency field to given value.

### HasScanFrequency

`func (o *ScanOptions) HasScanFrequency() bool`

HasScanFrequency returns a boolean if a field has been set.

### GetScanStart

`func (o *ScanOptions) GetScanStart() int32`

GetScanStart returns the ScanStart field if non-nil, zero value otherwise.

### GetScanStartOk

`func (o *ScanOptions) GetScanStartOk() (*int32, bool)`

GetScanStartOk returns a tuple with the ScanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStart

`func (o *ScanOptions) SetScanStart(v int32)`

SetScanStart sets ScanStart field to given value.

### HasScanStart

`func (o *ScanOptions) HasScanStart() bool`

HasScanStart returns a boolean if a field has been set.

### GetAgent

`func (o *ScanOptions) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ScanOptions) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ScanOptions) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ScanOptions) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetRate

`func (o *ScanOptions) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ScanOptions) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ScanOptions) SetRate(v int32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ScanOptions) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetMaxHostRate

`func (o *ScanOptions) GetMaxHostRate() int32`

GetMaxHostRate returns the MaxHostRate field if non-nil, zero value otherwise.

### GetMaxHostRateOk

`func (o *ScanOptions) GetMaxHostRateOk() (*int32, bool)`

GetMaxHostRateOk returns a tuple with the MaxHostRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHostRate

`func (o *ScanOptions) SetMaxHostRate(v int32)`

SetMaxHostRate sets MaxHostRate field to given value.

### HasMaxHostRate

`func (o *ScanOptions) HasMaxHostRate() bool`

HasMaxHostRate returns a boolean if a field has been set.

### GetPasses

`func (o *ScanOptions) GetPasses() int32`

GetPasses returns the Passes field if non-nil, zero value otherwise.

### GetPassesOk

`func (o *ScanOptions) GetPassesOk() (*int32, bool)`

GetPassesOk returns a tuple with the Passes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasses

`func (o *ScanOptions) SetPasses(v int32)`

SetPasses sets Passes field to given value.

### HasPasses

`func (o *ScanOptions) HasPasses() bool`

HasPasses returns a boolean if a field has been set.

### GetMaxSockets

`func (o *ScanOptions) GetMaxSockets() int32`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *ScanOptions) GetMaxSocketsOk() (*int32, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *ScanOptions) SetMaxSockets(v int32)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *ScanOptions) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### GetMaxGroupSize

`func (o *ScanOptions) GetMaxGroupSize() int32`

GetMaxGroupSize returns the MaxGroupSize field if non-nil, zero value otherwise.

### GetMaxGroupSizeOk

`func (o *ScanOptions) GetMaxGroupSizeOk() (*int32, bool)`

GetMaxGroupSizeOk returns a tuple with the MaxGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGroupSize

`func (o *ScanOptions) SetMaxGroupSize(v int32)`

SetMaxGroupSize sets MaxGroupSize field to given value.

### HasMaxGroupSize

`func (o *ScanOptions) HasMaxGroupSize() bool`

HasMaxGroupSize returns a boolean if a field has been set.

### GetTcpPorts

`func (o *ScanOptions) GetTcpPorts() string`

GetTcpPorts returns the TcpPorts field if non-nil, zero value otherwise.

### GetTcpPortsOk

`func (o *ScanOptions) GetTcpPortsOk() (*string, bool)`

GetTcpPortsOk returns a tuple with the TcpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPorts

`func (o *ScanOptions) SetTcpPorts(v string)`

SetTcpPorts sets TcpPorts field to given value.

### HasTcpPorts

`func (o *ScanOptions) HasTcpPorts() bool`

HasTcpPorts returns a boolean if a field has been set.

### GetScreenshots

`func (o *ScanOptions) GetScreenshots() bool`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *ScanOptions) GetScreenshotsOk() (*bool, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *ScanOptions) SetScreenshots(v bool)`

SetScreenshots sets Screenshots field to given value.

### HasScreenshots

`func (o *ScanOptions) HasScreenshots() bool`

HasScreenshots returns a boolean if a field has been set.

### GetNameservers

`func (o *ScanOptions) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ScanOptions) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ScanOptions) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *ScanOptions) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetProbes

`func (o *ScanOptions) GetProbes() string`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *ScanOptions) GetProbesOk() (*string, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *ScanOptions) SetProbes(v string)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *ScanOptions) HasProbes() bool`

HasProbes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


