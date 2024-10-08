/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisSuggestionGetV2ResponseData
type ToolsDiagnosisSuggestionGetV2ResponseData struct {
	// 诊断ID
	DiagnosisId *string `json:"diagnosis_id,omitempty"`
	// 诊断ID对应的过期时间
	ExpireTimestamp *string `json:"expire_timestamp,omitempty"`
	// 实际建议
	SuggestionList []*ToolsDiagnosisSuggestionGetV2ResponseDataSuggestionListInner `json:"suggestion_list,omitempty"`
}
