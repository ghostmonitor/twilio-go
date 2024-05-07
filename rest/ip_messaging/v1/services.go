/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}

//
func (c *ApiService) CreateService(params *CreateServiceParams) (*IpMessagingV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
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

//
func (c *ApiService) FetchService(Sid string) (*IpMessagingV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParams) SetLimit(Limit int) *ListServiceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(
	params *ListServiceParams,
	pageToken, pageNumber string,
) (*ListServiceResponse, error) {
	path := "/v1/Services"

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

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]IpMessagingV1Service, error) {
	response, errors := c.StreamService(params)

	records := make([]IpMessagingV1Service, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan IpMessagingV1Service, chan error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IpMessagingV1Service, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageService(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamService(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamService(
	response *ListServiceResponse,
	params *ListServiceParams,
	recordChannel chan IpMessagingV1Service,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Services
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	DefaultServiceRoleSid *string `json:"DefaultServiceRoleSid,omitempty"`
	//
	DefaultChannelRoleSid *string `json:"DefaultChannelRoleSid,omitempty"`
	//
	DefaultChannelCreatorRoleSid *string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	//
	ReadStatusEnabled *bool `json:"ReadStatusEnabled,omitempty"`
	//
	ReachabilityEnabled *bool `json:"ReachabilityEnabled,omitempty"`
	//
	TypingIndicatorTimeout *int `json:"TypingIndicatorTimeout,omitempty"`
	//
	ConsumptionReportInterval *int `json:"ConsumptionReportInterval,omitempty"`
	//
	NotificationsNewMessageEnabled *bool `json:"Notifications.NewMessage.Enabled,omitempty"`
	//
	NotificationsNewMessageTemplate *string `json:"Notifications.NewMessage.Template,omitempty"`
	//
	NotificationsAddedToChannelEnabled *bool `json:"Notifications.AddedToChannel.Enabled,omitempty"`
	//
	NotificationsAddedToChannelTemplate *string `json:"Notifications.AddedToChannel.Template,omitempty"`
	//
	NotificationsRemovedFromChannelEnabled *bool `json:"Notifications.RemovedFromChannel.Enabled,omitempty"`
	//
	NotificationsRemovedFromChannelTemplate *string `json:"Notifications.RemovedFromChannel.Template,omitempty"`
	//
	NotificationsInvitedToChannelEnabled *bool `json:"Notifications.InvitedToChannel.Enabled,omitempty"`
	//
	NotificationsInvitedToChannelTemplate *string `json:"Notifications.InvitedToChannel.Template,omitempty"`
	//
	PreWebhookUrl *string `json:"PreWebhookUrl,omitempty"`
	//
	PostWebhookUrl *string `json:"PostWebhookUrl,omitempty"`
	//
	WebhookMethod *string `json:"WebhookMethod,omitempty"`
	//
	WebhookFilters *[]string `json:"WebhookFilters,omitempty"`
	//
	WebhooksOnMessageSendUrl *string `json:"Webhooks.OnMessageSend.Url,omitempty"`
	//
	WebhooksOnMessageSendMethod *string `json:"Webhooks.OnMessageSend.Method,omitempty"`
	//
	WebhooksOnMessageUpdateUrl *string `json:"Webhooks.OnMessageUpdate.Url,omitempty"`
	//
	WebhooksOnMessageUpdateMethod *string `json:"Webhooks.OnMessageUpdate.Method,omitempty"`
	//
	WebhooksOnMessageRemoveUrl *string `json:"Webhooks.OnMessageRemove.Url,omitempty"`
	//
	WebhooksOnMessageRemoveMethod *string `json:"Webhooks.OnMessageRemove.Method,omitempty"`
	//
	WebhooksOnChannelAddUrl *string `json:"Webhooks.OnChannelAdd.Url,omitempty"`
	//
	WebhooksOnChannelAddMethod *string `json:"Webhooks.OnChannelAdd.Method,omitempty"`
	//
	WebhooksOnChannelDestroyUrl *string `json:"Webhooks.OnChannelDestroy.Url,omitempty"`
	//
	WebhooksOnChannelDestroyMethod *string `json:"Webhooks.OnChannelDestroy.Method,omitempty"`
	//
	WebhooksOnChannelUpdateUrl *string `json:"Webhooks.OnChannelUpdate.Url,omitempty"`
	//
	WebhooksOnChannelUpdateMethod *string `json:"Webhooks.OnChannelUpdate.Method,omitempty"`
	//
	WebhooksOnMemberAddUrl *string `json:"Webhooks.OnMemberAdd.Url,omitempty"`
	//
	WebhooksOnMemberAddMethod *string `json:"Webhooks.OnMemberAdd.Method,omitempty"`
	//
	WebhooksOnMemberRemoveUrl *string `json:"Webhooks.OnMemberRemove.Url,omitempty"`
	//
	WebhooksOnMemberRemoveMethod *string `json:"Webhooks.OnMemberRemove.Method,omitempty"`
	//
	WebhooksOnMessageSentUrl *string `json:"Webhooks.OnMessageSent.Url,omitempty"`
	//
	WebhooksOnMessageSentMethod *string `json:"Webhooks.OnMessageSent.Method,omitempty"`
	//
	WebhooksOnMessageUpdatedUrl *string `json:"Webhooks.OnMessageUpdated.Url,omitempty"`
	//
	WebhooksOnMessageUpdatedMethod *string `json:"Webhooks.OnMessageUpdated.Method,omitempty"`
	//
	WebhooksOnMessageRemovedUrl *string `json:"Webhooks.OnMessageRemoved.Url,omitempty"`
	//
	WebhooksOnMessageRemovedMethod *string `json:"Webhooks.OnMessageRemoved.Method,omitempty"`
	//
	WebhooksOnChannelAddedUrl *string `json:"Webhooks.OnChannelAdded.Url,omitempty"`
	//
	WebhooksOnChannelAddedMethod *string `json:"Webhooks.OnChannelAdded.Method,omitempty"`
	//
	WebhooksOnChannelDestroyedUrl *string `json:"Webhooks.OnChannelDestroyed.Url,omitempty"`
	//
	WebhooksOnChannelDestroyedMethod *string `json:"Webhooks.OnChannelDestroyed.Method,omitempty"`
	//
	WebhooksOnChannelUpdatedUrl *string `json:"Webhooks.OnChannelUpdated.Url,omitempty"`
	//
	WebhooksOnChannelUpdatedMethod *string `json:"Webhooks.OnChannelUpdated.Method,omitempty"`
	//
	WebhooksOnMemberAddedUrl *string `json:"Webhooks.OnMemberAdded.Url,omitempty"`
	//
	WebhooksOnMemberAddedMethod *string `json:"Webhooks.OnMemberAdded.Method,omitempty"`
	//
	WebhooksOnMemberRemovedUrl *string `json:"Webhooks.OnMemberRemoved.Url,omitempty"`
	//
	WebhooksOnMemberRemovedMethod *string `json:"Webhooks.OnMemberRemoved.Method,omitempty"`
	//
	LimitsChannelMembers *int `json:"Limits.ChannelMembers,omitempty"`
	//
	LimitsUserChannels *int `json:"Limits.UserChannels,omitempty"`
}

func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetDefaultServiceRoleSid(DefaultServiceRoleSid string) *UpdateServiceParams {
	params.DefaultServiceRoleSid = &DefaultServiceRoleSid
	return params
}
func (params *UpdateServiceParams) SetDefaultChannelRoleSid(DefaultChannelRoleSid string) *UpdateServiceParams {
	params.DefaultChannelRoleSid = &DefaultChannelRoleSid
	return params
}
func (params *UpdateServiceParams) SetDefaultChannelCreatorRoleSid(DefaultChannelCreatorRoleSid string) *UpdateServiceParams {
	params.DefaultChannelCreatorRoleSid = &DefaultChannelCreatorRoleSid
	return params
}
func (params *UpdateServiceParams) SetReadStatusEnabled(ReadStatusEnabled bool) *UpdateServiceParams {
	params.ReadStatusEnabled = &ReadStatusEnabled
	return params
}
func (params *UpdateServiceParams) SetReachabilityEnabled(ReachabilityEnabled bool) *UpdateServiceParams {
	params.ReachabilityEnabled = &ReachabilityEnabled
	return params
}
func (params *UpdateServiceParams) SetTypingIndicatorTimeout(TypingIndicatorTimeout int) *UpdateServiceParams {
	params.TypingIndicatorTimeout = &TypingIndicatorTimeout
	return params
}
func (params *UpdateServiceParams) SetConsumptionReportInterval(ConsumptionReportInterval int) *UpdateServiceParams {
	params.ConsumptionReportInterval = &ConsumptionReportInterval
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageEnabled(NotificationsNewMessageEnabled bool) *UpdateServiceParams {
	params.NotificationsNewMessageEnabled = &NotificationsNewMessageEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsNewMessageTemplate(NotificationsNewMessageTemplate string) *UpdateServiceParams {
	params.NotificationsNewMessageTemplate = &NotificationsNewMessageTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelEnabled(NotificationsAddedToChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsAddedToChannelEnabled = &NotificationsAddedToChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsAddedToChannelTemplate(NotificationsAddedToChannelTemplate string) *UpdateServiceParams {
	params.NotificationsAddedToChannelTemplate = &NotificationsAddedToChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelEnabled(NotificationsRemovedFromChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelEnabled = &NotificationsRemovedFromChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsRemovedFromChannelTemplate(NotificationsRemovedFromChannelTemplate string) *UpdateServiceParams {
	params.NotificationsRemovedFromChannelTemplate = &NotificationsRemovedFromChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelEnabled(NotificationsInvitedToChannelEnabled bool) *UpdateServiceParams {
	params.NotificationsInvitedToChannelEnabled = &NotificationsInvitedToChannelEnabled
	return params
}
func (params *UpdateServiceParams) SetNotificationsInvitedToChannelTemplate(NotificationsInvitedToChannelTemplate string) *UpdateServiceParams {
	params.NotificationsInvitedToChannelTemplate = &NotificationsInvitedToChannelTemplate
	return params
}
func (params *UpdateServiceParams) SetPreWebhookUrl(PreWebhookUrl string) *UpdateServiceParams {
	params.PreWebhookUrl = &PreWebhookUrl
	return params
}
func (params *UpdateServiceParams) SetPostWebhookUrl(PostWebhookUrl string) *UpdateServiceParams {
	params.PostWebhookUrl = &PostWebhookUrl
	return params
}
func (params *UpdateServiceParams) SetWebhookMethod(WebhookMethod string) *UpdateServiceParams {
	params.WebhookMethod = &WebhookMethod
	return params
}
func (params *UpdateServiceParams) SetWebhookFilters(WebhookFilters []string) *UpdateServiceParams {
	params.WebhookFilters = &WebhookFilters
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageSendUrl(WebhooksOnMessageSendUrl string) *UpdateServiceParams {
	params.WebhooksOnMessageSendUrl = &WebhooksOnMessageSendUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageSendMethod(WebhooksOnMessageSendMethod string) *UpdateServiceParams {
	params.WebhooksOnMessageSendMethod = &WebhooksOnMessageSendMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageUpdateUrl(WebhooksOnMessageUpdateUrl string) *UpdateServiceParams {
	params.WebhooksOnMessageUpdateUrl = &WebhooksOnMessageUpdateUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageUpdateMethod(WebhooksOnMessageUpdateMethod string) *UpdateServiceParams {
	params.WebhooksOnMessageUpdateMethod = &WebhooksOnMessageUpdateMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageRemoveUrl(WebhooksOnMessageRemoveUrl string) *UpdateServiceParams {
	params.WebhooksOnMessageRemoveUrl = &WebhooksOnMessageRemoveUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageRemoveMethod(WebhooksOnMessageRemoveMethod string) *UpdateServiceParams {
	params.WebhooksOnMessageRemoveMethod = &WebhooksOnMessageRemoveMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelAddUrl(WebhooksOnChannelAddUrl string) *UpdateServiceParams {
	params.WebhooksOnChannelAddUrl = &WebhooksOnChannelAddUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelAddMethod(WebhooksOnChannelAddMethod string) *UpdateServiceParams {
	params.WebhooksOnChannelAddMethod = &WebhooksOnChannelAddMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelDestroyUrl(WebhooksOnChannelDestroyUrl string) *UpdateServiceParams {
	params.WebhooksOnChannelDestroyUrl = &WebhooksOnChannelDestroyUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelDestroyMethod(WebhooksOnChannelDestroyMethod string) *UpdateServiceParams {
	params.WebhooksOnChannelDestroyMethod = &WebhooksOnChannelDestroyMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelUpdateUrl(WebhooksOnChannelUpdateUrl string) *UpdateServiceParams {
	params.WebhooksOnChannelUpdateUrl = &WebhooksOnChannelUpdateUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelUpdateMethod(WebhooksOnChannelUpdateMethod string) *UpdateServiceParams {
	params.WebhooksOnChannelUpdateMethod = &WebhooksOnChannelUpdateMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberAddUrl(WebhooksOnMemberAddUrl string) *UpdateServiceParams {
	params.WebhooksOnMemberAddUrl = &WebhooksOnMemberAddUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberAddMethod(WebhooksOnMemberAddMethod string) *UpdateServiceParams {
	params.WebhooksOnMemberAddMethod = &WebhooksOnMemberAddMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberRemoveUrl(WebhooksOnMemberRemoveUrl string) *UpdateServiceParams {
	params.WebhooksOnMemberRemoveUrl = &WebhooksOnMemberRemoveUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberRemoveMethod(WebhooksOnMemberRemoveMethod string) *UpdateServiceParams {
	params.WebhooksOnMemberRemoveMethod = &WebhooksOnMemberRemoveMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageSentUrl(WebhooksOnMessageSentUrl string) *UpdateServiceParams {
	params.WebhooksOnMessageSentUrl = &WebhooksOnMessageSentUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageSentMethod(WebhooksOnMessageSentMethod string) *UpdateServiceParams {
	params.WebhooksOnMessageSentMethod = &WebhooksOnMessageSentMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageUpdatedUrl(WebhooksOnMessageUpdatedUrl string) *UpdateServiceParams {
	params.WebhooksOnMessageUpdatedUrl = &WebhooksOnMessageUpdatedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageUpdatedMethod(WebhooksOnMessageUpdatedMethod string) *UpdateServiceParams {
	params.WebhooksOnMessageUpdatedMethod = &WebhooksOnMessageUpdatedMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageRemovedUrl(WebhooksOnMessageRemovedUrl string) *UpdateServiceParams {
	params.WebhooksOnMessageRemovedUrl = &WebhooksOnMessageRemovedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMessageRemovedMethod(WebhooksOnMessageRemovedMethod string) *UpdateServiceParams {
	params.WebhooksOnMessageRemovedMethod = &WebhooksOnMessageRemovedMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelAddedUrl(WebhooksOnChannelAddedUrl string) *UpdateServiceParams {
	params.WebhooksOnChannelAddedUrl = &WebhooksOnChannelAddedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelAddedMethod(WebhooksOnChannelAddedMethod string) *UpdateServiceParams {
	params.WebhooksOnChannelAddedMethod = &WebhooksOnChannelAddedMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelDestroyedUrl(WebhooksOnChannelDestroyedUrl string) *UpdateServiceParams {
	params.WebhooksOnChannelDestroyedUrl = &WebhooksOnChannelDestroyedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelDestroyedMethod(WebhooksOnChannelDestroyedMethod string) *UpdateServiceParams {
	params.WebhooksOnChannelDestroyedMethod = &WebhooksOnChannelDestroyedMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelUpdatedUrl(WebhooksOnChannelUpdatedUrl string) *UpdateServiceParams {
	params.WebhooksOnChannelUpdatedUrl = &WebhooksOnChannelUpdatedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnChannelUpdatedMethod(WebhooksOnChannelUpdatedMethod string) *UpdateServiceParams {
	params.WebhooksOnChannelUpdatedMethod = &WebhooksOnChannelUpdatedMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberAddedUrl(WebhooksOnMemberAddedUrl string) *UpdateServiceParams {
	params.WebhooksOnMemberAddedUrl = &WebhooksOnMemberAddedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberAddedMethod(WebhooksOnMemberAddedMethod string) *UpdateServiceParams {
	params.WebhooksOnMemberAddedMethod = &WebhooksOnMemberAddedMethod
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberRemovedUrl(WebhooksOnMemberRemovedUrl string) *UpdateServiceParams {
	params.WebhooksOnMemberRemovedUrl = &WebhooksOnMemberRemovedUrl
	return params
}
func (params *UpdateServiceParams) SetWebhooksOnMemberRemovedMethod(WebhooksOnMemberRemovedMethod string) *UpdateServiceParams {
	params.WebhooksOnMemberRemovedMethod = &WebhooksOnMemberRemovedMethod
	return params
}
func (params *UpdateServiceParams) SetLimitsChannelMembers(LimitsChannelMembers int) *UpdateServiceParams {
	params.LimitsChannelMembers = &LimitsChannelMembers
	return params
}
func (params *UpdateServiceParams) SetLimitsUserChannels(LimitsUserChannels int) *UpdateServiceParams {
	params.LimitsUserChannels = &LimitsUserChannels
	return params
}

//
func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*IpMessagingV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.DefaultServiceRoleSid != nil {
		data.Set("DefaultServiceRoleSid", *params.DefaultServiceRoleSid)
	}
	if params != nil && params.DefaultChannelRoleSid != nil {
		data.Set("DefaultChannelRoleSid", *params.DefaultChannelRoleSid)
	}
	if params != nil && params.DefaultChannelCreatorRoleSid != nil {
		data.Set("DefaultChannelCreatorRoleSid", *params.DefaultChannelCreatorRoleSid)
	}
	if params != nil && params.ReadStatusEnabled != nil {
		data.Set("ReadStatusEnabled", fmt.Sprint(*params.ReadStatusEnabled))
	}
	if params != nil && params.ReachabilityEnabled != nil {
		data.Set("ReachabilityEnabled", fmt.Sprint(*params.ReachabilityEnabled))
	}
	if params != nil && params.TypingIndicatorTimeout != nil {
		data.Set("TypingIndicatorTimeout", fmt.Sprint(*params.TypingIndicatorTimeout))
	}
	if params != nil && params.ConsumptionReportInterval != nil {
		data.Set("ConsumptionReportInterval", fmt.Sprint(*params.ConsumptionReportInterval))
	}
	if params != nil && params.NotificationsNewMessageEnabled != nil {
		data.Set("Notifications.NewMessage.Enabled", fmt.Sprint(*params.NotificationsNewMessageEnabled))
	}
	if params != nil && params.NotificationsNewMessageTemplate != nil {
		data.Set("Notifications.NewMessage.Template", *params.NotificationsNewMessageTemplate)
	}
	if params != nil && params.NotificationsAddedToChannelEnabled != nil {
		data.Set("Notifications.AddedToChannel.Enabled", fmt.Sprint(*params.NotificationsAddedToChannelEnabled))
	}
	if params != nil && params.NotificationsAddedToChannelTemplate != nil {
		data.Set("Notifications.AddedToChannel.Template", *params.NotificationsAddedToChannelTemplate)
	}
	if params != nil && params.NotificationsRemovedFromChannelEnabled != nil {
		data.Set("Notifications.RemovedFromChannel.Enabled", fmt.Sprint(*params.NotificationsRemovedFromChannelEnabled))
	}
	if params != nil && params.NotificationsRemovedFromChannelTemplate != nil {
		data.Set("Notifications.RemovedFromChannel.Template", *params.NotificationsRemovedFromChannelTemplate)
	}
	if params != nil && params.NotificationsInvitedToChannelEnabled != nil {
		data.Set("Notifications.InvitedToChannel.Enabled", fmt.Sprint(*params.NotificationsInvitedToChannelEnabled))
	}
	if params != nil && params.NotificationsInvitedToChannelTemplate != nil {
		data.Set("Notifications.InvitedToChannel.Template", *params.NotificationsInvitedToChannelTemplate)
	}
	if params != nil && params.PreWebhookUrl != nil {
		data.Set("PreWebhookUrl", *params.PreWebhookUrl)
	}
	if params != nil && params.PostWebhookUrl != nil {
		data.Set("PostWebhookUrl", *params.PostWebhookUrl)
	}
	if params != nil && params.WebhookMethod != nil {
		data.Set("WebhookMethod", *params.WebhookMethod)
	}
	if params != nil && params.WebhookFilters != nil {
		for _, item := range *params.WebhookFilters {
			data.Add("WebhookFilters", item)
		}
	}
	if params != nil && params.WebhooksOnMessageSendUrl != nil {
		data.Set("Webhooks.OnMessageSend.Url", *params.WebhooksOnMessageSendUrl)
	}
	if params != nil && params.WebhooksOnMessageSendMethod != nil {
		data.Set("Webhooks.OnMessageSend.Method", *params.WebhooksOnMessageSendMethod)
	}
	if params != nil && params.WebhooksOnMessageUpdateUrl != nil {
		data.Set("Webhooks.OnMessageUpdate.Url", *params.WebhooksOnMessageUpdateUrl)
	}
	if params != nil && params.WebhooksOnMessageUpdateMethod != nil {
		data.Set("Webhooks.OnMessageUpdate.Method", *params.WebhooksOnMessageUpdateMethod)
	}
	if params != nil && params.WebhooksOnMessageRemoveUrl != nil {
		data.Set("Webhooks.OnMessageRemove.Url", *params.WebhooksOnMessageRemoveUrl)
	}
	if params != nil && params.WebhooksOnMessageRemoveMethod != nil {
		data.Set("Webhooks.OnMessageRemove.Method", *params.WebhooksOnMessageRemoveMethod)
	}
	if params != nil && params.WebhooksOnChannelAddUrl != nil {
		data.Set("Webhooks.OnChannelAdd.Url", *params.WebhooksOnChannelAddUrl)
	}
	if params != nil && params.WebhooksOnChannelAddMethod != nil {
		data.Set("Webhooks.OnChannelAdd.Method", *params.WebhooksOnChannelAddMethod)
	}
	if params != nil && params.WebhooksOnChannelDestroyUrl != nil {
		data.Set("Webhooks.OnChannelDestroy.Url", *params.WebhooksOnChannelDestroyUrl)
	}
	if params != nil && params.WebhooksOnChannelDestroyMethod != nil {
		data.Set("Webhooks.OnChannelDestroy.Method", *params.WebhooksOnChannelDestroyMethod)
	}
	if params != nil && params.WebhooksOnChannelUpdateUrl != nil {
		data.Set("Webhooks.OnChannelUpdate.Url", *params.WebhooksOnChannelUpdateUrl)
	}
	if params != nil && params.WebhooksOnChannelUpdateMethod != nil {
		data.Set("Webhooks.OnChannelUpdate.Method", *params.WebhooksOnChannelUpdateMethod)
	}
	if params != nil && params.WebhooksOnMemberAddUrl != nil {
		data.Set("Webhooks.OnMemberAdd.Url", *params.WebhooksOnMemberAddUrl)
	}
	if params != nil && params.WebhooksOnMemberAddMethod != nil {
		data.Set("Webhooks.OnMemberAdd.Method", *params.WebhooksOnMemberAddMethod)
	}
	if params != nil && params.WebhooksOnMemberRemoveUrl != nil {
		data.Set("Webhooks.OnMemberRemove.Url", *params.WebhooksOnMemberRemoveUrl)
	}
	if params != nil && params.WebhooksOnMemberRemoveMethod != nil {
		data.Set("Webhooks.OnMemberRemove.Method", *params.WebhooksOnMemberRemoveMethod)
	}
	if params != nil && params.WebhooksOnMessageSentUrl != nil {
		data.Set("Webhooks.OnMessageSent.Url", *params.WebhooksOnMessageSentUrl)
	}
	if params != nil && params.WebhooksOnMessageSentMethod != nil {
		data.Set("Webhooks.OnMessageSent.Method", *params.WebhooksOnMessageSentMethod)
	}
	if params != nil && params.WebhooksOnMessageUpdatedUrl != nil {
		data.Set("Webhooks.OnMessageUpdated.Url", *params.WebhooksOnMessageUpdatedUrl)
	}
	if params != nil && params.WebhooksOnMessageUpdatedMethod != nil {
		data.Set("Webhooks.OnMessageUpdated.Method", *params.WebhooksOnMessageUpdatedMethod)
	}
	if params != nil && params.WebhooksOnMessageRemovedUrl != nil {
		data.Set("Webhooks.OnMessageRemoved.Url", *params.WebhooksOnMessageRemovedUrl)
	}
	if params != nil && params.WebhooksOnMessageRemovedMethod != nil {
		data.Set("Webhooks.OnMessageRemoved.Method", *params.WebhooksOnMessageRemovedMethod)
	}
	if params != nil && params.WebhooksOnChannelAddedUrl != nil {
		data.Set("Webhooks.OnChannelAdded.Url", *params.WebhooksOnChannelAddedUrl)
	}
	if params != nil && params.WebhooksOnChannelAddedMethod != nil {
		data.Set("Webhooks.OnChannelAdded.Method", *params.WebhooksOnChannelAddedMethod)
	}
	if params != nil && params.WebhooksOnChannelDestroyedUrl != nil {
		data.Set("Webhooks.OnChannelDestroyed.Url", *params.WebhooksOnChannelDestroyedUrl)
	}
	if params != nil && params.WebhooksOnChannelDestroyedMethod != nil {
		data.Set("Webhooks.OnChannelDestroyed.Method", *params.WebhooksOnChannelDestroyedMethod)
	}
	if params != nil && params.WebhooksOnChannelUpdatedUrl != nil {
		data.Set("Webhooks.OnChannelUpdated.Url", *params.WebhooksOnChannelUpdatedUrl)
	}
	if params != nil && params.WebhooksOnChannelUpdatedMethod != nil {
		data.Set("Webhooks.OnChannelUpdated.Method", *params.WebhooksOnChannelUpdatedMethod)
	}
	if params != nil && params.WebhooksOnMemberAddedUrl != nil {
		data.Set("Webhooks.OnMemberAdded.Url", *params.WebhooksOnMemberAddedUrl)
	}
	if params != nil && params.WebhooksOnMemberAddedMethod != nil {
		data.Set("Webhooks.OnMemberAdded.Method", *params.WebhooksOnMemberAddedMethod)
	}
	if params != nil && params.WebhooksOnMemberRemovedUrl != nil {
		data.Set("Webhooks.OnMemberRemoved.Url", *params.WebhooksOnMemberRemovedUrl)
	}
	if params != nil && params.WebhooksOnMemberRemovedMethod != nil {
		data.Set("Webhooks.OnMemberRemoved.Method", *params.WebhooksOnMemberRemovedMethod)
	}
	if params != nil && params.LimitsChannelMembers != nil {
		data.Set("Limits.ChannelMembers", fmt.Sprint(*params.LimitsChannelMembers))
	}
	if params != nil && params.LimitsUserChannels != nil {
		data.Set("Limits.UserChannels", fmt.Sprint(*params.LimitsUserChannels))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
