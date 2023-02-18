# LandingPageReport

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this landing page. | [optional] [default to null]
**Name** | **string** | The name of this landing page the user will see. | [optional] [default to null]
**Title** | **string** | The name of the landing page the user&#39;s customers will see. | [optional] [default to null]
**Url** | **string** | The landing page url. | [optional] [default to null]
**PublishedAt** | [**time.Time**](time.Time.md) | The time this landing page was published. | [optional] [default to null]
**UnpublishedAt** | [**time.Time**](time.Time.md) | The time this landing page was unpublished. | [optional] [default to null]
**Status** | **string** | The status of the landing page. | [optional] [default to null]
**ListId** | **string** | The list id connected to this landing page. | [optional] [default to null]
**Visits** | **int32** | The number of visits to this landing pages. | [optional] [default to null]
**UniqueVisits** | **int32** | The number of unique visits to this landing pages. | [optional] [default to null]
**Subscribes** | **int32** | The number of subscribes to this landing pages. | [optional] [default to null]
**Clicks** | **int32** | The number of clicks to this landing pages. | [optional] [default to null]
**ConversionRate** | **float32** | The percentage of people who visited your landing page and were added to your list. | [optional] [default to null]
**Timeseries** | [***LandingPageReportTimeseries**](Landing Page Report_timeseries.md) |  | [optional] [default to null]
**Ecommerce** | [***LandingPageReportEcommerce**](Landing Page Report_ecommerce.md) |  | [optional] [default to null]
**WebId** | **int32** | The ID used in the Mailchimp web application. | [optional] [default to null]
**ListName** | **string** | List Name | [optional] [default to null]
**SignupTags** | [**[]Tag**](Tag.md) | A list of tags associated to the landing page. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


