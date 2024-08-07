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

// ToolsLandingGroupCreateV2ExperimentSiteType
type ToolsLandingGroupCreateV2ExperimentSiteType string

// List of tools_landing_group_create_v2_experiment_site_type
const (
	OWN_ToolsLandingGroupCreateV2ExperimentSiteType         ToolsLandingGroupCreateV2ExperimentSiteType = "OWN"
	THIRD_PARTY_ToolsLandingGroupCreateV2ExperimentSiteType ToolsLandingGroupCreateV2ExperimentSiteType = "THIRD_PARTY"
)

// All allowed values of ToolsLandingGroupCreateV2ExperimentSiteType enum
var AllowedToolsLandingGroupCreateV2ExperimentSiteTypeEnumValues = []ToolsLandingGroupCreateV2ExperimentSiteType{
	"OWN",
	"THIRD_PARTY",
}

// NewToolsLandingGroupCreateV2ExperimentSiteTypeFromValue returns a pointer to a valid ToolsLandingGroupCreateV2ExperimentSiteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupCreateV2ExperimentSiteTypeFromValue(v string) (*ToolsLandingGroupCreateV2ExperimentSiteType, error) {
	ev := ToolsLandingGroupCreateV2ExperimentSiteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupCreateV2ExperimentSiteType: valid values are %v", v, AllowedToolsLandingGroupCreateV2ExperimentSiteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupCreateV2ExperimentSiteType) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupCreateV2ExperimentSiteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_create_v2_experiment_site_type value
func (v ToolsLandingGroupCreateV2ExperimentSiteType) Ptr() *ToolsLandingGroupCreateV2ExperimentSiteType {
	return &v
}
