/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationGetV30DataSubjectStatus
type AdvertiserQualificationGetV30DataSubjectStatus string

// List of advertiser_qualification_get_v3.0_data_subject_status
const (
	STATUS_CONFIRM_AdvertiserQualificationGetV30DataSubjectStatus         AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserQualificationGetV30DataSubjectStatus    AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserQualificationGetV30DataSubjectStatus      AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserQualificationGetV30DataSubjectStatus AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserQualificationGetV30DataSubjectStatus    AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_WAIT_CONFIRM"
)

// Ptr returns reference to advertiser_qualification_get_v3.0_data_subject_status value
func (v AdvertiserQualificationGetV30DataSubjectStatus) Ptr() *AdvertiserQualificationGetV30DataSubjectStatus {
	return &v
}
