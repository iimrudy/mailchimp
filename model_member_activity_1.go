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

// A summary of the interaction with the campaign.
type MemberActivity1 struct {
	// The date and time recorded for the action in ISO 8601 format.
	Timestamp time.Time `json:"timestamp,omitempty"`
}
