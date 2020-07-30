# \PublicApi

All URIs are relative to *https://console.rumble.run/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestAgentVersion**](PublicApi.md#GetLatestAgentVersion) | **Get** /releases/agent/version | Returns latest agent version.
[**GetLatestPlatformVersion**](PublicApi.md#GetLatestPlatformVersion) | **Get** /releases/platform/version | Returns latest platform version.
[**GetLatestScannerVersion**](PublicApi.md#GetLatestScannerVersion) | **Get** /releases/scanner/version | Returns latest scanner version.



## GetLatestAgentVersion

> ComponentVersion GetLatestAgentVersion(ctx).Execute()

Returns latest agent version.

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
    resp, r, err := api_client.PublicApi.GetLatestAgentVersion(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetLatestAgentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestAgentVersion`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetLatestAgentVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestAgentVersionRequest struct via the builder pattern


### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestPlatformVersion

> ComponentVersion GetLatestPlatformVersion(ctx).Execute()

Returns latest platform version.

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
    resp, r, err := api_client.PublicApi.GetLatestPlatformVersion(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetLatestPlatformVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestPlatformVersion`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetLatestPlatformVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestPlatformVersionRequest struct via the builder pattern


### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestScannerVersion

> ComponentVersion GetLatestScannerVersion(ctx).Execute()

Returns latest scanner version.

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
    resp, r, err := api_client.PublicApi.GetLatestScannerVersion(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicApi.GetLatestScannerVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestScannerVersion`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `PublicApi.GetLatestScannerVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestScannerVersionRequest struct via the builder pattern


### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

