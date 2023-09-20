/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListAc
type QianchuanOrientationPackageGetV10DataListAc string

// List of qianchuan_orientation_package_get_v1.0_data_list_ac
const (
	Enum_2_G_QianchuanOrientationPackageGetV10DataListAc QianchuanOrientationPackageGetV10DataListAc = "2G"
	Enum_3_G_QianchuanOrientationPackageGetV10DataListAc QianchuanOrientationPackageGetV10DataListAc = "3G"
	Enum_4_G_QianchuanOrientationPackageGetV10DataListAc QianchuanOrientationPackageGetV10DataListAc = "4G"
	WIFI_QianchuanOrientationPackageGetV10DataListAc     QianchuanOrientationPackageGetV10DataListAc = "WIFI"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListAc enum
var AllowedQianchuanOrientationPackageGetV10DataListAcEnumValues = []QianchuanOrientationPackageGetV10DataListAc{
	"2G",
	"3G",
	"4G",
	"WIFI",
}

// NewQianchuanOrientationPackageGetV10DataListAcFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListAcFromValue(v string) (*QianchuanOrientationPackageGetV10DataListAc, error) {
	ev := QianchuanOrientationPackageGetV10DataListAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListAc: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListAc) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_ac value
func (v QianchuanOrientationPackageGetV10DataListAc) Ptr() *QianchuanOrientationPackageGetV10DataListAc {
	return &v
}
