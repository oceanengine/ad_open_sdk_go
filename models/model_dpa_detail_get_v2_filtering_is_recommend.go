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

// DpaDetailGetV2FilteringIsRecommend
type DpaDetailGetV2FilteringIsRecommend int64

// List of dpa_detail_get_v2_filtering_is_recommend
const (
	Enum_1_DpaDetailGetV2FilteringIsRecommend DpaDetailGetV2FilteringIsRecommend = 1
	Enum_0_DpaDetailGetV2FilteringIsRecommend DpaDetailGetV2FilteringIsRecommend = 0
)

// All allowed values of DpaDetailGetV2FilteringIsRecommend enum
var AllowedDpaDetailGetV2FilteringIsRecommendEnumValues = []DpaDetailGetV2FilteringIsRecommend{
	1,
	0,
}

// NewDpaDetailGetV2FilteringIsRecommendFromValue returns a pointer to a valid DpaDetailGetV2FilteringIsRecommend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaDetailGetV2FilteringIsRecommendFromValue(v int64) (*DpaDetailGetV2FilteringIsRecommend, error) {
	ev := DpaDetailGetV2FilteringIsRecommend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaDetailGetV2FilteringIsRecommend: valid values are %v", v, AllowedDpaDetailGetV2FilteringIsRecommendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaDetailGetV2FilteringIsRecommend) IsValid() bool {
	for _, existing := range AllowedDpaDetailGetV2FilteringIsRecommendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_detail_get_v2_filtering_is_recommend value
func (v DpaDetailGetV2FilteringIsRecommend) Ptr() *DpaDetailGetV2FilteringIsRecommend {
	return &v
}
