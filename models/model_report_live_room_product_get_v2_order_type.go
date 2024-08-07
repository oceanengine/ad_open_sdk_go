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

// ReportLiveRoomProductGetV2OrderType
type ReportLiveRoomProductGetV2OrderType string

// List of report_live_room_product_get_v2_order_type
const (
	ASC_ReportLiveRoomProductGetV2OrderType  ReportLiveRoomProductGetV2OrderType = "ASC"
	DESC_ReportLiveRoomProductGetV2OrderType ReportLiveRoomProductGetV2OrderType = "DESC"
)

// All allowed values of ReportLiveRoomProductGetV2OrderType enum
var AllowedReportLiveRoomProductGetV2OrderTypeEnumValues = []ReportLiveRoomProductGetV2OrderType{
	"ASC",
	"DESC",
}

// NewReportLiveRoomProductGetV2OrderTypeFromValue returns a pointer to a valid ReportLiveRoomProductGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomProductGetV2OrderTypeFromValue(v string) (*ReportLiveRoomProductGetV2OrderType, error) {
	ev := ReportLiveRoomProductGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomProductGetV2OrderType: valid values are %v", v, AllowedReportLiveRoomProductGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomProductGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomProductGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_product_get_v2_order_type value
func (v ReportLiveRoomProductGetV2OrderType) Ptr() *ReportLiveRoomProductGetV2OrderType {
	return &v
}
