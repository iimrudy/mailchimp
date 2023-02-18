# Statistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberCount** | **int32** | The number of active members in the list. | [optional] [default to null]
**TotalContacts** | **int32** | The number of contacts in the list, including subscribed, unsubscribed, pending, cleaned, deleted, transactional, and those that need to be reconfirmed. Requires include_total_contacts query parameter to be included. | [optional] [default to null]
**UnsubscribeCount** | **int32** | The number of members who have unsubscribed from the list. | [optional] [default to null]
**CleanedCount** | **int32** | The number of members cleaned from the list. | [optional] [default to null]
**MemberCountSinceSend** | **int32** | The number of active members in the list since the last campaign was sent. | [optional] [default to null]
**UnsubscribeCountSinceSend** | **int32** | The number of members who have unsubscribed since the last campaign was sent. | [optional] [default to null]
**CleanedCountSinceSend** | **int32** | The number of members cleaned from the list since the last campaign was sent. | [optional] [default to null]
**CampaignCount** | **int32** | The number of campaigns in any status that use this list. | [optional] [default to null]
**CampaignLastSent** | [**time.Time**](time.Time.md) | The date and time the last campaign was sent to this list in ISO 8601 format. This is updated when a campaign is sent to 10 or more recipients. | [optional] [default to null]
**MergeFieldCount** | **int32** | The number of merge fields ([audience field](https://mailchimp.com/help/getting-started-with-merge-tags/)) for this list (doesn&#39;t include EMAIL). | [optional] [default to null]
**AvgSubRate** | **float32** | The average number of subscriptions per month for the list (not returned if we haven&#39;t calculated it yet). | [optional] [default to null]
**AvgUnsubRate** | **float32** | The average number of unsubscriptions per month for the list (not returned if we haven&#39;t calculated it yet). | [optional] [default to null]
**TargetSubRate** | **float32** | The target number of subscriptions per month for the list to keep it growing (not returned if we haven&#39;t calculated it yet). | [optional] [default to null]
**OpenRate** | **float32** | The average open rate (a percentage represented as a number between 0 and 100) per campaign for the list (not returned if we haven&#39;t calculated it yet). | [optional] [default to null]
**ClickRate** | **float32** | The average click rate (a percentage represented as a number between 0 and 100) per campaign for the list (not returned if we haven&#39;t calculated it yet). | [optional] [default to null]
**LastSubDate** | [**time.Time**](time.Time.md) | The date and time of the last time someone subscribed to this list in ISO 8601 format. | [optional] [default to null]
**LastUnsubDate** | [**time.Time**](time.Time.md) | The date and time of the last time someone unsubscribed from this list in ISO 8601 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


