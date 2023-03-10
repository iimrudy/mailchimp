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

// A single question and the response to that question.
type Response struct {
	// The unique ID for this question.
	QuestionId string `json:"question_id,omitempty"`
	// The type of question this is.
	QuestionType string `json:"question_type,omitempty"`
	// The survey question.
	Query string `json:"query,omitempty"`
	// The answer to this survey question.
	Answer string `json:"answer,omitempty"`
}
