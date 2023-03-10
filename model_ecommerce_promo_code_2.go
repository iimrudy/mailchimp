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

// Information about an Ecommerce Store's specific Promo Code.
type EcommercePromoCode2 struct {
	// The discount code. Restricted to UTF-8 characters with max length 50.
	Code string `json:"code,omitempty"`
	// The url that should be used in the promotion campaign restricted to UTF-8 characters with max length 2000.
	RedemptionUrl string `json:"redemption_url,omitempty"`
	// Number of times promo code has been used.
	UsageCount int32 `json:"usage_count,omitempty"`
	// Whether the promo code is currently enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The date and time the promotion was created in ISO 8601 format.
	CreatedAtForeign time.Time `json:"created_at_foreign,omitempty"`
	// The date and time the promotion was updated in ISO 8601 format.
	UpdatedAtForeign time.Time `json:"updated_at_foreign,omitempty"`
}
