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
	"net/url"
	"strings"
)

// Allows to cancel a port in request phone number by SID
func (c *ApiService) DeletePortingPortInPhoneNumber(PortInRequestSid string, PhoneNumberSid string) error {
	path := "/v1/Porting/PortIn/{PortInRequestSid}/PhoneNumber/{PhoneNumberSid}"
	path = strings.Replace(path, "{"+"PortInRequestSid"+"}", PortInRequestSid, -1)
	path = strings.Replace(path, "{"+"PhoneNumberSid"+"}", PhoneNumberSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}