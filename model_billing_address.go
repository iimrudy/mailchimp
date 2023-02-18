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

// The billing address for the order.
type BillingAddress struct {
	// The name associated with an order's billing address.
	Name string `json:"name,omitempty"`
	// The billing address for the order.
	Address1 string `json:"address1,omitempty"`
	// An additional field for the billing address.
	Address2 string `json:"address2,omitempty"`
	// The city in the billing address.
	City string `json:"city,omitempty"`
	// The state or normalized province in the billing address.
	Province string `json:"province,omitempty"`
	// The two-letter code for the province or state in the billing address.
	ProvinceCode string `json:"province_code,omitempty"`
	// The postal or zip code in the billing address.
	PostalCode string `json:"postal_code,omitempty"`
	// The country in the billing address.
	Country string `json:"country,omitempty"`
	// The two-letter code for the country in the billing address.
	CountryCode string `json:"country_code,omitempty"`
	// The longitude for the billing address location.
	Longitude float32 `json:"longitude,omitempty"`
	// The latitude for the billing address location.
	Latitude float32 `json:"latitude,omitempty"`
	// The phone number for the billing address.
	Phone string `json:"phone,omitempty"`
	// The company associated with the billing address.
	Company string `json:"company,omitempty"`
}
