/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2LocationType
type AudiencePackageCreateV2LocationType string

// List of audience_package_create_v2_location_type
const (
	ALL_AudiencePackageCreateV2LocationType     AudiencePackageCreateV2LocationType = "ALL"
	TRAVEL_AudiencePackageCreateV2LocationType  AudiencePackageCreateV2LocationType = "TRAVEL"
	HOME_AudiencePackageCreateV2LocationType    AudiencePackageCreateV2LocationType = "HOME"
	CURRENT_AudiencePackageCreateV2LocationType AudiencePackageCreateV2LocationType = "CURRENT"
)

// Ptr returns reference to audience_package_create_v2_location_type value
func (v AudiencePackageCreateV2LocationType) Ptr() *AudiencePackageCreateV2LocationType {
	return &v
}
