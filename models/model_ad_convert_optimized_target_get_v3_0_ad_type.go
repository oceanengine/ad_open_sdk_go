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

// AdConvertOptimizedTargetGetV30AdType
type AdConvertOptimizedTargetGetV30AdType string

// List of ad_convert_optimized_target_get_v3.0_ad_type
const (
	ALL_AdConvertOptimizedTargetGetV30AdType AdConvertOptimizedTargetGetV30AdType = "ALL"
)

// All allowed values of AdConvertOptimizedTargetGetV30AdType enum
var AllowedAdConvertOptimizedTargetGetV30AdTypeEnumValues = []AdConvertOptimizedTargetGetV30AdType{
	"ALL",
}

// NewAdConvertOptimizedTargetGetV30AdTypeFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30AdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30AdTypeFromValue(v string) (*AdConvertOptimizedTargetGetV30AdType, error) {
	ev := AdConvertOptimizedTargetGetV30AdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30AdType: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30AdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30AdType) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30AdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_ad_type value
func (v AdConvertOptimizedTargetGetV30AdType) Ptr() *AdConvertOptimizedTargetGetV30AdType {
	return &v
}
