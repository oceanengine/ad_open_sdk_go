/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2MarketingGoal
type AudiencePackageCreateV2MarketingGoal string

// List of audience_package_create_v2_marketing_goal
const (
	LIVE_AudiencePackageCreateV2MarketingGoal            AudiencePackageCreateV2MarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_AudiencePackageCreateV2MarketingGoal AudiencePackageCreateV2MarketingGoal = "VIDEO_AND_IMAGE"
)

// Ptr returns reference to audience_package_create_v2_marketing_goal value
func (v AudiencePackageCreateV2MarketingGoal) Ptr() *AudiencePackageCreateV2MarketingGoal {
	return &v
}
