/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInner struct for ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInner
type ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInner struct {
	// 计划ID
	AdId *string `json:"ad_id,omitempty"`
	// 建议列表
	SceneList []*ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInner `json:"scene_list,omitempty"`
}
