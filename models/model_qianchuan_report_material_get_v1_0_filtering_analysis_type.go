/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

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

// Ptr returns reference to qianchuan_report_material_get_v1.0_filtering_analysis_type value
func (v QianchuanReportMaterialGetV10FilteringAnalysisType) Ptr() *QianchuanReportMaterialGetV10FilteringAnalysisType {
	return &v
}
