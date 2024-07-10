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

// StarReportCustomDataTopicReportV2DataDataMetricsType
type StarReportCustomDataTopicReportV2DataDataMetricsType string

// List of star_report_custom_data_topic_report_v2_data_data_metrics_type
const (
	FLOAT64_StarReportCustomDataTopicReportV2DataDataMetricsType StarReportCustomDataTopicReportV2DataDataMetricsType = "FLOAT64"
	INT64_StarReportCustomDataTopicReportV2DataDataMetricsType   StarReportCustomDataTopicReportV2DataDataMetricsType = "INT64"
	JSON_StarReportCustomDataTopicReportV2DataDataMetricsType    StarReportCustomDataTopicReportV2DataDataMetricsType = "JSON"
	STRING_StarReportCustomDataTopicReportV2DataDataMetricsType  StarReportCustomDataTopicReportV2DataDataMetricsType = "STRING"
)

// All allowed values of StarReportCustomDataTopicReportV2DataDataMetricsType enum
var AllowedStarReportCustomDataTopicReportV2DataDataMetricsTypeEnumValues = []StarReportCustomDataTopicReportV2DataDataMetricsType{
	"FLOAT64",
	"INT64",
	"JSON",
	"STRING",
}

// NewStarReportCustomDataTopicReportV2DataDataMetricsTypeFromValue returns a pointer to a valid StarReportCustomDataTopicReportV2DataDataMetricsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportCustomDataTopicReportV2DataDataMetricsTypeFromValue(v string) (*StarReportCustomDataTopicReportV2DataDataMetricsType, error) {
	ev := StarReportCustomDataTopicReportV2DataDataMetricsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportCustomDataTopicReportV2DataDataMetricsType: valid values are %v", v, AllowedStarReportCustomDataTopicReportV2DataDataMetricsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportCustomDataTopicReportV2DataDataMetricsType) IsValid() bool {
	for _, existing := range AllowedStarReportCustomDataTopicReportV2DataDataMetricsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_custom_data_topic_report_v2_data_data_metrics_type value
func (v StarReportCustomDataTopicReportV2DataDataMetricsType) Ptr() *StarReportCustomDataTopicReportV2DataDataMetricsType {
	return &v
}
