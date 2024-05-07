/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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

// Optional parameters for the method 'CreateSmsCommand'
type CreateSmsCommandParams struct {
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) to send the SMS Command to.
	Sim *string `json:"Sim,omitempty"`
	// The message body of the SMS Command.
	Payload *string `json:"Payload,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_method` after we have sent the command.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
}

func (params *CreateSmsCommandParams) SetSim(Sim string) *CreateSmsCommandParams {
	params.Sim = &Sim
	return params
}
func (params *CreateSmsCommandParams) SetPayload(Payload string) *CreateSmsCommandParams {
	params.Payload = &Payload
	return params
}
func (params *CreateSmsCommandParams) SetCallbackMethod(CallbackMethod string) *CreateSmsCommandParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateSmsCommandParams) SetCallbackUrl(CallbackUrl string) *CreateSmsCommandParams {
	params.CallbackUrl = &CallbackUrl
	return params
}

// Send SMS Command to a Sim.
func (c *ApiService) CreateSmsCommand(params *CreateSmsCommandParams) (*SupersimV1SmsCommand, error) {
	path := "/v1/SmsCommands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Payload != nil {
		data.Set("Payload", *params.Payload)
	}
	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1SmsCommand{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch SMS Command instance from your account.
func (c *ApiService) FetchSmsCommand(Sid string) (*SupersimV1SmsCommand, error) {
	path := "/v1/SmsCommands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1SmsCommand{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSmsCommand'
type ListSmsCommandParams struct {
	// The SID or unique name of the Sim resource that SMS Command was sent to or from.
	Sim *string `json:"Sim,omitempty"`
	// The status of the SMS Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [SMS Command Status Values](https://www.twilio.com/docs/iot/supersim/api/smscommand-resource#status-values) for a description of each.
	Status *string `json:"Status,omitempty"`
	// The direction of the SMS Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSmsCommandParams) SetSim(Sim string) *ListSmsCommandParams {
	params.Sim = &Sim
	return params
}
func (params *ListSmsCommandParams) SetStatus(Status string) *ListSmsCommandParams {
	params.Status = &Status
	return params
}
func (params *ListSmsCommandParams) SetDirection(Direction string) *ListSmsCommandParams {
	params.Direction = &Direction
	return params
}
func (params *ListSmsCommandParams) SetPageSize(PageSize int) *ListSmsCommandParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSmsCommandParams) SetLimit(Limit int) *ListSmsCommandParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SmsCommand records from the API. Request is executed immediately.
func (c *ApiService) PageSmsCommand(
	params *ListSmsCommandParams,
	pageToken, pageNumber string,
) (*ListSmsCommandResponse, error) {
	path := "/v1/SmsCommands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
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

	ps := &ListSmsCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SmsCommand records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSmsCommand(params *ListSmsCommandParams) ([]SupersimV1SmsCommand, error) {
	response, errors := c.StreamSmsCommand(params)

	records := make([]SupersimV1SmsCommand, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SmsCommand records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSmsCommand(params *ListSmsCommandParams) (chan SupersimV1SmsCommand, chan error) {
	if params == nil {
		params = &ListSmsCommandParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1SmsCommand, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSmsCommand(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSmsCommand(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSmsCommand(
	response *ListSmsCommandResponse,
	params *ListSmsCommandParams,
	recordChannel chan SupersimV1SmsCommand,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.SmsCommands
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSmsCommandResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSmsCommandResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSmsCommandResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSmsCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
