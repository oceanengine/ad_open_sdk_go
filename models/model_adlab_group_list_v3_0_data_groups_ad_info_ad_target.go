/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoAdTarget
type AdlabGroupListV30DataGroupsAdInfoAdTarget string

// List of adlab_group_list_v3.0_data_groups_ad_info_ad_target
const (
	SMART_BID_CONSERVATIVE_AdlabGroupListV30DataGroupsAdInfoAdTarget AdlabGroupListV30DataGroupsAdInfoAdTarget = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_AdlabGroupListV30DataGroupsAdInfoAdTarget       AdlabGroupListV30DataGroupsAdInfoAdTarget = "SMART_BID_CUSTOM"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAdTarget enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAdTargetEnumValues = []AdlabGroupListV30DataGroupsAdInfoAdTarget{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewAdlabGroupListV30DataGroupsAdInfoAdTargetFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAdTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAdTargetFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAdTarget, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAdTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAdTarget: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAdTargetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAdTarget) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAdTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_ad_target value
func (v AdlabGroupListV30DataGroupsAdInfoAdTarget) Ptr() *AdlabGroupListV30DataGroupsAdInfoAdTarget {
	return &v
}
