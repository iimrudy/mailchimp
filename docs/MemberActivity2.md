# MemberActivity2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | One of the following actions: &#39;open&#39;, &#39;click&#39;, or &#39;bounce&#39; | [optional] [default to null]
**Type_** | **string** | If the action is a &#39;bounce&#39;, the type of bounce received: &#39;hard&#39;, &#39;soft&#39;. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | The date and time recorded for the action in ISO 8601 format. | [optional] [default to null]
**Url** | **string** | If the action is a &#39;click&#39;, the URL on which the member clicked. | [optional] [default to null]
**Ip** | **string** | The IP address recorded for the action. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


