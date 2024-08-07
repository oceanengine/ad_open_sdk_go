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

// QianchuanOrientationPackageGetV10DataListElectricFenceRegion
type QianchuanOrientationPackageGetV10DataListElectricFenceRegion int64

// List of qianchuan_orientation_package_get_v1.0_data_list_electric_fence_region
const (
	Enum_1_QianchuanOrientationPackageGetV10DataListElectricFenceRegion QianchuanOrientationPackageGetV10DataListElectricFenceRegion = 1
	Enum_2_QianchuanOrientationPackageGetV10DataListElectricFenceRegion QianchuanOrientationPackageGetV10DataListElectricFenceRegion = 2
)

// All allowed values of QianchuanOrientationPackageGetV10DataListElectricFenceRegion enum
var AllowedQianchuanOrientationPackageGetV10DataListElectricFenceRegionEnumValues = []QianchuanOrientationPackageGetV10DataListElectricFenceRegion{
	1,
	2,
}

// NewQianchuanOrientationPackageGetV10DataListElectricFenceRegionFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListElectricFenceRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListElectricFenceRegionFromValue(v int64) (*QianchuanOrientationPackageGetV10DataListElectricFenceRegion, error) {
	ev := QianchuanOrientationPackageGetV10DataListElectricFenceRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListElectricFenceRegion: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListElectricFenceRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListElectricFenceRegion) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListElectricFenceRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_electric_fence_region value
func (v QianchuanOrientationPackageGetV10DataListElectricFenceRegion) Ptr() *QianchuanOrientationPackageGetV10DataListElectricFenceRegion {
	return &v
}
