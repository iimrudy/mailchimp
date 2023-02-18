# \CampaignsApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCampaignsId**](CampaignsApi.md#DeleteCampaignsId) | **Delete** /campaigns/{campaign_id} | Delete campaign
[**DeleteCampaignsIdFeedbackId**](CampaignsApi.md#DeleteCampaignsIdFeedbackId) | **Delete** /campaigns/{campaign_id}/feedback/{feedback_id} | Delete campaign feedback message
[**GetCampaigns**](CampaignsApi.md#GetCampaigns) | **Get** /campaigns | List campaigns
[**GetCampaignsId**](CampaignsApi.md#GetCampaignsId) | **Get** /campaigns/{campaign_id} | Get campaign info
[**GetCampaignsIdContent**](CampaignsApi.md#GetCampaignsIdContent) | **Get** /campaigns/{campaign_id}/content | Get campaign content
[**GetCampaignsIdFeedback**](CampaignsApi.md#GetCampaignsIdFeedback) | **Get** /campaigns/{campaign_id}/feedback | List campaign feedback
[**GetCampaignsIdFeedbackId**](CampaignsApi.md#GetCampaignsIdFeedbackId) | **Get** /campaigns/{campaign_id}/feedback/{feedback_id} | Get campaign feedback message
[**GetCampaignsIdSendChecklist**](CampaignsApi.md#GetCampaignsIdSendChecklist) | **Get** /campaigns/{campaign_id}/send-checklist | Get campaign send checklist
[**PatchCampaignsId**](CampaignsApi.md#PatchCampaignsId) | **Patch** /campaigns/{campaign_id} | Update campaign settings
[**PatchCampaignsIdFeedbackId**](CampaignsApi.md#PatchCampaignsIdFeedbackId) | **Patch** /campaigns/{campaign_id}/feedback/{feedback_id} | Update campaign feedback message
[**PostCampaigns**](CampaignsApi.md#PostCampaigns) | **Post** /campaigns | Add campaign
[**PostCampaignsIdActionsCancelSend**](CampaignsApi.md#PostCampaignsIdActionsCancelSend) | **Post** /campaigns/{campaign_id}/actions/cancel-send | Cancel campaign
[**PostCampaignsIdActionsCreateResend**](CampaignsApi.md#PostCampaignsIdActionsCreateResend) | **Post** /campaigns/{campaign_id}/actions/create-resend | Resend campaign
[**PostCampaignsIdActionsPause**](CampaignsApi.md#PostCampaignsIdActionsPause) | **Post** /campaigns/{campaign_id}/actions/pause | Pause rss campaign
[**PostCampaignsIdActionsReplicate**](CampaignsApi.md#PostCampaignsIdActionsReplicate) | **Post** /campaigns/{campaign_id}/actions/replicate | Replicate campaign
[**PostCampaignsIdActionsResume**](CampaignsApi.md#PostCampaignsIdActionsResume) | **Post** /campaigns/{campaign_id}/actions/resume | Resume rss campaign
[**PostCampaignsIdActionsSchedule**](CampaignsApi.md#PostCampaignsIdActionsSchedule) | **Post** /campaigns/{campaign_id}/actions/schedule | Schedule campaign
[**PostCampaignsIdActionsSend**](CampaignsApi.md#PostCampaignsIdActionsSend) | **Post** /campaigns/{campaign_id}/actions/send | Send campaign
[**PostCampaignsIdActionsTest**](CampaignsApi.md#PostCampaignsIdActionsTest) | **Post** /campaigns/{campaign_id}/actions/test | Send test email
[**PostCampaignsIdActionsUnschedule**](CampaignsApi.md#PostCampaignsIdActionsUnschedule) | **Post** /campaigns/{campaign_id}/actions/unschedule | Unschedule campaign
[**PostCampaignsIdFeedback**](CampaignsApi.md#PostCampaignsIdFeedback) | **Post** /campaigns/{campaign_id}/feedback | Add campaign feedback
[**PutCampaignsIdContent**](CampaignsApi.md#PutCampaignsIdContent) | **Put** /campaigns/{campaign_id}/content | Set campaign content


# **DeleteCampaignsId**
> DeleteCampaignsId(ctx, campaignId)
Delete campaign

Remove a campaign from your Mailchimp account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCampaignsIdFeedbackId**
> DeleteCampaignsIdFeedbackId(ctx, campaignId, feedbackId)
Delete campaign feedback message

Remove a specific feedback message for a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **feedbackId** | **string**| The unique id for the feedback message. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaigns**
> InlineResponse2005 GetCampaigns(ctx, optional)
List campaigns

Get all campaigns in an account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CampaignsApiGetCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| The campaign type. | 
 **status** | **optional.String**| The status of the campaign. | 
 **beforeSendTime** | **optional.Time**| Restrict the response to campaigns sent before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceSendTime** | **optional.Time**| Restrict the response to campaigns sent after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeCreateTime** | **optional.Time**| Restrict the response to campaigns created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreateTime** | **optional.Time**| Restrict the response to campaigns created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **listId** | **optional.String**| The unique id for the list. | 
 **folderId** | **optional.String**| The unique folder id. | 
 **memberId** | **optional.String**| Retrieve campaigns sent to a particular list member. Member ID is The MD5 hash of the lowercase version of the list member’s email address. | 
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsId**
> Campaign GetCampaignsId(ctx, campaignId, optional)
Get campaign info

Get information about a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***CampaignsApiGetCampaignsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsIdContent**
> CampaignContent GetCampaignsIdContent(ctx, campaignId, optional)
Get campaign content

Get the the HTML and plain-text content for a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***CampaignsApiGetCampaignsIdContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsIdContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignContent**](Campaign Content.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsIdFeedback**
> CampaignReports GetCampaignsIdFeedback(ctx, campaignId, optional)
List campaign feedback

Get team feedback while you're working together on a Mailchimp campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***CampaignsApiGetCampaignsIdFeedbackOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsIdFeedbackOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignReports**](Campaign Reports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsIdFeedbackId**
> CampaignFeedback2 GetCampaignsIdFeedbackId(ctx, campaignId, feedbackId, optional)
Get campaign feedback message

Get a specific feedback message from a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **feedbackId** | **string**| The unique id for the feedback message. | 
 **optional** | ***CampaignsApiGetCampaignsIdFeedbackIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsIdFeedbackIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CampaignFeedback2**](Campaign Feedback_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsIdSendChecklist**
> SendChecklist GetCampaignsIdSendChecklist(ctx, campaignId, optional)
Get campaign send checklist

Review the send checklist for a campaign, and resolve any issues before sending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
 **optional** | ***CampaignsApiGetCampaignsIdSendChecklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiGetCampaignsIdSendChecklistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SendChecklist**](Send Checklist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCampaignsId**
> Campaign PatchCampaignsId(ctx, campaignId, body)
Update campaign settings

Update some or all of the settings for a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **body** | [**Campaign2**](Campaign2.md)|  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCampaignsIdFeedbackId**
> CampaignFeedback2 PatchCampaignsIdFeedbackId(ctx, campaignId, feedbackId, body)
Update campaign feedback message

Update a specific feedback message for a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **feedbackId** | **string**| The unique id for the feedback message. | 
  **body** | [**CampaignFeedback3**](CampaignFeedback3.md)|  | 

### Return type

[**CampaignFeedback2**](Campaign Feedback_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaigns**
> Campaign PostCampaigns(ctx, body)
Add campaign

Create a new Mailchimp campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Campaign1**](Campaign1.md)|  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsCancelSend**
> PostCampaignsIdActionsCancelSend(ctx, campaignId)
Cancel campaign

Cancel a Regular or Plain-Text Campaign after you send, before all of your recipients receive it. This feature is included with Mailchimp Pro.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsCreateResend**
> Campaign3 PostCampaignsIdActionsCreateResend(ctx, campaignId)
Resend campaign

Creates a Resend to Non-Openers version of this campaign. We will also check if this campaign meets the criteria for Resend to Non-Openers campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

[**Campaign3**](Campaign_3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsPause**
> PostCampaignsIdActionsPause(ctx, campaignId)
Pause rss campaign

Pause an RSS-Driven campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsReplicate**
> Campaign3 PostCampaignsIdActionsReplicate(ctx, campaignId)
Replicate campaign

Replicate a campaign in saved or send status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

[**Campaign3**](Campaign_3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsResume**
> PostCampaignsIdActionsResume(ctx, campaignId)
Resume rss campaign

Resume an RSS-Driven campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsSchedule**
> PostCampaignsIdActionsSchedule(ctx, campaignId, body)
Schedule campaign

Schedule a campaign for delivery. If you're using Multivariate Campaigns to test send times or sending RSS Campaigns, use the send action instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **body** | [**Body1**](Body1.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsSend**
> PostCampaignsIdActionsSend(ctx, campaignId)
Send campaign

Send a Mailchimp campaign. For RSS Campaigns, the campaign will send according to its schedule. All other campaigns will send immediately.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsTest**
> PostCampaignsIdActionsTest(ctx, campaignId, body)
Send test email

Send a test email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **body** | [**Body2**](Body2.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdActionsUnschedule**
> PostCampaignsIdActionsUnschedule(ctx, campaignId)
Unschedule campaign

Unschedule a scheduled campaign that hasn't started sending.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCampaignsIdFeedback**
> CampaignFeedback2 PostCampaignsIdFeedback(ctx, campaignId, body)
Add campaign feedback

Add feedback on a specific campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **body** | [**CampaignFeedback1**](CampaignFeedback1.md)|  | 

### Return type

[**CampaignFeedback2**](Campaign Feedback_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCampaignsIdContent**
> CampaignContent PutCampaignsIdContent(ctx, campaignId, body)
Set campaign content

Set the content for a campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **string**| The unique id for the campaign. | 
  **body** | [**CampaignContent1**](CampaignContent1.md)|  | 

### Return type

[**CampaignContent**](Campaign Content.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

