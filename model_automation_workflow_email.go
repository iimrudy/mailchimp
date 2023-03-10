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

// A summary of an individual Automation workflow email.
type AutomationWorkflowEmail struct {
	// A string that uniquely identifies the Automation email.
	Id string `json:"id,omitempty"`
	// The ID used in the Mailchimp web application. View this automation in your Mailchimp account at `https://{dc}.admin.mailchimp.com/campaigns/show/?id={web_id}`.
	WebId int32 `json:"web_id,omitempty"`
	// A string that uniquely identifies an Automation workflow.
	WorkflowId string `json:"workflow_id,omitempty"`
	// The position of an Automation email in a workflow.
	Position int32 `json:"position,omitempty"`
	Delay *AutomationDelay `json:"delay,omitempty"`
	// The date and time the campaign was created in ISO 8601 format.
	CreateTime time.Time `json:"create_time,omitempty"`
	// The date and time the campaign was started in ISO 8601 format.
	StartTime time.Time `json:"start_time,omitempty"`
	// The link to the campaign's archive version in ISO 8601 format.
	ArchiveUrl string `json:"archive_url,omitempty"`
	// The current status of the campaign.
	Status string `json:"status,omitempty"`
	// The total number of emails sent for this campaign.
	EmailsSent int32 `json:"emails_sent,omitempty"`
	//  The date and time a campaign was sent in ISO 8601 format
	SendTime time.Time `json:"send_time,omitempty"`
	// How the campaign's content is put together ('template', 'drag_and_drop', 'html', 'url').
	ContentType string `json:"content_type,omitempty"`
	// Determines if the automation email needs its blocks refreshed by opening the web-based campaign editor.
	NeedsBlockRefresh bool `json:"needs_block_refresh,omitempty"`
	// Determines if the campaign contains the *|BRAND:LOGO|* merge tag.
	HasLogoMergeTag bool `json:"has_logo_merge_tag,omitempty"`
	Recipients *List2 `json:"recipients,omitempty"`
	Settings *CampaignSettings `json:"settings,omitempty"`
	Tracking *CampaignTrackingOptions `json:"tracking,omitempty"`
	SocialCard *CampaignSocialCard `json:"social_card,omitempty"`
	TriggerSettings *AutomationTrigger `json:"trigger_settings,omitempty"`
	ReportSummary *CampaignReportSummary1 `json:"report_summary,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
