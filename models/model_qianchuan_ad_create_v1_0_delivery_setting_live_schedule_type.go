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

// QianchuanAdCreateV10DeliverySettingLiveScheduleType
type QianchuanAdCreateV10DeliverySettingLiveScheduleType string

// List of qianchuan_ad_create_v1.0_delivery_setting_live_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdCreateV10DeliverySettingLiveScheduleType        QianchuanAdCreateV10DeliverySettingLiveScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdCreateV10DeliverySettingLiveScheduleType       QianchuanAdCreateV10DeliverySettingLiveScheduleType = "SCHEDULE_START_END"
	SCHEDULE_TIME_FIXEDRANGE_QianchuanAdCreateV10DeliverySettingLiveScheduleType QianchuanAdCreateV10DeliverySettingLiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingLiveScheduleType enum
var AllowedQianchuanAdCreateV10DeliverySettingLiveScheduleTypeEnumValues = []QianchuanAdCreateV10DeliverySettingLiveScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
	"SCHEDULE_TIME_FIXEDRANGE",
}

// NewQianchuanAdCreateV10DeliverySettingLiveScheduleTypeFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingLiveScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingLiveScheduleTypeFromValue(v string) (*QianchuanAdCreateV10DeliverySettingLiveScheduleType, error) {
	ev := QianchuanAdCreateV10DeliverySettingLiveScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingLiveScheduleType: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingLiveScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingLiveScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingLiveScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_live_schedule_type value
func (v QianchuanAdCreateV10DeliverySettingLiveScheduleType) Ptr() *QianchuanAdCreateV10DeliverySettingLiveScheduleType {
	return &v
}
