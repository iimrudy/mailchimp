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

// The average campaign statistics for your list. This won't be present if we haven't calculated it yet for this list.
type ListStats struct {
	// The average number of subscriptions per month for the list.
	SubRate float32 `json:"sub_rate,omitempty"`
	// The average number of unsubscriptions per month for the list.
	UnsubRate float32 `json:"unsub_rate,omitempty"`
	// The average open rate (a percentage represented as a number between 0 and 100) per campaign for the list.
	OpenRate float32 `json:"open_rate,omitempty"`
	// The average click rate (a percentage represented as a number between 0 and 100) per campaign for the list.
	ClickRate float32 `json:"click_rate,omitempty"`
}
