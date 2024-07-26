/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/ghostmonitor/twilio-go/client"
)

// Optional parameters for the method 'CreateRateLimit'
type CreateRateLimitParams struct {
	// Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**
	UniqueName *string `json:"UniqueName,omitempty"`
	// Description of this Rate Limit
	Description *string `json:"Description,omitempty"`
}

func (params *CreateRateLimitParams) SetUniqueName(UniqueName string) *CreateRateLimitParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateRateLimitParams) SetDescription(Description string) *CreateRateLimitParams {
	params.Description = &Description
	return params
}

// Create a new Rate Limit for a Service
func (c *ApiService) CreateRateLimit(ServiceSid string, params *CreateRateLimitParams) (*VerifyV2RateLimit, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2RateLimit{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Rate Limit.
func (c *ApiService) DeleteRateLimit(ServiceSid string, Sid string) error {
	path := "/v2/Services/{ServiceSid}/RateLimits/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Rate Limit.
func (c *ApiService) FetchRateLimit(ServiceSid string, Sid string) (*VerifyV2RateLimit, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2RateLimit{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRateLimit'
type ListRateLimitParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRateLimitParams) SetPageSize(PageSize int) *ListRateLimitParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRateLimitParams) SetLimit(Limit int) *ListRateLimitParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RateLimit records from the API. Request is executed immediately.
func (c *ApiService) PageRateLimit(
	ServiceSid string,
	params *ListRateLimitParams,
	pageToken, pageNumber string,
) (*ListRateLimitResponse, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

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

	ps := &ListRateLimitResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RateLimit records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRateLimit(ServiceSid string, params *ListRateLimitParams) ([]VerifyV2RateLimit, error) {
	response, errors := c.StreamRateLimit(ServiceSid, params)

	records := make([]VerifyV2RateLimit, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams RateLimit records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRateLimit(
	ServiceSid string,
	params *ListRateLimitParams,
) (chan VerifyV2RateLimit, chan error) {
	if params == nil {
		params = &ListRateLimitParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VerifyV2RateLimit, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRateLimit(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRateLimit(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRateLimit(
	response *ListRateLimitResponse,
	params *ListRateLimitParams,
	recordChannel chan VerifyV2RateLimit,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.RateLimits
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRateLimitResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRateLimitResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRateLimitResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRateLimitResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRateLimit'
type UpdateRateLimitParams struct {
	// Description of this Rate Limit
	Description *string `json:"Description,omitempty"`
}

func (params *UpdateRateLimitParams) SetDescription(Description string) *UpdateRateLimitParams {
	params.Description = &Description
	return params
}

// Update a specific Rate Limit.
func (c *ApiService) UpdateRateLimit(
	ServiceSid string,
	Sid string,
	params *UpdateRateLimitParams,
) (*VerifyV2RateLimit, error) {
	path := "/v2/Services/{ServiceSid}/RateLimits/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2RateLimit{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
