# \VerifiedDomainsApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerifiedDomain**](VerifiedDomainsApi.md#CreateVerifiedDomain) | **Post** /verified-domains | Add domain to account
[**DeleteVerifiedDomain**](VerifiedDomainsApi.md#DeleteVerifiedDomain) | **Delete** /verified-domains/{domain_name} | Delete domain
[**GetVerifiedDomain**](VerifiedDomainsApi.md#GetVerifiedDomain) | **Get** /verified-domains/{domain_name} | Get domain info
[**GetVerifiedDomains**](VerifiedDomainsApi.md#GetVerifiedDomains) | **Get** /verified-domains | List sending domains
[**VerifyDomain**](VerifiedDomainsApi.md#VerifyDomain) | **Post** /verified-domains/{domain_name}/actions/verify | Verify domain


# **CreateVerifiedDomain**
> VerifiedDomains CreateVerifiedDomain(ctx, body)
Add domain to account

Add a domain to the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifiedDomains2**](VerifiedDomains2.md)|  | 

### Return type

[**VerifiedDomains**](Verified Domains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVerifiedDomain**
> DeleteVerifiedDomain(ctx, domainName)
Delete domain

Delete a verified domain from the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainName** | **string**| The domain name. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerifiedDomain**
> VerifiedDomains GetVerifiedDomain(ctx, domainName)
Get domain info

Get the details for a single domain on the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainName** | **string**| The domain name. | 

### Return type

[**VerifiedDomains**](Verified Domains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerifiedDomains**
> VerifiedDomains1 GetVerifiedDomains(ctx, )
List sending domains

Get all of the sending domains on the account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VerifiedDomains1**](Verified Domains_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyDomain**
> VerifiedDomains VerifyDomain(ctx, domainName, body)
Verify domain

Verify a domain for sending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainName** | **string**| The domain name. | 
  **body** | [**VerifyADomainForSending_**](VerifyADomainForSending_.md)|  | 

### Return type

[**VerifiedDomains**](Verified Domains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

