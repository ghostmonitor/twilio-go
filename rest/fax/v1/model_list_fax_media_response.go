/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFaxMediaResponse struct for ListFaxMediaResponse
type ListFaxMediaResponse struct {
	Media []FaxV1FaxMedia     `json:"media,omitempty"`
	Meta  ListFaxResponseMeta `json:"meta,omitempty"`
}
