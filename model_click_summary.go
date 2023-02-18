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

// A summary of the click-throughs on the campaign's URL.
type ClickSummary struct {
	// The total number of clicks to the campaign's URL.
	Clicks int32 `json:"clicks,omitempty"`
	// The timestamp for the first click to the URL.
	FirstClick time.Time `json:"first_click,omitempty"`
	// The timestamp for the last click to the URL.
	LastClick time.Time `json:"last_click,omitempty"`
	// A summary of the top click locations.
	Locations []Location3 `json:"locations,omitempty"`
}
