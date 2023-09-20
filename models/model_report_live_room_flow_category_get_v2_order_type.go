/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportLiveRoomFlowCategoryGetV2OrderType
type ReportLiveRoomFlowCategoryGetV2OrderType string

// List of report_live_room_flow_category_get_v2_order_type
const (
	ASC_ReportLiveRoomFlowCategoryGetV2OrderType  ReportLiveRoomFlowCategoryGetV2OrderType = "ASC"
	DESC_ReportLiveRoomFlowCategoryGetV2OrderType ReportLiveRoomFlowCategoryGetV2OrderType = "DESC"
)

// All allowed values of ReportLiveRoomFlowCategoryGetV2OrderType enum
var AllowedReportLiveRoomFlowCategoryGetV2OrderTypeEnumValues = []ReportLiveRoomFlowCategoryGetV2OrderType{
	"ASC",
	"DESC",
}

// NewReportLiveRoomFlowCategoryGetV2OrderTypeFromValue returns a pointer to a valid ReportLiveRoomFlowCategoryGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportLiveRoomFlowCategoryGetV2OrderTypeFromValue(v string) (*ReportLiveRoomFlowCategoryGetV2OrderType, error) {
	ev := ReportLiveRoomFlowCategoryGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportLiveRoomFlowCategoryGetV2OrderType: valid values are %v", v, AllowedReportLiveRoomFlowCategoryGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportLiveRoomFlowCategoryGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportLiveRoomFlowCategoryGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_live_room_flow_category_get_v2_order_type value
func (v ReportLiveRoomFlowCategoryGetV2OrderType) Ptr() *ReportLiveRoomFlowCategoryGetV2OrderType {
	return &v
}
