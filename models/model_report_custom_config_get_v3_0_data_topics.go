/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCustomConfigGetV30DataTopics
type ReportCustomConfigGetV30DataTopics string

// List of report_custom_config_get_v3.0_data_topics
const (
	BASIC_DATA_ReportCustomConfigGetV30DataTopics           ReportCustomConfigGetV30DataTopics = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomConfigGetV30DataTopics         ReportCustomConfigGetV30DataTopics = "BIDWORD_DATA"
	CREATIVE_DATA_ReportCustomConfigGetV30DataTopics        ReportCustomConfigGetV30DataTopics = "CREATIVE_DATA"
	DMP_DATA_ReportCustomConfigGetV30DataTopics             ReportCustomConfigGetV30DataTopics = "DMP_DATA"
	MATERIAL_DATA_ReportCustomConfigGetV30DataTopics        ReportCustomConfigGetV30DataTopics = "MATERIAL_DATA"
	ONE_KEY_BOOST_DATA_ReportCustomConfigGetV30DataTopics   ReportCustomConfigGetV30DataTopics = "ONE_KEY_BOOST_DATA"
	PRODUCT_DATA_ReportCustomConfigGetV30DataTopics         ReportCustomConfigGetV30DataTopics = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomConfigGetV30DataTopics           ReportCustomConfigGetV30DataTopics = "QUERY_DATA"
	VIDEO_DUARATION_DATA_ReportCustomConfigGetV30DataTopics ReportCustomConfigGetV30DataTopics = "VIDEO_DUARATION_DATA"
)

// All allowed values of ReportCustomConfigGetV30DataTopics enum
var AllowedReportCustomConfigGetV30DataTopicsEnumValues = []ReportCustomConfigGetV30DataTopics{
	"BASIC_DATA",
	"BIDWORD_DATA",
	"CREATIVE_DATA",
	"DMP_DATA",
	"MATERIAL_DATA",
	"ONE_KEY_BOOST_DATA",
	"PRODUCT_DATA",
	"QUERY_DATA",
	"VIDEO_DUARATION_DATA",
}

// NewReportCustomConfigGetV30DataTopicsFromValue returns a pointer to a valid ReportCustomConfigGetV30DataTopics
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomConfigGetV30DataTopicsFromValue(v string) (*ReportCustomConfigGetV30DataTopics, error) {
	ev := ReportCustomConfigGetV30DataTopics(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomConfigGetV30DataTopics: valid values are %v", v, AllowedReportCustomConfigGetV30DataTopicsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomConfigGetV30DataTopics) IsValid() bool {
	for _, existing := range AllowedReportCustomConfigGetV30DataTopicsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_config_get_v3.0_data_topics value
func (v ReportCustomConfigGetV30DataTopics) Ptr() *ReportCustomConfigGetV30DataTopics {
	return &v
}
