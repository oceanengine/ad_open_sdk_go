/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInner struct for ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInner
type ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInner struct {
	Scene *ToolsDiagnosisSuggestionGetV2DataSuggestionListSceneListScene `json:"scene,omitempty"`
	// 建议对应的工具列表
	Suggestions []*ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner `json:"suggestions,omitempty"`
}
