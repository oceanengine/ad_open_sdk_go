/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdStatExtraInfoGetV2ResponseData
type ToolsAdStatExtraInfoGetV2ResponseData struct {
	//
	AdId          *int64                                      `json:"ad_id,omitempty"`
	LearningPhase *ToolsAdStatExtraInfoGetV2DataLearningPhase `json:"learning_phase,omitempty"`
}
