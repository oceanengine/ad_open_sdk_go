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

// StarReportCustomDataTopicReportV2DataDataDataTopic
type StarReportCustomDataTopicReportV2DataDataDataTopic string

// List of star_report_custom_data_topic_report_v2_data_data_data_topic
const (
	BASIC_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic             StarReportCustomDataTopicReportV2DataDataDataTopic = "BASIC_DATA"
	COMMENT_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic           StarReportCustomDataTopicReportV2DataDataDataTopic = "COMMENT_DATA"
	CONVERT_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic           StarReportCustomDataTopicReportV2DataDataDataTopic = "CONVERT_DATA"
	DY_SHOP_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic           StarReportCustomDataTopicReportV2DataDataDataTopic = "DY_SHOP_DATA"
	FLOW_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic              StarReportCustomDataTopicReportV2DataDataDataTopic = "FLOW_DATA"
	INDEX_SCORE_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic       StarReportCustomDataTopicReportV2DataDataDataTopic = "INDEX_SCORE_DATA"
	RECOMMEND_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic         StarReportCustomDataTopicReportV2DataDataDataTopic = "RECOMMEND_DATA"
	SEARCH_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic            StarReportCustomDataTopicReportV2DataDataDataTopic = "SEARCH_DATA"
	USER_DISTRIBUTION_DATA_StarReportCustomDataTopicReportV2DataDataDataTopic StarReportCustomDataTopicReportV2DataDataDataTopic = "USER_DISTRIBUTION_DATA"
)

// All allowed values of StarReportCustomDataTopicReportV2DataDataDataTopic enum
var AllowedStarReportCustomDataTopicReportV2DataDataDataTopicEnumValues = []StarReportCustomDataTopicReportV2DataDataDataTopic{
	"BASIC_DATA",
	"COMMENT_DATA",
	"CONVERT_DATA",
	"DY_SHOP_DATA",
	"FLOW_DATA",
	"INDEX_SCORE_DATA",
	"RECOMMEND_DATA",
	"SEARCH_DATA",
	"USER_DISTRIBUTION_DATA",
}

// NewStarReportCustomDataTopicReportV2DataDataDataTopicFromValue returns a pointer to a valid StarReportCustomDataTopicReportV2DataDataDataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportCustomDataTopicReportV2DataDataDataTopicFromValue(v string) (*StarReportCustomDataTopicReportV2DataDataDataTopic, error) {
	ev := StarReportCustomDataTopicReportV2DataDataDataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportCustomDataTopicReportV2DataDataDataTopic: valid values are %v", v, AllowedStarReportCustomDataTopicReportV2DataDataDataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportCustomDataTopicReportV2DataDataDataTopic) IsValid() bool {
	for _, existing := range AllowedStarReportCustomDataTopicReportV2DataDataDataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_custom_data_topic_report_v2_data_data_data_topic value
func (v StarReportCustomDataTopicReportV2DataDataDataTopic) Ptr() *StarReportCustomDataTopicReportV2DataDataDataTopic {
	return &v
}
