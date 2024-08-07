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

// ReportStardeliveryTaskVideoDataGetV30OrderType
type ReportStardeliveryTaskVideoDataGetV30OrderType string

// List of report_stardelivery_task_video_data_get_v3.0_order_type
const (
	ASC_ReportStardeliveryTaskVideoDataGetV30OrderType  ReportStardeliveryTaskVideoDataGetV30OrderType = "ASC"
	DESC_ReportStardeliveryTaskVideoDataGetV30OrderType ReportStardeliveryTaskVideoDataGetV30OrderType = "DESC"
)

// All allowed values of ReportStardeliveryTaskVideoDataGetV30OrderType enum
var AllowedReportStardeliveryTaskVideoDataGetV30OrderTypeEnumValues = []ReportStardeliveryTaskVideoDataGetV30OrderType{
	"ASC",
	"DESC",
}

// NewReportStardeliveryTaskVideoDataGetV30OrderTypeFromValue returns a pointer to a valid ReportStardeliveryTaskVideoDataGetV30OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportStardeliveryTaskVideoDataGetV30OrderTypeFromValue(v string) (*ReportStardeliveryTaskVideoDataGetV30OrderType, error) {
	ev := ReportStardeliveryTaskVideoDataGetV30OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportStardeliveryTaskVideoDataGetV30OrderType: valid values are %v", v, AllowedReportStardeliveryTaskVideoDataGetV30OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportStardeliveryTaskVideoDataGetV30OrderType) IsValid() bool {
	for _, existing := range AllowedReportStardeliveryTaskVideoDataGetV30OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_stardelivery_task_video_data_get_v3.0_order_type value
func (v ReportStardeliveryTaskVideoDataGetV30OrderType) Ptr() *ReportStardeliveryTaskVideoDataGetV30OrderType {
	return &v
}
