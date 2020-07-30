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
**CreatedAt** | Pointer to **int64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Stats** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Recur** | Pointer to **bool** |  | [optional] 
**RecurFrequency** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**RecurLast** | Pointer to **int64** |  | [optional] 
**RecurNext** | Pointer to **int64** |  | [optional] 
**RecurLastTaskId** | Pointer to **string** |  | [optional] 

## Methods

### NewTask

`func NewTask(id string, ) *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClientId

`func (o *Task) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Task) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Task) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Task) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Task) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Task) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Task) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Task) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetAgentId

`func (o *Task) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *Task) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *Task) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *Task) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetSiteId

`func (o *Task) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Task) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Task) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Task) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCruncherId

`func (o *Task) GetCruncherId() string`

GetCruncherId returns the CruncherId field if non-nil, zero value otherwise.

### GetCruncherIdOk

`func (o *Task) GetCruncherIdOk() (*string, bool)`

GetCruncherIdOk returns a tuple with the CruncherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCruncherId

`func (o *Task) SetCruncherId(v string)`

SetCruncherId sets CruncherId field to given value.

### HasCruncherId

`func (o *Task) HasCruncherId() bool`

HasCruncherId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Task) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Task) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Task) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Task) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Task) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *Task) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *Task) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *Task) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *Task) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Task) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Task) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Task) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Task) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *Task) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Task) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Task) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Task) HasError() bool`

HasError returns a boolean if a field has been set.

### GetParams

`func (o *Task) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Task) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Task) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *Task) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetStats

`func (o *Task) GetStats() map[string]map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Task) GetStatsOk() (*map[string]map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Task) SetStats(v map[string]map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *Task) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetHidden

`func (o *Task) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Task) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Task) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Task) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetParentId

`func (o *Task) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Task) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Task) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Task) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRecur

`func (o *Task) GetRecur() bool`

GetRecur returns the Recur field if non-nil, zero value otherwise.

### GetRecurOk

`func (o *Task) GetRecurOk() (*bool, bool)`

GetRecurOk returns a tuple with the Recur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecur

`func (o *Task) SetRecur(v bool)`

SetRecur sets Recur field to given value.

### HasRecur

`func (o *Task) HasRecur() bool`

HasRecur returns a boolean if a field has been set.

### GetRecurFrequency

`func (o *Task) GetRecurFrequency() string`

GetRecurFrequency returns the RecurFrequency field if non-nil, zero value otherwise.

### GetRecurFrequencyOk

`func (o *Task) GetRecurFrequencyOk() (*string, bool)`

GetRecurFrequencyOk returns a tuple with the RecurFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurFrequency

`func (o *Task) SetRecurFrequency(v string)`

SetRecurFrequency sets RecurFrequency field to given value.

### HasRecurFrequency

`func (o *Task) HasRecurFrequency() bool`

HasRecurFrequency returns a boolean if a field has been set.

### GetStartTime

`func (o *Task) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Task) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Task) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Task) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetRecurLast

`func (o *Task) GetRecurLast() int64`

GetRecurLast returns the RecurLast field if non-nil, zero value otherwise.

### GetRecurLastOk

`func (o *Task) GetRecurLastOk() (*int64, bool)`

GetRecurLastOk returns a tuple with the RecurLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLast

`func (o *Task) SetRecurLast(v int64)`

SetRecurLast sets RecurLast field to given value.

### HasRecurLast

`func (o *Task) HasRecurLast() bool`

HasRecurLast returns a boolean if a field has been set.

### GetRecurNext

`func (o *Task) GetRecurNext() int64`

GetRecurNext returns the RecurNext field if non-nil, zero value otherwise.

### GetRecurNextOk

`func (o *Task) GetRecurNextOk() (*int64, bool)`

GetRecurNextOk returns a tuple with the RecurNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurNext

`func (o *Task) SetRecurNext(v int64)`

SetRecurNext sets RecurNext field to given value.

### HasRecurNext

`func (o *Task) HasRecurNext() bool`

HasRecurNext returns a boolean if a field has been set.

### GetRecurLastTaskId

`func (o *Task) GetRecurLastTaskId() string`

GetRecurLastTaskId returns the RecurLastTaskId field if non-nil, zero value otherwise.

### GetRecurLastTaskIdOk

`func (o *Task) GetRecurLastTaskIdOk() (*string, bool)`

GetRecurLastTaskIdOk returns a tuple with the RecurLastTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurLastTaskId

`func (o *Task) SetRecurLastTaskId(v string)`

SetRecurLastTaskId sets RecurLastTaskId field to given value.

### HasRecurLastTaskId

`func (o *Task) HasRecurLastTaskId() bool`

HasRecurLastTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


