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
	IOS_AudiencePackageUpdateV2Platform     AudiencePackageUpdateV2Platform = "IOS"
	IPAD_AudiencePackageUpdateV2Platform    AudiencePackageUpdateV2Platform = "IPAD"
	PC_AudiencePackageUpdateV2Platform      AudiencePackageUpdateV2Platform = "PC"
	ANDROID_AudiencePackageUpdateV2Platform AudiencePackageUpdateV2Platform = "ANDROID"
	WAP_AudiencePackageUpdateV2Platform     AudiencePackageUpdateV2Platform = "WAP"
)

// Ptr returns reference to audience_package_update_v2_platform value
func (v AudiencePackageUpdateV2Platform) Ptr() *AudiencePackageUpdateV2Platform {
	return &v
}
