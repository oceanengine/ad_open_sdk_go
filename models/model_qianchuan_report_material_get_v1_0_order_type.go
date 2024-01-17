/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10OrderType
type QianchuanReportMaterialGetV10OrderType string

// List of qianchuan_report_material_get_v1.0_order_type
const (
	ASC_QianchuanReportMaterialGetV10OrderType  QianchuanReportMaterialGetV10OrderType = "ASC"
	DESC_QianchuanReportMaterialGetV10OrderType QianchuanReportMaterialGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportMaterialGetV10OrderType enum
var AllowedQianchuanReportMaterialGetV10OrderTypeEnumValues = []QianchuanReportMaterialGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportMaterialGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportMaterialGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10OrderTypeFromValue(v string) (*QianchuanReportMaterialGetV10OrderType, error) {
	ev := QianchuanReportMaterialGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10OrderType: valid values are %v", v, AllowedQianchuanReportMaterialGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_order_type value
func (v QianchuanReportMaterialGetV10OrderType) Ptr() *QianchuanReportMaterialGetV10OrderType {
	return &v
}
