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

type CampaignscampaignIdcontentVariateContents struct {
	// The label used to identify the content option.
	ContentLabel string `json:"content_label"`
	// The plain-text portion of the campaign. If left unspecified, we'll generate this automatically.
	PlainText string `json:"plain_text,omitempty"`
	// The raw HTML for the campaign.
	Html string `json:"html,omitempty"`
	// When importing a campaign, the URL for the HTML.
	Url string `json:"url,omitempty"`
	Template *TemplateContent1 `json:"template,omitempty"`
	Archive *UploadArchive `json:"archive,omitempty"`
}
