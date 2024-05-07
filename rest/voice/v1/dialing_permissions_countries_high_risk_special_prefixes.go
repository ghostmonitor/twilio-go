/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Voice
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

// Optional parameters for the method 'ListDialingPermissionsHrsPrefixes'
type ListDialingPermissionsHrsPrefixesParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDialingPermissionsHrsPrefixesParams) SetPageSize(PageSize int) *ListDialingPermissionsHrsPrefixesParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDialingPermissionsHrsPrefixesParams) SetLimit(Limit int) *ListDialingPermissionsHrsPrefixesParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DialingPermissionsHrsPrefixes records from the API. Request is executed immediately.
func (c *ApiService) PageDialingPermissionsHrsPrefixes(
	IsoCode string,
	params *ListDialingPermissionsHrsPrefixesParams,
	pageToken, pageNumber string,
) (*ListDialingPermissionsHrsPrefixesResponse, error) {
	path := "/v1/DialingPermissions/Countries/{IsoCode}/HighRiskSpecialPrefixes"

	path = strings.Replace(path, "{"+"IsoCode"+"}", IsoCode, -1)

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

	ps := &ListDialingPermissionsHrsPrefixesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DialingPermissionsHrsPrefixes records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDialingPermissionsHrsPrefixes(
	IsoCode string,
	params *ListDialingPermissionsHrsPrefixesParams,
) ([]VoiceV1DialingPermissionsHrsPrefixes, error) {
	response, errors := c.StreamDialingPermissionsHrsPrefixes(IsoCode, params)

	records := make([]VoiceV1DialingPermissionsHrsPrefixes, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DialingPermissionsHrsPrefixes records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDialingPermissionsHrsPrefixes(
	IsoCode string,
	params *ListDialingPermissionsHrsPrefixesParams,
) (chan VoiceV1DialingPermissionsHrsPrefixes, chan error) {
	if params == nil {
		params = &ListDialingPermissionsHrsPrefixesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VoiceV1DialingPermissionsHrsPrefixes, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDialingPermissionsHrsPrefixes(IsoCode, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDialingPermissionsHrsPrefixes(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDialingPermissionsHrsPrefixes(
	response *ListDialingPermissionsHrsPrefixesResponse,
	params *ListDialingPermissionsHrsPrefixesParams,
	recordChannel chan VoiceV1DialingPermissionsHrsPrefixes,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Content
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDialingPermissionsHrsPrefixesResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDialingPermissionsHrsPrefixesResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDialingPermissionsHrsPrefixesResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDialingPermissionsHrsPrefixesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
