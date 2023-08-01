/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceSuperiorPopularityType
type ProjectCreateV30AudienceSuperiorPopularityType string

// List of project_create_v3.0_audience_superior_popularity_type
const (
	GAME_ProjectCreateV30AudienceSuperiorPopularityType ProjectCreateV30AudienceSuperiorPopularityType = "GAME"
	NONE_ProjectCreateV30AudienceSuperiorPopularityType ProjectCreateV30AudienceSuperiorPopularityType = "NONE"
)

// All allowed values of ProjectCreateV30AudienceSuperiorPopularityType enum
var AllowedProjectCreateV30AudienceSuperiorPopularityTypeEnumValues = []ProjectCreateV30AudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewProjectCreateV30AudienceSuperiorPopularityTypeFromValue returns a pointer to a valid ProjectCreateV30AudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceSuperiorPopularityTypeFromValue(v string) (*ProjectCreateV30AudienceSuperiorPopularityType, error) {
	ev := ProjectCreateV30AudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceSuperiorPopularityType: valid values are %v", v, AllowedProjectCreateV30AudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_superior_popularity_type value
func (v ProjectCreateV30AudienceSuperiorPopularityType) Ptr() *ProjectCreateV30AudienceSuperiorPopularityType {
	return &v
}
