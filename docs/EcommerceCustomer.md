# EcommerceCustomer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the customer. | [optional] [default to null]
**EmailAddress** | **string** | The customer&#39;s email address. | [optional] [default to null]
**OptInStatus** | **bool** | The customer&#39;s opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don&#39;t opt in to your Mailchimp list [will be added as &#x60;Transactional&#x60; members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers). | [optional] [default to null]
**Company** | **string** | The customer&#39;s company. | [optional] [default to null]
**FirstName** | **string** | The customer&#39;s first name. | [optional] [default to null]
**LastName** | **string** | The customer&#39;s last name. | [optional] [default to null]
**OrdersCount** | **int32** | The customer&#39;s total order count. | [optional] [default to null]
**TotalSpent** | **float32** | The total amount the customer has spent. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the customer was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the customer was last updated in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


