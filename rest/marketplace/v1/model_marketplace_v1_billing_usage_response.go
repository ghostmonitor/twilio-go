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

// MarketplaceV1BillingUsageResponse struct for MarketplaceV1BillingUsageResponse
type MarketplaceV1BillingUsageResponse struct {
	BillableItems []MarketplaceV1InstalledAddOnBillingUsageResponseBillableItems `json:"billable_items,omitempty"`
	// Represents the total quantity submitted.
	TotalSubmitted float32 `json:"total_submitted,omitempty"`
}

func (response *MarketplaceV1BillingUsageResponse) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		BillableItems  []MarketplaceV1InstalledAddOnBillingUsageResponseBillableItems `json:"billable_items"`
		TotalSubmitted interface{}                                                    `json:"total_submitted"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = MarketplaceV1BillingUsageResponse{
		BillableItems: raw.BillableItems,
	}

	responseTotalSubmitted, err := client.UnmarshalFloat32(&raw.TotalSubmitted)
	if err != nil {
		return err
	}
	response.TotalSubmitted = *responseTotalSubmitted

	return
}