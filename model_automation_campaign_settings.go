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
type AutomationCampaignSettings struct {
	// The title of the Automation.
	Title string `json:"title,omitempty"`
	// The 'from' name for the Automation (not an email address).
	FromName string `json:"from_name,omitempty"`
	// The reply-to email address for the Automation.
	ReplyTo string `json:"reply_to,omitempty"`
	// Whether to use Mailchimp Conversation feature to manage replies
	UseConversation bool `json:"use_conversation,omitempty"`
	// The Automation's custom 'To' name, typically the first name [audience field](https://mailchimp.com/help/getting-started-with-merge-tags/).
	ToName string `json:"to_name,omitempty"`
	// Whether Mailchimp [authenticated](https://mailchimp.com/help/about-email-authentication/) the Automation. Defaults to `true`.
	Authenticate bool `json:"authenticate,omitempty"`
	// Whether to automatically append Mailchimp's [default footer](https://mailchimp.com/help/about-campaign-footers/) to the Automation.
	AutoFooter bool `json:"auto_footer,omitempty"`
	// Whether to automatically inline the CSS included with the Automation content.
	InlineCss bool `json:"inline_css,omitempty"`
}