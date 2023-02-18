# EcommercePromoCode1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the promo code. Restricted to UTF-8 characters with max length 50. | [default to null]
**Code** | **string** | The discount code. Restricted to UTF-8 characters with max length 50. | [default to null]
**RedemptionUrl** | **string** | The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000. | [default to null]
**UsageCount** | **int32** | Number of times promo code has been used. | [optional] [default to null]
**Enabled** | **bool** | Whether the promo code is currently enabled. | [optional] [default to null]
**CreatedAtForeign** | [**time.Time**](time.Time.md) | The date and time the promotion was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAtForeign** | [**time.Time**](time.Time.md) | The date and time the promotion was updated in ISO 8601 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


