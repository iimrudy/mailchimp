# EcommercePromoRule1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the promo rule. If Ecommerce platform does not support promo rule, use promo code id as promo rule id. Restricted to UTF-8 characters with max length 50. | [default to null]
**Title** | **string** | The title that will show up in promotion campaign. Restricted to UTF-8 characters with max length of 100 bytes. | [optional] [default to null]
**Description** | **string** | The description of a promotion restricted to UTF-8 characters with max length 255. | [default to null]
**StartsAt** | [**time.Time**](time.Time.md) | The date and time when the promotion is in effect in ISO 8601 format. | [optional] [default to null]
**EndsAt** | **string** | The date and time when the promotion ends. Must be after starts_at and in ISO 8601 format. | [optional] [default to null]
**Amount** | **float32** | The amount of the promo code discount. If &#39;type&#39; is &#39;fixed&#39;, the amount is treated as a monetary value. If &#39;type&#39; is &#39;percentage&#39;, amount must be a decimal value between 0.0 and 1.0, inclusive. | [default to null]
**Type_** | **string** | Type of discount. For free shipping set type to fixed. | [default to null]
**Target** | **string** | The target that the discount applies to. | [default to null]
**Enabled** | **bool** | Whether the promo rule is currently enabled. | [optional] [default to null]
**CreatedAtForeign** | [**time.Time**](time.Time.md) | The date and time the promotion was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAtForeign** | [**time.Time**](time.Time.md) | The date and time the promotion was updated in ISO 8601 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


