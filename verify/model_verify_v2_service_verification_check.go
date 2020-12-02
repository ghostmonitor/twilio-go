/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
import (
	"time"
)
// VerifyV2ServiceVerificationCheck struct for VerifyV2ServiceVerificationCheck
type VerifyV2ServiceVerificationCheck struct {
	AccountSid string `json:"account_sid,omitempty"`
	Amount string `json:"amount,omitempty"`
	Channel string `json:"channel,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Payee string `json:"payee,omitempty"`
	ServiceSid string `json:"service_sid,omitempty"`
	Sid string `json:"sid,omitempty"`
	Status string `json:"status,omitempty"`
	To string `json:"to,omitempty"`
	Valid bool `json:"valid,omitempty"`
}