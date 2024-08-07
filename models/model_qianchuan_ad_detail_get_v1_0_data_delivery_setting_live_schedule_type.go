/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType
type QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType        QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType       QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType = "SCHEDULE_START_END"
	SCHEDULE_TIME_FIXEDRANGE_QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingLiveScheduleTypeEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
	"SCHEDULE_TIME_FIXEDRANGE",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingLiveScheduleTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingLiveScheduleTypeFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingLiveScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingLiveScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_live_schedule_type value
func (v QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingLiveScheduleType {
	return &v
}
