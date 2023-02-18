# AbuseComplaint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id for the abuse report | [optional] [default to null]
**CampaignId** | **string** | The campaign id for the abuse report | [optional] [default to null]
**ListId** | **string** | The list id for the abuse report. | [optional] [default to null]
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | [optional] [default to null]
**EmailAddress** | **string** | Email address for a subscriber. | [optional] [default to null]
**MergeFields** | **map[string]interface{}** | A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure. | [optional] [default to null]
**Vip** | **bool** | [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber. | [optional] [default to null]
**Date** | **string** | Date for the abuse report | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


