# \BatchWebhooksApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBatchWebhookId**](BatchWebhooksApi.md#DeleteBatchWebhookId) | **Delete** /batch-webhooks/{batch_webhook_id} | Delete batch webhook
[**GetBatchWebhook**](BatchWebhooksApi.md#GetBatchWebhook) | **Get** /batch-webhooks/{batch_webhook_id} | Get batch webhook info
[**GetBatchWebhooks**](BatchWebhooksApi.md#GetBatchWebhooks) | **Get** /batch-webhooks | List batch webhooks
[**PatchBatchWebhooks**](BatchWebhooksApi.md#PatchBatchWebhooks) | **Patch** /batch-webhooks/{batch_webhook_id} | Update batch webhook
[**PostBatchWebhooks**](BatchWebhooksApi.md#PostBatchWebhooks) | **Post** /batch-webhooks | Add batch webhook


# **DeleteBatchWebhookId**
> DeleteBatchWebhookId(ctx, batchWebhookId)
Delete batch webhook

Remove a batch webhook. Webhooks will no longer be sent to the given URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchWebhookId** | **string**| The unique id for the batch webhook. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchWebhook**
> BatchWebhook GetBatchWebhook(ctx, batchWebhookId, optional)
Get batch webhook info

Get information about a specific batch webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchWebhookId** | **string**| The unique id for the batch webhook. | 
 **optional** | ***BatchWebhooksApiGetBatchWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchWebhooksApiGetBatchWebhookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**BatchWebhook**](Batch Webhook.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchWebhooks**
> BatchWebhooks GetBatchWebhooks(ctx, optional)
List batch webhooks

Get all webhooks that have been configured for batches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BatchWebhooksApiGetBatchWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchWebhooksApiGetBatchWebhooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**BatchWebhooks**](Batch Webhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchBatchWebhooks**
> BatchWebhook PatchBatchWebhooks(ctx, batchWebhookId, body)
Update batch webhook

Update a webhook that will fire whenever any batch request completes processing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchWebhookId** | **string**| The unique id for the batch webhook. | 
  **body** | [**BatchWebhook2**](BatchWebhook2.md)|  | 

### Return type

[**BatchWebhook**](Batch Webhook.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostBatchWebhooks**
> BatchWebhook PostBatchWebhooks(ctx, body)
Add batch webhook

Configure a webhook that will fire whenever any batch request completes processing.  You may only have a maximum of 20 batch webhooks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchWebhook1**](BatchWebhook1.md)|  | 

### Return type

[**BatchWebhook**](Batch Webhook.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

