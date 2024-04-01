/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// FlexV1PluginArchive struct for FlexV1PluginArchive
type FlexV1PluginArchive struct {
	// The unique string that we created to identify the Flex Plugin resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Plugin resource and owns this resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The name that uniquely identifies this Flex Plugin resource.
	UniqueName *string `json:"unique_name,omitempty"`
	// The friendly name this Flex Plugin resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A descriptive string that you create to describe the plugin resource. It can be up to 500 characters long
	Description *string `json:"description,omitempty"`
	// Whether the Flex Plugin is archived. The default value is false.
	Archived *bool `json:"archived,omitempty"`
	// The date and time in GMT when the Flex Plugin was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Flex Plugin was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Flex Plugin resource.
	Url *string `json:"url,omitempty"`
}