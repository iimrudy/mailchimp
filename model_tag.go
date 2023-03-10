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

type Tag struct {
	// The unique id for the tag.
	TagId int32 `json:"tag_id,omitempty"`
	// The name of the tag.
	TagName string `json:"tag_name,omitempty"`
}
