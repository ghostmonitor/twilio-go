/*
 * Twilio - Wireless
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

// Delete a Sim resource on your Account.
func (c *ApiService) DeleteSim(Sid string) error {
	path := "/v1/Sims/{Sid}"
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

// Fetch a Sim resource on your Account.
func (c *ApiService) FetchSim(Sid string) (*WirelessV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSim'
type ListSimParams struct {
	// Only return Sim resources with this status.
	Status *string `json:"Status,omitempty"`
	// Only return Sim resources with this ICCID. This will return a list with a maximum size of 1.
	Iccid *string `json:"Iccid,omitempty"`
	// The SID or unique name of a [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource). Only return Sim resources assigned to this RatePlan resource.
	RatePlan *string `json:"RatePlan,omitempty"`
	// Deprecated.
	EId *string `json:"EId,omitempty"`
	// Only return Sim resources with this registration code. This will return a list with a maximum size of 1.
	SimRegistrationCode *string `json:"SimRegistrationCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSimParams) SetStatus(Status string) *ListSimParams {
	params.Status = &Status
	return params
}
func (params *ListSimParams) SetIccid(Iccid string) *ListSimParams {
	params.Iccid = &Iccid
	return params
}
func (params *ListSimParams) SetRatePlan(RatePlan string) *ListSimParams {
	params.RatePlan = &RatePlan
	return params
}
func (params *ListSimParams) SetEId(EId string) *ListSimParams {
	params.EId = &EId
	return params
}
func (params *ListSimParams) SetSimRegistrationCode(SimRegistrationCode string) *ListSimParams {
	params.SimRegistrationCode = &SimRegistrationCode
	return params
}
func (params *ListSimParams) SetPageSize(PageSize int) *ListSimParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSimParams) SetLimit(Limit int) *ListSimParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Sim records from the API. Request is executed immediately.
func (c *ApiService) PageSim(params *ListSimParams, pageToken, pageNumber string) (*ListSimResponse, error) {
	path := "/v1/Sims"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Iccid != nil {
		data.Set("Iccid", *params.Iccid)
	}
	if params != nil && params.RatePlan != nil {
		data.Set("RatePlan", *params.RatePlan)
	}
	if params != nil && params.EId != nil {
		data.Set("EId", *params.EId)
	}
	if params != nil && params.SimRegistrationCode != nil {
		data.Set("SimRegistrationCode", *params.SimRegistrationCode)
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

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Sim records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSim(params *ListSimParams) ([]WirelessV1Sim, error) {
	response, err := c.StreamSim(params)
	if err != nil {
		return nil, err
	}

	records := make([]WirelessV1Sim, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Sim records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSim(params *ListSimParams) (chan WirelessV1Sim, error) {
	if params == nil {
		params = &ListSimParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSim(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan WirelessV1Sim, 1)

	go func() {
		for response != nil {
			responseRecords := response.Sims
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListSimResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSimResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSimResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSimResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSim'
type UpdateSimParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource's status is `new`. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts).
	AccountSid *string `json:"AccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `POST` or `GET`. The default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_url` when the SIM has finished updating. When the SIM transitions from `new` to `ready` or from any status to `deactivated`, we call this URL when the status changes to an intermediate status (`ready` or `deactivated`) and again when the status changes to its final status (`active` or `canceled`).
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The HTTP method we should use to call `commands_callback_url`. Can be: `POST` or `GET`. The default is `POST`.
	CommandsCallbackMethod *string `json:"CommandsCallbackMethod,omitempty"`
	// The URL we should call using the `commands_callback_method` when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored.
	CommandsCallbackUrl *string `json:"CommandsCallbackUrl,omitempty"`
	// A descriptive string that you create to describe the Sim resource. It does not need to be unique.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned.
	RatePlan *string `json:"RatePlan,omitempty"`
	// Initiate a connectivity reset on the SIM. Set to `resetting` to initiate a connectivity reset on the SIM. No other value is valid.
	ResetStatus *string `json:"ResetStatus,omitempty"`
	// The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. Default is `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL we should call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. Default is `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we should call using the `sms_method` when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource).
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The new status of the Sim resource. Can be: `ready`, `active`, `suspended`, or `deactivated`.
	Status *string `json:"Status,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
	// Deprecated.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// Deprecated.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// Deprecated.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// Deprecated.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *UpdateSimParams) SetAccountSid(AccountSid string) *UpdateSimParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *UpdateSimParams) SetCallbackMethod(CallbackMethod string) *UpdateSimParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *UpdateSimParams) SetCallbackUrl(CallbackUrl string) *UpdateSimParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *UpdateSimParams) SetCommandsCallbackMethod(CommandsCallbackMethod string) *UpdateSimParams {
	params.CommandsCallbackMethod = &CommandsCallbackMethod
	return params
}
func (params *UpdateSimParams) SetCommandsCallbackUrl(CommandsCallbackUrl string) *UpdateSimParams {
	params.CommandsCallbackUrl = &CommandsCallbackUrl
	return params
}
func (params *UpdateSimParams) SetFriendlyName(FriendlyName string) *UpdateSimParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateSimParams) SetRatePlan(RatePlan string) *UpdateSimParams {
	params.RatePlan = &RatePlan
	return params
}
func (params *UpdateSimParams) SetResetStatus(ResetStatus string) *UpdateSimParams {
	params.ResetStatus = &ResetStatus
	return params
}
func (params *UpdateSimParams) SetSmsFallbackMethod(SmsFallbackMethod string) *UpdateSimParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}
func (params *UpdateSimParams) SetSmsFallbackUrl(SmsFallbackUrl string) *UpdateSimParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *UpdateSimParams) SetSmsMethod(SmsMethod string) *UpdateSimParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *UpdateSimParams) SetSmsUrl(SmsUrl string) *UpdateSimParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *UpdateSimParams) SetStatus(Status string) *UpdateSimParams {
	params.Status = &Status
	return params
}
func (params *UpdateSimParams) SetUniqueName(UniqueName string) *UpdateSimParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateSimParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *UpdateSimParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *UpdateSimParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *UpdateSimParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *UpdateSimParams) SetVoiceMethod(VoiceMethod string) *UpdateSimParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *UpdateSimParams) SetVoiceUrl(VoiceUrl string) *UpdateSimParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

// Updates the given properties of a Sim resource on your Account.
func (c *ApiService) UpdateSim(Sid string, params *UpdateSimParams) (*WirelessV1Sim, error) {
	path := "/v1/Sims/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.CommandsCallbackMethod != nil {
		data.Set("CommandsCallbackMethod", *params.CommandsCallbackMethod)
	}
	if params != nil && params.CommandsCallbackUrl != nil {
		data.Set("CommandsCallbackUrl", *params.CommandsCallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RatePlan != nil {
		data.Set("RatePlan", *params.RatePlan)
	}
	if params != nil && params.ResetStatus != nil {
		data.Set("ResetStatus", *params.ResetStatus)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Sim{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
