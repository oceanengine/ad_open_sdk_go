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

// QianchuanAdUpdateV10DeliverySettingVideoScheduleType
type QianchuanAdUpdateV10DeliverySettingVideoScheduleType string

// List of qianchuan_ad_update_v1.0_delivery_setting_video_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdUpdateV10DeliverySettingVideoScheduleType  QianchuanAdUpdateV10DeliverySettingVideoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdUpdateV10DeliverySettingVideoScheduleType QianchuanAdUpdateV10DeliverySettingVideoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of QianchuanAdUpdateV10DeliverySettingVideoScheduleType enum
var AllowedQianchuanAdUpdateV10DeliverySettingVideoScheduleTypeEnumValues = []QianchuanAdUpdateV10DeliverySettingVideoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewQianchuanAdUpdateV10DeliverySettingVideoScheduleTypeFromValue returns a pointer to a valid QianchuanAdUpdateV10DeliverySettingVideoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10DeliverySettingVideoScheduleTypeFromValue(v string) (*QianchuanAdUpdateV10DeliverySettingVideoScheduleType, error) {
	ev := QianchuanAdUpdateV10DeliverySettingVideoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10DeliverySettingVideoScheduleType: valid values are %v", v, AllowedQianchuanAdUpdateV10DeliverySettingVideoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10DeliverySettingVideoScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10DeliverySettingVideoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_delivery_setting_video_schedule_type value
func (v QianchuanAdUpdateV10DeliverySettingVideoScheduleType) Ptr() *QianchuanAdUpdateV10DeliverySettingVideoScheduleType {
	return &v
}
