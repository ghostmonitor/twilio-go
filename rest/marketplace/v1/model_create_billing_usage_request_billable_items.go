/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Marketplace
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"

	"github.com/ghostmonitor/twilio-go/client"
)

// CreateBillingUsageRequestBillableItems struct for CreateBillingUsageRequestBillableItems
type CreateBillingUsageRequestBillableItems struct {
	//
	Quantity float32 `json:"quantity"`
	//
	Sid string `json:"sid"`
}

func (response *CreateBillingUsageRequestBillableItems) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		Quantity interface{} `json:"quantity"`
		Sid      string      `json:"sid"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = CreateBillingUsageRequestBillableItems{
		Sid: raw.Sid,
	}

	responseQuantity, err := client.UnmarshalFloat32(&raw.Quantity)
	if err != nil {
		return err
	}
	response.Quantity = *responseQuantity

	return
}
