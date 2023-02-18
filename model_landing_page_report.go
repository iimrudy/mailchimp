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

// A summary of an individual landing page's settings and content.
type LandingPageReport struct {
	// A string that uniquely identifies this landing page.
	Id string `json:"id,omitempty"`
	// The name of this landing page the user will see.
	Name string `json:"name,omitempty"`
	// The name of the landing page the user's customers will see.
	Title string `json:"title,omitempty"`
	// The landing page url.
	Url string `json:"url,omitempty"`
	// The time this landing page was published.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// The time this landing page was unpublished.
	UnpublishedAt time.Time `json:"unpublished_at,omitempty"`
	// The status of the landing page.
	Status string `json:"status,omitempty"`
	// The list id connected to this landing page.
	ListId string `json:"list_id,omitempty"`
	// The number of visits to this landing pages.
	Visits int32 `json:"visits,omitempty"`
	// The number of unique visits to this landing pages.
	UniqueVisits int32 `json:"unique_visits,omitempty"`
	// The number of subscribes to this landing pages.
	Subscribes int32 `json:"subscribes,omitempty"`
	// The number of clicks to this landing pages.
	Clicks int32 `json:"clicks,omitempty"`
	// The percentage of people who visited your landing page and were added to your list.
	ConversionRate float32 `json:"conversion_rate,omitempty"`
	Timeseries *LandingPageReportTimeseries `json:"timeseries,omitempty"`
	Ecommerce *LandingPageReportEcommerce `json:"ecommerce,omitempty"`
	// The ID used in the Mailchimp web application.
	WebId int32 `json:"web_id,omitempty"`
	// List Name
	ListName string `json:"list_name,omitempty"`
	// A list of tags associated to the landing page.
	SignupTags []Tag `json:"signup_tags,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
