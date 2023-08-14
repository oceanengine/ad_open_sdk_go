/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceInterestActionMode
type ProjectListV30DataListAudienceInterestActionMode string

// List of project_list_v3.0_data_list_audience_interest_action_mode
const (
	CUSTOM_ProjectListV30DataListAudienceInterestActionMode    ProjectListV30DataListAudienceInterestActionMode = "CUSTOM"
	RECOMMEND_ProjectListV30DataListAudienceInterestActionMode ProjectListV30DataListAudienceInterestActionMode = "RECOMMEND"
	UNLIMITED_ProjectListV30DataListAudienceInterestActionMode ProjectListV30DataListAudienceInterestActionMode = "UNLIMITED"
)

// All allowed values of ProjectListV30DataListAudienceInterestActionMode enum
var AllowedProjectListV30DataListAudienceInterestActionModeEnumValues = []ProjectListV30DataListAudienceInterestActionMode{
	"CUSTOM",
	"RECOMMEND",
	"UNLIMITED",
}

// NewProjectListV30DataListAudienceInterestActionModeFromValue returns a pointer to a valid ProjectListV30DataListAudienceInterestActionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceInterestActionModeFromValue(v string) (*ProjectListV30DataListAudienceInterestActionMode, error) {
	ev := ProjectListV30DataListAudienceInterestActionMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceInterestActionMode: valid values are %v", v, AllowedProjectListV30DataListAudienceInterestActionModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceInterestActionMode) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceInterestActionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_interest_action_mode value
func (v ProjectListV30DataListAudienceInterestActionMode) Ptr() *ProjectListV30DataListAudienceInterestActionMode {
	return &v
}
