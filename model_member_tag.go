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

// Add or remove tags on a member by declaring a tag either active or inactive on a member.
type MemberTag struct {
	// The name of the tag.
	Name string `json:"name"`
	// The status for the tag on the member, pass in active to add a tag or inactive to remove it.
	Status string `json:"status"`
}
