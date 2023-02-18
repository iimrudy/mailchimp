# InlineResponse2001Exports

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportId** | **int32** | The ID for the export. | [optional] [default to null]
**Started** | [**time.Time**](time.Time.md) | Start time for the export. | [optional] [default to null]
**Finished** | [**time.Time**](time.Time.md) | If finished, the finish time for the export. | [optional] [default to null]
**SizeInBytes** | **int32** | The size of the uncompressed export in bytes. | [optional] [default to null]
**DownloadUrl** | **string** | If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


