# CreateAnAccountExport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeStages** | **[]string** | The stages of an account export to include. | [default to null]
**SinceTimestamp** | [**time.Time**](time.Time.md) | An ISO 8601 date that will limit the export to only records created after a given time. For instance, the reports stage will contain any campaign sent after the given timestamp. Audiences, however, are excluded from this limit. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


