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

// Extra options for some merge field types.
type MergeFieldOptions struct {
	// In an address field, the default country code if none supplied.
	DefaultCountry int32 `json:"default_country,omitempty"`
	// In a phone field, the phone number type: US or International.
	PhoneFormat string `json:"phone_format,omitempty"`
	// In a date or birthday field, the format of the date.
	DateFormat string `json:"date_format,omitempty"`
	// In a radio or dropdown non-group field, the available options for contacts to pick from.
	Choices []string `json:"choices,omitempty"`
	// In a text field, the default length of the text field.
	Size int32 `json:"size,omitempty"`
}
