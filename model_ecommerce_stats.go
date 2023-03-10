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

// Ecommerce stats for the list member if the list is attached to a store.
type EcommerceStats struct {
	// The total revenue the list member has brought in.
	TotalRevenue float32 `json:"total_revenue,omitempty"`
	// The total number of orders placed by the list member.
	NumberOfOrders float32 `json:"number_of_orders,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode string `json:"currency_code,omitempty"`
}
