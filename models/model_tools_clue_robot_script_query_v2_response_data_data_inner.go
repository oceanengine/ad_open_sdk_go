/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRobotScriptQueryV2ResponseDataDataInner struct for ToolsClueRobotScriptQueryV2ResponseDataDataInner
type ToolsClueRobotScriptQueryV2ResponseDataDataInner struct {
	// 一个号码默认外呼的次数
	CallTimes *int32 `json:"call_times,omitempty"`
	// 剧本ID
	Id *int64 `json:"id,omitempty"`
	// 剧本名称
	Name *string `json:"name,omitempty"`
	// 变量名称列表
	ScriptVariables []string `json:"script_variables,omitempty"`
}
