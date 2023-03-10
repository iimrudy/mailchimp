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

// A summary of an individual campaign's settings and content.
type Campaign2 struct {
	Recipients *List5 `json:"recipients,omitempty"`
	Settings *CampaignSettings4 `json:"settings"`
	VariateSettings *AbTestOptions1 `json:"variate_settings,omitempty"`
	Tracking *CampaignTrackingOptions1 `json:"tracking,omitempty"`
	RssOpts *RssOptions2 `json:"rss_opts,omitempty"`
	SocialCard *CampaignSocialCard `json:"social_card,omitempty"`
}
