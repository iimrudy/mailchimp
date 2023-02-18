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

// An individual store in an account.
type EcommerceStore1 struct {
	// The unique identifier for the store.
	Id string `json:"id"`
	// The unique identifier for the list associated with the store. The `list_id` for a specific store cannot change.
	ListId string `json:"list_id"`
	// The name of the store.
	Name string `json:"name"`
	// The e-commerce platform of the store.
	Platform string `json:"platform,omitempty"`
	// The store domain. This parameter is required for Connected Sites and Google Ads.
	Domain string `json:"domain,omitempty"`
	// Whether to disable automations because the store is currently [syncing](https://mailchimp.com/developer/marketing/docs/e-commerce/#pausing-store-automations).
	IsSyncing bool `json:"is_syncing,omitempty"`
	// The email address for the store.
	EmailAddress string `json:"email_address,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode string `json:"currency_code"`
	// The currency format for the store. For example: `$`, `£`, etc.
	MoneyFormat string `json:"money_format,omitempty"`
	// The primary locale for the store. For example: `en`, `de`, etc.
	PrimaryLocale string `json:"primary_locale,omitempty"`
	// The timezone for the store.
	Timezone string `json:"timezone,omitempty"`
	// The store phone number.
	Phone string `json:"phone,omitempty"`
	Address *Address1 `json:"address,omitempty"`
}
