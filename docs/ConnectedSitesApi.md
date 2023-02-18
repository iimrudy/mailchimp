# \ConnectedSitesApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnectedSitesId**](ConnectedSitesApi.md#DeleteConnectedSitesId) | **Delete** /connected-sites/{connected_site_id} | Delete connected site
[**GetConnectedSites**](ConnectedSitesApi.md#GetConnectedSites) | **Get** /connected-sites | List connected sites
[**GetConnectedSitesId**](ConnectedSitesApi.md#GetConnectedSitesId) | **Get** /connected-sites/{connected_site_id} | Get connected site
[**PostConnectedSites**](ConnectedSitesApi.md#PostConnectedSites) | **Post** /connected-sites | Add connected site
[**PostConnectedSitesIdActionsVerifyScriptInstallation**](ConnectedSitesApi.md#PostConnectedSitesIdActionsVerifyScriptInstallation) | **Post** /connected-sites/{connected_site_id}/actions/verify-script-installation | Verify connected site script


# **DeleteConnectedSitesId**
> DeleteConnectedSitesId(ctx, connectedSiteId)
Delete connected site

Remove a connected site from your Mailchimp account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectedSiteId** | **string**| The unique identifier for the site. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectedSites**
> ConnectedSites GetConnectedSites(ctx, optional)
List connected sites

Get all connected sites in an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectedSitesApiGetConnectedSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectedSitesApiGetConnectedSitesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**ConnectedSites**](Connected Sites.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectedSitesId**
> ConnectedSite GetConnectedSitesId(ctx, connectedSiteId, optional)
Get connected site

Get information about a specific connected site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectedSiteId** | **string**| The unique identifier for the site. | 
 **optional** | ***ConnectedSitesApiGetConnectedSitesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectedSitesApiGetConnectedSitesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ConnectedSite**](Connected Site.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostConnectedSites**
> ConnectedSite PostConnectedSites(ctx, body)
Add connected site

Create a new Mailchimp connected site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectedSite1**](ConnectedSite1.md)|  | 

### Return type

[**ConnectedSite**](Connected Site.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostConnectedSitesIdActionsVerifyScriptInstallation**
> PostConnectedSitesIdActionsVerifyScriptInstallation(ctx, connectedSiteId)
Verify connected site script

Verify that the connected sites script has been installed, either via the script URL or fragment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectedSiteId** | **string**| The unique identifier for the site. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

