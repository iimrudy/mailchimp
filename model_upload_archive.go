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

// Available when uploading an archive to create campaign content. The archive should include all campaign content and images. [Learn more](https://mailchimp.com/help/import-a-custom-html-template/).
type UploadArchive struct {
	// The base64-encoded representation of the archive file.
	ArchiveContent string `json:"archive_content"`
	// The type of encoded file. Defaults to zip.
	ArchiveType string `json:"archive_type,omitempty"`
}
