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

// Optional parameters for the method 'CreateFlow'
type CreateFlowParams struct {
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
	// JSON representation of flow definition.
	Definition *interface{} `json:"Definition,omitempty"`
	// Description of change made in the revision.
	CommitMessage *string `json:"CommitMessage,omitempty"`
}

func (params *CreateFlowParams) SetFriendlyName(FriendlyName string) *CreateFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateFlowParams) SetStatus(Status string) *CreateFlowParams {
	params.Status = &Status
	return params
}
func (params *CreateFlowParams) SetDefinition(Definition interface{}) *CreateFlowParams {
	params.Definition = &Definition
	return params
}
func (params *CreateFlowParams) SetCommitMessage(CommitMessage string) *CreateFlowParams {
	params.CommitMessage = &CommitMessage
	return params
}

// Create a Flow.
func (c *ApiService) CreateFlow(params *CreateFlowParams) (*StudioV2Flow, error) {
	path := "/v2/Flows"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Definition != nil {
		v, err := json.Marshal(params.Definition)

		if err != nil {
			return nil, err
		}

		data.Set("Definition", string(v))
	}
	if params != nil && params.CommitMessage != nil {
		data.Set("CommitMessage", *params.CommitMessage)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Flow.
func (c *ApiService) DeleteFlow(Sid string) error {
	path := "/v2/Flows/{Sid}"
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

// Retrieve a specific Flow.
func (c *ApiService) FetchFlow(Sid string) (*StudioV2Flow, error) {
	path := "/v2/Flows/{Sid}"
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

	ps := &StudioV2Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFlowParams) SetLimit(Limit int) *ListFlowParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Flow records from the API. Request is executed immediately.
func (c *ApiService) PageFlow(params *ListFlowParams, pageToken, pageNumber string) (*ListFlowResponse, error) {
	path := "/v2/Flows"

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

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Flow records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlow(params *ListFlowParams) ([]StudioV2Flow, error) {
	response, errors := c.StreamFlow(params)

	records := make([]StudioV2Flow, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Flow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlow(params *ListFlowParams) (chan StudioV2Flow, chan error) {
	if params == nil {
		params = &ListFlowParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan StudioV2Flow, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFlow(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFlow(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFlow(
	response *ListFlowResponse,
	params *ListFlowParams,
	recordChannel chan StudioV2Flow,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Flows
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListFlowResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFlowResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFlowResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateFlow'
type UpdateFlowParams struct {
	//
	Status *string `json:"Status,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// JSON representation of flow definition.
	Definition *interface{} `json:"Definition,omitempty"`
	// Description of change made in the revision.
	CommitMessage *string `json:"CommitMessage,omitempty"`
}

func (params *UpdateFlowParams) SetStatus(Status string) *UpdateFlowParams {
	params.Status = &Status
	return params
}
func (params *UpdateFlowParams) SetFriendlyName(FriendlyName string) *UpdateFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFlowParams) SetDefinition(Definition interface{}) *UpdateFlowParams {
	params.Definition = &Definition
	return params
}
func (params *UpdateFlowParams) SetCommitMessage(CommitMessage string) *UpdateFlowParams {
	params.CommitMessage = &CommitMessage
	return params
}

// Update a Flow.
func (c *ApiService) UpdateFlow(Sid string, params *UpdateFlowParams) (*StudioV2Flow, error) {
	path := "/v2/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Definition != nil {
		v, err := json.Marshal(params.Definition)

		if err != nil {
			return nil, err
		}

		data.Set("Definition", string(v))
	}
	if params != nil && params.CommitMessage != nil {
		data.Set("CommitMessage", *params.CommitMessage)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
