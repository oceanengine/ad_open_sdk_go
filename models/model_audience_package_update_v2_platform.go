/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2Platform
type AudiencePackageUpdateV2Platform string

// List of audience_package_update_v2_platform
const (
	ANDROID_AudiencePackageUpdateV2Platform AudiencePackageUpdateV2Platform = "ANDROID"
	PC_AudiencePackageUpdateV2Platform      AudiencePackageUpdateV2Platform = "PC"
	IPAD_AudiencePackageUpdateV2Platform    AudiencePackageUpdateV2Platform = "IPAD"
	WAP_AudiencePackageUpdateV2Platform     AudiencePackageUpdateV2Platform = "WAP"
	IOS_AudiencePackageUpdateV2Platform     AudiencePackageUpdateV2Platform = "IOS"
)

// Ptr returns reference to audience_package_update_v2_platform value
func (v AudiencePackageUpdateV2Platform) Ptr() *AudiencePackageUpdateV2Platform {
	return &v
}
