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

// ProjectCreateV30AudienceConvertedTimeDuration
type ProjectCreateV30AudienceConvertedTimeDuration string

// List of project_create_v3.0_audience_converted_time_duration
const (
	NONE_ProjectCreateV30AudienceConvertedTimeDuration         ProjectCreateV30AudienceConvertedTimeDuration = "NONE"
	ONE_MONTH_ProjectCreateV30AudienceConvertedTimeDuration    ProjectCreateV30AudienceConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_ProjectCreateV30AudienceConvertedTimeDuration    ProjectCreateV30AudienceConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_ProjectCreateV30AudienceConvertedTimeDuration    ProjectCreateV30AudienceConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_ProjectCreateV30AudienceConvertedTimeDuration  ProjectCreateV30AudienceConvertedTimeDuration = "THREE_MONTH"
	TODAY_ProjectCreateV30AudienceConvertedTimeDuration        ProjectCreateV30AudienceConvertedTimeDuration = "TODAY"
	TWELVE_MONTH_ProjectCreateV30AudienceConvertedTimeDuration ProjectCreateV30AudienceConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of ProjectCreateV30AudienceConvertedTimeDuration enum
var AllowedProjectCreateV30AudienceConvertedTimeDurationEnumValues = []ProjectCreateV30AudienceConvertedTimeDuration{
	"NONE",
	"ONE_MONTH",
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"TODAY",
	"TWELVE_MONTH",
}

// NewProjectCreateV30AudienceConvertedTimeDurationFromValue returns a pointer to a valid ProjectCreateV30AudienceConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceConvertedTimeDurationFromValue(v string) (*ProjectCreateV30AudienceConvertedTimeDuration, error) {
	ev := ProjectCreateV30AudienceConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceConvertedTimeDuration: valid values are %v", v, AllowedProjectCreateV30AudienceConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_converted_time_duration value
func (v ProjectCreateV30AudienceConvertedTimeDuration) Ptr() *ProjectCreateV30AudienceConvertedTimeDuration {
	return &v
}
