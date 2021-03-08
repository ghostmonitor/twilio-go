/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListSampleResponse struct for ListSampleResponse
type ListSampleResponse struct {
	Meta    ListDayResponseMeta                    `json:"Meta,omitempty"`
	Samples []PreviewUnderstandAssistantTaskSample `json:"Samples,omitempty"`
}