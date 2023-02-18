# \SearchCampaignsApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchCampaigns**](SearchCampaignsApi.md#GetSearchCampaigns) | **Get** /search-campaigns | Search campaigns


# **GetSearchCampaigns**
> Campaigns GetSearchCampaigns(ctx, query, optional)
Search campaigns

Search all campaigns for the specified query terms.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| The search query used to filter results. | 
 **optional** | ***SearchCampaignsApiGetSearchCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchCampaignsApiGetSearchCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Campaigns**](Campaigns.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

