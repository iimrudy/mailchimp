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

// Collection of Content for List Signup Forms.
type CollectionOfContentForListSignupForms struct {
	// The content section name.
	Section string `json:"section,omitempty"`
	// The content section text.
	Value string `json:"value,omitempty"`
}
