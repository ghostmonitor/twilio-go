/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Serverless
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

// Optional parameters for the method 'CreateEnvironment'
type CreateEnvironmentParams struct {
	// A user-defined string that uniquely identifies the Environment resource. It can be a maximum of 100 characters.
	UniqueName *string `json:"UniqueName,omitempty"`
	// A URL-friendly name that represents the environment and forms part of the domain name. It can be a maximum of 16 characters.
	DomainSuffix *string `json:"DomainSuffix,omitempty"`
}

func (params *CreateEnvironmentParams) SetUniqueName(UniqueName string) *CreateEnvironmentParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateEnvironmentParams) SetDomainSuffix(DomainSuffix string) *CreateEnvironmentParams {
	params.DomainSuffix = &DomainSuffix
	return params
}

// Create a new environment.
func (c *ApiService) CreateEnvironment(
	ServiceSid string,
	params *CreateEnvironmentParams,
) (*ServerlessV1Environment, error) {
	path := "/v1/Services/{ServiceSid}/Environments"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.DomainSuffix != nil {
		data.Set("DomainSuffix", *params.DomainSuffix)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Environment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific environment.
func (c *ApiService) DeleteEnvironment(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Environments/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Retrieve a specific environment.
func (c *ApiService) FetchEnvironment(ServiceSid string, Sid string) (*ServerlessV1Environment, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Environment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEnvironment'
type ListEnvironmentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEnvironmentParams) SetPageSize(PageSize int) *ListEnvironmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEnvironmentParams) SetLimit(Limit int) *ListEnvironmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Environment records from the API. Request is executed immediately.
func (c *ApiService) PageEnvironment(
	ServiceSid string,
	params *ListEnvironmentParams,
	pageToken, pageNumber string,
) (*ListEnvironmentResponse, error) {
	path := "/v1/Services/{ServiceSid}/Environments"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListEnvironmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Environment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEnvironment(
	ServiceSid string,
	params *ListEnvironmentParams,
) ([]ServerlessV1Environment, error) {
	response, errors := c.StreamEnvironment(ServiceSid, params)

	records := make([]ServerlessV1Environment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Environment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEnvironment(
	ServiceSid string,
	params *ListEnvironmentParams,
) (chan ServerlessV1Environment, chan error) {
	if params == nil {
		params = &ListEnvironmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ServerlessV1Environment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageEnvironment(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamEnvironment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamEnvironment(
	response *ListEnvironmentResponse,
	params *ListEnvironmentParams,
	recordChannel chan ServerlessV1Environment,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Environments
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListEnvironmentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListEnvironmentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListEnvironmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEnvironmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
