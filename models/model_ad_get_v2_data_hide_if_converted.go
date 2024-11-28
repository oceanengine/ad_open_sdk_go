/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataHideIfConverted
type AdGetV2DataHideIfConverted string

// List of ad_get_v2_data_hide_if_converted
const (
	AD_AdGetV2DataHideIfConverted           AdGetV2DataHideIfConverted = "AD"
	CAMPAIGN_AdGetV2DataHideIfConverted     AdGetV2DataHideIfConverted = "CAMPAIGN"
	CUSTOMER_AdGetV2DataHideIfConverted     AdGetV2DataHideIfConverted = "CUSTOMER"
	APP_AdGetV2DataHideIfConverted          AdGetV2DataHideIfConverted = "APP"
	NO_EXCLUDE_AdGetV2DataHideIfConverted   AdGetV2DataHideIfConverted = "NO_EXCLUDE"
	ADVERTISER_AdGetV2DataHideIfConverted   AdGetV2DataHideIfConverted = "ADVERTISER"
	ORGANIZATION_AdGetV2DataHideIfConverted AdGetV2DataHideIfConverted = "ORGANIZATION"
	GLOBAL_APP_AdGetV2DataHideIfConverted   AdGetV2DataHideIfConverted = "GLOBAL_APP"
)

// Ptr returns reference to ad_get_v2_data_hide_if_converted value
func (v AdGetV2DataHideIfConverted) Ptr() *AdGetV2DataHideIfConverted {
	return &v
}
