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

// For email question types, how many are new, known, or unknown contacts.
type ContactCounts struct {
	// The number of known contacts that responded to this survey.
	Known int32 `json:"known,omitempty"`
	// The number of unknown contacts that responded to this survey.
	Unknown int32 `json:"unknown,omitempty"`
	// The number of new contacts that responded to this survey.
	New int32 `json:"new,omitempty"`
}
