/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsKeywordsBidRatioCreateV30ResponseDataConflictInner struct for ToolsKeywordsBidRatioCreateV30ResponseDataConflictInner
type ToolsKeywordsBidRatioCreateV30ResponseDataConflictInner struct {
	//
	AdvName *string `json:"advName,omitempty"`
	//
	BidCoefs  []float32                                           `json:"bidCoefs"`
	Dimension ToolsKeywordsBidRatioCreateV30DataConflictDimension `json:"dimension"`
	//
	ProjectId *int64 `json:"projectId,omitempty"`
	//
	ProjectName *string `json:"projectName,omitempty"`
	//
	Word string `json:"word"`
}
