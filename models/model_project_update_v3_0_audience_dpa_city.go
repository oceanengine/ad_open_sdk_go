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

// ProjectUpdateV30AudienceDpaCity
type ProjectUpdateV30AudienceDpaCity string

// List of project_update_v3.0_audience_dpa_city
const (
	OFF_ProjectUpdateV30AudienceDpaCity ProjectUpdateV30AudienceDpaCity = "OFF"
	ON_ProjectUpdateV30AudienceDpaCity  ProjectUpdateV30AudienceDpaCity = "ON"
)

// All allowed values of ProjectUpdateV30AudienceDpaCity enum
var AllowedProjectUpdateV30AudienceDpaCityEnumValues = []ProjectUpdateV30AudienceDpaCity{
	"OFF",
	"ON",
}

// NewProjectUpdateV30AudienceDpaCityFromValue returns a pointer to a valid ProjectUpdateV30AudienceDpaCity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceDpaCityFromValue(v string) (*ProjectUpdateV30AudienceDpaCity, error) {
	ev := ProjectUpdateV30AudienceDpaCity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceDpaCity: valid values are %v", v, AllowedProjectUpdateV30AudienceDpaCityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceDpaCity) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceDpaCityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_dpa_city value
func (v ProjectUpdateV30AudienceDpaCity) Ptr() *ProjectUpdateV30AudienceDpaCity {
	return &v
}
