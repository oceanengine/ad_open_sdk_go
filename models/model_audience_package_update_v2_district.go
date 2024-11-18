/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2District
type AudiencePackageUpdateV2District string

// List of audience_package_update_v2_district
const (
	NONE_AudiencePackageUpdateV2District              AudiencePackageUpdateV2District = "NONE"
	REGION_AudiencePackageUpdateV2District            AudiencePackageUpdateV2District = "REGION"
	BUSINESS_DISTRICT_AudiencePackageUpdateV2District AudiencePackageUpdateV2District = "BUSINESS_DISTRICT"
	OVERSEA_AudiencePackageUpdateV2District           AudiencePackageUpdateV2District = "OVERSEA"
)

// Ptr returns reference to audience_package_update_v2_district value
func (v AudiencePackageUpdateV2District) Ptr() *AudiencePackageUpdateV2District {
	return &v
}
