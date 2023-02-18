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

// [RSS](https://mailchimp.com/help/share-your-blog-posts-with-mailchimp/) options for a campaign.
type RssOptions3 struct {
	// The URL for the RSS feed.
	FeedUrl string `json:"feed_url"`
	// The frequency of the RSS Campaign.
	Frequency string `json:"frequency"`
	Schedule *SendingSchedule1 `json:"schedule,omitempty"`
	// The date the campaign was last sent.
	LastSent time.Time `json:"last_sent,omitempty"`
	// Whether to add CSS to images in the RSS feed to constrain their width in campaigns.
	ConstrainRssImg bool `json:"constrain_rss_img,omitempty"`
}
