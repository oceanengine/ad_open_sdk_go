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

// QianchuanAdDetailGetV10DataRoomInfoRoomStatus
type QianchuanAdDetailGetV10DataRoomInfoRoomStatus string

// List of qianchuan_ad_detail_get_v1.0_data_room_info_room_status
const (
	ALL_QianchuanAdDetailGetV10DataRoomInfoRoomStatus     QianchuanAdDetailGetV10DataRoomInfoRoomStatus = "ALL"
	FINISH_QianchuanAdDetailGetV10DataRoomInfoRoomStatus  QianchuanAdDetailGetV10DataRoomInfoRoomStatus = "FINISH"
	LIVING_QianchuanAdDetailGetV10DataRoomInfoRoomStatus  QianchuanAdDetailGetV10DataRoomInfoRoomStatus = "LIVING"
	PAUSE_QianchuanAdDetailGetV10DataRoomInfoRoomStatus   QianchuanAdDetailGetV10DataRoomInfoRoomStatus = "PAUSE"
	PREPARE_QianchuanAdDetailGetV10DataRoomInfoRoomStatus QianchuanAdDetailGetV10DataRoomInfoRoomStatus = "PREPARE"
	UNKNOWN_QianchuanAdDetailGetV10DataRoomInfoRoomStatus QianchuanAdDetailGetV10DataRoomInfoRoomStatus = "UNKNOWN"
)

// All allowed values of QianchuanAdDetailGetV10DataRoomInfoRoomStatus enum
var AllowedQianchuanAdDetailGetV10DataRoomInfoRoomStatusEnumValues = []QianchuanAdDetailGetV10DataRoomInfoRoomStatus{
	"ALL",
	"FINISH",
	"LIVING",
	"PAUSE",
	"PREPARE",
	"UNKNOWN",
}

// NewQianchuanAdDetailGetV10DataRoomInfoRoomStatusFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataRoomInfoRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataRoomInfoRoomStatusFromValue(v string) (*QianchuanAdDetailGetV10DataRoomInfoRoomStatus, error) {
	ev := QianchuanAdDetailGetV10DataRoomInfoRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataRoomInfoRoomStatus: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataRoomInfoRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataRoomInfoRoomStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataRoomInfoRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_room_info_room_status value
func (v QianchuanAdDetailGetV10DataRoomInfoRoomStatus) Ptr() *QianchuanAdDetailGetV10DataRoomInfoRoomStatus {
	return &v
}
