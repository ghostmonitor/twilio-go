/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAssistantResponse struct for ListAssistantResponse
type ListAssistantResponse struct {
	Assistants []AutopilotV1Assistant    `json:"assistants,omitempty"`
	Meta       ListAssistantResponseMeta `json:"meta,omitempty"`
}
