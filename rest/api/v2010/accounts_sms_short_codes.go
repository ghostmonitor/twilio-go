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

// Optional parameters for the method 'FetchShortCode'
type FetchShortCodeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchShortCodeParams) SetPathAccountSid(PathAccountSid string) *FetchShortCodeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a short code
func (c *ApiService) FetchShortCode(Sid string, params *FetchShortCodeParams) (*ApiV2010ShortCode, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json"
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

	ps := &ApiV2010ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListShortCode'
type ListShortCodeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The string that identifies the ShortCode resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Only show the ShortCode resources that match this pattern. You can specify partial numbers and use '*' as a wildcard for any digit.
	ShortCode *string `json:"ShortCode,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListShortCodeParams) SetPathAccountSid(PathAccountSid string) *ListShortCodeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListShortCodeParams) SetFriendlyName(FriendlyName string) *ListShortCodeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListShortCodeParams) SetShortCode(ShortCode string) *ListShortCodeParams {
	params.ShortCode = &ShortCode
	return params
}
func (params *ListShortCodeParams) SetPageSize(PageSize int) *ListShortCodeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListShortCodeParams) SetLimit(Limit int) *ListShortCodeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ShortCode records from the API. Request is executed immediately.
func (c *ApiService) PageShortCode(
	params *ListShortCodeParams,
	pageToken, pageNumber string,
) (*ListShortCodeResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ShortCode != nil {
		data.Set("ShortCode", *params.ShortCode)
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

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ShortCode records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListShortCode(params *ListShortCodeParams) ([]ApiV2010ShortCode, error) {
	response, errors := c.StreamShortCode(params)

	records := make([]ApiV2010ShortCode, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ShortCode records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamShortCode(params *ListShortCodeParams) (chan ApiV2010ShortCode, chan error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010ShortCode, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageShortCode(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamShortCode(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamShortCode(
	response *ListShortCodeResponse,
	params *ListShortCodeParams,
	recordChannel chan ApiV2010ShortCode,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.ShortCodes
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListShortCodeResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListShortCodeResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListShortCodeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateShortCode'
type UpdateShortCodeParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the `FriendlyName` is the short code.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`.
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// The URL we should call when receiving an incoming SMS message to this short code.
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The HTTP method we should use when calling the `sms_url`. Can be: `GET` or `POST`.
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL that we should call if an error occurs while retrieving or executing the TwiML from `sms_url`.
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method that we should use to call the `sms_fallback_url`. Can be: `GET` or `POST`.
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
}

func (params *UpdateShortCodeParams) SetPathAccountSid(PathAccountSid string) *UpdateShortCodeParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateShortCodeParams) SetFriendlyName(FriendlyName string) *UpdateShortCodeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateShortCodeParams) SetApiVersion(ApiVersion string) *UpdateShortCodeParams {
	params.ApiVersion = &ApiVersion
	return params
}
func (params *UpdateShortCodeParams) SetSmsUrl(SmsUrl string) *UpdateShortCodeParams {
	params.SmsUrl = &SmsUrl
	return params
}
func (params *UpdateShortCodeParams) SetSmsMethod(SmsMethod string) *UpdateShortCodeParams {
	params.SmsMethod = &SmsMethod
	return params
}
func (params *UpdateShortCodeParams) SetSmsFallbackUrl(SmsFallbackUrl string) *UpdateShortCodeParams {
	params.SmsFallbackUrl = &SmsFallbackUrl
	return params
}
func (params *UpdateShortCodeParams) SetSmsFallbackMethod(SmsFallbackMethod string) *UpdateShortCodeParams {
	params.SmsFallbackMethod = &SmsFallbackMethod
	return params
}

// Update a short code with the following parameters
func (c *ApiService) UpdateShortCode(Sid string, params *UpdateShortCodeParams) (*ApiV2010ShortCode, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json"
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

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ApiVersion != nil {
		data.Set("ApiVersion", *params.ApiVersion)
	}
	if params != nil && params.SmsUrl != nil {
		data.Set("SmsUrl", *params.SmsUrl)
	}
	if params != nil && params.SmsMethod != nil {
		data.Set("SmsMethod", *params.SmsMethod)
	}
	if params != nil && params.SmsFallbackUrl != nil {
		data.Set("SmsFallbackUrl", *params.SmsFallbackUrl)
	}
	if params != nil && params.SmsFallbackMethod != nil {
		data.Set("SmsFallbackMethod", *params.SmsFallbackMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010ShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
