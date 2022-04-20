/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateBundle'
type CreateBundleParams struct {
	// The email address that will receive updates when the Bundle resource changes status.
	Email *string `json:"Email,omitempty"`
	// The [type of End User](https://www.twilio.com/docs/phone-numbers/regulatory/api/end-user-types) of the Bundle resource.
	EndUserType *string `json:"EndUserType,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Bundle's phone number country ownership request.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The type of phone number of the Bundle's ownership request. Can be `local`, `mobile`, `national`, or `toll free`.
	NumberType *string `json:"NumberType,omitempty"`
	// The unique string of a regulation that is associated to the Bundle resource.
	RegulationSid *string `json:"RegulationSid,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

func (params *CreateBundleParams) SetEmail(Email string) *CreateBundleParams {
	params.Email = &Email
	return params
}
func (params *CreateBundleParams) SetEndUserType(EndUserType string) *CreateBundleParams {
	params.EndUserType = &EndUserType
	return params
}
func (params *CreateBundleParams) SetFriendlyName(FriendlyName string) *CreateBundleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateBundleParams) SetIsoCountry(IsoCountry string) *CreateBundleParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *CreateBundleParams) SetNumberType(NumberType string) *CreateBundleParams {
	params.NumberType = &NumberType
	return params
}
func (params *CreateBundleParams) SetRegulationSid(RegulationSid string) *CreateBundleParams {
	params.RegulationSid = &RegulationSid
	return params
}
func (params *CreateBundleParams) SetStatusCallback(StatusCallback string) *CreateBundleParams {
	params.StatusCallback = &StatusCallback
	return params
}

