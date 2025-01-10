/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2FilterEvent
type AudiencePackageUpdateV2FilterEvent string

// List of audience_package_update_v2_filter_event
const (
	AD_CONVERT_EXTERNAL_ACTION_AudiencePackageUpdateV2FilterEvent      AudiencePackageUpdateV2FilterEvent = "AD_CONVERT_EXTERNAL_ACTION"
	AD_CONVERT_TYPE_PAY_AudiencePackageUpdateV2FilterEvent             AudiencePackageUpdateV2FilterEvent = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_AudiencePackageUpdateV2FilterEvent AudiencePackageUpdateV2FilterEvent = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ACTIVE_AudiencePackageUpdateV2FilterEvent          AudiencePackageUpdateV2FilterEvent = "AD_CONVERT_TYPE_ACTIVE"
)

// Ptr returns reference to audience_package_update_v2_filter_event value
func (v AudiencePackageUpdateV2FilterEvent) Ptr() *AudiencePackageUpdateV2FilterEvent {
	return &v
}
