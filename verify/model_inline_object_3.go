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
// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// Details provided to give context about the Challenge. Shown to the end user. It must be a stringified JSON with the following structure: {\"message\": \"string\", \"fields\": [ { \"label\": \"string\", \"value\": \"string\"}]}. `message` is required. If you send the `fields` property, each field has to include `label` and `value` properties. If you had set `include_date=true` in the `push` configuration of the [service](https://www.twilio.com/docs/verify/api/service), the response will also include the challenge's date created value as an additional field called `date`
	Details string `json:"Details,omitempty"`
	// The future date in which this Challenge will expire, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	ExpirationDate time.Time `json:"ExpirationDate,omitempty"`
	// The unique SID identifier of the Factor.
	FactorSid string `json:"FactorSid"`
	// Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. `{\"ip\": \"172.168.1.234\"}`
	HiddenDetails string `json:"HiddenDetails,omitempty"`
}