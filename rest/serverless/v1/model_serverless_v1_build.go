/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ServerlessV1Build struct for ServerlessV1Build
type ServerlessV1Build struct {
	// The SID of the Account that created the Build resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The list of Asset Version resource SIDs that are included in the Build
	AssetVersions *[]map[string]interface{} `json:"asset_versions,omitempty"`
	// The ISO 8601 date and time in GMT when the Build resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Build resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A list of objects that describe the Dependencies included in the Build
	Dependencies *[]map[string]interface{} `json:"dependencies,omitempty"`
	// The list of Function Version resource SIDs that are included in the Build
	FunctionVersions *[]map[string]interface{} `json:"function_versions,omitempty"`
	Links            *map[string]interface{}   `json:"links,omitempty"`
	// The Runtime version that will be used to run the Build.
	Runtime *string `json:"runtime,omitempty"`
	// The SID of the Service that the Build resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Build resource
	Sid *string `json:"sid,omitempty"`
	// The status of the Build
	Status *string `json:"status,omitempty"`
	// The absolute URL of the Build resource
	Url *string `json:"url,omitempty"`
}
