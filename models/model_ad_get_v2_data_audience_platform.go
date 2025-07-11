/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudiencePlatform
type AdGetV2DataAudiencePlatform string

// List of ad_get_v2_data_audience_platform
const (
	IOS_AdGetV2DataAudiencePlatform     AdGetV2DataAudiencePlatform = "IOS"
	HARMONY_AdGetV2DataAudiencePlatform AdGetV2DataAudiencePlatform = "HARMONY"
	WAP_AdGetV2DataAudiencePlatform     AdGetV2DataAudiencePlatform = "WAP"
	PC_AdGetV2DataAudiencePlatform      AdGetV2DataAudiencePlatform = "PC"
	ANDROID_AdGetV2DataAudiencePlatform AdGetV2DataAudiencePlatform = "ANDROID"
	IPAD_AdGetV2DataAudiencePlatform    AdGetV2DataAudiencePlatform = "IPAD"
	NONE_AdGetV2DataAudiencePlatform    AdGetV2DataAudiencePlatform = "NONE"
)

// Ptr returns reference to ad_get_v2_data_audience_platform value
func (v AdGetV2DataAudiencePlatform) Ptr() *AdGetV2DataAudiencePlatform {
	return &v
}
