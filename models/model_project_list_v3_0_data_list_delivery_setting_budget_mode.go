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

// ProjectListV30DataListDeliverySettingBudgetMode
type ProjectListV30DataListDeliverySettingBudgetMode string

// List of project_list_v3.0_data_list_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_ProjectListV30DataListDeliverySettingBudgetMode      ProjectListV30DataListDeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_ProjectListV30DataListDeliverySettingBudgetMode ProjectListV30DataListDeliverySettingBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_ProjectListV30DataListDeliverySettingBudgetMode    ProjectListV30DataListDeliverySettingBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of ProjectListV30DataListDeliverySettingBudgetMode enum
var AllowedProjectListV30DataListDeliverySettingBudgetModeEnumValues = []ProjectListV30DataListDeliverySettingBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewProjectListV30DataListDeliverySettingBudgetModeFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingBudgetModeFromValue(v string) (*ProjectListV30DataListDeliverySettingBudgetMode, error) {
	ev := ProjectListV30DataListDeliverySettingBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingBudgetMode: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingBudgetMode) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_budget_mode value
func (v ProjectListV30DataListDeliverySettingBudgetMode) Ptr() *ProjectListV30DataListDeliverySettingBudgetMode {
	return &v
}
