/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
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

// Fetch a specific Country.
func (c *ApiService) FetchTrunkingCountry(IsoCountry string) (*PricingV2TrunkingCountryInstance, error) {
	path := "/v2/Trunking/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2TrunkingCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrunkingCountry'
type ListTrunkingCountryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrunkingCountryParams) SetPageSize(PageSize int) *ListTrunkingCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrunkingCountryParams) SetLimit(Limit int) *ListTrunkingCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrunkingCountry records from the API. Request is executed immediately.
func (c *ApiService) PageTrunkingCountry(
	params *ListTrunkingCountryParams,
	pageToken, pageNumber string,
) (*ListTrunkingCountryResponse, error) {
	path := "/v2/Trunking/Countries"

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

	ps := &ListTrunkingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrunkingCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrunkingCountry(params *ListTrunkingCountryParams) ([]PricingV2TrunkingCountry, error) {
	response, errors := c.StreamTrunkingCountry(params)

	records := make([]PricingV2TrunkingCountry, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TrunkingCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrunkingCountry(params *ListTrunkingCountryParams) (chan PricingV2TrunkingCountry, chan error) {
	if params == nil {
		params = &ListTrunkingCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan PricingV2TrunkingCountry, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTrunkingCountry(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTrunkingCountry(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTrunkingCountry(
	response *ListTrunkingCountryResponse,
	params *ListTrunkingCountryParams,
	recordChannel chan PricingV2TrunkingCountry,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Countries
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTrunkingCountryResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTrunkingCountryResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTrunkingCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrunkingCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
