/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectAddV30RequestProjectListInner struct for ToolsPrivativeWordProjectAddV30RequestProjectListInner
type ToolsPrivativeWordProjectAddV30RequestProjectListInner struct {
	// 短语否定词列表
	PhraseWords []string `json:"phrase_words,omitempty"`
	// 精确否定词列表
	PreciseWords []string `json:"precise_words,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id"`
}
