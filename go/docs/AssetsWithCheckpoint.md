# AssetsWithCheckpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | Pointer to **int64** |  | 
**Assets** | Pointer to [**[]Asset**](Asset.md) |  | 

## Methods

### NewAssetsWithCheckpoint

`func NewAssetsWithCheckpoint(since int64, assets []Asset, ) *AssetsWithCheckpoint`

NewAssetsWithCheckpoint instantiates a new AssetsWithCheckpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsWithCheckpointWithDefaults

`func NewAssetsWithCheckpointWithDefaults() *AssetsWithCheckpoint`

NewAssetsWithCheckpointWithDefaults instantiates a new AssetsWithCheckpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *AssetsWithCheckpoint) GetSince() int64`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *AssetsWithCheckpoint) GetSinceOk() (*int64, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *AssetsWithCheckpoint) SetSince(v int64)`

SetSince sets Since field to given value.


### GetAssets

`func (o *AssetsWithCheckpoint) GetAssets() []Asset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AssetsWithCheckpoint) GetAssetsOk() (*[]Asset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AssetsWithCheckpoint) SetAssets(v []Asset)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


