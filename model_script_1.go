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

// The script used to connect your site with Mailchimp.
type Script1 struct {
	// The URL used for any integrations that offer built-in support for connected sites.
	Url string `json:"url,omitempty"`
	// A pre-built script that you can copy-and-paste into your site to integrate it with Mailchimp.
	Fragment string `json:"fragment,omitempty"`
}
