# \TemplateFoldersApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplateFoldersId**](TemplateFoldersApi.md#DeleteTemplateFoldersId) | **Delete** /template-folders/{folder_id} | Delete template folder
[**GetTemplateFolders**](TemplateFoldersApi.md#GetTemplateFolders) | **Get** /template-folders | List template folders
[**GetTemplateFoldersId**](TemplateFoldersApi.md#GetTemplateFoldersId) | **Get** /template-folders/{folder_id} | Get template folder
[**PatchTemplateFoldersId**](TemplateFoldersApi.md#PatchTemplateFoldersId) | **Patch** /template-folders/{folder_id} | Update template folder
[**PostTemplateFolders**](TemplateFoldersApi.md#PostTemplateFolders) | **Post** /template-folders | Add template folder


# **DeleteTemplateFoldersId**
> DeleteTemplateFoldersId(ctx, folderId)
Delete template folder

Delete a specific template folder, and mark all the templates in the folder as 'unfiled'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the template folder. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateFolders**
> TemplateFolders GetTemplateFolders(ctx, optional)
List template folders

Get all folders used to organize templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TemplateFoldersApiGetTemplateFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplateFoldersApiGetTemplateFoldersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**TemplateFolders**](Template Folders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplateFoldersId**
> TemplateFolder GetTemplateFoldersId(ctx, folderId, optional)
Get template folder

Get information about a specific folder used to organize templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the template folder. | 
 **optional** | ***TemplateFoldersApiGetTemplateFoldersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplateFoldersApiGetTemplateFoldersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**TemplateFolder**](Template Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchTemplateFoldersId**
> TemplateFolder PatchTemplateFoldersId(ctx, folderId, body)
Update template folder

Update a specific folder used to organize templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the template folder. | 
  **body** | [**TemplateFolder2**](TemplateFolder2.md)|  | 

### Return type

[**TemplateFolder**](Template Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTemplateFolders**
> TemplateFolder PostTemplateFolders(ctx, body)
Add template folder

Create a new template folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TemplateFolder1**](TemplateFolder1.md)|  | 

### Return type

[**TemplateFolder**](Template Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

