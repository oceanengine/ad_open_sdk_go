/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus
type QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus string

// List of qianchuan_tools_smart_boost_ad_boost_status_get_v1.0_data_raise_details_status
const (
	CANNOT_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "CANNOT_RAISE"
	CAN_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus    QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "CAN_RAISE"
	FINISH_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "FINISH_RAISE"
	HASPRE_RAISE_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "HASPRE_RAISE"
	RAISE_FAILED_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "RAISE_FAILED"
	RAISING_QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus      QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus = "RAISING"
)

// Ptr returns reference to qianchuan_tools_smart_boost_ad_boost_status_get_v1.0_data_raise_details_status value
func (v QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus) Ptr() *QianchuanToolsSmartBoostAdBoostStatusGetV10DataRaiseDetailsStatus {
	return &v
}
