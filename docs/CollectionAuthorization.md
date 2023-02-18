# CollectionAuthorization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MayCreate** | **bool** | May the user create additional instances of this resource? | [default to null]
**MaxInstances** | **int32** | How many total instances of this resource are allowed? This is independent of any filter conditions applied to the query. As a special case, -1 indicates unlimited. | [default to null]
**CurrentTotalInstances** | **int32** | How many total instances of this resource are already in use? This is independent of any filter conditions applied to the query. Value may be larger than max_instances. As a special case, -1 is returned when access is unlimited. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


