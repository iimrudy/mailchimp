# \EcommerceApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEcommerceStoresId**](EcommerceApi.md#DeleteEcommerceStoresId) | **Delete** /ecommerce/stores/{store_id} | Delete store
[**DeleteEcommerceStoresIdCartsId**](EcommerceApi.md#DeleteEcommerceStoresIdCartsId) | **Delete** /ecommerce/stores/{store_id}/carts/{cart_id} | Delete cart
[**DeleteEcommerceStoresIdCartsLinesId**](EcommerceApi.md#DeleteEcommerceStoresIdCartsLinesId) | **Delete** /ecommerce/stores/{store_id}/carts/{cart_id}/lines/{line_id} | Delete cart line item
[**DeleteEcommerceStoresIdCustomersId**](EcommerceApi.md#DeleteEcommerceStoresIdCustomersId) | **Delete** /ecommerce/stores/{store_id}/customers/{customer_id} | Delete customer
[**DeleteEcommerceStoresIdOrdersId**](EcommerceApi.md#DeleteEcommerceStoresIdOrdersId) | **Delete** /ecommerce/stores/{store_id}/orders/{order_id} | Delete order
[**DeleteEcommerceStoresIdOrdersIdLinesId**](EcommerceApi.md#DeleteEcommerceStoresIdOrdersIdLinesId) | **Delete** /ecommerce/stores/{store_id}/orders/{order_id}/lines/{line_id} | Delete order line item
[**DeleteEcommerceStoresIdProductsId**](EcommerceApi.md#DeleteEcommerceStoresIdProductsId) | **Delete** /ecommerce/stores/{store_id}/products/{product_id} | Delete product
[**DeleteEcommerceStoresIdProductsIdImagesId**](EcommerceApi.md#DeleteEcommerceStoresIdProductsIdImagesId) | **Delete** /ecommerce/stores/{store_id}/products/{product_id}/images/{image_id} | Delete product image
[**DeleteEcommerceStoresIdProductsIdVariantsId**](EcommerceApi.md#DeleteEcommerceStoresIdProductsIdVariantsId) | **Delete** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Delete product variant
[**DeleteEcommerceStoresIdPromocodesId**](EcommerceApi.md#DeleteEcommerceStoresIdPromocodesId) | **Delete** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes/{promo_code_id} | Delete promo code
[**DeleteEcommerceStoresIdPromorulesId**](EcommerceApi.md#DeleteEcommerceStoresIdPromorulesId) | **Delete** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id} | Delete promo rule
[**GetEcommerceOrders**](EcommerceApi.md#GetEcommerceOrders) | **Get** /ecommerce/orders | List account orders
[**GetEcommerceStores**](EcommerceApi.md#GetEcommerceStores) | **Get** /ecommerce/stores | List stores
[**GetEcommerceStoresId**](EcommerceApi.md#GetEcommerceStoresId) | **Get** /ecommerce/stores/{store_id} | Get store info
[**GetEcommerceStoresIdCarts**](EcommerceApi.md#GetEcommerceStoresIdCarts) | **Get** /ecommerce/stores/{store_id}/carts | List carts
[**GetEcommerceStoresIdCartsId**](EcommerceApi.md#GetEcommerceStoresIdCartsId) | **Get** /ecommerce/stores/{store_id}/carts/{cart_id} | Get cart info
[**GetEcommerceStoresIdCartsIdLines**](EcommerceApi.md#GetEcommerceStoresIdCartsIdLines) | **Get** /ecommerce/stores/{store_id}/carts/{cart_id}/lines | List cart line items
[**GetEcommerceStoresIdCartsIdLinesId**](EcommerceApi.md#GetEcommerceStoresIdCartsIdLinesId) | **Get** /ecommerce/stores/{store_id}/carts/{cart_id}/lines/{line_id} | Get cart line item
[**GetEcommerceStoresIdCustomers**](EcommerceApi.md#GetEcommerceStoresIdCustomers) | **Get** /ecommerce/stores/{store_id}/customers | List customers
[**GetEcommerceStoresIdCustomersId**](EcommerceApi.md#GetEcommerceStoresIdCustomersId) | **Get** /ecommerce/stores/{store_id}/customers/{customer_id} | Get customer info
[**GetEcommerceStoresIdOrders**](EcommerceApi.md#GetEcommerceStoresIdOrders) | **Get** /ecommerce/stores/{store_id}/orders | List orders
[**GetEcommerceStoresIdOrdersId**](EcommerceApi.md#GetEcommerceStoresIdOrdersId) | **Get** /ecommerce/stores/{store_id}/orders/{order_id} | Get order info
[**GetEcommerceStoresIdOrdersIdLines**](EcommerceApi.md#GetEcommerceStoresIdOrdersIdLines) | **Get** /ecommerce/stores/{store_id}/orders/{order_id}/lines | List order line items
[**GetEcommerceStoresIdOrdersIdLinesId**](EcommerceApi.md#GetEcommerceStoresIdOrdersIdLinesId) | **Get** /ecommerce/stores/{store_id}/orders/{order_id}/lines/{line_id} | Get order line item
[**GetEcommerceStoresIdProducts**](EcommerceApi.md#GetEcommerceStoresIdProducts) | **Get** /ecommerce/stores/{store_id}/products | List product
[**GetEcommerceStoresIdProductsId**](EcommerceApi.md#GetEcommerceStoresIdProductsId) | **Get** /ecommerce/stores/{store_id}/products/{product_id} | Get product info
[**GetEcommerceStoresIdProductsIdImages**](EcommerceApi.md#GetEcommerceStoresIdProductsIdImages) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/images | List product images
[**GetEcommerceStoresIdProductsIdImagesId**](EcommerceApi.md#GetEcommerceStoresIdProductsIdImagesId) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/images/{image_id} | Get product image info
[**GetEcommerceStoresIdProductsIdVariants**](EcommerceApi.md#GetEcommerceStoresIdProductsIdVariants) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/variants | List product variants
[**GetEcommerceStoresIdProductsIdVariantsId**](EcommerceApi.md#GetEcommerceStoresIdProductsIdVariantsId) | **Get** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Get product variant info
[**GetEcommerceStoresIdPromocodes**](EcommerceApi.md#GetEcommerceStoresIdPromocodes) | **Get** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes | List promo codes
[**GetEcommerceStoresIdPromocodesId**](EcommerceApi.md#GetEcommerceStoresIdPromocodesId) | **Get** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes/{promo_code_id} | Get promo code
[**GetEcommerceStoresIdPromorules**](EcommerceApi.md#GetEcommerceStoresIdPromorules) | **Get** /ecommerce/stores/{store_id}/promo-rules | List promo rules
[**GetEcommerceStoresIdPromorulesId**](EcommerceApi.md#GetEcommerceStoresIdPromorulesId) | **Get** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id} | Get promo rule
[**PatchEcommerceStoresId**](EcommerceApi.md#PatchEcommerceStoresId) | **Patch** /ecommerce/stores/{store_id} | Update store
[**PatchEcommerceStoresIdCartsId**](EcommerceApi.md#PatchEcommerceStoresIdCartsId) | **Patch** /ecommerce/stores/{store_id}/carts/{cart_id} | Update cart
[**PatchEcommerceStoresIdCartsIdLinesId**](EcommerceApi.md#PatchEcommerceStoresIdCartsIdLinesId) | **Patch** /ecommerce/stores/{store_id}/carts/{cart_id}/lines/{line_id} | Update cart line item
[**PatchEcommerceStoresIdCustomersId**](EcommerceApi.md#PatchEcommerceStoresIdCustomersId) | **Patch** /ecommerce/stores/{store_id}/customers/{customer_id} | Update customer
[**PatchEcommerceStoresIdOrdersId**](EcommerceApi.md#PatchEcommerceStoresIdOrdersId) | **Patch** /ecommerce/stores/{store_id}/orders/{order_id} | Update order
[**PatchEcommerceStoresIdOrdersIdLinesId**](EcommerceApi.md#PatchEcommerceStoresIdOrdersIdLinesId) | **Patch** /ecommerce/stores/{store_id}/orders/{order_id}/lines/{line_id} | Update order line item
[**PatchEcommerceStoresIdProductsId**](EcommerceApi.md#PatchEcommerceStoresIdProductsId) | **Patch** /ecommerce/stores/{store_id}/products/{product_id} | Update product
[**PatchEcommerceStoresIdProductsIdImagesId**](EcommerceApi.md#PatchEcommerceStoresIdProductsIdImagesId) | **Patch** /ecommerce/stores/{store_id}/products/{product_id}/images/{image_id} | Update product image
[**PatchEcommerceStoresIdProductsIdVariantsId**](EcommerceApi.md#PatchEcommerceStoresIdProductsIdVariantsId) | **Patch** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Update product variant
[**PatchEcommerceStoresIdPromocodesId**](EcommerceApi.md#PatchEcommerceStoresIdPromocodesId) | **Patch** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes/{promo_code_id} | Update promo code
[**PatchEcommerceStoresIdPromorulesId**](EcommerceApi.md#PatchEcommerceStoresIdPromorulesId) | **Patch** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id} | Update promo rule
[**PostEcommerceStores**](EcommerceApi.md#PostEcommerceStores) | **Post** /ecommerce/stores | Add store
[**PostEcommerceStoresIdCarts**](EcommerceApi.md#PostEcommerceStoresIdCarts) | **Post** /ecommerce/stores/{store_id}/carts | Add cart
[**PostEcommerceStoresIdCartsIdLines**](EcommerceApi.md#PostEcommerceStoresIdCartsIdLines) | **Post** /ecommerce/stores/{store_id}/carts/{cart_id}/lines | Add cart line item
[**PostEcommerceStoresIdCustomers**](EcommerceApi.md#PostEcommerceStoresIdCustomers) | **Post** /ecommerce/stores/{store_id}/customers | Add customer
[**PostEcommerceStoresIdOrders**](EcommerceApi.md#PostEcommerceStoresIdOrders) | **Post** /ecommerce/stores/{store_id}/orders | Add order
[**PostEcommerceStoresIdOrdersIdLines**](EcommerceApi.md#PostEcommerceStoresIdOrdersIdLines) | **Post** /ecommerce/stores/{store_id}/orders/{order_id}/lines | Add order line item
[**PostEcommerceStoresIdProducts**](EcommerceApi.md#PostEcommerceStoresIdProducts) | **Post** /ecommerce/stores/{store_id}/products | Add product
[**PostEcommerceStoresIdProductsIdImages**](EcommerceApi.md#PostEcommerceStoresIdProductsIdImages) | **Post** /ecommerce/stores/{store_id}/products/{product_id}/images | Add product image
[**PostEcommerceStoresIdProductsIdVariants**](EcommerceApi.md#PostEcommerceStoresIdProductsIdVariants) | **Post** /ecommerce/stores/{store_id}/products/{product_id}/variants | Add product variant
[**PostEcommerceStoresIdPromocodes**](EcommerceApi.md#PostEcommerceStoresIdPromocodes) | **Post** /ecommerce/stores/{store_id}/promo-rules/{promo_rule_id}/promo-codes | Add promo code
[**PostEcommerceStoresIdPromorules**](EcommerceApi.md#PostEcommerceStoresIdPromorules) | **Post** /ecommerce/stores/{store_id}/promo-rules | Add promo rule
[**PutEcommerceStoresIdCustomersId**](EcommerceApi.md#PutEcommerceStoresIdCustomersId) | **Put** /ecommerce/stores/{store_id}/customers/{customer_id} | Add or update customer
[**PutEcommerceStoresIdProductsIdVariantsId**](EcommerceApi.md#PutEcommerceStoresIdProductsIdVariantsId) | **Put** /ecommerce/stores/{store_id}/products/{product_id}/variants/{variant_id} | Add or update product variant


# **DeleteEcommerceStoresId**
> interface{} DeleteEcommerceStoresId(ctx, storeId)
Delete store

Delete a store. Deleting a store will also delete any associated subresources, including Customers, Orders, Products, and Carts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 

### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdCartsId**
> DeleteEcommerceStoresIdCartsId(ctx, storeId, cartId)
Delete cart

Delete a cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdCartsLinesId**
> DeleteEcommerceStoresIdCartsLinesId(ctx, storeId, cartId, lineId)
Delete cart line item

Delete a specific cart line item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
  **lineId** | **string**| The id for the line item of a cart. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdCustomersId**
> DeleteEcommerceStoresIdCustomersId(ctx, storeId, customerId)
Delete customer

Delete a customer from a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **customerId** | **string**| The id for the customer of a store. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdOrdersId**
> DeleteEcommerceStoresIdOrdersId(ctx, storeId, orderId)
Delete order

Delete an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdOrdersIdLinesId**
> DeleteEcommerceStoresIdOrdersIdLinesId(ctx, storeId, orderId, lineId)
Delete order line item

Delete a specific order line item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
  **lineId** | **string**| The id for the line item of an order. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdProductsId**
> DeleteEcommerceStoresIdProductsId(ctx, storeId, productId)
Delete product

Delete a product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdProductsIdImagesId**
> DeleteEcommerceStoresIdProductsIdImagesId(ctx, storeId, productId, imageId)
Delete product image

Delete a product image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **imageId** | **string**| The id for the product image. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdProductsIdVariantsId**
> DeleteEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId)
Delete product variant

Delete a product variant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **variantId** | **string**| The id for the product variant. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdPromocodesId**
> DeleteEcommerceStoresIdPromocodesId(ctx, storeId, promoRuleId, promoCodeId)
Delete promo code

Delete a promo code from a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
  **promoCodeId** | **string**| The id for the promo code of a store. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEcommerceStoresIdPromorulesId**
> DeleteEcommerceStoresIdPromorulesId(ctx, storeId, promoRuleId)
Delete promo rule

Delete a promo rule from a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceOrders**
> Orders GetEcommerceOrders(ctx, optional)
List account orders

Get information about an account's orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EcommerceApiGetEcommerceOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **campaignId** | **optional.String**| Restrict results to orders with a specific &#x60;campaign_id&#x60; value. | 
 **outreachId** | **optional.String**| Restrict results to orders with a specific &#x60;outreach_id&#x60; value. | 
 **customerId** | **optional.String**| Restrict results to orders made by a specific customer. | 
 **hasOutreach** | **optional.Bool**| Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad. | 

### Return type

[**Orders**](Orders.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStores**
> EcommerceStores GetEcommerceStores(ctx, optional)
List stores

Get information about all stores in the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EcommerceApiGetEcommerceStoresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**EcommerceStores**](Ecommerce Stores.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresId**
> EcommerceStore GetEcommerceStoresId(ctx, storeId, optional)
Get store info

Get information about a specific store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceStore**](Ecommerce Store.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdCarts**
> Carts GetEcommerceStoresIdCarts(ctx, storeId, optional)
List carts

Get information about a store's carts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdCartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdCartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Carts**](Carts.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdCartsId**
> EcommerceCart GetEcommerceStoresIdCartsId(ctx, storeId, cartId, optional)
Get cart info

Get information about a specific cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdCartsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdCartsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceCart**](Ecommerce Cart.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdCartsIdLines**
> CartLines GetEcommerceStoresIdCartsIdLines(ctx, storeId, cartId, optional)
List cart line items

Get information about a cart's line items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdCartsIdLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdCartsIdLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CartLines**](Cart Lines.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdCartsIdLinesId**
> EcommerceCartLineItem GetEcommerceStoresIdCartsIdLinesId(ctx, storeId, cartId, lineId, optional)
Get cart line item

Get information about a specific cart line item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
  **lineId** | **string**| The id for the line item of a cart. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdCartsIdLinesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdCartsIdLinesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceCartLineItem**](Ecommerce Cart Line Item.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdCustomers**
> Customers GetEcommerceStoresIdCustomers(ctx, storeId, optional)
List customers

Get information about a store's customers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **emailAddress** | **optional.String**| Restrict the response to customers with the email address. | 

### Return type

[**Customers**](Customers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdCustomersId**
> EcommerceCustomer GetEcommerceStoresIdCustomersId(ctx, storeId, customerId, optional)
Get customer info

Get information about a specific customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **customerId** | **string**| The id for the customer of a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdCustomersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdCustomersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceCustomer**](Ecommerce Customer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdOrders**
> Orders1 GetEcommerceStoresIdOrders(ctx, storeId, optional)
List orders

Get information about a store's orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **customerId** | **optional.String**| Restrict results to orders made by a specific customer. | 
 **hasOutreach** | **optional.Bool**| Restrict results to orders that have an outreach attached. For example, an email campaign or Facebook ad. | 
 **campaignId** | **optional.String**| Restrict results to orders with a specific &#x60;campaign_id&#x60; value. | 
 **outreachId** | **optional.String**| Restrict results to orders with a specific &#x60;outreach_id&#x60; value. | 

### Return type

[**Orders1**](Orders_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdOrdersId**
> EcommerceOrder GetEcommerceStoresIdOrdersId(ctx, storeId, orderId, optional)
Get order info

Get information about a specific order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdOrdersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdOrdersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceOrder**](Ecommerce Order.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdOrdersIdLines**
> OrderLines GetEcommerceStoresIdOrdersIdLines(ctx, storeId, orderId, optional)
List order line items

Get information about an order's line items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdOrdersIdLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdOrdersIdLinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**OrderLines**](Order Lines.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdOrdersIdLinesId**
> EcommerceOrderLineItem GetEcommerceStoresIdOrdersIdLinesId(ctx, storeId, orderId, lineId, optional)
Get order line item

Get information about a specific order line item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
  **lineId** | **string**| The id for the line item of an order. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdOrdersIdLinesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdOrdersIdLinesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceOrderLineItem**](Ecommerce Order Line Item.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdProducts**
> Products GetEcommerceStoresIdProducts(ctx, storeId, optional)
List product

Get information about a store's products.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Products**](Products.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdProductsId**
> EcommerceProduct GetEcommerceStoresIdProductsId(ctx, storeId, productId, optional)
Get product info

Get information about a specific product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdProductsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdProductsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceProduct**](Ecommerce Product.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdProductsIdImages**
> EcommerceProductImages GetEcommerceStoresIdProductsIdImages(ctx, storeId, productId, optional)
List product images

Get information about a product's images.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdProductsIdImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdProductsIdImagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**EcommerceProductImages**](Ecommerce Product Images.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdProductsIdImagesId**
> EcommerceProductImage GetEcommerceStoresIdProductsIdImagesId(ctx, storeId, productId, imageId, optional)
Get product image info

Get information about a specific product image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **imageId** | **string**| The id for the product image. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdProductsIdImagesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdProductsIdImagesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceProductImage**](Ecommerce Product Image.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdProductsIdVariants**
> EcommerceProductVariants GetEcommerceStoresIdProductsIdVariants(ctx, storeId, productId, optional)
List product variants

Get information about a product's variants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdProductsIdVariantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdProductsIdVariantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**EcommerceProductVariants**](Ecommerce Product Variants.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdProductsIdVariantsId**
> EcommerceProductVariant GetEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId, optional)
Get product variant info

Get information about a specific product variant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **variantId** | **string**| The id for the product variant. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdProductsIdVariantsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdProductsIdVariantsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommerceProductVariant**](Ecommerce Product Variant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdPromocodes**
> PromoCodes GetEcommerceStoresIdPromocodes(ctx, promoRuleId, storeId, optional)
List promo codes

Get information about a store's promo codes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdPromocodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdPromocodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**PromoCodes**](Promo Codes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdPromocodesId**
> EcommercePromoCode GetEcommerceStoresIdPromocodesId(ctx, storeId, promoRuleId, promoCodeId, optional)
Get promo code

Get information about a specific promo code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
  **promoCodeId** | **string**| The id for the promo code of a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdPromocodesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdPromocodesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommercePromoCode**](Ecommerce Promo Code.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdPromorules**
> PromoRules GetEcommerceStoresIdPromorules(ctx, storeId, optional)
List promo rules

Get information about a store's promo rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdPromorulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdPromorulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**PromoRules**](Promo Rules.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEcommerceStoresIdPromorulesId**
> EcommercePromoRule GetEcommerceStoresIdPromorulesId(ctx, storeId, promoRuleId, optional)
Get promo rule

Get information about a specific promo rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
 **optional** | ***EcommerceApiGetEcommerceStoresIdPromorulesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EcommerceApiGetEcommerceStoresIdPromorulesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EcommercePromoRule**](Ecommerce Promo Rule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresId**
> EcommerceStore PatchEcommerceStoresId(ctx, storeId, body)
Update store

Update a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **body** | [**EcommerceStore2**](EcommerceStore2.md)|  | 

### Return type

[**EcommerceStore**](Ecommerce Store.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdCartsId**
> EcommerceCart PatchEcommerceStoresIdCartsId(ctx, storeId, cartId, body)
Update cart

Update a specific cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
  **body** | [**EcommerceCart2**](EcommerceCart2.md)|  | 

### Return type

[**EcommerceCart**](Ecommerce Cart.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdCartsIdLinesId**
> EcommerceCartLineItem PatchEcommerceStoresIdCartsIdLinesId(ctx, storeId, cartId, lineId, body)
Update cart line item

Update a specific cart line item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
  **lineId** | **string**| The id for the line item of a cart. | 
  **body** | [**EcommerceCartLineItem4**](EcommerceCartLineItem4.md)|  | 

### Return type

[**EcommerceCartLineItem**](Ecommerce Cart Line Item.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdCustomersId**
> EcommerceCustomer PatchEcommerceStoresIdCustomersId(ctx, storeId, customerId, body)
Update customer

Update a customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **customerId** | **string**| The id for the customer of a store. | 
  **body** | [**EcommerceCustomer5**](EcommerceCustomer5.md)|  | 

### Return type

[**EcommerceCustomer**](Ecommerce Customer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdOrdersId**
> EcommerceOrder PatchEcommerceStoresIdOrdersId(ctx, storeId, orderId, body)
Update order

Update a specific order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
  **body** | [**EcommerceOrder2**](EcommerceOrder2.md)|  | 

### Return type

[**EcommerceOrder**](Ecommerce Order.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdOrdersIdLinesId**
> EcommerceOrderLineItem PatchEcommerceStoresIdOrdersIdLinesId(ctx, storeId, orderId, lineId, body)
Update order line item

Update a specific order line item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
  **lineId** | **string**| The id for the line item of an order. | 
  **body** | [**EcommerceOrderLineItem4**](EcommerceOrderLineItem4.md)|  | 

### Return type

[**EcommerceOrderLineItem**](Ecommerce Order Line Item.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdProductsId**
> EcommerceProduct PatchEcommerceStoresIdProductsId(ctx, storeId, productId, body)
Update product

Update a specific product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **body** | [**EcommerceProduct2**](EcommerceProduct2.md)|  | 

### Return type

[**EcommerceProduct**](Ecommerce Product.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdProductsIdImagesId**
> EcommerceProductImage PatchEcommerceStoresIdProductsIdImagesId(ctx, storeId, productId, imageId, body)
Update product image

Update a product image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **imageId** | **string**| The id for the product image. | 
  **body** | [**EcommerceProductImage4**](EcommerceProductImage4.md)|  | 

### Return type

[**EcommerceProductImage**](Ecommerce Product Image.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdProductsIdVariantsId**
> EcommerceProductVariant PatchEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId, body)
Update product variant

Update a product variant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **variantId** | **string**| The id for the product variant. | 
  **body** | [**EcommerceProductVariant5**](EcommerceProductVariant5.md)|  | 

### Return type

[**EcommerceProductVariant**](Ecommerce Product Variant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdPromocodesId**
> EcommercePromoCode PatchEcommerceStoresIdPromocodesId(ctx, storeId, promoRuleId, promoCodeId, body)
Update promo code

Update a promo code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
  **promoCodeId** | **string**| The id for the promo code of a store. | 
  **body** | [**EcommercePromoCode2**](EcommercePromoCode2.md)|  | 

### Return type

[**EcommercePromoCode**](Ecommerce Promo Code.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEcommerceStoresIdPromorulesId**
> EcommercePromoRule PatchEcommerceStoresIdPromorulesId(ctx, storeId, promoRuleId, body)
Update promo rule

Update a promo rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
  **body** | [**EcommercePromoRule2**](EcommercePromoRule2.md)|  | 

### Return type

[**EcommercePromoRule**](Ecommerce Promo Rule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStores**
> EcommerceStore PostEcommerceStores(ctx, body)
Add store

Add a new store to your Mailchimp account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EcommerceStore1**](EcommerceStore1.md)|  | 

### Return type

[**EcommerceStore**](Ecommerce Store.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdCarts**
> EcommerceCart PostEcommerceStoresIdCarts(ctx, storeId, body)
Add cart

Add a new cart to a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **body** | [**EcommerceCart1**](EcommerceCart1.md)|  | 

### Return type

[**EcommerceCart**](Ecommerce Cart.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdCartsIdLines**
> EcommerceCartLineItem PostEcommerceStoresIdCartsIdLines(ctx, storeId, cartId, body)
Add cart line item

Add a new line item to an existing cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **cartId** | **string**| The id for the cart. | 
  **body** | [**EcommerceCartLineItem3**](EcommerceCartLineItem3.md)|  | 

### Return type

[**EcommerceCartLineItem**](Ecommerce Cart Line Item.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdCustomers**
> EcommerceCustomer PostEcommerceStoresIdCustomers(ctx, storeId, body)
Add customer

Add a new customer to a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **body** | [**EcommerceCustomer3**](EcommerceCustomer3.md)|  | 

### Return type

[**EcommerceCustomer**](Ecommerce Customer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdOrders**
> EcommerceOrder PostEcommerceStoresIdOrders(ctx, storeId, body)
Add order

Add a new order to a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **body** | [**EcommerceOrder1**](EcommerceOrder1.md)|  | 

### Return type

[**EcommerceOrder**](Ecommerce Order.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdOrdersIdLines**
> EcommerceOrderLineItem PostEcommerceStoresIdOrdersIdLines(ctx, storeId, orderId, body)
Add order line item

Add a new line item to an existing order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **orderId** | **string**| The id for the order in a store. | 
  **body** | [**EcommerceOrderLineItem3**](EcommerceOrderLineItem3.md)|  | 

### Return type

[**EcommerceOrderLineItem**](Ecommerce Order Line Item.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdProducts**
> EcommerceProduct PostEcommerceStoresIdProducts(ctx, storeId, body)
Add product

Add a new product to a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **body** | [**EcommerceProduct1**](EcommerceProduct1.md)|  | 

### Return type

[**EcommerceProduct**](Ecommerce Product.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdProductsIdImages**
> EcommerceProductImage PostEcommerceStoresIdProductsIdImages(ctx, storeId, productId, body)
Add product image

Add a new image to the product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **body** | [**EcommerceProductImage3**](EcommerceProductImage3.md)|  | 

### Return type

[**EcommerceProductImage**](Ecommerce Product Image.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdProductsIdVariants**
> EcommerceProductVariant PostEcommerceStoresIdProductsIdVariants(ctx, storeId, productId, body)
Add product variant

Add a new variant to the product.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **body** | [**EcommerceProductVariant3**](EcommerceProductVariant3.md)|  | 

### Return type

[**EcommerceProductVariant**](Ecommerce Product Variant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdPromocodes**
> EcommercePromoCode PostEcommerceStoresIdPromocodes(ctx, storeId, promoRuleId, body)
Add promo code

Add a new promo code to a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **promoRuleId** | **string**| The id for the promo rule of a store. | 
  **body** | [**EcommercePromoCode1**](EcommercePromoCode1.md)|  | 

### Return type

[**EcommercePromoCode**](Ecommerce Promo Code.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEcommerceStoresIdPromorules**
> EcommercePromoRule PostEcommerceStoresIdPromorules(ctx, storeId, body)
Add promo rule

Add a new promo rule to a store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **body** | [**EcommercePromoRule1**](EcommercePromoRule1.md)|  | 

### Return type

[**EcommercePromoRule**](Ecommerce Promo Rule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutEcommerceStoresIdCustomersId**
> EcommerceCustomer PutEcommerceStoresIdCustomersId(ctx, storeId, customerId, body)
Add or update customer

Add or update a customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **customerId** | **string**| The id for the customer of a store. | 
  **body** | [**EcommerceCustomer4**](EcommerceCustomer4.md)|  | 

### Return type

[**EcommerceCustomer**](Ecommerce Customer.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutEcommerceStoresIdProductsIdVariantsId**
> EcommerceProductVariant PutEcommerceStoresIdProductsIdVariantsId(ctx, storeId, productId, variantId, body)
Add or update product variant

Add or update a product variant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storeId** | **string**| The store id. | 
  **productId** | **string**| The id for the product of a store. | 
  **variantId** | **string**| The id for the product variant. | 
  **body** | [**EcommerceProductVariant4**](EcommerceProductVariant4.md)|  | 

### Return type

[**EcommerceProductVariant**](Ecommerce Product Variant.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

