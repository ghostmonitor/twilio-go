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

// ListWebhookResponse struct for ListWebhookResponse
type ListWebhookResponse struct {
	Meta     ListAssistantResponseMeta `json:"meta,omitempty"`
	Webhooks []AutopilotV1Webhook      `json:"webhooks,omitempty"`
}
