# \AutomationsApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAutomations**](AutomationsApi.md#ArchiveAutomations) | **Post** /automations/{workflow_id}/actions/archive | Archive automation
[**DeleteAutomationsIdEmailsId**](AutomationsApi.md#DeleteAutomationsIdEmailsId) | **Delete** /automations/{workflow_id}/emails/{workflow_email_id} | Delete workflow email
[**GetAutomations**](AutomationsApi.md#GetAutomations) | **Get** /automations | List automations
[**GetAutomationsId**](AutomationsApi.md#GetAutomationsId) | **Get** /automations/{workflow_id} | Get automation info
[**GetAutomationsIdEmails**](AutomationsApi.md#GetAutomationsIdEmails) | **Get** /automations/{workflow_id}/emails | List automated emails
[**GetAutomationsIdEmailsId**](AutomationsApi.md#GetAutomationsIdEmailsId) | **Get** /automations/{workflow_id}/emails/{workflow_email_id} | Get workflow email info
[**GetAutomationsIdEmailsIdQueue**](AutomationsApi.md#GetAutomationsIdEmailsIdQueue) | **Get** /automations/{workflow_id}/emails/{workflow_email_id}/queue | List automated email subscribers
[**GetAutomationsIdEmailsIdQueueId**](AutomationsApi.md#GetAutomationsIdEmailsIdQueueId) | **Get** /automations/{workflow_id}/emails/{workflow_email_id}/queue/{subscriber_hash} | Get automated email subscriber
[**GetAutomationsIdRemovedSubscribers**](AutomationsApi.md#GetAutomationsIdRemovedSubscribers) | **Get** /automations/{workflow_id}/removed-subscribers | List subscribers removed from workflow
[**GetAutomationsIdRemovedSubscribersId**](AutomationsApi.md#GetAutomationsIdRemovedSubscribersId) | **Get** /automations/{workflow_id}/removed-subscribers/{subscriber_hash} | Get subscriber removed from workflow
[**PatchAutomationEmailWorkflowId**](AutomationsApi.md#PatchAutomationEmailWorkflowId) | **Patch** /automations/{workflow_id}/emails/{workflow_email_id} | Update workflow email
[**PostAutomations**](AutomationsApi.md#PostAutomations) | **Post** /automations | Add automation
[**PostAutomationsIdActionsPauseAllEmails**](AutomationsApi.md#PostAutomationsIdActionsPauseAllEmails) | **Post** /automations/{workflow_id}/actions/pause-all-emails | Pause automation emails
[**PostAutomationsIdActionsStartAllEmails**](AutomationsApi.md#PostAutomationsIdActionsStartAllEmails) | **Post** /automations/{workflow_id}/actions/start-all-emails | Start automation emails
[**PostAutomationsIdEmailsIdActionsPause**](AutomationsApi.md#PostAutomationsIdEmailsIdActionsPause) | **Post** /automations/{workflow_id}/emails/{workflow_email_id}/actions/pause | Pause automated email
[**PostAutomationsIdEmailsIdActionsStart**](AutomationsApi.md#PostAutomationsIdEmailsIdActionsStart) | **Post** /automations/{workflow_id}/emails/{workflow_email_id}/actions/start | Start automated email
[**PostAutomationsIdEmailsIdQueue**](AutomationsApi.md#PostAutomationsIdEmailsIdQueue) | **Post** /automations/{workflow_id}/emails/{workflow_email_id}/queue | Add subscriber to workflow email
[**PostAutomationsIdRemovedSubscribers**](AutomationsApi.md#PostAutomationsIdRemovedSubscribers) | **Post** /automations/{workflow_id}/removed-subscribers | Remove subscriber from workflow


# **ArchiveAutomations**
> ArchiveAutomations(ctx, workflowId)
Archive automation

Archiving will permanently end your automation and keep the report data. You’ll be able to replicate your archived automation, but you can’t restart it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutomationsIdEmailsId**
> DeleteAutomationsIdEmailsId(ctx, workflowId, workflowEmailId)
Delete workflow email

Removes an individual classic automation workflow email. Emails from certain workflow types, including the Abandoned Cart Email (abandonedCart) and Product Retargeting Email (abandonedBrowse) Workflows, cannot be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomations**
> InlineResponse2003 GetAutomations(ctx, optional)
List automations

Get a summary of an account's classic automations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutomationsApiGetAutomationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutomationsApiGetAutomationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **beforeCreateTime** | **optional.Time**| Restrict the response to automations created before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCreateTime** | **optional.Time**| Restrict the response to automations created after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeStartTime** | **optional.Time**| Restrict the response to automations started before this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceStartTime** | **optional.Time**| Restrict the response to automations started after this time. Uses the ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **status** | **optional.String**| Restrict the results to automations with the specified status. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsId**
> AutomationWorkflow GetAutomationsId(ctx, workflowId, optional)
Get automation info

Get a summary of an individual classic automation workflow's settings and content. The `trigger_settings` object returns information for the first email in the workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
 **optional** | ***AutomationsApiGetAutomationsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutomationsApiGetAutomationsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**AutomationWorkflow**](Automation Workflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsIdEmails**
> AutomationEmails GetAutomationsIdEmails(ctx, workflowId)
List automated emails

Get a summary of the emails in a classic automation workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 

### Return type

[**AutomationEmails**](Automation Emails.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsIdEmailsId**
> AutomationWorkflowEmail GetAutomationsIdEmailsId(ctx, workflowId, workflowEmailId)
Get workflow email info

Get information about an individual classic automation workflow email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 

### Return type

[**AutomationWorkflowEmail**](Automation Workflow Email.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsIdEmailsIdQueue**
> InlineResponse2004 GetAutomationsIdEmailsIdQueue(ctx, workflowId, workflowEmailId)
List automated email subscribers

Get information about a classic automation email queue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsIdEmailsIdQueueId**
> SubscriberInAutomationQueue2 GetAutomationsIdEmailsIdQueueId(ctx, workflowId, workflowEmailId, subscriberHash)
Get automated email subscriber

Get information about a specific subscriber in a classic automation email queue.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Return type

[**SubscriberInAutomationQueue2**](Subscriber in Automation Queue_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsIdRemovedSubscribers**
> RemovedSubscribers GetAutomationsIdRemovedSubscribers(ctx, workflowId)
List subscribers removed from workflow

Get information about subscribers who were removed from a classic automation workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 

### Return type

[**RemovedSubscribers**](Removed Subscribers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAutomationsIdRemovedSubscribersId**
> SubscriberRemovedFromAutomationWorkflow GetAutomationsIdRemovedSubscribersId(ctx, workflowId, subscriberHash)
Get subscriber removed from workflow

Get information about a specific subscriber who was removed from a classic automation workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Return type

[**SubscriberRemovedFromAutomationWorkflow**](Subscriber Removed from Automation Workflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchAutomationEmailWorkflowId**
> AutomationWorkflowEmail PatchAutomationEmailWorkflowId(ctx, workflowId, workflowEmailId, body)
Update workflow email

Update settings for a classic automation workflow email.  Only works with workflows of type: abandonedBrowse, abandonedCart, emailFollowup, or singleWelcome.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 
  **body** | [**UpdateInformationAboutASpecificWorkflowEmail**](UpdateInformationAboutASpecificWorkflowEmail.md)|  | 

### Return type

[**AutomationWorkflowEmail**](Automation Workflow Email.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomations**
> AutomationWorkflow PostAutomations(ctx, body)
Add automation

Create a new classic automation in your Mailchimp account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AutomationWorkflow1**](AutomationWorkflow1.md)|  | 

### Return type

[**AutomationWorkflow**](Automation Workflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomationsIdActionsPauseAllEmails**
> PostAutomationsIdActionsPauseAllEmails(ctx, workflowId)
Pause automation emails

Pause all emails in a specific classic automation workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomationsIdActionsStartAllEmails**
> PostAutomationsIdActionsStartAllEmails(ctx, workflowId)
Start automation emails

Start all emails in a classic automation workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomationsIdEmailsIdActionsPause**
> PostAutomationsIdEmailsIdActionsPause(ctx, workflowId, workflowEmailId)
Pause automated email

Pause an automated email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomationsIdEmailsIdActionsStart**
> PostAutomationsIdEmailsIdActionsStart(ctx, workflowId, workflowEmailId)
Start automated email

Start an automated email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomationsIdEmailsIdQueue**
> SubscriberInAutomationQueue2 PostAutomationsIdEmailsIdQueue(ctx, workflowId, workflowEmailId, body)
Add subscriber to workflow email

Manually add a subscriber to a workflow, bypassing the default trigger settings. You can also use this endpoint to trigger a series of automated emails in an API 3.0 workflow type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **workflowEmailId** | **string**| The unique id for the Automation workflow email. | 
  **body** | [**SubscriberInAutomationQueue1**](SubscriberInAutomationQueue1.md)|  | 

### Return type

[**SubscriberInAutomationQueue2**](Subscriber in Automation Queue_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAutomationsIdRemovedSubscribers**
> SubscriberRemovedFromAutomationWorkflow PostAutomationsIdRemovedSubscribers(ctx, workflowId, body)
Remove subscriber from workflow

Remove a subscriber from a specific classic automation workflow. You can remove a subscriber at any point in an automation workflow, regardless of how many emails they've been sent from that workflow. Once they're removed, they can never be added back to the same workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**| The unique id for the Automation workflow. | 
  **body** | [**SubscriberInAutomationQueue3**](SubscriberInAutomationQueue3.md)|  | 

### Return type

[**SubscriberRemovedFromAutomationWorkflow**](Subscriber Removed from Automation Workflow.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

