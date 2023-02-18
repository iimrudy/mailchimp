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

// Information about a specific order.
type EcommerceOrder1 struct {
	// A unique identifier for the order.
	Id string `json:"id"`
	Customer *EcommerceCustomer1 `json:"customer"`
	// A string that uniquely identifies the campaign for an order.
	CampaignId string `json:"campaign_id,omitempty"`
	// The URL for the page where the buyer landed when entering the shop.
	LandingSite string `json:"landing_site,omitempty"`
	// The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FinancialStatus string `json:"financial_status,omitempty"`
	// The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FulfillmentStatus string `json:"fulfillment_status,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode string `json:"currency_code"`
	// The total for the order.
	OrderTotal float32 `json:"order_total"`
	// The URL for the order.
	OrderUrl string `json:"order_url,omitempty"`
	// The total amount of the discounts to be applied to the price of the order.
	DiscountTotal float32 `json:"discount_total,omitempty"`
	// The tax total for the order.
	TaxTotal float32 `json:"tax_total,omitempty"`
	// The shipping total for the order.
	ShippingTotal float32 `json:"shipping_total,omitempty"`
	// The Mailchimp tracking code for the order. Uses the 'mc_tc' parameter in E-Commerce tracking URLs.
	TrackingCode string `json:"tracking_code,omitempty"`
	// The date and time the order was processed in ISO 8601 format.
	ProcessedAtForeign time.Time `json:"processed_at_foreign,omitempty"`
	// The date and time the order was cancelled in ISO 8601 format. Note: passing a value for this parameter will cancel the order being created.
	CancelledAtForeign time.Time `json:"cancelled_at_foreign,omitempty"`
	// The date and time the order was updated in ISO 8601 format.
	UpdatedAtForeign time.Time `json:"updated_at_foreign,omitempty"`
	ShippingAddress *ShippingAddress1 `json:"shipping_address,omitempty"`
	BillingAddress *BillingAddress1 `json:"billing_address,omitempty"`
	// The promo codes applied on the order
	Promos []EcommercestoresstoreIdordersPromos `json:"promos,omitempty"`
	// An array of the order's line items.
	Lines []EcommerceOrderLineItem1 `json:"lines"`
	Outreach *Outreach1 `json:"outreach,omitempty"`
	// The tracking number associated with the order.
	TrackingNumber string `json:"tracking_number,omitempty"`
	// The tracking carrier associated with the order.
	TrackingCarrier string `json:"tracking_carrier,omitempty"`
	// The tracking URL associated with the order.
	TrackingUrl string `json:"tracking_url,omitempty"`
}