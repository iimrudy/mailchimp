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

// List signup form.
type SignupForm struct {
	Header *SignupFormHeaderOptions `json:"header,omitempty"`
	// The signup form body content.
	Contents []CollectionOfContentForListSignupForms `json:"contents,omitempty"`
	// An array of objects, each representing an element style for the signup form.
	Styles []CollectionOfElementStyleForListSignupForms `json:"styles,omitempty"`
	// Signup form URL.
	SignupFormUrl string `json:"signup_form_url,omitempty"`
	// The signup form's list id.
	ListId string `json:"list_id,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
