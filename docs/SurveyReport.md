# SurveyReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this survey. | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. View this survey report in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/lists/surveys/results?survey_id&#x3D;{web_id}&#x60;. | [optional] [default to null]
**ListId** | **string** | The ID of the list connected to this survey. | [optional] [default to null]
**ListName** | **string** | The name of the list connected to this survey. | [optional] [default to null]
**Title** | **string** | The title of the survey. | [optional] [default to null]
**Url** | **string** | The URL for the survey. | [optional] [default to null]
**Status** | **string** | The survey&#39;s status. | [optional] [default to null]
**PublishedAt** | [**time.Time**](time.Time.md) | The date and time the survey was published in ISO 8601 format. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the survey was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the survey was last updated in ISO 8601 format. | [optional] [default to null]
**TotalResponses** | **int32** | The total number of responses to this survey. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


