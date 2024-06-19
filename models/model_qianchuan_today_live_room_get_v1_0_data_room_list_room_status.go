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

// QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus
type QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus string

// List of qianchuan_today_live_room_get_v1.0_data_room_list_room_status
const (
	ALL_QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus     QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus = "ALL"
	FINISH_QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus  QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus = "FINISH"
	LIVING_QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus  QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus = "LIVING"
	PAUSE_QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus   QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus = "PAUSE"
	PREPARE_QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus = "PREPARE"
	UNKNOWN_QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus = "UNKNOWN"
)

// All allowed values of QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus enum
var AllowedQianchuanTodayLiveRoomGetV10DataRoomListRoomStatusEnumValues = []QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus{
	"ALL",
	"FINISH",
	"LIVING",
	"PAUSE",
	"PREPARE",
	"UNKNOWN",
}

// NewQianchuanTodayLiveRoomGetV10DataRoomListRoomStatusFromValue returns a pointer to a valid QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanTodayLiveRoomGetV10DataRoomListRoomStatusFromValue(v string) (*QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus, error) {
	ev := QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus: valid values are %v", v, AllowedQianchuanTodayLiveRoomGetV10DataRoomListRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanTodayLiveRoomGetV10DataRoomListRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_today_live_room_get_v1.0_data_room_list_room_status value
func (v QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus) Ptr() *QianchuanTodayLiveRoomGetV10DataRoomListRoomStatus {
	return &v
}
