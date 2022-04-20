/*
 * Twilio - Chat
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateMessage'
type CreateMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// The message to send to the channel. Can be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
	Body *string `json:"Body,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat's history is being recreated from a backup/separate source.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the new message's author. The default value is `system`.
	From *string `json:"From,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.
	LastUpdatedBy *string `json:"LastUpdatedBy,omitempty"`
	// The SID of the [Media](https://www.twilio.com/docs/chat/rest/media) to attach to the new Message.
	MediaSid *string `json:"MediaSid,omitempty"`
}

func (params *CreateMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateMessageParams) SetAttributes(Attributes string) *CreateMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateMessageParams) SetBody(Body string) *CreateMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageParams) SetDateCreated(DateCreated time.Time) *CreateMessageParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateMessageParams) SetDateUpdated(DateUpdated time.Time) *CreateMessageParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateMessageParams) SetFrom(From string) *CreateMessageParams {
	params.From = &From
	return params
}
func (params *CreateMessageParams) SetLastUpdatedBy(LastUpdatedBy string) *CreateMessageParams {
	params.LastUpdatedBy = &LastUpdatedBy
	return params
}
func (params *CreateMessageParams) SetMediaSid(MediaSid string) *CreateMessageParams {
	params.MediaSid = &MediaSid
	return params
}

func (c *ApiService) CreateMessage(ServiceSid string, ChannelSid string, params *CreateMessageParams) (*ChatV2Message, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.LastUpdatedBy != nil {
		data.Set("LastUpdatedBy", *params.LastUpdatedBy)
	}
	if params != nil && params.MediaSid != nil {
		data.Set("MediaSid", *params.MediaSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteMessage'
type DeleteMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

func (c *ApiService) DeleteMessage(ServiceSid string, ChannelSid string, Sid string, params *DeleteMessageParams) error {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchMessage(ServiceSid string, ChannelSid string, Sid string) (*ChatV2Message, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessage'
type ListMessageParams struct {
	// The sort order of the returned messages. Can be: `asc` (ascending) or `desc` (descending) with `asc` as the default.
	Order *string `json:"Order,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMessageParams) SetOrder(Order string) *ListMessageParams {
	params.Order = &Order
	return params
}
func (params *ListMessageParams) SetPageSize(PageSize int) *ListMessageParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMessageParams) SetLimit(Limit int) *ListMessageParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Message records from the API. Request is executed immediately.
func (c *ApiService) PageMessage(ServiceSid string, ChannelSid string, params *ListMessageParams, pageToken, pageNumber string) (*ListMessageResponse, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
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

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Message records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessage(ServiceSid string, ChannelSid string, params *ListMessageParams) ([]ChatV2Message, error) {
	response, err := c.StreamMessage(ServiceSid, ChannelSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]ChatV2Message, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Message records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessage(ServiceSid string, ChannelSid string, params *ListMessageParams) (chan ChatV2Message, error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessage(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan ChatV2Message, 1)

	go func() {
		for response != nil {
			responseRecords := response.Messages
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListMessageResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMessageResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMessageResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMessage'
type UpdateMessageParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// The message to send to the channel. Can be an empty string or `null`, which sets the value as an empty string. You can send structured data in the body by serializing it as a string.
	Body *string `json:"Body,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service. This parameter should only be used when a Chat's history is being recreated from a backup/separate source.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the message's author.
	From *string `json:"From,omitempty"`
	// The [Identity](https://www.twilio.com/docs/chat/identity) of the User who last updated the Message, if applicable.
	LastUpdatedBy *string `json:"LastUpdatedBy,omitempty"`
}

func (params *UpdateMessageParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateMessageParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateMessageParams) SetAttributes(Attributes string) *UpdateMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateMessageParams) SetBody(Body string) *UpdateMessageParams {
	params.Body = &Body
	return params
}
func (params *UpdateMessageParams) SetDateCreated(DateCreated time.Time) *UpdateMessageParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateMessageParams) SetDateUpdated(DateUpdated time.Time) *UpdateMessageParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *UpdateMessageParams) SetFrom(From string) *UpdateMessageParams {
	params.From = &From
	return params
}
func (params *UpdateMessageParams) SetLastUpdatedBy(LastUpdatedBy string) *UpdateMessageParams {
	params.LastUpdatedBy = &LastUpdatedBy
	return params
}

func (c *ApiService) UpdateMessage(ServiceSid string, ChannelSid string, Sid string, params *UpdateMessageParams) (*ChatV2Message, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.LastUpdatedBy != nil {
		data.Set("LastUpdatedBy", *params.LastUpdatedBy)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
