/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateRebateDownloadCreateTaskV2QueryType
type FileRebateRebateDownloadCreateTaskV2QueryType string

// List of file_rebate_rebate_download_create_task_v2_query_type
const (
	MONTH_PUNISH_DETAILS_FileRebateRebateDownloadCreateTaskV2QueryType        FileRebateRebateDownloadCreateTaskV2QueryType = "month_punish_details"
	QUARTER_PUNISH_DETAILS_FileRebateRebateDownloadCreateTaskV2QueryType      FileRebateRebateDownloadCreateTaskV2QueryType = "quarter_punish_details"
	MONTH_REBATE_DETAILS_FileRebateRebateDownloadCreateTaskV2QueryType        FileRebateRebateDownloadCreateTaskV2QueryType = "month_rebate_details"
	QUARTER_REBATE_DETAILS_FileRebateRebateDownloadCreateTaskV2QueryType      FileRebateRebateDownloadCreateTaskV2QueryType = "quarter_rebate_details"
	MONTH_PERFORMANCE_DETAILS_FileRebateRebateDownloadCreateTaskV2QueryType   FileRebateRebateDownloadCreateTaskV2QueryType = "month_performance_details"
	QUARTER_PERFORMANCE_DETAILS_FileRebateRebateDownloadCreateTaskV2QueryType FileRebateRebateDownloadCreateTaskV2QueryType = "quarter_performance_details"
)

// Ptr returns reference to file_rebate_rebate_download_create_task_v2_query_type value
func (v FileRebateRebateDownloadCreateTaskV2QueryType) Ptr() *FileRebateRebateDownloadCreateTaskV2QueryType {
	return &v
}
