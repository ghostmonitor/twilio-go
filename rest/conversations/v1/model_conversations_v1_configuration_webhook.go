/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ConversationsV1ConfigurationWebhook struct for ConversationsV1ConfigurationWebhook
type ConversationsV1ConfigurationWebhook struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The list of webhook event triggers that are enabled for this Service.
	Filters *[]string `json:"filters,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	Method *string `json:"method,omitempty"`
	// The absolute url the post-event webhook request should be sent to.
	PostWebhookUrl *string `json:"post_webhook_url,omitempty"`
	// The absolute url the pre-event webhook request should be sent to.
	PreWebhookUrl *string `json:"pre_webhook_url,omitempty"`
	// The routing target of the webhook.
	Target *string `json:"target,omitempty"`
	// An absolute URL for this webhook.
	Url *string `json:"url,omitempty"`
}
