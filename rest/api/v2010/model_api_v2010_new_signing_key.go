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

// ApiV2010NewSigningKey struct for ApiV2010NewSigningKey
type ApiV2010NewSigningKey struct {
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The secret your application uses to sign Access Tokens and to authenticate to the REST API.
	Secret *string `json:"secret,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
}
