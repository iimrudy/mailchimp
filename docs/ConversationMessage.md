# ConversationMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this message | [optional] [default to null]
**ConversationId** | **string** | A string that identifies this message&#39;s conversation | [optional] [default to null]
**ListId** | **int32** | The list&#39;s web ID | [optional] [default to null]
**FromLabel** | **string** | A label representing the sender of this message | [optional] [default to null]
**FromEmail** | **string** | A label representing the email of the sender of this message | [optional] [default to null]
**Subject** | **string** | The subject of this message | [optional] [default to null]
**Message** | **string** | The plain-text content of the message | [optional] [default to null]
**Read** | **bool** | Whether this message has been marked as read | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The date and time the message was either sent or received in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


