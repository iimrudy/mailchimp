# SurveyQuestionReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the survey question. | [optional] [default to null]
**SurveyId** | **string** | The unique ID of the survey. | [optional] [default to null]
**Query** | **string** | The query of the survey question. | [optional] [default to null]
**Type_** | **string** | The response type of the survey question. | [optional] [default to null]
**TotalResponses** | **int32** | The total number of responses to this question. | [optional] [default to null]
**IsRequired** | **bool** | Whether this survey question is required to answer. | [optional] [default to null]
**HasOther** | **bool** | Whether this survey question has an &#39;other&#39; option. | [optional] [default to null]
**OtherLabel** | **string** | Label used for the &#39;other&#39; option of this survey question. | [optional] [default to null]
**AverageRating** | **float32** | The average rating for this range question. | [optional] [default to null]
**RangeLowLabel** | **string** | Label for the low end of the range. | [optional] [default to null]
**RangeHighLabel** | **string** | Label for the high end of the range. | [optional] [default to null]
**PlaceholderLabel** | **string** | Placeholder text for this survey question&#39;s answer box. | [optional] [default to null]
**SubscribeCheckboxEnabled** | **bool** | Whether the subscribe checkbox is shown for this email question. | [optional] [default to null]
**SubscribeCheckboxLabel** | **string** | Label used for the subscribe checkbox for this email question. | [optional] [default to null]
**MergeField** | [***MergeField3**](Merge Field_3.md) |  | [optional] [default to null]
**Options** | [**[]InlineResponse20014Options**](inline_response_200_14_options.md) | The answer choices for this question. | [optional] [default to null]
**ContactCounts** | [***ContactCounts**](Contact Counts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


