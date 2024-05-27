/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAbTestInfoGetV2DataConclusion
type ToolsAbTestInfoGetV2DataConclusion string

// List of tools_ab_test_info_get_v2_data_conclusion
const (
	FAILED_ToolsAbTestInfoGetV2DataConclusion       ToolsAbTestInfoGetV2DataConclusion = "FAILED"
	INSUFFICIENT_ToolsAbTestInfoGetV2DataConclusion ToolsAbTestInfoGetV2DataConclusion = "INSUFFICIENT"
	NOT_START_ToolsAbTestInfoGetV2DataConclusion    ToolsAbTestInfoGetV2DataConclusion = "NOT_START"
	OBVIOUS_ToolsAbTestInfoGetV2DataConclusion      ToolsAbTestInfoGetV2DataConclusion = "OBVIOUS"
	TMP_OBVIOUS_ToolsAbTestInfoGetV2DataConclusion  ToolsAbTestInfoGetV2DataConclusion = "TMP_OBVIOUS"
)

// All allowed values of ToolsAbTestInfoGetV2DataConclusion enum
var AllowedToolsAbTestInfoGetV2DataConclusionEnumValues = []ToolsAbTestInfoGetV2DataConclusion{
	"FAILED",
	"INSUFFICIENT",
	"NOT_START",
	"OBVIOUS",
	"TMP_OBVIOUS",
}

// NewToolsAbTestInfoGetV2DataConclusionFromValue returns a pointer to a valid ToolsAbTestInfoGetV2DataConclusion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestInfoGetV2DataConclusionFromValue(v string) (*ToolsAbTestInfoGetV2DataConclusion, error) {
	ev := ToolsAbTestInfoGetV2DataConclusion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestInfoGetV2DataConclusion: valid values are %v", v, AllowedToolsAbTestInfoGetV2DataConclusionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestInfoGetV2DataConclusion) IsValid() bool {
	for _, existing := range AllowedToolsAbTestInfoGetV2DataConclusionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_info_get_v2_data_conclusion value
func (v ToolsAbTestInfoGetV2DataConclusion) Ptr() *ToolsAbTestInfoGetV2DataConclusion {
	return &v
}
