/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10DeliverySettingVideoScheduleType
type QianchuanAdCreateV10DeliverySettingVideoScheduleType string

// List of qianchuan_ad_create_v1.0_delivery_setting_video_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdCreateV10DeliverySettingVideoScheduleType  QianchuanAdCreateV10DeliverySettingVideoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdCreateV10DeliverySettingVideoScheduleType QianchuanAdCreateV10DeliverySettingVideoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingVideoScheduleType enum
var AllowedQianchuanAdCreateV10DeliverySettingVideoScheduleTypeEnumValues = []QianchuanAdCreateV10DeliverySettingVideoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewQianchuanAdCreateV10DeliverySettingVideoScheduleTypeFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingVideoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingVideoScheduleTypeFromValue(v string) (*QianchuanAdCreateV10DeliverySettingVideoScheduleType, error) {
	ev := QianchuanAdCreateV10DeliverySettingVideoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingVideoScheduleType: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingVideoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingVideoScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingVideoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_video_schedule_type value
func (v QianchuanAdCreateV10DeliverySettingVideoScheduleType) Ptr() *QianchuanAdCreateV10DeliverySettingVideoScheduleType {
	return &v
}
