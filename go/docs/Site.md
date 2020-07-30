# Site

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**Permanent** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Excludes** | Pointer to **string** |  | [optional] 

## Methods

### NewSite

`func NewSite(id string, name string, ) *Site`

NewSite instantiates a new Site object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteWithDefaults

`func NewSiteWithDefaults() *Site`

NewSiteWithDefaults instantiates a new Site object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Site) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Site) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Site) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Site) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Site) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Site) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Site) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Site) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Site) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Site) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Site) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPermanent

`func (o *Site) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *Site) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *Site) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *Site) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.

### GetName

`func (o *Site) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Site) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Site) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Site) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Site) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Site) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Site) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *Site) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Site) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Site) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Site) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExcludes

`func (o *Site) GetExcludes() string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *Site) GetExcludesOk() (*string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *Site) SetExcludes(v string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *Site) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


