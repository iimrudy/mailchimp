# EcommerceProduct

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the product. | [optional] [default to null]
**CurrencyCode** | **string** | The currency code | [optional] [default to null]
**Title** | **string** | The title of a product. | [optional] [default to null]
**Handle** | **string** | The handle of a product. | [optional] [default to null]
**Url** | **string** | The URL for a product. | [optional] [default to null]
**Description** | **string** | The description of a product. | [optional] [default to null]
**Type_** | **string** | The type of product. | [optional] [default to null]
**Vendor** | **string** | The vendor for a product. | [optional] [default to null]
**ImageUrl** | **string** | The image URL for a product. | [optional] [default to null]
**Variants** | [**[]EcommerceProductVariant**](Ecommerce Product Variant.md) | Returns up to 50 of the product&#39;s variants. To retrieve all variants use [Product Variants](https://mailchimp.com/developer/marketing/api/ecommerce-product-variants/). | [optional] [default to null]
**Images** | [**[]EcommerceProductImage**](Ecommerce Product Image.md) | An array of the product&#39;s images. | [optional] [default to null]
**PublishedAtForeign** | [**time.Time**](time.Time.md) | The date and time the product was published in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


