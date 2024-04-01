/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceElectricFenceRegion
type QianchuanAdUpdateV10AudienceElectricFenceRegion int64

// List of qianchuan_ad_update_v1.0_audience_electric_fence_region
const (
	Enum_1_QianchuanAdUpdateV10AudienceElectricFenceRegion QianchuanAdUpdateV10AudienceElectricFenceRegion = 1
	Enum_2_QianchuanAdUpdateV10AudienceElectricFenceRegion QianchuanAdUpdateV10AudienceElectricFenceRegion = 2
)

// All allowed values of QianchuanAdUpdateV10AudienceElectricFenceRegion enum
var AllowedQianchuanAdUpdateV10AudienceElectricFenceRegionEnumValues = []QianchuanAdUpdateV10AudienceElectricFenceRegion{
	1,
	2,
}

// NewQianchuanAdUpdateV10AudienceElectricFenceRegionFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceElectricFenceRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceElectricFenceRegionFromValue(v int64) (*QianchuanAdUpdateV10AudienceElectricFenceRegion, error) {
	ev := QianchuanAdUpdateV10AudienceElectricFenceRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceElectricFenceRegion: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceElectricFenceRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceElectricFenceRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceElectricFenceRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_electric_fence_region value
func (v QianchuanAdUpdateV10AudienceElectricFenceRegion) Ptr() *QianchuanAdUpdateV10AudienceElectricFenceRegion {
	return &v
}
