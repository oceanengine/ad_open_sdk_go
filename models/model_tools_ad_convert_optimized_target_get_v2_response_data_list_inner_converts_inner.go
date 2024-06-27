/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdConvertOptimizedTargetGetV2ResponseDataListInnerConvertsInner struct for ToolsAdConvertOptimizedTargetGetV2ResponseDataListInnerConvertsInner
type ToolsAdConvertOptimizedTargetGetV2ResponseDataListInnerConvertsInner struct {
	ConvertType *ToolsAdConvertOptimizedTargetGetV2DataListConvertsConvertType `json:"convert_type,omitempty"`
	//
	Disabled *bool `json:"disabled,omitempty"`
	//
	ExternalActions []*ToolsAdConvertOptimizedTargetGetV2ResponseDataListInnerConvertsInnerExternalActionsInner `json:"external_actions,omitempty"`
	//
	Page *int64 `json:"page,omitempty"`
	//
	PageSize *int64 `json:"page_size,omitempty"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
