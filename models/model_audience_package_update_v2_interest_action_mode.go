/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2InterestActionMode
type AudiencePackageUpdateV2InterestActionMode string

// List of audience_package_update_v2_interest_action_mode
const (
	UNLIMITED_AudiencePackageUpdateV2InterestActionMode AudiencePackageUpdateV2InterestActionMode = "UNLIMITED"
	RECOMMEND_AudiencePackageUpdateV2InterestActionMode AudiencePackageUpdateV2InterestActionMode = "RECOMMEND"
	CUSTOM_AudiencePackageUpdateV2InterestActionMode    AudiencePackageUpdateV2InterestActionMode = "CUSTOM"
)

// Ptr returns reference to audience_package_update_v2_interest_action_mode value
func (v AudiencePackageUpdateV2InterestActionMode) Ptr() *AudiencePackageUpdateV2InterestActionMode {
	return &v
}
