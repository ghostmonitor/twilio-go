/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PricingV2VoiceVoiceNumber struct for PricingV2VoiceVoiceNumber
type PricingV2VoiceVoiceNumber struct {
	Country            *string                   `json:"Country,omitempty"`
	DestinationNumber  *string                   `json:"DestinationNumber,omitempty"`
	InboundCallPrice   *InboundCallPrice         `json:"InboundCallPrice,omitempty"`
	IsoCountry         *string                   `json:"IsoCountry,omitempty"`
	OriginationNumber  *string                   `json:"OriginationNumber,omitempty"`
	OutboundCallPrices *[]map[string]interface{} `json:"OutboundCallPrices,omitempty"`
	PriceUnit          *string                   `json:"PriceUnit,omitempty"`
	Url                *string                   `json:"Url,omitempty"`
}