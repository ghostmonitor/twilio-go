/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListCallSummaries'
type ListCallSummariesParams struct {
	//
	From *string `json:"From,omitempty"`
	//
	To *string `json:"To,omitempty"`
	//
	FromCarrier *string `json:"FromCarrier,omitempty"`
	//
	ToCarrier *string `json:"ToCarrier,omitempty"`
	//
	FromCountryCode *string `json:"FromCountryCode,omitempty"`
	//
	ToCountryCode *string `json:"ToCountryCode,omitempty"`
	//
	Branded *bool `json:"Branded,omitempty"`
	//
	VerifiedCaller *bool `json:"VerifiedCaller,omitempty"`
	//
	HasTag *bool `json:"HasTag,omitempty"`
	//
	StartTime *string `json:"StartTime,omitempty"`
	//
	EndTime *string `json:"EndTime,omitempty"`
	//
	CallType *string `json:"CallType,omitempty"`
	//
	CallState *string `json:"CallState,omitempty"`
	//
	Direction *string `json:"Direction,omitempty"`
	//
	ProcessingState *string `json:"ProcessingState,omitempty"`
	//
	SortBy *string `json:"SortBy,omitempty"`
	//
	Subaccount *string `json:"Subaccount,omitempty"`
	//
	AbnormalSession *bool `json:"AbnormalSession,omitempty"`
	//
	AnsweredBy *string `json:"AnsweredBy,omitempty"`
	//
	ConnectivityIssue *string `json:"ConnectivityIssue,omitempty"`
	//
	QualityIssue *string `json:"QualityIssue,omitempty"`
	//
	Spam *bool `json:"Spam,omitempty"`
	//
	CallScore *string `json:"CallScore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCallSummariesParams) SetFrom(From string) *ListCallSummariesParams {
	params.From = &From
	return params
}
func (params *ListCallSummariesParams) SetTo(To string) *ListCallSummariesParams {
	params.To = &To
	return params
}
func (params *ListCallSummariesParams) SetFromCarrier(FromCarrier string) *ListCallSummariesParams {
	params.FromCarrier = &FromCarrier
	return params
}
func (params *ListCallSummariesParams) SetToCarrier(ToCarrier string) *ListCallSummariesParams {
	params.ToCarrier = &ToCarrier
	return params
}
func (params *ListCallSummariesParams) SetFromCountryCode(FromCountryCode string) *ListCallSummariesParams {
	params.FromCountryCode = &FromCountryCode
	return params
}
func (params *ListCallSummariesParams) SetToCountryCode(ToCountryCode string) *ListCallSummariesParams {
	params.ToCountryCode = &ToCountryCode
	return params
}
func (params *ListCallSummariesParams) SetBranded(Branded bool) *ListCallSummariesParams {
	params.Branded = &Branded
	return params
}
func (params *ListCallSummariesParams) SetVerifiedCaller(VerifiedCaller bool) *ListCallSummariesParams {
	params.VerifiedCaller = &VerifiedCaller
	return params
}
func (params *ListCallSummariesParams) SetHasTag(HasTag bool) *ListCallSummariesParams {
	params.HasTag = &HasTag
	return params
}
func (params *ListCallSummariesParams) SetStartTime(StartTime string) *ListCallSummariesParams {
	params.StartTime = &StartTime
	return params
}
func (params *ListCallSummariesParams) SetEndTime(EndTime string) *ListCallSummariesParams {
	params.EndTime = &EndTime
	return params
}
func (params *ListCallSummariesParams) SetCallType(CallType string) *ListCallSummariesParams {
	params.CallType = &CallType
	return params
}
func (params *ListCallSummariesParams) SetCallState(CallState string) *ListCallSummariesParams {
	params.CallState = &CallState
	return params
}
func (params *ListCallSummariesParams) SetDirection(Direction string) *ListCallSummariesParams {
	params.Direction = &Direction
	return params
}
func (params *ListCallSummariesParams) SetProcessingState(ProcessingState string) *ListCallSummariesParams {
	params.ProcessingState = &ProcessingState
	return params
}
func (params *ListCallSummariesParams) SetSortBy(SortBy string) *ListCallSummariesParams {
	params.SortBy = &SortBy
	return params
}
func (params *ListCallSummariesParams) SetSubaccount(Subaccount string) *ListCallSummariesParams {
	params.Subaccount = &Subaccount
	return params
}
func (params *ListCallSummariesParams) SetAbnormalSession(AbnormalSession bool) *ListCallSummariesParams {
	params.AbnormalSession = &AbnormalSession
	return params
}
func (params *ListCallSummariesParams) SetAnsweredBy(AnsweredBy string) *ListCallSummariesParams {
	params.AnsweredBy = &AnsweredBy
	return params
}
func (params *ListCallSummariesParams) SetConnectivityIssue(ConnectivityIssue string) *ListCallSummariesParams {
	params.ConnectivityIssue = &ConnectivityIssue
	return params
}
func (params *ListCallSummariesParams) SetQualityIssue(QualityIssue string) *ListCallSummariesParams {
	params.QualityIssue = &QualityIssue
	return params
}
func (params *ListCallSummariesParams) SetSpam(Spam bool) *ListCallSummariesParams {
	params.Spam = &Spam
	return params
}
func (params *ListCallSummariesParams) SetCallScore(CallScore string) *ListCallSummariesParams {
	params.CallScore = &CallScore
	return params
}
func (params *ListCallSummariesParams) SetPageSize(PageSize int) *ListCallSummariesParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCallSummariesParams) SetLimit(Limit int) *ListCallSummariesParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CallSummaries records from the API. Request is executed immediately.
func (c *ApiService) PageCallSummaries(params *ListCallSummariesParams, pageToken, pageNumber string) (*ListCallSummariesResponse, error) {
	path := "/v1/Voice/Summaries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}
	if params != nil && params.FromCarrier != nil {
		data.Set("FromCarrier", *params.FromCarrier)
	}
	if params != nil && params.ToCarrier != nil {
		data.Set("ToCarrier", *params.ToCarrier)
	}
	if params != nil && params.FromCountryCode != nil {
		data.Set("FromCountryCode", *params.FromCountryCode)
	}
	if params != nil && params.ToCountryCode != nil {
		data.Set("ToCountryCode", *params.ToCountryCode)
	}
	if params != nil && params.Branded != nil {
		data.Set("Branded", fmt.Sprint(*params.Branded))
	}
	if params != nil && params.VerifiedCaller != nil {
		data.Set("VerifiedCaller", fmt.Sprint(*params.VerifiedCaller))
	}
	if params != nil && params.HasTag != nil {
		data.Set("HasTag", fmt.Sprint(*params.HasTag))
	}
	if params != nil && params.StartTime != nil {
		data.Set("StartTime", *params.StartTime)
	}
	if params != nil && params.EndTime != nil {
		data.Set("EndTime", *params.EndTime)
	}
	if params != nil && params.CallType != nil {
		data.Set("CallType", *params.CallType)
	}
	if params != nil && params.CallState != nil {
		data.Set("CallState", *params.CallState)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.ProcessingState != nil {
		data.Set("ProcessingState", *params.ProcessingState)
	}
	if params != nil && params.SortBy != nil {
		data.Set("SortBy", *params.SortBy)
	}
	if params != nil && params.Subaccount != nil {
		data.Set("Subaccount", *params.Subaccount)
	}
	if params != nil && params.AbnormalSession != nil {
		data.Set("AbnormalSession", fmt.Sprint(*params.AbnormalSession))
	}
	if params != nil && params.AnsweredBy != nil {
		data.Set("AnsweredBy", *params.AnsweredBy)
	}
	if params != nil && params.ConnectivityIssue != nil {
		data.Set("ConnectivityIssue", *params.ConnectivityIssue)
	}
	if params != nil && params.QualityIssue != nil {
		data.Set("QualityIssue", *params.QualityIssue)
	}
	if params != nil && params.Spam != nil {
		data.Set("Spam", fmt.Sprint(*params.Spam))
	}
	if params != nil && params.CallScore != nil {
		data.Set("CallScore", *params.CallScore)
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

	ps := &ListCallSummariesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CallSummaries records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallSummaries(params *ListCallSummariesParams) ([]InsightsV1CallSummaries, error) {
	response, errors := c.StreamCallSummaries(params)

	records := make([]InsightsV1CallSummaries, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams CallSummaries records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallSummaries(params *ListCallSummariesParams) (chan InsightsV1CallSummaries, chan error) {
	if params == nil {
		params = &ListCallSummariesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan InsightsV1CallSummaries, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageCallSummaries(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamCallSummaries(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamCallSummaries(response *ListCallSummariesResponse, params *ListCallSummariesParams, recordChannel chan InsightsV1CallSummaries, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.CallSummaries
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListCallSummariesResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListCallSummariesResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListCallSummariesResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallSummariesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
