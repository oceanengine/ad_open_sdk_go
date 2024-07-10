/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10FilteringAnalysisType
type QianchuanReportMaterialGetV10FilteringAnalysisType string

// List of qianchuan_report_material_get_v1.0_filtering_analysis_type
const (
	FIRST_PUBLISH_QianchuanReportMaterialGetV10FilteringAnalysisType  QianchuanReportMaterialGetV10FilteringAnalysisType = "first_publish"
	HIGH_QUALITY_QianchuanReportMaterialGetV10FilteringAnalysisType   QianchuanReportMaterialGetV10FilteringAnalysisType = "high_quality"
	IMPROVABLE_QianchuanReportMaterialGetV10FilteringAnalysisType     QianchuanReportMaterialGetV10FilteringAnalysisType = "improvable"
	LOW_EFFICIENCY_QianchuanReportMaterialGetV10FilteringAnalysisType QianchuanReportMaterialGetV10FilteringAnalysisType = "low_efficiency"
	POOR_QUALITY_QianchuanReportMaterialGetV10FilteringAnalysisType   QianchuanReportMaterialGetV10FilteringAnalysisType = "poor_quality"
)

// All allowed values of QianchuanReportMaterialGetV10FilteringAnalysisType enum
var AllowedQianchuanReportMaterialGetV10FilteringAnalysisTypeEnumValues = []QianchuanReportMaterialGetV10FilteringAnalysisType{
	"first_publish",
	"high_quality",
	"improvable",
	"low_efficiency",
	"poor_quality",
}

// NewQianchuanReportMaterialGetV10FilteringAnalysisTypeFromValue returns a pointer to a valid QianchuanReportMaterialGetV10FilteringAnalysisType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10FilteringAnalysisTypeFromValue(v string) (*QianchuanReportMaterialGetV10FilteringAnalysisType, error) {
	ev := QianchuanReportMaterialGetV10FilteringAnalysisType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10FilteringAnalysisType: valid values are %v", v, AllowedQianchuanReportMaterialGetV10FilteringAnalysisTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10FilteringAnalysisType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10FilteringAnalysisTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_filtering_analysis_type value
func (v QianchuanReportMaterialGetV10FilteringAnalysisType) Ptr() *QianchuanReportMaterialGetV10FilteringAnalysisType {
	return &v
}
