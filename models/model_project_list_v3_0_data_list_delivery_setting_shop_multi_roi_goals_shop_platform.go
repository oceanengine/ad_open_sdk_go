/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform
type ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform string

// List of project_list_v3.0_data_list_delivery_setting_shop_multi_roi_goals_shop_platform
const (
	JD_ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform    ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform = "JD"
	OTHER_ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform = "OTHER"
	PDD_ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform   ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform = "PDD"
	TB_ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform    ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform = "TB"
)

// All allowed values of ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform enum
var AllowedProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatformEnumValues = []ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform{
	"JD",
	"OTHER",
	"PDD",
	"TB",
}

// NewProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatformFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatformFromValue(v string) (*ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform, error) {
	ev := ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_shop_multi_roi_goals_shop_platform value
func (v ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform) Ptr() *ProjectListV30DataListDeliverySettingShopMultiRoiGoalsShopPlatform {
	return &v
}
