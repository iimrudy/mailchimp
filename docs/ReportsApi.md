# \ReportsApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReports**](ReportsApi.md#GetReports) | **Get** /reports | List campaign reports
[**GetReportsId**](ReportsApi.md#GetReportsId) | **Get** /reports/{campaign_id} | Get campaign report
[**GetReportsIdAbuseReportsId**](ReportsApi.md#GetReportsIdAbuseReportsId) | **Get** /reports/{campaign_id}/abuse-reports | List abuse reports
[**GetReportsIdAbuseReportsIdId**](ReportsApi.md#GetReportsIdAbuseReportsIdId) | **Get** /reports/{campaign_id}/abuse-reports/{report_id} | Get abuse report
[**GetReportsIdAdvice**](ReportsApi.md#GetReportsIdAdvice) | **Get** /reports/{campaign_id}/advice | List campaign feedback
[**GetReportsIdClickDetails**](ReportsApi.md#GetReportsIdClickDetails) | **Get** /reports/{campaign_id}/click-details | List campaign details
[**GetReportsIdClickDetailsId**](ReportsApi.md#GetReportsIdClickDetailsId) | **Get** /reports/{campaign_id}/click-details/{link_id} | Get campaign link details
[**GetReportsIdClickDetailsIdMembers**](ReportsApi.md#GetReportsIdClickDetailsIdMembers) | **Get** /reports/{campaign_id}/click-details/{link_id}/members | List clicked link subscribers
[**GetReportsIdClickDetailsIdMembersId**](ReportsApi.md#GetReportsIdClickDetailsIdMembersId) | **Get** /reports/{campaign_id}/click-details/{link_id}/members/{subscriber_hash} | Get clicked link subscriber
[**GetReportsIdDomainPerformance**](ReportsApi.md#GetReportsIdDomainPerformance) | **Get** /reports/{campaign_id}/domain-performance | List domain performance stats
[**GetReportsIdEcommerceProductActivity**](ReportsApi.md#GetReportsIdEcommerceProductActivity) | **Get** /reports/{campaign_id}/ecommerce-product-activity | List campaign product activity
[**GetReportsIdEepurl**](ReportsApi.md#GetReportsIdEepurl) | **Get** /reports/{campaign_id}/eepurl | List EepURL activity
[**GetReportsIdEmailActivity**](ReportsApi.md#GetReportsIdEmailActivity) | **Get** /reports/{campaign_id}/email-activity | List email activity
[**GetReportsIdEmailActivityId**](ReportsApi.md#GetReportsIdEmailActivityId) | **Get** /reports/{campaign_id}/email-activity/{subscriber_hash} | Get subscriber email activity
[**GetReportsIdLocations**](ReportsApi.md#GetReportsIdLocations) | **Get** /reports/{campaign_id}/locations | List top open activities
[**GetReportsIdOpenDetails**](ReportsApi.md#GetReportsIdOpenDetails) | **Get** /reports/{campaign_id}/open-details | List campaign open details
[**GetReportsIdOpenDetailsIdMembersId**](ReportsApi.md#GetReportsIdOpenDetailsIdMembersId) | **Get** /reports/{campaign_id}/open-details/{subscriber_hash} | Get opened campaign subscriber
[**GetReportsIdSentTo**](ReportsApi.md#GetReportsIdSentTo) | **Get** /reports/{campaign_id}/sent-to | List campaign recipients
[**GetReportsIdSentToId**](ReportsApi.md#GetReportsIdSentToId) | **Get** /reports/{campaign_id}/sent-to/{subscriber_hash} | Get campaign recipient info
[**GetReportsIdSubReportsId**](ReportsApi.md#GetReportsIdSubReportsId) | **Get** /reports/{campaign_id}/sub-reports | List child campaign reports
[**GetReportsIdUnsubscribed**](ReportsApi.md#GetReportsIdUnsubscribed) | **Get** /reports/{campaign_id}/unsubscribed | List unsubscribed members
[**GetReportsIdUnsubscribedId**](ReportsApi.md#GetReportsIdUnsubscribedId) | **Get** /reports/{campaign_id}/unsubscribed/{subscriber_hash} | Get unsubscribed member


# **GetReports**
> CampaignReports1 GetReports(ctx, optional)
List campaign reports

Get campaign reports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiGetReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| The campaign type. | 
 **beforeSendTime** | **optional.Time**| Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceSendTime** | **optional.Time**| Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**CampaignReports1**](Campaign Reports_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsId**
> CampaignReport GetReportsId(ctx, campaignId, optional)
Get campaign report

Get report details for a specific sent campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignReport**](Campaign Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdAbuseReportsId**
> AbuseComplaints1 GetReportsIdAbuseReportsId(ctx, campaignId, optional)
List abuse reports

Get a list of abuse complaints for a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdAbuseReportsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdAbuseReportsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AbuseComplaints1**](Abuse Complaints_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdAbuseReportsIdId**
> AbuseComplaint1 GetReportsIdAbuseReportsIdId(ctx, campaignId, reportId, optional)
Get abuse report

Get information about a specific abuse report for a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **reportId** | **string**| The id for the abuse report. | 
 **optional** | ***ReportsApiGetReportsIdAbuseReportsIdIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdAbuseReportsIdIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AbuseComplaint1**](Abuse Complaint_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdAdvice**
> CampaignAdviceReport GetReportsIdAdvice(ctx, campaignId, optional)
List campaign feedback

Get feedback based on a campaign's statistics. Advice feedback is based on campaign stats like opens, clicks, unsubscribes, bounces, and more.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdAdviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdAdviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignAdviceReport**](Campaign Advice Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdClickDetails**
> ClickDetailReport GetReportsIdClickDetails(ctx, campaignId, optional)
List campaign details

Get information about clicks on specific links in your Mailchimp campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdClickDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdClickDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**ClickDetailReport**](Click Detail Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdClickDetailsId**
> ClickDetailReport GetReportsIdClickDetailsId(ctx, campaignId, linkId, optional)
Get campaign link details

Get click details for a specific link in a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **linkId** | **string**| The id for the link. | 
 **optional** | ***ReportsApiGetReportsIdClickDetailsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdClickDetailsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ClickDetailReport**](Click Detail Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdClickDetailsIdMembers**
> ClickDetailMembers GetReportsIdClickDetailsIdMembers(ctx, campaignId, linkId, optional)
List clicked link subscribers

Get information about list members who clicked on a specific link in a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **linkId** | **string**| The id for the link. | 
 **optional** | ***ReportsApiGetReportsIdClickDetailsIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdClickDetailsIdMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**ClickDetailMembers**](Click Detail Members.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdClickDetailsIdMembersId**
> ClickDetailMember GetReportsIdClickDetailsIdMembersId(ctx, campaignId, linkId, subscriberHash, optional)
Get clicked link subscriber

Get information about a specific subscriber who clicked a link in a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **linkId** | **string**| The id for the link. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ReportsApiGetReportsIdClickDetailsIdMembersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdClickDetailsIdMembersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ClickDetailMember**](Click Detail Member.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdDomainPerformance**
> DomainPerformance GetReportsIdDomainPerformance(ctx, campaignId, optional)
List domain performance stats

Get statistics for the top-performing email domains in a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdDomainPerformanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdDomainPerformanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**DomainPerformance**](Domain Performance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdEcommerceProductActivity**
> InlineResponse2007 GetReportsIdEcommerceProductActivity(ctx, campaignId, optional)
List campaign product activity

Get breakdown of product activity for a campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdEcommerceProductActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdEcommerceProductActivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdEepurl**
> EepurlActivity GetReportsIdEepurl(ctx, campaignId, optional)
List EepURL activity

Get a summary of social activity for the campaign, tracked by EepURL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdEepurlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdEepurlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EepurlActivity**](Eepurl Activity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdEmailActivity**
> EmailActivity GetReportsIdEmailActivity(ctx, campaignId, optional)
List email activity

Get a list of member's subscriber activity in a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdEmailActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdEmailActivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **since** | **optional.String**| Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**EmailActivity**](Email Activity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdEmailActivityId**
> EmailActivity GetReportsIdEmailActivityId(ctx, campaignId, subscriberHash, optional)
Get subscriber email activity

Get a specific list member's activity in a campaign including opens, clicks, and bounces.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ReportsApiGetReportsIdEmailActivityIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdEmailActivityIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **since** | **optional.String**| Restrict results to email activity events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**EmailActivity**](Email Activity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdLocations**
> OpenLocations GetReportsIdLocations(ctx, campaignId, optional)
List top open activities

Get top open locations for a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**OpenLocations**](Open Locations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdOpenDetails**
> OpenDetailReport GetReportsIdOpenDetails(ctx, campaignId, optional)
List campaign open details

Get detailed information about any campaign emails that were opened by a list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdOpenDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdOpenDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **since** | **optional.String**| Restrict results to campaign open events that occur after a specific time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**OpenDetailReport**](Open Detail Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdOpenDetailsIdMembersId**
> OpenActivity GetReportsIdOpenDetailsIdMembersId(ctx, campaignId, subscriberHash, optional)
Get opened campaign subscriber

Get information about a specific subscriber who opened a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ReportsApiGetReportsIdOpenDetailsIdMembersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdOpenDetailsIdMembersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**OpenActivity**](Open Activity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdSentTo**
> SentTo GetReportsIdSentTo(ctx, campaignId, optional)
List campaign recipients

Get information about campaign recipients.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdSentToOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdSentToOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**SentTo**](SentTo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdSentToId**
> SentTo GetReportsIdSentToId(ctx, campaignId, subscriberHash, optional)
Get campaign recipient info

Get information about a specific campaign recipient.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ReportsApiGetReportsIdSentToIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdSentToIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SentTo**](Sent To.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdSubReportsId**
> CampaignSubReports GetReportsIdSubReportsId(ctx, campaignId, optional)
List child campaign reports

Get a list of reports with child campaigns for a specific parent campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdSubReportsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdSubReportsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignSubReports**](Campaign SubReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdUnsubscribed**
> Unsubscribes GetReportsIdUnsubscribed(ctx, campaignId, optional)
List unsubscribed members

Get information about members who have unsubscribed from a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***ReportsApiGetReportsIdUnsubscribedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdUnsubscribedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Unsubscribes**](Unsubscribes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsIdUnsubscribedId**
> Unsubscribes GetReportsIdUnsubscribedId(ctx, campaignId, subscriberHash, optional)
Get unsubscribed member

Get information about a specific list member who unsubscribed from a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ReportsApiGetReportsIdUnsubscribedIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsIdUnsubscribedIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Unsubscribes**](Unsubscribes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

