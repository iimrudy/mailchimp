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

// The settings for the Automation workflow.
type AutomationCampaignSettings1 struct {
	// The 'from' name for the Automation (not an email address).
	FromName string `json:"from_name,omitempty"`
	// The reply-to email address for the Automation.
	ReplyTo string `json:"reply_to,omitempty"`
}