// Create a new Bundle.
func (c *ApiService) CreateBundle(params *CreateBundleParams) (*NumbersV2Bundle, error) {
	path := "/v2/RegulatoryCompliance/Bundles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.EndUserType != nil {
		data.Set("EndUserType", *params.EndUserType)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.NumberType != nil {
		data.Set("NumberType", *params.NumberType)
	}
	if params != nil && params.RegulationSid != nil {
		data.Set("RegulationSid", *params.RegulationSid)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Bundle{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Bundle.
func (c *ApiService) DeleteBundle(Sid string) error {
	path := "/v2/RegulatoryCompliance/Bundles/{Sid}"
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

// Fetch a specific Bundle instance.
func (c *ApiService) FetchBundle(Sid string) (*NumbersV2Bundle, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Bundle{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListBundle'
type ListBundleParams struct {
	// The verification status of the Bundle resource. Please refer to [Bundle Statuses](https://www.twilio.com/docs/phone-numbers/regulatory/api/bundles#bundle-statuses) for more details.
	Status *string `json:"Status,omitempty"`
	// The string that you assigned to describe the resource. The column can contain 255 variable characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The unique string of a [Regulation resource](https://www.twilio.com/docs/phone-numbers/regulatory/api/regulations) that is associated to the Bundle resource.
	RegulationSid *string `json:"RegulationSid,omitempty"`
	// The 2-digit [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Bundle's phone number country ownership request.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The type of phone number of the Bundle's ownership request. Can be `local`, `mobile`, `national`, or `tollfree`.
	NumberType *string `json:"NumberType,omitempty"`
	// Indicates that the Bundle is a valid Bundle until a specified expiration date.
	HasValidUntilDate *bool `json:"HasValidUntilDate,omitempty"`
	// Can be `valid-until` or `date-updated`. Defaults to `date-created`.
	SortBy *string `json:"SortBy,omitempty"`
	// Default is `DESC`. Can be `ASC` or `DESC`.
	SortDirection *string `json:"SortDirection,omitempty"`
	// Date to filter Bundles having their `valid_until_date` before or after the specified date. Can be `ValidUntilDate>=` or `ValidUntilDate<=`. Both can be used in conjunction as well. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) is the acceptable date format.
	ValidUntilDate *time.Time `json:"ValidUntilDate,omitempty"`
	// Date to filter Bundles having their `valid_until_date` before or after the specified date. Can be `ValidUntilDate>=` or `ValidUntilDate<=`. Both can be used in conjunction as well. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) is the acceptable date format.
	ValidUntilDateBefore *time.Time `json:"ValidUntilDate&lt;,omitempty"`
	// Date to filter Bundles having their `valid_until_date` before or after the specified date. Can be `ValidUntilDate>=` or `ValidUntilDate<=`. Both can be used in conjunction as well. [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) is the acceptable date format.
	ValidUntilDateAfter *time.Time `json:"ValidUntilDate&gt;,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListBundleParams) SetStatus(Status string) *ListBundleParams {
	params.Status = &Status
	return params
}
func (params *ListBundleParams) SetFriendlyName(FriendlyName string) *ListBundleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListBundleParams) SetRegulationSid(RegulationSid string) *ListBundleParams {
	params.RegulationSid = &RegulationSid
	return params
}
func (params *ListBundleParams) SetIsoCountry(IsoCountry string) *ListBundleParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListBundleParams) SetNumberType(NumberType string) *ListBundleParams {
	params.NumberType = &NumberType
	return params
}
func (params *ListBundleParams) SetHasValidUntilDate(HasValidUntilDate bool) *ListBundleParams {
	params.HasValidUntilDate = &HasValidUntilDate
	return params
}
func (params *ListBundleParams) SetSortBy(SortBy string) *ListBundleParams {
	params.SortBy = &SortBy
	return params
}
func (params *ListBundleParams) SetSortDirection(SortDirection string) *ListBundleParams {
	params.SortDirection = &SortDirection
	return params
}
func (params *ListBundleParams) SetValidUntilDate(ValidUntilDate time.Time) *ListBundleParams {
	params.ValidUntilDate = &ValidUntilDate
	return params
}
func (params *ListBundleParams) SetValidUntilDateBefore(ValidUntilDateBefore time.Time) *ListBundleParams {
	params.ValidUntilDateBefore = &ValidUntilDateBefore
	return params
}
func (params *ListBundleParams) SetValidUntilDateAfter(ValidUntilDateAfter time.Time) *ListBundleParams {
	params.ValidUntilDateAfter = &ValidUntilDateAfter
	return params
}
func (params *ListBundleParams) SetPageSize(PageSize int) *ListBundleParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListBundleParams) SetLimit(Limit int) *ListBundleParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Bundle records from the API. Request is executed immediately.
func (c *ApiService) PageBundle(params *ListBundleParams, pageToken, pageNumber string) (*ListBundleResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RegulationSid != nil {
		data.Set("RegulationSid", *params.RegulationSid)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.NumberType != nil {
		data.Set("NumberType", *params.NumberType)
	}
	if params != nil && params.HasValidUntilDate != nil {
		data.Set("HasValidUntilDate", fmt.Sprint(*params.HasValidUntilDate))
	}
	if params != nil && params.SortBy != nil {
		data.Set("SortBy", *params.SortBy)
	}
	if params != nil && params.SortDirection != nil {
		data.Set("SortDirection", *params.SortDirection)
	}
	if params != nil && params.ValidUntilDate != nil {
		data.Set("ValidUntilDate", fmt.Sprint((*params.ValidUntilDate).Format(time.RFC3339)))
	}
	if params != nil && params.ValidUntilDateBefore != nil {
		data.Set("ValidUntilDate<", fmt.Sprint((*params.ValidUntilDateBefore).Format(time.RFC3339)))
	}
	if params != nil && params.ValidUntilDateAfter != nil {
		data.Set("ValidUntilDate>", fmt.Sprint((*params.ValidUntilDateAfter).Format(time.RFC3339)))
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

	ps := &ListBundleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Bundle records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListBundle(params *ListBundleParams) ([]NumbersV2Bundle, error) {
	response, err := c.StreamBundle(params)
	if err != nil {
		return nil, err
	}

	records := make([]NumbersV2Bundle, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams Bundle records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamBundle(params *ListBundleParams) (chan NumbersV2Bundle, error) {
	if params == nil {
		params = &ListBundleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageBundle(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan NumbersV2Bundle, 1)

	go func() {
		for response != nil {
			responseRecords := response.Results
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListBundleResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListBundleResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListBundleResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListBundleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateBundle'
type UpdateBundleParams struct {
	// The email address that will receive updates when the Bundle resource changes status.
	Email *string `json:"Email,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The verification status of the Bundle resource.
	Status *string `json:"Status,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback *string `json:"StatusCallback,omitempty"`
}

func (params *UpdateBundleParams) SetEmail(Email string) *UpdateBundleParams {
	params.Email = &Email
	return params
}
func (params *UpdateBundleParams) SetFriendlyName(FriendlyName string) *UpdateBundleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateBundleParams) SetStatus(Status string) *UpdateBundleParams {
	params.Status = &Status
	return params
}
func (params *UpdateBundleParams) SetStatusCallback(StatusCallback string) *UpdateBundleParams {
	params.StatusCallback = &StatusCallback
	return params
}

// Updates a Bundle in an account.
func (c *ApiService) UpdateBundle(Sid string, params *UpdateBundleParams) (*NumbersV2Bundle, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Email != nil {
		data.Set("Email", *params.Email)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Bundle{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
