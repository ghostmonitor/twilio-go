/*
 * Twilio - Verify
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

// VerifyV2Service struct for VerifyV2Service
type VerifyV2Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The length of the verification code
	CodeLength *int `json:"code_length,omitempty"`
	// Whether to allow sending verifications with a custom code.
	CustomCodeEnabled *bool `json:"custom_code_enabled,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	DefaultTemplateSid *string    `json:"default_template_sid,omitempty"`
	// Whether to add a security warning at the end of an SMS.
	DoNotShareWarningEnabled *bool `json:"do_not_share_warning_enabled,omitempty"`
	// Whether to ask the user to press a number before delivering the verify code in a phone call
	DtmfInputRequired *bool `json:"dtmf_input_required,omitempty"`
	// The string that you assigned to describe the verification service
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether to perform a lookup with each verification
	LookupEnabled *bool `json:"lookup_enabled,omitempty"`
	// Whether to pass PSD2 transaction parameters when starting a verification
	Psd2Enabled *bool `json:"psd2_enabled,omitempty"`
	// The service level configuration of factor push type.
	Push *interface{} `json:"push,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Whether to skip sending SMS verifications to landlines
	SkipSmsToLandlines *bool `json:"skip_sms_to_landlines,omitempty"`
	// The service level configuration of factor TOTP type.
	Totp *interface{} `json:"totp,omitempty"`
	// The name of an alternative text-to-speech service to use in phone calls
	TtsName *string `json:"tts_name,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
