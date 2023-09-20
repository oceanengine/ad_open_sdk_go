/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceAutoExtendTargets
type ProjectUpdateV30AudienceAutoExtendTargets string

// List of project_update_v3.0_audience_auto_extend_targets
const (
	AD_TAG_ProjectUpdateV30AudienceAutoExtendTargets          ProjectUpdateV30AudienceAutoExtendTargets = "AD_TAG"
	AGE_ProjectUpdateV30AudienceAutoExtendTargets             ProjectUpdateV30AudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_ProjectUpdateV30AudienceAutoExtendTargets ProjectUpdateV30AudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_ProjectUpdateV30AudienceAutoExtendTargets          ProjectUpdateV30AudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_ProjectUpdateV30AudienceAutoExtendTargets ProjectUpdateV30AudienceAutoExtendTargets = "INTEREST_ACTION"
	INTEREST_TAG_ProjectUpdateV30AudienceAutoExtendTargets    ProjectUpdateV30AudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_ProjectUpdateV30AudienceAutoExtendTargets          ProjectUpdateV30AudienceAutoExtendTargets = "REGION"
)

// All allowed values of ProjectUpdateV30AudienceAutoExtendTargets enum
var AllowedProjectUpdateV30AudienceAutoExtendTargetsEnumValues = []ProjectUpdateV30AudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"INTEREST_TAG",
	"REGION",
}

// NewProjectUpdateV30AudienceAutoExtendTargetsFromValue returns a pointer to a valid ProjectUpdateV30AudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceAutoExtendTargetsFromValue(v string) (*ProjectUpdateV30AudienceAutoExtendTargets, error) {
	ev := ProjectUpdateV30AudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceAutoExtendTargets: valid values are %v", v, AllowedProjectUpdateV30AudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_auto_extend_targets value
func (v ProjectUpdateV30AudienceAutoExtendTargets) Ptr() *ProjectUpdateV30AudienceAutoExtendTargets {
	return &v
}
