/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceDpaCity
type ProjectCreateV30AudienceDpaCity string

// List of project_create_v3.0_audience_dpa_city
const (
	OFF_ProjectCreateV30AudienceDpaCity ProjectCreateV30AudienceDpaCity = "OFF"
	ON_ProjectCreateV30AudienceDpaCity  ProjectCreateV30AudienceDpaCity = "ON"
)

// All allowed values of ProjectCreateV30AudienceDpaCity enum
var AllowedProjectCreateV30AudienceDpaCityEnumValues = []ProjectCreateV30AudienceDpaCity{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceDpaCityFromValue returns a pointer to a valid ProjectCreateV30AudienceDpaCity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceDpaCityFromValue(v string) (*ProjectCreateV30AudienceDpaCity, error) {
	ev := ProjectCreateV30AudienceDpaCity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceDpaCity: valid values are %v", v, AllowedProjectCreateV30AudienceDpaCityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceDpaCity) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceDpaCityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_dpa_city value
func (v ProjectCreateV30AudienceDpaCity) Ptr() *ProjectCreateV30AudienceDpaCity {
	return &v
}
