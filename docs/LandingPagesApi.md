# \LandingPagesApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLandingPageId**](LandingPagesApi.md#DeleteLandingPageId) | **Delete** /landing-pages/{page_id} | Delete landing page
[**GetAllLandingPages**](LandingPagesApi.md#GetAllLandingPages) | **Get** /landing-pages | List landing pages
[**GetLandingPageId**](LandingPagesApi.md#GetLandingPageId) | **Get** /landing-pages/{page_id} | Get landing page info
[**GetLandingPageIdContent**](LandingPagesApi.md#GetLandingPageIdContent) | **Get** /landing-pages/{page_id}/content | Get landing page content
[**PatchLandingPageId**](LandingPagesApi.md#PatchLandingPageId) | **Patch** /landing-pages/{page_id} | Update landing page
[**PostAllLandingPages**](LandingPagesApi.md#PostAllLandingPages) | **Post** /landing-pages | Add landing page
[**PostLandingPageIdActionsPublish**](LandingPagesApi.md#PostLandingPageIdActionsPublish) | **Post** /landing-pages/{page_id}/actions/publish | Publish landing page
[**PostLandingPageIdActionsUnpublish**](LandingPagesApi.md#PostLandingPageIdActionsUnpublish) | **Post** /landing-pages/{page_id}/actions/unpublish | Unpublish landing page


# **DeleteLandingPageId**
> DeleteLandingPageId(ctx, pageId)
Delete landing page

Delete a landing page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| The unique id for the page. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllLandingPages**
> InlineResponse2006 GetAllLandingPages(ctx, optional)
List landing pages

Get all landing pages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LandingPagesApiGetAllLandingPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LandingPagesApiGetAllLandingPagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLandingPageId**
> LandingPage GetLandingPageId(ctx, pageId, optional)
Get landing page info

Get information about a specific page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| The unique id for the page. | 
 **optional** | ***LandingPagesApiGetLandingPageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LandingPagesApiGetLandingPageIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**LandingPage**](Landing Page.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLandingPageIdContent**
> LandingPageContent GetLandingPageIdContent(ctx, pageId, optional)
Get landing page content

Get the the HTML for your landing page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| The unique id for the page. | 
 **optional** | ***LandingPagesApiGetLandingPageIdContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LandingPagesApiGetLandingPageIdContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**LandingPageContent**](Landing Page Content.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchLandingPageId**
> LandingPage PatchLandingPageId(ctx, pageId, body)
Update landing page

Update a landing page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| The unique id for the page. | 
  **body** | [**LandingPage2**](LandingPage2.md)|  | 

### Return type

[**LandingPage**](Landing Page.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAllLandingPages**
> LandingPage PostAllLandingPages(ctx, body, optional)
Add landing page

Create a new Mailchimp landing page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LandingPage1**](LandingPage1.md)|  | 
 **optional** | ***LandingPagesApiPostAllLandingPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LandingPagesApiPostAllLandingPagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useDefaultList** | **optional.Bool**| Will create the Landing Page using the account&#39;s Default List instead of requiring a list_id. | 

### Return type

[**LandingPage**](Landing Page.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLandingPageIdActionsPublish**
> PostLandingPageIdActionsPublish(ctx, pageId)
Publish landing page

Publish a landing page that is in draft, unpublished, or has been previously published and edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| The unique id for the page. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLandingPageIdActionsUnpublish**
> PostLandingPageIdActionsUnpublish(ctx, pageId)
Unpublish landing page

Unpublish a landing page that is in draft or has been published.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| The unique id for the page. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

