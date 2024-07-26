/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// Fetch a Channel for an Interaction.
func (c *ApiService) FetchInteractionChannel(InteractionSid string, Sid string) (*FlexV1InteractionChannel, error) {
	path := "/v1/Interactions/{InteractionSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"InteractionSid"+"}", InteractionSid, -1)
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

	ps := &FlexV1InteractionChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInteractionChannel'
type ListInteractionChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInteractionChannelParams) SetPageSize(PageSize int) *ListInteractionChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInteractionChannelParams) SetLimit(Limit int) *ListInteractionChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InteractionChannel records from the API. Request is executed immediately.
func (c *ApiService) PageInteractionChannel(
	InteractionSid string,
	params *ListInteractionChannelParams,
	pageToken, pageNumber string,
) (*ListInteractionChannelResponse, error) {
	path := "/v1/Interactions/{InteractionSid}/Channels"

	path = strings.Replace(path, "{"+"InteractionSid"+"}", InteractionSid, -1)

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

	ps := &ListInteractionChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InteractionChannel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInteractionChannel(
	InteractionSid string,
	params *ListInteractionChannelParams,
) ([]FlexV1InteractionChannel, error) {
	response, errors := c.StreamInteractionChannel(InteractionSid, params)

	records := make([]FlexV1InteractionChannel, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InteractionChannel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInteractionChannel(
	InteractionSid string,
	params *ListInteractionChannelParams,
) (chan FlexV1InteractionChannel, chan error) {
	if params == nil {
		params = &ListInteractionChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InteractionChannel, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInteractionChannel(InteractionSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInteractionChannel(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInteractionChannel(
	response *ListInteractionChannelResponse,
	params *ListInteractionChannelParams,
	recordChannel chan FlexV1InteractionChannel,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Channels
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInteractionChannelResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInteractionChannelResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInteractionChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInteractionChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateInteractionChannel'
type UpdateInteractionChannelParams struct {
	//
	Status *string `json:"Status,omitempty"`
	// It changes the state of associated tasks. Routing status is required, When the channel status is set to `inactive`. Allowed Value for routing status is `closed`. Otherwise Optional, if not specified, all tasks will be set to `wrapping`.
	Routing *interface{} `json:"Routing,omitempty"`
}

func (params *UpdateInteractionChannelParams) SetStatus(Status string) *UpdateInteractionChannelParams {
	params.Status = &Status
	return params
}
func (params *UpdateInteractionChannelParams) SetRouting(Routing interface{}) *UpdateInteractionChannelParams {
	params.Routing = &Routing
	return params
}

// Update an existing Interaction Channel.
func (c *ApiService) UpdateInteractionChannel(
	InteractionSid string,
	Sid string,
	params *UpdateInteractionChannelParams,
) (*FlexV1InteractionChannel, error) {
	path := "/v1/Interactions/{InteractionSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"InteractionSid"+"}", InteractionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Routing != nil {
		v, err := json.Marshal(params.Routing)

		if err != nil {
			return nil, err
		}

		data.Set("Routing", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InteractionChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
