/*
 * Mailchimp Marketing API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.55
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package mailchimp

type BatchUpdateListMembersErrors struct {
	// The email address that could not be added or updated.
	EmailAddress string `json:"email_address,omitempty"`
	// The error message indicating why the email address could not be added or updated.
	Error_ string `json:"error,omitempty"`
	// A unique code that identifies this specifc error.
	ErrorCode string `json:"error_code,omitempty"`
	// If the error is field-related, information about which field is at issue.
	Field string `json:"field,omitempty"`
	// Message indicating how to resolve a field-related error.
	FieldMessage string `json:"field_message,omitempty"`
}
