/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrusthubV1CustomerProfile struct for TrusthubV1CustomerProfile
type TrusthubV1CustomerProfile struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The email address
	Email *string `json:"email,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of the Assigned Items of the Customer-Profile resource
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique string of a policy.
	PolicySid *string `json:"policy_sid,omitempty"`
	// The unique string that identifies the resource.
	Sid *string `json:"sid,omitempty"`
	// The verification status of the Customer-Profile resource
	Status *string `json:"status,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"status_callback,omitempty"`
	// The absolute URL of the Customer-Profile resource
	Url *string `json:"url,omitempty"`
	// The ISO 8601 date and time in GMT when the resource will be valid until.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
}
