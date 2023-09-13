/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode
type QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode string

// List of qianchuan_aweme_order_get_v1.0_data_list_delivery_setting_liveroom_heat_mode
const (
	BY_ROOM_QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode  QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode = "BY_ROOM"
	BY_VIDEO_QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode = "BY_VIDEO"
)

// All allowed values of QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode enum
var AllowedQianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatModeEnumValues = []QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode{
	"BY_ROOM",
	"BY_VIDEO",
}

// NewQianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatModeFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatModeFromValue(v string) (*QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode, error) {
	ev := QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_data_list_delivery_setting_liveroom_heat_mode value
func (v QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode) Ptr() *QianchuanAwemeOrderGetV10DataListDeliverySettingLiveroomHeatMode {
	return &v
}
