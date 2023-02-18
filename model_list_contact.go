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

// [Contact information displayed in campaign footers](https://mailchimp.com/help/about-campaign-footers/) to comply with international spam laws.
type ListContact struct {
	// The company name for the list.
	Company string `json:"company,omitempty"`
	// The street address for the list contact.
	Address1 string `json:"address1,omitempty"`
	// The street address for the list contact.
	Address2 string `json:"address2,omitempty"`
	// The city for the list contact.
	City string `json:"city,omitempty"`
	// The state for the list contact.
	State string `json:"state,omitempty"`
	// The postal or zip code for the list contact.
	Zip string `json:"zip,omitempty"`
	// A two-character ISO3166 country code. Defaults to US if invalid.
	Country string `json:"country,omitempty"`
	// The phone number for the list contact.
	Phone string `json:"phone,omitempty"`
}
