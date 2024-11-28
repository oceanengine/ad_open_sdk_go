/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2HideIfConverted
type AudiencePackageCreateV2HideIfConverted string

// List of audience_package_create_v2_hide_if_converted
const (
	AD_AudiencePackageCreateV2HideIfConverted           AudiencePackageCreateV2HideIfConverted = "AD"
	CAMPAIGN_AudiencePackageCreateV2HideIfConverted     AudiencePackageCreateV2HideIfConverted = "CAMPAIGN"
	CUSTOMER_AudiencePackageCreateV2HideIfConverted     AudiencePackageCreateV2HideIfConverted = "CUSTOMER"
	APP_AudiencePackageCreateV2HideIfConverted          AudiencePackageCreateV2HideIfConverted = "APP"
	NO_EXCLUDE_AudiencePackageCreateV2HideIfConverted   AudiencePackageCreateV2HideIfConverted = "NO_EXCLUDE"
	ADVERTISER_AudiencePackageCreateV2HideIfConverted   AudiencePackageCreateV2HideIfConverted = "ADVERTISER"
	ORGANIZATION_AudiencePackageCreateV2HideIfConverted AudiencePackageCreateV2HideIfConverted = "ORGANIZATION"
	GLOBAL_APP_AudiencePackageCreateV2HideIfConverted   AudiencePackageCreateV2HideIfConverted = "GLOBAL_APP"
)

// Ptr returns reference to audience_package_create_v2_hide_if_converted value
func (v AudiencePackageCreateV2HideIfConverted) Ptr() *AudiencePackageCreateV2HideIfConverted {
	return &v
}
