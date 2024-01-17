/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListPlatform
type QianchuanOrientationPackageGetV10DataListPlatform string

// List of qianchuan_orientation_package_get_v1.0_data_list_platform
const (
	ANDROID_QianchuanOrientationPackageGetV10DataListPlatform QianchuanOrientationPackageGetV10DataListPlatform = "ANDROID"
	IOS_QianchuanOrientationPackageGetV10DataListPlatform     QianchuanOrientationPackageGetV10DataListPlatform = "IOS"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListPlatform enum
var AllowedQianchuanOrientationPackageGetV10DataListPlatformEnumValues = []QianchuanOrientationPackageGetV10DataListPlatform{
	"ANDROID",
	"IOS",
}

// NewQianchuanOrientationPackageGetV10DataListPlatformFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListPlatformFromValue(v string) (*QianchuanOrientationPackageGetV10DataListPlatform, error) {
	ev := QianchuanOrientationPackageGetV10DataListPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListPlatform: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListPlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_platform value
func (v QianchuanOrientationPackageGetV10DataListPlatform) Ptr() *QianchuanOrientationPackageGetV10DataListPlatform {
	return &v
}
