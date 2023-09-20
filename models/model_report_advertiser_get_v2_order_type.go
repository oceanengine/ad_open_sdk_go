/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2OrderType
type ReportAdvertiserGetV2OrderType string

// List of report_advertiser_get_v2_order_type
const (
	ASC_ReportAdvertiserGetV2OrderType  ReportAdvertiserGetV2OrderType = "ASC"
	DESC_ReportAdvertiserGetV2OrderType ReportAdvertiserGetV2OrderType = "DESC"
)

// All allowed values of ReportAdvertiserGetV2OrderType enum
var AllowedReportAdvertiserGetV2OrderTypeEnumValues = []ReportAdvertiserGetV2OrderType{
	"ASC",
	"DESC",
}

// NewReportAdvertiserGetV2OrderTypeFromValue returns a pointer to a valid ReportAdvertiserGetV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2OrderTypeFromValue(v string) (*ReportAdvertiserGetV2OrderType, error) {
	ev := ReportAdvertiserGetV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2OrderType: valid values are %v", v, AllowedReportAdvertiserGetV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2OrderType) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_order_type value
func (v ReportAdvertiserGetV2OrderType) Ptr() *ReportAdvertiserGetV2OrderType {
	return &v
}