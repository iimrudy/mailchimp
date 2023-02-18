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

// A detailed report of any campaign emails that were opened by a list member.
type OpenDetailReport struct {
	// An array of objects, each representing a list member who opened a campaign email. Each members object will contain information about the number of total opens by a single member, as well as timestamps for each open event.
	Members []OpenActivity `json:"members,omitempty"`
	// The campaign id.
	CampaignId string `json:"campaign_id,omitempty"`
	// The total number of opens matching the query regardless of pagination.
	TotalOpens int32 `json:"total_opens,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
