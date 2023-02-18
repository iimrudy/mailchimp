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

// Open and click rates for this subscriber.
type SubscriberStats struct {
	// A subscriber's average open rate.
	AvgOpenRate float32 `json:"avg_open_rate,omitempty"`
	// A subscriber's average clickthrough rate.
	AvgClickRate float32 `json:"avg_click_rate,omitempty"`
}