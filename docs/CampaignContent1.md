# CampaignContent1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlainText** | **string** | The plain-text portion of the campaign. If left unspecified, we&#39;ll generate this automatically. | [optional] [default to null]
**Html** | **string** | The raw HTML for the campaign. | [optional] [default to null]
**Url** | **string** | When importing a campaign, the URL where the HTML lives. | [optional] [default to null]
**Template** | [***TemplateContent**](Template Content.md) |  | [optional] [default to null]
**Archive** | [***UploadArchive**](Upload Archive.md) |  | [optional] [default to null]
**VariateContents** | [**[]CampaignscampaignIdcontentVariateContents**](campaignscampaign_idcontent_variate_contents.md) | Content options for [Multivariate Campaigns](https://mailchimp.com/help/about-multivariate-campaigns/). Each content option must provide HTML content and may optionally provide plain text. For campaigns not testing content, only one object should be provided. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


