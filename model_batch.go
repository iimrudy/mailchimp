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

// The status of a batch request
type Batch struct {
	// A string that uniquely identifies this batch request.
	Id string `json:"id,omitempty"`
	// The status of the batch call. [Learn more](https://mailchimp.com/developer/marketing/guides/run-async-requests-batch-endpoint/#check-the-status-of-a-batch-operation) about the batch operation status.
	Status string `json:"status,omitempty"`
	// The total number of operations to complete as part of this batch request. For GET requests requiring pagination, each page counts as a separate operation.
	TotalOperations int32 `json:"total_operations,omitempty"`
	// The number of completed operations. This includes operations that returned an error.
	FinishedOperations int32 `json:"finished_operations,omitempty"`
	// The number of completed operations that returned an error.
	ErroredOperations int32 `json:"errored_operations,omitempty"`
	// The date and time when the server received the batch request in ISO 8601 format.
	SubmittedAt time.Time `json:"submitted_at,omitempty"`
	// The date and time when all operations in the batch request completed in ISO 8601 format.
	CompletedAt time.Time `json:"completed_at,omitempty"`
	// The URL of the gzipped archive of the results of all the operations.
	ResponseBodyUrl string `json:"response_body_url,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}
