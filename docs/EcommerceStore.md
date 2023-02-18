# EcommerceStore

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the store. | [optional] [default to null]
**ListId** | **string** | The unique identifier for the list that&#39;s associated with the store. The &#x60;list_id&#x60; for a specific store can&#39;t change. | [optional] [default to null]
**Name** | **string** | The name of the store. | [optional] [default to null]
**Platform** | **string** | The e-commerce platform of the store. | [optional] [default to null]
**Domain** | **string** | The store domain.  The store domain must be unique within a user account. | [optional] [default to null]
**IsSyncing** | **bool** | Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations). | [optional] [default to null]
**EmailAddress** | **string** | The email address for the store. | [optional] [default to null]
**CurrencyCode** | **string** | The three-letter ISO 4217 code for the currency that the store accepts. | [optional] [default to null]
**MoneyFormat** | **string** | The currency format for the store. For example: &#x60;$&#x60;, &#x60;£&#x60;, etc. | [optional] [default to null]
**PrimaryLocale** | **string** | The primary locale for the store. For example: &#x60;en&#x60;, &#x60;de&#x60;, etc. | [optional] [default to null]
**Timezone** | **string** | The timezone for the store. | [optional] [default to null]
**Phone** | **string** | The store phone number. | [optional] [default to null]
**Address** | [***Address1**](Address_1.md) |  | [optional] [default to null]
**ConnectedSite** | [***ConnectedSite2**](Connected Site_2.md) |  | [optional] [default to null]
**Automations** | [***Automations**](Automations.md) |  | [optional] [default to null]
**ListIsActive** | **bool** | The status of the list connected to the store, namely if it&#39;s deleted or disabled. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the store was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the store was last updated in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


