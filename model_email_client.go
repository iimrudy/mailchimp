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

// The email client.
type EmailClient struct {
	// The name of the email client.
	Client string `json:"client,omitempty"`
	// The number of subscribed members who used this email client.
	Members int32 `json:"members,omitempty"`
}
