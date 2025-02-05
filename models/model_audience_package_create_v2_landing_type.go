/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2LandingType
type AudiencePackageCreateV2LandingType string

// List of audience_package_create_v2_landing_type
const (
	SHOP_AudiencePackageCreateV2LandingType        AudiencePackageCreateV2LandingType = "SHOP"
	APP_IOS_AudiencePackageCreateV2LandingType     AudiencePackageCreateV2LandingType = "APP_IOS"
	AWEME_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "AWEME"
	APP_ANDROID_AudiencePackageCreateV2LandingType AudiencePackageCreateV2LandingType = "APP_ANDROID"
	LIVE_AudiencePackageCreateV2LandingType        AudiencePackageCreateV2LandingType = "LIVE"
	EXTERNAL_AudiencePackageCreateV2LandingType    AudiencePackageCreateV2LandingType = "EXTERNAL"
	QUICK_APP_AudiencePackageCreateV2LandingType   AudiencePackageCreateV2LandingType = "QUICK_APP"
	MICRO_GAME_AudiencePackageCreateV2LandingType  AudiencePackageCreateV2LandingType = "MICRO_GAME"
	GOODS_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "GOODS"
	ARTICLE_AudiencePackageCreateV2LandingType     AudiencePackageCreateV2LandingType = "ARTICLE"
	STORE_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "STORE"
	DPA_AudiencePackageCreateV2LandingType         AudiencePackageCreateV2LandingType = "DPA"
)

// Ptr returns reference to audience_package_create_v2_landing_type value
func (v AudiencePackageCreateV2LandingType) Ptr() *AudiencePackageCreateV2LandingType {
	return &v
}
