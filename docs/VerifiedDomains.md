# VerifiedDomains

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The name of this domain. | [optional] [default to null]
**Verified** | **bool** | Whether the domain has been verified for sending. | [optional] [default to null]
**Authenticated** | **bool** | Whether domain authentication is enabled for this domain. | [optional] [default to null]
**VerificationEmail** | **string** | The e-mail address receiving the two-factor challenge for this domain. | [optional] [default to null]
**VerificationSent** | [**time.Time**](time.Time.md) | The date/time that the two-factor challenge was sent to the verification email. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


