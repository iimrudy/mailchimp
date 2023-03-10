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

type Operations struct {
	// The HTTP method to use for the operation.
	Method string `json:"method"`
	// The relative path to use for the operation.
	Path string `json:"path"`
	// Any request query parameters. Example parameters: {\"count\":10, \"offset\":0}
	Params interface{} `json:"params,omitempty"`
	// A string containing the JSON body to use with the request.
	Body string `json:"body,omitempty"`
	// An optional client-supplied id returned with the operation results.
	OperationId string `json:"operation_id,omitempty"`
}
