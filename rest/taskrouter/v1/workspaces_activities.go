/*
 * Twilio - Taskrouter
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

// Optional parameters for the method 'CreateActivity'
type CreateActivityParams struct {
	// Whether the Worker should be eligible to receive a Task when it occupies the Activity. A value of `true`, `1`, or `yes` specifies the Activity is available. All other values specify that it is not. The value cannot be changed after the Activity is created.
	Available *bool `json:"Available,omitempty"`
	// A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: `on-call`, `break`, and `email`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateActivityParams) SetAvailable(Available bool) *CreateActivityParams {
	params.Available = &Available
	return params
}
func (params *CreateActivityParams) SetFriendlyName(FriendlyName string) *CreateActivityParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateActivity(WorkspaceSid string, params *CreateActivityParams) (*TaskrouterV1Activity, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Activities"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Available != nil {
		data.Set("Available", fmt.Sprint(*params.Available))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Activity{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteActivity(WorkspaceSid string, Sid string) error {
	path := "/v1/Workspaces/{WorkspaceSid}/Activities/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
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

func (c *ApiService) FetchActivity(WorkspaceSid string, Sid string) (*TaskrouterV1Activity, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Activities/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Activity{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListActivity'
type ListActivityParams struct {
	// The `friendly_name` of the Activity resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether return only Activity resources that are available or unavailable. A value of `true` returns only available activities. Values of '1' or `yes` also indicate `true`. All other values represent `false` and return activities that are unavailable.
	Available *string `json:"Available,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListActivityParams) SetFriendlyName(FriendlyName string) *ListActivityParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListActivityParams) SetAvailable(Available string) *ListActivityParams {
	params.Available = &Available
	return params
}
func (params *ListActivityParams) SetPageSize(PageSize int) *ListActivityParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListActivityParams) SetLimit(Limit int) *ListActivityParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Activity records from the API. Request is executed immediately.
func (c *ApiService) PageActivity(WorkspaceSid string, params *ListActivityParams, pageToken, pageNumber string) (*ListActivityResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Activities"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Available != nil {
		data.Set("Available", *params.Available)
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

	ps := &ListActivityResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Activity records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListActivity(WorkspaceSid string, params *ListActivityParams) ([]TaskrouterV1Activity, error) {
	response, err := c.StreamActivity(WorkspaceSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]TaskrouterV1Activity, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Activity records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamActivity(WorkspaceSid string, params *ListActivityParams) (chan TaskrouterV1Activity, error) {
	if params == nil {
		params = &ListActivityParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageActivity(WorkspaceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan TaskrouterV1Activity, 1)

	go func() {
		for response != nil {
			responseRecords := response.Activities
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListActivityResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListActivityResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListActivityResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListActivityResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateActivity'
type UpdateActivityParams struct {
	// A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: `on-call`, `break`, and `email`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateActivityParams) SetFriendlyName(FriendlyName string) *UpdateActivityParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateActivity(WorkspaceSid string, Sid string, params *UpdateActivityParams) (*TaskrouterV1Activity, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Activities/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Activity{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
