/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataAudienceElectricFenceRegion
type QianchuanAdDetailGetV10DataAudienceElectricFenceRegion int64

// List of qianchuan_ad_detail_get_v1.0_data_audience_electric_fence_region
const (
	Enum_1_QianchuanAdDetailGetV10DataAudienceElectricFenceRegion QianchuanAdDetailGetV10DataAudienceElectricFenceRegion = 1
	Enum_2_QianchuanAdDetailGetV10DataAudienceElectricFenceRegion QianchuanAdDetailGetV10DataAudienceElectricFenceRegion = 2
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceElectricFenceRegion enum
var AllowedQianchuanAdDetailGetV10DataAudienceElectricFenceRegionEnumValues = []QianchuanAdDetailGetV10DataAudienceElectricFenceRegion{
	1,
	2,
}

// NewQianchuanAdDetailGetV10DataAudienceElectricFenceRegionFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceElectricFenceRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceElectricFenceRegionFromValue(v int64) (*QianchuanAdDetailGetV10DataAudienceElectricFenceRegion, error) {
	ev := QianchuanAdDetailGetV10DataAudienceElectricFenceRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceElectricFenceRegion: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceElectricFenceRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceElectricFenceRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceElectricFenceRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_electric_fence_region value
func (v QianchuanAdDetailGetV10DataAudienceElectricFenceRegion) Ptr() *QianchuanAdDetailGetV10DataAudienceElectricFenceRegion {
	return &v
}
