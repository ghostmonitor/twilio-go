/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSessionResponse struct for ListSessionResponse
type ListSessionResponse struct {
	Meta     ListServiceResponseMeta `json:"meta,omitempty"`
	Sessions []ProxyV1Session        `json:"sessions,omitempty"`
}
