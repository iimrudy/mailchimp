# \ReportingApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReportingFacebookAds**](ReportingApi.md#GetReportingFacebookAds) | **Get** /reporting/facebook-ads | List facebook ads reports
[**GetReportingFacebookAdsId**](ReportingApi.md#GetReportingFacebookAdsId) | **Get** /reporting/facebook-ads/{outreach_id} | Get facebook ad report
[**GetReportingFacebookAdsIdEcommerceProductActivity**](ReportingApi.md#GetReportingFacebookAdsIdEcommerceProductActivity) | **Get** /reporting/facebook-ads/{outreach_id}/ecommerce-product-activity | List facebook ecommerce report
[**GetReportingLandingPages**](ReportingApi.md#GetReportingLandingPages) | **Get** /reporting/landing-pages | List landing pages reports
[**GetReportingLandingPagesId**](ReportingApi.md#GetReportingLandingPagesId) | **Get** /reporting/landing-pages/{outreach_id} | Get landing page report
[**GetReportingSurveys**](ReportingApi.md#GetReportingSurveys) | **Get** /reporting/surveys | List survey reports
[**GetReportingSurveysId**](ReportingApi.md#GetReportingSurveysId) | **Get** /reporting/surveys/{outreach_id} | Get survey report
[**GetReportingSurveysIdQuestions**](ReportingApi.md#GetReportingSurveysIdQuestions) | **Get** /reporting/surveys/{outreach_id}/questions | List survey question reports
[**GetReportingSurveysIdQuestionsId**](ReportingApi.md#GetReportingSurveysIdQuestionsId) | **Get** /reporting/surveys/{outreach_id}/questions/{question_id} | Get survey question report
[**GetReportingSurveysIdQuestionsIdAnswers**](ReportingApi.md#GetReportingSurveysIdQuestionsIdAnswers) | **Get** /reporting/surveys/{outreach_id}/questions/{question_id}/answers | List answers for question
[**GetReportingSurveysIdResponses**](ReportingApi.md#GetReportingSurveysIdResponses) | **Get** /reporting/surveys/{outreach_id}/responses | List survey responses
[**GetReportingSurveysIdResponsesId**](ReportingApi.md#GetReportingSurveysIdResponsesId) | **Get** /reporting/surveys/{outreach_id}/responses/{response_id} | Get survey response


# **GetReportingFacebookAds**
> InlineResponse20010 GetReportingFacebookAds(ctx, optional)
List facebook ads reports

Get reports of Facebook ads.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiGetReportingFacebookAdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingFacebookAdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingFacebookAdsId**
> InlineResponse20011 GetReportingFacebookAdsId(ctx, outreachId, optional)
Get facebook ad report

Get report of a Facebook ad.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
 **optional** | ***ReportingApiGetReportingFacebookAdsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingFacebookAdsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingFacebookAdsIdEcommerceProductActivity**
> InlineResponse2007 GetReportingFacebookAdsIdEcommerceProductActivity(ctx, outreachId, optional)
List facebook ecommerce report

Get breakdown of product activity for an outreach.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
 **optional** | ***ReportingApiGetReportingFacebookAdsIdEcommerceProductActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingFacebookAdsIdEcommerceProductActivityOpts struct

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

# **GetReportingLandingPages**
> InlineResponse20012 GetReportingLandingPages(ctx, optional)
List landing pages reports

Get reports of landing pages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiGetReportingLandingPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingLandingPagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingLandingPagesId**
> LandingPageReport GetReportingLandingPagesId(ctx, outreachId, optional)
Get landing page report

Get report of a landing page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
 **optional** | ***ReportingApiGetReportingLandingPagesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingLandingPagesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**LandingPageReport**](Landing Page Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveys**
> InlineResponse20013 GetReportingSurveys(ctx, optional)
List survey reports

Get reports for surveys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportingApiGetReportingSurveysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingSurveysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveysId**
> SurveyReport GetReportingSurveysId(ctx, outreachId, optional)
Get survey report

Get report for a survey.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
 **optional** | ***ReportingApiGetReportingSurveysIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingSurveysIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SurveyReport**](Survey Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveysIdQuestions**
> InlineResponse20014 GetReportingSurveysIdQuestions(ctx, outreachId, optional)
List survey question reports

Get reports for survey questions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
 **optional** | ***ReportingApiGetReportingSurveysIdQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingSurveysIdQuestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveysIdQuestionsId**
> SurveyQuestionReport GetReportingSurveysIdQuestionsId(ctx, outreachId, questionId, optional)
Get survey question report

Get report for a survey question.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
  **questionId** | **string**| The ID of the survey question. | 
 **optional** | ***ReportingApiGetReportingSurveysIdQuestionsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingSurveysIdQuestionsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**SurveyQuestionReport**](Survey Question Report.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveysIdQuestionsIdAnswers**
> InlineResponse20015 GetReportingSurveysIdQuestionsIdAnswers(ctx, outreachId, questionId, optional)
List answers for question

Get answers for a survey question.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
  **questionId** | **string**| The ID of the survey question. | 
 **optional** | ***ReportingApiGetReportingSurveysIdQuestionsIdAnswersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingSurveysIdQuestionsIdAnswersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **respondentFamiliarityIs** | **optional.String**| Filter survey responses by familiarity of the respondents. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveysIdResponses**
> InlineResponse20016 GetReportingSurveysIdResponses(ctx, outreachId, optional)
List survey responses

Get responses to a survey.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
 **optional** | ***ReportingApiGetReportingSurveysIdResponsesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiGetReportingSurveysIdResponsesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **answeredQuestion** | **optional.Int32**| The ID of the question that was answered. | 
 **choseAnswer** | **optional.String**| The ID of the option chosen to filter responses on. | 
 **respondentFamiliarityIs** | **optional.String**| Filter survey responses by familiarity of the respondents. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingSurveysIdResponsesId**
> SurveyResponse GetReportingSurveysIdResponsesId(ctx, outreachId, responseId)
Get survey response

Get a single survey response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outreachId** | **string**| The outreach id. | 
  **responseId** | **string**| The ID of the survey response. | 

### Return type

[**SurveyResponse**](Survey Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

