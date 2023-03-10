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

type InlineResponse2007Products struct {
	Title string `json:"title,omitempty"`
	Sku string `json:"sku,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	TotalRevenue float32 `json:"total_revenue,omitempty"`
	TotalPurchased float32 `json:"total_purchased,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
	RecommendationTotal int32 `json:"recommendation_total,omitempty"`
	RecommendationPurchased int32 `json:"recommendation_purchased,omitempty"`
}
