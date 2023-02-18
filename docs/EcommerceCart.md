# EcommerceCart

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the cart. | [optional] [default to null]
**Customer** | [***EcommerceCustomer**](Ecommerce Customer.md) |  | [optional] [default to null]
**CampaignId** | **string** | A string that uniquely identifies the campaign associated with a cart. | [optional] [default to null]
**CheckoutUrl** | **string** | The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations. | [optional] [default to null]
**CurrencyCode** | **string** | The three-letter ISO 4217 code for the currency that the cart uses. | [optional] [default to null]
**OrderTotal** | **float32** | The order total for the cart. | [optional] [default to null]
**TaxTotal** | **float32** | The total tax for the cart. | [optional] [default to null]
**Lines** | [**[]EcommerceCartLineItem**](Ecommerce Cart Line Item.md) | An array of the cart&#39;s line items. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the cart was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the cart was last updated in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


