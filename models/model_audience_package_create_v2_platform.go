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
	IPAD_AudiencePackageCreateV2Platform    AudiencePackageCreateV2Platform = "IPAD"
	PC_AudiencePackageCreateV2Platform      AudiencePackageCreateV2Platform = "PC"
	WAP_AudiencePackageCreateV2Platform     AudiencePackageCreateV2Platform = "WAP"
	ANDROID_AudiencePackageCreateV2Platform AudiencePackageCreateV2Platform = "ANDROID"
)

// Ptr returns reference to audience_package_create_v2_platform value
func (v AudiencePackageCreateV2Platform) Ptr() *AudiencePackageCreateV2Platform {
	return &v
}
