/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateAuthorizationDocument'
type CreateAuthorizationDocumentParams struct {
	// A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument.
	AddressSid *string `json:"AddressSid,omitempty"`
	// Email that this AuthorizationDocument will be sent to for signing.
	Email *string `json:"Email,omitempty"`
	// The contact phone number of the person authorized to sign the Authorization Document.
	ContactPhoneNumber *string `json:"ContactPhoneNumber,omitempty"`
	// A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio's platform.
	HostedNumberOrderSids *[]string `json:"HostedNumberOrderSids,omitempty"`
	// The title of the person authorized to sign the Authorization Document for this phone number.
	ContactTitle *string `json:"ContactTitle,omitempty"`
	// Email recipients who will be informed when an Authorization Document has been sent and signed.
	CcEmails *[]string `json:"CcEmails,omitempty"`
}

func (params *CreateAuthorizationDocumentParams) SetAddressSid(AddressSid string) *CreateAuthorizationDocumentParams {
	params.AddressSid = &AddressSid
	return params
}
func (params *CreateAuthorizationDocumentParams) SetEmail(Email string) *CreateAuthorizationDocumentParams {
	params.Email = &Email
	return params
}
func (params *CreateAuthorizationDocumentParams) SetContactPhoneNumber(ContactPhoneNumber string) *CreateAuthorizationDocumentParams {
	params.ContactPhoneNumber = &ContactPhoneNumber
	return params
}
func (params *CreateAuthorizationDocumentParams) SetHostedNumberOrderSids(HostedNumberOrderSids []string) *CreateAuthorizationDocumentParams {
	params.HostedNumberOrderSids = &HostedNumberOrderSids
	return params
}
func (params *CreateAuthorizationDocumentParams) SetContactTitle(ContactTitle string) *CreateAuthorizationDocumentParams {
	params.ContactTitle = &ContactTitle
	return params
}
func (params *CreateAuthorizationDocumentParams) SetCcEmails(CcEmails []string) *CreateAuthorizationDocumentParams {
	params.CcEmails = &CcEmails
	return params
}

// Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on Twilio's platform.
func (c *ApiService) CreateAuthorizationDocument(params *CreateAuthorizationDocumentParams) (*NumbersV2AuthorizationDocument, error) {
	path := "/v2/HostedNumber/AuthorizationDocuments"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AddressSid != nil {
		data.Set("AddressSid", *params.AddressSid)
	}
	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.ContactPhoneNumber != nil {
		data.Set("ContactPhoneNumber", *params.ContactPhoneNumber)
	}
	if params != nil && params.HostedNumberOrderSids != nil {
		for _, item := range *params.HostedNumberOrderSids {
			data.Add("HostedNumberOrderSids", item)
		}
	}
	if params != nil && params.ContactTitle != nil {
		data.Set("ContactTitle", *params.ContactTitle)
	}
	if params != nil && params.CcEmails != nil {
		for _, item := range *params.CcEmails {
			data.Add("CcEmails", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2AuthorizationDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Cancel the AuthorizationDocument request.
func (c *ApiService) DeleteAuthorizationDocument(Sid string) error {
	path := "/v2/HostedNumber/AuthorizationDocuments/{Sid}"
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

// Fetch a specific AuthorizationDocument.
func (c *ApiService) FetchAuthorizationDocument(Sid string) (*NumbersV2AuthorizationDocument, error) {
	path := "/v2/HostedNumber/AuthorizationDocuments/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2AuthorizationDocument{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAuthorizationDocument'
type ListAuthorizationDocumentParams struct {
	// Email that this AuthorizationDocument will be sent to for signing.
	Email *string `json:"Email,omitempty"`
	// Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/authorization-document-resource#status-values) for more information on each of these statuses.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAuthorizationDocumentParams) SetEmail(Email string) *ListAuthorizationDocumentParams {
	params.Email = &Email
	return params
}
func (params *ListAuthorizationDocumentParams) SetStatus(Status string) *ListAuthorizationDocumentParams {
	params.Status = &Status
	return params
}
func (params *ListAuthorizationDocumentParams) SetPageSize(PageSize int) *ListAuthorizationDocumentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAuthorizationDocumentParams) SetLimit(Limit int) *ListAuthorizationDocumentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AuthorizationDocument records from the API. Request is executed immediately.
func (c *ApiService) PageAuthorizationDocument(
	params *ListAuthorizationDocumentParams,
	pageToken, pageNumber string,
) (*ListAuthorizationDocumentResponse, error) {
	path := "/v2/HostedNumber/AuthorizationDocuments"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
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

	ps := &ListAuthorizationDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AuthorizationDocument records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAuthorizationDocument(params *ListAuthorizationDocumentParams) ([]NumbersV2AuthorizationDocument, error) {
	response, errors := c.StreamAuthorizationDocument(params)

	records := make([]NumbersV2AuthorizationDocument, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AuthorizationDocument records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAuthorizationDocument(params *ListAuthorizationDocumentParams) (chan NumbersV2AuthorizationDocument, chan error) {
	if params == nil {
		params = &ListAuthorizationDocumentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2AuthorizationDocument, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAuthorizationDocument(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAuthorizationDocument(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAuthorizationDocument(
	response *ListAuthorizationDocumentResponse,
	params *ListAuthorizationDocumentParams,
	recordChannel chan NumbersV2AuthorizationDocument,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Items
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListAuthorizationDocumentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAuthorizationDocumentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAuthorizationDocumentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAuthorizationDocumentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
