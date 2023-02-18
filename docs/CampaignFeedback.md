# CampaignFeedback

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackId** | **int32** | The individual id for the feedback item. | [optional] [default to null]
**ParentId** | **int32** | If a reply, the id of the parent feedback item. | [optional] [default to null]
**BlockId** | **int32** | The block id for the editable block that the feedback addresses. | [optional] [default to null]
**Message** | **string** | The content of the feedback. | [default to null]
**IsComplete** | **bool** | The status of feedback. | [optional] [default to null]
**CreatedBy** | **string** | The login name of the user who created the feedback. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the feedback item was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the feedback was last updated in ISO 8601 format. | [optional] [default to null]
**Source** | **string** | The source of the feedback. | [optional] [default to null]
**CampaignId** | **string** | The unique id for the campaign. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


