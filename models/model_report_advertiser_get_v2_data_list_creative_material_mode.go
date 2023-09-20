/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2DataListCreativeMaterialMode
type ReportAdvertiserGetV2DataListCreativeMaterialMode string

// List of report_advertiser_get_v2_data_list_creative_material_mode
const (
	STATIC_ASSEMBLE_ReportAdvertiserGetV2DataListCreativeMaterialMode ReportAdvertiserGetV2DataListCreativeMaterialMode = "STATIC_ASSEMBLE"
	CTR_ReportAdvertiserGetV2DataListCreativeMaterialMode             ReportAdvertiserGetV2DataListCreativeMaterialMode = "CTR"
	INTERVENE_ReportAdvertiserGetV2DataListCreativeMaterialMode       ReportAdvertiserGetV2DataListCreativeMaterialMode = "INTERVENE"
)

// All allowed values of ReportAdvertiserGetV2DataListCreativeMaterialMode enum
var AllowedReportAdvertiserGetV2DataListCreativeMaterialModeEnumValues = []ReportAdvertiserGetV2DataListCreativeMaterialMode{
	"STATIC_ASSEMBLE",
	"CTR",
	"INTERVENE",
}

// NewReportAdvertiserGetV2DataListCreativeMaterialModeFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListCreativeMaterialModeFromValue(v string) (*ReportAdvertiserGetV2DataListCreativeMaterialMode, error) {
	ev := ReportAdvertiserGetV2DataListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListCreativeMaterialMode: valid values are %v", v, AllowedReportAdvertiserGetV2DataListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_creative_material_mode value
func (v ReportAdvertiserGetV2DataListCreativeMaterialMode) Ptr() *ReportAdvertiserGetV2DataListCreativeMaterialMode {
	return &v
}
