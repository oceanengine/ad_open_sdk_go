/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRobotTaskCreateV2Request struct for ToolsClueRobotTaskCreateV2Request
type ToolsClueRobotTaskCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 剧本ID
	ScriptId int64 `json:"script_id"`
	// 剧本变量配置信息
	ScriptVariableConfig []*ToolsClueRobotTaskCreateV2RequestScriptVariableConfigInner `json:"script_variable_config,omitempty"`
}
