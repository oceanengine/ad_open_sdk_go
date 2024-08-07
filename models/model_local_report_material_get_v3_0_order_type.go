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

// LocalReportMaterialGetV30OrderType
type LocalReportMaterialGetV30OrderType string

// List of local_report_material_get_v3.0_order_type
const (
	ASC_LocalReportMaterialGetV30OrderType  LocalReportMaterialGetV30OrderType = "ASC"
	DESC_LocalReportMaterialGetV30OrderType LocalReportMaterialGetV30OrderType = "DESC"
)

// All allowed values of LocalReportMaterialGetV30OrderType enum
var AllowedLocalReportMaterialGetV30OrderTypeEnumValues = []LocalReportMaterialGetV30OrderType{
	"ASC",
	"DESC",
}

// NewLocalReportMaterialGetV30OrderTypeFromValue returns a pointer to a valid LocalReportMaterialGetV30OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportMaterialGetV30OrderTypeFromValue(v string) (*LocalReportMaterialGetV30OrderType, error) {
	ev := LocalReportMaterialGetV30OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportMaterialGetV30OrderType: valid values are %v", v, AllowedLocalReportMaterialGetV30OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportMaterialGetV30OrderType) IsValid() bool {
	for _, existing := range AllowedLocalReportMaterialGetV30OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_material_get_v3.0_order_type value
func (v LocalReportMaterialGetV30OrderType) Ptr() *LocalReportMaterialGetV30OrderType {
	return &v
}
