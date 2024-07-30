/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// Optional parameters for the method 'DeleteTranscription'
type DeleteTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteTranscriptionParams) SetPathAccountSid(PathAccountSid string) *DeleteTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a transcription from the account used to make the request
func (c *ApiService) DeleteTranscription(Sid string, params *DeleteTranscriptionParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchTranscription'
type FetchTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchTranscriptionParams) SetPathAccountSid(PathAccountSid string) *FetchTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a Transcription
func (c *ApiService) FetchTranscription(Sid string, params *FetchTranscriptionParams) (*ApiV2010Transcription, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

	ps := &ApiV2010Transcription{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTranscription'
type ListTranscriptionParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTranscriptionParams) SetPathAccountSid(PathAccountSid string) *ListTranscriptionParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListTranscriptionParams) SetPageSize(PageSize int) *ListTranscriptionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTranscriptionParams) SetLimit(Limit int) *ListTranscriptionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Transcription records from the API. Request is executed immediately.
func (c *ApiService) PageTranscription(
	params *ListTranscriptionParams,
	pageToken, pageNumber string,
) (*ListTranscriptionResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Transcriptions.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListTranscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Transcription records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTranscription(params *ListTranscriptionParams) ([]ApiV2010Transcription, error) {
	response, errors := c.StreamTranscription(params)

	records := make([]ApiV2010Transcription, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Transcription records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTranscription(params *ListTranscriptionParams) (chan ApiV2010Transcription, chan error) {
	if params == nil {
		params = &ListTranscriptionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010Transcription, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTranscription(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTranscription(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTranscription(
	response *ListTranscriptionResponse,
	params *ListTranscriptionParams,
	recordChannel chan ApiV2010Transcription,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Transcriptions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTranscriptionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTranscriptionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTranscriptionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTranscriptionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
