/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdConvertOptimizedTargetGetV30ResponseDataListInner struct for AdConvertOptimizedTargetGetV30ResponseDataListInner
type AdConvertOptimizedTargetGetV30ResponseDataListInner struct {
	// 优化来源下的转化目标列表
	Converts []*AdConvertOptimizedTargetGetV30ResponseDataListInnerConvertsInner `json:"converts,omitempty"`
	// 是否禁用, true 表示已经禁用，false 表示可用
	Disabled *bool `json:"disabled,omitempty"`
}
