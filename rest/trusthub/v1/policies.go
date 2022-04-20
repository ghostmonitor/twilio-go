/*
 * Twilio - Trusthub
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
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Fetch specific Policy Instance.
func (c *ApiService) FetchPolicies(Sid string) (*TrusthubV1Policies, error) {
	path := "/v1/Policies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1Policies{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPolicies'
type ListPoliciesParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPoliciesParams) SetPageSize(PageSize int) *ListPoliciesParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPoliciesParams) SetLimit(Limit int) *ListPoliciesParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Policies records from the API. Request is executed immediately.
func (c *ApiService) PagePolicies(params *ListPoliciesParams, pageToken, pageNumber string) (*ListPoliciesResponse, error) {
	path := "/v1/Policies"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPoliciesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Policies records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPolicies(params *ListPoliciesParams) ([]TrusthubV1Policies, error) {
	response, err := c.StreamPolicies(params)
	if err != nil {
		return nil, err
	}

	records := make([]TrusthubV1Policies, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Policies records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPolicies(params *ListPoliciesParams) (chan TrusthubV1Policies, error) {
	if params == nil {
		params = &ListPoliciesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PagePolicies(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1Policies, 1)

	go func() {
		for response != nil {
			responseRecords := response.Results
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListPoliciesResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListPoliciesResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListPoliciesResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPoliciesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
