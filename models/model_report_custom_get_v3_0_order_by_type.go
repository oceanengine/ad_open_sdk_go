/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCustomGetV30OrderByType
type ReportCustomGetV30OrderByType string

// List of report_custom_get_v3.0_order_by_type
const (
	ASC_ReportCustomGetV30OrderByType  ReportCustomGetV30OrderByType = "ASC"
	DESC_ReportCustomGetV30OrderByType ReportCustomGetV30OrderByType = "DESC"
)

// All allowed values of ReportCustomGetV30OrderByType enum
var AllowedReportCustomGetV30OrderByTypeEnumValues = []ReportCustomGetV30OrderByType{
	"ASC",
	"DESC",
}

// NewReportCustomGetV30OrderByTypeFromValue returns a pointer to a valid ReportCustomGetV30OrderByType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomGetV30OrderByTypeFromValue(v string) (*ReportCustomGetV30OrderByType, error) {
	ev := ReportCustomGetV30OrderByType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomGetV30OrderByType: valid values are %v", v, AllowedReportCustomGetV30OrderByTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomGetV30OrderByType) IsValid() bool {
	for _, existing := range AllowedReportCustomGetV30OrderByTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_get_v3.0_order_by_type value
func (v ReportCustomGetV30OrderByType) Ptr() *ReportCustomGetV30OrderByType {
	return &v
}
