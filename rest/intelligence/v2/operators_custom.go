/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Intelligence
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

// Optional parameters for the method 'CreateCustomOperator'
type CreateCustomOperatorParams struct {
	// A human readable description of the new Operator, up to 64 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Operator Type for this Operator. References an existing Operator Type resource.
	OperatorType *string `json:"OperatorType,omitempty"`
	// Operator configuration, following the schema defined by the Operator Type.
	Config *interface{} `json:"Config,omitempty"`
}

func (params *CreateCustomOperatorParams) SetFriendlyName(FriendlyName string) *CreateCustomOperatorParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateCustomOperatorParams) SetOperatorType(OperatorType string) *CreateCustomOperatorParams {
	params.OperatorType = &OperatorType
	return params
}
func (params *CreateCustomOperatorParams) SetConfig(Config interface{}) *CreateCustomOperatorParams {
	params.Config = &Config
	return params
}

// Create a new Custom Operator for the given Account
func (c *ApiService) CreateCustomOperator(params *CreateCustomOperatorParams) (*IntelligenceV2CustomOperator, error) {
	path := "/v2/Operators/Custom"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.OperatorType != nil {
		data.Set("OperatorType", *params.OperatorType)
	}
	if params != nil && params.Config != nil {
		v, err := json.Marshal(params.Config)

		if err != nil {
			return nil, err
		}

		data.Set("Config", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IntelligenceV2CustomOperator{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Custom Operator.
func (c *ApiService) DeleteCustomOperator(Sid string) error {
	path := "/v2/Operators/Custom/{Sid}"
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

// Fetch a specific Custom Operator.
func (c *ApiService) FetchCustomOperator(Sid string) (*IntelligenceV2CustomOperator, error) {
	path := "/v2/Operators/Custom/{Sid}"
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

	ps := &IntelligenceV2CustomOperator{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCustomOperator'
type ListCustomOperatorParams struct {
	// Returns Custom Operators with the provided availability type. Possible values: internal, beta, public, retired.
	Availability *string `json:"Availability,omitempty"`
	// Returns Custom Operators that support the provided language code.
	LanguageCode *string `json:"LanguageCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCustomOperatorParams) SetAvailability(Availability string) *ListCustomOperatorParams {
	params.Availability = &Availability
	return params
}
func (params *ListCustomOperatorParams) SetLanguageCode(LanguageCode string) *ListCustomOperatorParams {
	params.LanguageCode = &LanguageCode
	return params
}
func (params *ListCustomOperatorParams) SetPageSize(PageSize int) *ListCustomOperatorParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCustomOperatorParams) SetLimit(Limit int) *ListCustomOperatorParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CustomOperator records from the API. Request is executed immediately.
func (c *ApiService) PageCustomOperator(
	params *ListCustomOperatorParams,
	pageToken, pageNumber string,
) (*ListCustomOperatorResponse, error) {
	path := "/v2/Operators/Custom"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Availability != nil {
		data.Set("Availability", *params.Availability)
	}
	if params != nil && params.LanguageCode != nil {
		data.Set("LanguageCode", *params.LanguageCode)
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

	ps := &ListCustomOperatorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CustomOperator records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCustomOperator(params *ListCustomOperatorParams) ([]IntelligenceV2CustomOperator, error) {
	response, errors := c.StreamCustomOperator(params)

	records := make([]IntelligenceV2CustomOperator, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CustomOperator records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCustomOperator(params *ListCustomOperatorParams) (chan IntelligenceV2CustomOperator, chan error) {
	if params == nil {
		params = &ListCustomOperatorParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IntelligenceV2CustomOperator, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCustomOperator(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCustomOperator(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCustomOperator(
	response *ListCustomOperatorResponse,
	params *ListCustomOperatorParams,
	recordChannel chan IntelligenceV2CustomOperator,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Operators
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCustomOperatorResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCustomOperatorResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCustomOperatorResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomOperatorResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCustomOperator'
type UpdateCustomOperatorParams struct {
	// The If-Match HTTP request header
	IfMatch *string `json:"If-Match,omitempty"`
	// A human-readable name of this resource, up to 64 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Operator configuration, following the schema defined by the Operator Type.
	Config *interface{} `json:"Config,omitempty"`
}

func (params *UpdateCustomOperatorParams) SetIfMatch(IfMatch string) *UpdateCustomOperatorParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateCustomOperatorParams) SetFriendlyName(FriendlyName string) *UpdateCustomOperatorParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateCustomOperatorParams) SetConfig(Config interface{}) *UpdateCustomOperatorParams {
	params.Config = &Config
	return params
}

// Update a specific Custom Operator.
func (c *ApiService) UpdateCustomOperator(
	Sid string,
	params *UpdateCustomOperatorParams,
) (*IntelligenceV2CustomOperator, error) {
	path := "/v2/Operators/Custom/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Config != nil {
		v, err := json.Marshal(params.Config)

		if err != nil {
			return nil, err
		}

		data.Set("Config", string(v))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IntelligenceV2CustomOperator{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
