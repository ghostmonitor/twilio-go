/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.0
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

// Optional parameters for the method 'CreateCustomerProfileEvaluation'
type CreateCustomerProfileEvaluationParams struct {
	// The unique string of a policy that is associated to the customer_profile resource.
	PolicySid *string `json:"PolicySid,omitempty"`
}

func (params *CreateCustomerProfileEvaluationParams) SetPolicySid(PolicySid string) *CreateCustomerProfileEvaluationParams {
	params.PolicySid = &PolicySid
	return params
}

// Create a new Evaluation
func (c *ApiService) CreateCustomerProfileEvaluation(CustomerProfileSid string, params *CreateCustomerProfileEvaluationParams) (*TrusthubV1CustomerProfileEvaluation, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/Evaluations"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PolicySid != nil {
		data.Set("PolicySid", *params.PolicySid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileEvaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch specific Evaluation Instance.
func (c *ApiService) FetchCustomerProfileEvaluation(CustomerProfileSid string, Sid string) (*TrusthubV1CustomerProfileEvaluation, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/Evaluations/{Sid}"
	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1CustomerProfileEvaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCustomerProfileEvaluation'
type ListCustomerProfileEvaluationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCustomerProfileEvaluationParams) SetPageSize(PageSize int) *ListCustomerProfileEvaluationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCustomerProfileEvaluationParams) SetLimit(Limit int) *ListCustomerProfileEvaluationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CustomerProfileEvaluation records from the API. Request is executed immediately.
func (c *ApiService) PageCustomerProfileEvaluation(CustomerProfileSid string, params *ListCustomerProfileEvaluationParams, pageToken, pageNumber string) (*ListCustomerProfileEvaluationResponse, error) {
	path := "/v1/CustomerProfiles/{CustomerProfileSid}/Evaluations"

	path = strings.Replace(path, "{"+"CustomerProfileSid"+"}", CustomerProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListCustomerProfileEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CustomerProfileEvaluation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCustomerProfileEvaluation(CustomerProfileSid string, params *ListCustomerProfileEvaluationParams) ([]TrusthubV1CustomerProfileEvaluation, error) {
	if params == nil {
		params = &ListCustomerProfileEvaluationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCustomerProfileEvaluation(CustomerProfileSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TrusthubV1CustomerProfileEvaluation

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCustomerProfileEvaluationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCustomerProfileEvaluationResponse)
	}

	return records, err
}

// Streams CustomerProfileEvaluation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCustomerProfileEvaluation(CustomerProfileSid string, params *ListCustomerProfileEvaluationParams) (chan TrusthubV1CustomerProfileEvaluation, error) {
	if params == nil {
		params = &ListCustomerProfileEvaluationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCustomerProfileEvaluation(CustomerProfileSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TrusthubV1CustomerProfileEvaluation, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCustomerProfileEvaluationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCustomerProfileEvaluationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCustomerProfileEvaluationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCustomerProfileEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
