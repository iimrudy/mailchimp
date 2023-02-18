# EcommerceOrder2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [***EcommerceCustomer2**](Ecommerce Customer_2.md) |  | [optional] [default to null]
**CampaignId** | **string** | A string that uniquely identifies the campaign associated with an order. | [optional] [default to null]
**LandingSite** | **string** | The URL for the page where the buyer landed when entering the shop. | [optional] [default to null]
**FinancialStatus** | **string** | The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications). | [optional] [default to null]
**FulfillmentStatus** | **string** | The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications). | [optional] [default to null]
**CurrencyCode** | **string** | The three-letter ISO 4217 code for the currency that the store accepts. | [optional] [default to null]
**OrderTotal** | **float32** | The order total associated with an order. | [optional] [default to null]
**OrderUrl** | **string** | The URL for the order. | [optional] [default to null]
**DiscountTotal** | **float32** | The total amount of the discounts to be applied to the price of the order. | [optional] [default to null]
**TaxTotal** | **float32** | The tax total associated with an order. | [optional] [default to null]
**ShippingTotal** | **float32** | The shipping total for the order. | [optional] [default to null]
**TrackingCode** | **string** | The Mailchimp tracking code for the order. Uses the &#39;mc_tc&#39; parameter in E-Commerce tracking URLs. | [optional] [default to null]
**ProcessedAtForeign** | [**time.Time**](time.Time.md) | The date and time the order was processed in ISO 8601 format. | [optional] [default to null]
**CancelledAtForeign** | [**time.Time**](time.Time.md) | The date and time the order was cancelled in ISO 8601 format. Note: passing a value for this parameter will cancel the order being edited. | [optional] [default to null]
**UpdatedAtForeign** | [**time.Time**](time.Time.md) | The date and time the order was updated in ISO 8601 format. | [optional] [default to null]
**ShippingAddress** | [***ShippingAddress**](Shipping Address.md) |  | [optional] [default to null]
**BillingAddress** | [***BillingAddress**](Billing Address.md) |  | [optional] [default to null]
**Promos** | [**[]EcommercestoresstoreIdordersPromos**](ecommercestoresstore_idorders_promos.md) | The promo codes applied on the order. Note: Patch will completely replace the value of promos with the new one provided. | [optional] [default to null]
**Lines** | [**[]EcommerceOrderLineItem2**](Ecommerce Order Line Item_2.md) | An array of the order&#39;s line items. | [optional] [default to null]
**Outreach** | [***Outreach1**](Outreach_1.md) |  | [optional] [default to null]
**TrackingNumber** | **string** | The tracking number associated with the order. | [optional] [default to null]
**TrackingCarrier** | **string** | The tracking carrier associated with the order. | [optional] [default to null]
**TrackingUrl** | **string** | The tracking URL associated with the order. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


