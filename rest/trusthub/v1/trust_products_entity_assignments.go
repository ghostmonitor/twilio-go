/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trusthub
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

// Optional parameters for the method 'CreateTrustProductEntityAssignment'
type CreateTrustProductEntityAssignmentParams struct {
	// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

func (params *CreateTrustProductEntityAssignmentParams) SetObjectSid(ObjectSid string) *CreateTrustProductEntityAssignmentParams {
	params.ObjectSid = &ObjectSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateTrustProductEntityAssignment(
	TrustProductSid string,
	params *CreateTrustProductEntityAssignmentParams,
) (*TrusthubV1TrustProductEntityAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.ObjectSid != nil {
		data.Set("ObjectSid", *params.ObjectSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1TrustProductEntityAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteTrustProductEntityAssignment(TrustProductSid string, Sid string) error {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchTrustProductEntityAssignment(
	TrustProductSid string,
	Sid string,
) (*TrusthubV1TrustProductEntityAssignment, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments/{Sid}"
	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)
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

	ps := &TrusthubV1TrustProductEntityAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTrustProductEntityAssignment'
type ListTrustProductEntityAssignmentParams struct {
	// A string to filter the results by (EndUserType or SupportingDocumentType) machine-name. This is useful when you want to retrieve the entity-assignment of a specific end-user or supporting document.
	ObjectType *string `json:"ObjectType,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTrustProductEntityAssignmentParams) SetObjectType(ObjectType string) *ListTrustProductEntityAssignmentParams {
	params.ObjectType = &ObjectType
	return params
}
func (params *ListTrustProductEntityAssignmentParams) SetPageSize(PageSize int) *ListTrustProductEntityAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTrustProductEntityAssignmentParams) SetLimit(Limit int) *ListTrustProductEntityAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TrustProductEntityAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageTrustProductEntityAssignment(
	TrustProductSid string,
	params *ListTrustProductEntityAssignmentParams,
	pageToken, pageNumber string,
) (*ListTrustProductEntityAssignmentResponse, error) {
	path := "/v1/TrustProducts/{TrustProductSid}/EntityAssignments"

	path = strings.Replace(path, "{"+"TrustProductSid"+"}", TrustProductSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.ObjectType != nil {
		data.Set("ObjectType", *params.ObjectType)
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

	ps := &ListTrustProductEntityAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TrustProductEntityAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTrustProductEntityAssignment(
	TrustProductSid string,
	params *ListTrustProductEntityAssignmentParams,
) ([]TrusthubV1TrustProductEntityAssignment, error) {
	response, errors := c.StreamTrustProductEntityAssignment(TrustProductSid, params)

	records := make([]TrusthubV1TrustProductEntityAssignment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TrustProductEntityAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTrustProductEntityAssignment(
	TrustProductSid string,
	params *ListTrustProductEntityAssignmentParams,
) (chan TrusthubV1TrustProductEntityAssignment, chan error) {
	if params == nil {
		params = &ListTrustProductEntityAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1TrustProductEntityAssignment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTrustProductEntityAssignment(TrustProductSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTrustProductEntityAssignment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTrustProductEntityAssignment(
	response *ListTrustProductEntityAssignmentResponse,
	params *ListTrustProductEntityAssignmentParams,
	recordChannel chan TrusthubV1TrustProductEntityAssignment,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTrustProductEntityAssignmentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTrustProductEntityAssignmentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTrustProductEntityAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTrustProductEntityAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
