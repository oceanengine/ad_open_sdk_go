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

// ReportRtaGetV2OrderType
type ReportRtaGetV2OrderType string

// List of report_rta_get_v2_order_type
const (
	ASC_ReportRtaGetV2OrderType  ReportRtaGetV2OrderType = "ASC"
	DESC_ReportRtaGetV2OrderType ReportRtaGetV2OrderType = "DESC"
)

// All allowed values of ReportRtaGetV2OrderType enum
var AllowedReportRtaGetV2OrderTypeEnumValues = []ReportRtaGetV2OrderType{
	"ASC",
	"DESC",
}

// NewReportRtaGetV2OrderTypeFromValue returns a pointer to a valid ReportRtaGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRtaGetV2OrderTypeFromValue(v string) (*ReportRtaGetV2OrderType, error) {
	ev := ReportRtaGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRtaGetV2OrderType: valid values are %v", v, AllowedReportRtaGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRtaGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportRtaGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rta_get_v2_order_type value
func (v ReportRtaGetV2OrderType) Ptr() *ReportRtaGetV2OrderType {
	return &v
}
