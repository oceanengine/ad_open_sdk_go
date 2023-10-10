/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListSmartInterestAction
type QianchuanOrientationPackageGetV10DataListSmartInterestAction string

// List of qianchuan_orientation_package_get_v1.0_data_list_smart_interest_action
const (
	RECOMMEND_QianchuanOrientationPackageGetV10DataListSmartInterestAction QianchuanOrientationPackageGetV10DataListSmartInterestAction = "RECOMMEND"
	CUSTOM_QianchuanOrientationPackageGetV10DataListSmartInterestAction    QianchuanOrientationPackageGetV10DataListSmartInterestAction = "CUSTOM"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListSmartInterestAction enum
var AllowedQianchuanOrientationPackageGetV10DataListSmartInterestActionEnumValues = []QianchuanOrientationPackageGetV10DataListSmartInterestAction{
	"RECOMMEND",
	"CUSTOM",
}

// NewQianchuanOrientationPackageGetV10DataListSmartInterestActionFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListSmartInterestAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListSmartInterestActionFromValue(v string) (*QianchuanOrientationPackageGetV10DataListSmartInterestAction, error) {
	ev := QianchuanOrientationPackageGetV10DataListSmartInterestAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListSmartInterestAction: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListSmartInterestActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListSmartInterestAction) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListSmartInterestActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_smart_interest_action value
func (v QianchuanOrientationPackageGetV10DataListSmartInterestAction) Ptr() *QianchuanOrientationPackageGetV10DataListSmartInterestAction {
	return &v
}
