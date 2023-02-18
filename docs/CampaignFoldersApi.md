# \CampaignFoldersApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCampaignFoldersId**](CampaignFoldersApi.md#DeleteCampaignFoldersId) | **Delete** /campaign-folders/{folder_id} | Delete campaign folder
[**GetCampaignFolders**](CampaignFoldersApi.md#GetCampaignFolders) | **Get** /campaign-folders | List campaign folders
[**GetCampaignFoldersId**](CampaignFoldersApi.md#GetCampaignFoldersId) | **Get** /campaign-folders/{folder_id} | Get campaign folder
[**PatchCampaignFoldersId**](CampaignFoldersApi.md#PatchCampaignFoldersId) | **Patch** /campaign-folders/{folder_id} | Update campaign folder
[**PostCampaignFolders**](CampaignFoldersApi.md#PostCampaignFolders) | **Post** /campaign-folders | Add campaign folder


# **DeleteCampaignFoldersId**
> DeleteCampaignFoldersId(ctx, folderId)
Delete campaign folder

Delete a specific campaign folder, and mark all the campaigns in the folder as 'unfiled'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the campaign folder. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignFolders**
> CampaignFolders GetCampaignFolders(ctx, optional)
List campaign folders

Get all folders used to organize campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CampaignFoldersApiGetCampaignFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignFoldersApiGetCampaignFoldersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CampaignFolders**](Campaign Folders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignFoldersId**
> CampaignFolder GetCampaignFoldersId(ctx, folderId, optional)
Get campaign folder

Get information about a specific folder used to organize campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the campaign folder. | 
 **optional** | ***CampaignFoldersApiGetCampaignFoldersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignFoldersApiGetCampaignFoldersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignFolder**](Campaign Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCampaignFoldersId**
> CampaignFolder PatchCampaignFoldersId(ctx, folderId, body)
Update campaign folder

Update a specific folder used to organize campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the campaign folder. | 
  **body** | [**CampaignFolder2**](CampaignFolder2.md)|  | 

### Return type

[**CampaignFolder**](Campaign Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignFolders**
> CampaignFolder PostCampaignFolders(ctx, body)
Add campaign folder

Create a new campaign folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CampaignFolder1**](CampaignFolder1.md)|  | 

### Return type

[**CampaignFolder**](Campaign Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

