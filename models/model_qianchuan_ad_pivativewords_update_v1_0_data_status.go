/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdPivativewordsUpdateV10DataStatus
type QianchuanAdPivativewordsUpdateV10DataStatus string

// List of qianchuan_ad_pivativewords_update_v1.0_data_status
const (
	UNKNOWN_QianchuanAdPivativewordsUpdateV10DataStatus QianchuanAdPivativewordsUpdateV10DataStatus = "UNKNOWN"
	SUCCESS_QianchuanAdPivativewordsUpdateV10DataStatus QianchuanAdPivativewordsUpdateV10DataStatus = "SUCCESS"
	FAIL_QianchuanAdPivativewordsUpdateV10DataStatus    QianchuanAdPivativewordsUpdateV10DataStatus = "FAIL"
)

// Ptr returns reference to qianchuan_ad_pivativewords_update_v1.0_data_status value
func (v QianchuanAdPivativewordsUpdateV10DataStatus) Ptr() *QianchuanAdPivativewordsUpdateV10DataStatus {
	return &v
}
