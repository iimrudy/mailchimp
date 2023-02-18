# \ListsApi

All URIs are relative to *https://server.api.mailchimp.com/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListsId**](ListsApi.md#DeleteListsId) | **Delete** /lists/{list_id} | Delete list
[**DeleteListsIdInterestCategoriesId**](ListsApi.md#DeleteListsIdInterestCategoriesId) | **Delete** /lists/{list_id}/interest-categories/{interest_category_id} | Delete interest category
[**DeleteListsIdInterestCategoriesIdInterestsId**](ListsApi.md#DeleteListsIdInterestCategoriesIdInterestsId) | **Delete** /lists/{list_id}/interest-categories/{interest_category_id}/interests/{interest_id} | Delete interest in category
[**DeleteListsIdMembersId**](ListsApi.md#DeleteListsIdMembersId) | **Delete** /lists/{list_id}/members/{subscriber_hash} | Archive list member
[**DeleteListsIdMembersIdNotesId**](ListsApi.md#DeleteListsIdMembersIdNotesId) | **Delete** /lists/{list_id}/members/{subscriber_hash}/notes/{note_id} | Delete note
[**DeleteListsIdMergeFieldsId**](ListsApi.md#DeleteListsIdMergeFieldsId) | **Delete** /lists/{list_id}/merge-fields/{merge_id} | Delete merge field
[**DeleteListsIdSegmentsId**](ListsApi.md#DeleteListsIdSegmentsId) | **Delete** /lists/{list_id}/segments/{segment_id} | Delete segment
[**DeleteListsIdSegmentsIdMembersId**](ListsApi.md#DeleteListsIdSegmentsIdMembersId) | **Delete** /lists/{list_id}/segments/{segment_id}/members/{subscriber_hash} | Remove list member from segment
[**DeleteListsIdWebhooksId**](ListsApi.md#DeleteListsIdWebhooksId) | **Delete** /lists/{list_id}/webhooks/{webhook_id} | Delete webhook
[**GetListMemberTags**](ListsApi.md#GetListMemberTags) | **Get** /lists/{list_id}/members/{subscriber_hash}/tags | List member tags
[**GetLists**](ListsApi.md#GetLists) | **Get** /lists | Get lists info
[**GetListsId**](ListsApi.md#GetListsId) | **Get** /lists/{list_id} | Get list info
[**GetListsIdAbuseReports**](ListsApi.md#GetListsIdAbuseReports) | **Get** /lists/{list_id}/abuse-reports | List abuse reports
[**GetListsIdAbuseReportsId**](ListsApi.md#GetListsIdAbuseReportsId) | **Get** /lists/{list_id}/abuse-reports/{report_id} | Get abuse report
[**GetListsIdActivity**](ListsApi.md#GetListsIdActivity) | **Get** /lists/{list_id}/activity | List recent activity
[**GetListsIdClients**](ListsApi.md#GetListsIdClients) | **Get** /lists/{list_id}/clients | List top email clients
[**GetListsIdGrowthHistory**](ListsApi.md#GetListsIdGrowthHistory) | **Get** /lists/{list_id}/growth-history | List growth history data
[**GetListsIdGrowthHistoryId**](ListsApi.md#GetListsIdGrowthHistoryId) | **Get** /lists/{list_id}/growth-history/{month} | Get growth history by month
[**GetListsIdInterestCategories**](ListsApi.md#GetListsIdInterestCategories) | **Get** /lists/{list_id}/interest-categories | List interest categories
[**GetListsIdInterestCategoriesId**](ListsApi.md#GetListsIdInterestCategoriesId) | **Get** /lists/{list_id}/interest-categories/{interest_category_id} | Get interest category info
[**GetListsIdInterestCategoriesIdInterests**](ListsApi.md#GetListsIdInterestCategoriesIdInterests) | **Get** /lists/{list_id}/interest-categories/{interest_category_id}/interests | List interests in category
[**GetListsIdInterestCategoriesIdInterestsId**](ListsApi.md#GetListsIdInterestCategoriesIdInterestsId) | **Get** /lists/{list_id}/interest-categories/{interest_category_id}/interests/{interest_id} | Get interest in category
[**GetListsIdLocations**](ListsApi.md#GetListsIdLocations) | **Get** /lists/{list_id}/locations | List locations
[**GetListsIdMembers**](ListsApi.md#GetListsIdMembers) | **Get** /lists/{list_id}/members | List members info
[**GetListsIdMembersId**](ListsApi.md#GetListsIdMembersId) | **Get** /lists/{list_id}/members/{subscriber_hash} | Get member info
[**GetListsIdMembersIdActivity**](ListsApi.md#GetListsIdMembersIdActivity) | **Get** /lists/{list_id}/members/{subscriber_hash}/activity | View recent activity 50
[**GetListsIdMembersIdActivityFeed**](ListsApi.md#GetListsIdMembersIdActivityFeed) | **Get** /lists/{list_id}/members/{subscriber_hash}/activity-feed | View recent activity
[**GetListsIdMembersIdEvents**](ListsApi.md#GetListsIdMembersIdEvents) | **Get** /lists/{list_id}/members/{subscriber_hash}/events | List member events
[**GetListsIdMembersIdGoals**](ListsApi.md#GetListsIdMembersIdGoals) | **Get** /lists/{list_id}/members/{subscriber_hash}/goals | List member goal events
[**GetListsIdMembersIdNotes**](ListsApi.md#GetListsIdMembersIdNotes) | **Get** /lists/{list_id}/members/{subscriber_hash}/notes | List recent member notes
[**GetListsIdMembersIdNotesId**](ListsApi.md#GetListsIdMembersIdNotesId) | **Get** /lists/{list_id}/members/{subscriber_hash}/notes/{note_id} | Get member note
[**GetListsIdMergeFields**](ListsApi.md#GetListsIdMergeFields) | **Get** /lists/{list_id}/merge-fields | List merge fields
[**GetListsIdMergeFieldsId**](ListsApi.md#GetListsIdMergeFieldsId) | **Get** /lists/{list_id}/merge-fields/{merge_id} | Get merge field
[**GetListsIdSegmentsId**](ListsApi.md#GetListsIdSegmentsId) | **Get** /lists/{list_id}/segments/{segment_id} | Get segment info
[**GetListsIdSegmentsIdMembers**](ListsApi.md#GetListsIdSegmentsIdMembers) | **Get** /lists/{list_id}/segments/{segment_id}/members | List members in segment
[**GetListsIdSignupForms**](ListsApi.md#GetListsIdSignupForms) | **Get** /lists/{list_id}/signup-forms | List signup forms
[**GetListsIdSurveys**](ListsApi.md#GetListsIdSurveys) | **Get** /lists/{list_id}/surveys | Get information about all surveys for a list
[**GetListsIdSurveysId**](ListsApi.md#GetListsIdSurveysId) | **Get** /lists/{list_id}/surveys/{survey_id} | Get survey
[**GetListsIdWebhooks**](ListsApi.md#GetListsIdWebhooks) | **Get** /lists/{list_id}/webhooks | List webhooks
[**GetListsIdWebhooksId**](ListsApi.md#GetListsIdWebhooksId) | **Get** /lists/{list_id}/webhooks/{webhook_id} | Get webhook info
[**PatchListsId**](ListsApi.md#PatchListsId) | **Patch** /lists/{list_id} | Update lists
[**PatchListsIdInterestCategoriesId**](ListsApi.md#PatchListsIdInterestCategoriesId) | **Patch** /lists/{list_id}/interest-categories/{interest_category_id} | Update interest category
[**PatchListsIdInterestCategoriesIdInterestsId**](ListsApi.md#PatchListsIdInterestCategoriesIdInterestsId) | **Patch** /lists/{list_id}/interest-categories/{interest_category_id}/interests/{interest_id} | Update interest in category
[**PatchListsIdMembersId**](ListsApi.md#PatchListsIdMembersId) | **Patch** /lists/{list_id}/members/{subscriber_hash} | Update list member
[**PatchListsIdMembersIdNotesId**](ListsApi.md#PatchListsIdMembersIdNotesId) | **Patch** /lists/{list_id}/members/{subscriber_hash}/notes/{note_id} | Update note
[**PatchListsIdMergeFieldsId**](ListsApi.md#PatchListsIdMergeFieldsId) | **Patch** /lists/{list_id}/merge-fields/{merge_id} | Update merge field
[**PatchListsIdSegmentsId**](ListsApi.md#PatchListsIdSegmentsId) | **Patch** /lists/{list_id}/segments/{segment_id} | Update segment
[**PatchListsIdWebhooksId**](ListsApi.md#PatchListsIdWebhooksId) | **Patch** /lists/{list_id}/webhooks/{webhook_id} | Update webhook
[**PostListMemberEvents**](ListsApi.md#PostListMemberEvents) | **Post** /lists/{list_id}/members/{subscriber_hash}/events | Add event
[**PostListMemberTags**](ListsApi.md#PostListMemberTags) | **Post** /lists/{list_id}/members/{subscriber_hash}/tags | Add or remove member tags
[**PostLists**](ListsApi.md#PostLists) | **Post** /lists | Add list
[**PostListsId**](ListsApi.md#PostListsId) | **Post** /lists/{list_id} | Batch subscribe or unsubscribe
[**PostListsIdInterestCategories**](ListsApi.md#PostListsIdInterestCategories) | **Post** /lists/{list_id}/interest-categories | Add interest category
[**PostListsIdInterestCategoriesIdInterests**](ListsApi.md#PostListsIdInterestCategoriesIdInterests) | **Post** /lists/{list_id}/interest-categories/{interest_category_id}/interests | Add interest in category
[**PostListsIdMembers**](ListsApi.md#PostListsIdMembers) | **Post** /lists/{list_id}/members | Add member to list
[**PostListsIdMembersHashActionsDeletePermanent**](ListsApi.md#PostListsIdMembersHashActionsDeletePermanent) | **Post** /lists/{list_id}/members/{subscriber_hash}/actions/delete-permanent | Delete list member
[**PostListsIdMembersIdNotes**](ListsApi.md#PostListsIdMembersIdNotes) | **Post** /lists/{list_id}/members/{subscriber_hash}/notes | Add member note
[**PostListsIdMergeFields**](ListsApi.md#PostListsIdMergeFields) | **Post** /lists/{list_id}/merge-fields | Add merge field
[**PostListsIdSegments**](ListsApi.md#PostListsIdSegments) | **Post** /lists/{list_id}/segments | Add segment
[**PostListsIdSegmentsId**](ListsApi.md#PostListsIdSegmentsId) | **Post** /lists/{list_id}/segments/{segment_id} | Batch add or remove members
[**PostListsIdSegmentsIdMembers**](ListsApi.md#PostListsIdSegmentsIdMembers) | **Post** /lists/{list_id}/segments/{segment_id}/members | Add member to segment
[**PostListsIdSignupForms**](ListsApi.md#PostListsIdSignupForms) | **Post** /lists/{list_id}/signup-forms | Customize signup form
[**PostListsIdWebhooks**](ListsApi.md#PostListsIdWebhooks) | **Post** /lists/{list_id}/webhooks | Add webhook
[**PreviewASegment**](ListsApi.md#PreviewASegment) | **Get** /lists/{list_id}/segments | List segments
[**PutListsIdMembersId**](ListsApi.md#PutListsIdMembersId) | **Put** /lists/{list_id}/members/{subscriber_hash} | Add or update list member
[**SearchTagsByName**](ListsApi.md#SearchTagsByName) | **Get** /lists/{list_id}/tag-search | Search for tags on a list by name.


# **DeleteListsId**
> DeleteListsId(ctx, listId)
Delete list

Delete a list from your Mailchimp account. If you delete a list, you'll lose the list history—including subscriber activity, unsubscribes, complaints, and bounces. You’ll also lose subscribers’ email addresses, unless you exported and backed up your list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdInterestCategoriesId**
> DeleteListsIdInterestCategoriesId(ctx, listId, interestCategoryId)
Delete interest category

Delete a specific interest category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdInterestCategoriesIdInterestsId**
> DeleteListsIdInterestCategoriesIdInterestsId(ctx, listId, interestCategoryId, interestId)
Delete interest in category

Delete interests or group names in a specific category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
  **interestId** | **string**| The specific interest or &#39;group name&#39;. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdMembersId**
> DeleteListsIdMembersId(ctx, listId, subscriberHash)
Archive list member

Archive a list member. To permanently delete, use the delete-permanent action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdMembersIdNotesId**
> DeleteListsIdMembersIdNotesId(ctx, listId, subscriberHash, noteId)
Delete note

Delete a specific note for a specific list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
  **noteId** | **string**| The id for the note. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdMergeFieldsId**
> DeleteListsIdMergeFieldsId(ctx, listId, mergeId)
Delete merge field

Delete a specific merge field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **mergeId** | **string**| The id for the merge field. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdSegmentsId**
> DeleteListsIdSegmentsId(ctx, listId, segmentId)
Delete segment

Delete a specific segment in a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdSegmentsIdMembersId**
> DeleteListsIdSegmentsIdMembersId(ctx, listId, segmentId, subscriberHash)
Remove list member from segment

Remove a member from the specified static segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsIdWebhooksId**
> DeleteListsIdWebhooksId(ctx, listId, webhookId)
Delete webhook

Delete a specific webhook in a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **webhookId** | **string**| The webhook&#39;s id. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListMemberTags**
> CollectionOfTags GetListMemberTags(ctx, listId, subscriberHash, optional)
List member tags

Get the tags on a list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
 **optional** | ***ListsApiGetListMemberTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListMemberTagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CollectionOfTags**](Collection of Tags.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLists**
> SubscriberLists GetLists(ctx, optional)
Get lists info

Get information about all lists in the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListsApiGetListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **beforeDateCreated** | **optional.String**| Restrict response to lists created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceDateCreated** | **optional.String**| Restrict results to lists created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeCampaignLastSent** | **optional.String**| Restrict results to lists created before the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceCampaignLastSent** | **optional.String**| Restrict results to lists created after the last campaign send date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **email** | **optional.String**| Restrict results to lists that include a specific subscriber&#39;s email address. | 
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 
 **hasEcommerceStore** | **optional.Bool**| Restrict results to lists that contain an active, connected, undeleted ecommerce store. | 
 **includeTotalContacts** | **optional.Bool**| Return the total_contacts field in the stats response, which contains an approximate count of all contacts in any state. | 

### Return type

[**SubscriberLists**](Subscriber Lists.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsId**
> SubscriberList GetListsId(ctx, listId, optional)
Get list info

Get information about a specific list in your Mailchimp account. Results include list members who have signed up but haven't confirmed their subscription yet and unsubscribed or cleaned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **includeTotalContacts** | **optional.Bool**| Return the total_contacts field in the stats response, which contains an approximate count of all contacts in any state. | 

### Return type

[**SubscriberList**](Subscriber List.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdAbuseReports**
> AbuseComplaints GetListsIdAbuseReports(ctx, listId, optional)
List abuse reports

Get all abuse reports for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdAbuseReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdAbuseReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**AbuseComplaints**](Abuse Complaints.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdAbuseReportsId**
> AbuseComplaint GetListsIdAbuseReportsId(ctx, listId, reportId, optional)
Get abuse report

Get details about a specific abuse report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **reportId** | **string**| The id for the abuse report. | 
 **optional** | ***ListsApiGetListsIdAbuseReportsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdAbuseReportsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**AbuseComplaint**](Abuse Complaint.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdActivity**
> ListActivity GetListsIdActivity(ctx, listId, optional)
List recent activity

Get up to the previous 180 days of daily detailed aggregated activity stats for a list, not including Automation activity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdActivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ListActivity**](List Activity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdClients**
> EmailClients GetListsIdClients(ctx, listId, optional)
List top email clients

Get a list of the top email clients based on user-agent strings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdClientsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**EmailClients**](Email Clients.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdGrowthHistory**
> GrowthHistory GetListsIdGrowthHistory(ctx, listId, optional)
List growth history data

Get a month-by-month summary of a specific list's growth activity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdGrowthHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdGrowthHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 

### Return type

[**GrowthHistory**](Growth History.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdGrowthHistoryId**
> GrowthHistory GetListsIdGrowthHistoryId(ctx, listId, month, optional)
Get growth history by month

Get a summary of a specific list's growth activity for a specific month and year.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **month** | **string**| A specific month of list growth history. | 
 **optional** | ***ListsApiGetListsIdGrowthHistoryIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdGrowthHistoryIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**GrowthHistory**](Growth History.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdInterestCategories**
> InterestGroupings GetListsIdInterestCategories(ctx, listId, optional)
List interest categories

Get information about a list's interest categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdInterestCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdInterestCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| Restrict results a type of interest group | 

### Return type

[**InterestGroupings**](Interest Groupings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdInterestCategoriesId**
> InterestCategory GetListsIdInterestCategoriesId(ctx, listId, interestCategoryId, optional)
Get interest category info

Get information about a specific interest category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
 **optional** | ***ListsApiGetListsIdInterestCategoriesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdInterestCategoriesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**InterestCategory**](Interest Category.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdInterestCategoriesIdInterests**
> Interests GetListsIdInterestCategoriesIdInterests(ctx, listId, interestCategoryId, optional)
List interests in category

Get a list of this category's interests.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
 **optional** | ***ListsApiGetListsIdInterestCategoriesIdInterestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdInterestCategoriesIdInterestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**Interests**](Interests.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdInterestCategoriesIdInterestsId**
> Interest GetListsIdInterestCategoriesIdInterestsId(ctx, listId, interestCategoryId, interestId, optional)
Get interest in category

Get interests or 'group names' for a specific category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
  **interestId** | **string**| The specific interest or &#39;group name&#39;. | 
 **optional** | ***ListsApiGetListsIdInterestCategoriesIdInterestsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdInterestCategoriesIdInterestsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**Interest**](Interest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdLocations**
> ListLocations GetListsIdLocations(ctx, listId, optional)
List locations

Get the locations (countries) that the list's subscribers have been tagged to based on geocoding their IP address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ListLocations**](List Locations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembers**
> ListMembers2 GetListsIdMembers(ctx, listId, optional)
List members info

Get information about members in a specific Mailchimp list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **emailType** | **optional.String**| The email type. | 
 **status** | **optional.String**| The subscriber&#39;s status. | 
 **sinceTimestampOpt** | **optional.String**| Restrict results to subscribers who opted-in after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeTimestampOpt** | **optional.String**| Restrict results to subscribers who opted-in before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **sinceLastChanged** | **optional.String**| Restrict results to subscribers whose information changed after the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeLastChanged** | **optional.String**| Restrict results to subscribers whose information changed before the set timeframe. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **uniqueEmailId** | **optional.String**| A unique identifier for the email address across all Mailchimp lists. | 
 **vipOnly** | **optional.Bool**| A filter to return only the list&#39;s VIP members. Passing &#x60;true&#x60; will restrict results to VIP list members, passing &#x60;false&#x60; will return all list members. | 
 **interestCategoryId** | **optional.String**| The unique id for the interest category. | 
 **interestIds** | **optional.String**| Used to filter list members by interests. Must be accompanied by interest_category_id and interest_match. The value must be a comma separated list of interest ids present for any supplied interest categories. | 
 **interestMatch** | **optional.String**| Used to filter list members by interests. Must be accompanied by interest_category_id and interest_ids. \&quot;any\&quot; will match a member with any of the interest supplied, \&quot;all\&quot; will only match members with every interest supplied, and \&quot;none\&quot; will match members without any of the interest supplied. | 
 **sortField** | **optional.String**| Returns files sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 
 **sinceLastCampaign** | **optional.Bool**| Filter subscribers by those subscribed/unsubscribed/pending/cleaned since last email campaign send. Member status is required to use this filter. | 
 **unsubscribedSince** | **optional.String**| Filter subscribers by those unsubscribed since a specific date. Using any status other than unsubscribed with this filter will result in an error. | 

### Return type

[**ListMembers2**](List Members_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersId**
> ListMembers2 GetListsIdMembersId(ctx, listId, subscriberHash, optional)
Get member info

Get information about a specific list member, including a currently subscribed, unsubscribed, or bounced member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
 **optional** | ***ListsApiGetListsIdMembersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**ListMembers2**](List Members_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersIdActivity**
> MemberActivityEvents GetListsIdMembersIdActivity(ctx, listId, subscriberHash, optional)
View recent activity 50

Get the last 50 events of a member's activity on a specific list, including opens, clicks, and unsubscribes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
 **optional** | ***ListsApiGetListsIdMembersIdActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdActivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **action** | [**optional.Interface of []string**](string.md)| A comma seperated list of actions to return. | 

### Return type

[**MemberActivityEvents**](Member Activity Events.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersIdActivityFeed**
> MemberActivityEvents1 GetListsIdMembersIdActivityFeed(ctx, listId, subscriberHash, optional)
View recent activity

Get a member's activity on a specific list, including opens, clicks, and unsubscribes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ListsApiGetListsIdMembersIdActivityFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdActivityFeedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **activityFilters** | [**optional.Interface of []string**](string.md)| A comma-separated list of activity filters that correspond to a set of activity types, e.g \&quot;?activity_filters&#x3D;open,bounce,click\&quot;. | 

### Return type

[**MemberActivityEvents1**](Member Activity Events_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersIdEvents**
> CollectionOfEvents GetListsIdMembersIdEvents(ctx, listId, subscriberHash, optional)
List member events

Get events for a contact.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
 **optional** | ***ListsApiGetListsIdMembersIdEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CollectionOfEvents**](Collection of Events.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersIdGoals**
> CollectionOfMemberActivityEvents GetListsIdMembersIdGoals(ctx, listId, subscriberHash, optional)
List member goal events

Get the last 50 Goal events for a member on a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
 **optional** | ***ListsApiGetListsIdMembersIdGoalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdGoalsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**CollectionOfMemberActivityEvents**](Collection of Member Activity Events.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersIdNotes**
> CollectionOfNotes GetListsIdMembersIdNotes(ctx, listId, subscriberHash, optional)
List recent member notes

Get recent notes for a specific list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
 **optional** | ***ListsApiGetListsIdMembersIdNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortField** | **optional.String**| Returns notes sorted by the specified field. | 
 **sortDir** | **optional.String**| Determines the order direction for sorted results. | 
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]

### Return type

[**CollectionOfNotes**](Collection of Notes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMembersIdNotesId**
> MemberNotes GetListsIdMembersIdNotesId(ctx, listId, subscriberHash, noteId, optional)
Get member note

Get a specific note for a specific list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
  **noteId** | **string**| The id for the note. | 
 **optional** | ***ListsApiGetListsIdMembersIdNotesIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMembersIdNotesIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 

### Return type

[**MemberNotes**](Member Notes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMergeFields**
> CollectionOfMergeFields GetListsIdMergeFields(ctx, listId, optional)
List merge fields

Get a list of all merge fields for an audience.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiGetListsIdMergeFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMergeFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| The merge field type. | 
 **required** | **optional.Bool**| Whether it&#39;s a required merge field. | 

### Return type

[**CollectionOfMergeFields**](Collection of Merge Fields.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdMergeFieldsId**
> MergeField GetListsIdMergeFieldsId(ctx, listId, mergeId, optional)
Get merge field

Get information about a specific merge field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **mergeId** | **string**| The id for the merge field. | 
 **optional** | ***ListsApiGetListsIdMergeFieldsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdMergeFieldsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 

### Return type

[**MergeField**](Merge Field.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdSegmentsId**
> List7 GetListsIdSegmentsId(ctx, listId, segmentId, optional)
Get segment info

Get information about a specific segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 
 **optional** | ***ListsApiGetListsIdSegmentsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdSegmentsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **includeCleaned** | **optional.Bool**| Include cleaned members in response | 
 **includeTransactional** | **optional.Bool**| Include transactional members in response | 
 **includeUnsubscribed** | **optional.Bool**| Include unsubscribed members in response | 

### Return type

[**List7**](List_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdSegmentsIdMembers**
> SegmentMembers GetListsIdSegmentsIdMembers(ctx, listId, segmentId, optional)
List members in segment

Get information about members in a saved segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 
 **optional** | ***ListsApiGetListsIdSegmentsIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiGetListsIdSegmentsIdMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **includeCleaned** | **optional.Bool**| Include cleaned members in response | 
 **includeTransactional** | **optional.Bool**| Include transactional members in response | 
 **includeUnsubscribed** | **optional.Bool**| Include unsubscribed members in response | 

### Return type

[**SegmentMembers**](Segment Members.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdSignupForms**
> ListSignupForms GetListsIdSignupForms(ctx, listId)
List signup forms

Get signup forms for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 

### Return type

[**ListSignupForms**](List Signup Forms.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdSurveys**
> GetListsIdSurveys(ctx, listId)
Get information about all surveys for a list

Get information about all available surveys for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdSurveysId**
> GetListsIdSurveysId(ctx, listId, surveyId)
Get survey

Get details about a specific survey.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **surveyId** | **string**| The ID of the survey. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdWebhooks**
> ListWebhooks GetListsIdWebhooks(ctx, listId)
List webhooks

Get information about all webhooks for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 

### Return type

[**ListWebhooks**](List Webhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsIdWebhooksId**
> ListWebhooks GetListsIdWebhooksId(ctx, listId, webhookId)
Get webhook info

Get information about a specific webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **webhookId** | **string**| The webhook&#39;s id. | 

### Return type

[**ListWebhooks**](List Webhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsId**
> SubscriberList PatchListsId(ctx, listId, body)
Update lists

Update the settings for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**SubscriberList2**](SubscriberList2.md)|  | 

### Return type

[**SubscriberList**](Subscriber List.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdInterestCategoriesId**
> InterestCategory PatchListsIdInterestCategoriesId(ctx, listId, interestCategoryId, body)
Update interest category

Update a specific interest category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
  **body** | [**InterestCategory2**](InterestCategory2.md)|  | 

### Return type

[**InterestCategory**](Interest Category.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdInterestCategoriesIdInterestsId**
> Interest PatchListsIdInterestCategoriesIdInterestsId(ctx, listId, interestCategoryId, interestId, body)
Update interest in category

Update interests or 'group names' for a specific category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
  **interestId** | **string**| The specific interest or &#39;group name&#39;. | 
  **body** | [**Interest2**](Interest2.md)|  | 

### Return type

[**Interest**](Interest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdMembersId**
> ListMembers2 PatchListsIdMembersId(ctx, listId, subscriberHash, body, optional)
Update list member

Update information for a specific list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
  **body** | [**AddListMembers3**](AddListMembers3.md)|  | 
 **optional** | ***ListsApiPatchListsIdMembersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiPatchListsIdMembersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **skipMergeValidation** | **optional.Bool**| If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 

### Return type

[**ListMembers2**](List Members_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdMembersIdNotesId**
> MemberNotes PatchListsIdMembersIdNotesId(ctx, listId, subscriberHash, noteId, body)
Update note

Update a specific note for a specific list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
  **noteId** | **string**| The id for the note. | 
  **body** | [**MemberNotes2**](MemberNotes2.md)|  | 

### Return type

[**MemberNotes**](Member Notes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdMergeFieldsId**
> MergeField PatchListsIdMergeFieldsId(ctx, listId, mergeId, body)
Update merge field

Update a specific merge field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **mergeId** | **string**| The id for the merge field. | 
  **body** | [**MergeField2**](MergeField2.md)|  | 

### Return type

[**MergeField**](Merge Field.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdSegmentsId**
> List7 PatchListsIdSegmentsId(ctx, listId, segmentId, body)
Update segment

Update a specific segment in a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 
  **body** | [**List9**](List9.md)|  | 

### Return type

[**List7**](List_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchListsIdWebhooksId**
> ListWebhooks PatchListsIdWebhooksId(ctx, listId, webhookId, body)
Update webhook

Update the settings for an existing webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **webhookId** | **string**| The webhook&#39;s id. | 
  **body** | [**AddWebhook1**](AddWebhook1.md)|  | 

### Return type

[**ListWebhooks**](List Webhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListMemberEvents**
> PostListMemberEvents(ctx, listId, subscriberHash, body)
Add event

Add an event for a list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
  **body** | [**Events**](Events.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListMemberTags**
> PostListMemberTags(ctx, listId, subscriberHash, body)
Add or remove member tags

Add or remove tags from a list member. If a tag that does not exist is passed in and set as 'active', a new tag will be created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
  **body** | [**MemberTags**](MemberTags.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLists**
> SubscriberList PostLists(ctx, body)
Add list

Create a new list in your Mailchimp account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubscriberList1**](SubscriberList1.md)|  | 

### Return type

[**SubscriberList**](Subscriber List.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsId**
> BatchUpdateListMembers PostListsId(ctx, listId, body, optional)
Batch subscribe or unsubscribe

Batch subscribe or unsubscribe list members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**MembersToSubscribeUnsubscribeTofromAListInBatch**](MembersToSubscribeUnsubscribeTofromAListInBatch.md)|  | 
 **optional** | ***ListsApiPostListsIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiPostListsIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipMergeValidation** | **optional.Bool**| If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 
 **skipDuplicateCheck** | **optional.Bool**| If skip_duplicate_check is true, we will ignore duplicates sent in the request when using the batch sub/unsub on the lists endpoint. The status of the first appearance in the request will be saved. This defaults to false. | 

### Return type

[**BatchUpdateListMembers**](Batch update List members.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdInterestCategories**
> InterestCategory PostListsIdInterestCategories(ctx, listId, body)
Add interest category

Create a new interest category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**InterestCategory1**](InterestCategory1.md)|  | 

### Return type

[**InterestCategory**](Interest Category.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdInterestCategoriesIdInterests**
> Interest PostListsIdInterestCategoriesIdInterests(ctx, listId, interestCategoryId, body)
Add interest in category

Create a new interest or 'group name' for a specific category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **interestCategoryId** | **string**| The unique ID for the interest category. | 
  **body** | [**Interest1**](Interest1.md)|  | 

### Return type

[**Interest**](Interest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdMembers**
> ListMembers2 PostListsIdMembers(ctx, listId, body, optional)
Add member to list

Add a new member to the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**AddListMembers1**](AddListMembers1.md)|  | 
 **optional** | ***ListsApiPostListsIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiPostListsIdMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipMergeValidation** | **optional.Bool**| If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 

### Return type

[**ListMembers2**](List Members_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdMembersHashActionsDeletePermanent**
> PostListsIdMembersHashActionsDeletePermanent(ctx, listId, subscriberHash)
Delete list member

Delete all personally identifiable information related to a list member, and remove them from a list. This will make it impossible to re-import the list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdMembersIdNotes**
> MemberNotes PostListsIdMembersIdNotes(ctx, listId, subscriberHash, body)
Add member note

Add a new note for a specific subscriber.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. | 
  **body** | [**MemberNotes1**](MemberNotes1.md)|  | 

### Return type

[**MemberNotes**](Member Notes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdMergeFields**
> MergeField PostListsIdMergeFields(ctx, listId, body)
Add merge field

Add a new merge field for a specific audience.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**MergeField1**](MergeField1.md)|  | 

### Return type

[**MergeField**](Merge Field.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSegments**
> List7 PostListsIdSegments(ctx, listId, body)
Add segment

Create a new segment in a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**List8**](List8.md)|  | 

### Return type

[**List7**](List_7.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSegmentsId**
> BatchAddremoveListMembersTofromStaticSegment PostListsIdSegmentsId(ctx, body, listId, segmentId)
Batch add or remove members

Batch add/remove list members to static segment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MembersToAddremoveTofromAStaticSegment**](MembersToAddremoveTofromAStaticSegment.md)|  | 
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 

### Return type

[**BatchAddremoveListMembersTofromStaticSegment**](Batch addremove List members tofrom static segment.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSegmentsIdMembers**
> ListMembers1 PostListsIdSegmentsIdMembers(ctx, listId, segmentId, body)
Add member to segment

Add a member to a static segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **segmentId** | **string**| The unique id for the segment. | 
  **body** | [**Body3**](Body3.md)|  | 

### Return type

[**ListMembers1**](List Members_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdSignupForms**
> SignupForm PostListsIdSignupForms(ctx, listId, body)
Customize signup form

Customize a list's default signup form.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**SignupForm1**](SignupForm1.md)|  | 

### Return type

[**SignupForm**](Signup Form.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostListsIdWebhooks**
> ListWebhooks PostListsIdWebhooks(ctx, listId, body)
Add webhook

Create a new webhook for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **body** | [**AddWebhook**](AddWebhook.md)|  | 

### Return type

[**ListWebhooks**](List Webhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreviewASegment**
> CollectionOfSegments PreviewASegment(ctx, listId, optional)
List segments

Get information about all available segments for a specific list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiPreviewASegmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiPreviewASegmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation. | 
 **excludeFields** | [**optional.Interface of []string**](string.md)| A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation. | 
 **count** | **optional.Int32**| The number of records to return. Default value is 10. Maximum value is 1000 | [default to 10]
 **offset** | **optional.Int32**| Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0. | [default to 0]
 **type_** | **optional.String**| Limit results based on segment type. | 
 **sinceCreatedAt** | **optional.String**| Restrict results to segments created after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeCreatedAt** | **optional.String**| Restrict results to segments created before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **includeCleaned** | **optional.Bool**| Include cleaned members in response | 
 **includeTransactional** | **optional.Bool**| Include transactional members in response | 
 **includeUnsubscribed** | **optional.Bool**| Include unsubscribed members in response | 
 **sinceUpdatedAt** | **optional.String**| Restrict results to segments update after the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 
 **beforeUpdatedAt** | **optional.String**| Restrict results to segments update before the set time. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00. | 

### Return type

[**CollectionOfSegments**](Collection of Segments.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutListsIdMembersId**
> ListMembers2 PutListsIdMembersId(ctx, listId, subscriberHash, body, optional)
Add or update list member

Add or update a list member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
  **subscriberHash** | **string**| The MD5 hash of the lowercase version of the list member&#39;s email address. This endpoint also accepts a list member&#39;s email address or contact_id. | 
  **body** | [**AddListMembers2**](AddListMembers2.md)|  | 
 **optional** | ***ListsApiPutListsIdMembersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiPutListsIdMembersIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **skipMergeValidation** | **optional.Bool**| If skip_merge_validation is true, member data will be accepted without merge field values, even if the merge field is usually required. This defaults to false. | 

### Return type

[**ListMembers2**](List Members_2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTagsByName**
> TagSearchResults SearchTagsByName(ctx, listId, optional)
Search for tags on a list by name.

Search for tags on a list by name. If no name is provided, will return all tags on the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **string**| The unique ID for the list. | 
 **optional** | ***ListsApiSearchTagsByNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListsApiSearchTagsByNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The search query used to filter tags.  The search query will be compared to each tag as a prefix, so all tags that have a name starting with this field will be returned. | 

### Return type

[**TagSearchResults**](Tag search results.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

