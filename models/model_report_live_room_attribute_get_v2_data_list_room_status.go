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

// ReportLiveRoomAttributeGetV2DataListRoomStatus
type ReportLiveRoomAttributeGetV2DataListRoomStatus string

// List of report_live_room_attribute_get_v2_data_list_room_status
const (
	LIVING_ReportLiveRoomAttributeGetV2DataListRoomStatus    ReportLiveRoomAttributeGetV2DataListRoomStatus = "LIVING"
	PAUSE_ReportLiveRoomAttributeGetV2DataListRoomStatus     ReportLiveRoomAttributeGetV2DataListRoomStatus = "PAUSE"
	END_ReportLiveRoomAttributeGetV2DataListRoomStatus       ReportLiveRoomAttributeGetV2DataListRoomStatus = "END"
	PREPARING_ReportLiveRoomAttributeGetV2DataListRoomStatus ReportLiveRoomAttributeGetV2DataListRoomStatus = "PREPARING"
)

// All allowed values of ReportLiveRoomAttributeGetV2DataListRoomStatus enum
var AllowedReportLiveRoomAttributeGetV2DataListRoomStatusEnumValues = []ReportLiveRoomAttributeGetV2DataListRoomStatus{
	"LIVING",
	"PAUSE",
	"END",
	"PREPARING",
}

// NewReportLiveRoomAttributeGetV2DataListRoomStatusFromValue returns a pointer to a valid ReportLiveRoomAttributeGetV2DataListRoomStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomAttributeGetV2DataListRoomStatusFromValue(v string) (*ReportLiveRoomAttributeGetV2DataListRoomStatus, error) {
	ev := ReportLiveRoomAttributeGetV2DataListRoomStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomAttributeGetV2DataListRoomStatus: valid values are %v", v, AllowedReportLiveRoomAttributeGetV2DataListRoomStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomAttributeGetV2DataListRoomStatus) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomAttributeGetV2DataListRoomStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_attribute_get_v2_data_list_room_status value
func (v ReportLiveRoomAttributeGetV2DataListRoomStatus) Ptr() *ReportLiveRoomAttributeGetV2DataListRoomStatus {
	return &v
}
