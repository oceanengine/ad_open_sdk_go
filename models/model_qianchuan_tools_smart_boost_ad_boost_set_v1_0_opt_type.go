/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostSetV10OptType
type QianchuanToolsSmartBoostAdBoostSetV10OptType string

// List of qianchuan_tools_smart_boost_ad_boost_set_v1.0_opt_type
const (
	CANCEL_SUBSCRIBE_QianchuanToolsSmartBoostAdBoostSetV10OptType QianchuanToolsSmartBoostAdBoostSetV10OptType = "CANCEL_SUBSCRIBE"
	START_RAISE_QianchuanToolsSmartBoostAdBoostSetV10OptType      QianchuanToolsSmartBoostAdBoostSetV10OptType = "START_RAISE"
	STOP_RAISE_QianchuanToolsSmartBoostAdBoostSetV10OptType       QianchuanToolsSmartBoostAdBoostSetV10OptType = "STOP_RAISE"
	SUBSCRIBE_RAISE_QianchuanToolsSmartBoostAdBoostSetV10OptType  QianchuanToolsSmartBoostAdBoostSetV10OptType = "SUBSCRIBE_RAISE"
)

// Ptr returns reference to qianchuan_tools_smart_boost_ad_boost_set_v1.0_opt_type value
func (v QianchuanToolsSmartBoostAdBoostSetV10OptType) Ptr() *QianchuanToolsSmartBoostAdBoostSetV10OptType {
	return &v
}
