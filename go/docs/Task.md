# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**CruncherId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **float32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Stats** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Recur** | Pointer to **bool** |  | [optional] 
**RecurFrequency** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**RecurLast** | Pointer to **int64** |  | [optional] 
**RecurNext** | Pointer to **int64** |  | [optional] 
**RecurLastTaskId** | Pointer to **string** |  | [optional] 

## Methods

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetClientId

`func (o *Task) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Task) GetClientIdOk() (string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientId

`func (o *Task) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientId

`func (o *Task) SetClientId(v string)`

SetClientId gets a reference to the given string and assigns it to the ClientId field.

### GetOrganizationId

`func (o *Task) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Task) GetOrganizationIdOk() (string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationId

`func (o *Task) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationId

`func (o *Task) SetOrganizationId(v string)`

SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.

### GetAgentId

`func (o *Task) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *Task) GetAgentIdOk() (string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgentId

`func (o *Task) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentId

`func (o *Task) SetAgentId(v string)`

SetAgentId gets a reference to the given string and assigns it to the AgentId field.

### GetSiteId

`func (o *Task) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Task) GetSiteIdOk() (string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSiteId

`func (o *Task) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteId

`func (o *Task) SetSiteId(v string)`

SetSiteId gets a reference to the given string and assigns it to the SiteId field.

### GetCruncherId

`func (o *Task) GetCruncherId() string`

GetCruncherId returns the CruncherId field if non-nil, zero value otherwise.

### GetCruncherIdOk

`func (o *Task) GetCruncherIdOk() (string, bool)`

GetCruncherIdOk returns a tuple with the CruncherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCruncherId

`func (o *Task) HasCruncherId() bool`

HasCruncherId returns a boolean if a field has been set.

### SetCruncherId

`func (o *Task) SetCruncherId(v string)`

SetCruncherId gets a reference to the given string and assigns it to the CruncherId field.

### GetCreatedAt

`func (o *Task) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v float32)`

SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.

### GetCreatedBy

`func (o *Task) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Task) GetCreatedByOk() (string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *Task) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *Task) SetCreatedBy(v string)`

SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.

### GetCreatedByUserId

`func (o *Task) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *Task) GetCreatedByUserIdOk() (string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedByUserId

`func (o *Task) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### SetCreatedByUserId

`func (o *Task) SetCreatedByUserId(v string)`

SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v float32)`

SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetError

`func (o *Task) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Task) GetErrorOk() (string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *Task) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *Task) SetError(v string)`

SetError gets a reference to the given string and assigns it to the Error field.

### GetParams

`func (o *Task) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Task) GetParamsOk() (map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParams

`func (o *Task) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParams

`func (o *Task) SetParams(v map[string]string)`

SetParams gets a reference to the given map[string]string and assigns it to the Params field.

### GetStats

`func (o *Task) GetStats() map[string]map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Task) GetStatsOk() (map[string]map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStats

`func (o *Task) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStats

`func (o *Task) SetStats(v map[string]map[string]interface{})`

SetStats gets a reference to the given map[string]map[string]interface{} and assigns it to the Stats field.

### GetHidden

`func (o *Task) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Task) GetHiddenOk() (bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHidden

`func (o *Task) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHidden

`func (o *Task) SetHidden(v bool)`

SetHidden gets a reference to the given bool and assigns it to the Hidden field.

### GetParentId

`func (o *Task) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Task) GetParentIdOk() (string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentId

`func (o *Task) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentId

`func (o *Task) SetParentId(v string)`

SetParentId gets a reference to the given string and assigns it to the ParentId field.

### GetRecur

`func (o *Task) GetRecur() bool`

GetRecur returns the Recur field if non-nil, zero value otherwise.

### GetRecurOk

`func (o *Task) GetRecurOk() (bool, bool)`

GetRecurOk returns a tuple with the Recur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecur

`func (o *Task) HasRecur() bool`

HasRecur returns a boolean if a field has been set.

### SetRecur

`func (o *Task) SetRecur(v bool)`

SetRecur gets a reference to the given bool and assigns it to the Recur field.

### GetRecurFrequency

`func (o *Task) GetRecurFrequency() string`

GetRecurFrequency returns the RecurFrequency field if non-nil, zero value otherwise.

### GetRecurFrequencyOk

`func (o *Task) GetRecurFrequencyOk() (string, bool)`

GetRecurFrequencyOk returns a tuple with the RecurFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurFrequency

`func (o *Task) HasRecurFrequency() bool`

HasRecurFrequency returns a boolean if a field has been set.

### SetRecurFrequency

`func (o *Task) SetRecurFrequency(v string)`

SetRecurFrequency gets a reference to the given string and assigns it to the RecurFrequency field.

### GetStartTime

`func (o *Task) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Task) GetStartTimeOk() (int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *Task) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *Task) SetStartTime(v int64)`

SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.

### GetRecurLast

`func (o *Task) GetRecurLast() int64`

GetRecurLast returns the RecurLast field if non-nil, zero value otherwise.

### GetRecurLastOk

`func (o *Task) GetRecurLastOk() (int64, bool)`

GetRecurLastOk returns a tuple with the RecurLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurLast

`func (o *Task) HasRecurLast() bool`

HasRecurLast returns a boolean if a field has been set.

### SetRecurLast

`func (o *Task) SetRecurLast(v int64)`

SetRecurLast gets a reference to the given int64 and assigns it to the RecurLast field.

### GetRecurNext

`func (o *Task) GetRecurNext() int64`

GetRecurNext returns the RecurNext field if non-nil, zero value otherwise.

### GetRecurNextOk

`func (o *Task) GetRecurNextOk() (int64, bool)`

GetRecurNextOk returns a tuple with the RecurNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurNext

`func (o *Task) HasRecurNext() bool`

HasRecurNext returns a boolean if a field has been set.

### SetRecurNext

`func (o *Task) SetRecurNext(v int64)`

SetRecurNext gets a reference to the given int64 and assigns it to the RecurNext field.

### GetRecurLastTaskId

`func (o *Task) GetRecurLastTaskId() string`

GetRecurLastTaskId returns the RecurLastTaskId field if non-nil, zero value otherwise.

### GetRecurLastTaskIdOk

`func (o *Task) GetRecurLastTaskIdOk() (string, bool)`

GetRecurLastTaskIdOk returns a tuple with the RecurLastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurLastTaskId

`func (o *Task) HasRecurLastTaskId() bool`

HasRecurLastTaskId returns a boolean if a field has been set.

### SetRecurLastTaskId

`func (o *Task) SetRecurLastTaskId(v string)`

SetRecurLastTaskId gets a reference to the given string and assigns it to the RecurLastTaskId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


