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

// ToolsClueContactLogListV2DataListCallDirection
type ToolsClueContactLogListV2DataListCallDirection string

// List of tools_clue_contact_log_list_v2_data_list_call_direction
const (
	CALL_IN_ToolsClueContactLogListV2DataListCallDirection  ToolsClueContactLogListV2DataListCallDirection = "CALL_IN"
	CALL_OUT_ToolsClueContactLogListV2DataListCallDirection ToolsClueContactLogListV2DataListCallDirection = "CALL_OUT"
)

// All allowed values of ToolsClueContactLogListV2DataListCallDirection enum
var AllowedToolsClueContactLogListV2DataListCallDirectionEnumValues = []ToolsClueContactLogListV2DataListCallDirection{
	"CALL_IN",
	"CALL_OUT",
}

// NewToolsClueContactLogListV2DataListCallDirectionFromValue returns a pointer to a valid ToolsClueContactLogListV2DataListCallDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueContactLogListV2DataListCallDirectionFromValue(v string) (*ToolsClueContactLogListV2DataListCallDirection, error) {
	ev := ToolsClueContactLogListV2DataListCallDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueContactLogListV2DataListCallDirection: valid values are %v", v, AllowedToolsClueContactLogListV2DataListCallDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueContactLogListV2DataListCallDirection) IsValid() bool {
	for _, existing := range AllowedToolsClueContactLogListV2DataListCallDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_contact_log_list_v2_data_list_call_direction value
func (v ToolsClueContactLogListV2DataListCallDirection) Ptr() *ToolsClueContactLogListV2DataListCallDirection {
	return &v
}
