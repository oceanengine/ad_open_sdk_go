/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerCvrResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerCvrResult struct {
	//
	CvrConclusionDetail *string `json:"cvr_conclusion_detail,omitempty"`
	//
	CvrConclusionTag *string `json:"cvr_conclusion_tag,omitempty"`
	//
	CvrSuggestion []string `json:"cvr_suggestion,omitempty"`
}
