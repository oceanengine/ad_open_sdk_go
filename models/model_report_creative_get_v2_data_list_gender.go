/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2DataListGender
type ReportCreativeGetV2DataListGender string

// List of report_creative_get_v2_data_list_gender
const (
	GENDER_FEMALE_ReportCreativeGetV2DataListGender    ReportCreativeGetV2DataListGender = "GENDER_FEMALE"
	GENDER_MALE_ReportCreativeGetV2DataListGender      ReportCreativeGetV2DataListGender = "GENDER_MALE"
	GENDER_UNLIMITED_ReportCreativeGetV2DataListGender ReportCreativeGetV2DataListGender = "GENDER_UNLIMITED"
	NONE_ReportCreativeGetV2DataListGender             ReportCreativeGetV2DataListGender = "NONE"
)

// Ptr returns reference to report_creative_get_v2_data_list_gender value
func (v ReportCreativeGetV2DataListGender) Ptr() *ReportCreativeGetV2DataListGender {
	return &v
}
