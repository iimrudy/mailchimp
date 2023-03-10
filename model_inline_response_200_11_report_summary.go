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

// Report summary of facebook ad
type InlineResponse20011ReportSummary struct {
	Reach int32 `json:"reach,omitempty"`
	Impressions int32 `json:"impressions,omitempty"`
	Clicks int32 `json:"clicks,omitempty"`
	ClickRate float32 `json:"click_rate,omitempty"`
	UniqueClicks int32 `json:"unique_clicks,omitempty"`
	FirstTimeBuyers int32 `json:"first_time_buyers,omitempty"`
	Ecommerce *InlineResponse20011ReportSummaryEcommerce `json:"ecommerce,omitempty"`
	TotalProductsSold int32 `json:"total_products_sold,omitempty"`
	TotalOrders int32 `json:"total_orders,omitempty"`
	AverageOrderAmount *InlineResponse20011ReportSummaryAverageOrderAmount `json:"average_order_amount,omitempty"`
	CostPerClick *InlineResponse20011ReportSummaryAverageOrderAmount `json:"cost_per_click,omitempty"`
	AverageDailyBudget *InlineResponse20011ReportSummaryAverageOrderAmount `json:"average_daily_budget,omitempty"`
	Likes int32 `json:"likes,omitempty"`
	Comments int32 `json:"comments,omitempty"`
	Shares int32 `json:"shares,omitempty"`
	HasExtendedAdDuration bool `json:"has_extended_ad_duration,omitempty"`
	ExtendedAt *InlineResponse20011ReportSummaryExtendedAt `json:"extended_at,omitempty"`
	ReturnOnInvestment float32 `json:"return_on_investment,omitempty"`
}
