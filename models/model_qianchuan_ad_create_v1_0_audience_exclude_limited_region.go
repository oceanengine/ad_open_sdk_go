/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceExcludeLimitedRegion
type QianchuanAdCreateV10AudienceExcludeLimitedRegion int64

// List of qianchuan_ad_create_v1.0_audience_exclude_limited_region
const (
	Enum_0_QianchuanAdCreateV10AudienceExcludeLimitedRegion QianchuanAdCreateV10AudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanAdCreateV10AudienceExcludeLimitedRegion QianchuanAdCreateV10AudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAdCreateV10AudienceExcludeLimitedRegion enum
var AllowedQianchuanAdCreateV10AudienceExcludeLimitedRegionEnumValues = []QianchuanAdCreateV10AudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAdCreateV10AudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanAdCreateV10AudienceExcludeLimitedRegion, error) {
	ev := QianchuanAdCreateV10AudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_exclude_limited_region value
func (v QianchuanAdCreateV10AudienceExcludeLimitedRegion) Ptr() *QianchuanAdCreateV10AudienceExcludeLimitedRegion {
	return &v
}
