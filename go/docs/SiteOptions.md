# SiteOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Excludes** | Pointer to **string** |  | [optional] 

## Methods

### NewSiteOptions

`func NewSiteOptions(name string, ) *SiteOptions`

NewSiteOptions instantiates a new SiteOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteOptionsWithDefaults

`func NewSiteOptionsWithDefaults() *SiteOptions`

NewSiteOptionsWithDefaults instantiates a new SiteOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SiteOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteOptions) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SiteOptions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SiteOptions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SiteOptions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SiteOptions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *SiteOptions) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SiteOptions) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SiteOptions) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SiteOptions) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetExcludes

`func (o *SiteOptions) GetExcludes() string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *SiteOptions) GetExcludesOk() (*string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *SiteOptions) SetExcludes(v string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *SiteOptions) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


