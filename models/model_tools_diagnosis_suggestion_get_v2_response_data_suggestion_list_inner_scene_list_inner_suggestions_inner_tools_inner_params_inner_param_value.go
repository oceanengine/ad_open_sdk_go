/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInnerToolsInnerParamsInnerParamValue
type ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInnerSceneListInnerSuggestionsInnerToolsInnerParamsInnerParamValue struct {
	// 布尔类型参数
	BooleanParam *bool `json:"boolean_param,omitempty"`
	// 列表类型参数
	ListParam []string `json:"list_param,omitempty"`
	// 对象列表类型参数
	ObjectListParam []map[string]interface{} `json:"object_list_param,omitempty"`
	// 字符类型参数
	StringParam *string `json:"string_param,omitempty"`
}
