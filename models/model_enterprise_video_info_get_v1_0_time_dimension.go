/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseVideoInfoGetV10TimeDimension
type EnterpriseVideoInfoGetV10TimeDimension string

// List of enterprise_video_info_get_v1.0_time_dimension
const (
	SUM_EnterpriseVideoInfoGetV10TimeDimension   EnterpriseVideoInfoGetV10TimeDimension = "SUM"
	DAILY_EnterpriseVideoInfoGetV10TimeDimension EnterpriseVideoInfoGetV10TimeDimension = "DAILY"
)

// All allowed values of EnterpriseVideoInfoGetV10TimeDimension enum
var AllowedEnterpriseVideoInfoGetV10TimeDimensionEnumValues = []EnterpriseVideoInfoGetV10TimeDimension{
	"SUM",
	"DAILY",
}

// NewEnterpriseVideoInfoGetV10TimeDimensionFromValue returns a pointer to a valid EnterpriseVideoInfoGetV10TimeDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseVideoInfoGetV10TimeDimensionFromValue(v string) (*EnterpriseVideoInfoGetV10TimeDimension, error) {
	ev := EnterpriseVideoInfoGetV10TimeDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseVideoInfoGetV10TimeDimension: valid values are %v", v, AllowedEnterpriseVideoInfoGetV10TimeDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseVideoInfoGetV10TimeDimension) IsValid() bool {
	for _, existing := range AllowedEnterpriseVideoInfoGetV10TimeDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_video_info_get_v1.0_time_dimension value
func (v EnterpriseVideoInfoGetV10TimeDimension) Ptr() *EnterpriseVideoInfoGetV10TimeDimension {
	return &v
}
