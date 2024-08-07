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

// QianchuanOrientationPackageGetV10DataListDistrict
type QianchuanOrientationPackageGetV10DataListDistrict string

// List of qianchuan_orientation_package_get_v1.0_data_list_district
const (
	CITY_QianchuanOrientationPackageGetV10DataListDistrict   QianchuanOrientationPackageGetV10DataListDistrict = "CITY"
	COUNTY_QianchuanOrientationPackageGetV10DataListDistrict QianchuanOrientationPackageGetV10DataListDistrict = "COUNTY"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListDistrict enum
var AllowedQianchuanOrientationPackageGetV10DataListDistrictEnumValues = []QianchuanOrientationPackageGetV10DataListDistrict{
	"CITY",
	"COUNTY",
}

// NewQianchuanOrientationPackageGetV10DataListDistrictFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListDistrictFromValue(v string) (*QianchuanOrientationPackageGetV10DataListDistrict, error) {
	ev := QianchuanOrientationPackageGetV10DataListDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListDistrict: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_district value
func (v QianchuanOrientationPackageGetV10DataListDistrict) Ptr() *QianchuanOrientationPackageGetV10DataListDistrict {
	return &v
}
