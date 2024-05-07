/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
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

// Optional parameters for the method 'CreateServiceRole'
type CreateServiceRoleParams struct {
	// A descriptive string that you create to describe the new resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	// A permission that you grant to the new role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. The values for this parameter depend on the role's `type`.
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *CreateServiceRoleParams) SetFriendlyName(FriendlyName string) *CreateServiceRoleParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceRoleParams) SetType(Type string) *CreateServiceRoleParams {
	params.Type = &Type
	return params
}
func (params *CreateServiceRoleParams) SetPermission(Permission []string) *CreateServiceRoleParams {
	params.Permission = &Permission
	return params
}

// Create a new user role in your service
func (c *ApiService) CreateServiceRole(
	ChatServiceSid string,
	params *CreateServiceRoleParams,
) (*ConversationsV1ServiceRole, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a user role from your service
func (c *ApiService) DeleteServiceRole(ChatServiceSid string, Sid string) error {
	path := "/v1/Services/{ChatServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
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

// Fetch a user role from your service
func (c *ApiService) FetchServiceRole(ChatServiceSid string, Sid string) (*ConversationsV1ServiceRole, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListServiceRole'
type ListServiceRoleParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceRoleParams) SetPageSize(PageSize int) *ListServiceRoleParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceRoleParams) SetLimit(Limit int) *ListServiceRoleParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ServiceRole records from the API. Request is executed immediately.
func (c *ApiService) PageServiceRole(
	ChatServiceSid string,
	params *ListServiceRoleParams,
	pageToken, pageNumber string,
) (*ListServiceRoleResponse, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles"

	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)

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

	ps := &ListServiceRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ServiceRole records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListServiceRole(
	ChatServiceSid string,
	params *ListServiceRoleParams,
) ([]ConversationsV1ServiceRole, error) {
	response, errors := c.StreamServiceRole(ChatServiceSid, params)

	records := make([]ConversationsV1ServiceRole, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ServiceRole records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamServiceRole(
	ChatServiceSid string,
	params *ListServiceRoleParams,
) (chan ConversationsV1ServiceRole, chan error) {
	if params == nil {
		params = &ListServiceRoleParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ServiceRole, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageServiceRole(ChatServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamServiceRole(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamServiceRole(
	response *ListServiceRoleResponse,
	params *ListServiceRoleParams,
	recordChannel chan ConversationsV1ServiceRole,
	errorChannel chan error,
) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Roles
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceRoleResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceRoleResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceRoleResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceRoleResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateServiceRole'
type UpdateServiceRoleParams struct {
	// A permission that you grant to the role. Only one permission can be granted per parameter. To assign more than one permission, repeat this parameter for each permission value. Note that the update action replaces all previously assigned permissions with those defined in the update action. To remove a permission, do not include it in the subsequent update action. The values for this parameter depend on the role's `type`.
	Permission *[]string `json:"Permission,omitempty"`
}

func (params *UpdateServiceRoleParams) SetPermission(Permission []string) *UpdateServiceRoleParams {
	params.Permission = &Permission
	return params
}

// Update an existing user role in your service
func (c *ApiService) UpdateServiceRole(
	ChatServiceSid string,
	Sid string,
	params *UpdateServiceRoleParams,
) (*ConversationsV1ServiceRole, error) {
	path := "/v1/Services/{ChatServiceSid}/Roles/{Sid}"
	path = strings.Replace(path, "{"+"ChatServiceSid"+"}", ChatServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Permission != nil {
		for _, item := range *params.Permission {
			data.Add("Permission", item)
		}
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ServiceRole{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
