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

// ReportCustomGetV30DataTopic
type ReportCustomGetV30DataTopic string

// List of report_custom_get_v3.0_data_topic
const (
	BASIC_DATA_ReportCustomGetV30DataTopic           ReportCustomGetV30DataTopic = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomGetV30DataTopic         ReportCustomGetV30DataTopic = "BIDWORD_DATA"
	DMP_DATA_ReportCustomGetV30DataTopic             ReportCustomGetV30DataTopic = "DMP_DATA"
	MATERIAL_DATA_ReportCustomGetV30DataTopic        ReportCustomGetV30DataTopic = "MATERIAL_DATA"
	ONE_KEY_BOOST_DATA_ReportCustomGetV30DataTopic   ReportCustomGetV30DataTopic = "ONE_KEY_BOOST_DATA"
	PRODUCT_DATA_ReportCustomGetV30DataTopic         ReportCustomGetV30DataTopic = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomGetV30DataTopic           ReportCustomGetV30DataTopic = "QUERY_DATA"
	VIDEO_DUARATION_DATA_ReportCustomGetV30DataTopic ReportCustomGetV30DataTopic = "VIDEO_DUARATION_DATA"
)

// All allowed values of ReportCustomGetV30DataTopic enum
var AllowedReportCustomGetV30DataTopicEnumValues = []ReportCustomGetV30DataTopic{
	"BASIC_DATA",
	"BIDWORD_DATA",
	"DMP_DATA",
	"MATERIAL_DATA",
	"ONE_KEY_BOOST_DATA",
	"PRODUCT_DATA",
	"QUERY_DATA",
	"VIDEO_DUARATION_DATA",
}

// NewReportCustomGetV30DataTopicFromValue returns a pointer to a valid ReportCustomGetV30DataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomGetV30DataTopicFromValue(v string) (*ReportCustomGetV30DataTopic, error) {
	ev := ReportCustomGetV30DataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomGetV30DataTopic: valid values are %v", v, AllowedReportCustomGetV30DataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomGetV30DataTopic) IsValid() bool {
	for _, existing := range AllowedReportCustomGetV30DataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_get_v3.0_data_topic value
func (v ReportCustomGetV30DataTopic) Ptr() *ReportCustomGetV30DataTopic {
	return &v
}
