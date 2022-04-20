/*
 * Twilio - Verify
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
	"net/url"
	"strings"
)

// Fetch the forms for a specific Form Type.
func (c *ApiService) FetchForm(FormType string) (*VerifyV2Form, error) {
	path := "/v2/Forms/{FormType}"
	path = strings.Replace(path, "{"+"FormType"+"}", FormType, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Form{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
