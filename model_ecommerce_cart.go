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

import (
	"time"
)

// Information about a specific cart.
type EcommerceCart struct {
	// A unique identifier for the cart.
	Id string `json:"id,omitempty"`
	Customer *EcommerceCustomer `json:"customer,omitempty"`
	// A string that uniquely identifies the campaign associated with a cart.
	CampaignId string `json:"campaign_id,omitempty"`
	// The URL for the cart. This parameter is required for [Abandoned Cart](https://mailchimp.com/help/create-an-abandoned-cart-email/) automations.
	CheckoutUrl string `json:"checkout_url,omitempty"`
	// The three-letter ISO 4217 code for the currency that the cart uses.
	CurrencyCode string `json:"currency_code,omitempty"`
	// The order total for the cart.
	OrderTotal float32 `json:"order_total,omitempty"`
	// The total tax for the cart.
	TaxTotal float32 `json:"tax_total,omitempty"`
	// An array of the cart's line items.
	Lines []EcommerceCartLineItem `json:"lines,omitempty"`
	// The date and time the cart was created in ISO 8601 format.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The date and time the cart was last updated in ISO 8601 format.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
