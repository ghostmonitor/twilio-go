/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateCredentialAws'
type CreateCredentialAwsParams struct {
	// The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request.
	AccountSid *string `json:"AccountSid,omitempty"`
	// A string that contains the AWS access credentials in the format `<AWS_ACCESS_KEY_ID>:<AWS_SECRET_ACCESS_KEY>`. For example, `AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
	Credentials *string `json:"Credentials,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateCredentialAwsParams) SetAccountSid(AccountSid string) *CreateCredentialAwsParams {
	params.AccountSid = &AccountSid
	return params
}
func (params *CreateCredentialAwsParams) SetCredentials(Credentials string) *CreateCredentialAwsParams {
	params.Credentials = &Credentials
	return params
}
func (params *CreateCredentialAwsParams) SetFriendlyName(FriendlyName string) *CreateCredentialAwsParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new AWS Credential
func (c *ApiService) CreateCredentialAws(params *CreateCredentialAwsParams) (*AccountsV1CredentialCredentialAws, error) {
	path := "/v1/Credentials/AWS"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AccountSid != nil {
		data.Set("AccountSid", *params.AccountSid)
	}
	if params != nil && params.Credentials != nil {
		data.Set("Credentials", *params.Credentials)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialCredentialAws{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Credential from your account
func (c *ApiService) DeleteCredentialAws(Sid string) error {
	path := "/v1/Credentials/AWS/{Sid}"
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

// Fetch the AWS credentials specified by the provided Credential Sid
func (c *ApiService) FetchCredentialAws(Sid string) (*AccountsV1CredentialCredentialAws, error) {
	path := "/v1/Credentials/AWS/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AccountsV1CredentialCredentialAws{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCredentialAws'
type ListCredentialAwsParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListCredentialAwsParams) SetPageSize(PageSize int) *ListCredentialAwsParams {
	params.PageSize = &PageSize
	return params
}

// Retrieves a collection of AWS Credentials belonging to the account used to make the request
func (c *ApiService) ListCredentialAws(params *ListCredentialAwsParams) (*ListCredentialAwsResponse, error) {
	path := "/v1/Credentials/AWS"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCredentialAwsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateCredentialAws'
type UpdateCredentialAwsParams struct {
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateCredentialAwsParams) SetFriendlyName(FriendlyName string) *UpdateCredentialAwsParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Modify the properties of a given Account
func (c *ApiService) UpdateCredentialAws(Sid string, params *UpdateCredentialAwsParams) (*AccountsV1CredentialCredentialAws, error) {
	path := "/v1/Credentials/AWS/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

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

	ps := &AccountsV1CredentialCredentialAws{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}