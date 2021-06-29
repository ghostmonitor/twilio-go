/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
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
)

// Optional parameters for the method 'CreateMember'
type CreateMemberParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. The default value is `null`. Note that this parameter should only be used when a Member is being recreated from a backup/separate source and where a Member was previously updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
	Identity *string `json:"Identity,omitempty"`
	// The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) in the [Channel](https://www.twilio.com/docs/chat/channels) that the Member has read. This parameter should only be used when recreating a Member from a backup/separate source.
	LastConsumedMessageIndex *int `json:"LastConsumedMessageIndex,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
	LastConsumptionTimestamp *time.Time `json:"LastConsumptionTimestamp,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource).
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *CreateMemberParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateMemberParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateMemberParams) SetAttributes(Attributes string) *CreateMemberParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateMemberParams) SetDateCreated(DateCreated time.Time) *CreateMemberParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateMemberParams) SetDateUpdated(DateUpdated time.Time) *CreateMemberParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateMemberParams) SetIdentity(Identity string) *CreateMemberParams {
	params.Identity = &Identity
	return params
}
func (params *CreateMemberParams) SetLastConsumedMessageIndex(LastConsumedMessageIndex int) *CreateMemberParams {
	params.LastConsumedMessageIndex = &LastConsumedMessageIndex
	return params
}
func (params *CreateMemberParams) SetLastConsumptionTimestamp(LastConsumptionTimestamp time.Time) *CreateMemberParams {
	params.LastConsumptionTimestamp = &LastConsumptionTimestamp
	return params
}
func (params *CreateMemberParams) SetRoleSid(RoleSid string) *CreateMemberParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) CreateMember(ServiceSid string, ChannelSid string, params *CreateMemberParams) (*ChatV2ServiceChannelMember, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.LastConsumedMessageIndex != nil {
		data.Set("LastConsumedMessageIndex", fmt.Sprint(*params.LastConsumedMessageIndex))
	}
	if params != nil && params.LastConsumptionTimestamp != nil {
		data.Set("LastConsumptionTimestamp", fmt.Sprint((*params.LastConsumptionTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2ServiceChannelMember{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteMember'
type DeleteMemberParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteMemberParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteMemberParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

func (c *ApiService) DeleteMember(ServiceSid string, ChannelSid string, Sid string, params *DeleteMemberParams) error {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid}"
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

func (c *ApiService) FetchMember(ServiceSid string, ChannelSid string, Sid string) (*ChatV2ServiceChannelMember, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid}"
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

	ps := &ChatV2ServiceChannelMember{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMember'
type ListMemberParams struct {
	// The [User](https://www.twilio.com/docs/chat/rest/user-resource)'s `identity` value of the Member resources to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more details.
	Identity *[]string `json:"Identity,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListMemberParams) SetIdentity(Identity []string) *ListMemberParams {
	params.Identity = &Identity
	return params
}
func (params *ListMemberParams) SetPageSize(PageSize int) *ListMemberParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListMember(ServiceSid string, ChannelSid string, params *ListMemberParams) (*ListMemberResponse, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMemberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateMember'
type UpdateMemberParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	// A valid JSON string that contains application-specific data.
	Attributes *string `json:"Attributes,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this parameter should only be used when a Member is being recreated from a backup/separate source.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The index of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) that the Member has read within the [Channel](https://www.twilio.com/docs/chat/channels).
	LastConsumedMessageIndex *int `json:"LastConsumedMessageIndex,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp of the last [Message](https://www.twilio.com/docs/chat/rest/message-resource) read event for the Member within the [Channel](https://www.twilio.com/docs/chat/channels).
	LastConsumptionTimestamp *time.Time `json:"LastConsumptionTimestamp,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) to assign to the member. The default roles are those specified on the [Service](https://www.twilio.com/docs/chat/rest/service-resource).
	RoleSid *string `json:"RoleSid,omitempty"`
}

func (params *UpdateMemberParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateMemberParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateMemberParams) SetAttributes(Attributes string) *UpdateMemberParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateMemberParams) SetDateCreated(DateCreated time.Time) *UpdateMemberParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateMemberParams) SetDateUpdated(DateUpdated time.Time) *UpdateMemberParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *UpdateMemberParams) SetLastConsumedMessageIndex(LastConsumedMessageIndex int) *UpdateMemberParams {
	params.LastConsumedMessageIndex = &LastConsumedMessageIndex
	return params
}
func (params *UpdateMemberParams) SetLastConsumptionTimestamp(LastConsumptionTimestamp time.Time) *UpdateMemberParams {
	params.LastConsumptionTimestamp = &LastConsumptionTimestamp
	return params
}
func (params *UpdateMemberParams) SetRoleSid(RoleSid string) *UpdateMemberParams {
	params.RoleSid = &RoleSid
	return params
}

func (c *ApiService) UpdateMember(ServiceSid string, ChannelSid string, Sid string, params *UpdateMemberParams) (*ChatV2ServiceChannelMember, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{ChannelSid}/Members/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.LastConsumedMessageIndex != nil {
		data.Set("LastConsumedMessageIndex", fmt.Sprint(*params.LastConsumedMessageIndex))
	}
	if params != nil && params.LastConsumptionTimestamp != nil {
		data.Set("LastConsumptionTimestamp", fmt.Sprint((*params.LastConsumptionTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.RoleSid != nil {
		data.Set("RoleSid", *params.RoleSid)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ChatV2ServiceChannelMember{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}