/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerCtrResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerCtrResult struct {
	//
	CtrConclusionDetail *string `json:"ctr_conclusion_detail,omitempty"`
	//
	CtrConclusionTag *string `json:"ctr_conclusion_tag,omitempty"`
	//
	CtrSuggestion []string `json:"ctr_suggestion,omitempty"`
}
