# \FileManagerApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFileManagerFilesId**](FileManagerApi.md#DeleteFileManagerFilesId) | **Delete** /file-manager/files/{file_id} | Delete file
[**DeleteFileManagerFoldersId**](FileManagerApi.md#DeleteFileManagerFoldersId) | **Delete** /file-manager/folders/{folder_id} | Delete folder
[**GetFileManagerFiles**](FileManagerApi.md#GetFileManagerFiles) | **Get** /file-manager/files | List stored files
[**GetFileManagerFilesId**](FileManagerApi.md#GetFileManagerFilesId) | **Get** /file-manager/files/{file_id} | Get file
[**GetFileManagerFolders**](FileManagerApi.md#GetFileManagerFolders) | **Get** /file-manager/folders | List folders
[**GetFileManagerFoldersFiles**](FileManagerApi.md#GetFileManagerFoldersFiles) | **Get** /file-manager/folders/{folder_id}/files | List stored files
[**GetFileManagerFoldersId**](FileManagerApi.md#GetFileManagerFoldersId) | **Get** /file-manager/folders/{folder_id} | Get folder
[**PatchFileManagerFilesId**](FileManagerApi.md#PatchFileManagerFilesId) | **Patch** /file-manager/files/{file_id} | Update file
[**PatchFileManagerFoldersId**](FileManagerApi.md#PatchFileManagerFoldersId) | **Patch** /file-manager/folders/{folder_id} | Update folder
[**PostFileManagerFiles**](FileManagerApi.md#PostFileManagerFiles) | **Post** /file-manager/files | Add file
[**PostFileManagerFolders**](FileManagerApi.md#PostFileManagerFolders) | **Post** /file-manager/folders | Add folder


# **DeleteFileManagerFilesId**
> DeleteFileManagerFilesId(ctx, fileId)
Delete file

Remove a specific file from the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| The unique id for the File Manager file. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFileManagerFoldersId**
> DeleteFileManagerFoldersId(ctx, folderId)
Delete folder

Delete a specific folder in the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the File Manager folder. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileManagerFiles**
> FileManager GetFileManagerFiles(ctx, optional)
List stored files

Get a list of available images and files stored in the File Manager for the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileManagerApiGetFileManagerFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileManagerApiGetFileManagerFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| The file type for the File Manager file. | 
 **createdBy** | **optional.String**| The Mailchimp account user who created the File Manager file. | 
 **beforeCreatedAt** | **optional.String**| Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreatedAt** | **optional.String**| Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 

### Return type

[**FileManager**](File Manager.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileManagerFilesId**
> GalleryFile GetFileManagerFilesId(ctx, fileId, optional)
Get file

Get information about a specific file in the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| The unique id for the File Manager file. | 
 **optional** | ***FileManagerApiGetFileManagerFilesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileManagerApiGetFileManagerFilesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GalleryFile**](Gallery File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileManagerFolders**
> FileManagerFolders GetFileManagerFolders(ctx, optional)
List folders

Get a list of all folders in the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FileManagerApiGetFileManagerFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileManagerApiGetFileManagerFoldersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **createdBy** | **optional.String**| The Mailchimp account user who created the File Manager file. | 
 **beforeCreatedAt** | **optional.String**| Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreatedAt** | **optional.String**| Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**FileManagerFolders**](File Manager Folders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileManagerFoldersFiles**
> FileManager GetFileManagerFoldersFiles(ctx, folderId, optional)
List stored files

Get a list of available images and files stored in this folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the File Manager folder. | 
 **optional** | ***FileManagerApiGetFileManagerFoldersFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileManagerApiGetFileManagerFoldersFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| The file type for the File Manager file. | 
 **createdBy** | **optional.String**| The Mailchimp account user who created the File Manager file. | 
 **beforeCreatedAt** | **optional.String**| Restrict the response to files created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreatedAt** | **optional.String**| Restrict the response to files created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 

### Return type

[**FileManager**](File Manager.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileManagerFoldersId**
> GalleryFolder GetFileManagerFoldersId(ctx, folderId, optional)
Get folder

Get information about a specific folder in the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the File Manager folder. | 
 **optional** | ***FileManagerApiGetFileManagerFoldersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileManagerApiGetFileManagerFoldersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GalleryFolder**](Gallery Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFileManagerFilesId**
> GalleryFile PatchFileManagerFilesId(ctx, fileId, body)
Update file

Update a file in the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| The unique id for the File Manager file. | 
  **body** | [**GalleryFile2**](GalleryFile2.md)|  | 

### Return type

[**GalleryFile**](Gallery File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFileManagerFoldersId**
> GalleryFolder PatchFileManagerFoldersId(ctx, folderId, body)
Update folder

Update a specific File Manager folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **string**| The unique id for the File Manager folder. | 
  **body** | [**GalleryFolder2**](GalleryFolder2.md)|  | 

### Return type

[**GalleryFolder**](Gallery Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFileManagerFiles**
> GalleryFile PostFileManagerFiles(ctx, body)
Add file

Upload a new image or file to the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GalleryFile1**](GalleryFile1.md)|  | 

### Return type

[**GalleryFile**](Gallery File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFileManagerFolders**
> GalleryFolder PostFileManagerFolders(ctx, body)
Add folder

Create a new folder in the File Manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GalleryFolder1**](GalleryFolder1.md)|  | 

### Return type

[**GalleryFolder**](Gallery Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

