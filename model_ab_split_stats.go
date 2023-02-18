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

// General stats about different groups of an A/B Split campaign. Does not return information about Multivariate Campaigns.
type AbSplitStats struct {
	A *CampaignA `json:"a,omitempty"`
	B *CampaignB `json:"b,omitempty"`
}