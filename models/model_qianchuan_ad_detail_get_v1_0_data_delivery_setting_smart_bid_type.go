/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10DataDeliverySettingSmartBidType
type QianchuanAdDetailGetV10DataDeliverySettingSmartBidType string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanAdDetailGetV10DataDeliverySettingSmartBidType QianchuanAdDetailGetV10DataDeliverySettingSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanAdDetailGetV10DataDeliverySettingSmartBidType       QianchuanAdDetailGetV10DataDeliverySettingSmartBidType = "SMART_BID_CUSTOM"
)

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_smart_bid_type value
func (v QianchuanAdDetailGetV10DataDeliverySettingSmartBidType) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingSmartBidType {
	return &v
}
