/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10DeliverySettingLiveScheduleType
type QianchuanAdUpdateV10DeliverySettingLiveScheduleType string

// List of qianchuan_ad_update_v1.0_delivery_setting_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdUpdateV10DeliverySettingLiveScheduleType        QianchuanAdUpdateV10DeliverySettingLiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdUpdateV10DeliverySettingLiveScheduleType       QianchuanAdUpdateV10DeliverySettingLiveScheduleType = "SCHEDULE_START_END"
	SCHEDULE_TIME_FIXEDRANGE_QianchuanAdUpdateV10DeliverySettingLiveScheduleType QianchuanAdUpdateV10DeliverySettingLiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// Ptr returns reference to qianchuan_ad_update_v1.0_delivery_setting_live_schedule_type value
func (v QianchuanAdUpdateV10DeliverySettingLiveScheduleType) Ptr() *QianchuanAdUpdateV10DeliverySettingLiveScheduleType {
	return &v
}
