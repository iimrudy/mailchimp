# Events

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for this type of event (&#39;purchased&#39;, &#39;visited&#39;, etc). Must be 2-30 characters in length | [default to null]
**Properties** | **map[string]string** | An optional list of properties | [optional] [default to null]
**IsSyncing** | **bool** | Events created with the is_syncing value set to &#x60;true&#x60; will not trigger automations. | [optional] [default to null]
**OccurredAt** | [**time.Time**](time.Time.md) | The date and time the event occurred in ISO 8601 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


