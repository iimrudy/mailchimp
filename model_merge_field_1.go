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

// A [merge field](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.
type MergeField1 struct {
	// The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts).
	Tag string `json:"tag,omitempty"`
	// The name of the merge field (audience field).
	Name string `json:"name"`
	// The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field.
	Type_ string `json:"type"`
	// Whether the merge field is required to import a contact.
	Required bool `json:"required,omitempty"`
	// The default value for the merge field if `null`.
	DefaultValue string `json:"default_value,omitempty"`
	// Whether the merge field is displayed on the signup form.
	Public bool `json:"public,omitempty"`
	// The order that the merge field displays on the list signup form.
	DisplayOrder int32 `json:"display_order,omitempty"`
	Options *MergeFieldOptions1 `json:"options,omitempty"`
	// Extra text to help the subscriber fill out the form.
	HelpText string `json:"help_text,omitempty"`
}
