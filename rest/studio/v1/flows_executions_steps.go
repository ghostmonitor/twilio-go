/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Studio
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

// Retrieve a Step.
func (c *ApiService) FetchExecutionStep(
	FlowSid string,
	ExecutionSid string,
	Sid string,
) (*StudioV1ExecutionStep, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
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

	ps := &StudioV1ExecutionStep{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListExecutionStep'
type ListExecutionStepParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListExecutionStepParams) SetPageSize(PageSize int) *ListExecutionStepParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListExecutionStepParams) SetLimit(Limit int) *ListExecutionStepParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ExecutionStep records from the API. Request is executed immediately.
func (c *ApiService) PageExecutionStep(
	FlowSid string,
	ExecutionSid string,
	params *ListExecutionStepParams,
	pageToken, pageNumber string,
) (*ListExecutionStepResponse, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps"

	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)

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

	ps := &ListExecutionStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ExecutionStep records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListExecutionStep(
	FlowSid string,
	ExecutionSid string,
	params *ListExecutionStepParams,
) ([]StudioV1ExecutionStep, error) {
	response, errors := c.StreamExecutionStep(FlowSid, ExecutionSid, params)

	records := make([]StudioV1ExecutionStep, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ExecutionStep records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamExecutionStep(
	FlowSid string,
	ExecutionSid string,
	params *ListExecutionStepParams,
) (chan StudioV1ExecutionStep, chan error) {
	if params == nil {
		params = &ListExecutionStepParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan StudioV1ExecutionStep, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageExecutionStep(FlowSid, ExecutionSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamExecutionStep(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamExecutionStep(
	response *ListExecutionStepResponse,
	params *ListExecutionStepParams,
	recordChannel chan StudioV1ExecutionStep,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Steps
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListExecutionStepResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListExecutionStepResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListExecutionStepResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
