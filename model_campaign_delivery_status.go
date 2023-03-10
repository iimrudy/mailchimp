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

// Updates on campaigns in the process of sending.
type CampaignDeliveryStatus struct {
	// Whether Campaign Delivery Status is enabled for this account and campaign.
	Enabled bool `json:"enabled,omitempty"`
	// Whether a campaign send can be canceled.
	CanCancel bool `json:"can_cancel,omitempty"`
	// The current state of a campaign delivery.
	Status string `json:"status,omitempty"`
	// The total number of emails confirmed sent for this campaign so far.
	EmailsSent int32 `json:"emails_sent,omitempty"`
	// The total number of emails canceled for this campaign.
	EmailsCanceled int32 `json:"emails_canceled,omitempty"`
}
