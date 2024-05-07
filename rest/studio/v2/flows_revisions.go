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

// Retrieve a specific Flow revision.
func (c *ApiService) FetchFlowRevision(Sid string, Revision string) (*StudioV2FlowRevision, error) {
	path := "/v2/Flows/{Sid}/Revisions/{Revision}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)
	path = strings.Replace(path, "{"+"Revision"+"}", Revision, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2FlowRevision{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlowRevision'
type ListFlowRevisionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFlowRevisionParams) SetPageSize(PageSize int) *ListFlowRevisionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFlowRevisionParams) SetLimit(Limit int) *ListFlowRevisionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FlowRevision records from the API. Request is executed immediately.
func (c *ApiService) PageFlowRevision(
	Sid string,
	params *ListFlowRevisionParams,
	pageToken, pageNumber string,
) (*ListFlowRevisionResponse, error) {
	path := "/v2/Flows/{Sid}/Revisions"

	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

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

	ps := &ListFlowRevisionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FlowRevision records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFlowRevision(Sid string, params *ListFlowRevisionParams) ([]StudioV2FlowRevision, error) {
	response, errors := c.StreamFlowRevision(Sid, params)

	records := make([]StudioV2FlowRevision, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams FlowRevision records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFlowRevision(
	Sid string,
	params *ListFlowRevisionParams,
) (chan StudioV2FlowRevision, chan error) {
	if params == nil {
		params = &ListFlowRevisionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan StudioV2FlowRevision, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFlowRevision(Sid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFlowRevision(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFlowRevision(
	response *ListFlowRevisionResponse,
	params *ListFlowRevisionParams,
	recordChannel chan StudioV2FlowRevision,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Revisions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListFlowRevisionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFlowRevisionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFlowRevisionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowRevisionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
