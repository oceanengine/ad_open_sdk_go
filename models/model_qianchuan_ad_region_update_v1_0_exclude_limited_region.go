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

// QianchuanAdRegionUpdateV10ExcludeLimitedRegion
type QianchuanAdRegionUpdateV10ExcludeLimitedRegion int64

// List of qianchuan_ad_region_update_v1.0_exclude_limited_region
const (
	Enum_0_QianchuanAdRegionUpdateV10ExcludeLimitedRegion QianchuanAdRegionUpdateV10ExcludeLimitedRegion = 0
	Enum_1_QianchuanAdRegionUpdateV10ExcludeLimitedRegion QianchuanAdRegionUpdateV10ExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAdRegionUpdateV10ExcludeLimitedRegion enum
var AllowedQianchuanAdRegionUpdateV10ExcludeLimitedRegionEnumValues = []QianchuanAdRegionUpdateV10ExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAdRegionUpdateV10ExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAdRegionUpdateV10ExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRegionUpdateV10ExcludeLimitedRegionFromValue(v int64) (*QianchuanAdRegionUpdateV10ExcludeLimitedRegion, error) {
	ev := QianchuanAdRegionUpdateV10ExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRegionUpdateV10ExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAdRegionUpdateV10ExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRegionUpdateV10ExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRegionUpdateV10ExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_region_update_v1.0_exclude_limited_region value
func (v QianchuanAdRegionUpdateV10ExcludeLimitedRegion) Ptr() *QianchuanAdRegionUpdateV10ExcludeLimitedRegion {
	return &v
}
