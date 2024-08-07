/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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

// Optional parameters for the method 'CreateParticipant'
type CreateParticipantParams struct {
	// The phone number of the Participant.
	Identifier *string `json:"Identifier,omitempty"`
	// The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.**
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool.
	ProxyIdentifier *string `json:"ProxyIdentifier,omitempty"`
	// The SID of the Proxy Identifier to assign to the Participant.
	ProxyIdentifierSid *string `json:"ProxyIdentifierSid,omitempty"`
}

func (params *CreateParticipantParams) SetIdentifier(Identifier string) *CreateParticipantParams {
	params.Identifier = &Identifier
	return params
}
func (params *CreateParticipantParams) SetFriendlyName(FriendlyName string) *CreateParticipantParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateParticipantParams) SetProxyIdentifier(ProxyIdentifier string) *CreateParticipantParams {
	params.ProxyIdentifier = &ProxyIdentifier
	return params
}
func (params *CreateParticipantParams) SetProxyIdentifierSid(ProxyIdentifierSid string) *CreateParticipantParams {
	params.ProxyIdentifierSid = &ProxyIdentifierSid
	return params
}

// Add a new Participant to the Session
func (c *ApiService) CreateParticipant(
	ServiceSid string,
	SessionSid string,
	params *CreateParticipantParams,
) (*ProxyV1Participant, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Identifier != nil {
		data.Set("Identifier", *params.Identifier)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ProxyIdentifier != nil {
		data.Set("ProxyIdentifier", *params.ProxyIdentifier)
	}
	if params != nil && params.ProxyIdentifierSid != nil {
		data.Set("ProxyIdentifierSid", *params.ProxyIdentifierSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Participant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Participant. This is a soft-delete. The participant remains associated with the session and cannot be re-added. Participants are only permanently deleted when the [Session](https://www.twilio.com/docs/proxy/api/session) is deleted.
func (c *ApiService) DeleteParticipant(ServiceSid string, SessionSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
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

// Fetch a specific Participant.
func (c *ApiService) FetchParticipant(ServiceSid string, SessionSid string, Sid string) (*ProxyV1Participant, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
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

	ps := &ProxyV1Participant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListParticipant'
type ListParticipantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListParticipantParams) SetPageSize(PageSize int) *ListParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListParticipantParams) SetLimit(Limit int) *ListParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Participant records from the API. Request is executed immediately.
func (c *ApiService) PageParticipant(
	ServiceSid string,
	SessionSid string,
	params *ListParticipantParams,
	pageToken, pageNumber string,
) (*ListParticipantResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

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

	ps := &ListParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Participant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListParticipant(
	ServiceSid string,
	SessionSid string,
	params *ListParticipantParams,
) ([]ProxyV1Participant, error) {
	response, errors := c.StreamParticipant(ServiceSid, SessionSid, params)

	records := make([]ProxyV1Participant, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Participant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamParticipant(
	ServiceSid string,
	SessionSid string,
	params *ListParticipantParams,
) (chan ProxyV1Participant, chan error) {
	if params == nil {
		params = &ListParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ProxyV1Participant, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageParticipant(ServiceSid, SessionSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamParticipant(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamParticipant(
	response *ListParticipantResponse,
	params *ListParticipantParams,
	recordChannel chan ProxyV1Participant,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Participants
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListParticipantResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListParticipantResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListParticipantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
