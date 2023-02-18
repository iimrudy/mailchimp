# MemberNotes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The note id. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the note was created in ISO 8601 format. | [optional] [default to null]
**CreatedBy** | **string** | The author of the note. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the note was last updated in ISO 8601 format. | [optional] [default to null]
**Note** | **string** | The content of the note. | [optional] [default to null]
**ListId** | **string** | The unique id for the list. | [optional] [default to null]
**EmailId** | **string** | The MD5 hash of the lowercase version of the list member&#39;s email address. | [optional] [default to null]
**ContactId** | **string** | As Mailchimp evolves beyond email, you may eventually have contacts without email addresses. While the &#x60;email_id&#x60; is the MD5 hash of their email address, this &#x60;contact_id&#x60; is agnostic of contact’s inclusion of an email address. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


