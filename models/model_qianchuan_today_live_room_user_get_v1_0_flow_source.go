/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanTodayLiveRoomUserGetV10FlowSource
type QianchuanTodayLiveRoomUserGetV10FlowSource string

// List of qianchuan_today_live_room_user_get_v1.0_flow_source
const (
	ALL_QianchuanTodayLiveRoomUserGetV10FlowSource QianchuanTodayLiveRoomUserGetV10FlowSource = "ALL"
	PC_QianchuanTodayLiveRoomUserGetV10FlowSource  QianchuanTodayLiveRoomUserGetV10FlowSource = "PC"
)

// All allowed values of QianchuanTodayLiveRoomUserGetV10FlowSource enum
var AllowedQianchuanTodayLiveRoomUserGetV10FlowSourceEnumValues = []QianchuanTodayLiveRoomUserGetV10FlowSource{
	"ALL",
	"PC",
}

// NewQianchuanTodayLiveRoomUserGetV10FlowSourceFromValue returns a pointer to a valid QianchuanTodayLiveRoomUserGetV10FlowSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomUserGetV10FlowSourceFromValue(v string) (*QianchuanTodayLiveRoomUserGetV10FlowSource, error) {
	ev := QianchuanTodayLiveRoomUserGetV10FlowSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomUserGetV10FlowSource: valid values are %v", v, AllowedQianchuanTodayLiveRoomUserGetV10FlowSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomUserGetV10FlowSource) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomUserGetV10FlowSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_user_get_v1.0_flow_source value
func (v QianchuanTodayLiveRoomUserGetV10FlowSource) Ptr() *QianchuanTodayLiveRoomUserGetV10FlowSource {
	return &v
}
