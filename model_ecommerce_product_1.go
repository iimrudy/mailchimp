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

// Information about a specific product.
type EcommerceProduct1 struct {
	// A unique identifier for the product.
	Id string `json:"id"`
	// The title of a product.
	Title string `json:"title"`
	// The handle of a product.
	Handle string `json:"handle,omitempty"`
	// The URL for a product.
	Url string `json:"url,omitempty"`
	// The description of a product.
	Description string `json:"description,omitempty"`
	// The type of product.
	Type_ string `json:"type,omitempty"`
	// The vendor for a product.
	Vendor string `json:"vendor,omitempty"`
	// The image URL for a product.
	ImageUrl string `json:"image_url,omitempty"`
	// An array of the product's variants. At least one variant is required for each product. A variant can use the same `id` and `title` as the parent product.
	Variants []EcommerceProductVariant1 `json:"variants"`
	// An array of the product's images.
	Images []EcommerceProductImage1 `json:"images,omitempty"`
	// The date and time the product was published.
	PublishedAtForeign time.Time `json:"published_at_foreign,omitempty"`
}
