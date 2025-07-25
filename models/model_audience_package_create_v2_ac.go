/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2Ac
type AudiencePackageCreateV2Ac string

// List of audience_package_create_v2_ac
const (
	Enum_4_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "4G"
	Enum_3_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "3G"
	Enum_2_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "2G"
	Enum_5_G_AudiencePackageCreateV2Ac AudiencePackageCreateV2Ac = "5G"
	WIFI_AudiencePackageCreateV2Ac     AudiencePackageCreateV2Ac = "WIFI"
)

// Ptr returns reference to audience_package_create_v2_ac value
func (v AudiencePackageCreateV2Ac) Ptr() *AudiencePackageCreateV2Ac {
	return &v
}
