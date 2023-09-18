/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportLongTransferOrderGetV10OptimizationTimeScope
type QianchuanReportLongTransferOrderGetV10OptimizationTimeScope int64

// List of qianchuan_report_long_transfer_order_get_v1.0_optimization_time_scope
const (
	Enum_1_QianchuanReportLongTransferOrderGetV10OptimizationTimeScope  QianchuanReportLongTransferOrderGetV10OptimizationTimeScope = 1
	Enum_15_QianchuanReportLongTransferOrderGetV10OptimizationTimeScope QianchuanReportLongTransferOrderGetV10OptimizationTimeScope = 15
	Enum_3_QianchuanReportLongTransferOrderGetV10OptimizationTimeScope  QianchuanReportLongTransferOrderGetV10OptimizationTimeScope = 3
	Enum_30_QianchuanReportLongTransferOrderGetV10OptimizationTimeScope QianchuanReportLongTransferOrderGetV10OptimizationTimeScope = 30
	Enum_7_QianchuanReportLongTransferOrderGetV10OptimizationTimeScope  QianchuanReportLongTransferOrderGetV10OptimizationTimeScope = 7
)

// All allowed values of QianchuanReportLongTransferOrderGetV10OptimizationTimeScope enum
var AllowedQianchuanReportLongTransferOrderGetV10OptimizationTimeScopeEnumValues = []QianchuanReportLongTransferOrderGetV10OptimizationTimeScope{
	1,
	15,
	3,
	30,
	7,
}

// NewQianchuanReportLongTransferOrderGetV10OptimizationTimeScopeFromValue returns a pointer to a valid QianchuanReportLongTransferOrderGetV10OptimizationTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportLongTransferOrderGetV10OptimizationTimeScopeFromValue(v int64) (*QianchuanReportLongTransferOrderGetV10OptimizationTimeScope, error) {
	ev := QianchuanReportLongTransferOrderGetV10OptimizationTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportLongTransferOrderGetV10OptimizationTimeScope: valid values are %v", v, AllowedQianchuanReportLongTransferOrderGetV10OptimizationTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportLongTransferOrderGetV10OptimizationTimeScope) IsValid() bool {
	for _, existing := range AllowedQianchuanReportLongTransferOrderGetV10OptimizationTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_long_transfer_order_get_v1.0_optimization_time_scope value
func (v QianchuanReportLongTransferOrderGetV10OptimizationTimeScope) Ptr() *QianchuanReportLongTransferOrderGetV10OptimizationTimeScope {
	return &v
}