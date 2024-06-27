/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerTargetResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerTargetResult struct {
	//
	AudNum *float64 `json:"aud_num,omitempty"`
	//
	MatchPassRate *float64 `json:"match_pass_rate,omitempty"`
	//
	OtherPassRate *float64 `json:"other_pass_rate,omitempty"`
	//
	Precision *float64 `json:"precision,omitempty"`
	//
	TargetingConclusionDetail *string `json:"targeting_conclusion_detail,omitempty"`
	//
	TargetingConclusionTag *string `json:"targeting_conclusion_tag,omitempty"`
	//
	TargetingSuggestion []string `json:"targeting_suggestion,omitempty"`
}
