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

// A single marketing permission a subscriber has either opted-in to or opted-out of.
type MarketingPermission1 struct {
	// The id for the marketing permission on the list
	MarketingPermissionId string `json:"marketing_permission_id,omitempty"`
	// If the subscriber has opted-in to the marketing permission.
	Enabled bool `json:"enabled,omitempty"`
}
