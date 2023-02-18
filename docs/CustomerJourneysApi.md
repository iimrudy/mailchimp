# \CustomerJourneysApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCustomerJourneysJourneysIdStepsIdActionsTrigger**](CustomerJourneysApi.md#PostCustomerJourneysJourneysIdStepsIdActionsTrigger) | **Post** /customer-journeys/journeys/{journey_id}/steps/{step_id}/actions/trigger | Customer Journeys API trigger for a contact


# **PostCustomerJourneysJourneysIdStepsIdActionsTrigger**
> interface{} PostCustomerJourneysJourneysIdStepsIdActionsTrigger(ctx, journeyId, stepId, body)
Customer Journeys API trigger for a contact

A step trigger in a Customer Journey. To use it, create a starting point or step from the Customer Journey builder in the app using the Customer Journeys API condition. We’ll provide a url during the process that includes the {journey_id} and {step_id}. You’ll then be able to use this endpoint to trigger the condition for the posted contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **journeyId** | **int32**| The id for the Journey. | 
  **stepId** | **int32**| The id for the Step. | 
  **body** | [**SubscriberInCustomerJourneysAudience**](SubscriberInCustomerJourneysAudience.md)|  | 

### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

