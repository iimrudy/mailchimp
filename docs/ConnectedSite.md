# ConnectedSite

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignId** | **string** | The unique identifier for the site. | [optional] [default to null]
**StoreId** | **string** | The unique identifier for the ecommerce store that&#39;s associated with the connected site (if any). The store_id for a specific connected site can&#39;t change. | [optional] [default to null]
**Platform** | **string** | The platform of the connected site. | [optional] [default to null]
**Domain** | **string** | The connected site domain. | [optional] [default to null]
**SiteScript** | [***Script**](Script.md) |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the connected site was created in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the connected site was last updated in ISO 8601 format. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


