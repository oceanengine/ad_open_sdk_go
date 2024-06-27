/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertOptimizedTargetGetV2AppType
type ToolsAdConvertOptimizedTargetGetV2AppType string

// List of tools_ad_convert_optimized_target_get_v2_app_type
const (
	APP_IOS_ToolsAdConvertOptimizedTargetGetV2AppType     ToolsAdConvertOptimizedTargetGetV2AppType = "APP_IOS"
	APP_ANDROID_ToolsAdConvertOptimizedTargetGetV2AppType ToolsAdConvertOptimizedTargetGetV2AppType = "APP_ANDROID"
)

// All allowed values of ToolsAdConvertOptimizedTargetGetV2AppType enum
var AllowedToolsAdConvertOptimizedTargetGetV2AppTypeEnumValues = []ToolsAdConvertOptimizedTargetGetV2AppType{
	"APP_IOS",
	"APP_ANDROID",
}

// NewToolsAdConvertOptimizedTargetGetV2AppTypeFromValue returns a pointer to a valid ToolsAdConvertOptimizedTargetGetV2AppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertOptimizedTargetGetV2AppTypeFromValue(v string) (*ToolsAdConvertOptimizedTargetGetV2AppType, error) {
	ev := ToolsAdConvertOptimizedTargetGetV2AppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertOptimizedTargetGetV2AppType: valid values are %v", v, AllowedToolsAdConvertOptimizedTargetGetV2AppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertOptimizedTargetGetV2AppType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertOptimizedTargetGetV2AppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_optimized_target_get_v2_app_type value
func (v ToolsAdConvertOptimizedTargetGetV2AppType) Ptr() *ToolsAdConvertOptimizedTargetGetV2AppType {
	return &v
}
