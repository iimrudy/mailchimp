# EcommerceCart2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [***EcommerceCustomer2**](Ecommerce Customer_2.md) |  | [optional] [default to null]
**CampaignId** | **string** | A string that uniquely identifies the campaign associated with a cart. | [optional] [default to null]
**CheckoutUrl** | **string** | The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations. | [optional] [default to null]
**CurrencyCode** | **string** | The three-letter ISO 4217 code for the currency that the cart uses. | [optional] [default to null]
**OrderTotal** | **float32** | The order total for the cart. | [optional] [default to null]
**TaxTotal** | **float32** | The total tax for the cart. | [optional] [default to null]
**Lines** | [**[]EcommerceCartLineItem2**](Ecommerce Cart Line Item_2.md) | An array of the cart&#39;s line items. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


