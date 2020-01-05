# \ExportApi

All URIs are relative to *https://console.rumble.run/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
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



## ExportAssetsCSV

> *os.File ExportAssetsCSV(ctx, optional)

Asset inventory as CSV.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportAssetsCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportAssetsCSVOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsJSON

> []Asset ExportAssetsJSON(ctx, optional)

Exports the asset inventory.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportAssetsJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportAssetsJSONOpts struct


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


## ExportAssetsJSONL

> *os.File ExportAssetsJSONL(ctx, optional)

Asset inventory as JSON line-delimited.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportAssetsJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportAssetsJSONLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAssetsNmapXML

> *os.File ExportAssetsNmapXML(ctx, optional)

Asset inventory as Nmap-style XML.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportAssetsNmapXMLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportAssetsNmapXMLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServicesCSV

> *os.File ExportServicesCSV(ctx, optional)

Service inventory as CSV.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportServicesCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportServicesCSVOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportServicesJSON

> []Service ExportServicesJSON(ctx, optional)

Service inventory as JSON.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportServicesJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportServicesJSONOpts struct


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


## ExportServicesJSONL

> *os.File ExportServicesJSONL(ctx, optional)

Service inventory as JSON line-delimited.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportServicesJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportServicesJSONLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSitesCSV

> *os.File ExportSitesCSV(ctx, )

Site list as CSV.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSitesJSON

> []Site ExportSitesJSON(ctx, )

Export all sites.

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


## ExportSitesJSONL

> *os.File ExportSitesJSONL(ctx, )

Site list as JSON line-delimited.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWirelessCSV

> *os.File ExportWirelessCSV(ctx, optional)

Wireless inventory as CSV.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportWirelessCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportWirelessCSVOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWirelessJSON

> []Wireless ExportWirelessJSON(ctx, optional)

Wireless inventory as JSON.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportWirelessJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportWirelessJSONOpts struct


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


## ExportWirelessJSONL

> *os.File ExportWirelessJSONL(ctx, optional)

Wireless inventory as JSON line-delimited.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportWirelessJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportWirelessJSONLOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

