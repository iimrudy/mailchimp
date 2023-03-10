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

type OrdersPromos struct {
	// The Promo Code
	Code string `json:"code,omitempty"`
	// The amount of discount applied on the total price. For example if the total cost was $100 and the customer paid $95.5, amount_discounted will be 4.5 For free shipping set amount_discounted to 0
	AmountDiscounted float32 `json:"amount_discounted,omitempty"`
	// Type of discount. For free shipping set type to fixed
	Type_ string `json:"type,omitempty"`
}
