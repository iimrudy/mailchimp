# \ActivityFeedApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityFeedChimpChatter**](ActivityFeedApi.md#GetActivityFeedChimpChatter) | **Get** /activity-feed/chimp-chatter | Get latest chimp chatter


# **GetActivityFeedChimpChatter**
> InlineResponse200 GetActivityFeedChimpChatter(ctx, optional)
Get latest chimp chatter

Return the Chimp Chatter for this account ordered by most recent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActivityFeedApiGetActivityFeedChimpChatterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivityFeedApiGetActivityFeedChimpChatterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

