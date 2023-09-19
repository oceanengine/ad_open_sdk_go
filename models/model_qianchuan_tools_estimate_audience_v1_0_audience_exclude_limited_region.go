/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion
type QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion int64

// List of qianchuan_tools_estimate_audience_v1.0_audience_exclude_limited_region
const (
	Enum_0_QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegionEnumValues = []QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_exclude_limited_region value
func (v QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion) Ptr() *QianchuanToolsEstimateAudienceV10AudienceExcludeLimitedRegion {
	return &v
}
