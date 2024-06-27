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

// AdConvertOptimizedTargetGetV30LandingType
type AdConvertOptimizedTargetGetV30LandingType string

// List of ad_convert_optimized_target_get_v3.0_landing_type
const (
	APP_AdConvertOptimizedTargetGetV30LandingType AdConvertOptimizedTargetGetV30LandingType = "APP"
)

// All allowed values of AdConvertOptimizedTargetGetV30LandingType enum
var AllowedAdConvertOptimizedTargetGetV30LandingTypeEnumValues = []AdConvertOptimizedTargetGetV30LandingType{
	"APP",
}

// NewAdConvertOptimizedTargetGetV30LandingTypeFromValue returns a pointer to a valid AdConvertOptimizedTargetGetV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdConvertOptimizedTargetGetV30LandingTypeFromValue(v string) (*AdConvertOptimizedTargetGetV30LandingType, error) {
	ev := AdConvertOptimizedTargetGetV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdConvertOptimizedTargetGetV30LandingType: valid values are %v", v, AllowedAdConvertOptimizedTargetGetV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdConvertOptimizedTargetGetV30LandingType) IsValid() bool {
	for _, existing := range AllowedAdConvertOptimizedTargetGetV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_convert_optimized_target_get_v3.0_landing_type value
func (v AdConvertOptimizedTargetGetV30LandingType) Ptr() *AdConvertOptimizedTargetGetV30LandingType {
	return &v
}
