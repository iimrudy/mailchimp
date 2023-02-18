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

// Information about a specific template.
type TemplateInstance struct {
	// The individual id for the template.
	Id int32 `json:"id,omitempty"`
	// The type of template (user, base, or gallery).
	Type_ string `json:"type,omitempty"`
	// The name of the template.
	Name string `json:"name,omitempty"`
	// Whether the template uses the drag and drop editor.
	DragAndDrop bool `json:"drag_and_drop,omitempty"`
	// Whether the template contains media queries to make it responsive.
	Responsive bool `json:"responsive,omitempty"`
	// If available, the category the template is listed in.
	Category string `json:"category,omitempty"`
	// The date and time the template was created in ISO 8601 format.
	DateCreated time.Time `json:"date_created,omitempty"`
	// The date and time the template was edited in ISO 8601 format.
	DateEdited time.Time `json:"date_edited,omitempty"`
	// The login name for template's creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The login name who last edited the template.
	EditedBy string `json:"edited_by,omitempty"`
	// User templates are not 'deleted,' but rather marked as 'inactive.' Returns whether the template is still active.
	Active bool `json:"active,omitempty"`
	// The id of the folder the template is currently in.
	FolderId string `json:"folder_id,omitempty"`
	// If available, the URL for a thumbnail of the template.
	Thumbnail string `json:"thumbnail,omitempty"`
	// The URL used for [template sharing](https://mailchimp.com/help/share-a-template/).
	ShareUrl string `json:"share_url,omitempty"`
	// How the template's content is put together.
	ContentType string `json:"content_type,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}