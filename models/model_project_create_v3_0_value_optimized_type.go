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

// ProjectCreateV30ValueOptimizedType
type ProjectCreateV30ValueOptimizedType string

// List of project_create_v3.0_value_optimized_type
const (
	ACTION_ProjectCreateV30ValueOptimizedType ProjectCreateV30ValueOptimizedType = "ACTION"
	OFF_ProjectCreateV30ValueOptimizedType    ProjectCreateV30ValueOptimizedType = "OFF"
)

// All allowed values of ProjectCreateV30ValueOptimizedType enum
var AllowedProjectCreateV30ValueOptimizedTypeEnumValues = []ProjectCreateV30ValueOptimizedType{
	"ACTION",
	"OFF",
}

// NewProjectCreateV30ValueOptimizedTypeFromValue returns a pointer to a valid ProjectCreateV30ValueOptimizedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30ValueOptimizedTypeFromValue(v string) (*ProjectCreateV30ValueOptimizedType, error) {
	ev := ProjectCreateV30ValueOptimizedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30ValueOptimizedType: valid values are %v", v, AllowedProjectCreateV30ValueOptimizedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30ValueOptimizedType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30ValueOptimizedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_value_optimized_type value
func (v ProjectCreateV30ValueOptimizedType) Ptr() *ProjectCreateV30ValueOptimizedType {
	return &v
}
