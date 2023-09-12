/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial
type QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial int64

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_enable_follow_material
const (
	Enum_0_QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial = 0
	Enum_1_QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial = 1
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterialEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterialFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterialFromValue(v int64) (*QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_enable_follow_material value
func (v QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingEnableFollowMaterial {
	return &v
}
