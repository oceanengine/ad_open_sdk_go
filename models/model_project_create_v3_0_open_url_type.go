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

// ProjectCreateV30OpenUrlType
type ProjectCreateV30OpenUrlType string

// List of project_create_v3.0_open_url_type
const (
	CUSTOM_ProjectCreateV30OpenUrlType ProjectCreateV30OpenUrlType = "CUSTOM"
	DPA_ProjectCreateV30OpenUrlType    ProjectCreateV30OpenUrlType = "DPA"
	NONE_ProjectCreateV30OpenUrlType   ProjectCreateV30OpenUrlType = "NONE"
)

// All allowed values of ProjectCreateV30OpenUrlType enum
var AllowedProjectCreateV30OpenUrlTypeEnumValues = []ProjectCreateV30OpenUrlType{
	"CUSTOM",
	"DPA",
	"NONE",
}

// NewProjectCreateV30OpenUrlTypeFromValue returns a pointer to a valid ProjectCreateV30OpenUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30OpenUrlTypeFromValue(v string) (*ProjectCreateV30OpenUrlType, error) {
	ev := ProjectCreateV30OpenUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30OpenUrlType: valid values are %v", v, AllowedProjectCreateV30OpenUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30OpenUrlType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30OpenUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_open_url_type value
func (v ProjectCreateV30OpenUrlType) Ptr() *ProjectCreateV30OpenUrlType {
	return &v
}
