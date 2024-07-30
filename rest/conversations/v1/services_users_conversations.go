/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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
	"time"

	"github.com/ghostmonitor/twilio-go/client"
)

// Delete a specific User Conversation.
func (c *ApiService) DeleteServiceUserConversation(
	ChatServiceSid string,
	UserSid string,
	ConversationSid string,
) error {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

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

// Fetch a specific User Conversation.
func (c *ApiService) FetchServiceUserConversation(
	ChatServiceSid string,
	UserSid string,
	ConversationSid string,
) (*ConversationsV1ServiceUserConversation, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceUserConversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceUserConversation'
type ListServiceUserConversationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceUserConversationParams) SetPageSize(PageSize int) *ListServiceUserConversationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceUserConversationParams) SetLimit(Limit int) *ListServiceUserConversationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceUserConversation records from the API. Request is executed immediately.
func (c *ApiService) PageServiceUserConversation(
	ChatServiceSid string,
	UserSid string,
	params *ListServiceUserConversationParams,
	pageToken, pageNumber string,
) (*ListServiceUserConversationResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

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

	ps := &ListServiceUserConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceUserConversation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceUserConversation(
	ChatServiceSid string,
	UserSid string,
	params *ListServiceUserConversationParams,
) ([]ConversationsV1ServiceUserConversation, error) {
	response, errors := c.StreamServiceUserConversation(ChatServiceSid, UserSid, params)

	records := make([]ConversationsV1ServiceUserConversation, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ServiceUserConversation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceUserConversation(
	ChatServiceSid string,
	UserSid string,
	params *ListServiceUserConversationParams,
) (chan ConversationsV1ServiceUserConversation, chan error) {
	if params == nil {
		params = &ListServiceUserConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ServiceUserConversation, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageServiceUserConversation(ChatServiceSid, UserSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamServiceUserConversation(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamServiceUserConversation(
	response *ListServiceUserConversationResponse,
	params *ListServiceUserConversationParams,
	recordChannel chan ConversationsV1ServiceUserConversation,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Conversations
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceUserConversationResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceUserConversationResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceUserConversationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceUserConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceUserConversation'
type UpdateServiceUserConversationParams struct {
	//
	NotificationLevel *string `json:"NotificationLevel,omitempty"`
	// The date of the last message read in conversation by the user, given in ISO 8601 format.
	LastReadTimestamp *time.Time `json:"LastReadTimestamp,omitempty"`
	// The index of the last Message in the Conversation that the Participant has read.
	LastReadMessageIndex *int `json:"LastReadMessageIndex,omitempty"`
}

func (params *UpdateServiceUserConversationParams) SetNotificationLevel(NotificationLevel string) *UpdateServiceUserConversationParams {
	params.NotificationLevel = &NotificationLevel
	return params
}
func (params *UpdateServiceUserConversationParams) SetLastReadTimestamp(LastReadTimestamp time.Time) *UpdateServiceUserConversationParams {
	params.LastReadTimestamp = &LastReadTimestamp
	return params
}
func (params *UpdateServiceUserConversationParams) SetLastReadMessageIndex(LastReadMessageIndex int) *UpdateServiceUserConversationParams {
	params.LastReadMessageIndex = &LastReadMessageIndex
	return params
}

// Update a specific User Conversation.
func (c *ApiService) UpdateServiceUserConversation(
	ChatServiceSid string,
	UserSid string,
	ConversationSid string,
	params *UpdateServiceUserConversationParams,
) (*ConversationsV1ServiceUserConversation, error) {
	path := "/v1/Services/{ChatServiceSid}/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.NotificationLevel != nil {
		data.Set("NotificationLevel", *params.NotificationLevel)
	}
	if params != nil && params.LastReadTimestamp != nil {
		data.Set("LastReadTimestamp", fmt.Sprint((*params.LastReadTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.LastReadMessageIndex != nil {
		data.Set("LastReadMessageIndex", fmt.Sprint(*params.LastReadMessageIndex))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceUserConversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
