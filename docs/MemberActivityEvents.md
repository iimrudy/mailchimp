# MemberActivityEvents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | [**[]MemberActivity**](Member Activity.md) | An array of objects, each representing a member event. | [optional] [default to null]
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | [optional] [default to null]
**ContactId** | **string** | As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the &#x60;email_id&#x60; is the MD5 hash of their email address, this &#x60;contact_id&#x60; is agnostic of contact’s inclusion of an email address. | [optional] [default to null]
**ListId** | **string** | The list id. | [optional] [default to null]
**TotalItems** | **int32** | The total number of items matching the query regardless of pagination. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


