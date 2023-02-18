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

import (
	"time"
)

// Report details about a sent campaign.
type CampaignReport struct {
	// A string that uniquely identifies this campaign.
	Id string `json:"id,omitempty"`
	// The title of the campaign.
	CampaignTitle string `json:"campaign_title,omitempty"`
	// The type of campaign (regular, plain-text, ab_split, rss, automation, variate, or auto).
	Type_ string `json:"type,omitempty"`
	// The unique list id.
	ListId string `json:"list_id,omitempty"`
	// The status of the list used, namely if it's deleted or disabled.
	ListIsActive bool `json:"list_is_active,omitempty"`
	// The name of the list.
	ListName string `json:"list_name,omitempty"`
	// The subject line for the campaign.
	SubjectLine string `json:"subject_line,omitempty"`
	// The preview text for the campaign.
	PreviewText string `json:"preview_text,omitempty"`
	// The total number of emails sent for this campaign.
	EmailsSent int32 `json:"emails_sent,omitempty"`
	// The number of abuse reports generated for this campaign.
	AbuseReports int32 `json:"abuse_reports,omitempty"`
	// The total number of unsubscribed members for this campaign.
	Unsubscribed int32 `json:"unsubscribed,omitempty"`
	// The date and time a campaign was sent in ISO 8601 format.
	SendTime time.Time `json:"send_time,omitempty"`
	// For RSS campaigns, the date and time of the last send in ISO 8601 format.
	RssLastSend time.Time `json:"rss_last_send,omitempty"`
	Bounces *Bounces `json:"bounces,omitempty"`
	Forwards *Forwards `json:"forwards,omitempty"`
	Opens *Opens `json:"opens,omitempty"`
	Clicks *Clicks `json:"clicks,omitempty"`
	FacebookLikes *FacebookLikes `json:"facebook_likes,omitempty"`
	IndustryStats *IndustryStats1 `json:"industry_stats,omitempty"`
	ListStats *ListStats `json:"list_stats,omitempty"`
	AbSplit *AbSplitStats `json:"ab_split,omitempty"`
	// An hourly breakdown of sends, opens, and clicks if a campaign is sent using timewarp.
	Timewarp []CampaignReports1Timewarp `json:"timewarp,omitempty"`
	// An hourly breakdown of the performance of the campaign over the first 24 hours.
	Timeseries []CampaignReports1Timeseries `json:"timeseries,omitempty"`
	ShareReport *ShareReport `json:"share_report,omitempty"`
	Ecommerce *ECommerceReport1 `json:"ecommerce,omitempty"`
	DeliveryStatus *CampaignDeliveryStatus `json:"delivery_status,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
