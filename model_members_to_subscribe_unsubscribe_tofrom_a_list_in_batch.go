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

// Members to subscribe to or unsubscribe from a list.
type MembersToSubscribeUnsubscribeTofromAListInBatch struct {
	// An array of objects, each representing an email address and the subscription status for a specific list. Up to 500 members may be added or updated with each API call.
	Members []AddListMembers `json:"members"`
	// Whether this batch operation will replace all existing tags with tags in request.
	SyncTags bool `json:"sync_tags,omitempty"`
	// Whether this batch operation will change existing members' subscription status.
	UpdateExisting bool `json:"update_existing,omitempty"`
}
