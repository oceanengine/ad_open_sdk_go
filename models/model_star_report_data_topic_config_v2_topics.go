/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportDataTopicConfigV2Topics
type StarReportDataTopicConfigV2Topics string

// List of star_report_data_topic_config_v2_topics
const (
	BASIC_DATA_StarReportDataTopicConfigV2Topics             StarReportDataTopicConfigV2Topics = "BASIC_DATA"
	COMMENT_DATA_StarReportDataTopicConfigV2Topics           StarReportDataTopicConfigV2Topics = "COMMENT_DATA"
	CONVERT_DATA_StarReportDataTopicConfigV2Topics           StarReportDataTopicConfigV2Topics = "CONVERT_DATA"
	DY_SHOP_DATA_StarReportDataTopicConfigV2Topics           StarReportDataTopicConfigV2Topics = "DY_SHOP_DATA"
	FLOW_DATA_StarReportDataTopicConfigV2Topics              StarReportDataTopicConfigV2Topics = "FLOW_DATA"
	INDEX_SCORE_DATA_StarReportDataTopicConfigV2Topics       StarReportDataTopicConfigV2Topics = "INDEX_SCORE_DATA"
	RECOMMEND_DATA_StarReportDataTopicConfigV2Topics         StarReportDataTopicConfigV2Topics = "RECOMMEND_DATA"
	SEARCH_DATA_StarReportDataTopicConfigV2Topics            StarReportDataTopicConfigV2Topics = "SEARCH_DATA"
	USER_DISTRIBUTION_DATA_StarReportDataTopicConfigV2Topics StarReportDataTopicConfigV2Topics = "USER_DISTRIBUTION_DATA"
)

// Ptr returns reference to star_report_data_topic_config_v2_topics value
func (v StarReportDataTopicConfigV2Topics) Ptr() *StarReportDataTopicConfigV2Topics {
	return &v
}
