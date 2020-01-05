# \OrganizationApi

All URIs are relative to *https://console.rumble.run/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScan**](OrganizationApi.md#CreateScan) | **Put** /org/sites/{site_id}/scan | Create a scan task for a given site.
[**CreateSite**](OrganizationApi.md#CreateSite) | **Put** /org/sites | Create a new site.
[**GetAgent**](OrganizationApi.md#GetAgent) | **Get** /org/agents/{agent_id} | Get details for a single agent.
[**GetAgents**](OrganizationApi.md#GetAgents) | **Get** /org/agents | Get all agents.
[**GetAsset**](OrganizationApi.md#GetAsset) | **Get** /org/assets/{asset_id} | Get asset details.
[**GetAssets**](OrganizationApi.md#GetAssets) | **Get** /org/assets | Get all assets.
[**GetKey**](OrganizationApi.md#GetKey) | **Get** /org/key | Get API key details.
[**GetOrganization**](OrganizationApi.md#GetOrganization) | **Get** /org | Get organization details.
[**GetService**](OrganizationApi.md#GetService) | **Get** /org/services/{service_id} | Get service details.
[**GetServices**](OrganizationApi.md#GetServices) | **Get** /org/services | Get all services.
[**GetSite**](OrganizationApi.md#GetSite) | **Get** /org/sites/{site_id} | Get site details.
[**GetSites**](OrganizationApi.md#GetSites) | **Get** /org/sites | Get all sites.
[**GetTask**](OrganizationApi.md#GetTask) | **Get** /org/tasks/{task_id} | Get task details.
[**GetTaskChangeReport**](OrganizationApi.md#GetTaskChangeReport) | **Get** /org/tasks/{task_id}/changes | Returns a temporary task change report data url.
[**GetTaskScanData**](OrganizationApi.md#GetTaskScanData) | **Get** /org/tasks/{task_id}/data | Returns a temporary task scan data url.
[**GetTasks**](OrganizationApi.md#GetTasks) | **Get** /org/tasks | Get all tasks (last 1000).
[**GetWirelessLAN**](OrganizationApi.md#GetWirelessLAN) | **Get** /org/wirelesss/{wireless_id} | Get wireless LAN details.
[**GetWirelessLANs**](OrganizationApi.md#GetWirelessLANs) | **Get** /org/wireless | Get all wireless LANs.
[**HideTask**](OrganizationApi.md#HideTask) | **Post** /org/tasks/{task_id}/hide | Signal that a completed task should be hidden.
[**ImportScanData**](OrganizationApi.md#ImportScanData) | **Put** /org/sites/{site_id}/import | Import a scan data file into a site.
[**RemoveAgent**](OrganizationApi.md#RemoveAgent) | **Delete** /org/agents/{agent_id} | Remove and uninstall an agent.
[**RemoveAsset**](OrganizationApi.md#RemoveAsset) | **Delete** /org/assets/{asset_id} | Remove an asset.
[**RemoveService**](OrganizationApi.md#RemoveService) | **Delete** /org/services/{service_id} | Remove a service.
[**RemoveSite**](OrganizationApi.md#RemoveSite) | **Delete** /org/sites/{site_id} | Remove a site and associated assets.
[**RemoveWirelessLAN**](OrganizationApi.md#RemoveWirelessLAN) | **Delete** /org/wirelesss/{wireless_id} | Remove a wireless LAN.
[**StopTask**](OrganizationApi.md#StopTask) | **Post** /org/tasks/{task_id}/stop | Signal that a task should be stopped or canceled.
[**UpdateAgentSite**](OrganizationApi.md#UpdateAgentSite) | **Patch** /org/agents/{agent_id} | Update the site associated with agent.
[**UpdateAssetComments**](OrganizationApi.md#UpdateAssetComments) | **Patch** /org/assets/{asset_id}/comments | Update asset comments.
[**UpdateAssetTags**](OrganizationApi.md#UpdateAssetTags) | **Patch** /org/assets/{asset_id}/tags | Update asset tags.
[**UpdateOrganization**](OrganizationApi.md#UpdateOrganization) | **Patch** /org | Update organization details.
[**UpdateSite**](OrganizationApi.md#UpdateSite) | **Patch** /org/sites/{site_id} | Update a site definition.
[**UpdateTask**](OrganizationApi.md#UpdateTask) | **Patch** /org/tasks/{task_id} | Update task parameters.
[**UpgradeAgent**](OrganizationApi.md#UpgradeAgent) | **Post** /org/agents/{agent_id}/update | Force an agent to update and restart.



## CreateScan

> Task CreateScan(ctx, siteId, optional)

Create a scan task for a given site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md)| UUID or name of the site to scan | 
 **optional** | ***CreateScanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateScanOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanOptions** | [**optional.Interface of ScanOptions**](ScanOptions.md)|  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSite

> Site CreateSite(ctx, siteOptions)

Create a new site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteOptions** | [**SiteOptions**](SiteOptions.md)| site definition | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> Agent GetAgent(ctx, agentId)

Get details for a single agent.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md)| UUID of the agent | 

### Return type

[**Agent**](Agent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> []Agent GetAgents(ctx, )

Get all agents.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Agent**](Agent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> Asset GetAsset(ctx, assetId)

Get asset details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md)| UUID of the asset to retrieve | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> []Asset GetAssets(ctx, optional)

Get all assets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAssetsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKey

> ApiKey GetKey(ctx, )

Get API key details.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ApiKey**](APIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> Organization GetOrganization(ctx, )

Get organization details.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> Service GetService(ctx, serviceId)

Get service details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | [**string**](.md)| UUID of the service to retrieve | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []Service GetServices(ctx, optional)

Get all services.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetServicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSite

> Site GetSite(ctx, siteId)

Get site details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md)| UUID or name of the site | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSites

> []Site GetSites(ctx, )

Get all sites.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, taskId)

Get task details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md)| UUID of the task to retrieve | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskChangeReport

> GetTaskChangeReport(ctx, taskId)

Returns a temporary task change report data url.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md)| UUID of the task | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskScanData

> GetTaskScanData(ctx, taskId)

Returns a temporary task scan data url.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md)| UUID of the task | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []Task GetTasks(ctx, optional)

Get all tasks (last 1000).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTasksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| an optional status string for filtering results | 

### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessLAN

> Wireless GetWirelessLAN(ctx, wirelessId)

Get wireless LAN details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wirelessId** | [**string**](.md)| UUID of the wireless LAN to retrieve | 

### Return type

[**Wireless**](Wireless.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessLANs

> []Wireless GetWirelessLANs(ctx, optional)

Get all wireless LANs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWirelessLANsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWirelessLANsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Wireless**](Wireless.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideTask

> HideTask(ctx, taskId)

Signal that a completed task should be hidden.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md)| UUID of the task to hide | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportScanData

> Task ImportScanData(ctx, siteId, optional)

Import a scan data file into a site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md)| UUID or name of the site to import scan data into | 
 **optional** | ***ImportScanDataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ImportScanDataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAgent

> RemoveAgent(ctx, agentId)

Remove and uninstall an agent.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md)| UUID of the agent to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAsset

> RemoveAsset(ctx, assetId)

Remove an asset.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md)| UUID of the asset to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveService

> RemoveService(ctx, serviceId)

Remove a service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | [**string**](.md)| UUID of the service to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSite

> RemoveSite(ctx, siteId)

Remove a site and associated assets.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md)| UUID or name of the site to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWirelessLAN

> RemoveWirelessLAN(ctx, wirelessId)

Remove a wireless LAN.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wirelessId** | [**string**](.md)| UUID of the wireless LAN to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopTask

> StopTask(ctx, taskId)

Signal that a task should be stopped or canceled.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md)| UUID of the task to stop | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentSite

> Agent UpdateAgentSite(ctx, agentId, agentSiteId)

Update the site associated with agent.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md)| UUID of the agent to update | 
**agentSiteId** | [**AgentSiteId**](AgentSiteId.md)| site_id to associate with the agent | 

### Return type

[**Agent**](Agent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetComments

> Asset UpdateAssetComments(ctx, assetId, assetComments)

Update asset comments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md)| UUID of the asset to update | 
**assetComments** | [**AssetComments**](AssetComments.md)| comments to apply to the asset | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetTags

> Asset UpdateAssetTags(ctx, assetId, assetTags)

Update asset tags.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md)| UUID of the agent to update | 
**assetTags** | [**AssetTags**](AssetTags.md)| tags to apply to the asset | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> Organization UpdateOrganization(ctx, orgOptions)

Update organization details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgOptions** | [**OrgOptions**](OrgOptions.md)| organization options | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSite

> Site UpdateSite(ctx, siteId, siteOptions)

Update a site definition.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md)| UUID or name of the site to update | 
**siteOptions** | [**SiteOptions**](SiteOptions.md)| site object | 

### Return type

[**Site**](Site.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> Task UpdateTask(ctx, taskId, task)

Update task parameters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md)| UUID of the task to update | 
**task** | [**Task**](Task.md)| task object | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeAgent

> UpgradeAgent(ctx, agentId)

Force an agent to update and restart.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md)| UUID of the agent to update | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

