/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringCreativeMaterialModes
type ReportAdGetV2FilteringCreativeMaterialModes string

// List of report_ad_get_v2_filtering_creative_material_modes
const (
	INTERVENE_ReportAdGetV2FilteringCreativeMaterialModes       ReportAdGetV2FilteringCreativeMaterialModes = "INTERVENE"
	STATIC_ASSEMBLE_ReportAdGetV2FilteringCreativeMaterialModes ReportAdGetV2FilteringCreativeMaterialModes = "STATIC_ASSEMBLE"
	CTR_ReportAdGetV2FilteringCreativeMaterialModes             ReportAdGetV2FilteringCreativeMaterialModes = "CTR"
)

// Ptr returns reference to report_ad_get_v2_filtering_creative_material_modes value
func (v ReportAdGetV2FilteringCreativeMaterialModes) Ptr() *ReportAdGetV2FilteringCreativeMaterialModes {
	return &v
}
