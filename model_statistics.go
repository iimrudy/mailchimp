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

// Stats for the list. Many of these are cached for at least five minutes.
type Statistics struct {
	// The number of active members in the list.
	MemberCount int32 `json:"member_count,omitempty"`
	// The number of contacts in the list, including subscribed, unsubscribed, pending, cleaned, deleted, transactional, and those that need to be reconfirmed. Requires include_total_contacts query parameter to be included.
	TotalContacts int32 `json:"total_contacts,omitempty"`
	// The number of members who have unsubscribed from the list.
	UnsubscribeCount int32 `json:"unsubscribe_count,omitempty"`
	// The number of members cleaned from the list.
	CleanedCount int32 `json:"cleaned_count,omitempty"`
	// The number of active members in the list since the last campaign was sent.
	MemberCountSinceSend int32 `json:"member_count_since_send,omitempty"`
	// The number of members who have unsubscribed since the last campaign was sent.
	UnsubscribeCountSinceSend int32 `json:"unsubscribe_count_since_send,omitempty"`
	// The number of members cleaned from the list since the last campaign was sent.
	CleanedCountSinceSend int32 `json:"cleaned_count_since_send,omitempty"`
	// The number of campaigns in any status that use this list.
	CampaignCount int32 `json:"campaign_count,omitempty"`
	// The date and time the last campaign was sent to this list in ISO 8601 format. This is updated when a campaign is sent to 10 or more recipients.
	CampaignLastSent time.Time `json:"campaign_last_sent,omitempty"`
	// The number of merge fields ([audience field](https://mailchimp.com/help/getting-started-with-merge-tags/)) for this list (doesn't include EMAIL).
	MergeFieldCount int32 `json:"merge_field_count,omitempty"`
	// The average number of subscriptions per month for the list (not returned if we haven't calculated it yet).
	AvgSubRate float32 `json:"avg_sub_rate,omitempty"`
	// The average number of unsubscriptions per month for the list (not returned if we haven't calculated it yet).
	AvgUnsubRate float32 `json:"avg_unsub_rate,omitempty"`
	// The target number of subscriptions per month for the list to keep it growing (not returned if we haven't calculated it yet).
	TargetSubRate float32 `json:"target_sub_rate,omitempty"`
	// The average open rate (a percentage represented as a number between 0 and 100) per campaign for the list (not returned if we haven't calculated it yet).
	OpenRate float32 `json:"open_rate,omitempty"`
	// The average click rate (a percentage represented as a number between 0 and 100) per campaign for the list (not returned if we haven't calculated it yet).
	ClickRate float32 `json:"click_rate,omitempty"`
	// The date and time of the last time someone subscribed to this list in ISO 8601 format.
	LastSubDate time.Time `json:"last_sub_date,omitempty"`
	// The date and time of the last time someone unsubscribed from this list in ISO 8601 format.
	LastUnsubDate time.Time `json:"last_unsub_date,omitempty"`
}
