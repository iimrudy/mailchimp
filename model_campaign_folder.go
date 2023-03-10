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

// A folder used to organize campaigns.
type CampaignFolder struct {
	// The name of the folder.
	Name string `json:"name,omitempty"`
	// A string that uniquely identifies this campaign folder.
	Id string `json:"id,omitempty"`
	// The number of campaigns in the folder.
	Count int32 `json:"count,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
