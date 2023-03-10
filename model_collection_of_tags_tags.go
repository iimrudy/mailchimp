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

type CollectionOfTagsTags struct {
	// The unique id for the tag.
	Id int32 `json:"id,omitempty"`
	// The name of the tag.
	Name string `json:"name,omitempty"`
	// The date and time the tag was added to the list member in ISO 8601 format.
	DateAdded time.Time `json:"date_added,omitempty"`
}
