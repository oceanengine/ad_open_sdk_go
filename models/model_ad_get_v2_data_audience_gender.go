/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceGender
type AdGetV2DataAudienceGender string

// List of ad_get_v2_data_audience_gender
const (
	NONE_AdGetV2DataAudienceGender             AdGetV2DataAudienceGender = "NONE"
	GENDER_UNLIMITED_AdGetV2DataAudienceGender AdGetV2DataAudienceGender = "GENDER_UNLIMITED"
	GENDER_MALE_AdGetV2DataAudienceGender      AdGetV2DataAudienceGender = "GENDER_MALE"
	GENDER_FEMALE_AdGetV2DataAudienceGender    AdGetV2DataAudienceGender = "GENDER_FEMALE"
)

// Ptr returns reference to ad_get_v2_data_audience_gender value
func (v AdGetV2DataAudienceGender) Ptr() *AdGetV2DataAudienceGender {
	return &v
}
