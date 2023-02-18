# \SearchMembersApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchMembers**](SearchMembersApi.md#GetSearchMembers) | **Get** /search-members | Search members


# **GetSearchMembers**
> Members GetSearchMembers(ctx, query, optional)
Search members

Search for list members. This search can be restricted to a specific list, or can be used to search across all lists in an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| The search query used to filter results. Query should be a valid email, or a string representing a contact&#39;s first or last name. | 
 **optional** | ***SearchMembersApiGetSearchMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchMembersApiGetSearchMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **listId** | **optional.String**| The unique id for the list. | 

### Return type

[**Members**](Members.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

