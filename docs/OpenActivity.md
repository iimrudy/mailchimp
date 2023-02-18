# OpenActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** | The unique id for the campaign. | [optional] [default to null]
**ListId** | **string** | The unique id for the list. | [optional] [default to null]
**ListIsActive** | **bool** | The status of the list used, namely if it&#39;s deleted or disabled. | [optional] [default to null]
**ContactStatus** | **string** | The status of the member, namely if they are subscribed, unsubscribed, deleted, non-subscribed, transactional, pending, or need reconfirmation. | [optional] [default to null]
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**MergeFields** | **map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**OpensCount** | **int32** | The total number of times the this campaign was opened by the list member. | [optional] [default to null]
**Opens** | [**[]MemberActivity1**](Member Activity_1.md) | An array of timestamps for each time a list member opened the campaign. If a list member opens an email multiple times, this will return a separate timestamp for each open event. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


