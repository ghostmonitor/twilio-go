/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Content
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

	"github.com/ghostmonitor/twilio-go/client"
)

// Optional parameters for the method 'ListContentAndApprovals'
type ListContentAndApprovalsParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListContentAndApprovalsParams) SetPageSize(PageSize int) *ListContentAndApprovalsParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListContentAndApprovalsParams) SetLimit(Limit int) *ListContentAndApprovalsParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ContentAndApprovals records from the API. Request is executed immediately.
func (c *ApiService) PageContentAndApprovals(
	params *ListContentAndApprovalsParams,
	pageToken, pageNumber string,
) (*ListContentAndApprovalsResponse, error) {
	path := "/v1/ContentAndApprovals"

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

	ps := &ListContentAndApprovalsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ContentAndApprovals records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListContentAndApprovals(params *ListContentAndApprovalsParams) ([]ContentV1ContentAndApprovals, error) {
	response, errors := c.StreamContentAndApprovals(params)

	records := make([]ContentV1ContentAndApprovals, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ContentAndApprovals records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamContentAndApprovals(params *ListContentAndApprovalsParams) (chan ContentV1ContentAndApprovals, chan error) {
	if params == nil {
		params = &ListContentAndApprovalsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ContentV1ContentAndApprovals, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageContentAndApprovals(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamContentAndApprovals(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamContentAndApprovals(
	response *ListContentAndApprovalsResponse,
	params *ListContentAndApprovalsParams,
	recordChannel chan ContentV1ContentAndApprovals,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Contents
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListContentAndApprovalsResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListContentAndApprovalsResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListContentAndApprovalsResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListContentAndApprovalsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
