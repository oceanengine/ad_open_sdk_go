/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion
type QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion int64

// List of qianchuan_aweme_suggest_bid_v1.0_audience_exclude_limited_region
const (
	Enum_0_QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion = 0
	Enum_1_QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion = 1
)

// All allowed values of QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion enum
var AllowedQianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegionEnumValues = []QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion{
	0,
	1,
}

// NewQianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegionFromValue returns a pointer to a valid QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegionFromValue(v int64) (*QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion, error) {
	ev := QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion: valid values are %v", v, AllowedQianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_suggest_bid_v1.0_audience_exclude_limited_region value
func (v QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion) Ptr() *QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion {
	return &v
}
