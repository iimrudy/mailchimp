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

// Available triggers for Automation workflows.
type AutomationTrigger struct {
	// The type of Automation workflow.
	WorkflowType string `json:"workflow_type"`
	// The title of the workflow type.
	WorkflowTitle string `json:"workflow_title,omitempty"`
	Runtime *AutomationWorkflowRuntimeSettings `json:"runtime,omitempty"`
	// The number of emails in the Automation workflow.
	WorkflowEmailsCount int32 `json:"workflow_emails_count,omitempty"`
}
