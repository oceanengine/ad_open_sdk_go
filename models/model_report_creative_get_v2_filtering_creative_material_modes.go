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

// ReportCreativeGetV2FilteringCreativeMaterialModes
type ReportCreativeGetV2FilteringCreativeMaterialModes string

// List of report_creative_get_v2_filtering_creative_material_modes
const (
	STATIC_ASSEMBLE_ReportCreativeGetV2FilteringCreativeMaterialModes ReportCreativeGetV2FilteringCreativeMaterialModes = "STATIC_ASSEMBLE"
	CTR_ReportCreativeGetV2FilteringCreativeMaterialModes             ReportCreativeGetV2FilteringCreativeMaterialModes = "CTR"
	INTERVENE_ReportCreativeGetV2FilteringCreativeMaterialModes       ReportCreativeGetV2FilteringCreativeMaterialModes = "INTERVENE"
)

// All allowed values of ReportCreativeGetV2FilteringCreativeMaterialModes enum
var AllowedReportCreativeGetV2FilteringCreativeMaterialModesEnumValues = []ReportCreativeGetV2FilteringCreativeMaterialModes{
	"STATIC_ASSEMBLE",
	"CTR",
	"INTERVENE",
}

// NewReportCreativeGetV2FilteringCreativeMaterialModesFromValue returns a pointer to a valid ReportCreativeGetV2FilteringCreativeMaterialModes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringCreativeMaterialModesFromValue(v string) (*ReportCreativeGetV2FilteringCreativeMaterialModes, error) {
	ev := ReportCreativeGetV2FilteringCreativeMaterialModes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringCreativeMaterialModes: valid values are %v", v, AllowedReportCreativeGetV2FilteringCreativeMaterialModesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringCreativeMaterialModes) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringCreativeMaterialModesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_creative_material_modes value
func (v ReportCreativeGetV2FilteringCreativeMaterialModes) Ptr() *ReportCreativeGetV2FilteringCreativeMaterialModes {
	return &v
}
