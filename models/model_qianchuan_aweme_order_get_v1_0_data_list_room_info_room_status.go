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

// QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus
type QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus string

// List of qianchuan_aweme_order_get_v1.0_data_list_room_info_room_status
const (
	FINISH_QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus  QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus = "FINISH"
	LIVING_QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus  QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus = "LIVING"
	PAUSE_QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus   QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus = "PAUSE"
	PREPARE_QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus = "PREPARE"
	UNKNOW_QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus  QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus = "UNKNOW"
)

// All allowed values of QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus enum
var AllowedQianchuanAwemeOrderGetV10DataListRoomInfoRoomStatusEnumValues = []QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus{
	"FINISH",
	"LIVING",
	"PAUSE",
	"PREPARE",
	"UNKNOW",
}

// NewQianchuanAwemeOrderGetV10DataListRoomInfoRoomStatusFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10DataListRoomInfoRoomStatusFromValue(v string) (*QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus, error) {
	ev := QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10DataListRoomInfoRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10DataListRoomInfoRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_data_list_room_info_room_status value
func (v QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus) Ptr() *QianchuanAwemeOrderGetV10DataListRoomInfoRoomStatus {
	return &v
}
