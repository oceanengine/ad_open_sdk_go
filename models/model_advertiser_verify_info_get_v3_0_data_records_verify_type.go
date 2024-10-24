/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserVerifyInfoGetV30DataRecordsVerifyType
type AdvertiserVerifyInfoGetV30DataRecordsVerifyType string

// List of advertiser_verify_info_get_v3.0_data_records_verify_type
const (
	AD_ARK_SKIP_AdvertiserVerifyInfoGetV30DataRecordsVerifyType      AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "AD_ARK_SKIP"
	AUTH_AdvertiserVerifyInfoGetV30DataRecordsVerifyType             AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "AUTH"
	AUTH_LETTER_AdvertiserVerifyInfoGetV30DataRecordsVerifyType      AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "AUTH_LETTER"
	BYTEDANCE_PAY_AdvertiserVerifyInfoGetV30DataRecordsVerifyType    AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "BYTEDANCE_PAY"
	CUSTOMER_PAY_AdvertiserVerifyInfoGetV30DataRecordsVerifyType     AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "CUSTOMER_PAY"
	ELEC_CERT_SKIP_AdvertiserVerifyInfoGetV30DataRecordsVerifyType   AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "ELEC_CERT_SKIP"
	FACE_DETECT_AdvertiserVerifyInfoGetV30DataRecordsVerifyType      AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "FACE_DETECT"
	PERSON_VERIFY3_AdvertiserVerifyInfoGetV30DataRecordsVerifyType   AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "PERSON_VERIFY3"
	PERSON_VERIFY4_AdvertiserVerifyInfoGetV30DataRecordsVerifyType   AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "PERSON_VERIFY4"
	SKIP_AdvertiserVerifyInfoGetV30DataRecordsVerifyType             AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "SKIP"
	THRID_PARTY_AUTH_AdvertiserVerifyInfoGetV30DataRecordsVerifyType AdvertiserVerifyInfoGetV30DataRecordsVerifyType = "THRID_PARTY_AUTH"
)

// Ptr returns reference to advertiser_verify_info_get_v3.0_data_records_verify_type value
func (v AdvertiserVerifyInfoGetV30DataRecordsVerifyType) Ptr() *AdvertiserVerifyInfoGetV30DataRecordsVerifyType {
	return &v
}