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

// Information about the account contact.
type AccountContact struct {
	// The company name for the account.
	Company string `json:"company,omitempty"`
	// The street address for the account contact.
	Addr1 string `json:"addr1,omitempty"`
	// The street address for the account contact.
	Addr2 string `json:"addr2,omitempty"`
	// The city for the account contact.
	City string `json:"city,omitempty"`
	// The state for the account contact.
	State string `json:"state,omitempty"`
	// The zip code for the account contact.
	Zip string `json:"zip,omitempty"`
	// The country for the account contact.
	Country string `json:"country,omitempty"`
}
