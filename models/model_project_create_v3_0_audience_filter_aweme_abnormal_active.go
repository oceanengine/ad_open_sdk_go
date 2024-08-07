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

// ProjectCreateV30AudienceFilterAwemeAbnormalActive
type ProjectCreateV30AudienceFilterAwemeAbnormalActive string

// List of project_create_v3.0_audience_filter_aweme_abnormal_active
const (
	OFF_ProjectCreateV30AudienceFilterAwemeAbnormalActive ProjectCreateV30AudienceFilterAwemeAbnormalActive = "OFF"
	ON_ProjectCreateV30AudienceFilterAwemeAbnormalActive  ProjectCreateV30AudienceFilterAwemeAbnormalActive = "ON"
)

// All allowed values of ProjectCreateV30AudienceFilterAwemeAbnormalActive enum
var AllowedProjectCreateV30AudienceFilterAwemeAbnormalActiveEnumValues = []ProjectCreateV30AudienceFilterAwemeAbnormalActive{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceFilterAwemeAbnormalActiveFromValue returns a pointer to a valid ProjectCreateV30AudienceFilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceFilterAwemeAbnormalActiveFromValue(v string) (*ProjectCreateV30AudienceFilterAwemeAbnormalActive, error) {
	ev := ProjectCreateV30AudienceFilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceFilterAwemeAbnormalActive: valid values are %v", v, AllowedProjectCreateV30AudienceFilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceFilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceFilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_filter_aweme_abnormal_active value
func (v ProjectCreateV30AudienceFilterAwemeAbnormalActive) Ptr() *ProjectCreateV30AudienceFilterAwemeAbnormalActive {
	return &v
}
