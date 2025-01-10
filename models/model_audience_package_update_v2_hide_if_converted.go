/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2HideIfConverted
type AudiencePackageUpdateV2HideIfConverted string

// List of audience_package_update_v2_hide_if_converted
const (
	ADVERTISER_AudiencePackageUpdateV2HideIfConverted   AudiencePackageUpdateV2HideIfConverted = "ADVERTISER"
	AD_AudiencePackageUpdateV2HideIfConverted           AudiencePackageUpdateV2HideIfConverted = "AD"
	CAMPAIGN_AudiencePackageUpdateV2HideIfConverted     AudiencePackageUpdateV2HideIfConverted = "CAMPAIGN"
	GLOBAL_APP_AudiencePackageUpdateV2HideIfConverted   AudiencePackageUpdateV2HideIfConverted = "GLOBAL_APP"
	NO_EXCLUDE_AudiencePackageUpdateV2HideIfConverted   AudiencePackageUpdateV2HideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_AudiencePackageUpdateV2HideIfConverted AudiencePackageUpdateV2HideIfConverted = "ORGANIZATION"
	APP_AudiencePackageUpdateV2HideIfConverted          AudiencePackageUpdateV2HideIfConverted = "APP"
	CUSTOMER_AudiencePackageUpdateV2HideIfConverted     AudiencePackageUpdateV2HideIfConverted = "CUSTOMER"
)

// Ptr returns reference to audience_package_update_v2_hide_if_converted value
func (v AudiencePackageUpdateV2HideIfConverted) Ptr() *AudiencePackageUpdateV2HideIfConverted {
	return &v
}
