# SubscriberList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this list. | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. View this list in your Mailchimp account at &#x60;https://{dc}.admin.mailchimp.com/lists/members/?id&#x3D;{web_id}&#x60;. | [optional] [default to null]
**Name** | **string** | The name of the list. | [optional] [default to null]
**Contact** | [***ListContact**](List Contact.md) |  | [optional] [default to null]
**PermissionReminder** | **string** | The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list. | [optional] [default to null]
**UseArchiveBar** | **bool** | Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default. | [optional] [default to null]
**CampaignDefaults** | [***CampaignDefaults**](Campaign Defaults.md) |  | [optional] [default to null]
**NotifyOnSubscribe** | **string** | The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to. | [optional] [default to null]
**NotifyOnUnsubscribe** | **string** | The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to. | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) | The date and time that this list was created in ISO 8601 format. | [optional] [default to null]
**ListRating** | **int32** | An auto-generated activity score for the list (0-5). | [optional] [default to null]
**EmailTypeOption** | **bool** | Whether the list supports [multiple formats for emails](https://mailchimp.com/help/change-audience-name-defaults/). When set to &#x60;true&#x60;, subscribers can choose whether they want to receive HTML or plain-text emails. When set to &#x60;false&#x60;, subscribers will receive HTML emails, with a plain-text alternative backup. | [optional] [default to null]
**SubscribeUrlShort** | **string** | Our [url shortened](https://mailchimp.com/help/share-your-signup-form/) version of this list&#39;s subscribe form. | [optional] [default to null]
**SubscribeUrlLong** | **string** | The full version of this list&#39;s subscribe form (host will vary). | [optional] [default to null]
**BeamerAddress** | **string** | The list&#39;s [Email Beamer](https://mailchimp.com/help/use-email-beamer-to-create-a-campaign/) address. | [optional] [default to null]
**Visibility** | **string** | Legacy - visibility settings are no longer used | [optional] [default to null]
**DoubleOptin** | **bool** | Whether or not to require the subscriber to confirm subscription via email. | [optional] [default to null]
**HasWelcome** | **bool** | Whether or not this list has a welcome automation connected. Welcome Automations: welcomeSeries, singleWelcome, emailFollowup. | [optional] [default to null]
**MarketingPermissions** | **bool** | Whether or not the list has marketing permissions (eg. GDPR) enabled. | [optional] [default to null]
**Modules** | **[]string** | Any list-specific modules installed for this list. | [optional] [default to null]
**Stats** | [***Statistics**](Statistics.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


