/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30MarketingGoal
type ProjectCreateV30MarketingGoal string

// List of project_create_v3.0_marketing_goal
const (
	LIVE_ProjectCreateV30MarketingGoal            ProjectCreateV30MarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_ProjectCreateV30MarketingGoal ProjectCreateV30MarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of ProjectCreateV30MarketingGoal enum
var AllowedProjectCreateV30MarketingGoalEnumValues = []ProjectCreateV30MarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewProjectCreateV30MarketingGoalFromValue returns a pointer to a valid ProjectCreateV30MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30MarketingGoalFromValue(v string) (*ProjectCreateV30MarketingGoal, error) {
	ev := ProjectCreateV30MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30MarketingGoal: valid values are %v", v, AllowedProjectCreateV30MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30MarketingGoal) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_marketing_goal value
func (v ProjectCreateV30MarketingGoal) Ptr() *ProjectCreateV30MarketingGoal {
	return &v
}
