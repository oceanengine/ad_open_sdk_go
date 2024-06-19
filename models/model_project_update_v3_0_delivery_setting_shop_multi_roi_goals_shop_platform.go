/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform
type ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform string

// List of project_update_v3.0_delivery_setting_shop_multi_roi_goals_shop_platform
const (
	JD_ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform    ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "JD"
	OTHER_ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "OTHER"
	PDD_ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform   ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "PDD"
	TB_ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform    ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform = "TB"
)

// All allowed values of ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform enum
var AllowedProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatformEnumValues = []ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform{
	"JD",
	"OTHER",
	"PDD",
	"TB",
}

// NewProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatformFromValue returns a pointer to a valid ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatformFromValue(v string) (*ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform, error) {
	ev := ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform: valid values are %v", v, AllowedProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_delivery_setting_shop_multi_roi_goals_shop_platform value
func (v ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform) Ptr() *ProjectUpdateV30DeliverySettingShopMultiRoiGoalsShopPlatform {
	return &v
}
