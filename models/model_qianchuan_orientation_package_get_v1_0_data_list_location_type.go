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

// QianchuanOrientationPackageGetV10DataListLocationType
type QianchuanOrientationPackageGetV10DataListLocationType string

// List of qianchuan_orientation_package_get_v1.0_data_list_location_type
const (
	CURRENT_QianchuanOrientationPackageGetV10DataListLocationType QianchuanOrientationPackageGetV10DataListLocationType = "CURRENT"
	HOME_QianchuanOrientationPackageGetV10DataListLocationType    QianchuanOrientationPackageGetV10DataListLocationType = "HOME"
	TRAVEL_QianchuanOrientationPackageGetV10DataListLocationType  QianchuanOrientationPackageGetV10DataListLocationType = "TRAVEL"
	ALL_QianchuanOrientationPackageGetV10DataListLocationType     QianchuanOrientationPackageGetV10DataListLocationType = "ALL"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListLocationType enum
var AllowedQianchuanOrientationPackageGetV10DataListLocationTypeEnumValues = []QianchuanOrientationPackageGetV10DataListLocationType{
	"CURRENT",
	"HOME",
	"TRAVEL",
	"ALL",
}

// NewQianchuanOrientationPackageGetV10DataListLocationTypeFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListLocationTypeFromValue(v string) (*QianchuanOrientationPackageGetV10DataListLocationType, error) {
	ev := QianchuanOrientationPackageGetV10DataListLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListLocationType: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListLocationType) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_location_type value
func (v QianchuanOrientationPackageGetV10DataListLocationType) Ptr() *QianchuanOrientationPackageGetV10DataListLocationType {
	return &v
}
