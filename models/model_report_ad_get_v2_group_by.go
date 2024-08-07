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

// ReportAdGetV2GroupBy
type ReportAdGetV2GroupBy string

// List of report_ad_get_v2_group_by
const (
	STAT_GROUP_BY_IMAGE_MODE_ReportAdGetV2GroupBy             ReportAdGetV2GroupBy = "STAT_GROUP_BY_IMAGE_MODE"
	STAT_GROUP_BY_AC_ReportAdGetV2GroupBy                     ReportAdGetV2GroupBy = "STAT_GROUP_BY_AC"
	STAT_GROUP_BY_PRICING_ReportAdGetV2GroupBy                ReportAdGetV2GroupBy = "STAT_GROUP_BY_PRICING"
	STAT_GROUP_BY_FIELD_STAT_TIME_ReportAdGetV2GroupBy        ReportAdGetV2GroupBy = "STAT_GROUP_BY_FIELD_STAT_TIME"
	STAT_GROUP_BY_CREATIVE_MATERIAL_MODE_ReportAdGetV2GroupBy ReportAdGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_MATERIAL_MODE"
	STAT_GROUP_BY_FIELD_ID_ReportAdGetV2GroupBy               ReportAdGetV2GroupBy = "STAT_GROUP_BY_FIELD_ID"
	STAT_GROUP_BY_MATERIAL_ID_ReportAdGetV2GroupBy            ReportAdGetV2GroupBy = "STAT_GROUP_BY_MATERIAL_ID"
	STAT_GROUP_BY_GENDER_ReportAdGetV2GroupBy                 ReportAdGetV2GroupBy = "STAT_GROUP_BY_GENDER"
	STAT_GROUP_BY_EXTERNAL_ACTION_ReportAdGetV2GroupBy        ReportAdGetV2GroupBy = "STAT_GROUP_BY_EXTERNAL_ACTION"
	STAT_GROUP_BY_CREATIVE_COMPONENT_ID_ReportAdGetV2GroupBy  ReportAdGetV2GroupBy = "STAT_GROUP_BY_CREATIVE_COMPONENT_ID"
	STAT_GROUP_BY_AGE_ReportAdGetV2GroupBy                    ReportAdGetV2GroupBy = "STAT_GROUP_BY_AGE"
	STAT_GROUP_BY_PRICING_CATEGORY_ReportAdGetV2GroupBy       ReportAdGetV2GroupBy = "STAT_GROUP_BY_PRICING_CATEGORY"
	STAT_GROUP_BY_PROVINCE_NAME_ReportAdGetV2GroupBy          ReportAdGetV2GroupBy = "STAT_GROUP_BY_PROVINCE_NAME"
	STAT_GROUP_BY_CAMPAIGN_TYPE_ReportAdGetV2GroupBy          ReportAdGetV2GroupBy = "STAT_GROUP_BY_CAMPAIGN_TYPE"
	STAT_GROUP_BY_PLATFORM_ReportAdGetV2GroupBy               ReportAdGetV2GroupBy = "STAT_GROUP_BY_PLATFORM"
	STAT_GROUP_BY_LANDING_TYPE_ReportAdGetV2GroupBy           ReportAdGetV2GroupBy = "STAT_GROUP_BY_LANDING_TYPE"
	STAT_GROUP_BY_INVENTORY_ReportAdGetV2GroupBy              ReportAdGetV2GroupBy = "STAT_GROUP_BY_INVENTORY"
	STAT_GROUP_BY_CITY_NAME_ReportAdGetV2GroupBy              ReportAdGetV2GroupBy = "STAT_GROUP_BY_CITY_NAME"
)

// All allowed values of ReportAdGetV2GroupBy enum
var AllowedReportAdGetV2GroupByEnumValues = []ReportAdGetV2GroupBy{
	"STAT_GROUP_BY_IMAGE_MODE",
	"STAT_GROUP_BY_AC",
	"STAT_GROUP_BY_PRICING",
	"STAT_GROUP_BY_FIELD_STAT_TIME",
	"STAT_GROUP_BY_CREATIVE_MATERIAL_MODE",
	"STAT_GROUP_BY_FIELD_ID",
	"STAT_GROUP_BY_MATERIAL_ID",
	"STAT_GROUP_BY_GENDER",
	"STAT_GROUP_BY_EXTERNAL_ACTION",
	"STAT_GROUP_BY_CREATIVE_COMPONENT_ID",
	"STAT_GROUP_BY_AGE",
	"STAT_GROUP_BY_PRICING_CATEGORY",
	"STAT_GROUP_BY_PROVINCE_NAME",
	"STAT_GROUP_BY_CAMPAIGN_TYPE",
	"STAT_GROUP_BY_PLATFORM",
	"STAT_GROUP_BY_LANDING_TYPE",
	"STAT_GROUP_BY_INVENTORY",
	"STAT_GROUP_BY_CITY_NAME",
}

// NewReportAdGetV2GroupByFromValue returns a pointer to a valid ReportAdGetV2GroupBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2GroupByFromValue(v string) (*ReportAdGetV2GroupBy, error) {
	ev := ReportAdGetV2GroupBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2GroupBy: valid values are %v", v, AllowedReportAdGetV2GroupByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2GroupBy) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2GroupByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_group_by value
func (v ReportAdGetV2GroupBy) Ptr() *ReportAdGetV2GroupBy {
	return &v
}
