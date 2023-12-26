/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListMarketingGoal
type ProjectListV30DataListMarketingGoal string

// List of project_list_v3.0_data_list_marketing_goal
const (
	LIVE_ProjectListV30DataListMarketingGoal            ProjectListV30DataListMarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_ProjectListV30DataListMarketingGoal ProjectListV30DataListMarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of ProjectListV30DataListMarketingGoal enum
var AllowedProjectListV30DataListMarketingGoalEnumValues = []ProjectListV30DataListMarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewProjectListV30DataListMarketingGoalFromValue returns a pointer to a valid ProjectListV30DataListMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListMarketingGoalFromValue(v string) (*ProjectListV30DataListMarketingGoal, error) {
	ev := ProjectListV30DataListMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListMarketingGoal: valid values are %v", v, AllowedProjectListV30DataListMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListMarketingGoal) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_marketing_goal value
func (v ProjectListV30DataListMarketingGoal) Ptr() *ProjectListV30DataListMarketingGoal {
	return &v
}
