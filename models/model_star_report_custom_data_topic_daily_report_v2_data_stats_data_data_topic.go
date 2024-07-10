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

// StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic
type StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic string

// List of star_report_custom_data_topic_daily_report_v2_data_stats_data_data_topic
const (
	BASIC_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic             StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "BASIC_DATA"
	COMMENT_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic           StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "COMMENT_DATA"
	CONVERT_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic           StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "CONVERT_DATA"
	DY_SHOP_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic           StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "DY_SHOP_DATA"
	FLOW_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic              StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "FLOW_DATA"
	INDEX_SCORE_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic       StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "INDEX_SCORE_DATA"
	RECOMMEND_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic         StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "RECOMMEND_DATA"
	SEARCH_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic            StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "SEARCH_DATA"
	USER_DISTRIBUTION_DATA_StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic = "USER_DISTRIBUTION_DATA"
)

// All allowed values of StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic enum
var AllowedStarReportCustomDataTopicDailyReportV2DataStatsDataDataTopicEnumValues = []StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic{
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

// NewStarReportCustomDataTopicDailyReportV2DataStatsDataDataTopicFromValue returns a pointer to a valid StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarReportCustomDataTopicDailyReportV2DataStatsDataDataTopicFromValue(v string) (*StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic, error) {
	ev := StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic: valid values are %v", v, AllowedStarReportCustomDataTopicDailyReportV2DataStatsDataDataTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic) IsValid() bool {
	for _, existing := range AllowedStarReportCustomDataTopicDailyReportV2DataStatsDataDataTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_report_custom_data_topic_daily_report_v2_data_stats_data_data_topic value
func (v StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic) Ptr() *StarReportCustomDataTopicDailyReportV2DataStatsDataDataTopic {
	return &v
}
