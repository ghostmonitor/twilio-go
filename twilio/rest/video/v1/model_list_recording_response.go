/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListRecordingResponse struct for ListRecordingResponse
type ListRecordingResponse struct {
	Meta       ListCompositionHookResponseMeta `json:"Meta,omitempty"`
	Recordings []VideoV1Recording              `json:"Recordings,omitempty"`
}