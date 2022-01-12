/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListConversationMessageReceiptResponse struct for ListConversationMessageReceiptResponse
type ListConversationMessageReceiptResponse struct {
	DeliveryReceipts []ConversationsV1ConversationMessageReceipt `json:"delivery_receipts,omitempty"`
	Meta             ListConversationResponseMeta                `json:"meta,omitempty"`
}
