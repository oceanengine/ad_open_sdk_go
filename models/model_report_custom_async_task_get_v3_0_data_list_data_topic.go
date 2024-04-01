/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCustomAsyncTaskGetV30DataListDataTopic
type ReportCustomAsyncTaskGetV30DataListDataTopic string

// List of report_custom_async_task_get_v3.0_data_list_data_topic
const (
	BASIC_DATA_ReportCustomAsyncTaskGetV30DataListDataTopic         ReportCustomAsyncTaskGetV30DataListDataTopic = "BASIC_DATA"
	BIDWORD_DATA_ReportCustomAsyncTaskGetV30DataListDataTopic       ReportCustomAsyncTaskGetV30DataListDataTopic = "BIDWORD_DATA"
	MATERIAL_DATA_ReportCustomAsyncTaskGetV30DataListDataTopic      ReportCustomAsyncTaskGetV30DataListDataTopic = "MATERIAL_DATA"
	ONE_KEY_BOOST_DATA_ReportCustomAsyncTaskGetV30DataListDataTopic ReportCustomAsyncTaskGetV30DataListDataTopic = "ONE_KEY_BOOST_DATA"
	PRODUCT_DATA_ReportCustomAsyncTaskGetV30DataListDataTopic       ReportCustomAsyncTaskGetV30DataListDataTopic = "PRODUCT_DATA"
	QUERY_DATA_ReportCustomAsyncTaskGetV30DataListDataTopic         ReportCustomAsyncTaskGetV30DataListDataTopic = "QUERY_DATA"
)

// All allowed values of ReportCustomAsyncTaskGetV30DataListDataTopic enum
var AllowedReportCustomAsyncTaskGetV30DataListDataTopicEnumValues = []ReportCustomAsyncTaskGetV30DataListDataTopic{
	"BASIC_DATA",
	"BIDWORD_DATA",
	"MATERIAL_DATA",
	"ONE_KEY_BOOST_DATA",
	"PRODUCT_DATA",
	"QUERY_DATA",
}

// NewReportCustomAsyncTaskGetV30DataListDataTopicFromValue returns a pointer to a valid ReportCustomAsyncTaskGetV30DataListDataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCustomAsyncTaskGetV30DataListDataTopicFromValue(v string) (*ReportCustomAsyncTaskGetV30DataListDataTopic, error) {
	ev := ReportCustomAsyncTaskGetV30DataListDataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCustomAsyncTaskGetV30DataListDataTopic: valid values are %v", v, AllowedReportCustomAsyncTaskGetV30DataListDataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCustomAsyncTaskGetV30DataListDataTopic) IsValid() bool {
	for _, existing := range AllowedReportCustomAsyncTaskGetV30DataListDataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_custom_async_task_get_v3.0_data_list_data_topic value
func (v ReportCustomAsyncTaskGetV30DataListDataTopic) Ptr() *ReportCustomAsyncTaskGetV30DataListDataTopic {
	return &v
}
