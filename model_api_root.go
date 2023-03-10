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

// The API root resource links to all other resources available in the API.
type ApiRoot struct {
	// The Mailchimp account ID.
	AccountId string `json:"account_id,omitempty"`
	// The ID associated with the user who owns this API key. If you can login to multiple accounts, this ID will be the same for each account.
	LoginId string `json:"login_id,omitempty"`
	// The name of the account.
	AccountName string `json:"account_name,omitempty"`
	// The account email address.
	Email string `json:"email,omitempty"`
	// The first name tied to the account.
	FirstName string `json:"first_name,omitempty"`
	// The last name tied to the account.
	LastName string `json:"last_name,omitempty"`
	// The username tied to the account.
	Username string `json:"username,omitempty"`
	// URL of the avatar for the user.
	AvatarUrl string `json:"avatar_url,omitempty"`
	// The [user role](https://mailchimp.com/help/manage-user-levels-in-your-account/) for the account.
	Role string `json:"role,omitempty"`
	// The date and time that the account was created in ISO 8601 format.
	MemberSince time.Time `json:"member_since,omitempty"`
	// The type of pricing plan the account is on.
	PricingPlanType string `json:"pricing_plan_type,omitempty"`
	// Date of first payment for monthly plans.
	FirstPayment time.Time `json:"first_payment,omitempty"`
	// The timezone currently set for the account.
	AccountTimezone string `json:"account_timezone,omitempty"`
	// The user-specified industry associated with the account.
	AccountIndustry string `json:"account_industry,omitempty"`
	Contact *AccountContact `json:"contact,omitempty"`
	// Legacy - whether the account includes [Mailchimp Pro](https://mailchimp.com/help/about-mailchimp-pro/).
	ProEnabled bool `json:"pro_enabled,omitempty"`
	// The date and time of the last login for this account in ISO 8601 format.
	LastLogin time.Time `json:"last_login,omitempty"`
	// The total number of subscribers across all lists in the account.
	TotalSubscribers int32 `json:"total_subscribers,omitempty"`
	IndustryStats *IndustryStats `json:"industry_stats,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
