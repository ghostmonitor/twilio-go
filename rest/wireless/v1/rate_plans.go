/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Wireless
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

// Optional parameters for the method 'CreateRatePlan'
type CreateRatePlanParams struct {
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
	// A descriptive string that you create to describe the resource. It does not have to be unique.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
	DataEnabled *bool `json:"DataEnabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is `1000`.
	DataLimit *int `json:"DataLimit,omitempty"`
	// The model used to meter data usage. Can be: `payg` and `quota-1`, `quota-10`, and `quota-50`. Learn more about the available [data metering models](https://www.twilio.com/docs/iot/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
	DataMetering *string `json:"DataMetering,omitempty"`
	// Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/iot/wireless/api/command-resource).
	MessagingEnabled *bool `json:"MessagingEnabled,omitempty"`
	// Deprecated.
	VoiceEnabled *bool `json:"VoiceEnabled,omitempty"`
	// Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/iot/wireless/api/rateplan-resource#national-roaming).
	NationalRoamingEnabled *bool `json:"NationalRoamingEnabled,omitempty"`
	// The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can contain: `data` and `messaging`.
	InternationalRoaming *[]string `json:"InternationalRoaming,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/iot/wireless/api/rateplan-resource#national-roaming) for more info.
	NationalRoamingDataLimit *int `json:"NationalRoamingDataLimit,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.
	InternationalRoamingDataLimit *int `json:"InternationalRoamingDataLimit,omitempty"`
}

func (params *CreateRatePlanParams) SetUniqueName(UniqueName string) *CreateRatePlanParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateRatePlanParams) SetFriendlyName(FriendlyName string) *CreateRatePlanParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateRatePlanParams) SetDataEnabled(DataEnabled bool) *CreateRatePlanParams {
	params.DataEnabled = &DataEnabled
	return params
}
func (params *CreateRatePlanParams) SetDataLimit(DataLimit int) *CreateRatePlanParams {
	params.DataLimit = &DataLimit
	return params
}
func (params *CreateRatePlanParams) SetDataMetering(DataMetering string) *CreateRatePlanParams {
	params.DataMetering = &DataMetering
	return params
}
func (params *CreateRatePlanParams) SetMessagingEnabled(MessagingEnabled bool) *CreateRatePlanParams {
	params.MessagingEnabled = &MessagingEnabled
	return params
}
func (params *CreateRatePlanParams) SetVoiceEnabled(VoiceEnabled bool) *CreateRatePlanParams {
	params.VoiceEnabled = &VoiceEnabled
	return params
}
func (params *CreateRatePlanParams) SetNationalRoamingEnabled(NationalRoamingEnabled bool) *CreateRatePlanParams {
	params.NationalRoamingEnabled = &NationalRoamingEnabled
	return params
}
func (params *CreateRatePlanParams) SetInternationalRoaming(InternationalRoaming []string) *CreateRatePlanParams {
	params.InternationalRoaming = &InternationalRoaming
	return params
}
func (params *CreateRatePlanParams) SetNationalRoamingDataLimit(NationalRoamingDataLimit int) *CreateRatePlanParams {
	params.NationalRoamingDataLimit = &NationalRoamingDataLimit
	return params
}
func (params *CreateRatePlanParams) SetInternationalRoamingDataLimit(InternationalRoamingDataLimit int) *CreateRatePlanParams {
	params.InternationalRoamingDataLimit = &InternationalRoamingDataLimit
	return params
}

//
func (c *ApiService) CreateRatePlan(params *CreateRatePlanParams) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.DataEnabled != nil {
		data.Set("DataEnabled", fmt.Sprint(*params.DataEnabled))
	}
	if params != nil && params.DataLimit != nil {
		data.Set("DataLimit", fmt.Sprint(*params.DataLimit))
	}
	if params != nil && params.DataMetering != nil {
		data.Set("DataMetering", *params.DataMetering)
	}
	if params != nil && params.MessagingEnabled != nil {
		data.Set("MessagingEnabled", fmt.Sprint(*params.MessagingEnabled))
	}
	if params != nil && params.VoiceEnabled != nil {
		data.Set("VoiceEnabled", fmt.Sprint(*params.VoiceEnabled))
	}
	if params != nil && params.NationalRoamingEnabled != nil {
		data.Set("NationalRoamingEnabled", fmt.Sprint(*params.NationalRoamingEnabled))
	}
	if params != nil && params.InternationalRoaming != nil {
		for _, item := range *params.InternationalRoaming {
			data.Add("InternationalRoaming", item)
		}
	}
	if params != nil && params.NationalRoamingDataLimit != nil {
		data.Set("NationalRoamingDataLimit", fmt.Sprint(*params.NationalRoamingDataLimit))
	}
	if params != nil && params.InternationalRoamingDataLimit != nil {
		data.Set("InternationalRoamingDataLimit", fmt.Sprint(*params.InternationalRoamingDataLimit))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteRatePlan(Sid string) error {
	path := "/v1/RatePlans/{Sid}"
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
func (c *ApiService) FetchRatePlan(Sid string) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans/{Sid}"
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

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRatePlan'
type ListRatePlanParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRatePlanParams) SetPageSize(PageSize int) *ListRatePlanParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRatePlanParams) SetLimit(Limit int) *ListRatePlanParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RatePlan records from the API. Request is executed immediately.
func (c *ApiService) PageRatePlan(
	params *ListRatePlanParams,
	pageToken, pageNumber string,
) (*ListRatePlanResponse, error) {
	path := "/v1/RatePlans"

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

	ps := &ListRatePlanResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RatePlan records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRatePlan(params *ListRatePlanParams) ([]WirelessV1RatePlan, error) {
	response, errors := c.StreamRatePlan(params)

	records := make([]WirelessV1RatePlan, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams RatePlan records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRatePlan(params *ListRatePlanParams) (chan WirelessV1RatePlan, chan error) {
	if params == nil {
		params = &ListRatePlanParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan WirelessV1RatePlan, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageRatePlan(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamRatePlan(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamRatePlan(
	response *ListRatePlanResponse,
	params *ListRatePlanParams,
	recordChannel chan WirelessV1RatePlan,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.RatePlans
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListRatePlanResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListRatePlanResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListRatePlanResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRatePlanResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRatePlan'
type UpdateRatePlanParams struct {
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
	// A descriptive string that you create to describe the resource. It does not have to be unique.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateRatePlanParams) SetUniqueName(UniqueName string) *UpdateRatePlanParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateRatePlanParams) SetFriendlyName(FriendlyName string) *UpdateRatePlanParams {
	params.FriendlyName = &FriendlyName
	return params
}

//
func (c *ApiService) UpdateRatePlan(Sid string, params *UpdateRatePlanParams) (*WirelessV1RatePlan, error) {
	path := "/v1/RatePlans/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1RatePlan{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
