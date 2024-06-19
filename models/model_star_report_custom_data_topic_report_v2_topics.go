/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarReportCustomDataTopicReportV2Topics
type StarReportCustomDataTopicReportV2Topics string

// List of star_report_custom_data_topic_report_v2_topics
const (
	BASIC_DATA_StarReportCustomDataTopicReportV2Topics             StarReportCustomDataTopicReportV2Topics = "BASIC_DATA"
	COMMENT_DATA_StarReportCustomDataTopicReportV2Topics           StarReportCustomDataTopicReportV2Topics = "COMMENT_DATA"
	CONVERT_DATA_StarReportCustomDataTopicReportV2Topics           StarReportCustomDataTopicReportV2Topics = "CONVERT_DATA"
	DY_SHOP_DATA_StarReportCustomDataTopicReportV2Topics           StarReportCustomDataTopicReportV2Topics = "DY_SHOP_DATA"
	FLOW_DATA_StarReportCustomDataTopicReportV2Topics              StarReportCustomDataTopicReportV2Topics = "FLOW_DATA"
	INDEX_SCORE_DATA_StarReportCustomDataTopicReportV2Topics       StarReportCustomDataTopicReportV2Topics = "INDEX_SCORE_DATA"
	RECOMMEND_DATA_StarReportCustomDataTopicReportV2Topics         StarReportCustomDataTopicReportV2Topics = "RECOMMEND_DATA"
	SEARCH_DATA_StarReportCustomDataTopicReportV2Topics            StarReportCustomDataTopicReportV2Topics = "SEARCH_DATA"
	USER_DISTRIBUTION_DATA_StarReportCustomDataTopicReportV2Topics StarReportCustomDataTopicReportV2Topics = "USER_DISTRIBUTION_DATA"
)

// All allowed values of StarReportCustomDataTopicReportV2Topics enum
var AllowedStarReportCustomDataTopicReportV2TopicsEnumValues = []StarReportCustomDataTopicReportV2Topics{
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

// NewStarReportCustomDataTopicReportV2TopicsFromValue returns a pointer to a valid StarReportCustomDataTopicReportV2Topics
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportCustomDataTopicReportV2TopicsFromValue(v string) (*StarReportCustomDataTopicReportV2Topics, error) {
	ev := StarReportCustomDataTopicReportV2Topics(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportCustomDataTopicReportV2Topics: valid values are %v", v, AllowedStarReportCustomDataTopicReportV2TopicsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportCustomDataTopicReportV2Topics) IsValid() bool {
	for _, existing := range AllowedStarReportCustomDataTopicReportV2TopicsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_custom_data_topic_report_v2_topics value
func (v StarReportCustomDataTopicReportV2Topics) Ptr() *StarReportCustomDataTopicReportV2Topics {
	return &v
}
