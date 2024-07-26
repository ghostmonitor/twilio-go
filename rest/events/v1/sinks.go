/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Events
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

// Optional parameters for the method 'CreateSink'
type CreateSinkParams struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
	// The information required for Twilio to connect to the provided Sink encoded as JSON.
	SinkConfiguration *interface{} `json:"SinkConfiguration,omitempty"`
	//
	SinkType *string `json:"SinkType,omitempty"`
}

func (params *CreateSinkParams) SetDescription(Description string) *CreateSinkParams {
	params.Description = &Description
	return params
}
func (params *CreateSinkParams) SetSinkConfiguration(SinkConfiguration interface{}) *CreateSinkParams {
	params.SinkConfiguration = &SinkConfiguration
	return params
}
func (params *CreateSinkParams) SetSinkType(SinkType string) *CreateSinkParams {
	params.SinkType = &SinkType
	return params
}

// Create a new Sink
func (c *ApiService) CreateSink(params *CreateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkConfiguration != nil {
		v, err := json.Marshal(params.SinkConfiguration)

		if err != nil {
			return nil, err
		}

		data.Set("SinkConfiguration", string(v))
	}
	if params != nil && params.SinkType != nil {
		data.Set("SinkType", *params.SinkType)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Sink.
func (c *ApiService) DeleteSink(Sid string) error {
	path := "/v1/Sinks/{Sid}"
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

// Fetch a specific Sink.
func (c *ApiService) FetchSink(Sid string) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
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

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSink'
type ListSinkParams struct {
	// A boolean query parameter filtering the results to return sinks used/not used by a subscription.
	InUse *bool `json:"InUse,omitempty"`
	// A String query parameter filtering the results by status `initialized`, `validating`, `active` or `failed`.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSinkParams) SetInUse(InUse bool) *ListSinkParams {
	params.InUse = &InUse
	return params
}
func (params *ListSinkParams) SetStatus(Status string) *ListSinkParams {
	params.Status = &Status
	return params
}
func (params *ListSinkParams) SetPageSize(PageSize int) *ListSinkParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSinkParams) SetLimit(Limit int) *ListSinkParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Sink records from the API. Request is executed immediately.
func (c *ApiService) PageSink(params *ListSinkParams, pageToken, pageNumber string) (*ListSinkResponse, error) {
	path := "/v1/Sinks"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.InUse != nil {
		data.Set("InUse", fmt.Sprint(*params.InUse))
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
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

	ps := &ListSinkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Sink records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSink(params *ListSinkParams) ([]EventsV1Sink, error) {
	response, errors := c.StreamSink(params)

	records := make([]EventsV1Sink, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Sink records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSink(params *ListSinkParams) (chan EventsV1Sink, chan error) {
	if params == nil {
		params = &ListSinkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan EventsV1Sink, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSink(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSink(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSink(
	response *ListSinkResponse,
	params *ListSinkParams,
	recordChannel chan EventsV1Sink,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Sinks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSinkResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSinkResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSinkResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSinkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSink'
type UpdateSinkParams struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
}

func (params *UpdateSinkParams) SetDescription(Description string) *UpdateSinkParams {
	params.Description = &Description
	return params
}

// Update a specific Sink
func (c *ApiService) UpdateSink(Sid string, params *UpdateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
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

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
