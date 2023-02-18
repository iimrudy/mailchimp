# BatchAddremoveListMembersTofromStaticSegment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembersAdded** | [**[]ListMembers**](List Members.md) | An array of objects, each representing a new member that was added to the static segment. | [optional] [default to null]
**MembersRemoved** | [**[]ListMembers**](List Members.md) | An array of objects, each representing an existing list member that got deleted from the static segment. | [optional] [default to null]
**Errors** | [**[]BatchAddremoveListMembersTofromStaticSegmentErrors**](Batch addremove List members tofrom static segment_errors.md) | An array of objects, each representing an array of email addresses that could not be added to the segment or removed and an error message providing more details. | [optional] [default to null]
**TotalAdded** | **int32** | The total number of items matching the query, irrespective of pagination. | [optional] [default to null]
**TotalRemoved** | **int32** | The total number of items matching the query, irrespective of pagination. | [optional] [default to null]
**ErrorCount** | **int32** | The total number of items matching the query, irrespective of pagination. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


