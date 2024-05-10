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

// ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch
type ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch string

// List of project_list_v3.0_data_list_delivery_setting_budget_optimize_switch
const (
	OFF_ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch = "OFF"
	ON_ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch  ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch = "ON"
)

// All allowed values of ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch enum
var AllowedProjectListV30DataListDeliverySettingBudgetOptimizeSwitchEnumValues = []ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch{
	"OFF",
	"ON",
}

// NewProjectListV30DataListDeliverySettingBudgetOptimizeSwitchFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingBudgetOptimizeSwitchFromValue(v string) (*ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch, error) {
	ev := ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingBudgetOptimizeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingBudgetOptimizeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_budget_optimize_switch value
func (v ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch) Ptr() *ProjectListV30DataListDeliverySettingBudgetOptimizeSwitch {
	return &v
}
