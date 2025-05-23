/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateCommonDownloadCreateTaskV2Request struct for FileRebateCommonDownloadCreateTaskV2Request
type FileRebateCommonDownloadCreateTaskV2Request struct {
	// 代理商帐户ID
	AgentId int64 `json:"agent_id"`
	// 月度，可以传多个，以逗号分隔，如1,2,3
	Month *string `json:"month,omitempty"`
	// 季度
	Quarter int64 `json:"quarter"`
	// 等待最新数据参数，默认等待
	WaitLatest *bool `json:"wait_latest,omitempty"`
	// 年
	Year int64 `json:"year"`
}
