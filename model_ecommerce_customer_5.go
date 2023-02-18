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

// Information about a specific customer. Orders for existing customers should include only the `id` parameter in the `customer` object body.
type EcommerceCustomer5 struct {
	// The customer's opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don't opt in to your Mailchimp list [will be added as `Transactional` members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).
	OptInStatus bool `json:"opt_in_status,omitempty"`
	// The customer's company.
	Company string `json:"company,omitempty"`
	// The customer's first name.
	FirstName string `json:"first_name,omitempty"`
	// The customer's last name.
	LastName string `json:"last_name,omitempty"`
	Address *Address `json:"address,omitempty"`
}
