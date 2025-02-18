/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceCareer
type AdGetV2DataAudienceCareer string

// List of ad_get_v2_data_audience_career
const (
	FINANCIAL_AdGetV2DataAudienceCareer       AdGetV2DataAudienceCareer = "FINANCIAL"
	COLLEGE_STUDENT_AdGetV2DataAudienceCareer AdGetV2DataAudienceCareer = "COLLEGE_STUDENT"
	IT_AdGetV2DataAudienceCareer              AdGetV2DataAudienceCareer = "IT"
	MEDICAL_STAFF_AdGetV2DataAudienceCareer   AdGetV2DataAudienceCareer = "MEDICAL_STAFF"
	TEACHER_AdGetV2DataAudienceCareer         AdGetV2DataAudienceCareer = "TEACHER"
	CIVIL_SERVANTS_AdGetV2DataAudienceCareer  AdGetV2DataAudienceCareer = "CIVIL_SERVANTS"
)

// Ptr returns reference to ad_get_v2_data_audience_career value
func (v AdGetV2DataAudienceCareer) Ptr() *AdGetV2DataAudienceCareer {
	return &v
}
