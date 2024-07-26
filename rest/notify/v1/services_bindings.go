/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Notify
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

// Optional parameters for the method 'CreateBinding'
type CreateBindingParams struct {
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Up to 20 Bindings can be created for the same Identity in a given Service.
	Identity *string `json:"Identity,omitempty"`
	//
	BindingType *string `json:"BindingType,omitempty"`
	// The channel-specific address. For APNS, the device token. For FCM and GCM, the registration token. For SMS, a phone number in E.164 format. For Facebook Messenger, the Messenger ID of the user or a phone number in E.164 format.
	Address *string `json:"Address,omitempty"`
	// A tag that can be used to select the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 20 tags.
	Tag *[]string `json:"Tag,omitempty"`
	// The protocol version to use to send the notification. This defaults to the value of `default_xxxx_notification_protocol_version` for the protocol in the [Service](https://www.twilio.com/docs/notify/api/service-resource). The current version is `\\\"3\\\"` for `apn`, `fcm`, and `gcm` type Bindings. The parameter is not applicable to `sms` and `facebook-messenger` type Bindings as the data format is fixed.
	NotificationProtocolVersion *string `json:"NotificationProtocolVersion,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) resource to be used to send notifications to this Binding. If present, this overrides the Credential specified in the Service resource. Applies to only `apn`, `fcm`, and `gcm` type Bindings.
	CredentialSid *string `json:"CredentialSid,omitempty"`
	// Deprecated.
	Endpoint *string `json:"Endpoint,omitempty"`
}

func (params *CreateBindingParams) SetIdentity(Identity string) *CreateBindingParams {
	params.Identity = &Identity
	return params
}
func (params *CreateBindingParams) SetBindingType(BindingType string) *CreateBindingParams {
	params.BindingType = &BindingType
	return params
}
func (params *CreateBindingParams) SetAddress(Address string) *CreateBindingParams {
	params.Address = &Address
	return params
}
func (params *CreateBindingParams) SetTag(Tag []string) *CreateBindingParams {
	params.Tag = &Tag
	return params
}
func (params *CreateBindingParams) SetNotificationProtocolVersion(NotificationProtocolVersion string) *CreateBindingParams {
	params.NotificationProtocolVersion = &NotificationProtocolVersion
	return params
}
func (params *CreateBindingParams) SetCredentialSid(CredentialSid string) *CreateBindingParams {
	params.CredentialSid = &CredentialSid
	return params
}
func (params *CreateBindingParams) SetEndpoint(Endpoint string) *CreateBindingParams {
	params.Endpoint = &Endpoint
	return params
}

//
func (c *ApiService) CreateBinding(ServiceSid string, params *CreateBindingParams) (*NotifyV1Binding, error) {
	path := "/v1/Services/{ServiceSid}/Bindings"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.BindingType != nil {
		data.Set("BindingType", *params.BindingType)
	}
	if params != nil && params.Address != nil {
		data.Set("Address", *params.Address)
	}
	if params != nil && params.Tag != nil {
		for _, item := range *params.Tag {
			data.Add("Tag", item)
		}
	}
	if params != nil && params.NotificationProtocolVersion != nil {
		data.Set("NotificationProtocolVersion", *params.NotificationProtocolVersion)
	}
	if params != nil && params.CredentialSid != nil {
		data.Set("CredentialSid", *params.CredentialSid)
	}
	if params != nil && params.Endpoint != nil {
		data.Set("Endpoint", *params.Endpoint)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NotifyV1Binding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteBinding(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

//
func (c *ApiService) FetchBinding(ServiceSid string, Sid string) (*NotifyV1Binding, error) {
	path := "/v1/Services/{ServiceSid}/Bindings/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

	ps := &NotifyV1Binding{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBinding'
type ListBindingParams struct {
	// Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`.
	StartDate *string `json:"StartDate,omitempty"`
	// Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.
	EndDate *string `json:"EndDate,omitempty"`
	// The [User](https://www.twilio.com/docs/chat/rest/user-resource)'s `identity` value of the resources to read.
	Identity *[]string `json:"Identity,omitempty"`
	// Only list Bindings that have all of the specified Tags. The following implicit tags are available: `all`, `apn`, `fcm`, `gcm`, `sms`, `facebook-messenger`. Up to 5 tags are allowed.
	Tag *[]string `json:"Tag,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBindingParams) SetStartDate(StartDate string) *ListBindingParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListBindingParams) SetEndDate(EndDate string) *ListBindingParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListBindingParams) SetIdentity(Identity []string) *ListBindingParams {
	params.Identity = &Identity
	return params
}
func (params *ListBindingParams) SetTag(Tag []string) *ListBindingParams {
	params.Tag = &Tag
	return params
}
func (params *ListBindingParams) SetPageSize(PageSize int) *ListBindingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBindingParams) SetLimit(Limit int) *ListBindingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Binding records from the API. Request is executed immediately.
func (c *ApiService) PageBinding(
	ServiceSid string,
	params *ListBindingParams,
	pageToken, pageNumber string,
) (*ListBindingResponse, error) {
	path := "/v1/Services/{ServiceSid}/Bindings"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
	if params != nil && params.Tag != nil {
		for _, item := range *params.Tag {
			data.Add("Tag", item)
		}
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

	ps := &ListBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Binding records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBinding(ServiceSid string, params *ListBindingParams) ([]NotifyV1Binding, error) {
	response, errors := c.StreamBinding(ServiceSid, params)

	records := make([]NotifyV1Binding, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Binding records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBinding(ServiceSid string, params *ListBindingParams) (chan NotifyV1Binding, chan error) {
	if params == nil {
		params = &ListBindingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NotifyV1Binding, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageBinding(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamBinding(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamBinding(
	response *ListBindingResponse,
	params *ListBindingParams,
	recordChannel chan NotifyV1Binding,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Bindings
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListBindingResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListBindingResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListBindingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBindingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
