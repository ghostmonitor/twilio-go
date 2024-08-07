/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Microvisor
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

// Fetch a specific Device.
func (c *ApiService) FetchDevice(Sid string) (*MicrovisorV1Device, error) {
	path := "/v1/Devices/{Sid}"
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

	ps := &MicrovisorV1Device{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDevice'
type ListDeviceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDeviceParams) SetPageSize(PageSize int) *ListDeviceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDeviceParams) SetLimit(Limit int) *ListDeviceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Device records from the API. Request is executed immediately.
func (c *ApiService) PageDevice(params *ListDeviceParams, pageToken, pageNumber string) (*ListDeviceResponse, error) {
	path := "/v1/Devices"

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

	ps := &ListDeviceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Device records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDevice(params *ListDeviceParams) ([]MicrovisorV1Device, error) {
	response, errors := c.StreamDevice(params)

	records := make([]MicrovisorV1Device, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Device records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDevice(params *ListDeviceParams) (chan MicrovisorV1Device, chan error) {
	if params == nil {
		params = &ListDeviceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MicrovisorV1Device, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDevice(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDevice(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDevice(
	response *ListDeviceResponse,
	params *ListDeviceParams,
	recordChannel chan MicrovisorV1Device,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Devices
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDeviceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDeviceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDeviceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDeviceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateDevice'
type UpdateDeviceParams struct {
	// A unique and addressable name to be assigned to this Device by the developer. It may be used in place of the Device SID.
	UniqueName *string `json:"UniqueName,omitempty"`
	// The SID or unique name of the App to be targeted to the Device.
	TargetApp *string `json:"TargetApp,omitempty"`
	// A Boolean flag specifying whether to enable application logging. Logs will be enabled or extended for 24 hours.
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty"`
	// Set to true to restart the App running on the Device.
	RestartApp *bool `json:"RestartApp,omitempty"`
}

func (params *UpdateDeviceParams) SetUniqueName(UniqueName string) *UpdateDeviceParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateDeviceParams) SetTargetApp(TargetApp string) *UpdateDeviceParams {
	params.TargetApp = &TargetApp
	return params
}
func (params *UpdateDeviceParams) SetLoggingEnabled(LoggingEnabled bool) *UpdateDeviceParams {
	params.LoggingEnabled = &LoggingEnabled
	return params
}
func (params *UpdateDeviceParams) SetRestartApp(RestartApp bool) *UpdateDeviceParams {
	params.RestartApp = &RestartApp
	return params
}

// Update a specific Device.
func (c *ApiService) UpdateDevice(Sid string, params *UpdateDeviceParams) (*MicrovisorV1Device, error) {
	path := "/v1/Devices/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.TargetApp != nil {
		data.Set("TargetApp", *params.TargetApp)
	}
	if params != nil && params.LoggingEnabled != nil {
		data.Set("LoggingEnabled", fmt.Sprint(*params.LoggingEnabled))
	}
	if params != nil && params.RestartApp != nil {
		data.Set("RestartApp", fmt.Sprint(*params.RestartApp))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MicrovisorV1Device{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
