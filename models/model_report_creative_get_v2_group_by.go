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

// ReportCreativeGetV2GroupBy
type ReportCreativeGetV2GroupBy string

// List of report_creative_get_v2_group_by
const (
	STAT_GROUP_BY_PRICING_CATEGORY_ReportCreativeGetV2GroupBy       ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_PRICING_CATEGORY"
	STAT_GROUP_BY_CREATIVE_MATERIAL_MODE_ReportCreativeGetV2GroupBy ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_MATERIAL_MODE"
	STAT_GROUP_BY_FIELD_STAT_TIME_ReportCreativeGetV2GroupBy        ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_FIELD_STAT_TIME"
	STAT_GROUP_BY_FIELD_ID_ReportCreativeGetV2GroupBy               ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_FIELD_ID"
	STAT_GROUP_BY_EXTERNAL_ACTION_ReportCreativeGetV2GroupBy        ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_EXTERNAL_ACTION"
	STAT_GROUP_BY_LANDING_TYPE_ReportCreativeGetV2GroupBy           ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_LANDING_TYPE"
	STAT_GROUP_BY_MATERIAL_ID_ReportCreativeGetV2GroupBy            ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_MATERIAL_ID"
	STAT_GROUP_BY_PRICING_ReportCreativeGetV2GroupBy                ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_PRICING"
	STAT_GROUP_BY_INVENTORY_ReportCreativeGetV2GroupBy              ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_INVENTORY"
	STAT_GROUP_BY_CAMPAIGN_TYPE_ReportCreativeGetV2GroupBy          ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_CAMPAIGN_TYPE"
	STAT_GROUP_BY_IMAGE_MODE_ReportCreativeGetV2GroupBy             ReportCreativeGetV2GroupBy = "STAT_GROUP_BY_IMAGE_MODE"
)

// All allowed values of ReportCreativeGetV2GroupBy enum
var AllowedReportCreativeGetV2GroupByEnumValues = []ReportCreativeGetV2GroupBy{
	"STAT_GROUP_BY_PRICING_CATEGORY",
	"STAT_GROUP_BY_CREATIVE_MATERIAL_MODE",
	"STAT_GROUP_BY_FIELD_STAT_TIME",
	"STAT_GROUP_BY_FIELD_ID",
	"STAT_GROUP_BY_EXTERNAL_ACTION",
	"STAT_GROUP_BY_LANDING_TYPE",
	"STAT_GROUP_BY_MATERIAL_ID",
	"STAT_GROUP_BY_PRICING",
	"STAT_GROUP_BY_INVENTORY",
	"STAT_GROUP_BY_CAMPAIGN_TYPE",
	"STAT_GROUP_BY_IMAGE_MODE",
}

// NewReportCreativeGetV2GroupByFromValue returns a pointer to a valid ReportCreativeGetV2GroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2GroupByFromValue(v string) (*ReportCreativeGetV2GroupBy, error) {
	ev := ReportCreativeGetV2GroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2GroupBy: valid values are %v", v, AllowedReportCreativeGetV2GroupByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2GroupBy) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2GroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_group_by value
func (v ReportCreativeGetV2GroupBy) Ptr() *ReportCreativeGetV2GroupBy {
	return &v
}
