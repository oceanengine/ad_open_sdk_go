/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdStatusUpdateV10OptStatus
type QianchuanAdStatusUpdateV10OptStatus string

// List of qianchuan_ad_status_update_v1.0_opt_status
const (
	DISABLE_QianchuanAdStatusUpdateV10OptStatus QianchuanAdStatusUpdateV10OptStatus = "DISABLE"
	DELETE_QianchuanAdStatusUpdateV10OptStatus  QianchuanAdStatusUpdateV10OptStatus = "DELETE"
	ENABLE_QianchuanAdStatusUpdateV10OptStatus  QianchuanAdStatusUpdateV10OptStatus = "ENABLE"
	REVIVE_QianchuanAdStatusUpdateV10OptStatus  QianchuanAdStatusUpdateV10OptStatus = "REVIVE"
)

// Ptr returns reference to qianchuan_ad_status_update_v1.0_opt_status value
func (v QianchuanAdStatusUpdateV10OptStatus) Ptr() *QianchuanAdStatusUpdateV10OptStatus {
	return &v
}
