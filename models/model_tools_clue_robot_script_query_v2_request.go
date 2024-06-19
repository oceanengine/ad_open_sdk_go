/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRobotScriptQueryV2Request struct for ToolsClueRobotScriptQueryV2Request
type ToolsClueRobotScriptQueryV2Request struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 页面编号
	PageNumber *int32 `json:"page_number,omitempty"`
	// 页面大小
	PageSize *int32 `json:"page_size,omitempty"`
}
