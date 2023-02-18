# \TemplatesApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplatesId**](TemplatesApi.md#DeleteTemplatesId) | **Delete** /templates/{template_id} | Delete template
[**GetTemplates**](TemplatesApi.md#GetTemplates) | **Get** /templates | List templates
[**GetTemplatesId**](TemplatesApi.md#GetTemplatesId) | **Get** /templates/{template_id} | Get template info
[**GetTemplatesIdDefaultContent**](TemplatesApi.md#GetTemplatesIdDefaultContent) | **Get** /templates/{template_id}/default-content | View default content
[**PatchTemplatesId**](TemplatesApi.md#PatchTemplatesId) | **Patch** /templates/{template_id} | Update template
[**PostTemplates**](TemplatesApi.md#PostTemplates) | **Post** /templates | Add template


# **DeleteTemplatesId**
> DeleteTemplatesId(ctx, templateId)
Delete template

Delete a specific template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| The unique id for the template. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplates**
> Templates GetTemplates(ctx, optional)
List templates

Get a list of an account's available templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TemplatesApiGetTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplatesApiGetTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **createdBy** | **optional.String**| The Mailchimp account user who created the template. | 
 **sinceDateCreated** | **optional.String**| Restrict the response to templates created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeDateCreated** | **optional.String**| Restrict the response to templates created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **type_** | **optional.String**| Limit results based on template type. | 
 **category** | **optional.String**| Limit results based on category. | 
 **folderId** | **optional.String**| The unique folder id. | 
 **sortField** | **optional.String**| Returns user templates sorted by the specified field. | 
 **contentType** | **optional.String**| Limit results based on how the template&#39;s content is put together. Only templates of type &#x60;user&#x60; can be filtered by &#x60;content_type&#x60;. If you want to retrieve saved templates created with the legacy email editor, then filter &#x60;content_type&#x60; to &#x60;template&#x60;. If you&#39;d rather pull your saved templates for the new editor, filter to &#x60;multichannel&#x60;. For code your own templates, filter to &#x60;html&#x60;. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 

### Return type

[**Templates**](Templates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplatesId**
> TemplateInstance GetTemplatesId(ctx, templateId, optional)
Get template info

Get information about a specific template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| The unique id for the template. | 
 **optional** | ***TemplatesApiGetTemplatesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplatesApiGetTemplatesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**TemplateInstance**](Template Instance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplatesIdDefaultContent**
> TemplateDefaultContent GetTemplatesIdDefaultContent(ctx, templateId, optional)
View default content

Get the sections that you can edit in a template, including each section's default content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| The unique id for the template. | 
 **optional** | ***TemplatesApiGetTemplatesIdDefaultContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplatesApiGetTemplatesIdDefaultContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**TemplateDefaultContent**](Template Default Content.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTemplatesId**
> TemplateInstance PatchTemplatesId(ctx, templateId, body)
Update template

Update the name, HTML, or `folder_id` of an existing template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| The unique id for the template. | 
  **body** | [**TemplateInstance2**](TemplateInstance2.md)|  | 

### Return type

[**TemplateInstance**](Template Instance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplates**
> TemplateInstance PostTemplates(ctx, body)
Add template

Create a new template for the account. Only Classic templates are supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TemplateInstance1**](TemplateInstance1.md)|  | 

### Return type

[**TemplateInstance**](Template Instance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

