# CampaignSettings2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectLine** | **string** | The subject line for the campaign. | [optional] [default to null]
**PreviewText** | **string** | The preview text for the campaign. | [optional] [default to null]
**Title** | **string** | The title of the campaign. | [optional] [default to null]
**FromName** | **string** | The &#39;from&#39; name on the campaign (not an email address). | [optional] [default to null]
**ReplyTo** | **string** | The reply-to email address for the campaign. | [optional] [default to null]
**UseConversation** | **bool** | Use Mailchimp Conversation feature to manage out-of-office replies. | [optional] [default to null]
**ToName** | **string** | The campaign&#39;s custom &#39;To&#39; name. Typically the first name [audience field](https://mailchimp.com/help/getting-started-with-merge-tags/). | [optional] [default to null]
**FolderId** | **string** | If the campaign is listed in a folder, the id for that folder. | [optional] [default to null]
**Authenticate** | **bool** | Whether Mailchimp [authenticated](https://mailchimp.com/help/about-email-authentication/) the campaign. Defaults to &#x60;true&#x60;. | [optional] [default to null]
**AutoFooter** | **bool** | Automatically append Mailchimp&#39;s [default footer](https://mailchimp.com/help/about-campaign-footers/) to the campaign. | [optional] [default to null]
**InlineCss** | **bool** | Automatically inline the CSS included with the campaign content. | [optional] [default to null]
**AutoTweet** | **bool** | Automatically tweet a link to the [campaign archive](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) page when the campaign is sent. | [optional] [default to null]
**AutoFbPost** | **[]string** | An array of [Facebook](https://mailchimp.com/help/connect-or-disconnect-the-facebook-integration/) page ids to auto-post to. | [optional] [default to null]
**FbComments** | **bool** | Allows Facebook comments on the campaign (also force-enables the Campaign Archive toolbar). Defaults to &#x60;true&#x60;. | [optional] [default to null]
**Timewarp** | **bool** | Send this campaign using [Timewarp](https://mailchimp.com/help/use-timewarp/). | [optional] [default to null]
**TemplateId** | **int32** | The id for the template used in this campaign. | [optional] [default to null]
**DragAndDrop** | **bool** | Whether the campaign uses the drag-and-drop editor. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


