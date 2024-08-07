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

// ProjectUpdateV30DeliverySettingBudgetMode
type ProjectUpdateV30DeliverySettingBudgetMode string

// List of project_update_v3.0_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_ProjectUpdateV30DeliverySettingBudgetMode      ProjectUpdateV30DeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_ProjectUpdateV30DeliverySettingBudgetMode ProjectUpdateV30DeliverySettingBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_ProjectUpdateV30DeliverySettingBudgetMode    ProjectUpdateV30DeliverySettingBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of ProjectUpdateV30DeliverySettingBudgetMode enum
var AllowedProjectUpdateV30DeliverySettingBudgetModeEnumValues = []ProjectUpdateV30DeliverySettingBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewProjectUpdateV30DeliverySettingBudgetModeFromValue returns a pointer to a valid ProjectUpdateV30DeliverySettingBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30DeliverySettingBudgetModeFromValue(v string) (*ProjectUpdateV30DeliverySettingBudgetMode, error) {
	ev := ProjectUpdateV30DeliverySettingBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30DeliverySettingBudgetMode: valid values are %v", v, AllowedProjectUpdateV30DeliverySettingBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30DeliverySettingBudgetMode) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30DeliverySettingBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_delivery_setting_budget_mode value
func (v ProjectUpdateV30DeliverySettingBudgetMode) Ptr() *ProjectUpdateV30DeliverySettingBudgetMode {
	return &v
}
