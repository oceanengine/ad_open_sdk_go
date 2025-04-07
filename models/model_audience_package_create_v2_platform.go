/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2Platform
type AudiencePackageCreateV2Platform string

// List of audience_package_create_v2_platform
const (
	IOS_AudiencePackageCreateV2Platform     AudiencePackageCreateV2Platform = "IOS"
	PC_AudiencePackageCreateV2Platform      AudiencePackageCreateV2Platform = "PC"
	ANDROID_AudiencePackageCreateV2Platform AudiencePackageCreateV2Platform = "ANDROID"
	WAP_AudiencePackageCreateV2Platform     AudiencePackageCreateV2Platform = "WAP"
	IPAD_AudiencePackageCreateV2Platform    AudiencePackageCreateV2Platform = "IPAD"
)

// Ptr returns reference to audience_package_create_v2_platform value
func (v AudiencePackageCreateV2Platform) Ptr() *AudiencePackageCreateV2Platform {
	return &v
}
