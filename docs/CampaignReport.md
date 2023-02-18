# CampaignReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this campaign. | [optional] [default to null]
**CampaignTitle** | **string** | The title of the campaign. | [optional] [default to null]
**Type_** | **string** | The type of campaign (regular, plain-text, ab_split, rss, automation, variate, or auto). | [optional] [default to null]
**ListId** | **string** | The unique list id. | [optional] [default to null]
**ListIsActive** | **bool** | The status of the list used, namely if it&#39;s deleted or disabled. | [optional] [default to null]
**ListName** | **string** | The name of the list. | [optional] [default to null]
**SubjectLine** | **string** | The subject line for the campaign. | [optional] [default to null]
**PreviewText** | **string** | The preview text for the campaign. | [optional] [default to null]
**EmailsSent** | **int32** | The total number of emails sent for this campaign. | [optional] [default to null]
**AbuseReports** | **int32** | The number of abuse reports generated for this campaign. | [optional] [default to null]
**Unsubscribed** | **int32** | The total number of unsubscribed members for this campaign. | [optional] [default to null]
**SendTime** | [**time.Time**](time.Time.md) | The date and time a campaign was sent in ISO 8601 format. | [optional] [default to null]
**RssLastSend** | [**time.Time**](time.Time.md) | For RSS campaigns, the date and time of the last send in ISO 8601 format. | [optional] [default to null]
**Bounces** | [***Bounces**](Bounces.md) |  | [optional] [default to null]
**Forwards** | [***Forwards**](Forwards.md) |  | [optional] [default to null]
**Opens** | [***Opens**](Opens.md) |  | [optional] [default to null]
**Clicks** | [***Clicks**](Clicks.md) |  | [optional] [default to null]
**FacebookLikes** | [***FacebookLikes**](Facebook Likes.md) |  | [optional] [default to null]
**IndustryStats** | [***IndustryStats1**](Industry Stats_1.md) |  | [optional] [default to null]
**ListStats** | [***ListStats**](List Stats.md) |  | [optional] [default to null]
**AbSplit** | [***AbSplitStats**](AB Split Stats.md) |  | [optional] [default to null]
**Timewarp** | [**[]CampaignReports1Timewarp**](Campaign Reports_1_timewarp.md) | An hourly breakdown of sends, opens, and clicks if a campaign is sent using timewarp. | [optional] [default to null]
**Timeseries** | [**[]CampaignReports1Timeseries**](Campaign Reports_1_timeseries.md) | An hourly breakdown of the performance of the campaign over the first 24 hours. | [optional] [default to null]
**ShareReport** | [***ShareReport**](Share Report.md) |  | [optional] [default to null]
**Ecommerce** | [***ECommerceReport1**](ECommerce Report_1.md) |  | [optional] [default to null]
**DeliveryStatus** | [***CampaignDeliveryStatus**](Campaign Delivery Status.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


