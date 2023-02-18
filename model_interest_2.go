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

// Assign subscribers to interests to group them together. Interests are referred to as 'group names' in the Mailchimp application.
type Interest2 struct {
	// The name of the interest. This can be shown publicly on a subscription form.
	Name string `json:"name"`
	// The display order for interests.
	DisplayOrder int32 `json:"display_order,omitempty"`
}