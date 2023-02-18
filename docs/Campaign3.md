# Campaign3

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this campaign. | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. View this campaign in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/campaigns/show/?id&#x3D;{web_id}&#x60;. | [optional] [default to null]
**ParentCampaignId** | **string** | If this campaign is the child of another campaign, this identifies the parent campaign. For Example, for RSS or Automation children. | [optional] [default to null]
**Type_** | **string** | There are four types of [campaigns](https://mailchimp.com/help/getting-started-with-campaigns/) you can create in Mailchimp. A/B Split campaigns have been deprecated and variate campaigns should be used instead. | [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | The date and time the campaign was created in ISO 8601 format. | [optional] [default to null]
**ArchiveUrl** | **string** | The link to the campaign&#39;s archive version. | [optional] [default to null]
**LongArchiveUrl** | **string** | The original link to the campaign&#39;s archive version. | [optional] [default to null]
**Status** | **string** | The current status of the campaign. | [optional] [default to null]
**EmailsSent** | **int32** | The total number of emails sent for this campaign. | [optional] [default to null]
**SendTime** | [**time.Time**](time.Time.md) | The date and time a campaign was sent in ISO 8601 format. | [optional] [default to null]
**ContentType** | **string** | How the campaign&#39;s content is put together (&#39;template&#39;, &#39;drag_and_drop&#39;, &#39;html&#39;, &#39;url&#39;). | [optional] [default to null]
**NeedsBlockRefresh** | **bool** | Determines if the campaign needs its blocks refreshed by opening the web-based campaign editor. Deprecated and will always return false. | [optional] [default to null]
**Resendable** | **bool** | Determines if the campaign qualifies to be resent to non-openers. | [optional] [default to null]
**Recipients** | [***List6**](List_6.md) |  | [optional] [default to null]
**Settings** | [***CampaignSettings5**](Campaign Settings_5.md) |  | [optional] [default to null]
**VariateSettings** | [***AbTestOptions2**](AB Test Options_2.md) |  | [optional] [default to null]
**Tracking** | [***CampaignTrackingOptions1**](Campaign Tracking Options_1.md) |  | [optional] [default to null]
**RssOpts** | [***RssOptions3**](RSS Options_3.md) |  | [optional] [default to null]
**AbSplitOpts** | [***AbTestingOptions**](AB Testing Options.md) |  | [optional] [default to null]
**SocialCard** | [***CampaignSocialCard**](Campaign Social Card.md) |  | [optional] [default to null]
**ReportSummary** | [***CampaignReportSummary3**](Campaign Report Summary_3.md) |  | [optional] [default to null]
**DeliveryStatus** | [***CampaignDeliveryStatus**](Campaign Delivery Status.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


