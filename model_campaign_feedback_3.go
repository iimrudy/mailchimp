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

// A specific feedback message from a specific campaign.
type CampaignFeedback3 struct {
	// The block id for the editable block that the feedback addresses.
	BlockId int32 `json:"block_id,omitempty"`
	// The content of the feedback.
	Message string `json:"message,omitempty"`
	// The status of feedback.
	IsComplete bool `json:"is_complete,omitempty"`
}
