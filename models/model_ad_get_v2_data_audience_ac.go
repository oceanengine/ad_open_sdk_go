/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceAc
type AdGetV2DataAudienceAc string

// List of ad_get_v2_data_audience_ac
const (
	WIFI_AdGetV2DataAudienceAc     AdGetV2DataAudienceAc = "WIFI"
	Enum_2_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "2G"
	Enum_3_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "3G"
	Enum_5_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "5G"
	Enum_4_G_AdGetV2DataAudienceAc AdGetV2DataAudienceAc = "4G"
)

// Ptr returns reference to ad_get_v2_data_audience_ac value
func (v AdGetV2DataAudienceAc) Ptr() *AdGetV2DataAudienceAc {
	return &v
}
