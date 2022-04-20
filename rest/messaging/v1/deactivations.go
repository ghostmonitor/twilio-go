/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Optional parameters for the method 'FetchDeactivation'
type FetchDeactivationParams struct {
	// The request will return a list of all United States Phone Numbers that were deactivated on the day specified by this parameter. This date should be specified in YYYY-MM-DD format.
	Date *string `json:"Date,omitempty"`
}

func (params *FetchDeactivationParams) SetDate(Date string) *FetchDeactivationParams {
	params.Date = &Date
	return params
}

// Fetch a list of all United States numbers that have been deactivated on a specific date.
func (c *ApiService) FetchDeactivation(params *FetchDeactivationParams) (*MessagingV1Deactivation, error) {
	path := "/v1/Deactivations"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Date != nil {
		data.Set("Date", fmt.Sprint(*params.Date))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1Deactivation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
