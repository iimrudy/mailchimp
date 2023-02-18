# ListMembers1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**UniqueEmailId** | **string** | An identifier for the address across all of Mailchimp. | [optional] [default to null]
**EmailType** | **string** | Type of email this member asked to get (&#39;html&#39; or &#39;text&#39;). | [optional] [default to null]
**Status** | **string** | Subscriber&#39;s current status. | [optional] [default to null]
**MergeFields** | **map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Interests** | **map[string]bool** | The key of this object&#39;s properties is the ID of the interest in question. | [optional] [default to null]
**Stats** | [***SubscriberStats**](Subscriber Stats.md) |  | [optional] [default to null]
**IpSignup** | **string** | IP address the subscriber signed up from. | [optional] [default to null]
**TimestampSignup** | [**time.Time**](time.Time.md) | The date and time the subscriber signed up for the list in ISO 8601 format. | [optional] [default to null]
**IpOpt** | **string** | The IP address the subscriber used to confirm their opt-in status. | [optional] [default to null]
**TimestampOpt** | [**time.Time**](time.Time.md) | The date and time the subscriber confirmed their opt-in status in ISO 8601 format. | [optional] [default to null]
**MemberRating** | **int32** | Star rating for this member, between 1 and 5. | [optional] [default to null]
**LastChanged** | [**time.Time**](time.Time.md) | The date and time the member&#39;s info was last changed in ISO 8601 format. | [optional] [default to null]
**Language** | **string** | If set/detected, the [subscriber&#39;s language](https://mailchimp.com/help/view-and-edit-contact-languages/). | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**EmailClient** | **string** | The list member&#39;s email client. | [optional] [default to null]
**Location** | [***Location1**](Location_1.md) |  | [optional] [default to null]
**LastNote** | [***Notes**](Notes.md) |  | [optional] [default to null]
**ListId** | **string** | The list id. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


