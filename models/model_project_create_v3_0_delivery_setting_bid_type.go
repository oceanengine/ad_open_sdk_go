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

// ProjectCreateV30DeliverySettingBidType
type ProjectCreateV30DeliverySettingBidType string

// List of project_create_v3.0_delivery_setting_bid_type
const (
	CONSERVATIVE_ProjectCreateV30DeliverySettingBidType    ProjectCreateV30DeliverySettingBidType = "CONSERVATIVE"
	CUSTOM_ProjectCreateV30DeliverySettingBidType          ProjectCreateV30DeliverySettingBidType = "CUSTOM"
	EXPLORE_UPGRADE_ProjectCreateV30DeliverySettingBidType ProjectCreateV30DeliverySettingBidType = "EXPLORE_UPGRADE"
	NO_BID_ProjectCreateV30DeliverySettingBidType          ProjectCreateV30DeliverySettingBidType = "NO_BID"
	OPTIMAL_COST_ProjectCreateV30DeliverySettingBidType    ProjectCreateV30DeliverySettingBidType = "OPTIMAL_COST"
	UPPER_CONTROL_ProjectCreateV30DeliverySettingBidType   ProjectCreateV30DeliverySettingBidType = "UPPER_CONTROL"
)

// All allowed values of ProjectCreateV30DeliverySettingBidType enum
var AllowedProjectCreateV30DeliverySettingBidTypeEnumValues = []ProjectCreateV30DeliverySettingBidType{
	"CONSERVATIVE",
	"CUSTOM",
	"EXPLORE_UPGRADE",
	"NO_BID",
	"OPTIMAL_COST",
	"UPPER_CONTROL",
}

// NewProjectCreateV30DeliverySettingBidTypeFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingBidTypeFromValue(v string) (*ProjectCreateV30DeliverySettingBidType, error) {
	ev := ProjectCreateV30DeliverySettingBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingBidType: valid values are %v", v, AllowedProjectCreateV30DeliverySettingBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingBidType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_bid_type value
func (v ProjectCreateV30DeliverySettingBidType) Ptr() *ProjectCreateV30DeliverySettingBidType {
	return &v
}
