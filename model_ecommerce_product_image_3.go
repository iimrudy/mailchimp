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

// Information about a specific product image.
type EcommerceProductImage3 struct {
	// A unique identifier for the product image.
	Id string `json:"id"`
	// The URL for a product image.
	Url string `json:"url"`
	// The list of product variants using the image.
	VariantIds []string `json:"variant_ids,omitempty"`
}
