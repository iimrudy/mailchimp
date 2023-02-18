# \SurveysApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostListsIdSurveysIdActionsCreateEmail**](SurveysApi.md#PostListsIdSurveysIdActionsCreateEmail) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/create-email | Create a Survey Campaign
[**PostListsIdSurveysIdActionsPublish**](SurveysApi.md#PostListsIdSurveysIdActionsPublish) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/publish | Publish a Survey
[**PostListsIdSurveysIdActionsUnpublish**](SurveysApi.md#PostListsIdSurveysIdActionsUnpublish) | **Post** /lists/{list_id}/surveys/{survey_id}/actions/unpublish | Unpublish a Survey


# **PostListsIdSurveysIdActionsCreateEmail**
> Campaign3 PostListsIdSurveysIdActionsCreateEmail(ctx, listId, surveyId)
Create a Survey Campaign

Utilize the List ID and Survey ID to generate a Campaign that links to your survey.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **surveyId** | **string**| The ID of the survey. | 

### Return type

[**Campaign3**](Campaign_3.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSurveysIdActionsPublish**
> PostListsIdSurveysIdActionsPublish(ctx, listId, surveyId)
Publish a Survey

Publish a survey that is in draft, unpublished, or has been previously published and edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **surveyId** | **string**| The ID of the survey. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSurveysIdActionsUnpublish**
> PostListsIdSurveysIdActionsUnpublish(ctx, listId, surveyId)
Unpublish a Survey

Unpublish a survey that has been published.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **surveyId** | **string**| The ID of the survey. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

