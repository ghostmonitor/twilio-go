/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
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

// Optional parameters for the method 'CreateDocument'
type CreateDocumentParams struct {
	// A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (the Sync Document's time-to-live).
	Ttl *int `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the Sync Document
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateDocumentParams) SetData(Data map[string]interface{}) *CreateDocumentParams {
	params.Data = &Data
	return params
}
func (params *CreateDocumentParams) SetTtl(Ttl int) *CreateDocumentParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateDocumentParams) SetUniqueName(UniqueName string) *CreateDocumentParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateDocument(ServiceSid string, params *CreateDocumentParams) (*SyncV1Document, error) {
	path := "/v1/Services/{ServiceSid}/Documents"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1Document{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteDocument(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Documents/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchDocument(ServiceSid string, Sid string) (*SyncV1Document, error) {
	path := "/v1/Services/{ServiceSid}/Documents/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1Document{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDocument'
type ListDocumentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDocumentParams) SetPageSize(PageSize int) *ListDocumentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDocumentParams) SetLimit(Limit int) *ListDocumentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Document records from the API. Request is executed immediately.
func (c *ApiService) PageDocument(ServiceSid string, params *ListDocumentParams, pageToken, pageNumber string) (*ListDocumentResponse, error) {
	path := "/v1/Services/{ServiceSid}/Documents"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Document records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDocument(ServiceSid string, params *ListDocumentParams) ([]SyncV1Document, error) {
	if params == nil {
		params = &ListDocumentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDocument(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1Document

	for response != nil {
		records = append(records, response.Documents...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListDocumentResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListDocumentResponse)
	}

	return records, err
}

// Streams Document records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDocument(ServiceSid string, params *ListDocumentParams) (chan SyncV1Document, error) {
	if params == nil {
		params = &ListDocumentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageDocument(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1Document, 1)

	go func() {
		for response != nil {
			for item := range response.Documents {
				channel <- response.Documents[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListDocumentResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDocumentResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDocumentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateDocument'
type UpdateDocumentParams struct {
	// The If-Match HTTP request header
	IfMatch *string `json:"If-Match,omitempty"`
	// A JSON string that represents an arbitrary, schema-less object that the Sync Document stores. Can be up to 16 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync Document expires and is deleted (time-to-live).
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateDocumentParams) SetIfMatch(IfMatch string) *UpdateDocumentParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateDocumentParams) SetData(Data map[string]interface{}) *UpdateDocumentParams {
	params.Data = &Data
	return params
}
func (params *UpdateDocumentParams) SetTtl(Ttl int) *UpdateDocumentParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) UpdateDocument(ServiceSid string, Sid string, params *UpdateDocumentParams) (*SyncV1Document, error) {
	path := "/v1/Services/{ServiceSid}/Documents/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1Document{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
