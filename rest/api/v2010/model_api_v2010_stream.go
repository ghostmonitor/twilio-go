/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Stream struct for ApiV2010Stream
type ApiV2010Stream struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The name of this resource
	Name *string `json:"name,omitempty"`
	// The SID of the Stream resource.
	Sid *string `json:"sid,omitempty"`
	// The status - one of `stopped`, `in-progress`
	Status *string `json:"status,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
