/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// Optional parameters for the method 'CreateInsightsAssessments'
type CreateInsightsAssessmentsParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The SID of the category
	CategorySid *string `json:"CategorySid,omitempty"`
	// The name of the category
	CategoryName *string `json:"CategoryName,omitempty"`
	// Segment Id of the conversation
	SegmentId *string `json:"SegmentId,omitempty"`
	// The id of the Agent
	AgentId *string `json:"AgentId,omitempty"`
	// The offset of the conversation.
	Offset *float32 `json:"Offset,omitempty"`
	// The question SID selected for assessment
	MetricId *string `json:"MetricId,omitempty"`
	// The question name of the assessment
	MetricName *string `json:"MetricName,omitempty"`
	// The answer text selected by user
	AnswerText *string `json:"AnswerText,omitempty"`
	// The id of the answer selected by user
	AnswerId *string `json:"AnswerId,omitempty"`
	// Questionnaire SID of the associated question
	QuestionnaireSid *string `json:"QuestionnaireSid,omitempty"`
}

func (params *CreateInsightsAssessmentsParams) SetAuthorization(Authorization string) *CreateInsightsAssessmentsParams {
	params.Authorization = &Authorization
	return params
}
func (params *CreateInsightsAssessmentsParams) SetCategorySid(CategorySid string) *CreateInsightsAssessmentsParams {
	params.CategorySid = &CategorySid
	return params
}
func (params *CreateInsightsAssessmentsParams) SetCategoryName(CategoryName string) *CreateInsightsAssessmentsParams {
	params.CategoryName = &CategoryName
	return params
}
func (params *CreateInsightsAssessmentsParams) SetSegmentId(SegmentId string) *CreateInsightsAssessmentsParams {
	params.SegmentId = &SegmentId
	return params
}
func (params *CreateInsightsAssessmentsParams) SetAgentId(AgentId string) *CreateInsightsAssessmentsParams {
	params.AgentId = &AgentId
	return params
}
func (params *CreateInsightsAssessmentsParams) SetOffset(Offset float32) *CreateInsightsAssessmentsParams {
	params.Offset = &Offset
	return params
}
func (params *CreateInsightsAssessmentsParams) SetMetricId(MetricId string) *CreateInsightsAssessmentsParams {
	params.MetricId = &MetricId
	return params
}
func (params *CreateInsightsAssessmentsParams) SetMetricName(MetricName string) *CreateInsightsAssessmentsParams {
	params.MetricName = &MetricName
	return params
}
func (params *CreateInsightsAssessmentsParams) SetAnswerText(AnswerText string) *CreateInsightsAssessmentsParams {
	params.AnswerText = &AnswerText
	return params
}
func (params *CreateInsightsAssessmentsParams) SetAnswerId(AnswerId string) *CreateInsightsAssessmentsParams {
	params.AnswerId = &AnswerId
	return params
}
func (params *CreateInsightsAssessmentsParams) SetQuestionnaireSid(QuestionnaireSid string) *CreateInsightsAssessmentsParams {
	params.QuestionnaireSid = &QuestionnaireSid
	return params
}

