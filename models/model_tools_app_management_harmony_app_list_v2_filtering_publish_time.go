/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementHarmonyAppListV2FilteringPublishTime
type ToolsAppManagementHarmonyAppListV2FilteringPublishTime struct {
	// 发布结束时间，格式：%Y-%m-%d
	EndTime *string `json:"end_time,omitempty"`
	// 发布起始时间，格式：%Y-%m-%d
	StartTime *string `json:"start_time,omitempty"`
}
