# AutomationWorkflowEmail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies the Automation email. | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. View this automation in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;. | [optional] [default to null]
**WorkflowId** | **string** | A string that uniquely identifies an Automation workflow. | [optional] [default to null]
**Position** | **int32** | The position of an Automation email in a workflow. | [optional] [default to null]
**Delay** | [***AutomationDelay**](Automation Delay.md) |  | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The date and time the campaign was created in ISO 8601 format. | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The date and time the campaign was started in ISO 8601 format. | [optional] [default to null]
**ArchiveUrl** | **string** | The link to the campaign&#39;s archive version in ISO 8601 format. | [optional] [default to null]
**Status** | **string** | The current status of the campaign. | [optional] [default to null]
**EmailsSent** | **int32** | The total number of emails sent for this campaign. | [optional] [default to null]
**SendTime** | [**time.Time**](time.Time.md) |  The date and time a campaign was sent in ISO 8601 format | [optional] [default to null]
**ContentType** | **string** | How the campaign&#39;s content is put together (&#39;template&#39;, &#39;drag_and_drop&#39;, &#39;html&#39;, &#39;url&#39;). | [optional] [default to null]
**NeedsBlockRefresh** | **bool** | Determines if the automation email needs its blocks refreshed by opening the web-based campaign editor. | [optional] [default to null]
**HasLogoMergeTag** | **bool** | Determines if the campaign contains the *|BRAND:LOGO|* merge tag. | [optional] [default to null]
**Recipients** | [***List2**](List_2.md) |  | [optional] [default to null]
**Settings** | [***CampaignSettings**](Campaign Settings.md) |  | [optional] [default to null]
**Tracking** | [***CampaignTrackingOptions**](Campaign Tracking Options.md) |  | [optional] [default to null]
**SocialCard** | [***CampaignSocialCard**](Campaign Social Card.md) |  | [optional] [default to null]
**TriggerSettings** | [***AutomationTrigger**](Automation Trigger.md) |  | [optional] [default to null]
**ReportSummary** | [***CampaignReportSummary1**](Campaign Report Summary_1.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


