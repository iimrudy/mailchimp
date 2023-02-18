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

// Information about a specific order line.
type EcommerceOrderLineItem struct {
	// A unique identifier for an order line item.
	Id string `json:"id,omitempty"`
	// A unique identifier for the product associated with an order line item.
	ProductId string `json:"product_id,omitempty"`
	// The name of the product for an order line item.
	ProductTitle string `json:"product_title,omitempty"`
	// A unique identifier for the product variant associated with an order line item.
	ProductVariantId string `json:"product_variant_id,omitempty"`
	// The name of the product variant for an order line item.
	ProductVariantTitle string `json:"product_variant_title,omitempty"`
	// The image URL for a product.
	ImageUrl string `json:"image_url,omitempty"`
	// The order line item quantity.
	Quantity int32 `json:"quantity,omitempty"`
	// The order line item price.
	Price float32 `json:"price,omitempty"`
	// The total discount amount applied to a line item.
	Discount float32 `json:"discount,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
