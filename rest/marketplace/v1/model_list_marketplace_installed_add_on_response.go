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

// ListMarketplaceInstalledAddOnResponse struct for ListMarketplaceInstalledAddOnResponse
type ListMarketplaceInstalledAddOnResponse struct {
	InstalledAddOns []MarketplaceInstalledAddOn               `json:"installed_add_ons,omitempty"`
	Meta            ListMarketplaceAvailableAddOnResponseMeta `json:"meta,omitempty"`
}