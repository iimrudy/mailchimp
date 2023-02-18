# EcommerceProduct1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the product. | [default to null]
**Title** | **string** | The title of a product. | [default to null]
**Handle** | **string** | The handle of a product. | [optional] [default to null]
**Url** | **string** | The URL for a product. | [optional] [default to null]
**Description** | **string** | The description of a product. | [optional] [default to null]
**Type_** | **string** | The type of product. | [optional] [default to null]
**Vendor** | **string** | The vendor for a product. | [optional] [default to null]
**ImageUrl** | **string** | The image URL for a product. | [optional] [default to null]
**Variants** | [**[]EcommerceProductVariant1**](Ecommerce Product Variant_1.md) | An array of the product&#39;s variants. At least one variant is required for each product. A variant can use the same &#x60;id&#x60; and &#x60;title&#x60; as the parent product. | [default to null]
**Images** | [**[]EcommerceProductImage1**](Ecommerce Product Image_1.md) | An array of the product&#39;s images. | [optional] [default to null]
**PublishedAtForeign** | [**time.Time**](time.Time.md) | The date and time the product was published. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


