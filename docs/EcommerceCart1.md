# EcommerceCart1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the cart. | [default to null]
**Customer** | [***EcommerceCustomer1**](Ecommerce Customer_1.md) |  | [default to null]
**CampaignId** | **string** | A string that uniquely identifies the campaign for a cart. | [optional] [default to null]
**CheckoutUrl** | **string** | The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations. | [optional] [default to null]
**CurrencyCode** | **string** | The three-letter ISO 4217 code for the currency that the cart uses. | [default to null]
**OrderTotal** | **float32** | The order total for the cart. | [default to null]
**TaxTotal** | **float32** | The total tax for the cart. | [optional] [default to null]
**Lines** | [**[]EcommerceCartLineItem1**](Ecommerce Cart Line Item_1.md) | An array of the cart&#39;s line items. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


