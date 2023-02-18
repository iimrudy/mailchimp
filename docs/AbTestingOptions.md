# AbTestingOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SplitTest** | **string** | The type of AB split to run. | [optional] [default to null]
**PickWinner** | **string** | How we should evaluate a winner. Based on &#39;opens&#39;, &#39;clicks&#39;, or &#39;manual&#39;. | [optional] [default to null]
**WaitUnits** | **string** | How unit of time for measuring the winner (&#39;hours&#39; or &#39;days&#39;). This cannot be changed after a campaign is sent. | [optional] [default to null]
**WaitTime** | **int32** | The amount of time to wait before picking a winner. This cannot be changed after a campaign is sent. | [optional] [default to null]
**SplitSize** | **int32** | The size of the split groups. Campaigns split based on &#39;schedule&#39; are forced to have a 50/50 split. Valid split integers are between 1-50. | [optional] [default to null]
**FromNameA** | **string** | For campaigns split on &#39;From Name&#39;, the name for Group A. | [optional] [default to null]
**FromNameB** | **string** | For campaigns split on &#39;From Name&#39;, the name for Group B. | [optional] [default to null]
**ReplyEmailA** | **string** | For campaigns split on &#39;From Name&#39;, the reply-to address for Group A. | [optional] [default to null]
**ReplyEmailB** | **string** | For campaigns split on &#39;From Name&#39;, the reply-to address for Group B. | [optional] [default to null]
**SubjectA** | **string** | For campaigns split on &#39;Subject Line&#39;, the subject line for Group A. | [optional] [default to null]
**SubjectB** | **string** | For campaigns split on &#39;Subject Line&#39;, the subject line for Group B. | [optional] [default to null]
**SendTimeA** | [**time.Time**](time.Time.md) | The send time for Group A. | [optional] [default to null]
**SendTimeB** | [**time.Time**](time.Time.md) | The send time for Group B. | [optional] [default to null]
**SendTimeWinner** | **string** | The send time for the winning version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


