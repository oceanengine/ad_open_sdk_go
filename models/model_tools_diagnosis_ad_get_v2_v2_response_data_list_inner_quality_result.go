/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsDiagnosisAdGetV2V2ResponseDataListInnerQualityResult
type ToolsDiagnosisAdGetV2V2ResponseDataListInnerQualityResult struct {
	//
	QualityConclusion *string `json:"quality_conclusion,omitempty"`
	//
	QualityConclusionTag *string `json:"quality_conclusion_tag,omitempty"`
	//
	QualitySuggestion []string `json:"quality_suggestion,omitempty"`
}
