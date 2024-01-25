/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// IntelligenceV2Transcript struct for IntelligenceV2Transcript
type IntelligenceV2Transcript struct {
	// The unique SID identifier of the Account.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique SID identifier of the Service.
	ServiceSid *string `json:"service_sid,omitempty"`
	// A 34 character string that uniquely identifies this Transcript.
	Sid *string `json:"sid,omitempty"`
	// The date that this Transcript was created, given in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Transcript was updated, given in ISO 8601 format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Status      *string    `json:"status,omitempty"`
	// Media Channel describing Transcript Source and Participant Mapping
	Channel *interface{} `json:"channel,omitempty"`
	// Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
	DataLogging *bool `json:"data_logging,omitempty"`
	// The default language code of the audio.
	LanguageCode *string `json:"language_code,omitempty"`
	CustomerKey  *string `json:"customer_key,omitempty"`
	// The date that this Transcript's media was started, given in ISO 8601 format.
	MediaStartTime *time.Time `json:"media_start_time,omitempty"`
	// The duration of this Transcript's source
	Duration *int `json:"duration,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// If the transcript has been redacted, a redacted alternative of the transcript will be available.
	Redaction *bool                   `json:"redaction,omitempty"`
	Links     *map[string]interface{} `json:"links,omitempty"`
}
