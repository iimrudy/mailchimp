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

// [A/B Testing](https://mailchimp.com/help/about-ab-testing-campaigns/) options for a campaign.
type AbTestingOptions struct {
	// The type of AB split to run.
	SplitTest string `json:"split_test,omitempty"`
	// How we should evaluate a winner. Based on 'opens', 'clicks', or 'manual'.
	PickWinner string `json:"pick_winner,omitempty"`
	// How unit of time for measuring the winner ('hours' or 'days'). This cannot be changed after a campaign is sent.
	WaitUnits string `json:"wait_units,omitempty"`
	// The amount of time to wait before picking a winner. This cannot be changed after a campaign is sent.
	WaitTime int32 `json:"wait_time,omitempty"`
	// The size of the split groups. Campaigns split based on 'schedule' are forced to have a 50/50 split. Valid split integers are between 1-50.
	SplitSize int32 `json:"split_size,omitempty"`
	// For campaigns split on 'From Name', the name for Group A.
	FromNameA string `json:"from_name_a,omitempty"`
	// For campaigns split on 'From Name', the name for Group B.
	FromNameB string `json:"from_name_b,omitempty"`
	// For campaigns split on 'From Name', the reply-to address for Group A.
	ReplyEmailA string `json:"reply_email_a,omitempty"`
	// For campaigns split on 'From Name', the reply-to address for Group B.
	ReplyEmailB string `json:"reply_email_b,omitempty"`
	// For campaigns split on 'Subject Line', the subject line for Group A.
	SubjectA string `json:"subject_a,omitempty"`
	// For campaigns split on 'Subject Line', the subject line for Group B.
	SubjectB string `json:"subject_b,omitempty"`
	// The send time for Group A.
	SendTimeA time.Time `json:"send_time_a,omitempty"`
	// The send time for Group B.
	SendTimeB time.Time `json:"send_time_b,omitempty"`
	// The send time for the winning version.
	SendTimeWinner string `json:"send_time_winner,omitempty"`
}
