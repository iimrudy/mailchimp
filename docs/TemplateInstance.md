# TemplateInstance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The individual id for the template. | [optional] [default to null]
**Type_** | **string** | The type of template (user, base, or gallery). | [optional] [default to null]
**Name** | **string** | The name of the template. | [optional] [default to null]
**DragAndDrop** | **bool** | Whether the template uses the drag and drop editor. | [optional] [default to null]
**Responsive** | **bool** | Whether the template contains media queries to make it responsive. | [optional] [default to null]
**Category** | **string** | If available, the category the template is listed in. | [optional] [default to null]
**DateCreated** | [**time.Time**](time.Time.md) | The date and time the template was created in ISO 8601 format. | [optional] [default to null]
**DateEdited** | [**time.Time**](time.Time.md) | The date and time the template was edited in ISO 8601 format. | [optional] [default to null]
**CreatedBy** | **string** | The login name for template&#39;s creator. | [optional] [default to null]
**EditedBy** | **string** | The login name who last edited the template. | [optional] [default to null]
**Active** | **bool** | User templates are not &#39;deleted,&#39; but rather marked as &#39;inactive.&#39; Returns whether the template is still active. | [optional] [default to null]
**FolderId** | **string** | The id of the folder the template is currently in. | [optional] [default to null]
**Thumbnail** | **string** | If available, the URL for a thumbnail of the template. | [optional] [default to null]
**ShareUrl** | **string** | The URL used for [template sharing](https://mailchimp.com/help/share-a-template/). | [optional] [default to null]
**ContentType** | **string** | How the template&#39;s content is put together. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


