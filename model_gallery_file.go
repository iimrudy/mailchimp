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

// An individual file listed in the File Manager.
type GalleryFile struct {
	// The unique id of the file.
	Id int32 `json:"id,omitempty"`
	// The id of the folder.
	FolderId int32 `json:"folder_id,omitempty"`
	// The type of file in the File Manager.
	Type_ string `json:"type,omitempty"`
	// The name of the file.
	Name string `json:"name,omitempty"`
	// The url of the full-size file.
	FullSizeUrl string `json:"full_size_url,omitempty"`
	// The url of the thumbnail preview.
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	// The size of the file in bytes.
	Size int32 `json:"size,omitempty"`
	// The date and time a file was added to the File Manager in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The username of the profile that uploaded the file.
	CreatedBy string `json:"created_by,omitempty"`
	// The width of the image.
	Width int32 `json:"width,omitempty"`
	// The height of an image.
	Height int32 `json:"height,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
