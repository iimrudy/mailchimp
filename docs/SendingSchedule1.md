# SendingSchedule1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hour** | **int32** | The hour to send the campaign in local time. Acceptable hours are 0-23. For example, &#39;4&#39; would be 4am in [your account&#39;s default time zone](https://mailchimp.com/help/set-account-details/). | [optional] [default to null]
**DailySend** | [***DailySendingDays**](Daily Sending Days.md) |  | [optional] [default to null]
**WeeklySendDay** | **string** | The day of the week to send a weekly RSS Campaign. | [optional] [default to null]
**MonthlySendDate** | **float32** | The day of the month to send a monthly RSS Campaign. Acceptable days are 0-31, where &#39;0&#39; is always the last day of a month. Months with fewer than the selected number of days will not have an RSS campaign sent out that day. For example, RSS Campaigns set to send on the 30th will not go out in February. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


