/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListGender
type ReportAdGetV2DataListGender string

// List of report_ad_get_v2_data_list_gender
const (
	NONE_ReportAdGetV2DataListGender             ReportAdGetV2DataListGender = "NONE"
	GENDER_MALE_ReportAdGetV2DataListGender      ReportAdGetV2DataListGender = "GENDER_MALE"
	GENDER_FEMALE_ReportAdGetV2DataListGender    ReportAdGetV2DataListGender = "GENDER_FEMALE"
	GENDER_UNLIMITED_ReportAdGetV2DataListGender ReportAdGetV2DataListGender = "GENDER_UNLIMITED"
)

// Ptr returns reference to report_ad_get_v2_data_list_gender value
func (v ReportAdGetV2DataListGender) Ptr() *ReportAdGetV2DataListGender {
	return &v
}
