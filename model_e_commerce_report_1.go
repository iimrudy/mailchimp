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

// E-Commerce stats for a campaign.
type ECommerceReport1 struct {
	// The total orders for a campaign.
	TotalOrders int32 `json:"total_orders,omitempty"`
	// The total spent for a campaign. Calculated as the sum of all order totals with no deductions.
	TotalSpent float32 `json:"total_spent,omitempty"`
	// The total revenue for a campaign. Calculated as the sum of all order totals minus shipping and tax totals.
	TotalRevenue float32 `json:"total_revenue,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}
