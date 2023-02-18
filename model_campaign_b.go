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

// Stats for Campaign B.
type CampaignB struct {
	// Bounces for Campaign B.
	Bounces int32 `json:"bounces,omitempty"`
	// Abuse reports for Campaign B.
	AbuseReports int32 `json:"abuse_reports,omitempty"`
	// Unsubscribes for Campaign B.
	Unsubs int32 `json:"unsubs,omitempty"`
	// Recipients clicks for Campaign B.
	RecipientClicks int32 `json:"recipient_clicks,omitempty"`
	// Forwards for Campaign B.
	Forwards int32 `json:"forwards,omitempty"`
	// Opens for forwards from Campaign B.
	ForwardsOpens int32 `json:"forwards_opens,omitempty"`
	// Opens for Campaign B.
	Opens int32 `json:"opens,omitempty"`
	// The last open for Campaign B.
	LastOpen string `json:"last_open,omitempty"`
	// Unique opens for Campaign B.
	UniqueOpens int32 `json:"unique_opens,omitempty"`
}