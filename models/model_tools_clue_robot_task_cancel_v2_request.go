/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRobotTaskCancelV2Request struct for ToolsClueRobotTaskCancelV2Request
type ToolsClueRobotTaskCancelV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 机器人任务ID
	TaskId int64 `json:"task_id"`
}
