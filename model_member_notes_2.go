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

// A specific note for a specific member.
type MemberNotes2 struct {
	// The content of the note. Note length is limited to 1,000 characters.
	Note string `json:"note,omitempty"`
}
