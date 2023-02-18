# \PingApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPing**](PingApi.md#GetPing) | **Get** /ping | Ping


# **GetPing**
> ApiHealthStatus GetPing(ctx, )
Ping

A health check for the API that won't return any account-specific information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiHealthStatus**](API health status.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

