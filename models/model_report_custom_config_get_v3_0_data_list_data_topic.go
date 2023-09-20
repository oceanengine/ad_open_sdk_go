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

// ReportCustomConfigGetV30DataListDataTopic
type ReportCustomConfigGetV30DataListDataTopic string

// List of report_custom_config_get_v3.0_data_list_data_topic
const (
	BASIC_DATA_ReportCustomConfigGetV30DataListDataTopic           ReportCustomConfigGetV30DataListDataTopic = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomConfigGetV30DataListDataTopic         ReportCustomConfigGetV30DataListDataTopic = "BIDWORD_DATA"
	CREATIVE_DATA_ReportCustomConfigGetV30DataListDataTopic        ReportCustomConfigGetV30DataListDataTopic = "CREATIVE_DATA"
	DMP_DATA_ReportCustomConfigGetV30DataListDataTopic             ReportCustomConfigGetV30DataListDataTopic = "DMP_DATA"
	MATERIAL_DATA_ReportCustomConfigGetV30DataListDataTopic        ReportCustomConfigGetV30DataListDataTopic = "MATERIAL_DATA"
	ONE_KEY_BOOST_DATA_ReportCustomConfigGetV30DataListDataTopic   ReportCustomConfigGetV30DataListDataTopic = "ONE_KEY_BOOST_DATA"
	PRODUCT_DATA_ReportCustomConfigGetV30DataListDataTopic         ReportCustomConfigGetV30DataListDataTopic = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomConfigGetV30DataListDataTopic           ReportCustomConfigGetV30DataListDataTopic = "QUERY_DATA"
	VIDEO_DUARATION_DATA_ReportCustomConfigGetV30DataListDataTopic ReportCustomConfigGetV30DataListDataTopic = "VIDEO_DUARATION_DATA"
)

// All allowed values of ReportCustomConfigGetV30DataListDataTopic enum
var AllowedReportCustomConfigGetV30DataListDataTopicEnumValues = []ReportCustomConfigGetV30DataListDataTopic{
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

// NewReportCustomConfigGetV30DataListDataTopicFromValue returns a pointer to a valid ReportCustomConfigGetV30DataListDataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomConfigGetV30DataListDataTopicFromValue(v string) (*ReportCustomConfigGetV30DataListDataTopic, error) {
	ev := ReportCustomConfigGetV30DataListDataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomConfigGetV30DataListDataTopic: valid values are %v", v, AllowedReportCustomConfigGetV30DataListDataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomConfigGetV30DataListDataTopic) IsValid() bool {
	for _, existing := range AllowedReportCustomConfigGetV30DataListDataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_config_get_v3.0_data_list_data_topic value
func (v ReportCustomConfigGetV30DataListDataTopic) Ptr() *ReportCustomConfigGetV30DataListDataTopic {
	return &v
}
