/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListMessagingConfigurationResponse struct for ListMessagingConfigurationResponse
type ListMessagingConfigurationResponse struct {
	MessagingConfigurations []VerifyV2MessagingConfiguration    `json:"messaging_configurations,omitempty"`
	Meta                    ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
}
