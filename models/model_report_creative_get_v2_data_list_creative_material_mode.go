/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2DataListCreativeMaterialMode
type ReportCreativeGetV2DataListCreativeMaterialMode string

// List of report_creative_get_v2_data_list_creative_material_mode
const (
	STATIC_ASSEMBLE_ReportCreativeGetV2DataListCreativeMaterialMode ReportCreativeGetV2DataListCreativeMaterialMode = "STATIC_ASSEMBLE"
	CTR_ReportCreativeGetV2DataListCreativeMaterialMode             ReportCreativeGetV2DataListCreativeMaterialMode = "CTR"
	INTERVENE_ReportCreativeGetV2DataListCreativeMaterialMode       ReportCreativeGetV2DataListCreativeMaterialMode = "INTERVENE"
)

// Ptr returns reference to report_creative_get_v2_data_list_creative_material_mode value
func (v ReportCreativeGetV2DataListCreativeMaterialMode) Ptr() *ReportCreativeGetV2DataListCreativeMaterialMode {
	return &v
}
