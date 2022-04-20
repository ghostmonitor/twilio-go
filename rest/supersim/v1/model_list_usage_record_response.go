/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUsageRecordResponse struct for ListUsageRecordResponse
type ListUsageRecordResponse struct {
	Meta         ListEsimProfileResponseMeta `json:"meta,omitempty"`
	UsageRecords []SupersimV1UsageRecord     `json:"usage_records,omitempty"`
}
