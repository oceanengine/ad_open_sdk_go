/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionGetV30ResponseData
type ToolsPromotionDiagnosisSuggestionGetV30ResponseData struct {
	//
	DiagnosisId *string `json:"diagnosis_id,omitempty"`
	//
	ExpireTimestamp *string `json:"expire_timestamp,omitempty"`
	//
	SuggestionList []*ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInner `json:"suggestion_list,omitempty"`
}
