/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2FilteringCreativeMaterialModes
type ReportAdGetV2FilteringCreativeMaterialModes string

// List of report_ad_get_v2_filtering_creative_material_modes
const (
	INTERVENE_ReportAdGetV2FilteringCreativeMaterialModes       ReportAdGetV2FilteringCreativeMaterialModes = "INTERVENE"
	STATIC_ASSEMBLE_ReportAdGetV2FilteringCreativeMaterialModes ReportAdGetV2FilteringCreativeMaterialModes = "STATIC_ASSEMBLE"
	CTR_ReportAdGetV2FilteringCreativeMaterialModes             ReportAdGetV2FilteringCreativeMaterialModes = "CTR"
)

// All allowed values of ReportAdGetV2FilteringCreativeMaterialModes enum
var AllowedReportAdGetV2FilteringCreativeMaterialModesEnumValues = []ReportAdGetV2FilteringCreativeMaterialModes{
	"INTERVENE",
	"STATIC_ASSEMBLE",
	"CTR",
}

// NewReportAdGetV2FilteringCreativeMaterialModesFromValue returns a pointer to a valid ReportAdGetV2FilteringCreativeMaterialModes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringCreativeMaterialModesFromValue(v string) (*ReportAdGetV2FilteringCreativeMaterialModes, error) {
	ev := ReportAdGetV2FilteringCreativeMaterialModes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringCreativeMaterialModes: valid values are %v", v, AllowedReportAdGetV2FilteringCreativeMaterialModesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringCreativeMaterialModes) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringCreativeMaterialModesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_creative_material_modes value
func (v ReportAdGetV2FilteringCreativeMaterialModes) Ptr() *ReportAdGetV2FilteringCreativeMaterialModes {
	return &v
}
