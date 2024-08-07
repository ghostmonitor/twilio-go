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

// Optional parameters for the method 'CreateSubscribedEvent'
type CreateSubscribedEventParams struct {
	// Type of event being subscribed to.
	Type *string `json:"Type,omitempty"`
	// The schema version that the Subscription should use.
	SchemaVersion *int `json:"SchemaVersion,omitempty"`
}

func (params *CreateSubscribedEventParams) SetType(Type string) *CreateSubscribedEventParams {
	params.Type = &Type
	return params
}
func (params *CreateSubscribedEventParams) SetSchemaVersion(SchemaVersion int) *CreateSubscribedEventParams {
	params.SchemaVersion = &SchemaVersion
	return params
}

// Add an event type to a Subscription.
func (c *ApiService) CreateSubscribedEvent(
	SubscriptionSid string,
	params *CreateSubscribedEventParams,
) (*EventsV1SubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.SchemaVersion != nil {
		data.Set("SchemaVersion", fmt.Sprint(*params.SchemaVersion))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SubscribedEvent{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an event type from a Subscription.
func (c *ApiService) DeleteSubscribedEvent(SubscriptionSid string, Type string) error {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

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

// Read an Event for a Subscription.
func (c *ApiService) FetchSubscribedEvent(SubscriptionSid string, Type string) (*EventsV1SubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SubscribedEvent{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSubscribedEvent'
type ListSubscribedEventParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSubscribedEventParams) SetPageSize(PageSize int) *ListSubscribedEventParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSubscribedEventParams) SetLimit(Limit int) *ListSubscribedEventParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SubscribedEvent records from the API. Request is executed immediately.
func (c *ApiService) PageSubscribedEvent(
	SubscriptionSid string,
	params *ListSubscribedEventParams,
	pageToken, pageNumber string,
) (*ListSubscribedEventResponse, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents"

	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)

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

	ps := &ListSubscribedEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SubscribedEvent records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSubscribedEvent(
	SubscriptionSid string,
	params *ListSubscribedEventParams,
) ([]EventsV1SubscribedEvent, error) {
	response, errors := c.StreamSubscribedEvent(SubscriptionSid, params)

	records := make([]EventsV1SubscribedEvent, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SubscribedEvent records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSubscribedEvent(
	SubscriptionSid string,
	params *ListSubscribedEventParams,
) (chan EventsV1SubscribedEvent, chan error) {
	if params == nil {
		params = &ListSubscribedEventParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan EventsV1SubscribedEvent, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSubscribedEvent(SubscriptionSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSubscribedEvent(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSubscribedEvent(
	response *ListSubscribedEventResponse,
	params *ListSubscribedEventParams,
	recordChannel chan EventsV1SubscribedEvent,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Types
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSubscribedEventResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSubscribedEventResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSubscribedEventResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSubscribedEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSubscribedEvent'
type UpdateSubscribedEventParams struct {
	// The schema version that the Subscription should use.
	SchemaVersion *int `json:"SchemaVersion,omitempty"`
}

func (params *UpdateSubscribedEventParams) SetSchemaVersion(SchemaVersion int) *UpdateSubscribedEventParams {
	params.SchemaVersion = &SchemaVersion
	return params
}

// Update an Event for a Subscription.
func (c *ApiService) UpdateSubscribedEvent(
	SubscriptionSid string,
	Type string,
	params *UpdateSubscribedEventParams,
) (*EventsV1SubscribedEvent, error) {
	path := "/v1/Subscriptions/{SubscriptionSid}/SubscribedEvents/{Type}"
	path = strings.Replace(path, "{"+"SubscriptionSid"+"}", SubscriptionSid, -1)
	path = strings.Replace(path, "{"+"Type"+"}", Type, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.SchemaVersion != nil {
		data.Set("SchemaVersion", fmt.Sprint(*params.SchemaVersion))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SubscribedEvent{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
