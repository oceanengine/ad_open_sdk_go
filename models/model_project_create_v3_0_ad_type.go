/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AdType
type ProjectCreateV30AdType string

// List of project_create_v3.0_ad_type
const (
	ALL_ProjectCreateV30AdType    ProjectCreateV30AdType = "ALL"
	SEARCH_ProjectCreateV30AdType ProjectCreateV30AdType = "SEARCH"
)

// All allowed values of ProjectCreateV30AdType enum
var AllowedProjectCreateV30AdTypeEnumValues = []ProjectCreateV30AdType{
	"ALL",
	"SEARCH",
}

// NewProjectCreateV30AdTypeFromValue returns a pointer to a valid ProjectCreateV30AdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AdTypeFromValue(v string) (*ProjectCreateV30AdType, error) {
	ev := ProjectCreateV30AdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AdType: valid values are %v", v, AllowedProjectCreateV30AdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AdType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_ad_type value
func (v ProjectCreateV30AdType) Ptr() *ProjectCreateV30AdType {
	return &v
}
