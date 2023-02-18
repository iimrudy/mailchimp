# SubscriberInAutomationQueue

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | [optional] [default to null]
**WorkflowId** | **string** | A string that uniquely identifies an Automation workflow. | [optional] [default to null]
**EmailId** | **string** | A string that uniquely identifies an email in an Automation workflow. | [optional] [default to null]
**ListId** | **string** | A string that uniquely identifies a list. | [optional] [default to null]
**EmailAddress** | **string** | The list member&#39;s email address. | [default to null]
**NextSend** | [**time.Time**](time.Time.md) | The date and time of the next send for the workflow email in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


