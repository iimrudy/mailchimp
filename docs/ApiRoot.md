# ApiRoot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Mailchimp account ID. | [optional] [default to null]
**LoginId** | **string** | The ID associated with the user who owns this API key. If you can login to multiple accounts, this ID will be the same for each account. | [optional] [default to null]
**AccountName** | **string** | The name of the account. | [optional] [default to null]
**Email** | **string** | The account email address. | [optional] [default to null]
**FirstName** | **string** | The first name tied to the account. | [optional] [default to null]
**LastName** | **string** | The last name tied to the account. | [optional] [default to null]
**Username** | **string** | The username tied to the account. | [optional] [default to null]
**AvatarUrl** | **string** | URL of the avatar for the user. | [optional] [default to null]
**Role** | **string** | The [user role](https://mailchimp.com/help/manage-user-levels-in-your-account/) for the account. | [optional] [default to null]
**MemberSince** | [**time.Time**](time.Time.md) | The date and time that the account was created in ISO 8601 format. | [optional] [default to null]
**PricingPlanType** | **string** | The type of pricing plan the account is on. | [optional] [default to null]
**FirstPayment** | [**time.Time**](time.Time.md) | Date of first payment for monthly plans. | [optional] [default to null]
**AccountTimezone** | **string** | The timezone currently set for the account. | [optional] [default to null]
**AccountIndustry** | **string** | The user-specified industry associated with the account. | [optional] [default to null]
**Contact** | [***AccountContact**](Account Contact.md) |  | [optional] [default to null]
**ProEnabled** | **bool** | Legacy - whether the account includes [Mailchimp Pro](https://mailchimp.com/help/about-mailchimp-pro/). | [optional] [default to null]
**LastLogin** | [**time.Time**](time.Time.md) | The date and time of the last login for this account in ISO 8601 format. | [optional] [default to null]
**TotalSubscribers** | **int32** | The total number of subscribers across all lists in the account. | [optional] [default to null]
**IndustryStats** | [***IndustryStats**](Industry Stats.md) |  | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


