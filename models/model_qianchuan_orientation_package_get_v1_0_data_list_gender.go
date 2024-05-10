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

// QianchuanOrientationPackageGetV10DataListGender
type QianchuanOrientationPackageGetV10DataListGender string

// List of qianchuan_orientation_package_get_v1.0_data_list_gender
const (
	GENDER_FEMALE_QianchuanOrientationPackageGetV10DataListGender QianchuanOrientationPackageGetV10DataListGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanOrientationPackageGetV10DataListGender   QianchuanOrientationPackageGetV10DataListGender = "GENDER_MALE"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListGender enum
var AllowedQianchuanOrientationPackageGetV10DataListGenderEnumValues = []QianchuanOrientationPackageGetV10DataListGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
}

// NewQianchuanOrientationPackageGetV10DataListGenderFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListGenderFromValue(v string) (*QianchuanOrientationPackageGetV10DataListGender, error) {
	ev := QianchuanOrientationPackageGetV10DataListGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListGender: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListGender) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_gender value
func (v QianchuanOrientationPackageGetV10DataListGender) Ptr() *QianchuanOrientationPackageGetV10DataListGender {
	return &v
}
