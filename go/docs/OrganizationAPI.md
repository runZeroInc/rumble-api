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

> Task CreateScan(ctx, siteId).ScanOptions(scanOptions).Execute()

Create a scan task for a given site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteId := TODO // string | UUID or name of the site to scan
    scanOptions := openapiclient.ScanOptions{Targets: "Targets_example", Excludes: "Excludes_example", ScanName: "ScanName_example", ScanDescription: "ScanDescription_example", ScanFrequency: "ScanFrequency_example", ScanStart: 123, Agent: "Agent_example", Rate: 123, MaxHostRate: 123, Passes: 123, MaxSockets: 123, MaxGroupSize: 123, TcpPorts: "TcpPorts_example", Screenshots: true, Nameservers: "Nameservers_example", Probes: "Probes_example"} // ScanOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.CreateScan(context.Background(), siteId).ScanOptions(scanOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.CreateScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScan`: Task
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.CreateScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md) | UUID or name of the site to scan | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanOptions** | [**ScanOptions**](ScanOptions.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: text/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSite

> Site CreateSite(ctx).SiteOptions(siteOptions).Execute()

Create a new site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteOptions := openapiclient.SiteOptions{Name: "Name_example", Description: "Description_example", Scope: "Scope_example", Excludes: "Excludes_example"} // SiteOptions | site definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.CreateSite(context.Background(), siteOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.CreateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.CreateSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteOptions** | [**SiteOptions**](SiteOptions.md) | site definition | 

### Return type

[**Site**](Site.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgent

> Agent GetAgent(ctx, agentId).Execute()

Get details for a single agent.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    agentId := TODO // string | UUID of the agent

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgent`: Agent
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md) | UUID of the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Agent**](Agent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> []Agent GetAgents(ctx).Execute()

Get all agents.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetAgents(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgents`: []Agent
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


### Return type

[**[]Agent**](Agent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsset

> Asset GetAsset(ctx, assetId).Execute()

Get asset details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetId := TODO // string | UUID of the asset to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetAsset(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsset`: Asset
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md) | UUID of the asset to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> []Asset GetAssets(ctx).Search(search).Execute()

Get all assets.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "search_example" // string | an optional search string for filtering results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetAssets(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssets`: []Asset
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKey

> APIKey GetKey(ctx).Execute()

Get API key details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetKey(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKey`: APIKey
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


### Return type

[**APIKey**](APIKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> Organization GetOrganization(ctx).Execute()

Get organization details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetOrganization(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetService

> Service GetService(ctx, serviceId).Execute()

Get service details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := TODO // string | UUID of the service to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: Service
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | [**string**](.md) | UUID of the service to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Service**](Service.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServices

> []Service GetServices(ctx).Search(search).Execute()

Get all services.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "search_example" // string | an optional search string for filtering results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetServices(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServices`: []Service
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Service**](Service.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSite

> Site GetSite(ctx, siteId).Execute()

Get site details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteId := TODO // string | UUID or name of the site

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetSite(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md) | UUID or name of the site | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Site**](Site.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSites

> []Site GetSites(ctx).Execute()

Get all sites.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetSites(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetSites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSites`: []Site
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesRequest struct via the builder pattern


### Return type

[**[]Site**](Site.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, taskId).Execute()

Get task details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskId := TODO // string | UUID of the task to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md) | UUID of the task to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Task**](Task.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskChangeReport

> GetTaskChangeReport(ctx, taskId).Execute()

Returns a temporary task change report data url.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskId := TODO // string | UUID of the task

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetTaskChangeReport(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetTaskChangeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md) | UUID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskChangeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskScanData

> GetTaskScanData(ctx, taskId).Execute()

Returns a temporary task scan data url.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskId := TODO // string | UUID of the task

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetTaskScanData(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetTaskScanData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md) | UUID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskScanDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []Task GetTasks(ctx).Status(status).Execute()

Get all tasks (last 1000).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    status := "status_example" // string | an optional status string for filtering results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetTasks(context.Background(), ).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: []Task
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | an optional status string for filtering results | 

### Return type

[**[]Task**](Task.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessLAN

> Wireless GetWirelessLAN(ctx, wirelessId).Execute()

Get wireless LAN details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wirelessId := TODO // string | UUID of the wireless LAN to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetWirelessLAN(context.Background(), wirelessId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetWirelessLAN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWirelessLAN`: Wireless
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetWirelessLAN`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wirelessId** | [**string**](.md) | UUID of the wireless LAN to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWirelessLANRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wireless**](Wireless.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWirelessLANs

> []Wireless GetWirelessLANs(ctx).Search(search).Execute()

Get all wireless LANs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "search_example" // string | an optional search string for filtering results (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.GetWirelessLANs(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.GetWirelessLANs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWirelessLANs`: []Wireless
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.GetWirelessLANs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWirelessLANsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[**[]Wireless**](Wireless.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideTask

> HideTask(ctx, taskId).Execute()

Signal that a completed task should be hidden.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskId := TODO // string | UUID of the task to hide

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.HideTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.HideTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md) | UUID of the task to hide | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportScanData

> Task ImportScanData(ctx, siteId).Body(body).Execute()

Import a scan data file into a site.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteId := TODO // string | UUID or name of the site to import scan data into
    body := 987 // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.ImportScanData(context.Background(), siteId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.ImportScanData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportScanData`: Task
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.ImportScanData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md) | UUID or name of the site to import scan data into | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportScanDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

[**Task**](Task.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAgent

> RemoveAgent(ctx, agentId).Execute()

Remove and uninstall an agent.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    agentId := TODO // string | UUID of the agent to remove

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.RemoveAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.RemoveAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md) | UUID of the agent to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAsset

> RemoveAsset(ctx, assetId).Execute()

Remove an asset.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetId := TODO // string | UUID of the asset to remove

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.RemoveAsset(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.RemoveAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md) | UUID of the asset to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveService

> RemoveService(ctx, serviceId).Execute()

Remove a service.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := TODO // string | UUID of the service to remove

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.RemoveService(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.RemoveService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | [**string**](.md) | UUID of the service to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSite

> RemoveSite(ctx, siteId).Execute()

Remove a site and associated assets.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteId := TODO // string | UUID or name of the site to remove

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.RemoveSite(context.Background(), siteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.RemoveSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md) | UUID or name of the site to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWirelessLAN

> RemoveWirelessLAN(ctx, wirelessId).Execute()

Remove a wireless LAN.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wirelessId := TODO // string | UUID of the wireless LAN to remove

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.RemoveWirelessLAN(context.Background(), wirelessId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.RemoveWirelessLAN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wirelessId** | [**string**](.md) | UUID of the wireless LAN to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWirelessLANRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopTask

> StopTask(ctx, taskId).Execute()

Signal that a task should be stopped or canceled.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskId := TODO // string | UUID of the task to stop

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.StopTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.StopTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md) | UUID of the task to stop | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentSite

> Agent UpdateAgentSite(ctx, agentId).AgentSiteID(agentSiteID).Execute()

Update the site associated with agent.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    agentId := TODO // string | UUID of the agent to update
    agentSiteID := openapiclient.AgentSiteID{SiteId: "SiteId_example"} // AgentSiteID | site_id to associate with the agent

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpdateAgentSite(context.Background(), agentId, agentSiteID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateAgentSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgentSite`: Agent
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateAgentSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md) | UUID of the agent to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentSiteID** | [**AgentSiteID**](AgentSiteID.md) | site_id to associate with the agent | 

### Return type

[**Agent**](Agent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetComments

> Asset UpdateAssetComments(ctx, assetId).AssetComments(assetComments).Execute()

Update asset comments.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetId := TODO // string | UUID of the asset to update
    assetComments := openapiclient.AssetComments{Comments: "Comments_example"} // AssetComments | comments to apply to the asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpdateAssetComments(context.Background(), assetId, assetComments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateAssetComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetComments`: Asset
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateAssetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md) | UUID of the asset to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetComments** | [**AssetComments**](AssetComments.md) | comments to apply to the asset | 

### Return type

[**Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetTags

> Asset UpdateAssetTags(ctx, assetId).AssetTags(assetTags).Execute()

Update asset tags.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetId := TODO // string | UUID of the agent to update
    assetTags := openapiclient.AssetTags{Tags: map[string]string{ "Key" = "Value" }} // AssetTags | tags to apply to the asset

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpdateAssetTags(context.Background(), assetId, assetTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateAssetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetTags`: Asset
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateAssetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | [**string**](.md) | UUID of the agent to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetTags** | [**AssetTags**](AssetTags.md) | tags to apply to the asset | 

### Return type

[**Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganization

> Organization UpdateOrganization(ctx).OrgOptions(orgOptions).Execute()

Update organization details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgOptions := openapiclient.OrgOptions{Name: "Name_example", Description: "Description_example"} // OrgOptions | organization options

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpdateOrganization(context.Background(), orgOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgOptions** | [**OrgOptions**](OrgOptions.md) | organization options | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSite

> Site UpdateSite(ctx, siteId).SiteOptions(siteOptions).Execute()

Update a site definition.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    siteId := TODO // string | UUID or name of the site to update
    siteOptions := openapiclient.SiteOptions{Name: "Name_example", Description: "Description_example", Scope: "Scope_example", Excludes: "Excludes_example"} // SiteOptions | site object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpdateSite(context.Background(), siteId, siteOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSite`: Site
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | [**string**](.md) | UUID or name of the site to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteOptions** | [**SiteOptions**](SiteOptions.md) | site object | 

### Return type

[**Site**](Site.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> Task UpdateTask(ctx, taskId).Task(task).Execute()

Update task parameters.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    taskId := TODO // string | UUID of the task to update
    task := openapiclient.Task{Id: "Id_example", Name: "Name_example", Description: "Description_example", ClientId: "ClientId_example", OrganizationId: "OrganizationId_example", AgentId: "AgentId_example", SiteId: "SiteId_example", CruncherId: "CruncherId_example", CreatedAt: int64(123), CreatedBy: "CreatedBy_example", CreatedByUserId: "CreatedByUserId_example", UpdatedAt: int64(123), Type: "Type_example", Status: "Status_example", Error: "Error_example", Params: map[string]string{ "Key" = "Value" }, Stats: map[string]string{ "Key" = "Value" }, Hidden: false, ParentId: "ParentId_example", Recur: false, RecurFrequency: "RecurFrequency_example", StartTime: int64(123), RecurLast: int64(123), RecurNext: int64(123), RecurLastTaskId: "RecurLastTaskId_example"} // Task | task object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpdateTask(context.Background(), taskId, task).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `OrganizationApi.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | [**string**](.md) | UUID of the task to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**Task**](Task.md) | task object | 

### Return type

[**Task**](Task.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeAgent

> UpgradeAgent(ctx, agentId).Execute()

Force an agent to update and restart.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    agentId := TODO // string | UUID of the agent to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationApi.UpgradeAgent(context.Background(), agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationApi.UpgradeAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | [**string**](.md) | UUID of the agent to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

