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

// Options for customizing your signup form header.
type SignupFormHeaderOptions struct {
	// Header image URL.
	ImageUrl string `json:"image_url,omitempty"`
	// Header text.
	Text string `json:"text,omitempty"`
	// Image width, in pixels.
	ImageWidth string `json:"image_width,omitempty"`
	// Image height, in pixels.
	ImageHeight string `json:"image_height,omitempty"`
	// Alt text for the image.
	ImageAlt string `json:"image_alt,omitempty"`
	// The URL that the header image will link to.
	ImageLink string `json:"image_link,omitempty"`
	// Image alignment.
	ImageAlign string `json:"image_align,omitempty"`
	// Image border width.
	ImageBorderWidth string `json:"image_border_width,omitempty"`
	// Image border style.
	ImageBorderStyle string `json:"image_border_style,omitempty"`
	// Image border color.
	ImageBorderColor string `json:"image_border_color,omitempty"`
	// Image link target.
	ImageTarget string `json:"image_target,omitempty"`
}
