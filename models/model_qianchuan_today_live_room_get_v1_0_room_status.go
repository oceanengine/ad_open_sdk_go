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

// QianchuanTodayLiveRoomGetV10RoomStatus
type QianchuanTodayLiveRoomGetV10RoomStatus string

// List of qianchuan_today_live_room_get_v1.0_room_status
const (
	ALL_QianchuanTodayLiveRoomGetV10RoomStatus    QianchuanTodayLiveRoomGetV10RoomStatus = "ALL"
	FINISH_QianchuanTodayLiveRoomGetV10RoomStatus QianchuanTodayLiveRoomGetV10RoomStatus = "FINISH"
	LIVING_QianchuanTodayLiveRoomGetV10RoomStatus QianchuanTodayLiveRoomGetV10RoomStatus = "LIVING"
)

// All allowed values of QianchuanTodayLiveRoomGetV10RoomStatus enum
var AllowedQianchuanTodayLiveRoomGetV10RoomStatusEnumValues = []QianchuanTodayLiveRoomGetV10RoomStatus{
	"ALL",
	"FINISH",
	"LIVING",
}

// NewQianchuanTodayLiveRoomGetV10RoomStatusFromValue returns a pointer to a valid QianchuanTodayLiveRoomGetV10RoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomGetV10RoomStatusFromValue(v string) (*QianchuanTodayLiveRoomGetV10RoomStatus, error) {
	ev := QianchuanTodayLiveRoomGetV10RoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomGetV10RoomStatus: valid values are %v", v, AllowedQianchuanTodayLiveRoomGetV10RoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomGetV10RoomStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomGetV10RoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_get_v1.0_room_status value
func (v QianchuanTodayLiveRoomGetV10RoomStatus) Ptr() *QianchuanTodayLiveRoomGetV10RoomStatus {
	return &v
}
