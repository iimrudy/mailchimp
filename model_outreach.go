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

// The outreach associated with this order. For example, an email campaign or Facebook ad.
type Outreach struct {
	// A unique identifier for the outreach. Can be an email campaign ID.
	Id string `json:"id,omitempty"`
	// The name for the outreach.
	Name string `json:"name,omitempty"`
	// The type of the outreach.
	Type_ string `json:"type,omitempty"`
	// The date and time the Outreach was published in ISO 8601 format.
	PublishedTime time.Time `json:"published_time,omitempty"`
}
