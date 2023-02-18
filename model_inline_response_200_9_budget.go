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

type InlineResponse2009Budget struct {
	// Duration of the ad in seconds
	Duration int32 `json:"duration,omitempty"`
	// Total budget of the ad
	TotalAmount float32 `json:"total_amount,omitempty"`
	// Currency code
	CurrencyCode string `json:"currency_code,omitempty"`
}
