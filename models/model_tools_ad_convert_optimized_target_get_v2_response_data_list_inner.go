/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdConvertOptimizedTargetGetV2ResponseDataListInner struct for ToolsAdConvertOptimizedTargetGetV2ResponseDataListInner
type ToolsAdConvertOptimizedTargetGetV2ResponseDataListInner struct {
	//
	Converts []*ToolsAdConvertOptimizedTargetGetV2ResponseDataListInnerConvertsInner `json:"converts,omitempty"`
	//
	Disabled         *bool                                                       `json:"disabled,omitempty"`
	MarketingPurpose *ToolsAdConvertOptimizedTargetGetV2DataListMarketingPurpose `json:"marketing_purpose,omitempty"`
}
