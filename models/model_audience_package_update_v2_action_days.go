/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2ActionDays
type AudiencePackageUpdateV2ActionDays int64

// List of audience_package_update_v2_action_days
const (
	Enum_7_AudiencePackageUpdateV2ActionDays   AudiencePackageUpdateV2ActionDays = 7
	Enum_15_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 15
	Enum_30_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 30
	Enum_60_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 60
	Enum_90_AudiencePackageUpdateV2ActionDays  AudiencePackageUpdateV2ActionDays = 90
	Enum_180_AudiencePackageUpdateV2ActionDays AudiencePackageUpdateV2ActionDays = 180
	Enum_365_AudiencePackageUpdateV2ActionDays AudiencePackageUpdateV2ActionDays = 365
)

// Ptr returns reference to audience_package_update_v2_action_days value
func (v AudiencePackageUpdateV2ActionDays) Ptr() *AudiencePackageUpdateV2ActionDays {
	return &v
}
