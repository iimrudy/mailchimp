# SurveyQuestionAnswer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the answer. | [optional] [default to null]
**Value** | **string** | The raw text answer. | [optional] [default to null]
**ResponseId** | **string** | The ID of the survey response. | [optional] [default to null]
**SubmittedAt** | [**time.Time**](time.Time.md) | The date and time when the survey response was submitted in ISO 8601 format. | [optional] [default to null]
**Contact** | [***Contact**](Contact.md) |  | [optional] [default to null]
**IsNewContact** | **bool** | If this contact was added to the Mailchimp audience via this survey. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


