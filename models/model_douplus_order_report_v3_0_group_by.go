/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderReportV30GroupBy
type DouplusOrderReportV30GroupBy string

// List of douplus_order_report_v3.0_group_by
const (
	GROUP_BY_AD_ID_DouplusOrderReportV30GroupBy       DouplusOrderReportV30GroupBy = "GROUP_BY_AD_ID"
	GROUP_BY_CREATIVE_ID_DouplusOrderReportV30GroupBy DouplusOrderReportV30GroupBy = "GROUP_BY_CREATIVE_ID"
)

// All allowed values of DouplusOrderReportV30GroupBy enum
var AllowedDouplusOrderReportV30GroupByEnumValues = []DouplusOrderReportV30GroupBy{
	"GROUP_BY_AD_ID",
	"GROUP_BY_CREATIVE_ID",
}

// NewDouplusOrderReportV30GroupByFromValue returns a pointer to a valid DouplusOrderReportV30GroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderReportV30GroupByFromValue(v string) (*DouplusOrderReportV30GroupBy, error) {
	ev := DouplusOrderReportV30GroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderReportV30GroupBy: valid values are %v", v, AllowedDouplusOrderReportV30GroupByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderReportV30GroupBy) IsValid() bool {
	for _, existing := range AllowedDouplusOrderReportV30GroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_report_v3.0_group_by value
func (v DouplusOrderReportV30GroupBy) Ptr() *DouplusOrderReportV30GroupBy {
	return &v
}
