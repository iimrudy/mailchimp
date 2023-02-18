# InlineResponse2009

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]
**Audience** | [***InlineResponse2009Audience**](inline_response_200_9_audience.md) |  | [optional] [default to null]
**Budget** | [***InlineResponse2009Budget**](inline_response_200_9_budget.md) |  | [optional] [default to null]
**CanceledAt** | [**time.Time**](time.Time.md) | The date and time the outreach was canceled in ISO 8601 format. | [optional] [default to null]
**Channel** | [***InlineResponse2009Channel**](inline_response_200_9_channel.md) |  | [optional] [default to null]
**Content** | [***InlineResponse2009Content**](inline_response_200_9_content.md) |  | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The date and time the outreach was created in ISO 8601 format. | [optional] [default to null]
**EmailSourceName** | **string** |  | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | The date and time the ad was ended in ISO 8601 format. | [optional] [default to null]
**Feedback** | [***InlineResponse2009Feedback**](inline_response_200_9_feedback.md) |  | [optional] [default to null]
**HasAudience** | **bool** | Check if this ad has audience setup | [optional] [default to null]
**HasContent** | **bool** | Check if this ad has content | [optional] [default to null]
**HasSegment** | **bool** | If this outreach targets a segment of your audience. | [optional] [default to null]
**Id** | **string** | Unique ID of an Outreach. | [optional] [default to null]
**IsConnected** | **bool** | Check if this ad is connected to a facebook page | [optional] [default to null]
**Name** | **string** | Title or name of an Outreach. | [optional] [default to null]
**NeedsAttention** | **bool** | If the ad has a problem and needs attention. | [optional] [default to null]
**PausedAt** | [**time.Time**](time.Time.md) | The date and time the ad was paused in ISO 8601 format. | [optional] [default to null]
**PublishedTime** | [**time.Time**](time.Time.md) | The date and time the outreach was (or will be) published in ISO 8601 format. | [optional] [default to null]
**Recipients** | [***Recipients**](Recipients.md) |  | [optional] [default to null]
**ReportSummary** | [***ReportSummary**](Report Summary.md) |  | [optional] [default to null]
**ShowReport** | **bool** | Outreach report availability. Note: This property is hotly debated in what it _should_ convey. See [MCP-1371](https://jira.mailchimp.com/browse/MCP-1371) for more context. | [optional] [default to null]
**Site** | [***InlineResponse2009Site**](inline_response_200_9_site.md) |  | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The date and time the outreach was started in ISO 8601 format. | [optional] [default to null]
**Status** | **string** | The status of this outreach. | [optional] [default to null]
**Thumbnail** | **string** | The URL of the thumbnail for this outreach. | [optional] [default to null]
**Type_** | **string** | The type of outreach this object is. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the outreach was last updated in ISO 8601 format. | [optional] [default to null]
**WasCanceledByFacebook** | **bool** |  | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. For example, for a &#x60;regular&#x60; outreach, you can view this campaign in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


