# \ExportApi

All URIs are relative to *https://console.rumble.run/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAssetSyncCreatedJSON**](ExportApi.md#ExportAssetSyncCreatedJSON) | **Get** /export/org/assets/sync/created/assets.json | Exports the asset inventory in a sync-friendly manner using created_at as a checkpoint.
[**ExportAssetSyncUpdatedJSON**](ExportApi.md#ExportAssetSyncUpdatedJSON) | **Get** /export/org/assets/sync/updated/assets.json | Exports the asset inventory in a sync-friendly manner using updated_at as a checkpoint.
[**ExportAssetsCSV**](ExportApi.md#ExportAssetsCSV) | **Get** /export/org/assets.csv | Asset inventory as CSV.
[**ExportAssetsJSON**](ExportApi.md#ExportAssetsJSON) | **Get** /export/org/assets.json | Exports the asset inventory.
[**ExportAssetsJSONL**](ExportApi.md#ExportAssetsJSONL) | **Get** /export/org/assets.jsonl | Asset inventory as JSON line-delimited.
[**ExportAssetsNmapXML**](ExportApi.md#ExportAssetsNmapXML) | **Get** /export/org/assets.nmap.xml | Asset inventory as Nmap-style XML.
[**ExportServicesCSV**](ExportApi.md#ExportServicesCSV) | **Get** /export/org/services.csv | Service inventory as CSV.
[**ExportServicesJSON**](ExportApi.md#ExportServicesJSON) | **Get** /export/org/services.json | Service inventory as JSON.
[**ExportServicesJSONL**](ExportApi.md#ExportServicesJSONL) | **Get** /export/org/services.jsonl | Service inventory as JSON line-delimited.
[**ExportSitesCSV**](ExportApi.md#ExportSitesCSV) | **Get** /export/org/sites.csv | Site list as CSV.
[**ExportSitesJSON**](ExportApi.md#ExportSitesJSON) | **Get** /export/org/sites.json | Export all sites.
[**ExportSitesJSONL**](ExportApi.md#ExportSitesJSONL) | **Get** /export/org/sites.jsonl | Site list as JSON line-delimited.
[**ExportWirelessCSV**](ExportApi.md#ExportWirelessCSV) | **Get** /export/org/wireless.csv | Wireless inventory as CSV.
[**ExportWirelessJSON**](ExportApi.md#ExportWirelessJSON) | **Get** /export/org/wireless.json | Wireless inventory as JSON.
[**ExportWirelessJSONL**](ExportApi.md#ExportWirelessJSONL) | **Get** /export/org/wireless.jsonl | Wireless inventory as JSON line-delimited.



## ExportAssetSyncCreatedJSON

> AssetsWithCheckpoint ExportAssetSyncCreatedJSON(ctx).Search(search).Fields(fields).Since(since).Execute()

Exports the asset inventory in a sync-friendly manner using created_at as a checkpoint.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)
    since := 987 // int64 | an optional unix timestamp to use as a checkpoint (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportAssetSyncCreatedJSON(context.Background(), ).Search(search).Fields(fields).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportAssetSyncCreatedJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAssetSyncCreatedJSON`: AssetsWithCheckpoint
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportAssetSyncCreatedJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetSyncCreatedJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 
 **since** | **int64** | an optional unix timestamp to use as a checkpoint | 

### Return type

[**AssetsWithCheckpoint**](AssetsWithCheckpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetSyncUpdatedJSON

> AssetsWithCheckpoint ExportAssetSyncUpdatedJSON(ctx).Search(search).Fields(fields).Since(since).Execute()

Exports the asset inventory in a sync-friendly manner using updated_at as a checkpoint.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)
    since := 987 // int64 | an optional unix timestamp to use as a checkpoint (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportAssetSyncUpdatedJSON(context.Background(), ).Search(search).Fields(fields).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportAssetSyncUpdatedJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAssetSyncUpdatedJSON`: AssetsWithCheckpoint
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportAssetSyncUpdatedJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetSyncUpdatedJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 
 **since** | **int64** | an optional unix timestamp to use as a checkpoint | 

### Return type

[**AssetsWithCheckpoint**](AssetsWithCheckpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsCSV

> *os.File ExportAssetsCSV(ctx).Search(search).Execute()

Asset inventory as CSV.

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
    resp, r, err := api_client.ExportApi.ExportAssetsCSV(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportAssetsCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAssetsCSV`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportAssetsCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsJSON

> []Asset ExportAssetsJSON(ctx).Search(search).Fields(fields).Execute()

Exports the asset inventory.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportAssetsJSON(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportAssetsJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAssetsJSON`: []Asset
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportAssetsJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

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


## ExportAssetsJSONL

> *os.File ExportAssetsJSONL(ctx).Search(search).Fields(fields).Execute()

Asset inventory as JSON line-delimited.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportAssetsJSONL(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportAssetsJSONL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAssetsJSONL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportAssetsJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsNmapXML

> *os.File ExportAssetsNmapXML(ctx).Search(search).Execute()

Asset inventory as Nmap-style XML.

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
    resp, r, err := api_client.ExportApi.ExportAssetsNmapXML(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportAssetsNmapXML``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAssetsNmapXML`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportAssetsNmapXML`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAssetsNmapXMLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServicesCSV

> *os.File ExportServicesCSV(ctx).Search(search).Execute()

Service inventory as CSV.

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
    resp, r, err := api_client.ExportApi.ExportServicesCSV(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportServicesCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportServicesCSV`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportServicesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServicesJSON

> []Service ExportServicesJSON(ctx).Search(search).Fields(fields).Execute()

Service inventory as JSON.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportServicesJSON(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportServicesJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportServicesJSON`: []Service
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportServicesJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

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


## ExportServicesJSONL

> *os.File ExportServicesJSONL(ctx).Search(search).Fields(fields).Execute()

Service inventory as JSON line-delimited.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportServicesJSONL(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportServicesJSONL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportServicesJSONL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportServicesJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportServicesJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSitesCSV

> *os.File ExportSitesCSV(ctx).Execute()

Site list as CSV.

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
    resp, r, err := api_client.ExportApi.ExportSitesCSV(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportSitesCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSitesCSV`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportSitesCSV`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportSitesCSVRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSitesJSON

> []Site ExportSitesJSON(ctx).Search(search).Fields(fields).Execute()

Export all sites.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportSitesJSON(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportSitesJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSitesJSON`: []Site
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportSitesJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSitesJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

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


## ExportSitesJSONL

> *os.File ExportSitesJSONL(ctx).Search(search).Fields(fields).Execute()

Site list as JSON line-delimited.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportSitesJSONL(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportSitesJSONL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSitesJSONL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportSitesJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSitesJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWirelessCSV

> *os.File ExportWirelessCSV(ctx).Search(search).Execute()

Wireless inventory as CSV.

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
    resp, r, err := api_client.ExportApi.ExportWirelessCSV(context.Background(), ).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportWirelessCSV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportWirelessCSV`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportWirelessCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWirelessCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWirelessJSON

> []Wireless ExportWirelessJSON(ctx).Search(search).Fields(fields).Execute()

Wireless inventory as JSON.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportWirelessJSON(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportWirelessJSON``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportWirelessJSON`: []Wireless
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportWirelessJSON`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWirelessJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

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


## ExportWirelessJSONL

> *os.File ExportWirelessJSONL(ctx).Search(search).Fields(fields).Execute()

Wireless inventory as JSON line-delimited.

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
    fields := "fields_example" // string | an optional list of fields to export, comma-separated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportApi.ExportWirelessJSONL(context.Background(), ).Search(search).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportApi.ExportWirelessJSONL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportWirelessJSONL`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ExportApi.ExportWirelessJSONL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWirelessJSONLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | an optional search string for filtering results | 
 **fields** | **string** | an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

