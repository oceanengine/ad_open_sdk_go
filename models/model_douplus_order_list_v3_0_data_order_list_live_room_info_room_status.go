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

// DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus
type DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus string

// List of douplus_order_list_v3.0_data_order_list_live_room_info_room_status
const (
	FINISH_DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus  DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus = "FINISH"
	LIVING_DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus  DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus = "LIVING"
	PAUSE_DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus   DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus = "PAUSE"
	PREPARE_DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus = "PREPARE"
)

// All allowed values of DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus enum
var AllowedDouplusOrderListV30DataOrderListLiveRoomInfoRoomStatusEnumValues = []DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus{
	"FINISH",
	"LIVING",
	"PAUSE",
	"PREPARE",
}

// NewDouplusOrderListV30DataOrderListLiveRoomInfoRoomStatusFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListLiveRoomInfoRoomStatusFromValue(v string) (*DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus, error) {
	ev := DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListLiveRoomInfoRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListLiveRoomInfoRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_live_room_info_room_status value
func (v DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus) Ptr() *DouplusOrderListV30DataOrderListLiveRoomInfoRoomStatus {
	return &v
}
