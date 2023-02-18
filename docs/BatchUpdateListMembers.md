# BatchUpdateListMembers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewMembers** | [**[]ListMembers**](List Members.md) | An array of objects, each representing a new member that was added to the list. | [optional] [default to null]
**UpdatedMembers** | [**[]ListMembers**](List Members.md) | An array of objects, each representing an existing list member whose subscription status was updated. | [optional] [default to null]
**Errors** | [**[]BatchUpdateListMembersErrors**](Batch update List members_errors.md) | An array of objects, each representing an email address that could not be added to the list or updated and an error message providing more details. | [optional] [default to null]
**TotalCreated** | **int32** | The total number of items matching the query, irrespective of pagination. | [optional] [default to null]
**TotalUpdated** | **int32** | The total number of items matching the query, irrespective of pagination. | [optional] [default to null]
**ErrorCount** | **int32** | The total number of items matching the query, irrespective of pagination. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


