# Batch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that uniquely identifies this batch request. | [optional] [default to null]
**Status** | **string** | The status of the batch call. [Learn more](https://mailchimp.com/developer/marketing/guides/run-async-requests-batch-endpoint/#check-the-status-of-a-batch-operation) about the batch operation status. | [optional] [default to null]
**TotalOperations** | **int32** | The total number of operations to complete as part of this batch request. For GET requests requiring pagination, each page counts as a separate operation. | [optional] [default to null]
**FinishedOperations** | **int32** | The number of completed operations. This includes operations that returned an error. | [optional] [default to null]
**ErroredOperations** | **int32** | The number of completed operations that returned an error. | [optional] [default to null]
**SubmittedAt** | [**time.Time**](time.Time.md) | The date and time when the server received the batch request in ISO 8601 format. | [optional] [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | The date and time when all operations in the batch request completed in ISO 8601 format. | [optional] [default to null]
**ResponseBodyUrl** | **string** | The URL of the gzipped archive of the results of all the operations. | [optional] [default to null]
**Links** | [**[]ResourceLink**](Resource Link.md) | A list of link types and descriptions for the API schema documents. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


