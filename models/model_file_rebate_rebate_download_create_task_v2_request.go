/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateRebateDownloadCreateTaskV2Request struct for FileRebateRebateDownloadCreateTaskV2Request
type FileRebateRebateDownloadCreateTaskV2Request struct {
	// 代理商账户id
	AgentId int64 `json:"agent_id"`
	// 月/季度，可以传多个，以逗号分隔，如传入   1,2,3
	MonthQuarter string                                        `json:"month_quarter"`
	QueryType    FileRebateRebateDownloadCreateTaskV2QueryType `json:"query_type"`
	// 年
	Year int32 `json:"year"`
}
