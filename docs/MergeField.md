# MergeField

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeId** | **int32** | An unchanging id for the merge field. | [optional] [default to null]
**Tag** | **string** | The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts). | [optional] [default to null]
**Name** | **string** | The name of the merge field (audience field). | [optional] [default to null]
**Type_** | **string** | The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field. | [optional] [default to null]
**Required** | **bool** | The boolean value if the merge field is required. | [optional] [default to null]
**DefaultValue** | **string** | The default value for the merge field if &#x60;null&#x60;. | [optional] [default to null]
**Public** | **bool** | Whether the merge field is displayed on the signup form. | [optional] [default to null]
**DisplayOrder** | **int32** | The order that the merge field displays on the list signup form. | [optional] [default to null]
**Options** | [***MergeFieldOptions**](Merge Field Options.md) |  | [optional] [default to null]
**HelpText** | **string** | Extra text to help the subscriber fill out the form. | [optional] [default to null]
**ListId** | **string** | The ID that identifies this merge field&#39;s audience&#39;. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


