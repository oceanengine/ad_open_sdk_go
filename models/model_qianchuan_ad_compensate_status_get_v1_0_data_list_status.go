/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCompensateStatusGetV10DataListStatus
type QianchuanAdCompensateStatusGetV10DataListStatus string

// List of qianchuan_ad_compensate_status_get_v1.0_data_list_status
const (
	SUCCESS_QianchuanAdCompensateStatusGetV10DataListStatus QianchuanAdCompensateStatusGetV10DataListStatus = "SUCCESS"
	FAILED_QianchuanAdCompensateStatusGetV10DataListStatus  QianchuanAdCompensateStatusGetV10DataListStatus = "FAILED"
)

// Ptr returns reference to qianchuan_ad_compensate_status_get_v1.0_data_list_status value
func (v QianchuanAdCompensateStatusGetV10DataListStatus) Ptr() *QianchuanAdCompensateStatusGetV10DataListStatus {
	return &v
}
