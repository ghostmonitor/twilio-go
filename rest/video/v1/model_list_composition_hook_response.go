/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCompositionHookResponse struct for ListCompositionHookResponse
type ListCompositionHookResponse struct {
	CompositionHooks []VideoV1CompositionHook        `json:"composition_hooks,omitempty"`
	Meta             ListCompositionHookResponseMeta `json:"meta,omitempty"`
}