// Add assessments against conversation to dynamo db. Used in assessments screen by user. Users can select the questionnaire and pick up answers for each and every question.
func (c *ApiService) CreateInsightsAssessments(params *CreateInsightsAssessmentsParams) (*FlexV1InsightsAssessments, error) {
	path := "/v1/Insights/QualityManagement/Assessments"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CategorySid != nil {
		data.Set("CategorySid", *params.CategorySid)
	}
	if params != nil && params.CategoryName != nil {
		data.Set("CategoryName", *params.CategoryName)
	}
	if params != nil && params.SegmentId != nil {
		data.Set("SegmentId", *params.SegmentId)
	}
	if params != nil && params.AgentId != nil {
		data.Set("AgentId", *params.AgentId)
	}
	if params != nil && params.Offset != nil {
		data.Set("Offset", fmt.Sprint(*params.Offset))
	}
	if params != nil && params.MetricId != nil {
		data.Set("MetricId", *params.MetricId)
	}
	if params != nil && params.MetricName != nil {
		data.Set("MetricName", *params.MetricName)
	}
	if params != nil && params.AnswerText != nil {
		data.Set("AnswerText", *params.AnswerText)
	}
	if params != nil && params.AnswerId != nil {
		data.Set("AnswerId", *params.AnswerId)
	}
	if params != nil && params.QuestionnaireSid != nil {
		data.Set("QuestionnaireSid", *params.QuestionnaireSid)
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsAssessments{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListInsightsAssessments'
type ListInsightsAssessmentsParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The id of the segment.
	SegmentId *string `json:"SegmentId,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListInsightsAssessmentsParams) SetAuthorization(Authorization string) *ListInsightsAssessmentsParams {
	params.Authorization = &Authorization
	return params
}
func (params *ListInsightsAssessmentsParams) SetSegmentId(SegmentId string) *ListInsightsAssessmentsParams {
	params.SegmentId = &SegmentId
	return params
}
func (params *ListInsightsAssessmentsParams) SetPageSize(PageSize int) *ListInsightsAssessmentsParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListInsightsAssessmentsParams) SetLimit(Limit int) *ListInsightsAssessmentsParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of InsightsAssessments records from the API. Request is executed immediately.
func (c *ApiService) PageInsightsAssessments(
	params *ListInsightsAssessmentsParams,
	pageToken, pageNumber string,
) (*ListInsightsAssessmentsResponse, error) {
	path := "/v1/Insights/QualityManagement/Assessments"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.SegmentId != nil {
		data.Set("SegmentId", *params.SegmentId)
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

	ps := &ListInsightsAssessmentsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists InsightsAssessments records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListInsightsAssessments(params *ListInsightsAssessmentsParams) ([]FlexV1InsightsAssessments, error) {
	response, errors := c.StreamInsightsAssessments(params)

	records := make([]FlexV1InsightsAssessments, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams InsightsAssessments records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamInsightsAssessments(params *ListInsightsAssessmentsParams) (chan FlexV1InsightsAssessments, chan error) {
	if params == nil {
		params = &ListInsightsAssessmentsParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FlexV1InsightsAssessments, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageInsightsAssessments(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamInsightsAssessments(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamInsightsAssessments(
	response *ListInsightsAssessmentsResponse,
	params *ListInsightsAssessmentsParams,
	recordChannel chan FlexV1InsightsAssessments,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Assessments
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListInsightsAssessmentsResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListInsightsAssessmentsResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListInsightsAssessmentsResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListInsightsAssessmentsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateInsightsAssessments'
type UpdateInsightsAssessmentsParams struct {
	// The Authorization HTTP request header
	Authorization *string `json:"Authorization,omitempty"`
	// The offset of the conversation
	Offset *float32 `json:"Offset,omitempty"`
	// The answer text selected by user
	AnswerText *string `json:"AnswerText,omitempty"`
	// The id of the answer selected by user
	AnswerId *string `json:"AnswerId,omitempty"`
}

func (params *UpdateInsightsAssessmentsParams) SetAuthorization(Authorization string) *UpdateInsightsAssessmentsParams {
	params.Authorization = &Authorization
	return params
}
func (params *UpdateInsightsAssessmentsParams) SetOffset(Offset float32) *UpdateInsightsAssessmentsParams {
	params.Offset = &Offset
	return params
}
func (params *UpdateInsightsAssessmentsParams) SetAnswerText(AnswerText string) *UpdateInsightsAssessmentsParams {
	params.AnswerText = &AnswerText
	return params
}
func (params *UpdateInsightsAssessmentsParams) SetAnswerId(AnswerId string) *UpdateInsightsAssessmentsParams {
	params.AnswerId = &AnswerId
	return params
}

// Update a specific Assessment assessed earlier
func (c *ApiService) UpdateInsightsAssessments(
	AssessmentSid string,
	params *UpdateInsightsAssessmentsParams,
) (*FlexV1InsightsAssessments, error) {
	path := "/v1/Insights/QualityManagement/Assessments/{AssessmentSid}"
	path = strings.Replace(path, "{"+"AssessmentSid"+"}", AssessmentSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Offset != nil {
		data.Set("Offset", fmt.Sprint(*params.Offset))
	}
	if params != nil && params.AnswerText != nil {
		data.Set("AnswerText", *params.AnswerText)
	}
	if params != nil && params.AnswerId != nil {
		data.Set("AnswerId", *params.AnswerId)
	}

	if params != nil && params.Authorization != nil {
		headers["Authorization"] = *params.Authorization
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1InsightsAssessments{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
