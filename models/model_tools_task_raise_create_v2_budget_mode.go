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

// ToolsTaskRaiseCreateV2BudgetMode
type ToolsTaskRaiseCreateV2BudgetMode string

// List of tools_task_raise_create_v2_budget_mode
const (
	LIMIT_ToolsTaskRaiseCreateV2BudgetMode    ToolsTaskRaiseCreateV2BudgetMode = "LIMIT"
	NO_LIMIT_ToolsTaskRaiseCreateV2BudgetMode ToolsTaskRaiseCreateV2BudgetMode = "NO_LIMIT"
)

// All allowed values of ToolsTaskRaiseCreateV2BudgetMode enum
var AllowedToolsTaskRaiseCreateV2BudgetModeEnumValues = []ToolsTaskRaiseCreateV2BudgetMode{
	"LIMIT",
	"NO_LIMIT",
}

// NewToolsTaskRaiseCreateV2BudgetModeFromValue returns a pointer to a valid ToolsTaskRaiseCreateV2BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseCreateV2BudgetModeFromValue(v string) (*ToolsTaskRaiseCreateV2BudgetMode, error) {
	ev := ToolsTaskRaiseCreateV2BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseCreateV2BudgetMode: valid values are %v", v, AllowedToolsTaskRaiseCreateV2BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseCreateV2BudgetMode) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseCreateV2BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_create_v2_budget_mode value
func (v ToolsTaskRaiseCreateV2BudgetMode) Ptr() *ToolsTaskRaiseCreateV2BudgetMode {
	return &v
}
